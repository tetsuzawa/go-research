#!/usr/bin/env bash

source /Users/tetsu/personal_files/Research/research_tools/venv/bin/activate;

APP_NAME="auto_on"

#for algo in LMS NLMS RLS; do
for len in 4 16 64 256 1024; do
  echo "${algo} start to plot with length ${len}"
  python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv_multi.py NLMS_${APP_NAME}_L-${len}_mse.csv AP_${APP_NAME}_L-${len}_order-8_mse.csv RLS_${APP_NAME}_L-${len}_mse.csv  -d ../../imgfiles/${APP_NAME}
  echo -e "\n\n\n"
done

deactivate;

