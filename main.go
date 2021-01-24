package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	// Just for -help message
	flag.Parse()

	if !IsExecutable("gh") {
		log.Fatal(`please install the github cli tool "gh"`)
	}

	if !GHIsLoggedIn() {
		log.Fatal(`please log in to github with "gh auth login"`)
	}

	user, _, err := GHGetConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(`getting pull requests for user "%s"`, user)
	prs, err := GHGetPulls(user)
	if err != nil {
		log.Fatal(err)
	}

	if len(prs) == 0 {
		log.Println("no prs to merge")
		os.Exit(0)
	}

	result, tui := NewTUI(prs)

	if err := tui.Start(); err != nil {
		log.Fatal(err)
	}

	r := <-result

	if r == nil {
		log.Println("nothing to do...exiting")
		os.Exit(0)
	}

	log.Printf("merging %s", r.url)
	GHMergePR(r.url)
}
