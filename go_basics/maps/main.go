package main

import "fmt"

type entry struct {
	key   string
	value int
}

type simpleMap []entry

func main() {
	firstEntry := entry{
		key:   "1",
		value: 0,
	}
	secondEntry := entry{
		key:   "2",
		value: 1,
	}

	m := simpleMap{firstEntry, secondEntry}
	fmt.Println(lookup(m, "2"))
}

func lookup(m simpleMap, key string) int {
	for _, e := range m {
		if e.key == key {
			return e.value
		}
	}
	return 0
}
