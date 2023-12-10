package es

import "shopping/internal/ddd"

type (
	Versioner interface {
		Version() int
		PendingVersion() int
	}

	Aggregate struct {
		ddd.Aggregate
		version int
	}
)

var _ interface {
	EventCommitter
	Versioner
	VersionSetter
} = (*Aggregate)(nil)

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Aggregate: ddd.NewAggregate(id, name),
		version:   0,
	}
}

func (a *Aggregate) AddEvent(name string, payload ddd.EventPayload, ops ...ddd.EventOption) {
	ops = append(ops, ddd.Metadata{
		ddd.AggregateVersionKey: a.PendingVersion() + 1,
	})
	a.Aggregate.AddEvent(name, payload, ops...)
}

func (a *Aggregate) CommitEvent() {
	a.version += len(a.Events())
	a.ClearEvents()
}

func (a *Aggregate) setVersion(version int) { a.version = version }

func (a Aggregate) Version() int        { return a.version }
func (a Aggregate) PendingVersion() int { return a.version + len(a.Events()) }
