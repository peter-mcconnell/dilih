#include "dilih_kern.h"

struct perf_trace_event {
	__u64 timestamp; // time elapsed since boot, excluding suspend time. see https://www.man7.org/linux/man-pages/man7/bpf-helpers.7.html
	__u32 processing_time_ns;
	__u8 type;
};

struct {
	__uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
	__uint(key_size, sizeof(int));
	__uint(value_size, sizeof(__u32));
	__uint(max_entries, 1024);
} output_map SEC(".maps");

SEC("xdp")
int xdp_dilih(struct xdp_md *ctx)
{
	struct perf_trace_event e = {};

	// perf event for entering xdp program
	e.timestamp = bpf_ktime_get_ns();
	e.type = 1;
	e.processing_time_ns = 0;
	bpf_perf_event_output(ctx, &output_map, BPF_F_CURRENT_CPU, &e, sizeof(e));
	
	if (bpf_get_prandom_u32() % 2 == 0) {

		// perf event for dropping packet
		e.type = 2;
		__u64 ts = bpf_ktime_get_ns();
		e.processing_time_ns = ts - e.timestamp;
		e.timestamp = ts;
		bpf_perf_event_output(ctx, &output_map, BPF_F_CURRENT_CPU, &e, sizeof(e));

		bpf_printk("dropping packet");
		return XDP_DROP;
	}

	// perf event for passing packet
	e.type = 3;
	__u64 ts = bpf_ktime_get_ns();
	e.processing_time_ns = ts - e.timestamp;
	e.timestamp = ts;

	bpf_perf_event_output(ctx, &output_map, BPF_F_CURRENT_CPU, &e, sizeof(e));
	bpf_printk("passing packet");

	return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
