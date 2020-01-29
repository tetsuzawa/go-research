for Angle in `seq 1 7`; do
	sox dr_level_0${Angle}.wav dr_level_0${Angle}_L.wav remix 1
	sox dr_level_0${Angle}.wav dr_level_0${Angle}_R.wav remix 2
done
