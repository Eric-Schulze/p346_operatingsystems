package L07

import (
	"testing"
	"utils/array_util"
)

func TestSortStrings(t *testing.T) {
	tables := []struct {
		a []string
		w []string
		n []float64
	}{
		{
			//Standard Test
			a: []string{"banana", "10", "apple", "23", "carrot", "34"},
			w: []string{"banana", "apple", "carrot"},
			n: []float64{10, 23, 34},
		},
		{
			//Decimal Test
			a: []string{"banana", "10.5", "apple", "23", "carrot", "daisy", "34.6", "45.3", "67.9"},
			w: []string{"banana", "apple", "carrot", "daisy"},
			n: []float64{10.5, 23, 34.6, 45.3, 67.9},
		},
		{
			//Only strings test
			a: []string{"banana", "apple", "carrot", "daisy"},
			w: []string{"banana", "apple", "carrot", "daisy"},
			n: []float64{},
		},
		{
			//Only numbers test
			a: []string{"10", "23", "34.4", "87.9"},
			w: []string{},
			n: []float64{10, 23, 34.4, 87.9},
		},
		{
			//Empty test
			a: []string{},
			w: []string{},
			n: []float64{},
		},
	}

	for _, table := range tables {
		words, nums := SortStrings(table.a)
		if !array_util.EqualStrings(words, table.w) {
			t.Errorf("Word Sort Error for %s: Expected %#v, Got %#v", table.a, table.w, words)
		}
		if !array_util.EqualFloats(nums, table.n) {
			t.Errorf("Number Sort Error for %s: Expected %#v, Got %#v", table.a, table.n, nums)
		}
	}
}

func TestSumSlice(t *testing.T) {
	tables := []struct {
		n []float64
		s float64
	}{
		{
			//Standard Test
			n: []float64{10, 23, 34},
			s: 67,
		},
		{
			//Decimal Test
			n: []float64{10.5, 23, 34.6, 45.3, 67.9},
			s: 181.3,
		},
		{
			//Negative numbers test
			n: []float64{10, -23, 34.4, -87.9},
			s: -66.5,
		},
		{
			//Empty test
			n: []float64{},
			s: 0,
		},
	}

	for _, table := range tables {
		sum := SumSlice(table.n)
		if sum != table.s {
			t.Errorf("Sum Error for %s: Expected %s, Got %s", table.n, table.s, sum)
		}
	}
}
