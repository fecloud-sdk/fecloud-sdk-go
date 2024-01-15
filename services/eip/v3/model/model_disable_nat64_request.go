package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DisableNat64Request Request Object
type DisableNat64Request struct {

	// 弹性公网ID
	PublicipId string `json:"publicip_id"`
}

func (o DisableNat64Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableNat64Request struct{}"
	}

	return strings.Join([]string{"DisableNat64Request", string(data)}, " ")
}
