package metrics

import (
	"sync/atomic"
)

func getOrRegisterRuntimeHistogram(name string, scale float64, r Registry) *runtimeHistogram {
	if r == nil {
		r = DefaultRegistry
	}
	constructor := func() Histogram { return newRuntimeHistogram(scale) }
	return r.GetOrRegister(name, constructor).(*runtimeHistogram)
}

// runtimeHistogram wraps a runtime/metrics histogram.
type runtimeHistogram struct {
	v           atomic.Value
	scaleFactor float64
}

func newRuntimeHistogram(scale float64) *runtimeHistogram {
	h := &runtimeHistogram{scaleFactor: scale}
	return h
}

// func (h *runtimeHistogram) update(mh *metrics.Float64Histogram) {
// 	if mh == nil {
// 		// The update value can be nil if the current Go version doesn't support a
// 		// requested metric. It's just easier to handle nil here than putting
// 		// conditionals everywhere.
// 		return
// 	}

// 	s := runtimeHistogramSnapshot{
// 		Counts:  make([]uint64, len(mh.Counts)),
// 		Buckets: make([]float64, len(mh.Buckets)),
// 	}
// 	copy(s.Counts, mh.Counts)
// 	copy(s.Buckets, mh.Buckets)
// 	for i, b := range s.Buckets {
// 		s.Buckets[i] = b * h.scaleFactor
// 	}
// 	h.v.Store(&s)
// }

// func (h *runtimeHistogram) load() *runtimeHistogramSnapshot {
// 	return h.v.Load().(*runtimeHistogramSnapshot)
// }

func (h *runtimeHistogram) Clear() {
	panic("runtimeHistogram does not support Clear")
}
func (h *runtimeHistogram) Update(int64) {
	panic("runtimeHistogram does not support Update")
}
func (h *runtimeHistogram) Sample() Sample {
	return NilSample{}
}

// Snapshot returns a non-changing cop of the histogram.
func (h *runtimeHistogram) Snapshot() Histogram {
	return NilHistogram{}
}

// Count returns the sample count.
func (h *runtimeHistogram) Count() int64 {
	return 0
}

// Mean returns an approximation of the mean.
func (h *runtimeHistogram) Mean() float64 {
	return 0
}

// StdDev approximates the standard deviation of the histogram.
func (h *runtimeHistogram) StdDev() float64 {
	return 0
}

// Variance approximates the variance of the histogram.
func (h *runtimeHistogram) Variance() float64 {
	return 0
}

// Percentile computes the p'th percentile value.
func (h *runtimeHistogram) Percentile(p float64) float64 {
	return 0
}

// Percentiles computes all requested percentile values.
func (h *runtimeHistogram) Percentiles(ps []float64) []float64 {
	return []float64{}
}

// Max returns the highest sample value.
func (h *runtimeHistogram) Max() int64 {
	return 0
}

// Min returns the lowest sample value.
func (h *runtimeHistogram) Min() int64 {
	return 0
}

// Sum returns the sum of all sample values.
func (h *runtimeHistogram) Sum() int64 {
	return 0
}
