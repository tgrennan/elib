// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=loop -id looperIn -d VecType=looperInVec -d Type=LooperIn github.com/platinasystems/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loop

import (
	"github.com/platinasystems/elib"
)

type looperInVec []LooperIn

func (p *looperInVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]LooperIn, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *looperInVec) validate(new_len uint, zero LooperIn) *LooperIn {
	old_cap := uint(cap(*p))
	old_len := uint(len(*p))
	if new_len <= old_cap {
		// Need to reslice to larger length?
		if new_len > old_len {
			*p = (*p)[:new_len]
			for i := old_len; i < new_len; i++ {
				(*p)[i] = zero
			}
		}
		return &(*p)[new_len-1]
	}
	return p.validateSlowPath(zero, old_cap, new_len, old_len)
}

func (p *looperInVec) validateSlowPath(zero LooperIn, old_cap, new_len, old_len uint) *LooperIn {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]LooperIn, new_cap, new_cap)
		copy(q, *p)
		for i := old_len; i < new_cap; i++ {
			q[i] = zero
		}
		*p = q[:new_len]
	}
	if new_len > old_len {
		*p = (*p)[:new_len]
	}
	return &(*p)[new_len-1]
}

func (p *looperInVec) Validate(i uint) *LooperIn {
	var zero LooperIn
	return p.validate(i+1, zero)
}

func (p *looperInVec) ValidateInit(i uint, zero LooperIn) *LooperIn {
	return p.validate(i+1, zero)
}

func (p *looperInVec) ValidateLen(l uint) (v *LooperIn) {
	if l > 0 {
		var zero LooperIn
		v = p.validate(l, zero)
	}
	return
}

func (p *looperInVec) ValidateLenInit(l uint, zero LooperIn) (v *LooperIn) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *looperInVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p looperInVec) Len() uint { return uint(len(p)) }
