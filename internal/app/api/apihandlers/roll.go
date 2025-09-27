package apihandlers

import (
	"context"
	"fmt"
	"go-backend/internal/app/database/dbrepos"
	"go-backend/internal/app/services"
	"go-backend/internal/app/utils/requestutils"
	"go-backend/internal/app/utils/responseutils"
	"net/http"
)

func GetRollRoutes() []string {
	return []string{
		"/api/v1/roll",
	}
}

type RollInput struct{}

type RollResponse struct {
	Status     string   `json:"status"`
	SlotValues []string `json:"slotValues"`
	Credits    int32    `json:"credits"`
}

type RollHandler struct{}

// Symbol rewards mapping
var symbolRewards = map[string]int32{
	"cherry":     10,
	"lemon":      20,
	"orange":     30,
	"watermelon": 40,
}

func (h RollHandler) Serve(ctx context.Context, input RollInput, w http.ResponseWriter, r *http.Request) {
	token, err := requestutils.GetCookieValue(r, "token")
	if err != nil {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "unauthenticated",
			Message: "No game session token found",
		})
		return
	}

	repo := dbrepos.NewGameSessionRepo()
	session, err := repo.GetByToken(ctx, token)
	if err != nil {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "invalid-session",
			Message: "Game session is not valid",
		})
		return
	}

	// Check if user has enough credits to roll
	const rollCost int32 = 1
	if session.Credits < rollCost {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "insufficient-credits",
			Message: "Not enough credits to roll",
		})
		return
	}

	// Deduct the roll cost
	success, err := repo.SubtractCredits(ctx, token, rollCost)
	if err != nil || !success {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "error",
			Message: "Failed to deduct credits",
		})
		return
	}

	// Get the reroll chance based on current credits
	rerollChance := h.getRerollChance(session.Credits)
	fmt.Printf("RerollChance: %v\n", rerollChance)

	rollService := services.RollService{}

	roll := rollService.NewRoll(3, symbolRewards)
	roll.Roll(rerollChance)

	// Update credits with winnings if any
	winnings := roll.CalculateWinnings()
	updatedCredits := session.Credits - rollCost
	if winnings > 0 {
		success, err = repo.AddCredits(ctx, token, winnings)
		if err != nil || !success {
			// we couldn't add credits so we return an error without deducting rollcost.
			responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
				Status:  "error",
				Message: "Failed to add credits",
			})
			return
		}
		updatedCredits += winnings
	}

	responseutils.WriteJSON(ctx, w, RollResponse{
		Status:     "ok",
		SlotValues: roll.GetSlotValues(),
		Credits:    updatedCredits,
	})
}

// getRerollChance returns the reroll chance based on credits
func (h RollHandler) getRerollChance(credits int32) float32 {
	if credits < 40 {
		return 0.0 // No cheating under 40 credits
	} else if credits <= 60 {
		return 0.30 // 30% chance to reroll winning rolls
	}
	return 0.60 // 60% chance to reroll winning rolls above 60 credits
}
