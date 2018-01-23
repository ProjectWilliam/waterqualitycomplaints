package main


import (
	"github.com/labstack/echo"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	e := echo.New()
	e.File("/", "public/index.html")
	e.GET("/data", getJSON)
	e.Static("/", "public")
	e.Logger.Fatal(e.Start(":8000"))
}
func getJSON(c echo.Context) error {
	url := "https://data.cityofnewyork.us/resource/qfe3-6dkn.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	return c.String(http.StatusOK, responseString)

}