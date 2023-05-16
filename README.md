load balancer
=============

build bpf module
----------------

```sh
(cd bpf && sudo make lb_kern.o)
```

build golang code
-----------------

```sh
go build
```
