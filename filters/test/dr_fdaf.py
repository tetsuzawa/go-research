import sys

import numpy as np
import matplotlib

matplotlib.use('Agg')
import matplotlib.pyplot as plt

sys.path.append('/Users/tetsu/personal_files/Research')
sys.path.append('/Users/tetsu/personal_files/Research/research_tools')
sys.path.append('/Users/tetsu/personal_files/Research/sample_wav')

from research_tools.wave_handler_multi_ch import WaveHandler

input_filename = sys.argv[1]
input_name_list = input_filename.split(".")
mu = 0.018
L = 64


def fdaf(data):
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

    idx = 0
    is_data_end = False

    try:
        # for k in range(k_max):
        while True:
            i += 1
            j = 1

            # 2.0 Initialize phi = 0s

            # 2.1 Iterate for i = 0, 1, 2, 3, ..., L-1 (k is the block index)
            for j in range(L):
                # print(i * L + j)
                j += 1

                # 2.1.0 Read/generate a new data pair
                # input = np.random.randn()  # input signal.  use as scalar. time
                # input = np.sin(2 * np.pi * 5000 * 1 / 48000) + 0.3 * np.sin(
                #     2 * np.pi * 3000 * 1 / 48000) + 0.1 * np.random.randn()
                if is_data_end:
                    print("\nAdaptation completed!!")
                    print("Please wait for plotting...")
                    break
                input = data[idx]
                if idx == len(data) - 1:
                    is_data_end = True
                idx += 1

                u = np.delete(u, 0)
                u = np.append(u, input)

            if is_data_end:
                break

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

            # ----------- progress bar -----------
            print(f'{(idx + 1) * 100 / len(data):3.0f}%\r', end='')


    except KeyboardInterrupt:
        print("Ctrl-C detected!!")

    plt.figure(facecolor='w')  # Backgroundcolor_white
    plt.plot(u_buf, "r--", label="input u(n)")
    plt.plot(d_buf, "g--", label="desired signal d(n)")
    plt.plot(y_buf, "b-", label="output y(n)")
    plt.plot(e_buf, "y-", label="error e(n)")
    plt.grid()
    plt.legend()
    plt.title('LMS Algorithm Online')
    plt.show()

    img_out_dir = "/Users/tetsu/personal_files/Research/filters/test/FDAF_img/"
    img_out_name = input_name_list[0] + f"_mu-{mu}_L-{L}.png"
    plt.savefig(img_out_dir + img_out_name)
    print("\nfilterd data plot is saved at: ", img_out_dir + img_out_name, "\n")

    return e_buf


def main():
    # wav = WaveHandler("/Users/tetsu/personal_files/Research/sample_wav/drone_10_11/wav/dr_level_01_L.wav")
    wav = WaveHandler(input_filename)
    data = wav.data
    filtered_data = fdaf(data)

    filtered_wav = WaveHandler()
    filtered_wav.ch = 1
    filtered_wav.width = 2
    filtered_wav.fs = 48000

    wav_out_dir = "/Users/tetsu/personal_files/Research/filters/test/FDAF_wav/"
    wav_out_name = input_name_list[0] + f"_mu-{mu}_L-{L}.wav"

    filtered_wav.wave_write(
        filename=wav_out_dir + wav_out_name,
        data_array=filtered_data)
    print("\nfilterd data is saved at: ", wav_out_dir + wav_out_name, "\n")


if __name__ == '__main__':
    main()
