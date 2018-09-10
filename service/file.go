package service

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"topMovieCr/movie"
)

func SaveToJson(movies []movie.Movie) error {
	//文件保存
	jsondata, jsonerr := json.Marshal(movies)
	if jsonerr != nil {
		panic(jsonerr)
	}
	fileObj, err := os.OpenFile("topmovie.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	_, writeErr := io.WriteString(fileObj, string(jsondata))
	return writeErr
}
