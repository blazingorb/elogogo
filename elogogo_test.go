package elogogo

import (
	"testing"
)

func TestCalculateKA(t *testing.T) {
	var testCases = []struct {
		rating    float64
		expectedK float64
		expectedA float64
		name      string
	}{
		{
			rating:    100,
			expectedK: 116,
			expectedA: 200,
			name:      "Rating 100",
		},
		{
			rating:    200,
			expectedK: 110,
			expectedA: 195,
			name:      "Rating 200",
		},
		{
			rating:    300,
			expectedK: 105,
			expectedA: 190,
			name:      "Rating 300",
		},
		{
			rating:    400,
			expectedK: 100,
			expectedA: 185,
			name:      "Rating 400",
		},
		{
			rating:    500,
			expectedK: 95,
			expectedA: 180,
			name:      "Rating 500",
		},
		{
			rating:    600,
			expectedK: 90,
			expectedA: 175,
			name:      "Rating 600",
		},
		{
			rating:    700,
			expectedK: 85,
			expectedA: 170,
			name:      "Rating 700",
		},
		{
			rating:    800,
			expectedK: 80,
			expectedA: 165,
			name:      "Rating 800",
		},
		{
			rating:    900,
			expectedK: 75,
			expectedA: 160,
			name:      "Rating 900",
		},
		{
			rating:    1000,
			expectedK: 70,
			expectedA: 155,
			name:      "Rating 1000",
		},
		{
			rating:    1100,
			expectedK: 65,
			expectedA: 150,
			name:      "Rating 1100",
		},
		{
			rating:    1200,
			expectedK: 60,
			expectedA: 145,
			name:      "Rating 1200",
		},
		{
			rating:    1300,
			expectedK: 55,
			expectedA: 140,
			name:      "Rating 1300",
		},
		{
			rating:    1400,
			expectedK: 51,
			expectedA: 135,
			name:      "Rating 1400",
		},
		{
			rating:    1500,
			expectedK: 47,
			expectedA: 130,
			name:      "Rating 1500",
		},
		{
			rating:    1600,
			expectedK: 43,
			expectedA: 125,
			name:      "Rating 1600",
		},
		{
			rating:    1700,
			expectedK: 39,
			expectedA: 120,
			name:      "Rating 1700",
		},
		{
			rating:    1800,
			expectedK: 35,
			expectedA: 115,
			name:      "Rating 1800",
		},
		{
			rating:    1900,
			expectedK: 31,
			expectedA: 110,
			name:      "Rating 1900",
		},
		{
			rating:    2000,
			expectedK: 27,
			expectedA: 105,
			name:      "Rating 2000",
		},
		{
			rating:    2100,
			expectedK: 24,
			expectedA: 100,
			name:      "Rating 2100",
		},
		{
			rating:    2200,
			expectedK: 21,
			expectedA: 95,
			name:      "Rating 2200",
		},
		{
			rating:    2300,
			expectedK: 18,
			expectedA: 90,
			name:      "Rating 2300",
		},
		{
			rating:    2400,
			expectedK: 15,
			expectedA: 85,
			name:      "Rating 2400",
		},
		{
			rating:    2500,
			expectedK: 13,
			expectedA: 80,
			name:      "Rating 2500",
		},
		{
			rating:    2600,
			expectedK: 11,
			expectedA: 75,
			name:      "Rating 2600",
		},
		{
			rating:    2700,
			expectedK: 10,
			expectedA: 70,
			name:      "Rating 2700",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			k, a := CalcKA(tc.rating)
			if tc.expectedK != k {
				t.Errorf("K for Rating %v, expected %v got %v", tc.rating, tc.expectedK, k)
			}
			if tc.expectedA != a {
				t.Errorf("A for Rating %v, expected %v got %v", tc.rating, tc.expectedA, a)
			}
		})
	}
}
