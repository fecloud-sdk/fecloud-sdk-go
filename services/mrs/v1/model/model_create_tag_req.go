package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateTagReq struct {
	Tag *Tag `json:"tag"`
}

func (o CreateTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagReq struct{}"
	}

	return strings.Join([]string{"CreateTagReq", string(data)}, " ")
}
