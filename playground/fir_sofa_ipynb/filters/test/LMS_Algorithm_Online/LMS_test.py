# Library_Import#############################
# from numpy import *
import numpy as np
import matplotlib.pyplot as plt
# Library_Import_end##########################

import sys
import os.path


# ディレクトリの絶対パスを取得
current_dir = os.path.dirname(os.path.abspath("__file__"))
# モジュールのあるパスを追加
sys.path.append(str(current_dir) + '/../../../programs')
print("Add path : /../../../programs")

plt.rcParams['font.family'] = 'IPAPGothic'  # 使用するフォント
# x軸の目盛線が内向き('in')か外向き('out')か双方向か('inout')
plt.rcParams['xtick.direction'] = 'in'
# y軸の目盛線が内向き('in')か外向き('out')か双方向か('inout')
plt.rcParams['ytick.direction'] = 'in'
plt.rcParams['xtick.top'] = True  # x軸の目盛線の上側を表示
plt.rcParams['ytick.right'] = True  # y軸の目盛線の右側を表示
plt.rcParams['xtick.major.width'] = 1.0  # x軸主目盛り線の線幅
plt.rcParams['ytick.major.width'] = 1.0  # y軸主目盛り線の線幅
plt.rcParams['font.size'] = 11  # フォントの大きさ
plt.rcParams['axes.linewidth'] = 1.0  # 軸の線幅edge linewidth。囲みの太さ
plt.rcParams['figure.figsize'] = (7, 5)
plt.rcParams['figure.dpi'] = 100  # dpiの設定
plt.rcParams['figure.subplot.hspace'] = 0.3  # 図と図の幅
plt.rcParams['figure.subplot.wspace'] = 0.3  # 図と図の幅

# fig = plt.figure(figsize=(8, 11))


def lms_agm(myu, update, end_con, smp_num):
    """
    myu:step size  , update:update count

    end_con:end condition ,smp_num:sample number

    """
    # global x
    # global d
    w = np.random.rand(d_size, 1)  # initial coefficient
    for i in np.arange(1, update+1):
        # find dot product of cofficients and numbers
        y = np.dot(np.array(w).T, np.array(x))
        e = d[smp_num, 0]-y  # find error
        w = w + x * np.array(e) * myu  # update w -> array(e)
        if(abs(e) < end_con):  # error threshold
            break

    y_opt = np.dot(np.array(w).T, np.array(x))  # ADF out

    return y_opt


def nlms_agm(alpha, update, end_con, smp_num):
    """
    alpha:step size 0 < alpha < 2 , update:update count

    end_con:end condition ,smp_num:sample number

    """
    # global x
    # global d

    w = np.random.rand(d_size, 1)  # initial coefficient
    for i in np.arange(1, update+1):
        # find dot product of cofficients and numbers
        y = np.dot(np.array(w).T, np.array(x))
        e = d[smp_num, 0]-y  # find error
        # update w -> array(e)
        # (+ 1e-8) : avoid dividing by 0
        w = w + alpha * np.array(e) * x / (x_norm_square + 1e-8)
        if(abs(e) < end_con):  # error threshold
            break

    y_opt = np.dot(np.array(w).T, np.array(x))  # ADF out

    return y_opt


d_size = 64

# static random function
np.random.seed(seed=10)

# define time samples
t = np.array(np.linspace(0, d_size, d_size)).T

# static random function
np.random.seed(seed=20)

# Make desired value
# d=array(sin(t)) #sine wave
d = np.random.rand(d_size, 1)

# static random function
np.random.seed(seed=30)

# Make filter input figures
x = np.random.rand(d_size, 1)

# find variance
x_var = np.var(a=x)
print(x_var)

# find norm square
x_norm_square = np.dot(x.T, x)

# ADF : Adaptive Filter
# Define output list
ADF_out = []
for j in np.arange(0, d_size, 1):
    end_con = float(lms_agm(myu=1/(x_norm_square),
                            update=20, end_con=0.5, smp_num=j))
    ADF_out.append(end_con)

ADF_out_arr = np.array(ADF_out)
ADF_out_nd = ADF_out_arr.reshape(len(ADF_out_arr), 1)

# _plot_command_############################
plt.figure(facecolor='w')  # Backgroundcolor_white
plt.plot(d, label="Desired Signal")
plt.plot(ADF_out_nd, "r--", label="LMS_online")
plt.plot(d-ADF_out_nd, "g--", label="LMS_online_filterd")
plt.grid()
plt.legend()
plt.title('LMS Algorithm Online')
try:
    plt.show()
except KeyboardInterrupt:
    plt.close('all')


# ADF : Adaptive Filter
# Define output list
NADF_out = []
for j in np.arange(0, d_size, 1):
    nend_con = float(nlms_agm(alpha=1, update=20, end_con=0.5, smp_num=j))
    NADF_out.append(nend_con)


NADF_out_arr = np.array(NADF_out)
NADF_out_nd = NADF_out_arr.reshape(len(NADF_out_arr), 1)
# _plot_command_############################
plt.figure(facecolor='w')  # Backgroundcolor_white
plt.plot(d, label="Desired Signal")
plt.plot(NADF_out_nd, "r--", label="NLMS_online")
plt.plot(d-NADF_out_nd, "g--", label="NLMS_online_filterd")
plt.grid()
plt.legend()
plt.title('NLMS Algorithm Online')
try:
    plt.show()
except KeyboardInterrupt:
    plt.close('all')
