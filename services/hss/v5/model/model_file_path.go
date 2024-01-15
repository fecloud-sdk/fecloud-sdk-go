package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FilePath struct {
}

func (o FilePath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilePath struct{}"
	}

	return strings.Join([]string{"FilePath", string(data)}, " ")
}
