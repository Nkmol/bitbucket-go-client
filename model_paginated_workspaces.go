/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bitbucket

// A paginated list of workspaces.
type PaginatedWorkspaces struct {
	// Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Size int32 `json:"size,omitempty"`
	// Page number of the current results. This is an optional element that is not provided in all responses.
	Page int32 `json:"page,omitempty"`
	// Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Pagelen int32 `json:"pagelen,omitempty"`
	// Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Next string `json:"next,omitempty"`
	// Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Previous string      `json:"previous,omitempty"`
	Values   []Workspace `json:"values,omitempty"`
}
