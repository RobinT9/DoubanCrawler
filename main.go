package main

import (
	"database/sql"
	"log"
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
	for _, v := range movies {
		prepare, preerr := doubanDB.Prepare("insert into top_movie_copy1(rank,title,foreititle,othertitle,rating_num,people_info,on_time,country,tag,info,img) values (?,?,?,?,?,?,?,?,?,?,?)")
		if preerr != nil {
			log.Print("Insert Error:", preerr)
			return
		}
		_, execerr := prepare.Exec(v.Rank, v.Title, v.ForeiTitle, v.OtherTitle, v.RatingNum, v.People, v.Time, v.Country, v.Tag, v.Info, v.Img)
		if execerr != nil {
			log.Print("Exec Error:", execerr.Error())
			return
		}
	}

	log.Println("Done")
}
