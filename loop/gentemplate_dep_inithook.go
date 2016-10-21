// autogenerated: do not edit!
// generated from gentemplate [gentemplate -id initHook -d Package=loop -d DepsType=initHookVec -d Type=initHook -d Data=hooks github.com/platinasystems/go/elib/dep/dep.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loop

import (
	"github.com/platinasystems/go/elib/dep"
)

type initHookVec struct {
	deps  dep.Deps
	hooks []initHook
}

func (t *initHookVec) Len() int {
	return t.deps.Len()
}

func (t *initHookVec) Get(i int) initHook {
	return t.hooks[t.deps.Index(i)]
}

func (t *initHookVec) Add(x initHook, ds ...*dep.Dep) {
	if len(ds) == 0 {
		t.deps.Add(&dep.Dep{})
	} else {
		t.deps.Add(ds[0])
	}
	t.hooks = append(t.hooks, x)
}
