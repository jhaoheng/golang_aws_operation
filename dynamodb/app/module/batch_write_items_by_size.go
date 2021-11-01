package module

import (
	"dynamodb/app/ddbservice"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// 寫入多少 size 的 datas
func WriteBySizeMB(size float64) {
	/*
		1 個 item 約 128 Byte
		要寫入 1MB 的資料, 需要 8192 個 items (=1 * 1024 * 1024 / 128)
		每次只能寫入 25 個
		總共需要寫 328 次 (=8192 / 25)
	*/

	// 執行次數
	execTimesBySize := func() int {
		var WriteInMB float64 = size // 總共寫入 2MB 的資料
		const itemSize = 128
		const BatchSize = 25
		execTimes := WriteInMB * 1024 * 1024 / itemSize / BatchSize
		return int(execTimes)
	}

	for i := 0; i <= execTimesBySize(); i++ {
		WriteDatas()
	}
}

/*
- batch_write_items, 每次只能寫入 25 個 items
- 參考 DynamoDB rejects the entire batch write operation
	- https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.BatchWriteItem
*/
func WriteDatas() error {
	//多久過期
	var expired_ts int64 = 60 * 60 * 24
	// 一次最多只能寫入 25 個 items
	items := make([]ddbservice.ITEM, 25)
	//
	for index := range items {
		did := uuid.New().String()
		created_ts_int64 := time.Now().Unix() - int64(len(items)) + int64(index)
		created_ts := strconv.Itoa(int(created_ts_int64))
		//
		item := ddbservice.ITEM{
			DidModuleID: did,
			CreatedTs:   fmt.Sprintf("%v.%v", created_ts, "001"),
			Did:         did,
			Value:       "0",
			ExpiredAt:   fmt.Sprintf("%v", created_ts_int64+expired_ts),
		}
		fmt.Println(item)
		//
		items[index] = item
	}

	//
	if ddbservice.DDBService == nil {
		ddbservice.NewDDB()
	}
	//
	output, err := ddbservice.DDBService.BatchWriteItem(items)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
