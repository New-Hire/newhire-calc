package hire

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcBasicScore(t *testing.T) {
	aaa2 := []Aaaa2{
		{RaterId: 2, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 6, Score2: 4},
		{RaterId: 5, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 10, Score2: 4},
		{RaterId: 3, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 4, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
	}

	expectedResult := 71884

	result := calcBasicScore(aaa2, 1)

	assert.Equal(t, expectedResult, result)
}

func TestCalcBasicScore2(t *testing.T) {
	aaa2 := []Aaaa2{
		{RaterId: 2, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 6, Score2: 4},
		{RaterId: 5, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 10, Score2: 4},
		{RaterId: 3, UserId: 2, Deep: 2, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		{RaterId: 4, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 8, Score2: 4},
		{RaterId: 5, UserId: 3, Deep: 3, RaterCompanyId: 1, UserCompanyId: 1, Score1: 4, Score2: 4},
		// 追加
		{RaterId: 6, UserId: 1, Deep: 1, RaterCompanyId: 1, UserCompanyId: 1, Score1: 2, Score2: 4},
	}

	expectedResult := 56400

	result := calcBasicScore(aaa2, 1)

	assert.Equal(t, expectedResult, result)
}
