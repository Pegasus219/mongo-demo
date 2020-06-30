package main

import (
	"encoding/json"
	"fmt"
	"mongo-db/models"
	"time"
)

func main() {
	//demo 1
	result, err := models.GetUserActivities(4019916, 12, []int{476, 486})
	if err != nil {
		fmt.Println(err)
	}
	strByte, _ := json.Marshal(result)
	fmt.Println(string(strByte))
	//demo 2
	startAt := time.Now().Unix() - 190*86400
	endAt := time.Now().Unix()
	result2, err := models.GetActivitiesDuring(4019916, 12, startAt, endAt)
	if err != nil {
		fmt.Println(err)
	}
	strByte2, _ := json.Marshal(result2)
	fmt.Println(string(strByte2))
}
