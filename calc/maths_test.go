package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcDiscrepancy0(t *testing.T) {
	result1, _ := calcDiscrepancy(1, []int{1, 3, 7, 8, 9, 10})
	result2, _ := calcDiscrepancy(5, []int{1, 3, 7, 8, 9, 10})
	result3, _ := calcDiscrepancy(10, []int{1, 3, 7, 8, 9, 10})
	result4, _ := calcDiscrepancy(8, []int{1, 3, 7, 8, 9, 10})
	result5, _ := calcDiscrepancy(6, []int{1, 3, 7, 8, 9, 10})

	assert.Equal(t, 1.4985372985307102, result1)
	assert.Equal(t, 0.3746343246326775, result2)
	assert.Equal(t, 1.0302443927398635, result3)
	assert.Equal(t, 0.4682929057908471, result4)
	assert.Equal(t, 0.09365858115816932, result5)
}

func TestCalcDiscrepancy1(t *testing.T) {
	expectedResult := 0.07856742013183861
	result, _ := calcDiscrepancy(5, []int{1, 10})
	assert.Equal(t, expectedResult, result)
}

func TestCalcDiscrepancy2(t *testing.T) {
	expectedResult := 0.0
	result, _ := calcDiscrepancy(10, []int{10, 10})
	assert.Equal(t, expectedResult, result)

	expectedResult2 := 0.0
	result2, _ := calcDiscrepancy(8, []int{10, 10})
	assert.Equal(t, expectedResult2, result2)
}

func TestCalcDiscrepancy3(t *testing.T) {
	expectedResult := 1.649915822768611
	result, _ := calcDiscrepancy(5, []int{7, 10})
	assert.Equal(t, expectedResult, result)
}

func TestCalcDiscrepancy4(t *testing.T) {
	expectedResult := 0.7071067811865476
	result, _ := calcDiscrepancy(7, []int{7, 10})
	assert.Equal(t, expectedResult, result)
}

func TestCalcDiscrepancy5(t *testing.T) {
	result, _ := calcDiscrepancy(9, []int{1, 10})
	result2, _ := calcDiscrepancy(8, []int{1, 10})
	assert.Greater(t, result, result2)
}

func TestCalcDiscrepancy6(t *testing.T) {
	result, _ := calcDiscrepancy(9, []int{1, 10})
	result2, _ := calcDiscrepancy(10, []int{1, 10})
	result3, _ := calcDiscrepancy(11, []int{1, 10})
	assert.Less(t, result, result2)
	assert.Greater(t, result3, result2)
	assert.Less(t, result, result3)
}

func TestCalcArrDiscrepancy1(t *testing.T) {
	expectedResult := 6.363961030678928
	result := calcArrDiscrepancy([]int{1, 10})
	assert.Equal(t, expectedResult, result)
}

func TestCalcArrDiscrepancy2(t *testing.T) {
	expectedResult := 0.0
	result := calcArrDiscrepancy([]int{5, 5, 5, 5})
	assert.Equal(t, expectedResult, result)
}
