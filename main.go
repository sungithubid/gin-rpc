package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/ipipdotnet/ipdb-go"
	"github.com/sungithubid/gin-rpc/pkg/setting"
)

// CityInfo Ip
type CityInfo struct {
	CountryName string `json:"countryName"`
	RegionName  string `json:"regionName"`
	CityName    string `json:"cityName"`
}

// init
func init() {
	setting.Load()
}

// IPLocate query
func IPLocate(ip string, c *gin.Context) string {
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
	data, err := json.Marshal(locate)

	return fmt.Sprintf("%s", data)
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
