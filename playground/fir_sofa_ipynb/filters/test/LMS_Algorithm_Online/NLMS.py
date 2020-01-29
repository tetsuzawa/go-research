# %%
import time
import numpy as np
import matplotlib.pyplot as plt


# %%
def nlms():
    """NLMS Algorithm
    :arg:
    input u(n)
    desired d(n)
    filter param w(n)
    output y(n)
    error d(n) - y(n)
    num_taps, I, E, M
    update param mu
    
    :return: 
    """

    # mu = 0.1
    # mu = 0.7
    mu = 1.0
    # mu = 1.9

    # 1 w(0). random value. use as vector
    n_max = 32  # ideal value is infinity
    w = np.zeros(shape=n_max, dtype=np.float64, order='C')

    # 2 iterate for n = 0, 1, 2, ... n_max
    # n_max = 32  # ideal value is infinity
    # num_taps = 1024  # same as n_max
    n = np.arange(n_max)

    idx = 0
    u = np.zeros(n_max)

    # output buffer
    u_buf = []
    d_buf = []
    y_buf = []
    e_buf = []

    i = 1

    input = np.random.randn()  # input signal (random signal). use as scalar. time
    y = 0

    # while True:
    for n in range(1000 * n_max):
        try:
            print(i)
            i += 1

            # 2.0 Read/generate a new data pair
            # input = np.random.randn()  # input signal (random signal). use as scalar. time
            input = np.sin(2 * np.pi * 5000 * 1 / 48000) + 0.3 * np.sin(
                2 * np.pi * 3000 * 1 / 48000) + 0.1 * np.random.randn()

            u = np.delete(u, 0)
            u = np.append(u, input)
            un = np.dot(u, u.T)

            # if idx == n-1: idx = 0
            # else: idx += 1

            # desired signal scalar
            # d = input + 0.1 * np.random.randn()
            d = input

            # 2.1 calc filter output
            # i:0 ~ num_taps までの内積
            y = np.dot(w[:n_max].T, u)

            # 2.2 calc error
            e = d - y  # scalar

            # 2.3 parameter adaption. calc w(n+1)
            w = w + mu / (1e-8 + un) * u * e  # [n_max, 1]

            u_buf.append(input)
            d_buf.append(d)
            y_buf.append(y)
            e_buf.append(e)

            time.sleep(0.0001)

        except KeyboardInterrupt:
            break

    plt.figure(facecolor='w')  # Backgroundcolor_white
    plt.plot(u_buf, "r--", label="input u(n)")
    plt.plot(d_buf, "g--", label="desired signal d(n)")
    plt.plot(y_buf, "b--", label="output y(n)")
    plt.plot(e_buf, "y--", label="error e(n)")
    plt.grid()
    plt.legend()
    plt.title('LMS Algorithm Online')
    plt.show()


def main():
    nlms()


if __name__ == '__main__':
    main()

# %%
