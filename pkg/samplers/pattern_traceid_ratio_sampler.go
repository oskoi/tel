package samplers

import (
	"github.com/gobwas/glob"
	"go.opentelemetry.io/otel/sdk/trace"
)

type PatternTraceOptions struct {
	Pattern     string
	Fraction    float64
	TrackStatus bool
}

type patternSampler struct {
	glob    glob.Glob
	sampler trace.Sampler
}

var neverSampler = trace.NeverSample()

var PatternTraceDefault = PatternTraceOptions{
	Pattern:     "**",
	Fraction:    0.1,
	TrackStatus: true,
}

// PatternTraceIDRatioBased samples traces by pattern of span name
// Example:
// tele, close := tel.New(
//
//		context.Background(),
//		tel.GetConfigFromEnv(),
//		tel.WithTraceSampler(
//			samplers.PatternTraceIDRatioBased(
//				samplers.PatternTraceOptions{
//					Pattern:  "GET:/foobar**",
//					Fraction: 1.0,
//				},
//				samplers.PatternTraceDefault,
//			),
//		),
//	)
//	defer close()
func PatternTraceIDRatioBased(options ...PatternTraceOptions) trace.Sampler {
	patternSamplers := make([]patternSampler, 0, len(options))
	for _, opt := range options {
		g := glob.MustCompile(opt.Pattern)

		sampler := trace.TraceIDRatioBased(opt.Fraction)
		if opt.TrackStatus {
			sampler = StatusTraceIDRatioBased(opt.Fraction)
		}

		patternSamplers = append(
			patternSamplers,
			patternSampler{glob: g, sampler: sampler},
		)
	}

	return patternTraceIDRatioSampler{
		patternSamplers: patternSamplers,
		description:     "PatternTraceIDRatioBased",
	}
}

type patternTraceIDRatioSampler struct {
	patternSamplers []patternSampler
	description     string
}

func (s patternTraceIDRatioSampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	for _, ps := range s.patternSamplers {
		if ps.glob.Match(p.Name) {
			return ps.sampler.ShouldSample(p)
		}
	}

	return neverSampler.ShouldSample(p)
}

func (ts patternTraceIDRatioSampler) Description() string {
	return ts.description
}
