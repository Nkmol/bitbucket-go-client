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

type Participant struct {
	Type_    string `json:"type"`
	User     *User  `json:"user,omitempty"`
	Role     string `json:"role,omitempty"`
	Approved bool   `json:"approved,omitempty"`
	State    string `json:"state,omitempty"`
	// The ISO8601 timestamp of the participant's action. For approvers, this is the time of their approval. For commenters and pull request reviewers who are not approvers, this is the time they last commented, or null if they have not commented.
	ParticipatedOn time.Time `json:"participated_on,omitempty"`
}
