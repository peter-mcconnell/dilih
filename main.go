package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
)

func main() {
	spec, err := ebpf.LoadCollectionSpec("bpf/dilih_kern.o")
	if err != nil {
		panic(err)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		panic(fmt.Sprintf("Failed to create new collection: %v\n", err))
	}
	defer coll.Close()

	prog := coll.Programs["xdp_dilih"]
	if prog == nil {
		panic("No program named 'xdp_dilih' found in collection")
	}

	iface := os.Getenv("INTERFACE")
	if iface == "" {
		panic("No interface specified. Please set the INTERFACE environment variable to the name of the interface to be use")
	}
	iface_idx, err := net.InterfaceByName(iface)
	if err != nil {
		panic(fmt.Sprintf("Failed to get interface %s: %v\n", iface, err))
	}
	opts := link.XDPOptions{
		Program:   prog,
		Interface: iface_idx.Index,
		// Flags is one of XDPAttachFlags (optional).
	}
	lnk, err := link.AttachXDP(opts)
	if err != nil {
		panic(err)
	}
	defer lnk.Close()

	fmt.Println("Successfully loaded and attached BPF program.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
