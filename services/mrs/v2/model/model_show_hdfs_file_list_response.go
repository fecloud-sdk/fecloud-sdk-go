package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowHdfsFileListResponse struct {
	TotalCount *int64 `json:"total_count,omitempty"`

	Files          *[]FileStatusV2 `json:"files,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowHdfsFileListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHdfsFileListResponse struct{}"
	}

	return strings.Join([]string{"ShowHdfsFileListResponse", string(data)}, " ")
}
