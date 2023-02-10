package example

import (
	"testing"

	"github.com/zcalusic/sysinfo/cpuid"
)

func hasLZCNT() bool {
	// https://www.amd.com/system/files/TechDocs/25481.pdf
	const (
		Fn8000_0001 = 0x8000_0001
		ABM         = 0x0400_0000 // ABM is the 5th bit, from the left, zero based, EDX register
		EDX         = 3
	)
	var out [4]uint32
	cpuid.CPUID(&out, Fn8000_0001)
	return out[EDX]&ABM != 0
}

func TestAmd64CpuId_cpuHasFeature(t *testing.T) {
	if !hasLZCNT() {
		panic("oh no")
	} else {
		println("all good")
	}
}
