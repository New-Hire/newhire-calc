package hire

import (
	log "github.com/sirupsen/logrus"
	"slices"
)

const (
	Score1Weight = 200
	Score2Weight = 100
	Decimal      = 10000
)

// 时间越近，权重越大
// 核心是自上往下计算，将上层能量传递给下层，层层相叠

/**
 * Deep越深,权重加权越大
 */
func calcWeightDeep(aaa2 []Aaaa2) int {
	maxDeep := slices.MaxFunc(aaa2, func(a, b Aaaa2) int {
		if a.Deep > b.Deep {
			return 1
		} else if a.Deep < b.Deep {
			return -1
		}
		return 0
	}).Deep
	const MinDeep int8 = 1
	// TODO: 单测 同层级
	w := int(maxDeep-MinDeep) * 100
	log.WithFields(log.Fields{
		"deepWeight": w,
	}).Debug("权重计算")
	return w
}

func groupByUserIdWhereDeep(students []Aaaa2, deep int8) map[int64][]Aaaa2 {
	groups := make(map[int64][]Aaaa2)
	for _, student := range students {
		if deep == student.Deep {
			groups[student.UserId] = append(groups[student.UserId], student)
		}
	}
	return groups
}

/**
 * 计算基础得分
 * 基础分只算直接打分
 */
func calcBasicScore(aaa2 []Aaaa2, userI int64) int {
	maxDeep := slices.MaxFunc(aaa2, func(a, b Aaaa2) int {
		if a.Deep > b.Deep {
			return 1
		} else if a.Deep < b.Deep {
			return -1
		}
		return 0
	}).Deep

	m := make(map[int64]int)
	for i := maxDeep; i > 0; i-- {
		log.Debug("deep=", i)
		vv := groupByUserIdWhereDeep(aaa2, i)
		//log.Debug("vv.count=", len(vv))
		for userId, value := range vv {
			//log.Debug("value.count=", len(value))
			//ucount := len(value)
			s := 0

			uuuu := 0
			for _, dd := range value {
				ssdd := m[dd.RaterId]
				if ssdd > 0 {
					uuuu += ssdd / 10 // 向下取整
				} else {
					uuuu += 5 * Decimal / 10 // 向下取整
				}
			}

			for _, dd := range value {
				//log.Debug("dd=", dd)
				ssdd := m[dd.RaterId]
				log.Debug("userId=", dd.RaterId, "->", dd.UserId)

				u1 := 0
				if ssdd > 0 {
					u1 = ssdd * Decimal / (uuuu * 10) // 向下取整
				} else {
					u1 = 5 * Decimal * Decimal / (uuuu * 10) // 向下取整
				}
				log.Debug("u1=", u1)
				log.Debug("uuuu=", uuuu)

				s += dd.Score1 * u1
				if ssdd > 0 {
					s += ssdd * u1 / (Decimal * 100)
					log.Debug("p=", ssdd*u1/(Decimal*100))
				} else {
					log.Debug("p=", ssdd)
				}
				//log.Debug("u1=", u1)
			}
			log.Debug("s=", s)
			m[userId] = s
			log.Debug("---------------------------")
			log.Debug("m=", m)
		}
	}
	return m[userI]
}

///**
// * 计算基础得分
// */
//func calcBasicScore(aaa2 []Aaaa2, deepWeight int) int {
//	for _, a2 := range aaa2 {
//		deepCoeff := int(a2.Deep) * deepWeight
//		s1 := deepCoeff * a2.Score1 * Score1Weight
//		s2 := deepCoeff * a2.Score2 * Score2Weight
//	}
//}
