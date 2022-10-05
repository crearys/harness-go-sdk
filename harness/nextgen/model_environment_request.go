/*
* Harness NextGen Software Delivery Platform API Reference
*
* This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
*
* API version: 3.0
* Contact: contact@harness.io
* Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// This is the Environment entity defined in Harness
type EnvironmentRequest struct {
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Identifier of the Environment.
	Identifier string `json:"identifier,omitempty"`
	// Tags
	Tags map[string]string `json:"tags,omitempty"`
	// Name of the Environment.
	Name string `json:"name,omitempty"`
	// Description of the entity
	Description string `json:"description,omitempty"`
	// Color of the Environment.
	Color string `json:"color,omitempty"`
	// Specify the environment type whether production or Preproduction.
	Type_ string `json:"type,omitempty"`
	// Yaml of this entity.
	Yaml string `json:"yaml,omitempty"`
}
