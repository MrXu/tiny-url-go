package models

import (
	"time"
)

type Url struct {
	Short_url string
	Full_url string
	create_time 	int64
	active 		int8
}

const (
	full_url_query 		string = "SELECT short_url, full_url, create_time, active FROM url_tab WHERE short_url=?"
	insert_url_query 	string = "INSERT INTO url_tab(short_url,full_url,create_time,active) VALUES(?,?,?,?)"
)

func GetFullUrl(urlName string) (*Url, error){

	var target_url Url
	err := db.QueryRow(full_url_query, urlName).Scan(&target_url.Short_url, &target_url.Full_url, &target_url.create_time, &target_url.active)

	if err != nil{
		return nil,err
	}

	return &target_url, nil
}

func CreateUrl(urlName string, shorten_url string) (*Url, error) {
	stmt, err := db.Prepare(insert_url_query)

	if err != nil {
		return nil, err
	}

	var time_stamp int64 = int64(time.Now().Unix())
	var short_url string = shorten_url
	_, err = stmt.Exec(short_url, urlName, time_stamp, 1)

	if err != nil {
		return nil, err
	}

	return &Url{
		Short_url:short_url,
		Full_url: urlName,
		create_time:time_stamp,
		active:1,
	}, nil
}
