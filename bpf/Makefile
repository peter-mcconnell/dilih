TARGET = dilih
BPF_TARGET = ${TARGET:=_kern}
BPF_C = ${BPF_TARGET:=.c}
BPF_OBJ = ${BPF_C:.c=.o}

BPF_PINNED_PATH := /sys/fs/bpf/$(TARGET)
XDP_NAME := dilih
DEV := ens160

xdp: $(BPF_OBJ)
	bpftool net detach $(XDP_NAME) dev $(DEV)
	rm -f $(BPF_PINNED_PATH)
	bpftool prog load $(BPF_OBJ) $(BPF_PINNED_PATH)
	bpftool net attach $(XDP_NAME) pinned $(BPF_PINNED_PATH) dev $(DEV)

$(BPF_OBJ): %.o: %.c
	clang -S \
		-g \
		-target bpf \
		-D __BPF_TRACING__ \
	  -I../libbpf/src\
		-Wall \
		-Wno-unused-value \
		-Wno-pointer-sign \
		-Wno-compare-distinct-pointer-types \
		-Werror \
		-O2 -emit-llvm -c -o ${@:.o=.ll} $<
	llc -march=bpf -filetype=obj -O2 -o $@ ${@:.o=.ll}

clean:
	-bpftool net detach $(XDP_NAME) dev $(DEV)
	sudo rm -f $(BPF_PINNED_PATH)
	rm -f $(BPF_OBJ)
	rm -f ${BPF_OBJ:.o=.ll}
