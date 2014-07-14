#!/bin/bash

dir=logistic_phase_png

rm ./$dir -fr 
mkdir $dir
for i in 1.0 1.9 2.0 2.25 3.0 3.5 4.0 ; do
    echo "shell para" $i
    ./logistic_phase -a $i
    gnuplot logistic_phase.gp 
    mv logistic_phase.png $dir/logistic_phase_$i.png
done
