// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mips64

import (
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/ssagen"
	"cmd/internal/obj/mips"
	"cmd/internal/objabi"
)

func Init(arch *ssagen.ArchInfo) {
	arch.LinkArch = &mips.Linkmips64
	if objabi.GOARCH == "mips64le" {
		arch.LinkArch = &mips.Linkmips64le
	}
	arch.REGSP = mips.REGSP
	arch.MAXWIDTH = 1 << 50
	arch.SoftFloat = objabi.GOMIPS64 == "softfloat"
	arch.ZeroRange = zerorange
	arch.Ginsnop = ginsnop
	arch.Ginsnopdefer = ginsnop

	arch.SSAMarkMoves = func(s *ssagen.State, b *ssa.Block) {}
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
