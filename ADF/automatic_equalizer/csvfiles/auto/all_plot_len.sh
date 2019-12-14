#!/usr/bin/env bash

source /Users/tetsu/personal_files/Research/research_tools/venv/bin/activate;

APP_NAME="auto"

for algo in NLMS RLS; do
#for len in 4 16 64 256 1024; do
  echo "${algo} start to plot with length ${len}"
  python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv_multi.py ${algo}_${APP_NAME}_L-4_mse.csv ${algo}_${APP_NAME}_L-64_mse.csv ${algo}_${APP_NAME}_L-1024_mse.csv  -d ../../imgfiles/${APP_NAME}
  echo -e "\n\n\n"
done

deactivate;

