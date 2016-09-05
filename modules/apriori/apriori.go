package apriori

import "sort"

type Transaction struct {
	Pattern []string
	Support float64
}

type Transactions []*Transaction

func (slice Transactions) Len() int {
	return len(slice)
}

func (slice Transactions) Less(i, j int) bool {
	return slice[i].Support > slice[j].Support
}

func (slice Transactions) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice Transactions) Patterns() [][]string {
	var patterns [][]string
	for i := 0; i < len(slice); i++ {
		patterns = append(patterns, slice[i].Pattern)
	}
	return patterns
}

func Subset(first, second []string) bool {
	set := make(map[string]int)
	for _, value := range second {
		set[value]++
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func Support(x []string, y [][]string) float64 {
	count := 0
	for i := 0; i < len(y); i++ {
		if Subset(x, y[i]) {
			count++
		}
	}
	sup := float64(count) / float64(len(y))
	return sup
}

func Result(dataset [][]string) [][]string {
	l := len(dataset)
	var combinations Transactions
	for i := 0; i < l; i++ {
		transaction := dataset[i]
		combinations = append(combinations, &Transaction{transaction, Support(transaction, dataset)})
	}
	sort.Sort(combinations)

	return combinations.Patterns()
}
