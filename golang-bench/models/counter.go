package models

type Stats struct {
    Qps      float64
    Duration float64
    Min      float64
    Mean     float64
    Stdev    float64
    Median   float64
    Max      float64
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
