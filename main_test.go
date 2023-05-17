package main

import (
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := &ringBuffer{}

	// Test adding values to the ring buffer
	rb.add(10)
	rb.add(20)
	rb.add(30)

	// Test the average calculation
	avg := rb.avg()
	expectedAvg := float32(20.0) // (10 + 20 + 30) / 3 = 20
	if avg != expectedAvg {
		t.Errorf("Average mismatch. Got: %f, Expected: %f", avg, expectedAvg)
	}
}

func TestRingBufferFull(t *testing.T) {
	rb := &ringBuffer{}

	for i := 1; i <= 128; i++ {  // 128 is the buffer size
		rb.add(uint32(i))
	}

	// Test the average calculation after exceeding the buffer size
	avg := rb.avg()
	// triange calc for 0..128 = 128 * (128 + 1) / 2 = 8256
	// 8256 / 128 = 64.5
	expectedAvg := float32(64.5)
	if avg != expectedAvg {
		t.Errorf("Average mismatch after exceeding buffer size. Got: %f, Expected: %f", avg, expectedAvg)
	}
}

func TestRingBufferOverflow(t *testing.T) {
	rb := &ringBuffer{}

	for i := 0; i < 130; i++ {  // 2 more than the buffer size
		rb.add(uint32(i + 1))
	}

	// Test the average calculation after exceeding the buffer size
	avg := rb.avg()
	// triange calc for 0..128 = 128 * (128 + 1) / 2 = 8256
	// replace first two values (1, 2) with (129, 130)
	// 8256 - 1 - 2 + 129 + 130 = 8256 + 256 = 8512
	// 8512 / 128 = 66.5
	expectedAvg := float32(66.5)
	if avg != expectedAvg {
		t.Errorf("Average mismatch after exceeding buffer size. Got: %f, Expected: %f", avg, expectedAvg)
	}
}
