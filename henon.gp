set term png  size 1920,1080
set output "henon.png"
set grid
set key left top

set multiplot layout 2,1 title "Henon Map"
set tmargin 2
set xrange [0:1.5]

set yrange [-1.5:2]
set title "Plot "
set xlabel "a"
set ylabel "xn"
plot "henon.dat" title "henon map" with dots
#dots

set title "Plot ly"
set yrange [-1:0.5]
set xlabel "a"
set ylabel "xn"
plot "henon_ly.dat" title "henon ly" with dots,\
     "henon_ly_zero.dat.mlt" with points ,\
    "" using 1:2:3 with labels offset 0,1
unset multiplot

unset output
