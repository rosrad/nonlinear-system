#!/bin/sh

if [[ $# != 0 ]] ; then
    echo "update the data"
    go build -o henon_ly ./henon.go
    ./henon_ly
fi
echo "plot the henon map"
gnuplot henon.gp
