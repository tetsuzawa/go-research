import numpy as np
import matplotlib.pyplot as plt


# 関数
def model(x, A, B, u):
    """
    State equation

    :type A: array(n, n) 
    """
    v = 2.0
    return A @ x + B * v  # x[1]= x=x+2vt, x[2]= y=y+vt, v=2


def true(x):
    noise = np.random.normal(loc=0.0, scale=2.0, size=(2, 1))  # loc:average, scale:variance
    return x + noise


def system(x, A, B, u):
    true_val = true(model(x, A, B, u))
    obs_val = true(true_val)
    return true_val, obs_val


def Kalman_Filter(m, V, y, A, B, u, Q, R):
    # 予測
    m_est = model(x, A, B, u)
    V_est = (A @ V) @ A.T + Q

    # 観測更新 (観測方程式無し)
    K = V_est @ np.linalg.inv(V_est + R)  # Kalman gain
    m_next = m_est + (K @ (y - m_est))  # Filtering 推定値
    V_next = (np.identity(2) - K) @ V_est

    # 観測更新 (観測方程式有り)
    """
    C = "DEFINE SOMTHNG HERE LIKE ARRAY"
    K = V_est @ (C.T @ np.linalg.inv((C @ (V_est @ C.T)) + R))
    m_next = m_est + np.dot(K, (y - np.dot(C, m_est)))
    V_next = np.dot((np.identity(2) - np.dot(K, C)), V_est)
    """

    return m_next, V_next


if __name__ == "__main__":
    A = np.identity(2)  # 状態方程式のA  単位行列(2,2)
    # B = np.ones((2, 1))  # 状態方程式のB
    B = np.array([2, 1]).reshape(2, 1)  # 状態方程式のB
    u = 2.0  # 速度
    x = np.zeros((2, 1))  # 真値
    m = np.zeros((2, 1))  # 推定値
    V = np.identity(2)  # 推定値の初期共分散行列(勝手に設定して良い)
    Q = 2 * np.identity(2)  # 入力誤差の共分散行列(今回はtrue()の中でnoiseの分散を2.0に設定したので)
    R = 2 * np.identity(2)  # 上と同じ

    # 記録用
    rec = np.empty((4, 30))

    # main loop
    for i in range(30):
        x, y = system(x, A, B, u)  # x:state value, y: observation value
        # x, y = system(x)
        m, V = Kalman_Filter(m, V, y, A, B, u, Q, R)  # m: estimated value

        rec[0, i] = x[0]
        rec[1, i] = x[1]
        rec[2, i] = m[0]
        rec[3, i] = m[1]

    # 描画
    plt.xlabel("x")
    plt.ylabel("y")
    plt.plot(rec[0, :], rec[1, :], color="blue", marker="o", label="true")
    plt.plot(rec[2, :], rec[3, :], color="red", marker="^", label="estimated")
    plt.legend()
    plt.show()
