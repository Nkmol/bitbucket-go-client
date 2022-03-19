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

type Group struct {
	Type_     string      `json:"type"`
	Links     *GroupLinks `json:"links,omitempty"`
	Owner     *Account    `json:"owner,omitempty"`
	Workspace *Workspace  `json:"workspace,omitempty"`
	Name      string      `json:"name,omitempty"`
	// The \"sluggified\" version of the group's name. This contains only ASCII characters and can therefore be slightly different than the name
	Slug string `json:"slug,omitempty"`
	// The concatenation of the workspace's slug and the group's slug, separated with a colon (e.g. `acme:developers`)
	FullSlug string `json:"full_slug,omitempty"`
}
