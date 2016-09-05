package basket

import (
	"github.com/entropyx/basket-analysis/modules/apriori"
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

func BasketAnalysis(items []string, combinations [][]string) []string {
	var list1 []string
	for i := 0; i < len(combinations); i++ {
		if apriori.Subset(items, combinations[i]) {
			list1 = append(list1, combinations[i]...)
		}
	}
	list1 = unique(list1)
	for j := 0; j < len(items); j++ {
		list1 = removeElement(items[j], list1)
	}
	return list1
}
