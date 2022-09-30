package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"sort"

	"gopkg.in/yaml.v3"
)

func GHIsLoggedIn() bool {
	cmd := exec.Command("gh", "auth", "status")
	err := cmd.Run()
	return err == nil // lazy but fine for now
}

func GHGetConfig() (string, string, error) {
	user, err := user.Current()
	if err != nil {
		return "", "", err
	}

	cfgPath := path.Join(user.HomeDir, ".config", "gh", "hosts.yml")

	cfgBytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return "", "", err
	}

	var cfgs map[string]struct {
		User  string `yaml:"user"`
		Token string `yaml:"oauth_token"`
	}

	if err = yaml.Unmarshal(cfgBytes, &cfgs); err != nil {
		return "", "", err
	}

	ghdotcomCfg, found := cfgs["github.com"]
	if !found {
		return "", "", fmt.Errorf(`could not find "github.com" host configuration for "gh" cli`)
	}

	return ghdotcomCfg.User, ghdotcomCfg.Token, nil
}

type GHOpenPR struct {
	org  string
	repo string

	title  string
	number int
	url    string
}

func (p GHOpenPR) Display() string {
	return fmt.Sprintf(
		"%s/%s [%d] %s %s",
		p.org, p.repo,
		p.number, p.url, p.title,
	)
}

func GHGetPulls(user string) ([]GHOpenPR, error) {
	query := fmt.Sprintf(`query={
    search(query: "org:%s is:pr is:open" type: ISSUE last: 100) {
        edges {
          node {
            ... on PullRequest {
              headRepository {
              isArchived
              isFork
              name
              owner { login }
            }
            number
            mergeable
            url
            title
          }
        }
      }
    }
	}`, user)

	cmd := exec.Command("gh", "api", "graphql", "-f", query)
	o, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var results struct {
		Data struct {
			Search struct {
				Edges []struct {
					Node struct {
						Repo struct {
							IsArchived bool   `json:"isArchived"`
							IsFork     bool   `json:"isFork"`
							Name       string `json:"name"`

							Owner struct {
								Login string `json:"login"`
							} `json:"owner"`
						} `json:"headRepository"`
						Mergeable string `json:"mergeable"`
						Number    int    `json:"number"`
						Title     string `json:"title"`
						URL       string `json:"url"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"search"`
		} `json:"data"`
	}

	if err := json.Unmarshal(o, &results); err != nil {
		return nil, err
	}

	prs := []GHOpenPR{}
	for _, n := range results.Data.Search.Edges {
		node := n.Node

		if node.Repo.IsArchived || node.Repo.IsFork || node.Mergeable != "MERGEABLE" {
			continue
		}

		prs = append(prs, GHOpenPR{
			org:    node.Repo.Owner.Login,
			repo:   node.Repo.Name,
			title:  node.Title,
			number: node.Number,
			url:    node.URL,
		})
	}

	sort.Slice(prs, func(i, j int) bool {
		return prs[i].repo <= prs[j].repo
	})

	return prs, nil
}

func GHMergePR(url string) error {
	// use a temporary directory so local branch detection doesn't recurse anywhere sensitive
	dir, err := ioutil.TempDir("", "merge-mgr")
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	cmd := exec.Command("gh", "pr", "merge", "-m", url)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir

	return cmd.Run()
}
