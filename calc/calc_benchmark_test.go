package calc

import (
	"math/rand"
	"newhire-rate/model"
	"testing"
)

const (
	MaxNodeCount = 1000000
	MaxUserCount = 1000000
	MaxDeep      = 60
)

// 测试大数据集
func BenchmarkHigh(b *testing.B) {
	var nodes []model.Node
	for i := 1; i < MaxNodeCount; i++ {
		d := model.Node{
			RaterId:        rand.Int63n(100) + 1,
			UserId:         rand.Int63n(100) + 1,
			Deep:           int8(rand.Intn(16)) + 1,
			RaterCompanyId: 1,
			UserCompanyId:  1,
			Score1:         rand.Intn(10) + 1,
			Score2:         4,
		}
		nodes = append(nodes, d)
	}
	_, _ = Calc(nodes)
}

// 测试大数据集+大用户数
func BenchmarkHigh2(b *testing.B) {
	var nodes []model.Node
	for i := 1; i < MaxNodeCount; i++ {
		d := model.Node{
			RaterId:        rand.Int63n(MaxUserCount) + 1,
			UserId:         rand.Int63n(MaxUserCount) + 1,
			Deep:           int8(rand.Intn(16)) + 1,
			RaterCompanyId: 1,
			UserCompanyId:  1,
			Score1:         rand.Intn(10) + 1,
			Score2:         4,
		}
		nodes = append(nodes, d)
	}
	_, _ = Calc(nodes)
}

// 测试大数据集+大复杂深度
func BenchmarkHigh3(b *testing.B) {
	var nodes []model.Node
	for i := 1; i < MaxNodeCount; i++ {
		d := model.Node{
			RaterId:        rand.Int63n(100) + 1,
			UserId:         rand.Int63n(100) + 1,
			Deep:           int8(rand.Intn(MaxDeep)) + 1,
			RaterCompanyId: 1,
			UserCompanyId:  1,
			Score1:         rand.Intn(10) + 1,
			Score2:         4,
		}
		nodes = append(nodes, d)
	}

	_, _ = Calc(nodes)
}

// 测试大数据集+大用户数+大复杂深度
func BenchmarkHigh4(b *testing.B) {
	var nodes []model.Node
	for i := 1; i < MaxNodeCount; i++ {
		d := model.Node{
			RaterId:        rand.Int63n(MaxUserCount) + 1,
			UserId:         rand.Int63n(MaxUserCount) + 1,
			Deep:           int8(rand.Intn(MaxDeep)) + 1,
			RaterCompanyId: 1,
			UserCompanyId:  1,
			Score1:         rand.Intn(10) + 1,
			Score2:         4,
		}
		nodes = append(nodes, d)
	}

	_, _ = Calc(nodes)
}
