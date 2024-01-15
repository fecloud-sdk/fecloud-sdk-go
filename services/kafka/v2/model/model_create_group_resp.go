package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateGroupResp struct {
}

func (o CreateGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupResp struct{}"
	}

	return strings.Join([]string{"CreateGroupResp", string(data)}, " ")
}
