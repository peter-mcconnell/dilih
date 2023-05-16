loady balancy
=============

![loady balancy logo](./media/logo.png "loady balancy")

build bpf module
----------------

```sh
(cd bpf && sudo make lb_kern.o)
```

build golang code
-----------------

```sh
CGO_ENABLED=0 go build
```

running with docker
-------------------

```sh
docker run --net=host --cap-add NET_ADMIN --cap-add BPF --cap-add PERFMON -u0 -e INTERFACE=ens160 --rm loadbalancer
```
