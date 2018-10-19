// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=mctree -id shared_pair_offsets_pool -d PoolType=shared_pair_offsets_pool -d Type=shared_pair_offsets -d Data=elts github.com/platinasystems/elib/pool.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mctree

import (
	"github.com/platinasystems/elib"
)

type shared_pair_offsets_pool struct {
	elib.Pool
	elts []shared_pair_offsets
}

func (p *shared_pair_offsets_pool) GetIndex() (i uint) {
	l := uint(len(p.elts))
	i = p.Pool.GetIndex(l)
	if i >= l {
		p.Validate(i)
	}
	return i
}

func (p *shared_pair_offsets_pool) PutIndex(i uint) (ok bool) {
	return p.Pool.PutIndex(i)
}

func (p *shared_pair_offsets_pool) IsFree(i uint) (v bool) {
	v = i >= uint(len(p.elts))
	if !v {
		v = p.Pool.IsFree(i)
	}
	return
}

func (p *shared_pair_offsets_pool) Resize(n uint) {
	c := uint(cap(p.elts))
	l := uint(len(p.elts) + int(n))
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]shared_pair_offsets, l, c)
		copy(q, p.elts)
		p.elts = q
	}
	p.elts = p.elts[:l]
}

func (p *shared_pair_offsets_pool) Validate(i uint) {
	c := uint(cap(p.elts))
	l := uint(i) + 1
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]shared_pair_offsets, l, c)
		copy(q, p.elts)
		p.elts = q
	}
	if l > uint(len(p.elts)) {
		p.elts = p.elts[:l]
	}
}

func (p *shared_pair_offsets_pool) Elts() uint {
	return uint(len(p.elts)) - p.FreeLen()
}

func (p *shared_pair_offsets_pool) Len() uint {
	return uint(len(p.elts))
}

func (p *shared_pair_offsets_pool) Foreach(f func(x shared_pair_offsets)) {
	for i := range p.elts {
		if !p.Pool.IsFree(uint(i)) {
			f(p.elts[i])
		}
	}
}

func (p *shared_pair_offsets_pool) ForeachIndex(f func(i uint)) {
	for i := range p.elts {
		if !p.Pool.IsFree(uint(i)) {
			f(uint(i))
		}
	}
}

func (p *shared_pair_offsets_pool) Reset() {
	p.Pool.Reset()
	if len(p.elts) > 0 {
		p.elts = p.elts[:0]
	}
}
