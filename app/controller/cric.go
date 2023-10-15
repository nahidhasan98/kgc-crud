package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type Team struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}

type Game struct {
	Bat  Team `json:"bat"`
	Bowl Team `json:"bowl"`
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetShortScore(c *gin.Context) {
	url := "https://www.cricbuzz.com/api/html/homepage-scag"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	checkError(err)

	response, err := client.Do(req)
	checkError(err)
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)

	var games = []Game{}

	document.Find("li.cb-match-card").Each(func(i int, s *goquery.Selection) {
		if i <= 1 {
			var temp = Game{}

			temp.Bat.Name = strings.TrimSpace(s.Find(".cb-hmscg-tm-bat-scr .text-normal").Text())
			s.Find(".cb-hmscg-tm-bat-scr .cb-ovr-flo").Each(func(ii int, ss *goquery.Selection) {
				if ii == 1 {
					temp.Bat.Score = strings.ReplaceAll(strings.TrimSpace(ss.Text()), " ", "")
				}
			})

			temp.Bowl.Name = strings.TrimSpace(s.Find(".cb-hmscg-tm-bwl-scr .text-normal").Text())
			s.Find(".cb-hmscg-tm-bwl-scr .cb-ovr-flo").Each(func(ii int, ss *goquery.Selection) {
				if ii == 1 {
					temp.Bowl.Score = strings.ReplaceAll(strings.TrimSpace(ss.Text()), " ", "")
				}
			})

			games = append(games, temp)
		}
	})

	res := ""

	for key, val := range games {
		res += val.Bat.Name
		if len(val.Bat.Score) > 0 {
			res += ": " + val.Bat.Score
		}

		res += " VS "

		res += val.Bowl.Name
		if len(val.Bowl.Score) > 0 {
			res += ": " + val.Bowl.Score
		}

		if key < len(games)-1 {
			res += " | "
		}
	}

	// fmt.Println(res)
	c.JSON(http.StatusOK, res)
}

func GetLongScore(c *gin.Context) {
	url := "https://www.cricbuzz.com/api/html/homepage-scag"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	checkError(err)

	response, err := client.Do(req)
	checkError(err)
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)

	var games = []Game{}

	document.Find("li.cb-match-card").Each(func(i int, s *goquery.Selection) {
		var temp = Game{}

		temp.Bat.Name = strings.TrimSpace(s.Find(".cb-hmscg-tm-bat-scr .text-normal").Text())
		s.Find(".cb-hmscg-tm-bat-scr .cb-ovr-flo").Each(func(ii int, ss *goquery.Selection) {
			if ii == 1 {
				temp.Bat.Score = strings.ReplaceAll(strings.TrimSpace(ss.Text()), " ", "")
			}
		})

		temp.Bowl.Name = strings.TrimSpace(s.Find(".cb-hmscg-tm-bwl-scr .text-normal").Text())
		s.Find(".cb-hmscg-tm-bwl-scr .cb-ovr-flo").Each(func(ii int, ss *goquery.Selection) {
			if ii == 1 {
				temp.Bowl.Score = strings.ReplaceAll(strings.TrimSpace(ss.Text()), " ", "")
			}
		})

		games = append(games, temp)
	})

	// fmt.Println(games)
	c.JSON(http.StatusOK, games)
}
