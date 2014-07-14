set term png  size 1920,1080
set output "henon_phase.png"
set grid
set key left top

set yrange [-1.5:2]
set xrange [-1.5:2]
plot "henon_phase.dat" title "the phase of henon map" with points
unset output
