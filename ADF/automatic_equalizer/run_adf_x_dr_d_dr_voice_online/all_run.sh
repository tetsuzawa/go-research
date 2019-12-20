#!/usr/bin/env bash


BIN_NAME="run_adf_x_dr_d_dr_voice_online"
DIR_NAME="auto_on_ref"

make build

#for algo in LMS NLMS RLS; do
algo="NLMS"
for len in 4 16 64 128 256; do
  echo "${algo} start calculation with length ${len}"
  ./build/${BIN_NAME} ../jsonfiles/${DIR_NAME}/${algo}_static_L-${len}.json ../csvfiles/${DIR_NAME}
done

algo="RLS"
for len in 4 16 64 128 256; do
  echo "${algo} start calculation with length ${len}"
  ./build/${BIN_NAME} ../jsonfiles/${DIR_NAME}/${algo}_static_L-${len}.json ../csvfiles/${DIR_NAME}
done

algo="AP"
for order in 8; do
  for len in 4 16 64 128 256; do
    echo "${algo} start calculation with length ${len}"
    ./build/${BIN_NAME} ../jsonfiles/${DIR_NAME}/${algo}_static_L-${len}_order-${order}.json ../csvfiles/${DIR_NAME}
  done
done
