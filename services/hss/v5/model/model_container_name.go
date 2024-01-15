package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ContainerName struct {
}

func (o ContainerName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerName struct{}"
	}

	return strings.Join([]string{"ContainerName", string(data)}, " ")
}
