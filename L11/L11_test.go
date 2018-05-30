package L11

import (
	"fmt"
	"testing"
)

type table struct {
	test    []string
	results []thread
}

func TestMakeThreads(t *testing.T) {
	tables := []table{
		{
			//Standard Test
			test:    []string{"apple", "10", "banana", "25", "carrot", "100"},
			results: []thread{thread{name: "apple", delay: 10}, thread{name: "banana", delay: 25}, thread{name: "carrot", delay: 100}},
		},
		{
			//No Thread Test
			test:    []string{},
			results: []thread{},
		},
		{
			//One Thread Test
			test:    []string{"apple", "10"},
			results: []thread{thread{name: "apple", delay: 10}},
		},
	}

	for _, ta := range tables {
		threads, err := MakeThreads(ta.test)
		if err != nil {
			fmt.Println("Make Threads Error: ")
			fmt.Println(err)
			return
		}

		equals := true
		for i, t := range threads {
			if t != ta.results[i] {
				equals = false
				break
			}

		}

		if !equals {
			t.Errorf("Make Threads Error for %s: Expected %#v, Got %#v", ta.test, ta.results, threads)
		}
	}
}
