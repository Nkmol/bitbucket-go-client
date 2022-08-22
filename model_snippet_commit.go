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

import (
	"time"
)

type SnippetCommit struct {
	Hash    string              `json:"hash,omitempty"`
	Date    time.Time           `json:"date,omitempty"`
	Author  *Author             `json:"author,omitempty"`
	Message string              `json:"message,omitempty"`
	Summary *CommentContent     `json:"summary,omitempty"`
	Parents []BaseCommit        `json:"parents,omitempty"`
	Links   *SnippetCommitLinks `json:"links,omitempty"`
	Snippet *Snippet            `json:"snippet,omitempty"`
}
