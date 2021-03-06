// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=cli -id file -d Data=Files -d PoolType=FilePool -d Type=File github.com/platinasystems/elib/pool.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cli

import (
	"github.com/platinasystems/elib"
)

type FilePool struct {
	elib.Pool
	Files []File
}

func (p *FilePool) GetIndex() (i uint) {
	l := uint(len(p.Files))
	i = p.Pool.GetIndex(l)
	if i >= l {
		p.Validate(i)
	}
	return i
}

func (p *FilePool) PutIndex(i uint) (ok bool) {
	return p.Pool.PutIndex(i)
}

func (p *FilePool) IsFree(i uint) (v bool) {
	v = i >= uint(len(p.Files))
	if !v {
		v = p.Pool.IsFree(i)
	}
	return
}

func (p *FilePool) Resize(n uint) {
	c := uint(cap(p.Files))
	l := uint(len(p.Files) + int(n))
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]File, l, c)
		copy(q, p.Files)
		p.Files = q
	}
	p.Files = p.Files[:l]
}

func (p *FilePool) Validate(i uint) {
	c := uint(cap(p.Files))
	l := uint(i) + 1
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]File, l, c)
		copy(q, p.Files)
		p.Files = q
	}
	if l > uint(len(p.Files)) {
		p.Files = p.Files[:l]
	}
}

func (p *FilePool) Elts() uint {
	return uint(len(p.Files)) - p.FreeLen()
}

func (p *FilePool) Len() uint {
	return uint(len(p.Files))
}

func (p *FilePool) Foreach(f func(x File)) {
	for i := range p.Files {
		if !p.Pool.IsFree(uint(i)) {
			f(p.Files[i])
		}
	}
}

func (p *FilePool) ForeachIndex(f func(i uint)) {
	for i := range p.Files {
		if !p.Pool.IsFree(uint(i)) {
			f(uint(i))
		}
	}
}

func (p *FilePool) Reset() {
	p.Pool.Reset()
	if len(p.Files) > 0 {
		p.Files = p.Files[:0]
	}
}
