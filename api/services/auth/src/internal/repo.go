package internal

import (
	"context"

	"github.com/ssibrahimbas/samchat-v3.shared/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	c   *mongo.Collection
	ctx context.Context
	db  *db.MongoDB
}

type RepoParams struct {
	C   *mongo.Collection
	Ctx context.Context
	Db  *db.MongoDB
}

func NewRepo(p *RepoParams) *Repo {
	return &Repo{
		c:   p.C,
		ctx: p.Ctx,
		db:  p.Db,
	}
}

