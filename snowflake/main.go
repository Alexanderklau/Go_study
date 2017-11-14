package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zheng-ji/goSnowFlake"
	"strconv"
)

var idWorkerMap = make(map[int]*goSnowFlake.IdWorker)

func main() {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Get ID
	r.GET("/worker/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Params.ByName("id"))
		value, ok := idWorkerMap[id]
		if ok {
			nid, _ := value.NextId()
			c.JSON(200, gin.H{"id": nid})
		} else {
			iw, err := goSnowFlake.NewIdWorker(int64(id))
			if err == nil {
				nid, _ := iw.NextId()
				idWorkerMap[id] = iw
				c.JSON(200, gin.H{"id": nid})
			} else {
				fmt.Println(err)
			}
		}
	})

	// Listen and Server in 0.0.0.0:8182
	r.Run(":8182")
}
