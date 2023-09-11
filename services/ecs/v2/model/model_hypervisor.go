package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Hypervisor hypervisor信息。
type Hypervisor struct {

	// hypervisor类型
	HypervisorType *string `json:"hypervisor_type,omitempty"`

	// hypervisor csd信息
	CsdHypervisor *string `json:"csd_hypervisor,omitempty"`
}

func (o Hypervisor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Hypervisor struct{}"
	}

	return strings.Join([]string{"Hypervisor", string(data)}, " ")
}
