#!/usr/bin/env bash

for algo in LMS NLMS AP RLS; do
  for L in 4 16 64 512; do
    echo "${algo} start calculation with length ${L}"
    ./raw_drone_convergence dr_static_20.wav ${algo} ${L} 1.0 8
  done
done
