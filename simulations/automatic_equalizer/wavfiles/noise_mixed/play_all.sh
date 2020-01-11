#!/usr/bin/env bash

dirs=`find *.wav -maxdepth 0 -type f`

for dir in ${dirs};
do
	echo ${dir}
	play ${dir}
done
