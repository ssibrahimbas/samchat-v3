package internal

import (
	"github.com/ssibrahimbas/samchat-v3.auth/src/config"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/http"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/i18n"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/jwt"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/validator"
)

type Handler struct {
	s    *Srv
	c    *config.App
	h    *http.Client
	v    *validator.Validator
	i18n *i18n.I18n
	jwt  *jwt.Jwt
}

type HandlerParams struct {
	Srv   *Srv
	Http  *http.Client
	I18n  *i18n.I18n
	Jwt   *jwt.Jwt
	Valid *validator.Validator
	Cnf   *config.App
}

func NewHandler(p *HandlerParams) *Handler {
	return &Handler{
		s:    p.Srv,
		h:    p.Http,
		i18n: p.I18n,
		jwt:  p.Jwt,
		v:    p.Valid,
		c:    p.Cnf,
	}
}

func (h *Handler) InitAllVersions() {
	h.initV1()
}

func (h *Handler) initV1() {}
