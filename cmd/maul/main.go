package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"

	"github.com/ymt2/maul/config"
	"github.com/ymt2/maul/maul"
)

var (
	repos maul.Repositories

	dueAfter = flag.Int("due-after", 0, "Due after n days")
	title    = flag.String("title", "", "Title")
)

func init() {
	flag.Var(&repos, "repo", "Some description for this param.")
}

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(args []string) int {
	flag.Parse()
	env := config.ReadFromEnv()

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: env.GitHubAuthToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	dueOn := time.Now().AddDate(0, 0, *dueAfter)
	ms := &github.Milestone{
		Title:       title,
		State:       github.String("open"),
		Description: github.String(""),
		DueOn:       &dueOn,
	}

	m := maul.New(client)
	if err := m.CreateMilestones(ctx, ms, repos); err != nil {
		log.Fatal(err)
	}

	return 0
}
