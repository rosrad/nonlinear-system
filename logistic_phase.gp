set term png  size 1920,1080
set output "logistic_phase.png"
set grid
set xrange [0:80]
#set yrange [-1:1]
set key left top

# set tmargin 2

set xlabel "interation"
set ylabel "x"
plot "logistic_phase.dat" title "the phase of logistic map" with line


# set ylabel "dx"
# plot "logistic_dx.dat" w d 
# unset multiplot

unset output
