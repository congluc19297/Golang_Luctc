package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

type matchToday struct {
	Venue             string      `json:"venue"`
	Location          string      `json:"location"`
	Status            string      `json:"status"`
	Time              string      `json:"time"`
	FifaID            string      `json:"fifa_id"`
	Datetime          string      `json:"datetime"`
	LastEventUpdateAt string      `json:"last_event_update_at"`
	LastScoreUpdateAt string      `json:"last_score_update_at"`
	HomeTeam          infoTeam    `json:"home_team"`
	AwayTeam          infoTeam    `json:"away_team"`
	Winner            string      `json:"winner"`
	WinnerCode        string      `json:"winner_code"`
	HomeTeamEvents    []teamEvent `json:"home_team_events"`
	AwayTeamEvents    []teamEvent `json:"away_team_events"`
}

// Last_event_update_at // Last_score_update_at
// home_team // away_team
// winner_code
// home_team_events // away_team_events

type infoTeam struct {
	Country string `json:"country"`
	Code    string `json:"code"`
	Goals   int    `json:"goals"`
}

type teamEvent struct {
	ID          string `json:"id"`
	TypeOfEvent string `json:"type_of_event"`
	Player      string `json:"player"`
	Time        string `json:"time"`
}

func main() {
	// fmt.Println("Before Get data")
	response, err := http.Get("http://worldcup.sfg.io/matches/today")
	// fmt.Println("After Get data")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}

	// fmt.Println(string(responseData))
	var matchesToday []matchToday
	err = json.Unmarshal(responseData, &matchesToday)
	if err != nil {
		panic(err)
	}
	fmt.Println(matchesToday)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseData))
	})
	http.ListenAndServe(":3000", r)
}
