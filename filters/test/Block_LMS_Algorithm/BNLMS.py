import time
import numpy as np
import matplotlib.pyplot as plt


def bnlms():
    """Block Normalized LMS Algorithm
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

    mu = 0.00001
    # mu = 0.100
    # mu = 1.0
    L = 512

    # 1 w(0). random value. use as vector
    k_max = L  # ideal value is infinity
    w = np.zeros(shape=L, dtype=np.float64, order='C')


    # output buffer
    u_buf = []
    d_buf = []
    y_buf = []
    e_buf = []

    i = 0

    input = np.random.randn()  # input signal (random signal). use as scalar. time
    y = 0


    # 2 Iterate for k = 0, 1, 2, 3, ..., k_max (k is the block index)

    # 2.0 Initialize phi = 0s
    phi = np.zeros(shape=L, dtype=np.float64, order='C')

    u = np.zeros(L)
    print(u)

    try:
        # for k in range(k_max):
        while True:
            i += 1
            j = 1


            # 2.1 Iterate for i = 0, 1, 2, 3, ..., L-1 (k is the block index)
            for j in range(L):
                print(i * L + j)
                j += 1

                # 2.1.0 Read/generate a new data pair
                # input = np.random.randn()  # input signal.  use as scalar. time
                input = np.sin(2 * np.pi * 5000 * 1 / 48000) + 0.3 * np.sin(
                    2 * np.pi * 3000 * 1 / 48000) + 0.1 * np.random.randn()

                u = np.delete(u, 0)
                u = np.append(u, input)
                u_norm = np.dot(u.T, u)

                # desired signal scalar
                # d = input + 0.1 * np.random.randn()
                d = input

                # 2.1.1 calc filter output
                # i:0 ~ num_taps までの内積
                y = np.dot(w.T, u)

                # 2.2 calc error
                e = d - y  # scalar

                # 2.3 parameter adaption. calc w(n+1)
                phi = phi + mu * u * e / (1e-8 + u_norm)  # [n_max, 0]

                u_buf.append(input)
                d_buf.append(d)
                y_buf.append(y)
                e_buf.append(e)

                # time.sleep(0.001)

            w = w + phi

    except KeyboardInterrupt:
        pass


    plt.figure(facecolor='w')  # Backgroundcolor_white
    plt.plot(u_buf, "r--", label="input u(n)")
    plt.plot(d_buf, "g--", label="desired signal d(n)")
    plt.plot(y_buf, "b-", label="output y(n)")
    plt.plot(e_buf, "y-", label="error e(n)")
    plt.grid()
    plt.legend()
    plt.title('LMS Algorithm Online')
    plt.show()


def main():
    bnlms()


if __name__ == '__main__':
    main()
