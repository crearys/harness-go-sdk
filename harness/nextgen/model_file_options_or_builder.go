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

type FileOptionsOrBuilder struct {
	JavaStringCheckUtf8              bool                           `json:"javaStringCheckUtf8,omitempty"`
	JavaPackage                      string                         `json:"javaPackage,omitempty"`
	JavaPackageBytes                 *ByteString                    `json:"javaPackageBytes,omitempty"`
	JavaOuterClassname               string                         `json:"javaOuterClassname,omitempty"`
	JavaOuterClassnameBytes          *ByteString                    `json:"javaOuterClassnameBytes,omitempty"`
	JavaMultipleFiles                bool                           `json:"javaMultipleFiles,omitempty"`
	JavaGenerateEqualsAndHash        bool                           `json:"javaGenerateEqualsAndHash,omitempty"`
	OptimizeFor                      string                         `json:"optimizeFor,omitempty"`
	GoPackage                        string                         `json:"goPackage,omitempty"`
	GoPackageBytes                   *ByteString                    `json:"goPackageBytes,omitempty"`
	CcGenericServices                bool                           `json:"ccGenericServices,omitempty"`
	JavaGenericServices              bool                           `json:"javaGenericServices,omitempty"`
	PyGenericServices                bool                           `json:"pyGenericServices,omitempty"`
	PhpGenericServices               bool                           `json:"phpGenericServices,omitempty"`
	Deprecated                       bool                           `json:"deprecated,omitempty"`
	CcEnableArenas                   bool                           `json:"ccEnableArenas,omitempty"`
	ObjcClassPrefix                  string                         `json:"objcClassPrefix,omitempty"`
	ObjcClassPrefixBytes             *ByteString                    `json:"objcClassPrefixBytes,omitempty"`
	CsharpNamespace                  string                         `json:"csharpNamespace,omitempty"`
	CsharpNamespaceBytes             *ByteString                    `json:"csharpNamespaceBytes,omitempty"`
	SwiftPrefix                      string                         `json:"swiftPrefix,omitempty"`
	SwiftPrefixBytes                 *ByteString                    `json:"swiftPrefixBytes,omitempty"`
	PhpClassPrefix                   string                         `json:"phpClassPrefix,omitempty"`
	PhpClassPrefixBytes              *ByteString                    `json:"phpClassPrefixBytes,omitempty"`
	PhpNamespace                     string                         `json:"phpNamespace,omitempty"`
	PhpNamespaceBytes                *ByteString                    `json:"phpNamespaceBytes,omitempty"`
	PhpMetadataNamespace             string                         `json:"phpMetadataNamespace,omitempty"`
	PhpMetadataNamespaceBytes        *ByteString                    `json:"phpMetadataNamespaceBytes,omitempty"`
	RubyPackage                      string                         `json:"rubyPackage,omitempty"`
	RubyPackageBytes                 *ByteString                    `json:"rubyPackageBytes,omitempty"`
	UninterpretedOptionList          []UninterpretedOption          `json:"uninterpretedOptionList,omitempty"`
	UninterpretedOptionCount         int32                          `json:"uninterpretedOptionCount,omitempty"`
	UninterpretedOptionOrBuilderList []UninterpretedOptionOrBuilder `json:"uninterpretedOptionOrBuilderList,omitempty"`
	DefaultInstanceForType           *Message                       `json:"defaultInstanceForType,omitempty"`
	AllFields                        map[string]interface{}         `json:"allFields,omitempty"`
	InitializationErrorString        string                         `json:"initializationErrorString,omitempty"`
	DescriptorForType                *Descriptor                    `json:"descriptorForType,omitempty"`
	UnknownFields                    *UnknownFieldSet               `json:"unknownFields,omitempty"`
	Initialized                      bool                           `json:"initialized,omitempty"`
}
