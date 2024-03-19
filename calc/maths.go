package calc

import (
	"errors"
	"math"
)

/**
 * 计算数字与数组离散程度，返回值越接近0，说明num在arr里越接近中心
 */
func calcDiscrepancy(num float64, arr []int) (float64, error) {
	if len(arr) == 0 {
		return 0, errors.New("数组要求必须有值")
	}

	// 计算平均值
	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))

	discrepancy := calcArrDiscrepancy(arr)

	// 计算num与平均值之差的标准差
	diffFromMean := math.Abs(num - mean)
	if discrepancy == 0 {
		return 0, nil
	}
	return diffFromMean / discrepancy, nil
}

// 等同于 calcDiscrepancy，节省discrepancy计算成本
func calcDiscrepancy2(num float64, arr []int, discrepancy float64) (float64, error) {
	if len(arr) == 0 {
		return 0, errors.New("数组要求必须有值")
	}

	// 计算平均值
	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))

	// 计算num与平均值之差的标准差
	diffFromMean := math.Abs(num - mean)
	if discrepancy == 0 {
		return 0, nil
	}
	return diffFromMean / discrepancy, nil
}

/**
 * 计算数组离散程度
 */
func calcArrDiscrepancy(arr []int) float64 {
	// 计算平均值
	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))

	// 计算所有元素与平均值之差的平方和
	var varianceSum float64
	for _, v := range arr {
		varianceSum += math.Pow(float64(v)-mean, 2)
	}

	// 计算方差并进一步得到标准差（离散程度）
	variance := varianceSum / float64(len(arr)-1)
	return math.Sqrt(variance)
}
