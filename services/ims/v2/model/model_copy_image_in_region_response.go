package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CopyImageInRegionResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyImageInRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageInRegionResponse struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionResponse", string(data)}, " ")
}
