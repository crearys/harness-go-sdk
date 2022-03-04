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

type ChildChainExecutableResponse struct {
	UnknownFields             *UnknownFieldSet                    `json:"unknownFields,omitempty"`
	Initialized               bool                                `json:"initialized,omitempty"`
	NextChildId               string                              `json:"nextChildId,omitempty"`
	NextChildIdBytes          *ByteString                         `json:"nextChildIdBytes,omitempty"`
	PreviousChildId           string                              `json:"previousChildId,omitempty"`
	PreviousChildIdBytes      *ByteString                         `json:"previousChildIdBytes,omitempty"`
	SerializedSize            int32                               `json:"serializedSize,omitempty"`
	ParserForType             *ParserChildChainExecutableResponse `json:"parserForType,omitempty"`
	DefaultInstanceForType    *ChildChainExecutableResponse       `json:"defaultInstanceForType,omitempty"`
	PassThroughData           *ByteString                         `json:"passThroughData,omitempty"`
	LastLink                  bool                                `json:"lastLink,omitempty"`
	Suspend                   bool                                `json:"suspend,omitempty"`
	AllFields                 map[string]interface{}              `json:"allFields,omitempty"`
	InitializationErrorString string                              `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor                         `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize    int32                               `json:"memoizedSerializedSize,omitempty"`
}
