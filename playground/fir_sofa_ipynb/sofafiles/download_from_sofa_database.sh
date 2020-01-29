#!/usr/bin/env bash

source /Users/tetsu/personal_files/Research/research_tools/venv/bin/activate

BASE_URL="https://sofacoustics.org/data/database/"

#for DB in ari cipic riec aachen; do
for DB in cipic riec aachen; do
  mkdir -p ${DB}
  echo -e "\n${DB} start to download\n"
  python /Users/tetsu/personal_files/Research/research_tools/download_sofa.py -d ${DB} ${BASE_URL}${DB}/
done

deactivate;
