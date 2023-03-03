// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package uservalidator

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/app-sre/go-qontract-reconcile/pkg/gql"
)

// GithubOrgsGithuborg_v1GithubOrg_v1 includes the requested fields of the GraphQL type GithubOrg_v1.
type GithubOrgsGithuborg_v1GithubOrg_v1 struct {
	Name    string                                                `json:"name"`
	Token   GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1 `json:"token"`
	Default bool                                                  `json:"default"`
}

// GetName returns GithubOrgsGithuborg_v1GithubOrg_v1.Name, and is useful for accessing the field via an interface.
func (v *GithubOrgsGithuborg_v1GithubOrg_v1) GetName() string { return v.Name }

// GetToken returns GithubOrgsGithuborg_v1GithubOrg_v1.Token, and is useful for accessing the field via an interface.
func (v *GithubOrgsGithuborg_v1GithubOrg_v1) GetToken() GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1 {
	return v.Token
}

// GetDefault returns GithubOrgsGithuborg_v1GithubOrg_v1.Default, and is useful for accessing the field via an interface.
func (v *GithubOrgsGithuborg_v1GithubOrg_v1) GetDefault() bool { return v.Default }

// GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1 includes the requested fields of the GraphQL type VaultSecret_v1.
type GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1 struct {
	Path    string `json:"path"`
	Field   string `json:"field"`
	Version int    `json:"version"`
	Format  string `json:"format"`
}

// GetPath returns GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1.Path, and is useful for accessing the field via an interface.
func (v *GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1) GetPath() string { return v.Path }

// GetField returns GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1.Field, and is useful for accessing the field via an interface.
func (v *GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1) GetField() string { return v.Field }

// GetVersion returns GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1.Version, and is useful for accessing the field via an interface.
func (v *GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1) GetVersion() int { return v.Version }

// GetFormat returns GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1.Format, and is useful for accessing the field via an interface.
func (v *GithubOrgsGithuborg_v1GithubOrg_v1TokenVaultSecret_v1) GetFormat() string { return v.Format }

// GithubOrgsResponse is returned by GithubOrgs on success.
type GithubOrgsResponse struct {
	Githuborg_v1 []GithubOrgsGithuborg_v1GithubOrg_v1 `json:"githuborg_v1"`
}

// GetGithuborg_v1 returns GithubOrgsResponse.Githuborg_v1, and is useful for accessing the field via an interface.
func (v *GithubOrgsResponse) GetGithuborg_v1() []GithubOrgsGithuborg_v1GithubOrg_v1 {
	return v.Githuborg_v1
}

// UsersResponse is returned by Users on success.
type UsersResponse struct {
	Users_v1 []UsersUsers_v1User_v1 `json:"users_v1"`
}

// GetUsers_v1 returns UsersResponse.Users_v1, and is useful for accessing the field via an interface.
func (v *UsersResponse) GetUsers_v1() []UsersUsers_v1User_v1 { return v.Users_v1 }

// UsersUsers_v1User_v1 includes the requested fields of the GraphQL type User_v1.
type UsersUsers_v1User_v1 struct {
	Path               string `json:"path"`
	Name               string `json:"name"`
	Org_username       string `json:"org_username"`
	Github_username    string `json:"github_username"`
	Slack_username     string `json:"slack_username"`
	Pagerduty_username string `json:"pagerduty_username"`
	Public_gpg_key     string `json:"public_gpg_key"`
}

// GetPath returns UsersUsers_v1User_v1.Path, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetPath() string { return v.Path }

// GetName returns UsersUsers_v1User_v1.Name, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetName() string { return v.Name }

// GetOrg_username returns UsersUsers_v1User_v1.Org_username, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetOrg_username() string { return v.Org_username }

// GetGithub_username returns UsersUsers_v1User_v1.Github_username, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetGithub_username() string { return v.Github_username }

// GetSlack_username returns UsersUsers_v1User_v1.Slack_username, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetSlack_username() string { return v.Slack_username }

// GetPagerduty_username returns UsersUsers_v1User_v1.Pagerduty_username, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetPagerduty_username() string { return v.Pagerduty_username }

// GetPublic_gpg_key returns UsersUsers_v1User_v1.Public_gpg_key, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetPublic_gpg_key() string { return v.Public_gpg_key }

func GithubOrgs(
	ctx context.Context,
) (*GithubOrgsResponse, error) {
	req := &graphql.Request{
		OpName: "GithubOrgs",
		Query: `
query GithubOrgs {
	githuborg_v1 {
		name
		token {
			path
			field
			version
			format
		}
		default
	}
}
`,
	}
	var err error
	var client graphql.Client

	client, err = gql.NewQontractClient(ctx)
	if err != nil {
		return nil, err
	}

	var data GithubOrgsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func Users(
	ctx context.Context,
) (*UsersResponse, error) {
	req := &graphql.Request{
		OpName: "Users",
		Query: `
query Users {
	users_v1 {
		path
		name
		org_username
		github_username
		slack_username
		pagerduty_username
		public_gpg_key
	}
}
`,
	}
	var err error
	var client graphql.Client

	client, err = gql.NewQontractClient(ctx)
	if err != nil {
		return nil, err
	}

	var data UsersResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}