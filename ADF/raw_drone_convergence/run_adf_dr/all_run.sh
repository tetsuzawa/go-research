#!/usr/bin/env sh

for algo in NLMS RLS; do
  for len in 4 16 64 256; do
    echo "${algo} start calculation with length ${len}"
    ./build/run_adf_dr ../jsonfiles/dr/${algo}_static_L-${len}.json ../csvfiles/dr_noise
  done
done

for algo in AP; do
  for order in 8; do
    for len in 4 16 64 256; do
      echo "${algo} start calculation with length ${len}"
      ./build/run_adf_dr ../jsonfiles/dr/${algo}_static_L-${len}_order-${order}.json ../csvfiles/dr_noise
    done
  done
done
