package study

import "math"

// ScheduledResult holds the calculated next state for a card
type ScheduledResult struct {
	Interval    int     // Days until next review
	EaseFactor  float64 // The multiplier
	Repetitions int     // Consecutive successful reviews
}

// CalculateNextReview applies the SM-2 algorithm.
// currentInterval: days since last review
// currentEase: the card's "stickiness" (starts at 2.5)
// repetitions: consecutive successes
// rating: 1 (Again), 2 (Hard), 3 (Good), 4 (Easy)
func CalculateNextReview(currentInterval int, currentEase float64, repetitions int, rating Rating) ScheduledResult {
	if rating == RatingAgain {
		return ScheduledResult{
			Interval:    0,           // Reset to 0 days (show again today/tomorrow)
			EaseFactor:  currentEase, // Don't punish ease too much on lapses
			Repetitions: 0,           // Reset streak
		}
	}

	// 1. Update Ease Factor
	// Formula: EF' = EF + (0.1 - (5-q) * (0.08 + (5-q)*0.02))
	// q = rating
	newEase := currentEase + (0.1 - (5.0-float64(rating))*(0.08+(5.0-float64(rating))*0.02))
	if newEase < 1.3 {
		newEase = 1.3 // Minimum floor
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
		// Interval[n] = Interval[n-1] * EaseFactor
		newInterval = int(math.Round(float64(currentInterval) * newEase))
	}

	return ScheduledResult{
		Interval:    newInterval,
		EaseFactor:  newEase,
		Repetitions: newRepetitions,
	}
}
