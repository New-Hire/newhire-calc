package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"newhire-rate/calc"
	"newhire-rate/model"
	"os"
	"strconv"
)

func init() {
	// TODO: 环境切换格式
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// 初始化数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/hire?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := connectToDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 初始化图数据库
	dgraphClient := createDgraphClient()

	consumer(func(body string) {
		userId, parseErr := strconv.ParseInt(body, 10, 64)
		if parseErr != nil {
			log.WithFields(log.Fields{
				"userId": userId,
			}).Error("数据异常未处理，已ACK")
			// 对于异常数据，记录日志并直接返回
			return
		}

		// 取图数据库
		nodes := getNodes(dgraphClient, userId)
		log.Debug("已取回图数据")

		// 计算
		m, calcErr := calc.Calc(nodes)
		if calcErr != nil {
			log.WithFields(log.Fields{
				"userId": userId,
			}).Error("分数结果计算异常，中断后续执行")
			// 对于异常数据，记录日志并直接返回
			return
		} else {
			log.Debug("分数结果已计算")
		}

		// 将结果写回数据库
		for anyUserId, score := range m {
			var userScore model.UserScore
			userScore.UserId = anyUserId
			userScore.Score = int(score * 10000)
			log.Debugln("userScore=", userScore)
			db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"score", "updated_at"}),
			}).Create(&userScore)
		}
		log.Debug("分数结果已存储")
	})
}

type NodeData struct {
	Data []model.User `json:"data"`
}

// TODO: 泛型
func runDgraphTxn(c *dgo.Dgraph, query string, queryParams map[string]string) model.User {
	txn := c.NewTxn()
	ctx := context.Background()

	// 查询我被多少人评分过
	res, err := txn.QueryWithVars(ctx, query, queryParams)
	if err != nil {
		log.Fatal(err)
	}

	var p NodeData
	err = json.Unmarshal(res.Json, &p)
	if err != nil {
		log.Fatal(err)
	}
	return p.Data[0]
}

func getNodes(c *dgo.Dgraph, userId int64) []model.Node {
	// 查询我被多少人评分过
	q := `query all($userId: string) {
		data(func: eq(User.id, $userId)) @recurse(loop: false) {
			UserId: User.id
			CompanyId: User.company
			Raters: User.raters @facets(Score1:score1,Score2:score2)
		}
	}`
	res := runDgraphTxn(c, q, map[string]string{"$userId": strconv.FormatInt(userId, 10)})
	var nodes []model.Node
	resolveRaters(res.UserId, res.Raters, &nodes, 0)
	return nodes
}

// 将图数据库结构转化成 model.Node 格式
func resolveRaters(userId int64, raters []model.User, nodes *[]model.Node, deep int8) {
	deep++
	for i := 0; i < len(raters); i++ {
		rater := raters[i]
		if rater.Raters != nil {
			resolveRaters(rater.UserId, rater.Raters, nodes, deep)
		}
		node := model.Node{
			RaterId:        rater.UserId,
			UserId:         userId,
			Deep:           deep,
			RaterCompanyId: rater.CompanyId,
			UserCompanyId:  userId,
			Score1:         rater.Score1,
			Score2:         rater.Score2,
		}
		fmt.Printf("%+v\n", node)
		*nodes = append(*nodes, node)
	}
}

func connectToDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func createDgraphClient() *dgo.Dgraph {
	// TODO: 密钥配置
	conn, err := dgo.DialCloud("green-feather-41381518.grpc.ap-south-1.aws.cloud.dgraph.io:443", "OGUzMTljOTAwY2EwNmMyZDAxYmUxYmYzZDBhMDE5Yjc=")
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	return dgo.NewDgraphClient(api.NewDgraphClient(conn))
}
