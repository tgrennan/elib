// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=event -id actor -d VecType=ActorVec -d Type=Actor github.com/platinasystems/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

import (
	"github.com/platinasystems/elib"
)

type ActorVec []Actor

func (p *ActorVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]Actor, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *ActorVec) validate(new_len uint, zero Actor) *Actor {
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

func (p *ActorVec) validateSlowPath(zero Actor, old_cap, new_len, old_len uint) *Actor {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]Actor, new_cap, new_cap)
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

func (p *ActorVec) Validate(i uint) *Actor {
	var zero Actor
	return p.validate(i+1, zero)
}

func (p *ActorVec) ValidateInit(i uint, zero Actor) *Actor {
	return p.validate(i+1, zero)
}

func (p *ActorVec) ValidateLen(l uint) (v *Actor) {
	if l > 0 {
		var zero Actor
		v = p.validate(l, zero)
	}
	return
}

func (p *ActorVec) ValidateLenInit(l uint, zero Actor) (v *Actor) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *ActorVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p ActorVec) Len() uint { return uint(len(p)) }
