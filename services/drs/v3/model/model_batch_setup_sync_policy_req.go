package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchSetupSyncPolicyReq struct {
	Jobs []SyncPolicyReq `json:"jobs"`
}

func (o BatchSetupSyncPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetupSyncPolicyReq struct{}"
	}

	return strings.Join([]string{"BatchSetupSyncPolicyReq", string(data)}, " ")
}
