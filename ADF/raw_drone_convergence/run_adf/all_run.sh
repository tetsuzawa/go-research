#!/usr/bin/env sh

for algo in LMS NLMS AP RLS; do
  for len in 4 16 64 256 1024; do
    echo "${algo} start calculation with length ${len}"
    ./run_linux ../jsonfiles/${algo}_static_L-${len}.json ../csvfiles
  done
done
