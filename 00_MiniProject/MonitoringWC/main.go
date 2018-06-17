package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/smtp"
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

type infoTeam struct {
	Country string `json:"country"`
	Code    string `json:"code"`
	Goals   int    `json:"goals"`
}

type teamEvent struct {
	ID          int    `json:"id"`
	TypeOfEvent string `json:"type_of_event"`
	Player      string `json:"player"`
	Time        string `json:"time"`
}

var listID []string

func main() {
	var matchsToday []matchToday
	// c := cron.New()
	// c.AddFunc("1 * * * * *", monitoringWC)
	// fmt.Println("Before Start")
	// c.Run()
	monitoringWC(&matchsToday)
	// c.Stop()
}

func monitoringWC(matchsToday *[]matchToday) {
	// fmt.Println((*matchsToday)[0])
	// fmt.Println("Len Before fetch data: ", len(*matchsToday))
	err := fetchMatchsToday(&(*matchsToday))
	if err != nil {
		panic(err)
	}
	if len(*matchsToday) != 0 {
		// fmt.Println("Len After fetch data: ", len(*matchsToday))
		var lenMatchs = len(*matchsToday)
		for i := 0; i < lenMatchs; i++ {
			if (*matchsToday)[i].Status == "completed" && !checkFifaIDSended((*matchsToday)[i].FifaID) {
				listID = append(listID, (*matchsToday)[i].FifaID)
				// fmt.Println(listID)
				err = sendMailToUsers((*matchsToday)[i])
				if err == nil {
					delElementArr(i, &(*matchsToday))
					lenMatchs--
					i--
					continue
				}
			}
		}
		// fmt.Println("Len After delete data: ", len(*matchsToday))
	}
}

func sendMailToUsers(match matchToday) error {
	// Set up authentication information.
	users := []string{"congluc19297@gmail.com", "1511917@hcmut.edu.vn"}

	auth := smtp.PlainAuth(
		"",
		"olivertran732169@gmail.com",
		"ttcnpm_nhom21",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: Kết quả thi đấu WorldCup" + "!\n"
	body := emailTemplateHTML(match)
	err := smtp.SendMail(
		"smtp.gmail.com:25",
		auth,
		"olivertran732169@gmail.com",
		users,
		[]byte(subject+mime+"\n"+body),
	)
	if err != nil {
		return err
	}
	return nil
}

func emailTemplateHTML(match matchToday) string {
	var err error
	t, err := template.ParseFiles("EmailTemplate.gohtml")
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, match); err != nil {
		panic(err)
	}

	result := tpl.String()
	return result
}

func delElementArr(index int, matchsToday *[]matchToday) {
	// func append(slice []Type, elems ...Type) []Type
	*matchsToday = append((*matchsToday)[:index], (*matchsToday)[index+1:]...)
}

func checkFifaIDSended(fifaID string) bool {
	var lenFifaID = len(listID)
	for i := 0; i < lenFifaID; i++ {
		if fifaID == listID[i] {
			return true
		}
	}
	return false
}

func fetchMatchsToday(matchsToday *[]matchToday) error {
	// fmt.Println((*matchsToday)[0])
	res, err := http.Get("http://worldcup.sfg.io/matches/today")
	if err != nil {
		return err
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resData, &(*matchsToday))
	if err != nil {
		return err
	}
	return nil
}
