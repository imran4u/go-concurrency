package main

import (
	"fmt"
	"slices"
)

type Product struct {
	name  string
	color string
	size  string
	price int
}

func main() {
	products := []Product{
		{
			"Long Sleeve Scoop Neck Cotton Tee",
			"White",
			"S",
			2500,
		},
		{
			"Long Sleeve Scoop Neck Cotton Tee",
			"White",
			"M",
			2500,
		},
		{
			"Long Sleeve Scoop Neck Cotton Tee",
			"White",
			"L",
			2500,
		},
		{
			"Long Sleeve Scoop Neck Cotton Tee",
			"White",
			"Med",
			2500,
		},
		{
			"Long Sleeve Scoop Neck Cotton Tee",
			"Blue",
			"L",
			2500,
		},
		{
			"Waxed Denim Slim Leg Jeans",
			"Blue",
			"8",
			19800,
		},
		{
			"Waxed Denim Slim Leg Jeans",
			"Blue",
			"8",
			19800,
		},
		{
			"Waxed Denim Slim Leg Jeans",
			"Blue",
			"10",
			19800,
		},
		{
			"Mera Tee",
			"Grey",
			"XS",
			22500,
		},
		{
			"Mera Tee",
			"Blue",
			"S",
			22500,
		},
		{
			"Mera Tee",
			"Grey",
			"S",
			22500,
		},
	}

	for _, p := range products {
		fmt.Println(p.name)
	}

	level1(products)
	fmt.Println("---- L2")

	level2(products, false)

	fmt.Println("---- L3")
	level2(products, true) // normalised
}

func level1(products []Product) {
	countMap := make(map[string]int)
	fmt.Println("Report by Product:")
	for _, p := range products {
		//countMap[p.name] = countMap[p.name] + 1
		countMap[p.name]++ // Thumb rule : x++ → valid but  y = x++ → invalid
	}

	for k, v := range countMap {
		fmt.Println(k, ":", v)
	}

}

func level2(products []Product, normalise bool) {
	// create map as per the out put.
	// sort the map keys so the you can print the value as per sortted values.
	output := make(map[string]map[string]map[string]int)

	for _, p := range products {
		if _, ok := output[p.name]; !ok {
			output[p.name] = make(map[string]map[string]int)
		}
		if _, ok := output[p.name][p.color]; !ok {
			output[p.name][p.color] = make(map[string]int)
		}
		output[p.name][p.color][p.size]++
	}

	productSortedKeys := sortedKeys(output)
	fmt.Println("Report by Product Color/Size:")

	for _, key := range productSortedKeys {
		colorMap := output[key]
		fmt.Println(key)
		colorSortedKyes := sortedKeys(colorMap)
		for _, color := range colorSortedKyes {
			fmt.Println(color)
			sizeMap := colorMap[color]
			// sizeSortedKey := sortedKeys(sizeMap)
			for key, size := range sizeMap {
				if normalise {
					key = normalizeSize(key)
				}
				fmt.Println(key, ":", size)
			}
		}
	}

}

// important : use generic here
func sortedKeys[T any](m map[string]T) []string {
	result := make([]string, 0)
	for k := range m {
		result = append(result, k)
	}
	//

	// result = sort.StringSlice(result) // old way
	// sort.Strings(result) // old way

	slices.Sort(result)
	return result
}

func normalizeSize(size string) string {
	switch size {
	case "XS":
		return "XS"
	case "S":
		return "S"
	case "M", "Med", "8", "10":
		return "M"
	case "L":
		return "L"
	case "XL":
		return "XL"
	default:
		return size
	}
}
