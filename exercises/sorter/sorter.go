package main

import (
	"fmt"
	"sort"
)
type people []string

func (p people) Len() int {
	return len(p)
}
func (p people) Swap(i, j int) {
	x := p[i]
	p[i] = p[j]
	p[j] = x
}
func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {

	// Sort using interfaces
	studygroup := people { "joe", "bill", "steve", "kathy", "kelly", "robin", "Robin"}
	fmt.Println(sort.StringsAreSorted(studygroup))
	sort.Sort(studygroup)
	fmt.Println(studygroup)
	fmt.Println(sort.StringsAreSorted(studygroup))

	// Sort using sort.Strings
	x := []string { "tom", "dick", "harry"}
	sort.Strings(x)
	fmt.Println(x)
}
