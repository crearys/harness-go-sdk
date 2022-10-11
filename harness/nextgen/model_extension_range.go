/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type ExtensionRange struct {
	UnknownFields             *UnknownFieldSet                `json:"unknownFields,omitempty"`
	Initialized               bool                            `json:"initialized,omitempty"`
	Options                   *ExtensionRangeOptions          `json:"options,omitempty"`
	End                       int32                           `json:"end,omitempty"`
	Start                     int32                           `json:"start,omitempty"`
	ParserForType             *ParserExtensionRange           `json:"parserForType,omitempty"`
	SerializedSize            int32                           `json:"serializedSize,omitempty"`
	DefaultInstanceForType    *ExtensionRange                 `json:"defaultInstanceForType,omitempty"`
	OptionsOrBuilder          *ExtensionRangeOptionsOrBuilder `json:"optionsOrBuilder,omitempty"`
	AllFields                 map[string]interface{}          `json:"allFields,omitempty"`
	DescriptorForType         *Descriptor                     `json:"descriptorForType,omitempty"`
	InitializationErrorString string                          `json:"initializationErrorString,omitempty"`
	MemoizedSerializedSize    int32                           `json:"memoizedSerializedSize,omitempty"`
}
