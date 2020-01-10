L: under
R: upper
python ~/personal_files/Research/research_tools/calc_subtracted_wav_2.py none.wav drone_r.wav drone_subtrancted_r.wav
python ~/personal_files/Research/research_tools/calc_subtracted_wav_2.py none.wav drone_l.wav drone_subtrancted_l.wav
sox drone_subtracted_r.wav drone_subtracted_r_20sec.wav trim 0 20
python drone_mic_ir.py drone_subtracted_l_20sec.wav drone_subtracted_r_20sec.wav pseude_ir.wav

SSTFをDDAからcsvにした
SSTFを1000サンプルに編集した
python csv_to_wav_each_column.py SSTF/SSTF_check_L.csv SSTF/SSTF_check_L.wav
python csv_to_wav_each_column.py SSTF/SSTF_check_R.csv SSTF/SSTF_check_R.wav
fast_convolution_wav -x drone_subtracted_r_20sec.wav -y drone_pseude_ir.wav -o drone_subtracted_l_20sec_ir_convolved_by_r.wav
