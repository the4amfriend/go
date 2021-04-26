+build nsx,386 nsx,arm

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func init() {
	// On 32-bit NSX systems, the fcntl syscall that matches Go's
	// Flock_t type is SYS_FCNTL64, not SYS_FCNTL. (Assumption based on Linux, if disproved - need to be fixed)
	fcntl64Syscall = SYS_FCNTL64
}
