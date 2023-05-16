dilih
=====

_drop it like it's hot_

![dilih balancy logo](./media/logo.png "drop it like its hot")

build dependencies
------------------

This project was built and tested in the following environment:

```
 - linux kernel 5.15.0-71-generic
 - llvm 16.0.3 with -bpf target support (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/llvm.yaml)
 - clang 16.0.3 (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/llvm.yaml)
 - golang go1.20.2 (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/golang.yaml)
 - docker 23.0.1 (https://github.com/peter-mcconnell/.dotfiles/blob/master/tasks/docker.yaml)
 - make 4.3
```

build everything (bpf program, go program, docker image)
--------------------------------------------------------

```sh
make
```

running with docker
-------------------

```sh
DEV=eth0 make run_docker

# you should now see xdpgeneric on the given interface - ensure you clean this up !
```

clean
-----

```sh
DEV=eth0 make clean
```
