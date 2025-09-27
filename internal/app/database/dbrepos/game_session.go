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

// gameSessionFromDBModel converts a database model to domain model
func gameSessionFromDBModel(dbSession dbmodels.GameSession) GameSession {
	result := GameSession{
		ID: int(dbSession.ID),
		GameSessionParams: GameSessionParams{
			Name:    dbSession.Name.String,
			Token:   dbSession.Token,
			Credits: dbSession.Credits,
		},
	}

	if dbSession.CreatedAt.Valid {
		result.CreatedAt = dbSession.CreatedAt.Time
	}

	return result
}

func (r GameSessionRepo) Create(ctx context.Context, session GameSessionParams) (GameSession, error) {
	db, err := database.GetDB(ctx)
	if err != nil {
		return GameSession{}, err
	}

	queries := dbmodels.New(db)

	result, creationErr := queries.CreateGameSession(ctx, dbmodels.CreateGameSessionParams{
		Name:    sql.NullString{String: session.Name, Valid: len(session.Name) > 0},
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

func (r GameSessionRepo) GetByToken(ctx context.Context, token string) (GameSession, error) {
	db, err := database.GetDB(ctx)
	if err != nil {
		return GameSession{}, err
	}

	queries := dbmodels.New(db)

	dbSession, err := queries.GetGameSessionByToken(ctx, token)
	if err != nil {
		return GameSession{}, err
	}

	return gameSessionFromDBModel(dbSession), nil
}

func (r GameSessionRepo) DeleteByToken(ctx context.Context, token string) (bool, error) {
	db, err := database.GetDB(ctx)
	if err != nil {
		return false, err
	}

	queries := dbmodels.New(db)

	result, err := queries.DeleteGameSessionByToken(ctx, token)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r GameSessionRepo) AddCredits(ctx context.Context, token string, credits int32) (bool, error) {
	db, err := database.GetDB(ctx)
	if err != nil {
		return false, err
	}

	queries := dbmodels.New(db)

	result, err := queries.AddCreditsToSession(ctx, dbmodels.AddCreditsToSessionParams{
		Credits: credits,
		Token:   token,
	})
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r GameSessionRepo) SubtractCredits(ctx context.Context, token string, credits int32) (bool, error) {
	db, err := database.GetDB(ctx)
	if err != nil {
		return false, err
	}

	queries := dbmodels.New(db)

	result, err := queries.SubtractCreditsFromSession(ctx, dbmodels.SubtractCreditsFromSessionParams{
		Credits:   credits,
		Token:     token,
		Credits_2: credits, // This ensures we don't go below 0
	})
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
