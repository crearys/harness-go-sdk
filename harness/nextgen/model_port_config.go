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

type PortConfig struct {
	Protocol       string        `json:"protocol,omitempty"`
	TargetProtocol string        `json:"target_protocol,omitempty"`
	Port           int32         `json:"port,omitempty"`
	TargetPort     int32         `json:"target_port,omitempty"`
	ServerName     string        `json:"server_name,omitempty"`
	Action         string        `json:"action,omitempty"`
	RedirectUrl    string        `json:"redirect_url,omitempty"`
	RoutingRules   []RoutingRule `json:"routing_rules,omitempty"`
}
