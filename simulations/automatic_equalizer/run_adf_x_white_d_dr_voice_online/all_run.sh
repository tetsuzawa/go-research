#!/usr/bin/env bash


BIN_NAME="run_adf_x_white_d_dr_voice_online"
DIR_NAME="auto_on"

make build

#for algo in LMS NLMS RLS; do
for algo in NLMS RLS; do
  for len in 4 16 64 128 256; do
    echo "${algo} start calculation with length ${len}"
    ./build/${BIN_NAME} ../jsonfiles/${DIR_NAME}/${algo}_static_L-${len}.json ../csvfiles/${DIR_NAME}
  done
done

for algo in AP; do
  for order in 8; do
    for len in 4 16 64 128 256; do
      echo "${algo} start calculation with length ${len}"
      ./build/${BIN_NAME} ../jsonfiles/${DIR_NAME}/${algo}_static_L-${len}_order-${order}.json ../csvfiles/${DIR_NAME}
    done
  done
done