#!/usr/bin/env bash

for algo in LMS NLMS AP RLS; do
  for len in 4 16 64 256 1024; do
    echo "${algo} start calculation with length ${len}"
    ./build/run_linux ../jsonfiles/${algo}_static_L-${len}.json ../csvfiles/white
  done
done
