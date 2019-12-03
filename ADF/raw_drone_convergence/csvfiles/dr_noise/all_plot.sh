#!/usr/bin/env bash

source /Users/tetsu/personal_files/Research/research_tools/venv/bin/activate;

#for algo in LMS NLMS RLS; do
for algo in NLMS RLS; do
  for len in 4 16 64 256 1024; do
    echo "${algo} start to plot with length ${len}"
    python /Users/tetsu/personal_files/Research/research_tools/plot_from_csv.py ${algo}_dr_noise_L-${len}.csv -d ../../imgfiles/dr_noise
    python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv.py ${algo}_dr_noise_L-${len}.csv -d ../../imgfiles/dr_noise -t 256
    echo -e "\n\n\n"
  done
done

for algo in AP; do
  for order in 8; do
    for len in 4 16 64 256 1024; do
      echo "${algo} start to plot with length ${len}"
      python /Users/tetsu/personal_files/Research/research_tools/plot_from_csv.py ${algo}_dr_noise_L-${len}_order-${order}.csv -d ../../imgfiles/dr_noise
    python /Users/tetsu/personal_files/Research/research_tools/plot_MSE_iter_from_csv.py ${algo}_dr_noise_L-${len}_order-${order}.csv -d ../../imgfiles/dr_noise -t 256
      echo -e "\n\n\n"
    done
  done
done

deactivate;

