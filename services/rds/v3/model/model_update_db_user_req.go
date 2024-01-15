package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateDbUserReq struct {
	Comment *string `json:"comment,omitempty"`
}

func (o UpdateDbUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbUserReq struct{}"
	}

	return strings.Join([]string{"UpdateDbUserReq", string(data)}, " ")
}
