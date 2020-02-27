package vector

import (
	"testing"
)

type UseCase struct {
	Nums   []int
	Desire int
}

func TestPivotIndex(t *testing.T) {
	var useCases []UseCase = []UseCase{
		UseCase{[]int{1, 2, 3}, -1},
		UseCase{[]int{1, 7, 3, 6, 5, 6}, 3},
		UseCase{[]int{-1, -1, 0, 0, -1, -1}, 2},
		UseCase{[]int{-1, -1, 1, 0, 0, -1}, 3},
		UseCase{[]int{-1, -1, 0, 1, -1, -1}, 1},
	}
	for index, useCase := range useCases {
		real := PivotIndex(useCase.Nums)
		if useCase.Desire != real {
			t.Fatalf("The %d is Fail \nExcept [%d] \nreal is [%d]", index, useCase.Desire, real)
		}

	}
}
