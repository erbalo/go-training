package location_test

import (
	"chapter5/src/domain"
	"chapter5/src/infra/services/location"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Athens = domain.City{
	Name: "Athens",
	Location: domain.Location{
		Lat: 37.983972,
		Lon: 23.727806,
	},
}

var Amsterdam = domain.City{
	Name: "Amsterdam",
	Location: domain.Location{
		Lat: 52.366667,
		Lon: 4.9,
	},
}

var Berlin = domain.City{
	Name: "Berlin",
	Location: domain.Location{
		Lat: 52.516667,
		Lon: 13.388889,
	},
}

func TestDistance(test *testing.T) {
	type args struct {
		startedCity domain.City
		finalCity   domain.City
	}

	testCases := []struct {
		name             string
		args             args
		expectedDistance float64
	}{
		{
			name: "Athens to Amsterdam",
			args: args{
				startedCity: Athens,
				finalCity:   Amsterdam,
			},
			expectedDistance: 2163.2310285824487,
		},
		{
			name: "Amsterdam to Berlin",
			args: args{
				startedCity: Amsterdam,
				finalCity:   Berlin,
			},
			expectedDistance: 575.2949643958797,
		},
		{
			name: "Berlin to Athens",
			args: args{
				startedCity: Berlin,
				finalCity:   Athens,
			},
			expectedDistance: 1803.1087879059255,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			service := location.New()

			startLocation := testCase.args.startedCity.Location
			endLocation := testCase.args.finalCity.Location

			if got := service.Distance(startLocation, endLocation); got != testCase.expectedDistance {
				assert.InDelta(t, testCase.expectedDistance, got, 0.002)
			}
		})
	}
}

var res float64

func BenchmarkAthensToAmsterdam(b *testing.B) {
	service := location.New()

	// Any initialization code comes here
	var res1 float64

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res1 = service.Distance(Athens.Location, Amsterdam.Location)
	}

	res = res1
}
