# -------------------------------------------------------------------------------
# Module Name: LMS Algorithm (offline)
# Author m_tsutsui
# -------------------------------------------------------------------------------

# Library_Import#############################
from numpy import*
import math
import numpy as np
import matplotlib.pyplot as plt
# Library_Import_end##########################


def lms_off(myu, update, samp_n):
    """
    myu:step size  , update:update count

    smp_n:desired signal sample number

    """
    w = np.random.rand(d_size, 1)  # initial coefficient

    for n in np.arange(1, update, 1):
        # w = (np.eye(d_size, d_size)-np.array(myu)*matrix(R)) * \
        # J    matrix(w)+array(myu)*d[samp_n, 0]*matrix(x)
        w = np.dot((np.eye(d_size, d_size) - np.dot(np.array(myu), np.array(R))), np.array(w)) \
            + np.dot(np.dot(np.array(myu), d[samp_n, 0]), np.array(x))
        w_opt = w

    y_opt = np.dot(np.array(w_opt).T, np.array(x))

    return y_opt  # ADF 1 sample out


if __name__ == '__main__':

    d_size = 50  # data size

    t = np.array(np.linspace(0, d_size, d_size)).T
    # d=array(sin(t))  #Desired Signal (sine wave)

    d = np.random.rand(d_size, 1)  # Desired Signal (random signal)

    x = np.random.rand(d_size, 1)  # ADF input

    R = np.dot(np.array(x), np.array(x).T)  # E[x・x']

    ADF_out = []
    for j in np.arange(0, d_size, 1):
        ADF_buf = float(lms_off(0.03, 7, j))
        ADF_out.append(ADF_buf)  # ADF out


# _plot_command_############################
    plt.figure(facecolor='w')  # Backgroundcolor_white
    d = np.array(d)
    ADF_out = np.array(ADF_out).reshape(len(ADF_out), 1)
    plt.plot(d)
    plt.plot(ADF_out, "r--")
    plt.plot(d - ADF_out, "g--")
    plt.grid()
    plt.legend(('Desired Signal', 'LMS'))
    plt.title('LMS Algorithm(off line)')
    plt.show()
# _end_#####################################
