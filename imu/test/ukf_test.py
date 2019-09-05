import numpy as np
import matplotlib.pyplot as plt




# UKF Parameters
ukf_lambda = 10.0  # UKFのλパラメータ
ukf_kappa = 0.1  # UKFのκパラメータ
ukf_alpha2 = (2.0 + ukf_lambda) / (2.0 + ukf_kappa)  # UKFのα^2パラメータ
ukf_w0_m = ukf_lambda / (2.0 + ukf_lambda)  # UKFの重みパラメータ
ukf_w0_c = ukf_w0_m + (1.0 - ukf_alpha2 + 2.0)  # UKFの重みパラメータ
ukf_wi = 1.0 / (2.0 * (2.0 + ukf_lambda))  # UKFの重みパラメータ
ukf_wm = np.zeros([5, 1])  # UKFの重みパラメータw_m
ukf_wc = np.zeros([5, 1])  # UKFの重みパラメータw_c
ukf_gamma = np.sqrt(3.0 + ukf_lambda)  # γ = √(n + λ)

# UKFの重みパラメータの設定
ukf_wm[0] = ukf_w0_m
ukf_wc[0] = ukf_w0_c
for i in range(1, 5):
    ukf_wm[i] = ukf_wi
    ukf_wc[i] = ukf_wi


def gen_sigma_point(mu, sigma):
    """
    Function to generate Matrix chi which is the set of the sigma point Vector.

    arguments :
    mu    : Vector of the mean values.
    sigma : Correlation Matrix.

    return :
    chi   : Matrix chi.
    """
    n = len(mu)  # 次元
    chi = np.zeros([n, 2 * n + 1])  # χ行列
    root_sigma = np.linalg.cholesky(sigma)  # √Σの計算

    chi[:, 0] = mu[:, 0]
    chi[:, 1:1 + n] = mu + ukf_gamma * root_sigma
    chi[:, 1 + n:2 * n + 1] = mu - ukf_gamma * root_sigma

    return chi


def gen_correlation_mat(mu, chi, R):
    """
    Function to describe the change of the correlation matrix.

    arguments :
    mu    : Vector of the mean values.
    chi   : Destribution of the sigma point.
    R     : Correlation Matrix of the observation noise.

    return :
    sigma_bar : Correlation Matrix after the change.
    """
    sigma_bar = R

    for i in range(chi.shape[1]):
        x = chi[:, i].reshape((chi.shape[0], 1)) - mu
        sigma_bar = sigma_bar + ukf_wc[i] * (np.dot(x, x.T))

    return sigma_bar


def UKF(mu, sigma, u, z, state_func, obs_func, R, Q, dt):
    """
    Unscented Kalman Filter Program.

    arguments :
    mu         : Vector which is the set of state values.
    sigma      : Correlation Matrix of the state values.
    u          : Control inputs.Vector/Float.
    z          : Observed values.
    state_func : Function which describe how the state changes.
    obs_func   : Function to get observed values from state values.
    R          : Correlation Matrix of the noise of the state change.
    Q          : Correlation Matrix of the observation noise.
    dt         : Time span of the integration.

    returns :
    mu    : Vector which is the set of state values.
    sigma : Correlation Matrix of the state values.
    ---------------
    """
    # Σ点生成
    chi = gen_sigma_point(mu, sigma)  # χ

    # Σ点の遷移
    chi_star = state_func(u, chi, dt)

    # 事前分布取得
    mu_bar = np.dot(chi_star, ukf_wm)  # 平均の事前分布
    sigma_bar = gen_correlation_mat(mu_bar, chi_star, R)  # 共分散行列の事前分布

    # Σ点の事前分布
    chi_bar = gen_sigma_point(mu_bar, sigma_bar)

    # 観測の予測分布
    z_bar = obs_func(chi_bar)

    # 観測の予測値
    z_hat = np.dot(z_bar, ukf_wm)

    # 観測の予測共分散行列
    S = gen_correlation_mat(z_hat, z_bar, Q)

    # Σ_xzの計算
    sigma_xz = np.zeros([mu.shape[0], z.shape[0]])
    for i in range(chi_bar.shape[1]):
        x_ = chi_bar[:, i] - mu_bar[:, 0]
        z_ = z_bar[:, i] - z_hat[:, 0]
        x_ = x_.reshape((x_.shape[0], 1))
        z_ = z_.reshape((z_.shape[0], 1))
        sigma_xz = sigma_xz + ukf_wc[i] * (np.dot(x_, z_.T))

    # カルマンゲインの計算
    K = np.dot(sigma_xz, np.linalg.inv(S))

    # 戻り値の計算
    mu = mu_bar + np.dot(K, (z - z_hat))
    sigma = sigma_bar - np.dot(np.dot(K, S), K.T)

    return mu, sigma


def state_func(u, chi, dt):
    L = 100
    omega = np.pi / 10
    return chi + np.array([[-L * omega * np.sin(omega * u)],
                           [L * omega * np.cos(omega * u)]])


def obs_func(chi_bar):
    obs = np.zeros_like(chi_bar)
    for i in range(chi_bar.shape[1]):
        obs[0][i] = np.sqrt(chi_bar[0, i] ** 2 + chi_bar[1, i] ** 2)
        obs[1][i] = np.arctan(chi_bar[1, i] / chi_bar[0, i])

    return obs




# 関数
def state_eq(x, L, omega, t):
    x_new = x + np.array([-L * omega * np.sin(omega * t),
                          L * omega * np.cos(omega * t)]).reshape([2, 1])
    return x_new


def obs_eq(x, obs_noise_y1, obs_noise_y2):
    x1 = x[0]
    x2 = x[1]
    y = np.array([np.sqrt(x1 ** 2 + x2 ** 2),
                  np.arctan(x2 / x1)]).reshape([2, 1])
    noise = np.array([np.random.normal(0, obs_noise_y1),
                      np.random.normal(0, obs_noise_y2)]).reshape([2, 1])
    return y + noise


def obs_eq_noiseless(x):
    x1 = x[0]
    x2 = x[1]
    y = np.array([np.sqrt(x1 ** 2 + x2 ** 2),
                  np.arctan(x2 / x1)]).reshape([2, 1])
    return y


def true(x, input_noise):
    noise = np.random.normal(0, input_noise, (2, 1))
    return x + noise


def state_jacobian():
    jacobian = np.identity(2)
    return jacobian


def obs_jacobian(x):
    jacobian = np.empty((2, 2))
    jacobian[0][0] = x[0] / np.sqrt(x[0] ** 2 + x[1] ** 2)
    jacobian[0][1] = x[1] / np.sqrt(x[0] ** 2 + x[1] ** 2)
    jacobian[1][0] = -x[1] / (x[0] ** 2 + x[1] ** 2)
    jacobian[1][1] = x[0] / (x[0] ** 2 + x[1] ** 2)

    return jacobian


def system(x, L, omega, t, input_noise, obs_noise_y1, obs_noise_y2):
    true_state = true(state_eq(x, L, omega, t), input_noise)
    obs = obs_eq(true_state, obs_noise_y1, obs_noise_y2)
    return true_state, obs


def EKF(m, V, y, Q, R, L, omega, t):
    # 予測ステップ
    m_est = state_eq(m, L, omega, t)
    A = state_jacobian()
    V_est = np.dot(np.dot(A, V), A.transpose()) + Q

    # 観測更新ステップ
    C = obs_jacobian(m_est)
    temp = np.dot(np.dot(C, V_est), C.transpose()) + R
    K = np.dot(np.dot(V_est, C.transpose()), np.linalg.inv(temp))
    m_next = m_est + np.dot(K, (y - obs_eq_noiseless(m_est)))
    V_next = np.dot(np.identity(V_est.shape[0]) - np.dot(K, C), V_est)

    return m_next, V_next


if __name__ == "__main__":
    x = np.array([100, 0]).reshape([2, 1])
    L = 100
    omega = np.pi / 10
    input_noise = 1.0 ** 2
    obs_noise_y1 = 10.0 ** 2
    obs_noise_y2 = (5.0 * np.pi / 180) ** 2
    m = np.array([100, 0]).reshape([2, 1])
    t = 0.0
    dt = 1.0
    V = np.identity(2) * 1.0 ** 2
    Q = np.identity(2) * input_noise
    R = np.array([[obs_noise_y1, 0],
                  [0, obs_noise_y2]])

    # 記録用
    rec = np.empty([4, 21])

    for i in range(21):
        rec[0, i] = x[0]
        rec[1, i] = x[1]
        rec[2, i] = m[0]
        rec[3, i] = m[1]

        x, y = system(x, L, omega, t, input_noise, obs_noise_y1, obs_noise_y2)
        m, V = EKF(m, V, y, Q, R, L, omega, t)

        t += dt

    plt.plot(rec[0, :], rec[1, :], color="blue", marker="o", label="true")
    plt.plot(rec[2, :], rec[3, :], color="red", marker="^", label="estimated")
    plt.legend()
    plt.show()
