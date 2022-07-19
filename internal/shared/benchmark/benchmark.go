package benchmark

type BenchmarkResultType string

const (
	Fps    BenchmarkResultType = "FPS"
	Rating BenchmarkResultType = "Rating"
)

type BenchmarkDefinition struct {
	Name string
	Type BenchmarkResultType
}

type BenchmarkValue = float64

type GpuName = string

type GpuWithBenchmarks struct {
	Name                  GpuName
	BenchmarkResultValues []BenchmarkValue
}

type Url = string

type GpuComparison struct {
	Source               Url
	Title                string
	BenchmarkDefinitions []BenchmarkDefinition
	GpusWithBenchmarks   []GpuWithBenchmarks
}
