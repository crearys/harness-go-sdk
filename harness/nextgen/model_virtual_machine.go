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

type VirtualMachine struct {
	AvgPrice             float64  `json:"avgPrice,omitempty"`
	Burst                bool     `json:"burst,omitempty"`
	Category             string   `json:"category,omitempty"`
	CpusPerVm            float64  `json:"cpusPerVm,omitempty"`
	CurrentGen           bool     `json:"currentGen,omitempty"`
	GpusPerVm            float64  `json:"gpusPerVm,omitempty"`
	MemPerVm             float64  `json:"memPerVm,omitempty"`
	AllocatableCpusPerVm float64  `json:"allocatableCpusPerVm,omitempty"`
	AllocatableMemPerVm  float64  `json:"allocatableMemPerVm,omitempty"`
	NetworkPerf          string   `json:"networkPerf,omitempty"`
	NetworkPerfCategory  string   `json:"networkPerfCategory,omitempty"`
	OnDemandPrice        float64  `json:"onDemandPrice,omitempty"`
	Type_                string   `json:"type,omitempty"`
	Zones                []string `json:"zones,omitempty"`
}
