package study

import "testing"

func TestCalculateNextReview(t *testing.T) {
	tests := []struct {
		name         string
		currInterval int
		currEase     float64
		reps         int
		rating       Rating
		wantInterval int
	}{
		{"Fail resets interval", 10, 2.5, 5, RatingAgain, 0},
		{"First success is 1 day", 0, 2.5, 0, RatingGood, 1},
		{"Second success is 6 days", 1, 2.5, 1, RatingGood, 6},
		{"Easy bonus", 6, 2.5, 2, RatingEasy, 21}, // 6 * (2.5 + bonus) approx
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateNextReview(tt.currInterval, tt.currEase, tt.reps, tt.rating)
			if got.Interval != tt.wantInterval {
				t.Errorf("got interval %d, want %d", got.Interval, tt.wantInterval)
			}
		})
	}
}
