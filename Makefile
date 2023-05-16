.PHONY: build
.DEFAULT_GOAL = build

INTERFACE := ens160

build_bpf:
	(cd bpf && sudo make dilih_kern.o)

build_go:
	CGO_ENABLED=0 go build
	
build_docker:
	docker build -t=dilih .

build: build_bpf build_go build_docker

run_docker:
	docker run --net=host --cap-add NET_ADMIN --cap-add BPF --cap-add PERFMON -u0 -e INTERFACE=$(INTERFACE) --rm dilih
