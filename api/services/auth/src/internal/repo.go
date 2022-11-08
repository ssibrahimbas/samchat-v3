package internal

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	c   *mongo.Collection
	ctx context.Context
}

type RepoParams struct {
	C   *mongo.Collection
	Ctx context.Context
}

func NewRepo(p *RepoParams) *Repo {
	return &Repo{
		c:   p.C,
		ctx: p.Ctx,
	}
}

