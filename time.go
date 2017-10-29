package gen

import (
	"math/rand"
	"time"
)

func Time(r *rand.Rand) *TimeGen {
	return &TimeGen{
		r,
		time.Unix(0, 0),
		time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
}

type TimeGen struct {
	r      *rand.Rand
	after  time.Time
	before time.Time
}

func (g *TimeGen) After(t time.Time) *TimeGen {
	g.after = t
	return g
}

func (g *TimeGen) Before(t time.Time) *TimeGen {
	g.before = t
	return g
}

func (g *TimeGen) Gen() *time.Time {
	d := g.before.Sub(g.after)
	t := g.after.Add(time.Duration(Int(g.r).Max64(int64(d)).Gen64()))
	return &t
}