#!/usr/bin/env sh

/Users/tetsu/personal_files/Research/venv/bin/activate;

for algo in LMS NLMS RLS; do
  for len in 4 16 64 256 1024; do
    echo "${algo} start to plot with length ${len}"
    python /Users/tetsu/personal_files/Research/research_tools/plot_from_csv.py ${algo}_static_L-${len}.csv -d ../imgfiles
    echo "\n\n\n"
  done
done

for algo in AP; do
  for order in 8; do
    for len in 4 16 64 256 1024; do
      echo "${algo} start to plot with length ${len}"
      python /Users/tetsu/personal_files/Research/research_tools/plot_from_csv.py ${algo}_static_L-${len}_order-${order}.csv -d ../imgfiles
      echo "\n\n\n"
    done
done

deactivate;
