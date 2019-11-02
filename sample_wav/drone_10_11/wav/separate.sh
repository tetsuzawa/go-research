for Angle in `seq 1 7`; do
	python /Users/tetsu/personal_files/Research/research_tools/calc_stereo2monoLR.py dr_level_0${Angle}.wav dr_level_0${Angle}_R.wav dr_level_0${Angle}_L.wav
done
