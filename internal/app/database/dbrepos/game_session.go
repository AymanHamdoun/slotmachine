package dbrepos

import (
	"context"
	"database/sql"
	"go-backend/internal/app/database"
	"go-backend/internal/app/database/dbmodels"
	"time"
)

type GameSessionParams struct {
	Name    string
	Token   string
	Credits int32
}
type GameSession struct {
	ID int
	GameSessionParams
	CreatedAt time.Time
}

type GameSessionRepo struct{}

func NewGameSessionRepo() GameSessionRepo {
	return GameSessionRepo{}
}

func (r GameSessionRepo) Create(ctx context.Context, session GameSessionParams) (GameSession, error) {
	db, err := database.GetDB(ctx)
	if err != nil {
		return GameSession{}, err
	}

	queries := dbmodels.New(db)

	result, creationErr := queries.CreateGameSession(ctx, dbmodels.CreateGameSessionParams{
		Name:    sql.NullString{String: session.Name},
		Token:   session.Token,
		Credits: session.Credits,
	})
	if creationErr != nil {
		return GameSession{}, creationErr
	}

	var sessionID int64
	sessionID, err = result.LastInsertId()
	if err != nil {
		return GameSession{}, err
	}

	return GameSession{
		ID: int(sessionID),
		GameSessionParams: GameSessionParams{
			Name:    session.Name,
			Token:   session.Token,
			Credits: session.Credits,
		},
	}, nil
}
