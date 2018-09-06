package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"test/hservice"
	"topMovieCr/movie"
)

var movies []movie.Movie

func main() {

	db := hservice.Db{
		ConnType: "mysql",
		Uname:    "root",
		Password: "",
		Ip:       "127.0.0.1",
		Port:     3306,
		Database: "douban",
	}
	var doubanDB *sql.DB
	doubanDB, err := db.DbInit()
	defer doubanDB.Close()
	if err != nil {
		panic(err)
	}

	movies = movie.GetMovie()
	//sql保存
	//for _, v := range movies {
	//	prepare, preerr := doubanDB.Prepare("insert into top_movie_copy1(rank,title,foreititle,othertitle,rating_num,people_info,on_time,country,tag,info,img) values (?,?,?,?,?,?,?,?,?,?,?)")
	//	if preerr != nil {
	//		log.Print("Insert Error:", preerr)
	//		return
	//	}
	//	_, execerr := prepare.Exec(v.Rank, v.Title, v.ForeiTitle, v.OtherTitle, v.RatingNum, v.People, v.Time, v.Country, v.Tag, v.Info, v.Img)
	//	if execerr != nil {
	//		log.Print("Exec Error:", execerr.Error())
	//		return
	//	}
	//}

	//文件保存
	jsondata, jsonerr := json.Marshal(movies)
	if jsonerr != nil {
		panic(jsonerr)
	}
	fmt.Printf("%s\n", jsondata)
	fileObj, err := os.OpenFile("topmovie.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	io.WriteString(fileObj, string(jsondata))

	log.Println("Done")
}
