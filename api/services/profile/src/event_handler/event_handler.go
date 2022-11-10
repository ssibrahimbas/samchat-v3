package event_handler

import (
	"fmt"
	"os"

	natsGo "github.com/nats-io/nats.go"
	"github.com/ssibrahimbas/samchat-v3.profile/src/internal"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/nats"
)

type Handler struct {
	n    *nats.Conn
	s    *internal.Srv
	subs []*natsGo.Subscription
}



func New(n *nats.Conn, s *internal.Srv) *Handler {
	return &Handler{
		n: n,
		s: s,
	}
}

func (eh *Handler) ListenAll() {}

func (eh *Handler) subscribeSub(s string, h natsGo.MsgHandler) {
	sub, err := eh.n.Subscribe(s, h)
	if err != nil {
		fmt.Println(err)
		return
	}
	eh.subs = append(eh.subs, sub)
}

func (eh *Handler) OnClose(s os.Signal) {
	for _, sub := range eh.subs {
		_ = sub.Unsubscribe()
	}
}

