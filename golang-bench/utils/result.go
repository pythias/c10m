package utils

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"sync"
)

func GetResult(m sync.Map) (float64, float64, string) {
	var min, max, sum, mean, stdev, median float64

	var milliseconds []float64
	m.Range(func(_, v interface{}) bool {
		f, _ := convertToFloat64(v)
		milliseconds = append(milliseconds, f)
		return true
	})

	sort.Float64s(milliseconds)
	min = milliseconds[0]
	max = milliseconds[len(milliseconds)-1]

	for _, v := range milliseconds {
		sum += v
	}
	mean = sum / float64(len(milliseconds))

	for _, v := range milliseconds {
		stdev += math.Pow(v-mean, 2)
	}
	stdev = math.Sqrt(stdev / float64(len(milliseconds)))

	if len(milliseconds)%2 == 0 {
		median = (milliseconds[len(milliseconds)/2] + milliseconds[len(milliseconds)/2+1]) / 2
	} else {
		median = milliseconds[1+len(milliseconds)/2]
	}

	cps := (1000 * float64(len(milliseconds))) / sum
	tpc := sum / float64(len(milliseconds))

	return cps, tpc, fmt.Sprintf("%7.2f %7.2f %7.2f %7.2f %7.2f", min, mean, stdev, median, max)
}

func convertToFloat64(v interface{}) (float64, error) {
	switch f := v.(type) {
	case float64:
		return f, nil
	case float32:
		return float64(f), nil
	case int64:
		return float64(f), nil
	default:
		return 0, errors.New("getFloat: unknown value is of incompatible type")
	}
}
