package calc

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"math"
	"newhire-rate/model"
	"os"
	"slices"
)

func init() {
	// TODO: 环境切换格式
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
}

const (
	MinDeep int8 = 1
)

func Calc(nodes []model.Node) (map[int64]float64, error) {
	if nodes == nil || len(nodes) == 0 {
		return nil, errors.New("空节点，无法计算")
	}
	maxDeep := slices.MaxFunc(nodes, func(a, b model.Node) int {
		if a.Deep > b.Deep {
			return 1
		} else if a.Deep < b.Deep {
			return -1
		}
		return 0
	}).Deep

	m := make(map[int64]float64)
	// 评分离散，越接近0代表所有评分集中，反之代表争议很大
	m2 := make(map[int64]float64)
	// 评分者个人对集体的离散，越接近0代表此人评分越从众，反之代表与他人有很大不同
	m3 := make(map[int64]float64)
	upCalc(nodes, maxDeep, m2, m3)
	downCalc(nodes, maxDeep, m, m2, m3)
	return m, nil
}

func groupByUserIdWhereDeep(nodes []model.Node, deep int8) map[int64][]model.Node {
	groups := make(map[int64][]model.Node)
	for _, node := range nodes {
		if deep == node.Deep {
			groups[node.UserId] = append(groups[node.UserId], node)
		}
	}
	return groups
}

// 计算节点权重
func calcNodeWeight(node model.Node, nodeScoreMap *map[int64]float64) float64 {
	scoreMap := *nodeScoreMap
	raterScore := scoreMap[node.RaterId]
	if raterScore > 0 {
		return raterScore
	} else {
		// 顶节点的默认分
		return 5
	}
}

// 计算来自上级节点传递的能量分
func calcNodeIncrScore(node model.Node, nodeScoreMap *map[int64]float64, nodeWeight float64) float64 {
	scoreMap := *nodeScoreMap
	raterScore := scoreMap[node.RaterId]
	if raterScore == 0 {
		// 顶节点给予默认分5
		raterScore = 5
	}
	log.Debug("raterScore=", raterScore)
	incrScore := raterScore * nodeWeight / 10
	log.Debug("incrScore=", incrScore)
	return incrScore
}

// 往下计算
func downCalc(nodes []model.Node, maxDeep int8, nodeScoreMap map[int64]float64, nodeDMap map[int64]float64, dMap map[int64]float64) {
	for currDeep := maxDeep; currDeep > 0; currDeep-- {
		log.Debug("-----------DOWN-DOWN-DOWN-DOWN-------------")
		log.Debug("currDeep=", currDeep)
		nodesGroup := groupByUserIdWhereDeep(nodes, currDeep)
		log.Debug("len(nodesGroup)=", len(nodesGroup))
		for userId, nodes := range nodesGroup {
			log.Debug("userId=", userId)
			log.Debug("nodes.count=", len(nodes))

			/**
			 * 形如 a / (a + b) 中的 (a+b)，用以取权重总和，方便计算动态权重
			 */
			var weightSum float64 = 0
			for _, node := range nodes {
				log.Debugf("node=%+v", node)
				weightSum += calcNodeWeight(node, &nodeScoreMap)
			}
			log.Debug("总权重=", weightSum)

			totalScore := 0.0

			// 确保 for nodes 是顺序的
			for _, node := range nodes {
				nodeWeight := calcNodeWeight(node, &nodeScoreMap) / weightSum
				log.Debug("node.RaterId=", node.RaterId)
				log.Debug("nodeWeight=", nodeWeight)
				nodeScore := float64(node.Score1) * nodeWeight
				totalScore += nodeScore
				log.Debug("nodeScore=", nodeScore)

				incrScore := calcNodeIncrScore(node, &nodeScoreMap, nodeWeight)
				log.Debug("incrScore=", incrScore)
				totalScore += incrScore

				// 识人之能
				if nodeDMap[node.UserId] > 0 {
					coeff2 := 1.0 / nodeDMap[node.UserId]
					log.Debug("coeff2=", coeff2)
					totalScore += 0.8 * coeff2
				}
			}

			// dMap越接近0，代表争议越小
			if dMap[userId] > 0 {
				coeff1 := 1.0 / dMap[userId]
				log.Debug("coeff1=", coeff1)
				// 争议附加分
				totalScore += 0.5 * coeff1
			}

			nodeScoreMap[userId] = totalScore
			log.Debug("totalScore=", totalScore)
			log.Debugf("map=%+v", nodeScoreMap)
		}
	}
}

// 往上计算
func upCalc(nodes []model.Node, maxDeep int8, nodeDMap map[int64]float64, dMap map[int64]float64) {
	for currDeep := MinDeep; currDeep <= maxDeep; currDeep++ {
		log.Debug("-----------UP-UP-UP-UP-------------")
		log.Debug("currDeep=", currDeep)
		nodesGroup := groupByUserIdWhereDeep(nodes, currDeep)
		log.Debug("len(nodesGroup)=", len(nodesGroup))
		for userId, nodes := range nodesGroup {
			log.Debug("userId=", userId)
			log.Debug("nodes.count=", len(nodes))

			if len(nodes) <= 1 {
				log.Debug("节点无多人评分，Skip...")
				continue
			}

			var scores []int
			for _, node := range nodes {
				scores = append(scores, node.Score1)
			}
			log.Debugln("scores=", scores)

			arrD := calcArrDiscrepancy(scores)
			dMap[userId] = arrD
			log.Debugln("评分离散=", arrD)

			// 根据被评者的评分离散，为评分者附加额外离散，表达 评分者 被 被评人 影响
			var appendD float64 = 0
			if nodeDMap[userId] > 0 {
				appendD = math.Sqrt(nodeDMap[userId] / 2)
			}
			log.Debugln("追加离散", appendD)

			// 确保 for nodes 是顺序的
			for _, node := range nodes {
				// 计算此人评分从众程度(离散程度)
				d, err := calcDiscrepancy2(float64(node.Score1), scores, arrD)
				if err != nil {
					// 即使错误也不能影响后续计算
					log.Fatal(err)
				}
				log.Debug("node.RaterId=", node.RaterId)
				log.Debug("原始d=", d)
				log.Debug("最终d=", d+appendD)
				nodeDMap[node.RaterId] = d + appendD
				log.Debugln("评分者个人对集体的离散", nodeDMap)
			}
		}
	}
}
