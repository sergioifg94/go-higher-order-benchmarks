package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/tools/benchmark/parse"
)

func formatBenchmark() error {
	cmd := exec.Command("go", "test", "-benchmem=true", "-bench=.", "./pkg/enumerable/.")

	output, err := cmd.Output()
	if err != nil {
		if exitErrror, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("Error: %s", string(exitErrror.Stderr))
		}

		return err
	}

	benchmarkSet, err := parse.ParseSet(bytes.NewReader(output))
	if err != nil {
		return err
	}

	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write([]string{
		"name",
		"size",
		"ncalls",
		"nsperop",
		"n",
		"bperop",
		"allocs",
	})

	for name, benchmarks := range benchmarkSet {
		for _, benchmark := range benchmarks {
			bmData := parseName(name)

			writer.Write([]string{
				bmData.name,
				fmt.Sprintf("%d", bmData.size),
				fmt.Sprintf("%d", bmData.calls),
				fmt.Sprintf("%.6f", benchmark.NsPerOp),
				fmt.Sprintf("%d", benchmark.N),
				fmt.Sprintf("%d", benchmark.AllocedBytesPerOp),
				fmt.Sprintf("%d", benchmark.AllocsPerOp),
			})
		}
	}

	return nil
}

func main() {
	if err := formatBenchmark(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func parseName(fullName string) struct {
	name  string
	size  int64
	calls int64
} {
	sep := strings.Split(fullName, "/")
	left := sep[0]
	right := sep[1]

	rightSep := strings.Split(right, "-")
	sizeAndCalls := strings.Split(rightSep[0], ",")

	size, _ := strconv.ParseInt(sizeAndCalls[0], 10, 64)
	calls, _ := strconv.ParseInt(sizeAndCalls[1], 10, 64)

	return struct {
		name  string
		size  int64
		calls int64
	}{
		name:  left,
		size:  size,
		calls: calls,
	}
}
