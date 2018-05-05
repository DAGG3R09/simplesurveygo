package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

// MgoSession : Sessoin used to export
var MgoSession *mgo.Session

func init() {
	if MgoSession == nil {
		var err error
		MgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
	}
}

type Movie struct {
	Title    string `json:"title" bson:"title"`
	Year     int    `json:"year" bson:"year"`
	Director string `json:"director" bson:"director"`
	Cast     string `json:"cast" bson:"cast"`
	Genre    string `json:"genre" bson:"genre"`
	Notes    string `json:"notes" bson:"notes"`
}

func InsertMovie(movie Movie) {
	session := MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("Movie").C("movies")
	clctn.Insert(movie)
}

func Insertor(movies []Movie) {
	fmt.Println("in Go routine")
	for _, movie := range movies {
		InsertMovie(movie)
	}
	<-counter
}

var counter = make(chan bool, 4)

func main() {
	url := "https://raw.githubusercontent.com/prust/wikipedia-movie-data/master/movies.json"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var movies []Movie

	err = json.Unmarshal(data, &movies)

	if err != nil {
		panic(err)
	}

	lenM := len(movies)
	chunckSize := lenM / 4

	// counter <- true
	// counter <- true
	// counter <- true
	// counter <- true

	fmt.Println(lenM)
	for i := 0; i < lenM; i += chunckSize {
		fmt.Println(i)
		go Insertor(movies[i : i+chunckSize])
	}

	fmt.Scanln()
}
