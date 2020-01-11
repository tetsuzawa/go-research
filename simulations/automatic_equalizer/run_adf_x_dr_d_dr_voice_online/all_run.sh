#!/usr/bin/env bash

GOOS="darwin"

BIN_NAME="run_adf_x_dr_d_dr_voice_online_${GOOS}"
DIR_NAME="auto_on_ref_convo"

make build

#	jsonName := os.Args[1]
#	dataDir := os.Args[2]
#	xWavPath := os.Args[3]
#	dWavPath := os.Args[4]

for SNR in $(seq -40 5 0); do
  echo -e "\n***start calculation with SN Rate ${SNR}***\n"

  SAVE_DIR_NAME="../csvfiles/${DIR_NAME}/SNR_${SNR}"
  mkdir -p ${SAVE_DIR_NAME}

  #for algo in LMS NLMS RLS; do
  algo="NLMS"
  for len in 4 16 64 128 256; do
    echo "${algo} start calculation with length ${len}"
    ./bin/${BIN_NAME} ../jsonfiles/${DIR_NAME}/${algo}_static_L-${len}.json ${SAVE_DIR_NAME} ../wavfiles/noise_mixed_convo_x/voice_r_20sec_ir_convolved_20sec_snr${SNR}.wav ../wavfiles/noise_mixed_convo_d/voice_l_20sec_ir_convolved_20sec_snr${SNR}.wav
  done

  algo="RLS"
  for len in 4 16 64 128 256; do
    echo "${algo} start calculation with length ${len}"
    ./bin/${BIN_NAME} ../jsonfiles/${DIR_NAME}/${algo}_static_L-${len}.json ${SAVE_DIR_NAME} ../wavfiles/noise_mixed_convo_x/voice_r_20sec_ir_convolved_20sec_snr${SNR}.wav ../wavfiles/noise_mixed_convo_d/voice_l_20sec_ir_convolved_20sec_snr${SNR}.wav
  done

  algo="AP"
  for order in 8; do
    for len in 4 16 64 128 256; do
      echo "${algo} start calculation with length ${len}"
      ./bin/${BIN_NAME} ../jsonfiles/${DIR_NAME}/${algo}_static_L-${len}_order-${order}.json ${SAVE_DIR_NAME} ../wavfiles/noise_mixed_convo_x/voice_r_20sec_ir_convolved_20sec_snr${SNR}.wav ../wavfiles/noise_mixed_convo_d/voice_l_20sec_ir_convolved_20sec_snr${SNR}.wav
    done
  done
done
