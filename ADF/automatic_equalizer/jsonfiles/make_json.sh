#!/usr/bin/env bash

make build

DIR_NAME="$1"

for len in 4 16 64 128 256; do

  dr_json -wav ../wavfiles/dr_static_20.wav -adf NLMS -mu 1.00 -L ${len} -order 8 -dir ${DIR_NAME}
  dr_json -wav ../wavfiles/dr_static_20.wav -adf AP -mu 1.00 -L ${len} -order 8 -dir ${DIR_NAME}
  dr_json -wav ../wavfiles/dr_static_20.wav -adf RLS -mu 1.00 -L ${len} -order 8 -dir ${DIR_NAME}

done

