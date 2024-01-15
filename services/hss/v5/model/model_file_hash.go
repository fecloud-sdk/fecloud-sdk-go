package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FileHash struct {
}

func (o FileHash) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileHash struct{}"
	}

	return strings.Join([]string{"FileHash", string(data)}, " ")
}
