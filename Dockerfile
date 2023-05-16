FROM scratch
COPY loadbalancer /loadbalancer
COPY bpf/lb_kern.o /bpf/lb_kern.o
CMD ["/loadbalancer"]
