set term png  size 1920,1080
set output "logistic.png"
set grid
set xrange [0.8:4]
set yrange [-1:1]
set key left top
# set multiplot layout 2,1 title "Multiplot layout "

# set tmargin 2
# set title "Plot x"
set xlabel "a"
set ylabel "x"
plot "logistic.dat" title "logistic map" with dots,\
     "logistic_ly.dat" title "Lyapunov of logistic " with line,\
     "logistic_ly_zero.dat.mlt" using 1:2 with points notitle, \
     "" using 1:2:3 with labels offset 0,1 notitle

# set title "Plot for dx"
# set ylabel "dx"
# plot "logistic_dx.dat" w d 
# unset multiplot

unset output
