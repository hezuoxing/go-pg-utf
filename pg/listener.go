package pg

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
)

type Listener interface {
	Channel() <-chan pg.Notification
	ChannelSize(size int) <-chan pg.Notification
	Close() error
	Listen(ctx context.Context, channels ...string) error
	Receive(ctx context.Context) (channel string, payload string, err error)
	ReceiveTimeout(ctx context.Context, timeout time.Duration) (channel, payload string, err error)
	String() string
	Unlisten(ctx context.Context, channels ...string) error
}

type ListenerWrap struct {
	l Listener
}

func NewListener(l Listener) *ListenerWrap {
	return &ListenerWrap{l}
}

func (l *ListenerWrap) Channel() <-chan pg.Notification {
	return l.l.Channel()
}

func (l *ListenerWrap) ChannelSize(size int) <-chan pg.Notification {
	return l.l.ChannelSize(size)
}

func (l *ListenerWrap) Close() error {
	return l.l.Close()
}

func (l *ListenerWrap) Listen(ctx context.Context, channels ...string) error {
	return l.l.Listen(ctx, channels...)
}

func (l *ListenerWrap) Receive(ctx context.Context) (channel string, payload string, err error) {
	return l.l.Receive(ctx)
}

func (l *ListenerWrap) ReceiveTimeout(ctx context.Context, timeout time.Duration) (channel, payload string, err error) {
	return l.l.ReceiveTimeout(ctx, timeout)
}

func (l *ListenerWrap) String() string {
	return l.l.String()
}

func (l *ListenerWrap) Unlisten(ctx context.Context, channels ...string) error {
	return l.l.Unlisten(ctx, channels...)
}
