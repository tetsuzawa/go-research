# -------------------------------------------------------------------------------
# Module Name:Affine projection algorithm
# Author:  m_tsutsui
# -------------------------------------------------------------------------------

# Library_Import###############
import numpy as np
# from numpy import *
import matplotlib.pyplot as plt


# Library_Import_end##################

def ap_main(mu, alpha, update_count, adf_N, d):
    """
    ap_algorithm function

    :parameter
        mu:step size,	alpha:regularisation  const
        update_count:Update count
    """

    def ap_algorithm():
        nonlocal w
        nonlocal x
        nonlocal mu
        nonlocal alpha
        nonlocal update_count
        nonlocal adf_N
        nonlocal d
        """
        nonlocal xx
        nonlocal id_matrix
        nonlocal up_inv
        """
        xx = np.dot(x.T, x)
        id_matrix = np.eye(adf_N, adf_N)
        up_inv = np.linalg.inv(alpha * id_matrix + xx)

        for _ in np.arange(1, update_count + 1, 1):
            e = d - np.dot(w.T, x)  # error vector
            w = w + mu * np.dot(x, np.dot(up_inv, e.T))

        y_opt = np.dot(w.T, x)  # filter out

        return y_opt



    # w = np.random.rand(adf_N, 1)  # initial coefficient
    # x = np.dot(w, d)  # input vector shape:(adf_N, adf_N)

    w = np.random.rand(adf_N, 1)  # initial coefficient
    w = np.array(w, dtype=np.float32)
    x = np.dot(w, d)  # input vector shape:(adf_N, adf_N)
    adf_out = ap_algorithm()

    # plot_command############################
    plt.figure(facecolor='w')
    plt.plot(d.T, "c--", alpha=0.5)
    plt.plot(adf_out.T, "r--", alpha=0.5)
    plt.grid()
    plt.legend(('desired signal', 'APA'))
    plt.title('Affine projection algorithm', fontsize=20)
    plt.show()

    plt.figure(facecolor='w')  # Back ground color_white
    plt.plot(d.T - adf_out.T, "g--")
    plt.grid()
    plt.legend(('desired signal', 'APA'))
    plt.title('Affine projection algorithm', fontsize=20)
    plt.show()


if __name__ == '__main__':
    adf_N = 64  # data size
    d = np.random.rand(1, adf_N)  # desired signal
    ap_main(mu=0.5, alpha=3, update_count=8, adf_N=adf_N, d=d)
