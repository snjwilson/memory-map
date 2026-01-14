package study

import "math"

// ScheduledResult holds the calculated next state for a card
type ScheduledResult struct {
	Interval    int     // Days until next review
	EaseFactor  float64 // The multiplier
	Repetitions int     // Consecutive successful reviews
}

func CalculateNextReview(currentInterval int, currentEase float64, repetitions int, rating Rating) ScheduledResult {
	// Handle "Again" (Fail)
	if rating == RatingAgain {
		return ScheduledResult{
			Interval:    1,                               // Show again in 1 day
			EaseFactor:  math.Max(1.3, currentEase-0.20), // Slightly penalize ease on fail
			Repetitions: 0,
		}
	}

	// Map your 1-3 scale to the SM-2 3-5 scale logic
	// Frontend 2 (Hard) -> SM-2 3 (Difficult but correct)
	// Frontend 3 (Easy) -> SM-2 5 (Perfect recall)
	var sm2Weight float64
	if rating == RatingHard {
		sm2Weight = 3.0
	} else {
		sm2Weight = 5.0 // RatingEasy
	}

	// 1. Update Ease Factor using the mapped weight
	newEase := currentEase + (0.1 - (5.0-sm2Weight)*(0.08+(5.0-sm2Weight)*0.02))
	if newEase < 1.3 {
		newEase = 1.3
	}

	// 2. Update Repetitions
	newRepetitions := repetitions + 1

	// 3. Calculate New Interval
	var newInterval int
	if newRepetitions == 1 {
		newInterval = 1
	} else if newRepetitions == 2 {
		newInterval = 6
	} else {
		newInterval = int(math.Ceil(float64(currentInterval) * newEase))
	}

	return ScheduledResult{
		Interval:    newInterval,
		EaseFactor:  newEase,
		Repetitions: newRepetitions,
	}
}
