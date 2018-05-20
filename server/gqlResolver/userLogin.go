package gqlResolver

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/utils"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// IMPORTANT DOCS
// oauth2 official google example https://github.com/golang/oauth2/blob/master/google/example_test.go
// OAuth 2.0 authentication in a Golang CLI https://gist.github.com/marians/3b55318106df0e4e648158f1ffb43d38
var PublicKey = []byte("secret")

var users = []model.User{
	{
		ID:       "1",
		Username: "seb",
		Password: "seb",
	},
	{
		ID:       "2",
		Username: "bb",
		Password: "bb",
	},
}

// UserLogin :
// - Returns nil if no args are provided
// - Inserts a new user into mongodb
// - Resolves the current user if args are provided
func (r *Resolver) UserLogin(ctx context.Context, args *struct {
	Input *model.UserLoginInput
}) (string, error) {

	conf := &oauth2.Config{
		ClientID:     viper.GetString("lychee.oauth.google_client_id"),
		ClientSecret: viper.GetString("lychee.oauth.google_client_secret"),
		RedirectURL:  "http://google.com",
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
	// // Redirect user to Google's consent page to ask for permission
	// // for the scopes specified above.
	// url := conf.AuthCodeURL("state")
	// log.Printf("Visit the URL for the auth dialog: %v", url)

	// // Handle the exchange code to initiate a transport.
	// tok, err := conf.Exchange(ctx, "authorization-code")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client := conf.Client(ctx, tok)

	// add transport for self-signed certificate to context
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)
	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	log.Println(color.CyanString("You will now be taken to your browser for authentication"))
	time.Sleep(1 * time.Second)
	open.Run(url)
	time.Sleep(1 * time.Second)
	log.Printf("Authentication URL: %s\n", url)
	spew.Dump(url)
	for _, user := range users {
		if user.Username == args.Input.Username {
			if user.Password == args.Input.Password {
				token, err := utils.GenerateToken(user)
				if err != nil {
					return "", err
				}
				return token, err
			} else {
				return "", errors.New("password is incorrect")
			}
		}
	}

	return "", errors.New("User not found")
}
