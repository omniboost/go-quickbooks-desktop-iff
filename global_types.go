package qdiif

import (
	"time"
)

type Date struct {
	time.Time
}

func (d Date) IIF() string {
	if d.IsZero() {
		return ""
	}

	return d.String()
}

type Time struct {
	time.Time
}

func (t Time) IIF() string {
	if t.IsZero() {
		return ""
	}

	return t.String()
}

type Bool bool

func (b Bool) IIF() string {
	if b {
		return "Y"
	}

	return "N"
}
