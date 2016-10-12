// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=loop -id eventElogEvent -d Type=eventElogEvent github.com/platinasystems/go/elib/elog/event.tmpl]

package loop

import (
	"github.com/platinasystems/go/elib/elog"
)

var eventElogEventType = &elog.EventType{
	Name: "loop.eventElogEvent",
}

func init() {
	t := eventElogEventType
	t.Stringer = stringer_eventElogEvent
	t.Encode = encode_eventElogEvent
	t.Decode = decode_eventElogEvent
	elog.RegisterType(eventElogEventType)
}

func stringer_eventElogEvent(e *elog.Event) string {
	var x eventElogEvent
	x.Decode(e.Data[:])
	return x.String()
}

func encode_eventElogEvent(b []byte, e *elog.Event) int {
	var x eventElogEvent
	x.Decode(e.Data[:])
	return x.Encode(b)
}

func decode_eventElogEvent(b []byte, e *elog.Event) int {
	var x eventElogEvent
	x.Decode(b)
	return x.Encode(e.Data[:])
}

func (x eventElogEvent) Log() { x.Logb(elog.DefaultBuffer) }

func (x eventElogEvent) Logb(b *elog.Buffer) {
	e := b.Add(eventElogEventType)
	x.Encode(e.Data[:])
}
