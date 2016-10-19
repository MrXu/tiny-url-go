package models

import (
	"time"
	"tinyUrl/utils"
)

type Url struct {
	short_url 	string
	full_url 	string
	create_time 	int64
	active 		int8
}

const (
	full_url_query 		string = "SELECT * FROM url_tab WHERE short_url=$1"
	insert_url_query 	string = "INSERT INTO url_tab(short_url,full_url,create_time,active) VALUES($1,$2,$3,$4)"
)

func GetFullUrl(urlName string) (*Url, error){

	var target_url Url
	err := db.QueryRow(full_url_query, urlName).Scan(&target_url)

	if err != nil{
		return nil,err
	}

	return target_url, nil
}

func CreateUrl(urlName string) (*Url, error) {
	stmt, err := db.Prepare(insert_url_query)

	if err != nil {
		return nil, err
	}

	var time_stamp int64 = int64(time.Now().Unix())
	var short_url string = utils.UrlShortener(urlName)
	_, err = stmt.Exec(short_url, urlName, time_stamp, 1)

	if err != nil {
		return nil, err
	}

	return &Url{
		short_url:short_url,
		full_url: urlName,
		create_time:time_stamp,
		active:1,
	}, nil
}
