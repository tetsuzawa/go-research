#! /usr/bin/env bash

# /usr/bin/env python3 ~/personal_files/Research/research_tools/calc_noisemix_list.py \
# 	--clean_file ~/.ghq/github.com/tetsuzawa/audio-SNR/data/source_clean/arctic_a0001.wav \
# 	--noise_file dr_static_5_16k.wav \
# 	--output_dir noise_mixed/ \
# 	--snr_start -40 \
# 	--snr_end 40 \
# 	--snr_div_num 19


calc_noise_mix \
 -clean /Users/tetsu/.ghq/github.com/tetsuzawa/audio-SNR/data/source_clean/arctic_a0001.wav \
 -noise /Users/tetsu/personal_files/Research/sample_wav/drone_10_11/wav/dr_static_5_16k.wav \
 -output /Users/tetsu/personal_files/Research/sample_wav/drone_10_11/wav/tmp\
 -start -20 \
 -end 20 \
 -div 9

