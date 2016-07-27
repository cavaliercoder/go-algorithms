# go-algorithms

This project contains implementations of common computer science related
algorithms written in Go. My aim is not to provide reference implementations,
but simply to practice while working my way through the third edition of 
*Introduction to Algorithms* from MIT Press.

Documentation is available on [godoc.org](https://godoc.org/github.com/cavaliercoder/go-algorithms).

To test all algorithms for correctness with pseudo-random input, run

    $ make test

To test all algorithms for correctness repeatedly with fuzzed input, run:

    $ make fuzz

To benchmark all algorithms for best, worst and pseudo-random (hopefully 
average) case, run:

    $ make bench
