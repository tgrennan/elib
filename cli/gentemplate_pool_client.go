// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=cli -id client -d Data=clients -d PoolType=clientPool -d Type=client github.com/platinasystems/elib/pool.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"github.com/platinasystems/elib"
)

type clientPool struct {
	elib.Pool
	clients []client
}

func (p *clientPool) GetIndex() (i uint) {
	l := uint(len(p.clients))
	i = p.Pool.GetIndex(l)
	if i >= l {
		p.Validate(i)
	}
	return i
}

func (p *clientPool) PutIndex(i uint) (ok bool) {
	return p.Pool.PutIndex(i)
}

func (p *clientPool) IsFree(i uint) (v bool) {
	v = i >= uint(len(p.clients))
	if !v {
		v = p.Pool.IsFree(i)
	}
	return
}

func (p *clientPool) Resize(n uint) {
	c := uint(cap(p.clients))
	l := uint(len(p.clients) + int(n))
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]client, l, c)
		copy(q, p.clients)
		p.clients = q
	}
	p.clients = p.clients[:l]
}

func (p *clientPool) Validate(i uint) {
	c := uint(cap(p.clients))
	l := uint(i) + 1
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]client, l, c)
		copy(q, p.clients)
		p.clients = q
	}
	if l > uint(len(p.clients)) {
		p.clients = p.clients[:l]
	}
}

func (p *clientPool) Elts() uint {
	return uint(len(p.clients)) - p.FreeLen()
}

func (p *clientPool) Len() uint {
	return uint(len(p.clients))
}

func (p *clientPool) Foreach(f func(x client)) {
	for i := range p.clients {
		if !p.Pool.IsFree(uint(i)) {
			f(p.clients[i])
		}
	}
}

func (p *clientPool) ForeachIndex(f func(i uint)) {
	for i := range p.clients {
		if !p.Pool.IsFree(uint(i)) {
			f(uint(i))
		}
	}
}

func (p *clientPool) Reset() {
	p.Pool.Reset()
	if len(p.clients) > 0 {
		p.clients = p.clients[:0]
	}
}
