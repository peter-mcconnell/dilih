FROM scratch
COPY dilih /dilih
COPY bpf/dilih_kern.o /bpf/dilih_kern.o
CMD ["/dilih"]
