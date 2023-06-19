package routeUtil

import (
	"main/dbUtil"
)

var dataCol = dbUtil.GetCollection("data")

//func AddFruits() gin.HandlerFunc{
//	return func(c *gin.Context){
//		var amqp storageUtil.Trade
//
//		err := c.BindJSON(&amqp)
//		if err != nil {
//			c.JSON(400, storageUtil.Response{Code: 400, Message: "Bad Request", Success: false, Data: map[string]interface{}{"data": 09990908888888888888888888888888888888888888888888888}})
//			return
//		}
//
//		res, err := dataCol.InsertOne(c, interface{}(amqp.Items))
//		if err != nil {
//			c.JSON(400, storageUtil.Response{Code: 400, Message: "Bad Request", Success: false, Data: map[string]interface{}{"data": 0999}})
//			return
//		}
//
//		c.JSON(200, storageUtil.Response{Code: 200, Message: "OK", Success: true, Data: map[string]interface{}{"data": res}})
//	}
//}
