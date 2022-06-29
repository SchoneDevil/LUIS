package repo

import "user_service/pkg/postgres"

type Repo struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg: pg}
}
