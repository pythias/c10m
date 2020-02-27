package utils

import (
	"errors"
	tm "github.com/buger/goterm"
	"github.com/pythias/c10m/golang-bench/models"
	"math"
	"sort"
	"sync"
)

func OutputResults(header string, results map[int]models.Counter) {
	if len(results) == 0 {
		tm.Clear()
		tm.MoveCursor(1, 1)
		_, _ = tm.Println(header)
		tm.Flush()
		return
	}

	stats0 := models.Stats{
		ConnectionPerSecond: 0,
		TimePerConnection:   0,
		Min:                 math.MaxFloat64,
		Mean:                0,
		Stdev:               0,
		Median:              0,
		Max:                 0,
	}
	counter0 := models.Counter{
		Stats:     stats0,
		Id:        0,
		Connected: 0,
		Failed:    0,
		Closed:    0,
		Sent:      0,
		Received:  0,
	}

	means := []float64{}
	for _, counter := range results {
		counter0.Connected += counter.Connected
		counter0.Failed += counter.Failed
		counter0.Closed += counter.Closed
		counter0.Sent += counter.Sent
		counter0.Received += counter.Received

		counter0.Min = math.Min(counter0.Min, counter.Min)
		counter0.Max = math.Max(counter0.Max, counter.Max)
		means = append(means, counter.Mean)
	}

	meanStats, _ := CalculateStats(means)
	counter0.Mean = meanStats.Mean
	counter0.Stdev = meanStats.Stdev
	counter0.Median = meanStats.Median

	tm.Clear()
	tm.MoveCursor(1, 1)
	_, _ = tm.Println(header)
	_, _ = tm.Printf("Completed connections: %d\n", counter0.Connected+counter0.Failed)
	_, _ = tm.Printf("Succeed connections: %d (%0.2f%%)\n", counter0.Connected, 100*float64(counter0.Connected)/float64(counter0.Connected+counter0.Failed))
	_, _ = tm.Println("Failed connections:", counter0.Failed)
	_, _ = tm.Printf("Connections per second: %0.2f\n", counter0.ConnectionPerSecond)
	_, _ = tm.Printf("Time per connection: %0.2f [ms]\n", counter0.TimePerConnection)
	_, _ = tm.Println("Connection Times (ms)")
	_, _ = tm.Println("             min    mean [+/-sd]  median     max")
	_, _ = tm.Printf("Connect  %7.2f %7.2f %7.2f %7.2f %7.2f\n", counter0.Min, counter0.Median, counter0.Stdev, counter0.Median, counter0.Max)
	tm.Flush()
}

func GetResult(m sync.Map) (models.Stats, error) {
	milliseconds := syncToNormal(m)
	return CalculateStats(milliseconds)
}

func CalculateStats(milliseconds []float64) (models.Stats, error) {
	if len(milliseconds) < 2 {
		return models.Stats{}, errors.New("Not enough values")
	}

	var min, max, sum, mean, stdev, median float64
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

	result := models.Stats{
		ConnectionPerSecond: cps,
		TimePerConnection:   tpc,
		Min:                 min,
		Mean:                mean,
		Stdev:               stdev,
		Median:              median,
		Max:                 max,
	}
	return result, nil
}

func syncToNormal(m sync.Map) []float64 {
	var milliseconds []float64
	m.Range(func(_, v interface{}) bool {
		f, _ := convertToFloat64(v)
		milliseconds = append(milliseconds, f)
		return true
	})

	return milliseconds
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
