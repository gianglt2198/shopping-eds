package am

import "time"

type AckType int

const (
	AckTypeAuto AckType = iota
	AckTypeManual
)

var defaultAckWait = 5 * time.Second
var defaultMaxRetentionDeliver = 5

type SubscriberConfig struct {
	msgFilter           []string
	groupName           string
	ackType             AckType
	ackWait             time.Duration
	maxRetentionDeliver int
}

func NewSubscriberConfig(ops []SubscriberOption) SubscriberConfig {
	cfg := SubscriberConfig{
		msgFilter:           []string{},
		groupName:           "",
		ackType:             AckTypeManual,
		ackWait:             defaultAckWait,
		maxRetentionDeliver: defaultMaxRetentionDeliver,
	}

	for _, op := range ops {
		op.configureSubscriberConfig(&cfg)
	}

	return cfg
}

type SubscriberOption interface {
	configureSubscriberConfig(*SubscriberConfig)
}

func (c SubscriberConfig) MessageFilters() []string {
	return c.msgFilter
}

func (c SubscriberConfig) GroupName() string {
	return c.groupName
}

func (c SubscriberConfig) AckType() AckType {
	return c.ackType
}

func (c SubscriberConfig) AckWait() time.Duration {
	return c.ackWait
}

func (c SubscriberConfig) MaxRetetionDeliver() int {
	return c.maxRetentionDeliver
}

type MessageFilter []string

func (s MessageFilter) configureSubscriberConfig(cfg *SubscriberConfig) {
	cfg.msgFilter = s
}

type GroupName string

func (n GroupName) configureSubscriberConfig(cfg *SubscriberConfig) {
	cfg.groupName = string(n)
}

func (t AckType) configureSubscriberConfig(cfg *SubscriberConfig) {
	cfg.ackType = t
}

type AckWait time.Duration

func (w AckWait) configureSubscriberConfig(cfg *SubscriberConfig) {
	cfg.ackWait = time.Duration(w)
}

type MaxRetetionDeliver int

func (i MaxRetetionDeliver) configureSubscriberConfig(cfg *SubscriberConfig) {
	cfg.maxRetentionDeliver = int(i)
}
