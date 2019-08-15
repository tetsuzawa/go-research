# -------------------------------------------------------------------------------
# Module Name:Affine projection algorithm
# Author:  m_tsutsui
# -------------------------------------------------------------------------------

# Library_Import###############
import numpy as np
from numpy import *
import matplotlib.pyplot as plt


# Library_Import_end##################

def ap_algorithm(mu, alpha, update_count):
    """
    ap_algorithm function

    :parameter
	    mu:step size,	alpha:regularisation  const
	    update_count:Update count
    """
    # nonlocal w
    # nonlocal d
    # nonlocal x
    global w
    global d
    global x

    for _ in arange(1, update_count + 1, 1):
        e = np.dot(d - w.T, x)  # error vector
        w = w + mu * np.dot(x, alpha * np.eye(adf_N, adf_N)) + np.dot(x.T, np.dot(np.linalg.inv(x), e.T))

    y_opt = np.dot(w.T, x)  # filter out

    return y_opt


if __name__ == '__main__':
    adf_N = 64  # data size

    w = np.random.rand(adf_N, 1)  # initial coefficient

    d = np.random.rand(1, adf_N)  # desired signal

    x = np.dot(w, d)  # input vector shape:(adf_N, adf_N)

    adf_out = ap_algorithm(mu=0.5, alpha=3, update_count=8)

    # plot_command############################
    plt.figure(facecolor='w')
    plt.plot(d.T, "c--")
    plt.plot(adf_out.T, "r--")
    plt.grid()
    plt.legend(('desired signal', 'APA'))
    plt.title('Affine projection algorithm', fontsize=20)

    plt.plot(d.T - adf_out.T, "g--")
    plt.grid()
    plt.legend(('desired signal', 'APA'))
    plt.title('Affine projection algorithm', fontsize=20)
    plt.show()
