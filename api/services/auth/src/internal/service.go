package internal

import (
	"github.com/ssibrahimbas/samchat-v3.auth/src/event_publisher"
	"github.com/ssibrahimbas/samchat-v3.auth/src/mapper"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/cipher"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/i18n"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/jwt"
)

type Srv struct {
	r    *Repo
	m    *mapper.Mapper
	i18n *i18n.I18n
	jwt  *jwt.Jwt
	ep     *event_publisher.Publisher
	cipher *cipher.Cipher
}

type SrvParams struct {
	Repo *Repo
	I18n *i18n.I18n
	Jwt  *jwt.Jwt
	EPub   *event_publisher.Publisher
	Cipher *cipher.Cipher
}

func NewSrv(p *SrvParams) *Srv {
	return &Srv{
		r:    p.Repo,
		i18n: p.I18n,
		jwt:  p.Jwt,
		m:    mapper.New(),
		ep:     p.EPub,
		cipher: p.Cipher,
	}
}
