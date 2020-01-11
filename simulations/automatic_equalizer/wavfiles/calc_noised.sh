#! /usr/bin/env bash

# /usr/bin/env python3 ~/personal_files/Research/research_tools/calc_noisemix_list.py \
# 	--clean_file ~/.ghq/github.com/tetsuzawa/audio-SNR/data/source_clean/arctic_a0001.wav \
# 	--noise_file dr_static_5_16k.wav \
# 	--output_dir noise_mixed/ \
# 	--snr_start -40 \
# 	--snr_end 40 \
# 	--snr_div_num 19

WAV_DIR="$(cd $(dirname $0); pwd)"

#calc_noise_mix \
# -clean /Users/tetsu/.ghq/github.com/tetsuzawa/audio-SNR/data/source_clean/arctic_a0001.wav \
# -noise ${WAV_DIR}/dr_static_5_16k.wav \
# -output ${WAV_DIR}/noise_mixed/ \
# -start -20 \
# -end 20 \
# -div 9


#calc_noise_mix \
# -clean /Users/tetsu/personal_files/Research/sample_wav/fukushima_20sec.wav \
# -noise ${WAV_DIR}/dr_static_20.wav \
# -output ${WAV_DIR}/noise_mixed/ \
# -start -40 \
# -end 0 \
# -div 9

#01/10 for real x and d
calc_noise_mix \
 -clean ../DRONE_01_09/voice_l_20sec_ir_convolved_20sec.wav \
 -noise ../DRONE_01_09/drone_subtracted_l_20sec_ir_convolved_by_r.wav \
 -output ${WAV_DIR}/noise_mixed_convo_d/ \
 -start -40 \
 -end 0 \
 -div 9

calc_noise_mix \
 -clean ../DRONE_01_09/voice_r_20sec_ir_convolved_20sec.wav \
 -noise ../DRONE_01_09/drone_subtracted_r_20sec.wav \
 -output ${WAV_DIR}/noise_mixed_convo_x/ \
 -start -40 \
 -end 0 \
 -div 9
