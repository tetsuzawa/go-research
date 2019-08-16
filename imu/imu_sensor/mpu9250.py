# coding: utf-8


class MPU9250(object):

    def initialize(self, port, address):
        pass

    def get_motion6(self):
        ma = [1, 3, 5]
        mg = [2, 4, 6]
        return ma, mg
