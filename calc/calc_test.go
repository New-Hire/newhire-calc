package calc

import (
	"github.com/stretchr/testify/assert"
	"newhire-rate/model"
	"testing"
)

func TestCalcBasicScore(t *testing.T) {
	aaa2 := []model.Aaaa2{
		{RaterId: 2, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 6, Score2: 4},
		{RaterId: 5, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 10, Score2: 4},
		{RaterId: 3, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 4, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
	}

	expectedResult := map[int64]int{
		1: 71884,
		2: 67593,
		3: 60000,
	}

	result := CalcBasicScore(aaa2, 1)

	assert.Equal(t, expectedResult, result)
}

func TestCalcBasicScore2(t *testing.T) {
	aaa2 := []model.Aaaa2{
		{RaterId: 2, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 6, Score2: 4},
		{RaterId: 5, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 10, Score2: 4},
		{RaterId: 3, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 4, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		// 追加
		{RaterId: 6, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 2, Score2: 4},
	}

	expectedResult := map[int64]int{
		1: 56400,
		2: 67593,
		3: 60000,
	}

	result := CalcBasicScore(aaa2, 1)

	assert.Equal(t, expectedResult, result)
}

func TestCalcBasicScore3(t *testing.T) {
	aaa2 := []model.Aaaa2{
		{RaterId: 2, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 6, Score2: 4},
		{RaterId: 5, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 10, Score2: 4},
		{RaterId: 3, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 4, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 6, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 2, Score2: 4},
		// 追加
		{RaterId: 3, UserId: 5, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
	}
	expectedResult := map[int64]int{
		1: 58283,
		2: 76158,
		3: 55805,
		5: 80558,
	}

	result := CalcBasicScore(aaa2, 1)

	assert.Equal(t, expectedResult, result)
}

// 追加1评7，看看其它节点有没有变化
func TestCalcBasicScore4(t *testing.T) {
	aaa2 := []model.Aaaa2{
		{RaterId: 2, UserId: 1, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 1, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 6, Score2: 4},
		{RaterId: 5, UserId: 2, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 10, Score2: 4},
		{RaterId: 3, UserId: 2, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 4, UserId: 3, Deep: 4, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 3, Deep: 4, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 6, UserId: 1, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 2, Score2: 4},
		{RaterId: 3, UserId: 5, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		// 追加
		{RaterId: 1, UserId: 7, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 9, Score2: 4},
	}
	expectedResult := map[int64]int{
		1: 58283,
		2: 76158,
		3: 55805,
		5: 80558,
		7: 90582,
	}

	result := CalcBasicScore(aaa2, 1)

	assert.Equal(t, expectedResult, result)
}

// 追加8评7(8是空节点)，看看其它节点有没有变化
func TestCalcBasicScore5(t *testing.T) {
	aaa2 := []model.Aaaa2{
		{RaterId: 2, UserId: 1, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 1, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 6, Score2: 4},
		{RaterId: 5, UserId: 2, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 10, Score2: 4},
		{RaterId: 3, UserId: 2, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 4, UserId: 3, Deep: 4, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 3, Deep: 4, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 6, UserId: 1, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 2, Score2: 4},
		{RaterId: 3, UserId: 5, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 1, UserId: 7, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 9, Score2: 4},
		// 追加
		{RaterId: 8, UserId: 7, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 2, Score2: 4},
	}
	expectedResult := map[int64]int{
		1: 58283,
		2: 76158,
		3: 55805,
		5: 80558,
		7: 57985,
	}

	result := CalcBasicScore(aaa2, 1)

	assert.Equal(t, expectedResult, result)
}

// 追加2评7(2是上层节点)，看看其它节点有没有变化
func TestCalcBasicScore6(t *testing.T) {
	aaa2 := []model.Aaaa2{
		{RaterId: 2, UserId: 1, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 1, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 6, Score2: 4},
		{RaterId: 5, UserId: 2, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 10, Score2: 4},
		{RaterId: 3, UserId: 2, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 4, UserId: 3, Deep: 4, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 3, Deep: 4, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 6, UserId: 1, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 2, Score2: 4},
		{RaterId: 3, UserId: 5, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 1, UserId: 7, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 9, Score2: 4},
		{RaterId: 8, UserId: 7, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 2, Score2: 4},
		// 追加
		{RaterId: 2, UserId: 7, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 10, Score2: 4},
	}
	expectedResult := map[int64]int{
		1: 58283,
		2: 76158,
		3: 55805,
		5: 80558,
		7: 75650,
	}

	result := CalcBasicScore(aaa2, 1)

	assert.Equal(t, expectedResult, result)
}
