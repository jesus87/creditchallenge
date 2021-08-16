package assigncredits

import (
	"errors"
	"fmt"
)

var (
	credits = map[int32]int32{
		0: 300,
		1: 500,
		2: 700,
	}
)

type DefaultAssignment struct {
	credits map[int32]int32
}

func New() *DefaultAssignment {
	return &DefaultAssignment{
		credits: map[int32]int32{
			0: 300,
			1: 500,
			2: 700,
		},
	}
}

func (a *DefaultAssignment) Asssign(investment int32) (int32, int32, int32, error) {
	if investment < credits[0] {
		return 0, 0, 0, fmt.Errorf("The investment must be greater than %v", credits[0])
	}

	if remainder := investment % 100; remainder > 0 {
		return 0, 0, 0, errors.New("The investment must be divisible by 100")
	}

	permutation := map[int32]int32{
		0: 0,
		1: 0,
		2: 0,
	}

	permutations := []map[int32]int32{}

	a.swap(investment, permutation, &permutations)

	if len(permutations) == 0 {
		return 0, 0, 0, errors.New("Investment cannot be applied")
	}

	permutation = permutations[0]

	return permutation[0],
		permutation[1],
		permutation[2],
		nil
}

func (a *DefaultAssignment) swap(investment int32, p map[int32]int32, permutations *[]map[int32]int32) {
	permutation := map[int32]int32{
		0: p[0],
		1: p[1],
		2: p[2],
	}

	for cType := int32(0); cType < int32(len(a.credits)); cType++ {
		credit := credits[cType]
		investment -= credit
		permutation[cType] += 1
		if investment > 0 {
			a.swap(investment, permutation, permutations)
			continue
		}

		if investment < 0 {
			investment += credit
			permutation[cType] -= 1
			continue
		}

		if investment == 0 {
			if !existsCombination(permutation, permutations) {
				*permutations = append(*permutations, permutation)
			}
			return
		}
	}
}

func existsCombination(p map[int32]int32, permutations *[]map[int32]int32) bool {
	for _, expectedPermutation := range *permutations {
		if len(p) == len(expectedPermutation) {
			for k, v := range expectedPermutation {
				if p[k] != v {
					continue
				}
				return true
			}
		}
	}
	return false
}
