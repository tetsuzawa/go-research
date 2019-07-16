# -*- coding: utf-8 -*-
# python3.7 VS_plot.DXB.py [FILE_NAME] [FFT_LENGTH]
#
import struct
import sys
import math
import string
import numpy as np
import os.path
import pandas as pd
import matplotlib.pyplot as plt
import scipy as sp
import wave


class wave_proccess():

    def __init__(self, filename):
        # open wave file
        wf = wave.open(filename, 'r')

        # waveファイルが持つ性質を取得
        self.filename = filename
        self.ch = wf.getnchannels()
        self.width = wf.getsampwidth()
        self.fr = wf.getframerate()
        self.chunk_size = wf.getnframes()
        # load wave data
        self.amp = (2**8) ** self.width / 2
        data = wf.readframes(self.chunk_size)   # バイナリ読み込み
        data = np.frombuffer(data, 'int16')  # intに変換
        data = data / self.amp                  # 振幅正規化
        self.data = data[::self.ch]
        wf.close()

        return

    def wave_write(self, filename, data_array):
        ww = wave.open(filename, 'w')
        ww.setnchannels(self.ch)
        ww.setsampwidth(self.width)
        ww.setframerate(self.fr)
        ww.writeframes(data_array)
        ww.close()
