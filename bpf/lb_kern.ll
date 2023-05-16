; ModuleID = 'lb_kern.c'
source_filename = "lb_kern.c"
target datalayout = "e-m:e-p:64:64-i64:64-i128:128-n32:64-S128"
target triple = "bpf"

@xdp_load_balancer.____fmt = internal constant [14 x i8] c"got something\00", align 1
@_license = dso_local global [4 x i8] c"GPL\00", section "license", align 1
@llvm.compiler.used = appending global [2 x ptr] [ptr @_license, ptr @xdp_load_balancer], section "llvm.metadata"

; Function Attrs: nounwind
define dso_local i32 @xdp_load_balancer(ptr nocapture readnone %0) #0 section "xdp_lb" {
  %2 = tail call i64 (ptr, i32, ...) inttoptr (i64 6 to ptr)(ptr noundef nonnull @xdp_load_balancer.____fmt, i32 noundef 14) #1
  ret i32 2
}

attributes #0 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" }
attributes #1 = { nounwind }

!llvm.module.flags = !{!0, !1}
!llvm.ident = !{!2}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{i32 7, !"frame-pointer", i32 2}
!2 = !{!"clang version 16.0.3 (https://github.com/llvm/llvm-project.git da3cd333bea572fb10470f610a27f22bcb84b08c)"}
