package codeowners

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/gofri/go-github-pagination/githubpagination"
	"github.com/google/go-github/v64/github"
)

type Options struct {
	User  string `short:"u" xor:"user,org" required:"" help:"Organization name."`
	Org   string `short:"o" xor:"org" required:"" help:"Organization name."`
	Token string `required:"" env:"GITHUB_TOKEN" help:"Authentication token for github.com API requests."`
}

type Codeowners struct {
	Repo          string `json:"repo"`
	HasCODEOWNERS bool   `json:"has_codeowners"`
	Content       string `json:"content,omitempty"`
}

func List(ctx context.Context, options *Options) ([]*Codeowners, error) {
	paginator := githubpagination.NewClient(nil)
	client := github.NewClient(paginator).WithAuthToken(options.Token)
	repos, err := listRepos(ctx, client, options)

	if err != nil {
		return nil, err
	}

	cos := []*Codeowners{}

	for _, r := range repos {
		if r.Archived != nil && *r.Archived {
			continue
		}

		co, err := getCodeowners(ctx, client, r)

		if err != nil {
			return nil, err
		}

		cos = append(cos, co)
	}

	return cos, nil
}

func listRepos(ctx context.Context, client *github.Client, options *Options) ([]*github.Repository, error) {
	var repos []*github.Repository
	var err error

	if options.Org != "" {
		repos, _, err = client.Repositories.ListByOrg(ctx, options.Org, &github.RepositoryListByOrgOptions{})
	} else {
		repos, _, err = client.Repositories.ListByUser(ctx, options.User, &github.RepositoryListByUserOptions{})
	}

	if err != nil {
		return nil, err
	}

	return repos, nil
}

func getCodeowners(ctx context.Context, client *github.Client, repo *github.Repository) (*Codeowners, error) {
	raw, _, _, err := client.Repositories.GetContents(ctx, *repo.Owner.Login, *repo.Name, ".github/CODEOWNERS", &github.RepositoryContentGetOptions{})

	co := &Codeowners{
		Repo:          *repo.Name,
		HasCODEOWNERS: false,
	}

	if err != nil {
		var errResp *github.ErrorResponse

		if errors.As(err, &errResp) && errResp.Response.StatusCode == http.StatusNotFound {
			return co, nil
		} else {
			return nil, err
		}
	}

	content, err := base64.StdEncoding.DecodeString(*raw.Content)

	if err != nil {
		panic(err)
	}

	co.Content = string(content)
	co.HasCODEOWNERS = true

	return co, nil
}
