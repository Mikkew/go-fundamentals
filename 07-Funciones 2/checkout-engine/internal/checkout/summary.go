package checkout

import (
	"fmt"
	"time"
)

func PrintHeader(title string) {
	fmt.Printf("\n=== %s ===\n", title)
}

func PrintDivider() {
	fmt.Println("————————————————————————————————————————————————————————")
}

func PrintKV(key string, value any) {
	fmt.Printf("%-30s : %v\n", key, value)
}

func PrintItems(key string, items []Item) {
	fmt.Printf("%-30s : ", key)
	fmt.Printf("[\n")

	for i, item := range items {
		fmt.Printf("%-32s  {\n", "")
		fmt.Printf("%-32s    \"SKU\": \"%s\",\n", "", item.SKU)
		fmt.Printf("%-32s    \"Name\": \"%s\",\n", "", item.Name)
		fmt.Printf("%-32s    \"Price\": %s,\n", "", StringUSD(item.Price))
		fmt.Printf("%-32s    \"Qty\": %d\n", "", item.Qty)

		if i < len(items)-1 {
			fmt.Printf("%-32s  },\n", "")
		} else {
			fmt.Printf("%-32s  }\n", "")
		}
	}

	fmt.Printf("%-33s]\n", "")
}

func PrintKV2(key string, value any, extra any) {
	fmt.Printf("%-30s : %v %v\n", key, value, extra)
}

func Track(name string) func() {
	start := time.Now()
	return func() {
		PrintKV("Time:"+name, time.Since(start))
	}
}
