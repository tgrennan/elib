// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=loop -id pollerElogEvent -d Type=pollerElogEvent github.com/platinasystems/go/elib/elog/event.tmpl]

package loop

import (
	"github.com/platinasystems/go/elib/elog"
)

var pollerElogEventType = &elog.EventType{
	Name: "loop.pollerElogEvent",
}

func init() {
	t := pollerElogEventType
	t.Stringer = stringer_pollerElogEvent
	t.Encode = encode_pollerElogEvent
	t.Decode = decode_pollerElogEvent
	elog.RegisterType(pollerElogEventType)
}

func stringer_pollerElogEvent(e *elog.Event) string {
	var x pollerElogEvent
	x.Decode(e.Data[:])
	return x.String()
}

func encode_pollerElogEvent(b []byte, e *elog.Event) int {
	var x pollerElogEvent
	x.Decode(e.Data[:])
	return x.Encode(b)
}

func decode_pollerElogEvent(b []byte, e *elog.Event) int {
	var x pollerElogEvent
	x.Decode(b)
	return x.Encode(e.Data[:])
}

func (x pollerElogEvent) Log() { x.Logb(elog.DefaultBuffer) }

func (x pollerElogEvent) Logb(b *elog.Buffer) {
	e := b.Add(pollerElogEventType)
	x.Encode(e.Data[:])
}
