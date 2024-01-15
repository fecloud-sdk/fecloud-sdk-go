package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PostPaidServerEip struct {
	Iptype string `json:"iptype"`

	Bandwidth *PostPaidServerEipBandwidth `json:"bandwidth"`

	Extendparam *PostPaidServerEipExtendParam `json:"extendparam,omitempty"`
}

func (o PostPaidServerEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerEip struct{}"
	}

	return strings.Join([]string{"PostPaidServerEip", string(data)}, " ")
}
