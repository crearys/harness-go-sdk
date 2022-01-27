/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Returns config details for the AWS KMS Secret Manager Connector.
type AwsKmsConnectorCredential struct {
	Type_ string                `json:"type"`
	Spec  *AwsKmsCredentialSpec `json:"spec,omitempty"`
}