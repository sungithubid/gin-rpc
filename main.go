package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/ipipdotnet/ipdb-go"
)

// CityInfo Ip
type CityInfo struct {
	CountryName string
	RegionName  string
	CityName    string
}

// IPLocate query
func IPLocate(ip string) CityInfo {
	db, err := ipdb.NewCity("runtime/ipipfree.ipdb")
	if err != nil {
		log.Fatal(err)
	}

	// db.Reload("runtime/ipipfree.ipdb")
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
	router.Any("/rpc", func(c *gin.Context) {
		service.ServeHTTP(c.Writer, c.Request)
	})
	router.Run(":8081")
}
