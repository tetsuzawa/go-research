#!/usr/bin/env bash

source /Users/tetsu/personal_files/Research/research_tools/venv/bin/activate;
subject="white"

for algo in NLMS RLS; do
  for len in 4 16 64 256 1024; do
    echo "${algo} start to calc with length ${len}"
    mse_csv -tap 512 ${algo}_${subject}_L-${len}.csv
    echo -e "\n\n\n"
  done
done

for algo in AP; do
  for order in 8; do
    for len in 4 16 64 256 1024; do
      echo "${algo} start to calc with length ${len}"
      mse_csv -tap 512 ${algo}_${subject}_L-${len}_order-${order}.csv
      echo -e "\n\n\n"
    done
  done
done

deactivate;

