#!/usr/bin/env bash

source /Users/tetsu/repo/research-tools/pyresearch/venv/bin/activate;

PWD=$(pwd)
APP_NAME=$(basename ${PWD})

for SNR in $(seq -40 5 0); do
  echo -e "\n\n*********************start calculation with SN Rate ${SNR}*********************\n\n"
  mkdir -p ../../imgfiles/"${APP_NAME}"/SNR_"${SNR}"
  for len in 4 16 64 128 256; do
    echo "${algo} start to plot with length ${len}"
    python /Users/tetsu/repo/research-tools/pyresearch/plot_MSE_iter_from_csv_multi.py SNR_"${SNR}"/NLMS_"${APP_NAME}"_L-${len}_mse.csv  SNR_"${SNR}"/AP_"${APP_NAME}"_L-${len}_order-8_mse.csv  SNR_"${SNR}"/RLS_"${APP_NAME}"_L-${len}_mse.csv  -d ../../imgfiles/"${APP_NAME}"/SNR_"${SNR}"
    echo -e "\n\n\n"
  done
done

deactivate;

