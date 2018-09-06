package movie

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Movie struct {
	Rank       string
	Title      string
	ForeiTitle string
	OtherTitle string
	RatingNum  string
	People     string
	Time       string
	Country    string
	Tag        string
	Info       string
	Img        string
}

var movies []Movie

func GetMovie() []Movie {
	for i := 0; i <= 250; i = i + 25 {
		//create request
		client := &http.Client{}
		//生成要访问的url
		url := "https://movie.douban.com/top250?start=" + strconv.Itoa(i) + "&filter="
		//提交请求
		reqest, err := http.NewRequest("GET", url, nil)

		//增加header选项
		reqest.Header.Add("Cookie", "bid=uAVGuPbz_Lg; douban-fav-remind=1; __utmc=30149280; __utmz=30149280.1536113992.2.2.utmcsr=github.com|utmccn=(referral)|utmcmd=referral|utmcct=/EDDYCJY/blog/blob/master/golang/crawler/2018-03-21-%E7%88%AC%E5%8F%96%E6%9C%80%E7%AE%80%E5%8D%95%E7%9A%84%E8%B1%86%E7%93%A3%E7%94%B5%E5%BD%B1-Top250.md; __utmc=223695111; __utmz=223695111.1536113992.1.1.utmcsr=github.com|utmccn=(referral)|utmcmd=referral|utmcct=/EDDYCJY/blog/blob/master/golang/crawler/2018-03-21-%E7%88%AC%E5%8F%96%E6%9C%80%E7%AE%80%E5%8D%95%E7%9A%84%E8%B1%86%E7%93%A3%E7%94%B5%E5%BD%B1-Top250.md; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1536129484%2C%22https%3A%2F%2Fgithub.com%2FEDDYCJY%2Fblog%2Fblob%2Fmaster%2Fgolang%2Fcrawler%2F2018-03-21-%25E7%2588%25AC%25E5%258F%2596%25E6%259C%2580%25E7%25AE%2580%25E5%258D%2595%25E7%259A%2584%25E8%25B1%2586%25E7%2593%25A3%25E7%2594%25B5%25E5%25BD%25B1-Top250.md%22%5D; _pk_ses.100001.4cf6=*; __utma=30149280.1320413867.1535526167.1536117507.1536129484.4; __utmb=30149280.0.10.1536129484; __utma=223695111.2044650758.1536113992.1536117507.1536129484.3; __utmb=223695111.0.10.1536129484; ap_v=0,6.0; _pk_id.100001.4cf6=82ec3166399112a1.1536113992.3.1536129966.1536117509.")
		reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

		if err != nil {
			panic(err)
		}
		//处理返回结果
		response, _ := client.Do(reqest)
		defer response.Body.Close()

		doc, err := goquery.NewDocumentFromReader(response.Body)

		if err != nil {

			log.Fatal(err)

		}

		doc.Find(".grid_view li").Each(func(i int, contentSelection *goquery.Selection) {

			rank := contentSelection.Find(".pic em").Text()
			title := contentSelection.Find(".title").Eq(0).Text()                                                //电影名
			foreititle := strings.TrimLeft(strings.TrimSpace(contentSelection.Find(".title").Eq(1).Text()), "/") //电影名

			otherTitle := strings.TrimSpace(contentSelection.Find(".other").Text()) //别名

			subscribe := contentSelection.Find(".bd p").Eq(0).Text() //导演，主演，时间，国家，tag混在一起

			subInfo := strings.Split(subscribe, "\n") //把导演和时间，国家，tag用\n先分开
			//混杂导演，主演
			allpeople := strings.TrimSpace(subInfo[1])

			//混杂时间，国家，tag
			timeTag := strings.Split(subInfo[2], "/")

			time := strings.TrimSpace(timeTag[0])
			country := strings.TrimSpace(timeTag[1])
			tag := strings.TrimSpace(timeTag[2])

			intro := contentSelection.Find(".bd p").Eq(1).Text() //简介
			intro = strings.TrimSpace(intro)
			ratingNum := contentSelection.Find(".rating_num").Text() //评分
			img, _ := contentSelection.Find(".pic a img").Attr("src")

			movie := Movie{
				Rank:       rank,
				Title:      title,
				ForeiTitle: foreititle,
				OtherTitle: otherTitle,
				RatingNum:  ratingNum,
				People:     allpeople,
				Time:       time,
				Country:    country,
				Tag:        tag,
				Info:       intro,
				Img:        img,
			}

			movies = append(movies, movie)
		})

	}

	return movies
}
