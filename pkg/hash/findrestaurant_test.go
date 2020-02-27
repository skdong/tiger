package hash

import "testing"

type RestaurantCase struct {
	list1  []string
	list2  []string
	desire []string
}

func TestFindRestaurant(t *testing.T) {
	cases := []RestaurantCase{
		RestaurantCase{list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2:  []string{"KFC", "Shogun", "Burger King"},
			desire: []string{"KFC", "Shogun", "Burger King"}},
	}
	for _, testCase := range cases {
		real := findRestaurant(testCase.list1, testCase.list2)
		t.Fatalf("%v", real)
	}
}
