#!/usr/bin/env bash

source /Users/tetsu/personal_files/Research/research_tools/venv/bin/activate;

APP_NAME="auto_on"

for algo in NLMS RLS; do
  echo "${algo} start to plot with length"
  #python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv_multi_v2.py ${algo}_${APP_NAME}_L-4_mse.csv ${algo}_${APP_NAME}_L-256_mse.csv ${algo}_${APP_NAME}_L-1024_mse.csv ${algo}_${APP_NAME}_L-8096_mse.csv --dst_path ../../imgfiles/${APP_NAME} --sample 2000
  python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv_multi_v2.py ${algo}_${APP_NAME}_L-4_mse.csv ${algo}_${APP_NAME}_L-256_mse.csv ${algo}_${APP_NAME}_L-1024_mse.csv ${algo}_${APP_NAME}_L-8096_mse.csv --dst_path ../../imgfiles/${APP_NAME}
  echo -e "\n\n\n"
done

for algo in AP; do
  echo "${algo} start to plot with length"
  #python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv_multi_v2.py ${algo}_${APP_NAME}_L-4_order-8_mse.csv ${algo}_${APP_NAME}_L-256_order-8_mse.csv ${algo}_${APP_NAME}_L-1024_order-8_mse.csv ${algo}_${APP_NAME}_L-8096_order-8_mse.csv --dst_path ../../imgfiles/${APP_NAME} --sample 2000
  python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv_multi_v2.py ${algo}_${APP_NAME}_L-4_order-8_mse.csv ${algo}_${APP_NAME}_L-256_order-8_mse.csv ${algo}_${APP_NAME}_L-1024_order-8_mse.csv ${algo}_${APP_NAME}_L-8096_order-8_mse.csv --dst_path ../../imgfiles/${APP_NAME}
  echo -e "\n\n\n"
done

deactivate;

