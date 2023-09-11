package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteServersRequest Request Object
type DeleteServersRequest struct {
	Body *DeleteServersRequestBody `json:"body,omitempty"`
}

func (o DeleteServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServersRequest struct{}"
	}

	return strings.Join([]string{"DeleteServersRequest", string(data)}, " ")
}
