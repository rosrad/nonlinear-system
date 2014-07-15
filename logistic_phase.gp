set term png  size 1920,1080
set output "logistic_phase.png"
set grid
# set xrange [0:80]
set yrange [-0.5:1.5]
set key left top

# set tmargin 2

set xlabel "interation"
set ylabel "x"
plot "logistic_phase.dat" title "the phase of logistic map" with points


# set ylabel "dx"
# plot "logistic_dx.dat" w d 
# unset multiplot

unset output
