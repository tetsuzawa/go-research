# %%
import sys
import os.path
import numpy as np
import pandas as pd
import wave
import matplotlib
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker
from numba import jit
from matplotlib.backends.backend_pdf import PdfPages
from collections import OrderedDict
from scipy import signal
from scipy.optimize import curve_fit
from functools import wraps

# ディレクトリの絶対パスを取得
current_dir = os.path.dirname(os.path.abspath("__file__"))
# モジュールのあるパスを追加
# sys.path.append(str(current_dir) + '/../../research_tools')
# sys.path.append(str(current_dir) + '/../../sample_wav')
sys.path.append(str(current_dir) + '/research_tools')
sys.path.append(str(current_dir) + '/sample_wav')

# get_ipython().run_line_magic('matplotlib', 'inline')
# %matplotlib inline

plt.rcParams['font.family'] = 'IPAPGothic'  # 使用するフォント
plt.rcParams['xtick.direction'] = 'in'  # x軸の目盛線が内向き('in')か外向き('out')か双方向か('inout')
plt.rcParams['ytick.direction'] = 'in'  # y軸の目盛線が内向き('in')か外向き('out')か双方向か('inout')
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

fig = plt.figure(figsize=(8, 11))

# %%
try:
    import decorators
    # import plot_tool
    import plot_tools
    import adaptive_filters
    import adaptive_filters_v2
    from wave_process import WaveHandler
except ModuleNotFoundError as err:
    print(err)
    sys.path.append(str(current_dir) + '/programs')
    print(sys.path)
    print("Add path : ./programs")
    import decorators
    # import plot_tool
    import plot_tools
    import adaptive_filters
    import wave_proce
# %%
dr_01_L_wav = WaveHandler(filename="sample_wav/drone_10_11/wav/dr_level_01_L.wav")

# %%
dr_01_L_gph = plot_tools.PlotTools(y=dr_01_L_wav.data, fs=48000, fft_N=len(dr_01_L_wav.data), stft_N=256, )
# %%
dr_01_L_gph.plot_spectrogram_acf()
# %%
dr_01_silent_gph = plot_tools.PlotTools(y=dr_01_L_wav.data[:180000], fs=48000, fft_N=len(dr_01_L_wav.data[:180000]),
                                        stft_N=256, )
# %%
dr_01_silent_gph.plot_spectrogram_acf()
# %%
dr_01_silent_data = dr_01_L_wav.data[:180000]
# %%
silent_filter_Y = 1 / dr_01_silent_gph.Y
# %%
silent_filter_y = np.fft.ifft(silent_filter_Y)
dr_01_L_filterd_data = np.convolve(dr_01_L_wav.data, silent_filter_y)
# %%
dr_01_L_filterd_data = np.real(dr_01_L_filterd_data)

# %%
dr_01_L_filterd_gph = plot_tools.PlotTools(y=dr_01_L_filterd_data, fs=48000, fft_N=len(dr_01_L_filterd_data),
                                           stft_N=256, )

# %%
dr_01_L_filterd_gph.plot_spectrogram_acf()

# %%
dr_01_L_filterd_wav = WaveHandler()
# %%
dr_01_L_filterd_wav.ch = 1
dr_01_L_filterd_wav.width = 2
dr_01_L_filterd_wav.fs = 48000
# %%
dr_01_L_filterd_wav.wave_write("/Users/tetsu/personal_files/Research/sample_wav/drone_10_11/dr_level_01_L_filterd.wav",
                               dr_01_L_filterd_data)
