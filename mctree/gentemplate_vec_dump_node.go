// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=mctree -id dump_node -d VecType=dump_node_vec -d Type=dump_node github.com/platinasystems/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mctree

import (
	"github.com/platinasystems/elib"
)

type dump_node_vec []dump_node

func (p *dump_node_vec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]dump_node, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *dump_node_vec) validate(new_len uint, zero dump_node) *dump_node {
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

func (p *dump_node_vec) validateSlowPath(zero dump_node, old_cap, new_len, old_len uint) *dump_node {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]dump_node, new_cap, new_cap)
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

func (p *dump_node_vec) Validate(i uint) *dump_node {
	var zero dump_node
	return p.validate(i+1, zero)
}

func (p *dump_node_vec) ValidateInit(i uint, zero dump_node) *dump_node {
	return p.validate(i+1, zero)
}

func (p *dump_node_vec) ValidateLen(l uint) (v *dump_node) {
	if l > 0 {
		var zero dump_node
		v = p.validate(l, zero)
	}
	return
}

func (p *dump_node_vec) ValidateLenInit(l uint, zero dump_node) (v *dump_node) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *dump_node_vec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p dump_node_vec) Len() uint { return uint(len(p)) }
