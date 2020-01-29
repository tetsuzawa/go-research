#------------------------------------------------
# Module Name: Newton's Method
# Author: m_tsutsui
#------------------------------------------------

#Library_Import_______________
from numpy import*
import math, numpy as np
import matplotlib.pyplot as plt
#_____________________________

def Newton_m(Ite_C,myu):
    """
    Ite_C:Update count  , myu:step size

    """
    global w_ini

    for Ite_C in range(1,Ite_C+1,1): #Filter_Coefficients_Renewal_loop
        w_ini=w_ini*(1-2*myu)+2*myu*1/(siguma_x)*siguma_d_x

    y_opt=w_ini.T*x #ADF_Out

    return y_opt



if __name__ == '__main__':

    Data=50 #Data Number
    sita=linspace(0,Data,200)

    #________desired signal_________
    d=sin(sita) #Desired_Signal
    d_M=matrix(d).T #array->matrix

    [Data_renew,Data_renew_p]=d_M.shape #size
    ones_block=matrix(np.ones((Data_renew,1)))
    L=np.size(d,0) #size

    #_________coefficient________________
    w_ini=np.random.rand(Data_renew,Data_renew) #ADF_Initial_coefficient

    #______________Noise________________
    NG=5    #noise gain
    v=NG*(np.random.rand(Data_renew,1)-np.random.rand(Data_renew,1)) #Noise
    v_M=matrix(v)#array->matrix

    x=(d_M+v_M) #ADF_In

    #_____________Expectation________________
    E_d=np.mean(d)#E_d
    E_x=np.mean(x)#E_x

    #_____________covariance________________
    siguma_x=np.cov(x.T)#x_covariance
    siguma_x_M=matrix(siguma_x)#array->matrix
    siguma_d_x=1/L*(x-E_x*ones_block)*(d_M-E_d*ones_block).T    #E[dx]


#plot_command__________________________
    plt.figure(facecolor='w')
    plt.plot(d)
    plt.plot(array(x),'k')
    plt.plot(array(Newton_m(5,0.3)),"r--")
    plt.grid()
    plt.legend(('Desired Signal','ADF Input','ADF Output'))
    plt.title('Newtons Method')
    plt.show()
