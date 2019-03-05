package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/ipipdotnet/ipdb-go"
	"github.com/sungithubid/gin-rpc/pkg/setting"
)

// CityInfo Ip
type CityInfo struct {
	CountryName string
	RegionName  string
	CityName    string
}

// init
func init() {
	setting.Load()
}

// IPLocate query
func IPLocate(ip string) CityInfo {
	ipdbFilePath := fmt.Sprintf("%s%s", setting.AppConfig.RuntimePath, "ipipfree.ipdb")
	db, err := ipdb.NewCity(ipdbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// db.Reload(ipdbFilePath)
	ipdbRes, ipdbErr := db.FindMap(ip, "CN")
	if ipdbErr != nil {
		log.Fatal(ipdbErr)
	}

	var locate CityInfo
	locate.CountryName = ipdbRes["country_name"]
	locate.RegionName = ipdbRes["region_name"]
	locate.CityName = ipdbRes["city_name"]
	return locate
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("ipLocate", IPLocate)
	// http.ListenAndServe(":8081", service)
	router := gin.Default()
	router.POST("/rpc", func(c *gin.Context) {
		service.ServeHTTP(c.Writer, c.Request)
	})
	router.Run(fmt.Sprintf(":%d", setting.AppConfig.HTTPPort))
}
