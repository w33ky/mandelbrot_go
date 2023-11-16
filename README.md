compile with `go build`

run `./mandelbrot -help` to get all command line options 

examples:
```
./mandelbrot -resX 2000 -resY 2000 -iterations 40 -posX -0.7 -posY 0 -scaleX 1.4 -scaleY 1.4
./mandelbrot -resX 2000 -resY 2000 -iterations 40 -posX -0.2 -posY 0.9 -scaleX 0.3 -scaleY 0.3
./mandelbrot -resX 4000 -resY 4000 -iterations 250 -posX -0.1437 -posY 0.8879 -scaleX 0.0005 -scaleY 0.0005 -multithread -colorPreset red
```