package services

import (
	"go-backend/internal/app/utils/maputils"
	"math/rand"
)

type RollService struct{}

type roll struct {
	SlotCount     int
	Slots         []string
	SymbolRewards map[string]int32
}

func (s RollService) NewRoll(slotCount int, symbolRewards map[string]int32) roll {
	return roll{
		SlotCount:     slotCount,
		Slots:         make([]string, 0, slotCount),
		SymbolRewards: symbolRewards,
	}
}

func (r *roll) GetSlotValues() []string {
	return r.Slots
}

func (r *roll) Roll(rerollChance float32) {
	r.FillSlotsRandomly()
	if r.CalculateWinnings() > 0 && rerollChance > 0 && rand.Float32() <= rerollChance {
		r.FillSlotsRandomly()
	}
}

// FillSlotsRandomly generates 3 random slot values
func (r *roll) FillSlotsRandomly() {
	possibleSlotValues := maputils.Keys(r.SymbolRewards)
	result := make([]string, r.SlotCount)
	for i := 0; i < r.SlotCount; i++ {
		result[i] = possibleSlotValues[rand.Intn(len(possibleSlotValues))]
	}

	r.Slots = result
}

// CalculateWinnings calculates the reward if all n slots match
func (r *roll) CalculateWinnings() int32 {
	if len(r.Slots) == 0 {
		return 0
	}

	lastSlot := ""
	for i, slot := range r.Slots {
		if i == 0 {
			lastSlot = slot
			continue
		}

		if lastSlot != slot {
			return 0
		}
	}

	if lastSlot == "" {
		return 0
	}

	return r.SymbolRewards[lastSlot]
}
