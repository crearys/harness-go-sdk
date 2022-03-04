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

// Returns AWS Secret Manager configuration details.
type AwsSecretManager struct {
	Credential *AwsSecretManagerCredential `json:"credential"`
	// Region for AWS Secret Manager.
	Region string `json:"region"`
	// Text that is appended to the Secret as a prefix.
	SecretNamePrefix string `json:"secretNamePrefix,omitempty"`
	// List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.
	DelegateSelectors []string `json:"delegateSelectors,omitempty"`
	Default_          bool     `json:"default,omitempty"`
}
