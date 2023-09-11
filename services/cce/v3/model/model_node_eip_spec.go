package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NodeEipSpec struct {
	Iptype string `json:"iptype"`

	Bandwidth *NodeBandwidth `json:"bandwidth,omitempty"`
}

func (o NodeEipSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeEipSpec struct{}"
	}

	return strings.Join([]string{"NodeEipSpec", string(data)}, " ")
}
