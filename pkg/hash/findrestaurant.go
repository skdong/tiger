package hash

type Restaurant struct {
	index int
	name  string
}

func findRestaurant(list1 []string, list2 []string) []string {
	cache := map[string]int{}
	var restaurant *Restaurant
	var target []string
	for i, resta := range list1 {
		cache[resta] = i
	}
	for index2, resta := range list2 {
		if index1, ok := cache[resta]; ok {
			index := index1 + index2
			if restaurant != nil {
				if restaurant.index > index {
					restaurant.name = resta
					restaurant.index = index
					target = []string{resta}
				} else if restaurant.index == index {
					target = append(target, resta)
				}
			} else {
				restaurant = &Restaurant{name: resta, index: index}
				target = []string{resta}
			}
		}
	}
	return target
}
