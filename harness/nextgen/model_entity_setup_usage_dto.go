/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type EntitySetupUsageDto struct {
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	ReferredEntity *EntityDetail `json:"referredEntity,omitempty"`
	ReferredByEntity *EntityDetail `json:"referredByEntity"`
	Detail *SetupUsageDetail `json:"detail,omitempty"`
	CreatedAt int64 `json:"createdAt,omitempty"`
}