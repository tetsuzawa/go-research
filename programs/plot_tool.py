# -*- coding: utf-8 -*-
# python3.7 VS_plot.DXB.py [FILE_NAME] [FFT_LENGTH]
#
import sys
import os.path
import math
import numpy as np
import pandas as pd
import wave
import struct
import string
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker
from matplotlib.backends.backend_pdf import PdfPages
from collections import OrderedDict
from scipy import signal
import time


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

fig = plt.figure(figsize=(8, 11))
# plt.gca().xaxis.set_major_formatter(plt.FormatStrFormatter('%.3f'))#y軸小数点以下3桁表示
# plt.gca().yaxis.set_major_formatter(plt.FormatStrFormatter('%.3f'))#y軸小数点以下3桁表示
# plt.gca().xaxis.get_major_formatter().set_useOffset(False)

# plt.add_axes([left,bottom,width,height],zorder=0)


def define_window_function(name, N, kaiser_para=5):
    if name == None:
        return 1
    elif name == "hamming":
        return np.hamming(M=N)
    elif name == "hanning":
        return np.hanning(M=N)
    elif name == "bartlett":
        return np.bartlett(M=N)
    elif name == "blackman":
        return np.blackman(M=N)
    elif name == "kaiser":
        return np.kaiser(N=N, beta=kaiser_para)


def plot_3charts(N, y, fs=48000, start_sec=0, window_func_name="hamming"):
    # Period
    dt = 1/fs
    # Define start sec
    start_pos = int(start_sec/dt)
    # Redefine y
    y = y[start_pos: N+start_pos]
    # Window function
    hamming_win = define_window_function(name=window_func_name, N=N)
    # Fourier transform
    Y = np.fft.fft(hamming_win * y)
    # Find a list of frequencies
    freqList = np.fft.fftfreq(N, d=dt)
    # Find the time for y
    t = np.arange(start_pos*dt, (N+start_pos)*dt, dt)

    # Complement 0 or less to display decibels
    y_abs = np.array(np.abs(y))
    u_0_list = np.where(y_abs <= 0)
    for u_0 in u_0_list:
        y_abs[u_0] = (y_abs[u_0-1] + y_abs[u_0+1]) / 2

    # y decivel desplay
    y_db = 20.0*np.log10(y_abs)

    # amplitudeSpectrum = [np.sqrt(c.real ** 2  + c.imag ** 2 ) for c in Y]
    # phaseSpectrum     = [np.arctan2(np.float64(c.imag),np.float64(c.real)) for c in Y]
    # Adjust the amplitude to the original signal.
    amplitudeSpectrum = np.abs(Y) / N * 2 
    amplitudeSpectrum[0] = amplitudeSpectrum[0] / 2 
    # amplitudeSpectrum = np.abs(Y) / np.max(amplitudeSpectrum)
    phaseSpectrum = np.rad2deg(np.angle(Y))
    decibelSpectrum = 20.0*np.log10(amplitudeSpectrum / np.max(amplitudeSpectrum))

    fig = plt.figure(figsize=(11, 8))

    '''
    ax1 = fig.add_subplot(311)
    ax1.plot(y)
    ax1.axis([0,N,np.amin(y),np.amax(y)])
    ax1.set_xlabel("time [sample]")
    ax1.set_ylabel("amplitude")
    '''

    ax1 = fig.add_subplot(321)
    ax1.plot(t, y_db, "-", markersize=1)
    ax1.axis([start_sec, (N+start_pos) * dt, np.amin(y_db), np.amax(y_db)+10])
    ax1.set_xlabel("Time [sec]")
    ax1.set_ylabel("Amplitude [dB]")

    ax2 = fig.add_subplot(322)
    ax2.set_xscale('log')
    ax2.axis([10, fs/2, np.amin(decibelSpectrum), np.amax(decibelSpectrum)+10])
    ax2.plot(freqList, decibelSpectrum, '-', markersize=1)
    ax2.set_xlabel("Frequency [Hz]")
    ax2.set_ylabel("Amplitude [dB]")

    ax3 = fig.add_subplot(323)
    ax3.plot(freqList, decibelSpectrum, '-', markersize=1)
    ax3.axis([0, fs/2, np.amin(decibelSpectrum), np.amax(decibelSpectrum)+10])
    ax3.set_xlabel("Frequency [Hz]")
    ax3.set_ylabel("Amplitude [dB]")

    ax4 = fig.add_subplot(324)
    ax4.set_xscale('log')
    ax4.axis([10, fs/2, -180, 180])
    ax4.set_yticks(np.linspace(-180, 180, 9))
    ax4.plot(freqList, phaseSpectrum, '-', markersize=1)
    ax4.set_xlabel("Frequency [Hz]")
    ax4.set_ylabel("Phase [deg]")

    ax5 = fig.add_subplot(325)
    ax5.plot(t, y, "-", markersize=1)
    ax5.axis([start_sec, (N+start_pos)*dt, np.amin(y)*0.9, np.amax(y)*1.1])
    ax5.set_xlabel("Time [sec]")
    ax5.set_ylabel("Amplitude")

    ax6 = fig.add_subplot(326)
    ax6.axis([10, fs/2, np.amin(amplitudeSpectrum)*0.9, np.amax(amplitudeSpectrum)*1.1])
    ax6.plot(freqList, amplitudeSpectrum, '-', markersize=1)
    ax6.set_xlabel("Frequency [Hz]")
    ax6.set_ylabel("Amplitude")

    # subplot(314)
    # xscale('linear')
    # plot(freqList, phaseSpectrum,".")
    # axis([0,fs/2,-np.pi,np.pi])
    # xlabel("frequency[Hz]")
    # ylabel("phase [rad]")

    try:
        plt.show()
    finally:
        plt.close()
