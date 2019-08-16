# coding: utf-8
import numpy as np


class MPU9250(object):

    def initialize(self, port, address):
        pass

    def get_motion6(self):
        # ma = [1, 3, 5]
        # mg = [2, 4, 6]
        ma = np.random.normal(loc=0, scale=1, size=(3,))
        mg = np.random.normal(loc=0, scale=1, size=(3,))
        return ma, mg
