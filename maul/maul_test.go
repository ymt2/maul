package maul_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v29/github"

	"github.com/ymt2/maul/maul"
)

func Test_maul_CreateMilestones(t *testing.T) {
	now := time.Now()

	handleFunc := func(ms *github.Milestone) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			v := new(github.Milestone)
			json.NewDecoder(r.Body).Decode(v)

			testMethod(t, r, "POST")
			if diff := cmp.Diff(ms, v); diff != "" {
				t.Errorf("differs: %s\n", diff)
			}

			fmt.Fprint(w, `{"number":1}`)
		}
	}

	type args struct {
		ctx   context.Context
		ms    *github.Milestone
		repos maul.Repositories
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should be pass",
			args: args{
				ctx: context.TODO(),
				ms: &github.Milestone{
					Title:       github.String("title"),
					State:       github.String("open"),
					Description: github.String(""),
					DueOn:       &now,
				},
				repos: maul.Repositories{
					maul.ExportNewRepository("o1", "n1"),
					maul.ExportNewRepository("o1", "n2"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc("/repos/o1/n1/milestones", handleFunc(tt.args.ms))
			mux.HandleFunc("/repos/o1/n2/milestones", handleFunc(tt.args.ms))

			m := maul.New(client)
			if err := m.CreateMilestones(tt.args.ctx, tt.args.ms, tt.args.repos); (err != nil) != tt.wantErr {
				t.Errorf("maul.CreateMilestones() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func setup() (client *github.Client, mux *http.ServeMux, serverURL string, teardown func()) {
	const baseURLPath = "/api-v3"

	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))

	server := httptest.NewServer(apiHandler)

	client = github.NewClient(nil)
	url, _ := url.Parse(server.URL + baseURLPath + "/")
	client.BaseURL = url
	client.UploadURL = url

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}
