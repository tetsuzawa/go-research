#!/usr/bin/env bash


APP_NAME="auto_on_ref"

source /Users/tetsu/personal_files/Research/research_tools/venv/bin/activate;

#for algo in LMS NLMS RLS; do
for algo in NLMS RLS; do
  for len in 4 16 64 128 256; do
    echo "${algo} start to plot with length ${len}"
    python /Users/tetsu/personal_files/Research/research_tools/plot_from_csv.py ${algo}_${APP_NAME}_L-${len}.csv -d ../../imgfiles/${APP_NAME}
    mse -tap 256 ${algo}_${APP_NAME}_L-${len}.csv
    python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv.py ${algo}_${APP_NAME}_L-${len}_mse.csv -d ../../imgfiles/${APP_NAME}
    echo -e "\n\n\n"
  done
done

for algo in AP; do
  for order in 8; do
    for len in 4 16 64 128 256; do
      echo "${algo} start to plot with length ${len}"
      python /Users/tetsu/personal_files/Research/research_tools/plot_from_csv.py ${algo}_${APP_NAME}_L-${len}_order-${order}.csv -d ../../imgfiles/${APP_NAME}
    mse -tap 256 ${algo}_${APP_NAME}_L-${len}_order-${order}.csv
    python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv.py ${algo}_${APP_NAME}_L-${len}_order-${order}_mse.csv -d ../../imgfiles/${APP_NAME}
      echo -e "\n\n\n"
    done
  done
done

deactivate;
