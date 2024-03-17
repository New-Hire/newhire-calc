package hire

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"newhire-rate/model"
	"os"
)

func init() {
	// TODO: 环境切换格式
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

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
	sdsd := runDgraphTxn(dgraphClient)
	//runDgraphTxn2(dgraphClient)
	calcBasicScore(sdsd, 1)
}

/**
 * Deep越深,权重加权越大,比如1-10为0.3，1-3为0.1
 */
func calcScore(aaa2 []Aaaa2) {
	deepW := calcWeightDeep(aaa2)
	fmt.Printf("%d", deepW)
}

type AaaaData struct {
	Data []Aaaa `json:"data"`
}

type Aaaa struct {
	UserId    int64
	CompanyId int64
	Raters    []Aaaa
	Score1    int `json:"score1,omitempty"`
	Score2    int `json:"score2,omitempty"`
}

type Aaaa2 struct {
	RaterId        int64
	UserId         int64
	Deep           int8
	RaterCompanyId int64
	UserCompanyId  int64
	Score1         int
	Score2         int
}

func runDgraphTxn(c *dgo.Dgraph) []Aaaa2 {
	txn := c.NewTxn()
	ctx := context.Background()

	// 查询我被多少人评分过
	q := `query all($userId: string) {
		data(func: eq(User.id, $userId)) @recurse(loop: false) {
			UserId: User.id
			CompanyId: User.company
			Raters: User.raters @facets(Score1:score1,Score2:score2)
		}
	}`
	res, err := txn.QueryWithVars(ctx, q, map[string]string{"$userId": "2"})
	if err != nil {
		log.Fatal(err)
	}

	var p AaaaData
	err2 := json.Unmarshal(res.Json, &p)
	if err2 != nil {
		log.Fatal(err2)
	}
	aaa := &[]Aaaa2{}
	model := p.Data[0]
	userId := model.UserId
	//companyId := model.CompanyId
	resolveRaters(userId, model.Raters, aaa, 0)
	return *aaa
}

//func runDgraphTxn2(c *dgo.Dgraph) {
//	txn := c.NewTxn()
//	ctx := context.Background()
//
//	// 查询我评分过多少人
//	q := `query all($userId: string) {
//		data(func: eq(User.id, $userId)) @recurse(loop: false) {
//			UserId: User.id
//			CompanyId: User.company
//			Raters: ~User.raters @facets(Score1:score1,Score2:score2)
//		}
//	}`
//	res, err := txn.QueryWithVars(ctx, q, map[string]string{"$userId": "3"})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var p AaaaData
//	err2 := json.Unmarshal(res.Json, &p)
//	if err2 != nil {
//		log.Fatal(err2)
//	}
//	var aaa []Aaaa2
//	model := p.Data[0]
//	userId := model.UserId
//	//companyId := model.CompanyId
//	resolveRaters2(userId, model.Raters, aaa, 0)
//}

func resolveRaters(userId int64, raters []Aaaa, aaa *[]Aaaa2, deep int8) {
	deep++
	for i := 0; i < len(raters); i++ {
		rater := raters[i]
		if rater.Raters != nil {
			resolveRaters(rater.UserId, rater.Raters, aaa, deep)
		}
		aa := Aaaa2{
			RaterId:        rater.UserId,
			UserId:         userId,
			Deep:           deep,
			RaterCompanyId: rater.CompanyId,
			UserCompanyId:  userId,
			Score1:         rater.Score1,
			Score2:         rater.Score2,
		}
		fmt.Printf("%+v\n", aa)
		*aaa = append(*aaa, aa)
	}
}

//func resolveRaters2(userId int64, raters []Aaaa, aaa []Aaaa2, deep int8) {
//	deep++
//	for i := 0; i < len(raters); i++ {
//		rater := raters[i]
//		if rater.Raters != nil {
//			resolveRaters(rater.UserId, rater.Raters, aaa, deep)
//		}
//		aa := Aaaa2{
//			RaterId:        userId,
//			UserId:         rater.UserId,
//			Deep:           deep,
//			RaterCompanyId: userId,
//			UserCompanyId:  rater.CompanyId,
//			Score1:         rater.Score1,
//			Score2:         rater.Score2,
//		}
//		aaaa2s := append(aaa, aa)
//		fmt.Printf("%+v\n", aaaa2s)
//	}
//}

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
