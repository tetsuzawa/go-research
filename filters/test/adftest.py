import sys
import os.path
import numpy as np
import matplotlib.pyplot as plt

# ディレクトリの絶対パスを取得
current_dir = os.path.dirname(os.path.abspath("__file__"))
# モジュールのあるパスを追加
sys.path.append( str(current_dir) + '/../../research_tools' )


plt.rcParams['font.family'] ='IPAPGothic' #使用するフォント
plt.rcParams['xtick.direction'] = 'in' #x軸の目盛線が内向き('in')か外向き('out')か双方向か('inout')
plt.rcParams['ytick.direction'] = 'in' #y軸の目盛線が内向き('in')か外向き('out')か双方向か('inout')
plt.rcParams['xtick.top'] = True #x軸の目盛線の上側を表示
plt.rcParams['ytick.right'] = True #y軸の目盛線の右側を表示
plt.rcParams['xtick.major.width'] = 1.0 #x軸主目盛り線の線幅
plt.rcParams['ytick.major.width'] = 1.0 #y軸主目盛り線の線幅
plt.rcParams['font.size'] = 11 #フォントの大きさ
plt.rcParams['axes.linewidth'] = 1.0 #軸の線幅edge linewidth。囲みの太さ
plt.rcParams['figure.figsize'] = (7,5)
plt.rcParams['figure.dpi'] = 100 #dpiの設定
plt.rcParams['figure.subplot.hspace'] = 0.3 # 図と図の幅
plt.rcParams['figure.subplot.wspace'] = 0.3 # 図と図の幅


try:
    import decorators
    # import plot_tool
    import plot_tools
    import adaptive_filters
    import adaptive_filters_v2
    import wave_process
except ModuleNotFoundError as err:
    print(err)
    sys.path.append( str(current_dir) + '/programs' )
    print(sys.path)
    print("Add path : ./programs")
    import decorators
    # import plot_tool
    import plot_tools
    import adaptive_filters
    import wave_process









N = 2**17  # サンプル数 528244
wav= wave_process.wave_process("../../sample_wav/fuku_white_noise.wav")
wav_noise = wav.data
wav_noise = np.array(wav_noise).reshape(len(wav_noise), 1)
adf_N = 8
ADF_out = decorators.stop_watch(
                                adaptive_filters_v2.nlms_agm_on)(
                                    alpha=0.7, update_count=1, threshold=0.01, d=wav_noise, adf_N=adf_N)


