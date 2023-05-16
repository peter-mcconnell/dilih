#include "dilih_kern.h"

SEC("xdp")
int xdp_dilih(struct xdp_md *ctx)
{
	if (bpf_get_prandom_u32() % 2 == 0) {
		bpf_printk("dropping packet");
		return XDP_DROP;
	}
	bpf_printk("passing packet");
	return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
