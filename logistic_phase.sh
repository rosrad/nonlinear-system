#!/bin/bash
dir=logistic_phase_png
rm ./$dir -fr 
mkdir $dir
for i in  $(cat logistic_ly_zero.dat.mlt |cut -d' ' -f1|sort|uniq); do
    echo "shell para" $i
    go run ./logistic.go -a $i
    gnuplot logistic_phase.gp 
    mv logistic_phase.png $dir/logistic_phase_$i.png
done
