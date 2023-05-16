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
	// Open compiled BPF object file.
	spec, err := ebpf.LoadCollectionSpec("bpf/lb_kern.o")
	if err != nil {
		panic(err)
	}

	// Load the BPF programs from the spec.
	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		panic(fmt.Sprintf("Failed to create new collection: %v\n", err))
	}
	defer coll.Close()

	// Get program named "xdp_lb" from the collection.
	prog := coll.Programs["xdp_load_balancer"]
	if prog == nil {
		panic("No program named 'xdp_load_balancer' found in collection")
	}

	// Open a network interface.
	iface := os.Getenv("INTERFACE")
	if iface == "" {
		panic("No interface specified. Please set the INTERFACE environment variable to the name of the interface to be use")
	}
	iface_idx, err := net.InterfaceByName(iface)
	if err != nil {
		panic(fmt.Sprintf("Failed to get interface %s: %v\n", iface, err))
	}
	opts := link.XDPOptions{
		// Program must be an XDP BPF program.
		Program: prog,
		
		// Interface is the interface index to attach program to.
		Interface: iface_idx.Index,
		
		// Flags is one of XDPAttachFlags (optional).
		//
		// Only one XDP mode should be set, without flag defaults
		// to driver/generic mode (best effort).
	}
	lnk, err := link.AttachXDP(opts)
	if err != nil {
		panic(err)
	}
	defer lnk.Close()

	fmt.Println("Successfully loaded and attached BPF program.")

	// Handle SIGINT signal (Ctrl+C), to gracefully shutdown the application.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
