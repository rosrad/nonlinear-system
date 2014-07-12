set term png  size 1920,1080
set output "henon.png"
set grid
set xrange [0:1.5]
set yrange [-2:4]
set key left top

set xlabel "a"
set ylabel "x/ly"
plot "henon.dat" title "henon map" with dots, "henon_ly.dat" title "henon ly" with dots 


# unset multiplot

unset output
