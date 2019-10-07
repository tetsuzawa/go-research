# -------------------------------------------------------------------------------
# Module Name: LMS Algorithm (online)
# Author m_tsutsui
# -------------------------------------------------------------------------------

# Library_Import#############################
# from numpy import *
import numpy as np
import matplotlib.pyplot as plt
# Library_Import_end##########################


def lms_agm(myu, update, end_con, smp_num):
    """
    myu:step size  , update:update count

    end_con:end condition ,smp_num:sample number

    """
    w = np.random.rand(d_size, 1)  # initial coefficient
    for i in np.arange(1, update+1):
        y = np.dot(np.array(w).T, np.array(x))
        e = d[smp_num, 0]-y  # error
        w = w+myu*np.array(e)*x  # e-> array(e)
        if(abs(e) < end_con):
            break

    y_opt = np.dot(np.array(w).T, np.array(x))  # ADF out

    return y_opt


if __name__ == '__main__':

    d_size = 80
    # define time samples
    t = np.array(np.linspace(0, d_size, d_size)).T

    # Make desired value
    # d=array(sin(t)) #sine wave
    d = np.random.rand(d_size, 1)

    # Make filter input figures
    x = np.random.rand(d_size, 1)

    # ADF : Adaptive Filter
    # Define output list
    ADF_out = []
    for j in np.arange(0, d_size, 1):
        end_con = float(lms_agm(myu=0.05, update=20, end_con=0.5, smp_num=j))
        ADF_out.append(end_con)

    # _plot_command_############################
    plt.figure(facecolor='w')  # Backgroundcolor_white
    plt.plot(np.array(d), label="Desired Signal")
    plt.plot(np.array(ADF_out), "r--", label="LMS_online")
    plt.grid()
    plt.legend()
    plt.title('LMS Algorithm')
    try:
        plt.show()
    except KeyboardInterrupt:
        plt.close('all')

    # _end_#####################################
