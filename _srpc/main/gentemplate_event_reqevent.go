// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=main -id ReqEvent -d Type=reqEvent github.com/platinasystems/elib/elog/event.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/platinasystems/elib/elog"
)

var reqEventType = &elog.EventType{
	Name: "main.reqEvent",
}

func init() {
	t := reqEventType
	t.Stringer = stringer_reqEvent
	t.Encode = encode_reqEvent
	t.Decode = decode_reqEvent
	elog.RegisterType(reqEventType)
}

func stringer_reqEvent(e *elog.Event) string {
	var x reqEvent
	x.Decode(e.Data[:])
	return x.String()
}

func encode_reqEvent(b []byte, e *elog.Event) int {
	var x reqEvent
	x.Decode(e.Data[:])
	return x.Encode(b)
}

func decode_reqEvent(b []byte, e *elog.Event) int {
	var x reqEvent
	x.Decode(b)
	return x.Encode(e.Data[:])
}

func (x reqEvent) Log() { x.Logb(elog.DefaultBuffer) }

func (x reqEvent) Logb(b *elog.Buffer) {
	e := b.Add(reqEventType)
	x.Encode(e.Data[:])
}
