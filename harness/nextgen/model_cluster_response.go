/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This is the ClusterRequest entity defined in Harness
type ClusterResponse struct {
	// identifier of the gitops cluster
	ClusterRef string `json:"clusterRef,omitempty"`
	// organization identifier of the cluster
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// project identifier of the cluster
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// account identifier of the cluster
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	// environment identifier of the cluster
	EnvRef string `json:"envRef"`
	// time at which the cluster was linked
	LinkedAt int64 `json:"linkedAt,omitempty"`
	// scope at which the cluster exists in harness gitops, project vs org vs account
	Scope string `json:"scope,omitempty"`
	// name of the gitops cluster
	Name string `json:"name,omitempty"`
}
