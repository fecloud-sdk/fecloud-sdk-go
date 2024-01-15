package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListServerBlockDevicesResponse struct {
	AttachableQuantity *BlockDeviceAttachableQuantity `json:"attachableQuantity,omitempty"`

	VolumeAttachments *[]ServerBlockDevice `json:"volumeAttachments,omitempty"`
	HttpStatusCode    int                  `json:"-"`
}

func (o ListServerBlockDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerBlockDevicesResponse struct{}"
	}

	return strings.Join([]string{"ListServerBlockDevicesResponse", string(data)}, " ")
}
