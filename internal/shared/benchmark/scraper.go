package benchmark

import "github.com/PuerkitoBio/goquery"

type BenchmarkScraper interface {
	Url() string
	GetGpuComparison(body *goquery.Selection) (string, GpuComparison)
}
