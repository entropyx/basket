package basket

import (
	"github.com/entropyx/basket/modules/apriori"
)

func unique(x []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range x {
		if encountered[x[v]] == true {
		} else {
			encountered[x[v]] = true
			result = append(result, x[v])
		}
	}
	return result
}

func removeElement(element string, array []string) []string {
	var out []string
	l1 := len(array)
	for i := 0; i < l1; i++ {
		if array[i] != element {
			out = append(out, array[i])
		}
	}
	return out
}

func Analyze(items []string, combinations [][]string) []string {
	var list []string
	for i := 0; i < len(combinations); i++ {
		if apriori.Subset(items, combinations[i]) {
			list = append(list, combinations[i]...)
		}
	}
	list = unique(list)
	for j := 0; j < len(items); j++ {
		list = removeElement(items[j], list)
	}
	return list
}
