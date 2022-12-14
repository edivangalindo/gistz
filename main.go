package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	// Check for stdin input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No users detected. Hint: cat users.txt | gistz")
		os.Exit(1)
	}

	// Read stdin
	var users []string
	for {
		var user string
		_, err := fmt.Scan(&user)
		if err != nil {
			break
		}
		users = append(users, user)
	}

	// Download all gists for each user
	for _, user := range users {
		downloadGists(user)
	}
}

// Download all gists for a user using Github API
func downloadGists(user string) {
	fmt.Printf("Downloading gists for %s\n", user)

	token := os.Getenv("GH_AUTH_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.GistListOptions{}

	gists, resp, err := client.Gists.List(ctx, user, opt)
	if err != nil {
		if resp.StatusCode == 403 {
			checkRateLimit(resp)
			return
		}

		fmt.Printf("Error: %s\n", err)
		return
	}

	for _, gist := range gists {

		gist, resp, err := client.Gists.Get(ctx, gist.GetID())
		if err != nil {
			if resp.StatusCode == 403 {
				checkRateLimit(resp)
				return
			}

			fmt.Printf("Error: %s\n", err)
			return
		}

		for _, f := range gist.Files {
			// Create directory if it doesn't exist
			os.MkdirAll("gists/"+user+"/"+gist.GetID(), 0755)

			// Write file to disk with filename and content
			fmt.Printf("Writing %s to %s\n", f.GetFilename(), "gists/"+user+"/"+gist.GetID()+"/"+f.GetFilename())
			file, err := os.Create("gists/" + user + "/" + gist.GetID() + "/" + f.GetFilename())
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}

			_, err = file.WriteString(f.GetContent())
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}

			file.Close()

		}
	}
}

func checkRateLimit(resp *github.Response) {
	if resp.Remaining == 0 {
		fmt.Printf("Rate limit exceeded, waitting for %+v minutes\n", time.Until(resp.Rate.Reset.Time).Round(time.Minute).Minutes())
		time.Sleep(time.Duration(time.Until(resp.Rate.Reset.Time).Minutes()) * time.Minute)
	}
}
