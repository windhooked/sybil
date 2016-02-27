// +build !profile

package sybil

var PROFILER_ENABLED bool
var PROFILE *bool = &PROFILER_ENABLED

type NoProfile struct{}

func (p NoProfile) Start() ProfilerStart {
	return NoProfile{}
}
func (p NoProfile) Stop() {

}

var RUN_PROFILER = func() ProfilerStop {
	return NoProfile{}
}
