package app

import (
	"context"
	"log"

	"github.com/ssibrahimbas/samchat-v3.profile/src/config"
	"github.com/ssibrahimbas/samchat-v3.profile/src/event_handler"
	"github.com/ssibrahimbas/samchat-v3.profile/src/event_publisher"
	"github.com/ssibrahimbas/samchat-v3.profile/src/internal"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/cipher"
	cnf "github.com/ssibrahimbas/samchat-v3.shared/pkg/config"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/db"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/helper"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/http"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/i18n"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/jwt"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/nats"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/validator"
)

type App struct {
	Cnf    *config.App
	Db     *db.MongoDB
	I18n   *i18n.I18n
	Http   *http.Client
	Cipher *cipher.Cipher
	Jwt    *jwt.Jwt
	Val    *validator.Validator
	Nats   *nats.Conn
	EPub   *event_publisher.Publisher
	EHnd   *event_handler.Handler
	Repo   *internal.Repo
	Srv    *internal.Srv
	Hnd    *internal.Handler
}

func New() *App {
	return &App{}
}

func (a *App) Init() *App {
	a.loadConfig()
	a.loadDb()
	a.loadI18n()
	a.loadHttp()
	a.loadCipher()
	a.loadJWT()
	a.loadValidator()
	a.loadNats()
	a.loadEvents()
	a.loadInternal()
	return a
}

func (a *App) Serve() {
	log.Fatal(a.Http.Listen(a.Cnf.Host + a.Cnf.Port))
}

func (a *App) loadConfig() {
	cnfg := config.App{}
	cnf.LoadConfig(&cnfg)
	a.Cnf = &cnfg
}

func (a *App) loadDb() {
	uri := db.CalcMongoUri(db.UriParams{
		Host: a.Cnf.Db.Host,
		Port: a.Cnf.Db.Port,
		User: a.Cnf.Db.Usr,
		Pass: a.Cnf.Db.Pw,
		Db:   a.Cnf.Db.Name,
	})
	d, err := db.NewMongo(uri, a.Cnf.Db.Name)
	helper.CheckErr(err)
	a.Db = d
}

func (a *App) loadI18n() {
	i := i18n.New(a.Cnf.I18n.Fallback)
	i.LoadLanguages(a.Cnf.I18n.LocaleDir, a.Cnf.I18n.Locales...)
	a.I18n = i
}

func (a *App) loadHttp() {
	a.Http = http.New(a.I18n)
}

func (a *App) loadCipher() {
	a.Cipher = cipher.New()
}

func (a *App) loadJWT() {
	a.Jwt = jwt.New(a.Cnf.JWT.Secret)
}

func (a *App) loadValidator() {
	val := validator.New(a.I18n)
	val.ConnectCustom()
	val.RegisterTagName()
	a.Val = val
}

func (a *App) loadNats() {
	a.Nats = nats.New()
}

func (a *App) loadEvents() {
	a.EPub = event_publisher.New(a.Nats)
}

func (a *App) loadInternal() {
	a.Repo = internal.NewRepo(&internal.RepoParams{
		C:   a.Db.GetCollection(a.Cnf.Db.Collection),
		Db:  a.Db,
		Ctx: context.TODO(),
	})
	a.Srv = internal.NewSrv(&internal.SrvParams{
		Repo:   a.Repo,
		I18n:   a.I18n,
		Jwt:    a.Jwt,
		EPub:   a.EPub,
		Cipher: a.Cipher,
	})
	a.Hnd = internal.NewHandler(&internal.HandlerParams{
		Srv:   a.Srv,
		Cnf:   a.Cnf,
		I18n:  a.I18n,
		Valid: a.Val,
		Http:  a.Http,
		Jwt:   a.Jwt,
	})
	a.Hnd.InitAllVersions()
}
