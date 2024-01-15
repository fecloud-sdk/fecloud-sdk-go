package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DedicatedResourceCapacity struct {
	Vcpus int32 `json:"vcpus"`

	Ram int32 `json:"ram"`

	Volume int32 `json:"volume"`
}

func (o DedicatedResourceCapacity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DedicatedResourceCapacity struct{}"
	}

	return strings.Join([]string{"DedicatedResourceCapacity", string(data)}, " ")
}
