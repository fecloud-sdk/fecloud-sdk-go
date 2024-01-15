package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FileAttr struct {
}

func (o FileAttr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileAttr struct{}"
	}

	return strings.Join([]string{"FileAttr", string(data)}, " ")
}
