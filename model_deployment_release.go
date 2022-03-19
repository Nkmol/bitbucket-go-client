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

type DeploymentRelease struct {
	Type_ string `json:"type"`
	// The UUID identifying the release.
	Uuid string `json:"uuid,omitempty"`
	// The name of the release.
	Name string `json:"name,omitempty"`
	// Link to the pipeline that produced the release.
	Url    string  `json:"url,omitempty"`
	Commit *Commit `json:"commit,omitempty"`
	// The timestamp when the release was created.
	CreatedOn time.Time `json:"created_on,omitempty"`
}
