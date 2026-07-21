package baggage

import (
	"context"
)

type baggageKeyType struct{}

var baggageKey = baggageKeyType{}

type Member struct {
	key   string
	value string
}

func NewMember(key, value string) (Member, error) {
	return Member{key: key, value: value}, nil
}

func (m Member) Key() string {
	return m.key
}

func (m Member) Value() string {
	return m.value
}

type Baggage struct {
	members map[string]Member
}

func New(members ...Member) (Baggage, error) {
	b := Baggage{
		members: make(map[string]Member),
	}
	for _, m := range members {
		b.members[m.key] = m
	}
	return b, nil
}

func (b Baggage) Members() []Member {
	result := make([]Member, 0, len(b.members))
	for _, m := range b.members {
		result = append(result, m)
	}
	return result
}

func (b Baggage) Member(key string) Member {
	return b.members[key]
}

func ContextWithBaggage(ctx context.Context, b Baggage) context.Context {
	return context.WithValue(ctx, baggageKey, b)
}

func FromContext(ctx context.Context) Baggage {
	if ctx == nil {
		return Baggage{members: make(map[string]Member)}
	}
	b, ok := ctx.Value(baggageKey).(Baggage)
	if !ok {
		return Baggage{members: make(map[string]Member)}
	}
	return b
}
