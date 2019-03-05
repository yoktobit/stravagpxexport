package strava

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/yoktobit/stravagpxexport/util"
	"golang.org/x/oauth2"
)

// Session stellt eine Strava-Session dar
type Session struct {
	Context context.Context
	Token   oauth2.Token
}

// Login logs you in into strava (using Windows default desktop browser)
func (s *Session) Login() *http.Client {
	s.Context = context.Background()

	var tokenChannel = make(chan string)

	token := new(oauth2.Token)
	err := util.ReadGob("token.db", token)
	if err != nil {
		token = nil
	}

	conf := &oauth2.Config{
		ClientID:     os.Getenv("STRAVA_CLIENT_ID"),
		ClientSecret: os.Getenv("STRAVA_CLIENT_SECRET"),
		Scopes:       []string{"activity:read_all"},
		RedirectURL:  "http://127.0.0.1:44444",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.strava.com/oauth/mobile/authorize",
			TokenURL: "https://www.strava.com/oauth/token",
		},
	}

	if token == nil {
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		go exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tokenCode := r.FormValue("code")
			fmt.Fprintf(w, `You can close your Browser now`)
			tokenChannel <- tokenCode
		})
		go http.ListenAndServe(":44444", nil)

		tokenCodeResult := <-tokenChannel

		token, err = conf.Exchange(s.Context, tokenCodeResult)
		if err != nil {
			panic(err)
		}
	}
	util.WriteGob("token.db", token)

	return conf.Client(s.Context, token)
}

// NewSession creates a new Strava API session
func NewSession() *Session {
	return new(Session)
}
