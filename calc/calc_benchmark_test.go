package calc

import (
	"math/rand"
	"newhire-rate/model"
	"testing"
)

// 测试大数据集 3.58s
func BenchmarkHigh(b *testing.B) {
	var aaa2 []model.Aaaa2
	for i := 1; i < 10000000; i++ {
		d := model.Aaaa2{
			RaterId:        rand.Int63n(100) + 1,
			UserId:         rand.Int63n(100) + 1,
			Deep:           int8(rand.Intn(16)) + 1,
			RaterCompanyId: 1,
			UserCompanyId:  1,
			Score1:         rand.Intn(10) + 1,
			Score2:         4,
		}
		aaa2 = append(aaa2, d)
	}
	CalcBasicScore(aaa2, 1)
}

// 测试大数据集+大用户数 9.49s
func BenchmarkHigh2(b *testing.B) {
	var aaa2 []model.Aaaa2
	for i := 1; i < 10000000; i++ {
		d := model.Aaaa2{
			RaterId:        rand.Int63n(1000000) + 1,
			UserId:         rand.Int63n(1000000) + 1,
			Deep:           int8(rand.Intn(16)) + 1,
			RaterCompanyId: 1,
			UserCompanyId:  1,
			Score1:         rand.Intn(10) + 1,
			Score2:         4,
		}
		aaa2 = append(aaa2, d)
	}
	CalcBasicScore(aaa2, 1)
}

// 测试大数据集+大复杂深度 8.27s
func BenchmarkHigh3(b *testing.B) {
	var aaa2 []model.Aaaa2
	for i := 1; i < 10000000; i++ {
		d := model.Aaaa2{
			RaterId:        rand.Int63n(100) + 1,
			UserId:         rand.Int63n(100) + 1,
			Deep:           int8(rand.Intn(160)) + 1,
			RaterCompanyId: 1,
			UserCompanyId:  1,
			Score1:         rand.Intn(10) + 1,
			Score2:         4,
		}
		aaa2 = append(aaa2, d)
	}

	CalcBasicScore(aaa2, 1)
}

// 测试大数据集+大用户数+大复杂深度 13.16s
func BenchmarkHigh4(b *testing.B) {
	var aaa2 []model.Aaaa2
	for i := 1; i < 10000000; i++ {
		d := model.Aaaa2{
			RaterId:        rand.Int63n(1000000) + 1,
			UserId:         rand.Int63n(1000000) + 1,
			Deep:           int8(rand.Intn(160)) + 1,
			RaterCompanyId: 1,
			UserCompanyId:  1,
			Score1:         rand.Intn(10) + 1,
			Score2:         4,
		}
		aaa2 = append(aaa2, d)
	}

	CalcBasicScore(aaa2, 1)
}
