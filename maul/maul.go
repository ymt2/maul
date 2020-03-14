package maul

import (
	"context"
	"fmt"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/google/go-github/v29/github"
	"golang.org/x/sync/errgroup"
)

type maul struct {
	client *github.Client
}

// CreateMilestone creates a new milestone on the specified repository.
func (m *maul) CreateMilestones(ctx context.Context, ms *github.Milestone, repos Repositories) error {
	var res sync.Map

	eg, ctx := errgroup.WithContext(ctx)
	for _, repo := range repos {
		r := repo
		eg.Go(func() error {
			_, _, err := m.client.Issues.CreateMilestone(ctx, r.owner, r.name, ms)
			res.Store(r, err)
			return err
		})
	}
	err := eg.Wait()

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintln(w, "REPOSITORY\tSTATUS")
	for _, r := range repos {
		var status string
		if v, ok := res.Load(r); ok {
			if e, ok := v.(error); ok {
				status = e.Error()
			} else {
				status = "Created"
			}
		} else {
			status = "Skipped"
		}

		name := r.owner + "/" + r.name
		fmt.Fprintf(w, "%s\t%s\n", name, status)
	}
	fmt.Fprintln(w, "")

	return err
}

// New creates a maul.
func New(client *github.Client) *maul {
	return &maul{
		client: client,
	}
}
