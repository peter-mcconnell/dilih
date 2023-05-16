drop it like its hot
====================

![dilih balancy logo](./media/logo.png "drop it like its hot")

build bpf module
----------------

```sh
(cd bpf && sudo make dilih_kern.o)
```

build golang code
-----------------

```sh
CGO_ENABLED=0 go build
```

running with docker
-------------------

```sh
docker build -t=dilih .
docker run --net=host --cap-add NET_ADMIN --cap-add BPF --cap-add PERFMON -u0 -e INTERFACE=ens160 --rm dilih
```
