.PHONY: build build_bpf build_go build_docker clean run_docker push_docker
.DEFAULT_GOAL = build

INTERFACE := ens160
TAG := v0.0.1-pre

build_bpf:
	$(MAKE) -C bpf dilih_kern.o

build_go:
	CGO_ENABLED=0 go build
	
build_docker:
	docker build -t=pemcconnell/dilih:$(TAG) .

build: build_bpf build_go build_docker

clean:
	$(MAKE) -C bpf clean

run_docker: build_docker
	docker run --net=host --cap-add NET_ADMIN --cap-add BPF --cap-add PERFMON -u0 -e INTERFACE=$(INTERFACE) --rm pemcconnell/dilih:$(TAG)

push_docker: build_docker
	docker push pemcconnell/dilih:$(TAG)
