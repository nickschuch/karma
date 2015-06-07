package main

import (
	"log"
	"net/http"
	"runtime"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v1"

	"github.com/nickschuch/karma/storage"
	_ "github.com/nickschuch/karma/storage/dynamodb"
	_ "github.com/nickschuch/karma/storage/memory"
)

var (
	cliPort     = kingpin.Flag("port", "Port to run the bot on.").Default("80").OverrideDefaultFromEnvar("KARMA_PORT").String()
	cliToken    = kingpin.Flag("token", "Token to keep this bot secure.").Default("").OverrideDefaultFromEnvar("KARMA_TOKEN").String()
	cliBackend  = kingpin.Flag("storage", "Storage backend for keeping karma.").Default("memory").OverrideDefaultFromEnvar("KARMA_STORAGE").String()
	cliCallback = kingpin.Flag("callback", "The URL as setup in https://api.slack.com/incoming-webhooks").Default("").OverrideDefaultFromEnvar("KARMA_CALLBACK").String()
	cliBotName  = kingpin.Flag("name", "Name your bot.").Default("Karma").OverrideDefaultFromEnvar("KARMA_NAME").String()
	cliBotEmoji = kingpin.Flag("emoji", "Give your bot a custom image.").Default(":slack:").OverrideDefaultFromEnvar("KARMA_EMOJI").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.CommandLine.Help = "Karma bot for Slack."
	kingpin.Parse()

	// This allows us to serve more than a single request at a time.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Start up the webserver.
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+*cliPort, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var user string
	var amount int

	// Get the values posted from Slack.
	r.ParseForm()
	token := r.Form.Get("token")
	text := r.Form.Get("text")
	author := r.Form.Get("user_name")
	channel := r.Form.Get("channel_name")

	// We need to ensure the the request has the correct token. Otherwise anyone
	// can steal our karma!
	if *cliToken != token {
		log.Println("Invalid token", token)
		return
	}

	// Build a response object which we can use to send back to Slack.
	response := Response{
		Username: *cliBotName,
		Emoji:    *cliBotEmoji,
		Channel:  "#" + channel,
	}
	response.Send(*cliCallback)

	// Now that we have gone through all the check we can connect to the backend.
	s, err := storage.New(*cliBackend)
	if err != nil {
		log.Println("Cannot start the backend: %v", cliBackend)
		return
	}

	// Now we attempt to find out which user this request is for.
	user, err = getUser(text)
	if err != nil {
		amount = s.Get(author)
		response.Text = author + " has " + strconv.Itoa(amount) + " karma"
		response.Send(*cliCallback)
		return
	}

	// Check for increase request.
	amount = increaseAmount(text)
	if amount > 0 {
		s.Increase(user, amount)
		response.Text = author + " gave " + user + " " + strconv.Itoa(amount) + " karma"
		response.Send(*cliCallback)
		return
	}

	// Check for decrease request.
	amount = decreaseAmount(text)
	if amount > 0 {
		s.Decrease(user, amount)
		response.Text = author + " deducted " + strconv.Itoa(amount) + " karma off " + user
		response.Send(*cliCallback)
		return
	}

	// By this stage I think we can assume the user wants the amount associated
	// with a user.
	amount = s.Get(user)
	response.Text = author + " has " + strconv.Itoa(amount) + " karma"
	response.Send(*cliCallback)
}
