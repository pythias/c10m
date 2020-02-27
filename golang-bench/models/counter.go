package models

type Stats struct {
	ConnectionPerSecond float64
	TimePerConnection   float64
	Min                 float64
	Mean                float64
	Stdev               float64
	Median              float64
	Max                 float64
}

type Counter struct {
	Stats

	Id        int
	Connected int64
	Failed    int64
	Closed    int64
	Sent      int64
	Received  int64
}
