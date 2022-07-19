package tomshardware

import (
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"

	"github.com/Tobias-Belch/gpu-comparator/internal/shared/benchmark"
)

type TomsHardwareBenchmarkScraper struct {
}

func (scraper TomsHardwareBenchmarkScraper) Url() string {
	return "https://www.tomshardware.com/reviews/gpu-hierarchy,4388.html"
}

func (scraper TomsHardwareBenchmarkScraper) GetGpuComparison(body *goquery.Selection) (string, benchmark.GpuComparison) {
	h2s := body.Find("h2[id=gpu-benchmarks-ranking-2022]")

	if len(h2s.Nodes) > 0 {
		resultType := benchmark.Fps
		title := h2s.Text()

		table := h2s.SiblingsFiltered(".widthsetter").First().
			ChildrenFiltered(".articletable").First().
			ChildrenFiltered("table").First()

		benchmarkResultNames := table.
			ChildrenFiltered("thead").
			ChildrenFiltered("tr").
			ChildrenFiltered("th").
			Slice(1, 4).
			Map(func(i int, selection *goquery.Selection) string { return selection.Text() })

		var benchmarkDefinitions []benchmark.BenchmarkDefinition
		for _, name := range benchmarkResultNames {
			benchmarkDefinitions = append(benchmarkDefinitions, benchmark.BenchmarkDefinition{Name: name, Type: resultType})
		}

		var gpusWithBenchmarks []benchmark.GpuWithBenchmarks

		table.
			ChildrenFiltered("tbody").
			ChildrenFiltered("tr").
			Each(func(_ int, row *goquery.Selection) {
				gpuBenchmarks := extractGpuBenchmarks(row)

				gpusWithBenchmarks = append(gpusWithBenchmarks, gpuBenchmarks)
			})

		return "", benchmark.GpuComparison{Source: scraper.Url(), Title: title, BenchmarkDefinitions: benchmarkDefinitions, GpusWithBenchmarks: gpusWithBenchmarks}
	}

	var emptyComparison benchmark.GpuComparison
	return "Article not found!", emptyComparison
}

func extractGpuBenchmarks(row *goquery.Selection) benchmark.GpuWithBenchmarks {
	var benchmarkValues []benchmark.BenchmarkValue

	cells := row.Children()

	name := cells.First().Text()

	re := regexp.MustCompile(`\(([\d.]*)fps\)`)

	cells.Slice(1, 4).
		Each(func(_ int, cell *goquery.Selection) {
			matches := re.FindStringSubmatch(cell.Text())

			if len(matches) > 1 {
				if value, err := strconv.ParseFloat(matches[1], 64); err == nil {
					benchmarkValues = append(benchmarkValues, value)
				}
			}
		})

	return benchmark.GpuWithBenchmarks{Name: name, BenchmarkResultValues: benchmarkValues}
}
