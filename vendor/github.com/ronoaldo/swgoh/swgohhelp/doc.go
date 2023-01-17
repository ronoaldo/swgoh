/*
Package swgohhelp provides the API client interface for the
https://api.swgoh.help/ service.


How to use

This package provides typed datas structures to interface with the
website payloads and higher level methods for consuming the services.

You start by initializing the client and authenticating:

	swcli := swgohhelp.New(context.Background())
	if err := swcli.SignIn(username, password); err != nil {
		log.Fatalf("Unable to authenticate: %v", err)
	}

There is no current session state so you may need to reauth by hand.


Fetching player profiles

You call the Client.Players() method, passing in the array of desired
ally codes to be fetched:

	players, err := swcli.Player("335-983-287")
	if err != nil {
		log.Fatalf("Unable to fetch players: %v". err)
	}
	log.Printf("Found %d players", len(players))
	for i := range players {
		log.Printf("Player #%d: %s, last updated at %v", i, players[i], players[i].UpdatedAt)
	}


Debugging your API calls

Sometimes it may be usefull to see the raw request and response
data sent/received by the client. Use the .SetDebug method to enable
file-based logging of each request and response payloads to the
default OS temporary folder.

*/
package swgohhelp
