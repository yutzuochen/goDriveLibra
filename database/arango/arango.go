package arango

import (
	"context"
	"example/dao"
	"example/database"
	"example/env"
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"gitlab.geax.io/demeter/gologger/logger"
)

type handler struct {
	db driver.Database
}

//func NewArango() (*handler, error) {
func NewArango(cfg *env.DatabaseConfig) (database.Manager, error) {
	this := &handler{}
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{cfg.AddressIP},
	})
	if err != nil {
		logger.Errorf("NewArango/NewConnection Error: %v", err)
		return nil, err
	}

	cli, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		logger.Errorf("NewArango/NewClient Error: %v", err)
		return nil, err
	}

	db, err := cli.Database(context.TODO(), "Database")
	if err != nil {
		logger.Errorf("NewArango/Database dbname=%v Error: %v", "Database", err)
		return nil, err
	}

	this.db = db
	//Instant = this

	logger.Infof("Arango connected to dbname=%v success", "Database")

	return this, nil
}

func (hdr *handler) GetPlayer(username string) []*dao.Player {
	ctx := context.Background()
	bindVars := map[string]interface{}{
		"name": username,
	}
	query := "FOR p IN Players FILTER p.nickname == @name  RETURN p"
	cursor, err := hdr.db.Query(ctx, query, bindVars)
	if err != nil {
		fmt.Println("query wrong: ", err)
	}
	fmt.Println("cursor: ", cursor)
	defer cursor.Close()
	list := []*dao.Player{}
	for {
		result := &dao.Player{} // 注意： 這裡一定要先初始化，否則該變數後面是不能用的，亦即該行不能是 var result *dao.Player
		_, err := cursor.ReadDocument(context.TODO(), result)
		if driver.IsNoMoreDocuments(err) {
			break
		}
		fmt.Println("here 3")
		fmt.Println("result: ", result.Nickname)
		list = append(list, result)

	}
	fmt.Println("list: ", list)

	return list

	/*
		var docList []Player
		for {
			var doc Player
			_, _ = cursor.ReadDocument(ctx, &doc)
			if driver.IsNoMoreDocuments(err) {
				fmt.Println("break!!")
				break
			} else if err != nil {
				// handle other errors
			}
			docList = append(docList, doc)
			//fmt.Println(docList)

		}
		fmt.Printf("docList", docList)
		return nil
	*/

}

/*
// reflect
func mapToPlayer(m map[string]interface{}) *dao.Player {
	var res *dao.Player
	for k, v := range m {
		switch t := v.(type) {
		case string:
			//if k == "nickname" {
			//	res.nickname = v.(string)
			//}
			res.k = t
		case int32:
			fmt.Println(k, t)
		case int64:
			fmt.Println(k, t)
		}
	}
	return res
}
*/

func (this *handler) GetPlayerBalance(playerId string) (int, error, int64) {
	fmt.Println("ok1")
	action := `function (Params) {
		const db = require('@arangodb').db;
		const playerCol = db._collection("Players");
		const player = playerCol.document(Params[0]);
		return player["outcome"];}`
	fmt.Println("ok2")
	// TODO: 確定 MaxTransactionSize 如何正確計算
	options := &driver.TransactionOptions{
		MaxTransactionSize: 1000,
		ReadCollections:    []string{"Players"},
		Params:             []interface{}{playerId},
		WaitForSync:        false,
	}
	fmt.Println("ok3")
	//result, err := this.db.Transaction(context.TODO(), action, options)
	result, err := this.db.Transaction(context.TODO(), action, options)
	fmt.Println("ok4")
	if err != nil {
		logger.Errorf("arango/GetPlayerBalance Error: %v", err)
		return 404, err, 0
	}

	return 200, nil, int64(result.(float64))
}
