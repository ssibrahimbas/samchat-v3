package event_publisher

import (
	"encoding/json"

	"github.com/ssibrahimbas/samchat-v3.shared/pkg/helper"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/nats"
)

type Publisher struct {
	n *nats.Conn
}

func New(n *nats.Conn) *Publisher {
	return &Publisher{
		n: n,
	}
}

func (ep *Publisher) parseAndPublish(s string, d interface{}) {
	p, err := json.Marshal(&d)
	helper.CheckErr(err)
	ep.n.Publish(s, p)
}

