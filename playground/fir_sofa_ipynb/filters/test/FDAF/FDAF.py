import numpy as np
import matplotlib.pyplot as plt


def fdaf():
    """Block LMS Algorithm
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

    mu = 0.000932
    L = 64
    # mu = 0.100

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

    u = np.zeros(shape=2 * L, dtype=np.float64, order='C')
    e = np.zeros(shape=2 * L, dtype=np.float64, order='C')
    phi = np.zeros(shape=L, dtype=np.float64, order='C')
    zeros = np.zeros(shape=L, dtype=np.float64, order='C')

    # 2 Iterate for k = 0, 1, 2, 3, ..., k_max (k is the block index)

    try:
        # for k in range(k_max):
        while True:
            i += 1
            j = 1

            # 2.0 Initialize phi = 0s

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

            # 1 compute the output of the filter for the block kM, ..., KM + M -1
            Y = np.fft.fft(np.concatenate([w[:L], zeros])) * np.fft.fft(u)  # W * U
            y = np.fft.ifft(Y)[L:]

            d = u
            e = d[L:] - y

            # 2 compute the correlation vector
            aux1 = np.concatenate([zeros, e])
            aux2 = np.fft.fft(u)
            EU = np.fft.fft(aux1) * np.conj(aux2)
            phi = np.fft.ifft(EU)[:L]

            # 3 update the parameters of the filter
            W = np.fft.fft(np.concatenate([w[:L], zeros])) + mu * np.fft.fft(np.concatenate([phi, zeros]))
            w = np.fft.ifft(W)

            u_buf.extend(u[L:])
            d_buf.extend(u[L:])
            y_buf.extend(y)
            e_buf.extend(e)

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
    fdaf()


if __name__ == '__main__':
    main()
