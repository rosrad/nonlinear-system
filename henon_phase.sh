#!/bin/bash

go build -o henon_phase ./henon.go
dir=henon_phase_png
rm ./$dir -fr 
mkdir $dir
for i in $(cat henon_ly_zero.dat.mlt |cut -d' ' -f1|sort|uniq) ; do
    echo "shell para" $i 
    ./henon_phase -a $i
    gnuplot henon_phase.gp 
    mv henon_phase.png $dir/henon_phase_$i.png
done
