package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/perf"
)

type event struct {
	TimeSinceBoot  uint64
	ProcessingTime uint32
	Type           uint8
}

const ringBufferSize = 128 // size of ring buffer used to calculate average processing times
type ringBuffer struct {
	data [ringBufferSize]uint32
	start int
	size int
}

func (rb *ringBuffer) add(val uint32) {
	if rb.size < ringBufferSize {
		rb.size++
	} else {
		rb.size = 1
	}
	rb.data[rb.size - 1] = val
}

func (rb *ringBuffer) avg() uint32 {
	if rb.size == 0 {
		return 0
	}
	sum := uint32(0)
	for i := 0; i < rb.size; i++ {
		sum += uint32(rb.data[i])
	}
	return sum / uint32(rb.size)
}

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

	// handle perf events
	outputMap, ok := coll.Maps["output_map"]
	if !ok {
		panic("No map named 'output_map' found in collection")
	}
	perfEvent, err := perf.NewReader(outputMap, 4096)
	if err != nil {
		panic(fmt.Sprintf("Failed to create perf event reader: %v\n", err))
	}
	defer perfEvent.Close()
	buckets := map[uint8]uint64{
		1:   0, // bpf program entered
		2:   0, // bpf program dropped
		3:   0, // bpf program passed
	}

	processingTimePassed := &ringBuffer{}
	processingTimeDropped := &ringBuffer{}

	go func() {
		// var event event
		for {
			record, err := perfEvent.Read()
			if err != nil {
				fmt.Println(err)
				continue
			}

			var e event
			if len(record.RawSample) < 9 {
				fmt.Println("Invalid sample size")
				continue
			}
			// time since boot in the first 8 bytes
			e.TimeSinceBoot = binary.LittleEndian.Uint64(record.RawSample[:8])
			// processing time in the next 4 bytes
			e.ProcessingTime = binary.LittleEndian.Uint32(record.RawSample[8:12])
			// type in the last byte
			e.Type = uint8(record.RawSample[12])

			if e.Type == 2 {
				processingTimeDropped.add(e.ProcessingTime)
			}
			if e.Type == 3 {
				processingTimePassed.add(e.ProcessingTime)
			}

			buckets[e.Type]++
			fmt.Print("\033[H\033[2J")
			fmt.Printf("total: %d. passed: %d. dropped: %d. passed processing time avg: %d. dropped processing time avg: %d\n", buckets[1], buckets[3], buckets[2], processingTimePassed.avg(), processingTimeDropped.avg())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
