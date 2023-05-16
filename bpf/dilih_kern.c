#include "dilih_kern.h"

SEC("xdp_lb")
int xdp_dilih(struct xdp_md *ctx)
{
	bpf_printk("got something");

	return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
