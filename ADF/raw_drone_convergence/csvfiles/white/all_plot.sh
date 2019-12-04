#!/usr/bin/env bash

source /Users/tetsu/personal_files/Research/research_tools/venv/bin/activate;
subject="white"

#for algo in LMS NLMS RLS; do
for algo in NLMS RLS; do
  for len in 4 16 64 256 1024; do
    echo "${algo} start to plot with length ${len}"
    python /Users/tetsu/personal_files/Research/research_tools/plot_from_csv.py ${algo}_${subject}_L-${len}.csv -d ../../imgfiles/${subject}
    mse_csv -tap 256 ${algo}_${subject}_L-${len}.csv 
    python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv.py ${algo}_${subject}_L-${len}_mse.csv -d ../../imgfiles/${subject} 
    echo -e "\n\n\n"
  done
done

for algo in AP; do
  for order in 8; do
    for len in 4 16 64 256 1024; do
      echo "${algo} start to plot with length ${len}"
      python /Users/tetsu/personal_files/Research/research_tools/plot_from_csv.py ${algo}_${subject}_L-${len}_order-${order}.csv -d ../../imgfiles/${subject} 
    mse_csv -tap 256 ${algo}_${subject}_L-${len}_order-${order}.csv
    python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv.py ${algo}_${subject}_L-${len}_order-${order}_mse.csv -d ../../imgfiles/${subject}
      echo -e "\n\n\n"
    done
  done
done

deactivate;

