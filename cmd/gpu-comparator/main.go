package main

import (
	"fmt"

	"github.com/gocolly/colly"

	"github.com/Tobias-Belch/gpu-comparator/internal/shared/benchmark"
	"github.com/Tobias-Belch/gpu-comparator/internal/tomshardware"
)

func main() {
	var benchmarkScraper benchmark.BenchmarkScraper = tomshardware.TomsHardwareBenchmarkScraper{}

	c := colly.NewCollector()

	c.OnHTML("body", func(body *colly.HTMLElement) {
		fmt.Println(benchmarkScraper.GetGpuComparison(body.DOM))
	})

	c.Visit(benchmarkScraper.Url())
}
