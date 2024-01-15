package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type JarPackageHostInfo struct {
	AgentId *string `json:"agent_id,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	FileName *string `json:"file_name,omitempty"`

	Name *string `json:"name,omitempty"`

	Catalogue *string `json:"catalogue,omitempty"`

	FileType *string `json:"file_type,omitempty"`

	Version *string `json:"version,omitempty"`

	Path *string `json:"path,omitempty"`

	Hash *string `json:"hash,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Uid *int32 `json:"uid,omitempty"`

	Gid *int32 `json:"gid,omitempty"`

	Mode *string `json:"mode,omitempty"`

	Pid *int32 `json:"pid,omitempty"`

	ProcPath *string `json:"proc_path,omitempty"`

	ContainerId *string `json:"container_id,omitempty"`

	ContainerName *string `json:"container_name,omitempty"`

	PackagePath *string `json:"package_path,omitempty"`

	IsEmbedded *int32 `json:"is_embedded,omitempty"`

	RecordTime *int64 `json:"record_time,omitempty"`
}

func (o JarPackageHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JarPackageHostInfo struct{}"
	}

	return strings.Join([]string{"JarPackageHostInfo", string(data)}, " ")
}
