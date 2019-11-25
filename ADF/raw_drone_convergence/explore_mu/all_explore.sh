#!/usr/bin/env bash

#for algo in LMS NLMS AP RLS; do
#  for len in 4 16 64 256 1024; do
#    echo "${algo} start calculation with length ${len}"
#    ./explore_mu -l ${len} -start 0.000001 -end 2 -step 100 ../wavfiles/dr_static_1.wav ${algo} ../jsonfiles
#  done
#done

algo="LMS"
for len in 4 16 64 256 1024; do
  echo "${algo} start calculation with length ${len}"
  ./build/explore_mu -l ${len} -start 0.000001 -end 2 -step 100 ../wavfiles/dr_static_1.wav ${algo} ../jsonfiles
done

algo="NLMS"
for len in 4 16 64 256 1024; do
  echo "${algo} start calculation with length ${len}"
  ./build/explore_mu -l ${len} -start 0.000001 -end 2 -step 100 ../wavfiles/dr_static_1.wav ${algo} ../jsonfiles
done

algo="AP"
for len in 4 16 64 256 1024; do
  echo "${algo} start calculation with length ${len}"
  ./build/explore_mu -l ${len} -start 0.000001 -end 30 -step 100 ../wavfiles/dr_static_1.wav ${algo} ../jsonfiles
done

algo="RLS"
for len in 4 16 64 256 1024; do
  echo "${algo} start calculation with length ${len}"
  ./build/explore_mu -l ${len} -start 0.000001 -end 30 -step 100 ../wavfiles/dr_static_1.wav ${algo} ../jsonfiles
done
