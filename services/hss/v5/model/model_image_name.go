package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ImageName struct {
}

func (o ImageName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageName struct{}"
	}

	return strings.Join([]string{"ImageName", string(data)}, " ")
}
