package main

import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


type Recording struct {
	EndTime string `xml:"end_time"`
	ID      string `xml:"id"`
	Meeting struct {
		_breakout   string `xml:"@breakout"`
		_externalID string `xml:"@externalId"`
		_id         string `xml:"@id"`
		_name       string `xml:"@name"`
	} `json:"meeting"`
	Meta struct {
		IsBreakout  string `xml:"isBreakout"`
		MeetingID   string `xml:"meetingId"`
		MeetingName string `xml:"meetingName"`
	} `xml:"meta"`
	Participants string `xml:"participants"`
	Playback     struct {
		Duration   string `xml:"duration"`
		Extensions struct {
			Preview struct {
				Images []struct {
					_text   string `xml:"#text"`
					_alt    string `xml:"@alt"`
					_height string `xml:"@height"`
					_width  string `xml:"@width"`
				} `xml:"images"`
			} `xml:"preview"`
		} `xml:"extensions"`
		Format         string `xml:"format"`
		Link           string `xml:"link"`
		ProcessingTime string `xml:"processing_time"`
		Size           string `xml:"size"`
	} `xml:"playback"`
	Published string `xml:"published"`
	RawSize   string `xml:"raw_size"`
	StartTime string `xml:"start_time"`
	State     string `xml:"state"`
}

func main() {





	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization"}
	config.AllowMethods = []string{"OPTIONS", "GET", "POST"}
	router.Use(cors.New(config))



	router.GET("/duration/:path", Duration)

	router.Run(":8070")




}


func Duration(c *gin.Context) {


	//u := c.PostForm("userid")

	var str Recording

	file:= c.Param("path") + "/metadata.xml"
	fmt.Println(file)
	bs, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("file not found")
		return
	}

	xml.Unmarshal(bs, &str)
	fmt.Println(str.Playback.Duration)

	c.String(200, str.Playback.Duration)

}