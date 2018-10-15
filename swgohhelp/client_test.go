package swgohhelp_test

import (
	"context"
	"flag"
	"testing"

	"github.com/ronoaldo/swgoh/swgohhelp"
)

var username, password, accessToken string

func init() {
	flag.StringVar(&username, "username", "", "Username to authenticate to the API")
	flag.StringVar(&password, "password", "", "Password to authenticate to the API")
	flag.StringVar(&accessToken, "access-token", "", "API Access Token to authenticate to the API")
}
func TestClientAuth(t *testing.T) {
	if username == "" || password == "" {
		t.Skip("Auth tests skipped due to missing credentials. Use -username and -passoword test flags.")
	}

	c := swgohhelp.New(context.Background())
	token, err := c.SignIn(username, password)
	if err != nil {
		t.Fatalf("Unexpected error signing in: %v", err)
	}
	if token == "" {
		t.Fatalf("Unexpected empty token after auth!")
	}

	t.Logf("Auth success; token=%s", token)
}

func TestPlayer(t *testing.T) {
	if accessToken == "" {
		t.Skip("Ignoring TestPlayer call due to missing authentication. Use -access-token test parameter.")
	}

	// c := swgohapi.New(context.Background())
}
