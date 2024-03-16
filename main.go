package main

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"newhire-rate/model"
)

func main() {
	fmt.Println("start")
	dsn := "root:123456@tcp(127.0.0.1:3306)/hire?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := connectToDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}

	var userRecommend model.UserRecommend
	res := db.First(&userRecommend)
	res.Row()

	dgraphClient := createDgraphClient()
	runDgraphTxn(dgraphClient)
}

func runDgraphTxn(c *dgo.Dgraph) {
	txn := c.NewTxn()
	ctx := context.Background()

	//query := `query {
	//	user as var(func: eq(User.id, "3"))
	//	rater as var(func: eq(User.id, "2"))
	//}`
	//mu := &api.Mutation{
	//	SetNquads: []byte(`
	//		uid(user) <User.id> "3" .
	//		uid(user) <User.company> "1" .
	//		uid(rater) <User.id> "2" .
	//		uid(rater) <User.company> "1" .
	//		uid(user) <User.raters> uid(rater) (score1=13, score2=14) .
	//	`),
	//}
	//req := &api.Request{
	//	Query:     query,
	//	Mutations: []*api.Mutation{mu},
	//	CommitNow: true,
	//}
	//if _, err := txn.Do(ctx, req); err != nil {
	//	log.Fatal(err)
	//}

	// 查询我被多少人评分过
	q := `query all($userId: string) {
		data(func: eq(User.id, $userId)) @recurse(loop: false) {
			User.id
			User.company
			raters: User.raters @facets
		}
	}`
	// 查询我评分过多少人
	q = `query all($userId: string) {
		data(func: eq(User.id, $userId)) @recurse(loop: false) {
			User.id
			User.company
			ratedUsers: ~User.raters @facets
		}
	}`
	res, err := txn.QueryWithVars(ctx, q, map[string]string{"$userId": "2"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", res.Json)
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
