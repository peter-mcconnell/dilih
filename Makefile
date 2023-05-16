.PHONY: build build_bpf build_go build_docker clean run_docker
.DEFAULT_GOAL = build

INTERFACE := ens160

build_bpf:
	$(MAKE) -C bpf dilih_kern.o

build_go:
	CGO_ENABLED=0 go build
	
build_docker:
	docker build -t=dilih .

build: build_bpf build_go build_docker

clean:
	$(MAKE) -C bpf clean

run_docker:
	docker run --net=host --cap-add NET_ADMIN --cap-add BPF --cap-add PERFMON -u0 -e INTERFACE=$(INTERFACE) --rm dilih
