package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	type row struct {
		xed          string
		conventional string
		filename     string
	}
	rows := make([]row, 0, len(avx512extensions))
	for xed, conventional := range avx512extensions {
		filename := conventional
		filename = strings.Replace(filename, "+AVX512VL", "", 1)
		filename = strings.Replace(filename, "+", "_", 1)
		filename = strings.ToLower(filename)
		rows = append(rows, row{
			xed:          xed,
			conventional: conventional,
			filename:     filename,
		})
	}
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].xed < rows[j].xed
	})
	fmt.Println(`"xed","conventional","filename"`)
	for _, row := range rows {
		fmt.Printf("%q,%q,%q\n", row.xed, row.conventional, row.filename)
	}
}
