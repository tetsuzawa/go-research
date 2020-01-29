import time
# import argparse
# import navio.mpu9250

import numpy as np

import ekf_functions as ekf_func
import imu_sensor.mpu9250

# IMU = navio.mpu9250.MPU9250()
IMU = imu_sensor.mpu9250.MPU9250()

# IMU.initialize(1, 0x06)  # lowpass->20Hz
time.sleep(1)

m6a0 = np.zeros((3,))
m6g0 = np.zeros((3,))

for _ in range(1, 1000):
    ma0, mg0 = IMU.get_motion6()
    m6a0 = m6a0 + np.array(ma0)
    m6g0 = m6g0 + np.array(mg0)
m6g0 = m6g0 / 1000
m6a0 = m6a0 / 1000
m6a0[0], m6a0[1] = m6a0[1], m6a0[0]
m6g0[0], m6g0[1] = m6g0[1], m6g0[0]

x0 = np.zeros((2,))
for _ in range(1, 1000):
    m6a, m6g = IMU.get_motion6()
    m6a[0], m6a[1] = -m6a[1], -m6a[0]
    m6a = np.array(m6a)
    x0 = x0 + ekf_func.get_angle_acc(m6a)
x0 = x0 / 1000

x = x0
P = np.zeros((2, 2))
Ts = 1.0 / 250.0
Yaw = 0
Tri = np.zeros((2, 3))

c = np.array([[1, 0], [0, 1]])
q = np.array([[1.74E-3 * Ts * Ts, 0], [0, 1.74E-3 * Ts * Ts]])
b = np.array([[1, 0], [0, 1]])
r = np.array([[1 * Ts * Ts, 0], [0, 1 * Ts * Ts]])
Cgy = np.eye(3, 3) * 1.74E-3 * Ts * Ts

while True:
    t_estimate_Attitude0 = time.time()
    m6a, m6g = IMU.get_motion6()
    m6a[0], m6a[1] = -m6a[1], -m6a[0]
    m6g[0], m6g[1] = m6g[1], m6g[0]

    m6a = np.array(m6a)
    m6g = np.array(m6g) - m6g0
    m6g[2] = -m6g[2]

    Tri = ekf_func.get_Trigonometrxic(x)
    J = ekf_func.Jacobian_forprocessvariance2(Tri)
    Jt = J.transpose()
    x, P = ekf_func.Kalman_filer2(x, ekf_func.get_angle_acc(m6a), m6g, c, b, q, r, P, Ts, Tri)
    Yaw = Yaw + (Tri[1, 1] / Tri[0, 0] * m6g[1] + Tri[1, 0] / Tri[0, 0] * m6g[2]) * Ts
    phi = np.array([Yaw, x[0] - x0[0], x[1] - x0[1]])
    print(phi)
    time.sleep(Ts - time.time() + t_estimate_Attitude0)
