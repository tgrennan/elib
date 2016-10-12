// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=loop -id looperIn -d VecType=looperInVec -d Type=LooperIn github.com/platinasystems/go/elib/vec.tmpl]

package loop

import (
	"github.com/platinasystems/go/elib"
)

type looperInVec []LooperIn

func (p *looperInVec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]LooperIn, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *looperInVec) validate(new_len uint, zero *LooperIn) *LooperIn {
	c := elib.Index(cap(*p))
	lʹ := elib.Index(len(*p))
	l := elib.Index(new_len)
	if l <= c {
		// Need to reslice to larger length?
		if l >= lʹ {
			*p = (*p)[:l]
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *looperInVec) validateSlowPath(zero *LooperIn,
	c, l, lʹ elib.Index) *LooperIn {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]LooperIn, cNext, cNext)
		copy(q, *p)
		if zero != nil {
			for i := c; i < cNext; i++ {
				q[i] = *zero
			}
		}
		*p = q[:l]
	}
	if l > lʹ {
		*p = (*p)[:l]
	}
	return &(*p)[l-1]
}

func (p *looperInVec) Validate(i uint) *LooperIn {
	return p.validate(i+1, (*LooperIn)(nil))
}

func (p *looperInVec) ValidateInit(i uint, zero LooperIn) *LooperIn {
	return p.validate(i+1, &zero)
}

func (p *looperInVec) ValidateLen(l uint) (v *LooperIn) {
	if l > 0 {
		v = p.validate(l, (*LooperIn)(nil))
	}
	return
}

func (p *looperInVec) ValidateLenInit(l uint, zero LooperIn) (v *LooperIn) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p looperInVec) Len() uint { return uint(len(p)) }