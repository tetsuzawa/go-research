# -------------------------------------------------------------------------------
# Module Name:Affine projection algorithm
# Author:  m_tsutsui
# -------------------------------------------------------------------------------

# Library_Import###############
import numpy as np
# from numpy import *
import matplotlib.pyplot as plt


# Library_Import_end##################

def APA(myu, arufa, UC):
    """
    APA function

	myu:step size,	arufa:regularisation  const

	UC:Update count
    """
    for i in np.arange(1, UC + 1, 1):
        global w_ini
        e = np.matrix(d) - np.matrix(w_ini).T * np.matrix(X)  # error vector
        w_ini = w_ini + myu * np.matrix(X) * (arufa * np.eye(d_size, d_size) + np.matrix(X).T * np.matrix(X)).I * np.matrix(e).T

    y_opt = np.matrix(w_ini).T * np.matrix(X)  # filter out

    return y_opt


if __name__ == '__main__':
    d_size = 80  # data size

    w_ini = np.random.rand(d_size, 1)  # initial coefficient

    d = np.random.rand(1, d_size)  # desired signal

    X = np.matrix(w_ini) * np.matrix(d)  # input vector

    APA_out = APA(0.5, 3, 8)

    # plot_command############################
    plt.figure(facecolor='w')
    plt.plot(d.T)
    plt.plot(APA_out.T, "r--")
    plt.grid()
    plt.legend(('desired signal', 'APA'))
    plt.title('Affine projection algorithm', fontsize=20)
    plt.show()
