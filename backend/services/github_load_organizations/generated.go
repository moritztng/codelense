// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// __getOrganizationsInput is used internally by genqlient
type __getOrganizationsInput struct {
	Cursor string `json:"cursor,omitempty"`
	Query  string `json:"query"`
}

// GetCursor returns __getOrganizationsInput.Cursor, and is useful for accessing the field via an interface.
func (v *__getOrganizationsInput) GetCursor() string { return v.Cursor }

// GetQuery returns __getOrganizationsInput.Query, and is useful for accessing the field via an interface.
func (v *__getOrganizationsInput) GetQuery() string { return v.Query }

// getOrganizationsResponse is returned by getOrganizations on success.
type getOrganizationsResponse struct {
	// Perform a search across resources, returning a maximum of 1,000 results.
	Search getOrganizationsSearchSearchResultItemConnection `json:"search"`
}

// GetSearch returns getOrganizationsResponse.Search, and is useful for accessing the field via an interface.
func (v *getOrganizationsResponse) GetSearch() getOrganizationsSearchSearchResultItemConnection {
	return v.Search
}

// getOrganizationsSearchSearchResultItemConnection includes the requested fields of the GraphQL type SearchResultItemConnection.
// The GraphQL type's documentation follows.
//
// A list of results that matched against a search query. Regardless of the number
// of matches, a maximum of 1,000 results will be available across all types,
// potentially split across many pages.
type getOrganizationsSearchSearchResultItemConnection struct {
	// The total number of users that matched the search query. Regardless of the
	// total number of matches, a maximum of 1,000 results will be available across all types.
	UserCount int `json:"userCount"`
	// A list of edges.
	Edges []getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge `json:"edges"`
	// Information to aid in pagination.
	PageInfo getOrganizationsSearchSearchResultItemConnectionPageInfo `json:"pageInfo"`
}

// GetUserCount returns getOrganizationsSearchSearchResultItemConnection.UserCount, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnection) GetUserCount() int { return v.UserCount }

// GetEdges returns getOrganizationsSearchSearchResultItemConnection.Edges, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnection) GetEdges() []getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge {
	return v.Edges
}

// GetPageInfo returns getOrganizationsSearchSearchResultItemConnection.PageInfo, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnection) GetPageInfo() getOrganizationsSearchSearchResultItemConnectionPageInfo {
	return v.PageInfo
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge includes the requested fields of the GraphQL type SearchResultItemEdge.
// The GraphQL type's documentation follows.
//
// An edge in a connection.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge struct {
	// The item at the end of the edge.
	Node getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem `json:"-"`
}

// GetNode returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge.Node, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge) GetNode() getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem {
	return v.Node
}

func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge
		Node json.RawMessage `json:"node"`
		graphql.NoUnmarshalJSON
	}
	firstPass.getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Node
		src := firstPass.Node
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalgetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge.Node: %w", err)
			}
		}
	}
	return nil
}

type __premarshalgetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge struct {
	Node json.RawMessage `json:"node"`
}

func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge) __premarshalJSON() (*__premarshalgetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge, error) {
	var retval __premarshalgetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge

	{

		dst := &retval.Node
		src := v.Node
		var err error
		*dst, err = __marshalgetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdge.Node: %w", err)
		}
	}
	return &retval, nil
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeApp includes the requested fields of the GraphQL type App.
// The GraphQL type's documentation follows.
//
// A GitHub App.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeApp struct {
	Typename string `json:"__typename"`
}

// GetTypename returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeApp.Typename, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeApp) GetTypename() string {
	return v.Typename
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeDiscussion includes the requested fields of the GraphQL type Discussion.
// The GraphQL type's documentation follows.
//
// A discussion in a repository.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeDiscussion struct {
	Typename string `json:"__typename"`
}

// GetTypename returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeDiscussion.Typename, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeDiscussion) GetTypename() string {
	return v.Typename
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeIssue includes the requested fields of the GraphQL type Issue.
// The GraphQL type's documentation follows.
//
// An Issue is a place to discuss ideas, enhancements, tasks, and bugs for a project.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeIssue struct {
	Typename string `json:"__typename"`
}

// GetTypename returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeIssue.Typename, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeIssue) GetTypename() string {
	return v.Typename
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeMarketplaceListing includes the requested fields of the GraphQL type MarketplaceListing.
// The GraphQL type's documentation follows.
//
// A listing in the GitHub integration marketplace.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeMarketplaceListing struct {
	Typename string `json:"__typename"`
}

// GetTypename returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeMarketplaceListing.Typename, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeMarketplaceListing) GetTypename() string {
	return v.Typename
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization includes the requested fields of the GraphQL type Organization.
// The GraphQL type's documentation follows.
//
// An account on GitHub, with one or more owners, that has repositories, members and teams.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization struct {
	Typename string `json:"__typename"`
	// Identifies the primary key from the database.
	DatabaseId int `json:"databaseId"`
	// The organization's login name.
	Login string `json:"login"`
	// Identifies the date and time when the object was created.
	CreatedAt time.Time `json:"createdAt"`
}

// GetTypename returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization.Typename, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization) GetTypename() string {
	return v.Typename
}

// GetDatabaseId returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization.DatabaseId, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization) GetDatabaseId() int {
	return v.DatabaseId
}

// GetLogin returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization.Login, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization) GetLogin() string {
	return v.Login
}

// GetCreatedAt returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization.CreatedAt, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization) GetCreatedAt() time.Time {
	return v.CreatedAt
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodePullRequest includes the requested fields of the GraphQL type PullRequest.
// The GraphQL type's documentation follows.
//
// A repository pull request.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodePullRequest struct {
	Typename string `json:"__typename"`
}

// GetTypename returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodePullRequest.Typename, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodePullRequest) GetTypename() string {
	return v.Typename
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository includes the requested fields of the GraphQL type Repository.
// The GraphQL type's documentation follows.
//
// A repository contains the content for a project.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository struct {
	Typename string `json:"__typename"`
}

// GetTypename returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository.Typename, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository) GetTypename() string {
	return v.Typename
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem includes the requested fields of the GraphQL interface SearchResultItem.
//
// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem is implemented by the following types:
// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeApp
// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeDiscussion
// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeIssue
// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeMarketplaceListing
// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization
// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodePullRequest
// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository
// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeUser
// The GraphQL type's documentation follows.
//
// The results of a search.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem interface {
	implementsGraphQLInterfacegetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
}

func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeApp) implementsGraphQLInterfacegetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem() {
}
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeDiscussion) implementsGraphQLInterfacegetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem() {
}
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeIssue) implementsGraphQLInterfacegetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem() {
}
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeMarketplaceListing) implementsGraphQLInterfacegetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem() {
}
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization) implementsGraphQLInterfacegetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem() {
}
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodePullRequest) implementsGraphQLInterfacegetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem() {
}
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository) implementsGraphQLInterfacegetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem() {
}
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeUser) implementsGraphQLInterfacegetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem() {
}

func __unmarshalgetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem(b []byte, v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "App":
		*v = new(getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeApp)
		return json.Unmarshal(b, *v)
	case "Discussion":
		*v = new(getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeDiscussion)
		return json.Unmarshal(b, *v)
	case "Issue":
		*v = new(getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeIssue)
		return json.Unmarshal(b, *v)
	case "MarketplaceListing":
		*v = new(getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeMarketplaceListing)
		return json.Unmarshal(b, *v)
	case "Organization":
		*v = new(getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization)
		return json.Unmarshal(b, *v)
	case "PullRequest":
		*v = new(getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodePullRequest)
		return json.Unmarshal(b, *v)
	case "Repository":
		*v = new(getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository)
		return json.Unmarshal(b, *v)
	case "User":
		*v = new(getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeUser)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing SearchResultItem.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem: "%v"`, tn.TypeName)
	}
}

func __marshalgetOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem(v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeApp:
		typename = "App"

		result := struct {
			TypeName string `json:"__typename"`
			*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeApp
		}{typename, v}
		return json.Marshal(result)
	case *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeDiscussion:
		typename = "Discussion"

		result := struct {
			TypeName string `json:"__typename"`
			*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeDiscussion
		}{typename, v}
		return json.Marshal(result)
	case *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeIssue:
		typename = "Issue"

		result := struct {
			TypeName string `json:"__typename"`
			*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeIssue
		}{typename, v}
		return json.Marshal(result)
	case *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeMarketplaceListing:
		typename = "MarketplaceListing"

		result := struct {
			TypeName string `json:"__typename"`
			*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeMarketplaceListing
		}{typename, v}
		return json.Marshal(result)
	case *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization:
		typename = "Organization"

		result := struct {
			TypeName string `json:"__typename"`
			*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeOrganization
		}{typename, v}
		return json.Marshal(result)
	case *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodePullRequest:
		typename = "PullRequest"

		result := struct {
			TypeName string `json:"__typename"`
			*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodePullRequest
		}{typename, v}
		return json.Marshal(result)
	case *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository:
		typename = "Repository"

		result := struct {
			TypeName string `json:"__typename"`
			*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeRepository
		}{typename, v}
		return json.Marshal(result)
	case *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeUser:
		typename = "User"

		result := struct {
			TypeName string `json:"__typename"`
			*getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeUser
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeSearchResultItem: "%T"`, v)
	}
}

// getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A user is an individual's account on GitHub that owns repositories and can make new content.
type getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeUser struct {
	Typename string `json:"__typename"`
}

// GetTypename returns getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeUser.Typename, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionEdgesSearchResultItemEdgeNodeUser) GetTypename() string {
	return v.Typename
}

// getOrganizationsSearchSearchResultItemConnectionPageInfo includes the requested fields of the GraphQL type PageInfo.
// The GraphQL type's documentation follows.
//
// Information about pagination in a connection.
type getOrganizationsSearchSearchResultItemConnectionPageInfo struct {
	// When paginating forwards, the cursor to continue.
	EndCursor string `json:"endCursor"`
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
}

// GetEndCursor returns getOrganizationsSearchSearchResultItemConnectionPageInfo.EndCursor, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionPageInfo) GetEndCursor() string {
	return v.EndCursor
}

// GetHasNextPage returns getOrganizationsSearchSearchResultItemConnectionPageInfo.HasNextPage, and is useful for accessing the field via an interface.
func (v *getOrganizationsSearchSearchResultItemConnectionPageInfo) GetHasNextPage() bool {
	return v.HasNextPage
}

// The query or mutation executed by getOrganizations.
const getOrganizations_Operation = `
query getOrganizations ($cursor: String, $query: String!) {
	search(query: $query, type: USER, first: 100, after: $cursor) {
		userCount
		edges {
			node {
				__typename
				... on Organization {
					databaseId
					login
					createdAt
				}
			}
		}
		pageInfo {
			endCursor
			hasNextPage
		}
	}
}
`

func getOrganizations(
	ctx context.Context,
	client graphql.Client,
	cursor string,
	query string,
) (*getOrganizationsResponse, error) {
	req := &graphql.Request{
		OpName: "getOrganizations",
		Query:  getOrganizations_Operation,
		Variables: &__getOrganizationsInput{
			Cursor: cursor,
			Query:  query,
		},
	}
	var err error

	var data getOrganizationsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
