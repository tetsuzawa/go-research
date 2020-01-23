#!/usr/bin/env bash

source /Users/tetsu/repo/research-tools/pyresearch/venv/bin/activate;

PWD=$(pwd)
APP_NAME=$(basename ${PWD})

# for SNR in $(seq -40 5 0); do
#   echo -e "\n\n*********************start calculation with SN Rate ${SNR}*********************\n\n"
#   for algo in NLMS RLS; do
#     for len in 4 16 64 128 256; do
#       echo "${algo} start to plot with length ${len}"
#       calc_mse_csv -tap "$1" SNR_"${SNR}"/${algo}_"${APP_NAME}"_L-${len}.csv SNR_"${SNR}"
#       echo -e "\n\n\n"
#     done
#   done
# done

for SNR in $(seq -40 5 0); do
  for algo in AP; do
    for order in 8; do
      for len in 4 16 64 128 256; do
        echo "${algo} start to plot with length ${len}"
        calc_mse_csv -tap "$1" SNR_"${SNR}"/${algo}_"${APP_NAME}"_L-${len}_order-${order}.csv SNR_"${SNR}"
        echo -e "\n\n\n"
      done
    done
  done
done

deactivate;

