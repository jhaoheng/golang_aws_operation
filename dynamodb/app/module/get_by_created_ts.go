package module

import (
	"dynamodb/app/ddbservice"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

/*
一個 item 約 126 Byte
一次拿 8000 個 items, 約 0.96 MB (=126 * 8000 / 1024 / 1024)
每
*/

// 計算多少錢取得多少資料
func GetItems() (allItems []ddbservice.ITEM, totalRCU float64) {
	/*
		- 每 10 元美金, 約可刪除 13477 個物件, 取 13000 個物件
		- 取得 13000 個物件, 每個物件 128 Byte, 總共 1625 KB
		- 每 4 KB 消耗 1 RCU, 總共約 406.5 RCU
		- 每次 scan 只能取得最多 1 MB 的資料
		- 1625 KB, 約 13000 個 item (= 1625*1024/128)
	*/
	var totalItemsCount int64 = 13000
	var ts int64 = time.Now().Unix()
	//
	var lastEvaluatedKey map[string]*dynamodb.AttributeValue = nil
	for {
		var items []ddbservice.ITEM
		var rcu float64 = 0
		items, lastEvaluatedKey, rcu = GetByCreatedTsWithUpperBounder(ts, totalItemsCount, lastEvaluatedKey)
		allItems = append(allItems, items...)
		totalItemsCount = totalItemsCount - int64(len(allItems))
		totalRCU = totalRCU + rcu
		if totalItemsCount <= 0 {
			break
		} else if lastEvaluatedKey == nil {
			break
		}
	}
	// fmt.Println(len(allItems))
	// for _, item := range allItems {
	// 	b, _ := json.Marshal(item)
	// 	fmt.Println(string(b))
	// }
	return
}

/*
	- 每次 scan 只能取得最多 1 MB 的資料
	- 需要將 LastEvaluatedKey 帶入 ExclusiveStartKey
*/
func GetByCreatedTsWithUpperBounder(ts int64, totalItemsCount int64, exclusiveStartKey map[string]*dynamodb.AttributeValue) (saveItems []ddbservice.ITEM, lastEvaluatedKey map[string]*dynamodb.AttributeValue, rcu float64) {

	// totalItemsCount = 13000, 耗損 $10 WCU 刪除
	var limitSize int64 = totalItemsCount
	//
	if ddbservice.DDBService == nil {
		ddbservice.NewDDB()
	}

	//
	output, err := ddbservice.DDBService.ScanByCreatedTs(ddbservice.Lower_than, ts, limitSize, exclusiveStartKey)
	if err != nil {
		panic(err)
	}
	//
	rcu = *output.ConsumedCapacity.CapacityUnits
	fmt.Printf("WCU=%v, itemsCount:%v\n", *output.ConsumedCapacity.CapacityUnits, *output.Count)
	//
	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &saveItems)
	if err != nil {
		panic(err)
	}
	//
	if output.LastEvaluatedKey != nil {
		lastEvaluatedKey = output.LastEvaluatedKey
	}

	return
}
