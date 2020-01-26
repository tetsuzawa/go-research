#!/usr/bin/env bash

source /Users/tetsu/repo/research-tools/pyresearch/venv/bin/activate;

#for algo in LMS NLMS RLS; do
for len in 4 16 64 256 1024; do
  echo "${algo} start to plot with length ${len}"
  #python /Users/tetsu/repo/research-tools/pyresearch/plot_MSE_iter_from_csv_multi.py NLMS_white_L-${len}_mse.csv AP_white_L-${len}_order-8_mse.csv RLS_white_L-${len}_mse.csv  -d ../../imgfiles/white
  python /Users/tetsu/repo/research-tools/pyresearch/plot_MSE_iter_from_csv_multi.py NLMS_white_L-${len}_mse.csv AP_white_L-${len}_order-8_mse.csv RLS_white_L-${len}_mse.csv  -d .
  echo -e "\n\n\n"
done

deactivate;

