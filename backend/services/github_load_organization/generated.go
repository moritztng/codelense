// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// __getOrganizationInput is used internally by genqlient
type __getOrganizationInput struct {
	Login string `json:"login"`
}

// GetLogin returns __getOrganizationInput.Login, and is useful for accessing the field via an interface.
func (v *__getOrganizationInput) GetLogin() string { return v.Login }

// getOrganizationOrganization includes the requested fields of the GraphQL type Organization.
// The GraphQL type's documentation follows.
//
// An account on GitHub, with one or more owners, that has repositories, members and teams.
type getOrganizationOrganization struct {
	// Identifies the primary key from the database.
	DatabaseId int `json:"databaseId"`
	// The organization's login name.
	Login string `json:"login"`
	// The organization's public email.
	Email string `json:"email"`
	// The organization's public profile name.
	Name string `json:"name"`
	// The organization's public profile location.
	Location string `json:"location"`
	// Identifies the date and time when the object was created.
	CreatedAt time.Time `json:"createdAt"`
	// A URL pointing to the organization's public avatar.
	AvatarUrl string `json:"avatarUrl"`
	// The organization's public profile description.
	Description string `json:"description"`
	// The organization's Twitter username.
	TwitterUsername string `json:"twitterUsername"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// The organization's public profile URL.
	WebsiteUrl string `json:"websiteUrl"`
	// The HTTP URL for this organization.
	Url string `json:"url"`
}

// GetDatabaseId returns getOrganizationOrganization.DatabaseId, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetDatabaseId() int { return v.DatabaseId }

// GetLogin returns getOrganizationOrganization.Login, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetLogin() string { return v.Login }

// GetEmail returns getOrganizationOrganization.Email, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetEmail() string { return v.Email }

// GetName returns getOrganizationOrganization.Name, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetName() string { return v.Name }

// GetLocation returns getOrganizationOrganization.Location, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetLocation() string { return v.Location }

// GetCreatedAt returns getOrganizationOrganization.CreatedAt, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetCreatedAt() time.Time { return v.CreatedAt }

// GetAvatarUrl returns getOrganizationOrganization.AvatarUrl, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetAvatarUrl() string { return v.AvatarUrl }

// GetDescription returns getOrganizationOrganization.Description, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetDescription() string { return v.Description }

// GetTwitterUsername returns getOrganizationOrganization.TwitterUsername, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetTwitterUsername() string { return v.TwitterUsername }

// GetUpdatedAt returns getOrganizationOrganization.UpdatedAt, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetUpdatedAt() time.Time { return v.UpdatedAt }

// GetWebsiteUrl returns getOrganizationOrganization.WebsiteUrl, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetWebsiteUrl() string { return v.WebsiteUrl }

// GetUrl returns getOrganizationOrganization.Url, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetUrl() string { return v.Url }

// getOrganizationResponse is returned by getOrganization on success.
type getOrganizationResponse struct {
	// Lookup a organization by login.
	Organization getOrganizationOrganization `json:"organization"`
}

// GetOrganization returns getOrganizationResponse.Organization, and is useful for accessing the field via an interface.
func (v *getOrganizationResponse) GetOrganization() getOrganizationOrganization {
	return v.Organization
}

// The query or mutation executed by getOrganization.
const getOrganization_Operation = `
query getOrganization ($login: String!) {
	organization(login: $login) {
		databaseId
		login
		email
		name
		location
		createdAt
		avatarUrl
		description
		twitterUsername
		updatedAt
		websiteUrl
		url
	}
}
`

func getOrganization(
	ctx context.Context,
	client graphql.Client,
	login string,
) (*getOrganizationResponse, error) {
	req := &graphql.Request{
		OpName: "getOrganization",
		Query:  getOrganization_Operation,
		Variables: &__getOrganizationInput{
			Login: login,
		},
	}
	var err error

	var data getOrganizationResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
