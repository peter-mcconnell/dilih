dilih
=====

A simple example project that shows how to build XDP eBPF programs in C and load them with Golang; accompanying article: https://www.petermcconnell.com/posts/writing-an-xdp-ebpf-program/.

This program will drop ~50% of packets on a given interface, allowing you to inspect the impact faulty network has on your applications. Note: you can also accomplish this functionality with `tc`.

_drop it like it's hot_

![dilih balancy logo](./media/logo.png "drop it like its hot")

build dependencies
------------------

This project was built and tested in the following environment:

```
 - x86_64
 - linux kernel 5.15.0-71-generic
 - llvm 16.0.3 with -bpf target support (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/llvm.yaml)
 - clang 16.0.3 (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/llvm.yaml)
 - golang go1.20.2 (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/golang.yaml)
 - docker 23.0.1 (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/docker.yaml)
 - make 4.3
```

build everything (bpf program, go program, docker image)
--------------------------------------------------------

If you haven't already, ensure the libbpf submodule is pulled:
```sh
$ git submodule init
$ git submodule update
```

Then proceed to build:
```sh
$ make
```

running with docker
-------------------

```sh
# set DEV= to your device (check 'ip link')
$ DEV=eth0 make run_docker
# press Ctrl+C when you want to resume normality
```

Or you can run the image from docker hub (no need to pull this repo):

```sh
# set INTERFACE= to your device (check 'ip link')
docker run --net=host --cap-add SYS_ADMIN -u0 -e INTERFACE=ens160 --rm pemcconnell/dilih:v0.0.1-pre
```

testing
-------

With the program running you should start to see disruption. For example, when applied to the hosts interface used for the internet connection we can see ~50% packet loss:

```sh
$ ping -c8 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=3 ttl=115 time=20.6 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=115 time=28.1 ms
64 bytes from 8.8.8.8: icmp_seq=7 ttl=115 time=20.8 ms
64 bytes from 8.8.8.8: icmp_seq=8 ttl=115 time=21.8 ms

--- 8.8.8.8 ping statistics ---
8 packets transmitted, 4 received, 50% packet loss, time 7073ms
rtt min/avg/max/mdev = 20.557/22.809/28.085/3.078 ms
```

_Note: it may not actually be 50% each time - the logic depends on randomness._

clean
-----

```sh
$ DEV=eth0 make clean
```
