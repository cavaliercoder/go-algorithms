# go-algorithms

This project contains implementations of common computer science related
algorithms written in Go. My aim is not to provide reference implementations,
but simply to practice while working my way through the third edition of 
*Introduction to Algorithms* from MIT Press.

Documentation is available on [godoc.org](https://godoc.org/github.com/cavaliercoder/go-algorithms).

To test all algorithms for correctness with random input run

    $ go test -v

To benchmark all algorithms for best, worst and pseudo-random (hopefully 
average) case, run:

    $ go test -v -bench=. -benchtime=10s
