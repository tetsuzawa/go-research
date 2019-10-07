                                                              #-------------------------------------------------------------------------------
# Name:Recursive least squares
# Author:  m_tsutsui
#-------------------------------------------------------------------------------

#Library_Import#############################
from numpy import*
import math, numpy as np
import matplotlib.pyplot as plt
#Library_Import_end##########################



def RLS(arufa,lamda,update):    #RLS Algorithm
    """
    arufa:α , lamda:forgetting factor

    update:Update Count

    """
    y_opt_buf=[]

    P_ini=1/arufa*np.eye(d_size,d_size)

    w_ini=np.zeros([d_size,1])  #initial coefficient

    for s_loop in np.arange(0,d_size,1):# d sample loop
        for i in np.arange(1,update+1,1):
            gain=(1/lamda*(matrix(P_ini)*matrix(x)))*(1+1/lamda*matrix(x).T*matrix(P_ini)*matrix(x)).I  #gain vector
            e=d[0,s_loop]-matrix(w_ini).T*matrix(x)     #error

            w_ini=matrix(w_ini)+matrix(gain)*e
            P_ini=1/lamda*(np.eye(d_size,d_size)-matrix(gain)*matrix(x).T)*matrix(P_ini)

        y_opt=float (matrix(w_ini).T*matrix(x)) #ADF out (1sample)
        y_opt_buf.append(y_opt) #ADF out vector

    return  y_opt_buf

if __name__ == '__main__':

    d_size=50   #data size

    d=np.random.rand(1,d_size)  #Desired Signal

    x=np.random.rand(d_size,1)  #filter input

#_plot_command_############################
    plt.figure(facecolor='w')
    plt.plot(array(d).T)
    plt.plot(array(RLS(0.01,0.4,5)).T,"r--")
    plt.grid()
    plt.legend(('Desired Signal','RLS'))
    plt.title('Recursive Least Squares')
    plt.show()
 #_end_#####################################
