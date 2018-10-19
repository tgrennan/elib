// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=loop -id Vi -d VecType=viVec -d Type=Vi github.com/platinasystems/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loop

import (
	"github.com/platinasystems/elib"
)

type viVec []Vi

func (p *viVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]Vi, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *viVec) validate(new_len uint, zero Vi) *Vi {
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

func (p *viVec) validateSlowPath(zero Vi, old_cap, new_len, old_len uint) *Vi {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]Vi, new_cap, new_cap)
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

func (p *viVec) Validate(i uint) *Vi {
	var zero Vi
	return p.validate(i+1, zero)
}

func (p *viVec) ValidateInit(i uint, zero Vi) *Vi {
	return p.validate(i+1, zero)
}

func (p *viVec) ValidateLen(l uint) (v *Vi) {
	if l > 0 {
		var zero Vi
		v = p.validate(l, zero)
	}
	return
}

func (p *viVec) ValidateLenInit(l uint, zero Vi) (v *Vi) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *viVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p viVec) Len() uint { return uint(len(p)) }
