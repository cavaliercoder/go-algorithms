test:
	go test -v -cover

bench:
	go test -bench=.

algorithms-fuzz.zip: *.go
	go-fuzz-build \
		-o algorithms-fuzz.zip \
		github.com/cavaliercoder/go-algorithms

fuzz: algorithms-fuzz.zip corpus
	go-fuzz \
		-workdir=./fuzz \
		-bin algorithms-fuzz.zip

corpus:
	rm -f fuzz/corpus/rand*
	mkdir -p fuzz/corpus
	head -c 16384 /dev/urandom > fuzz/corpus/rand1
	head -c 32768 /dev/urandom > fuzz/corpus/rand2
	head -c 65536 /dev/urandom > fuzz/corpus/rand3
	head -c 131072 /dev/urandom > fuzz/corpus/rand4

get-deps:
	go get github.com/dvyukov/go-fuzz/go-fuzz
	go get github.com/dvyukov/go-fuzz/go-fuzz-build


.PHONY: test bench fuzz corpus get-deps
