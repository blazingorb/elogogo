package elogogo

import (
	"math"
)

const (
	AMax      = 200
	AMin      = 70
	RatingMin = 100
	RatingMax = 2700

	Win  = 1
	Lose = 0
	Tie  = .5
)

var (
	kRange = []float64{116, 116, 110, 105, 100, 95, 90, 85, 80, 75, 70, 65, 60, 55, 51, 47, 43, 39, 35, 31, 27, 24, 21, 18, 15, 13, 11, 10}
)

func init() {
	if (len(kRange)-1)*100 != RatingMax {
		panic("kRange array length incorrect")
	}
}

// CalcKA get the K and A constants for a specified rating based on the EGF rating table
func CalcKA(rating float64) (k float64, a float64) {
	ratingInt := int(rating)
	if ratingInt < RatingMin {
		return kRange[1], AMax
	} else if ratingInt >= RatingMax {
		return kRange[len(kRange)-1], AMin
	}

	ratio := ratingInt / 100
	k = kRange[ratio]
	a = float64(AMax - ((ratio - 1) * 5))
	return
}

// CalcOffer ratingB is higher than ratingA
func CalcOffer(ratingB, ratingA float64, scoreB, scoreA float64) (selectedK float64, offerB float64, offerA float64) {
	// use lower Rating to get K/A
	var A float64
	selectedK, A = CalcKA(ratingA)

	probabilityA := 1 / (math.Exp((ratingB-ratingA)/A) + 1)
	probabilityB := 1 - probabilityA

	offerB = (selectedK * (scoreB - probabilityB))
	offerA = (selectedK * (scoreA - probabilityA))
	return
}
