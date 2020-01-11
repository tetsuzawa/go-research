#!/usr/bin/env bash

#for algo in LMS NLMS RLS; do
for algo in NLMS RLS; do
  for len in 4 16 64 256; do
    echo "${algo} start calculation with length ${len}"
    ./build/run_adf_white ../jsonfiles/white/${algo}_white_L-${len}.json ../csvfiles/white
  done
done

for algo in AP; do
  for order in 8; do
    for len in 4 16 64 256; do
      echo "${algo} start calculation with length ${len}"
      ./build/run_adf_white ../jsonfiles/white/${algo}_white_L-${len}_order-${order}.json ../csvfiles/white
    done
  done
done
