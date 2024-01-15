package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventProcessResponseInfo struct {
	ProcessName *string `json:"process_name,omitempty"`

	ProcessPath *string `json:"process_path,omitempty"`

	ProcessPid *int32 `json:"process_pid,omitempty"`

	ProcessUid *int32 `json:"process_uid,omitempty"`

	ProcessUsername *string `json:"process_username,omitempty"`

	ProcessCmdline *string `json:"process_cmdline,omitempty"`

	ProcessFilename *string `json:"process_filename,omitempty"`

	ProcessStartTime *int64 `json:"process_start_time,omitempty"`

	ProcessGid *int32 `json:"process_gid,omitempty"`

	ProcessEgid *int32 `json:"process_egid,omitempty"`

	ProcessEuid *int32 `json:"process_euid,omitempty"`

	ParentProcessName *string `json:"parent_process_name,omitempty"`

	ParentProcessPath *string `json:"parent_process_path,omitempty"`

	ParentProcessPid *int32 `json:"parent_process_pid,omitempty"`

	ParentProcessUid *int32 `json:"parent_process_uid,omitempty"`

	ParentProcessCmdline *string `json:"parent_process_cmdline,omitempty"`

	ParentProcessFilename *string `json:"parent_process_filename,omitempty"`

	ParentProcessStartTime *int64 `json:"parent_process_start_time,omitempty"`

	ParentProcessGid *int32 `json:"parent_process_gid,omitempty"`

	ParentProcessEgid *int32 `json:"parent_process_egid,omitempty"`

	ParentProcessEuid *int32 `json:"parent_process_euid,omitempty"`

	ChildProcessName *string `json:"child_process_name,omitempty"`

	ChildProcessPath *string `json:"child_process_path,omitempty"`

	ChildProcessPid *int32 `json:"child_process_pid,omitempty"`

	ChildProcessUid *int32 `json:"child_process_uid,omitempty"`

	ChildProcessCmdline *string `json:"child_process_cmdline,omitempty"`

	ChildProcessFilename *string `json:"child_process_filename,omitempty"`

	ChildProcessStartTime *int64 `json:"child_process_start_time,omitempty"`

	ChildProcessGid *int32 `json:"child_process_gid,omitempty"`

	ChildProcessEgid *int32 `json:"child_process_egid,omitempty"`

	ChildProcessEuid *int32 `json:"child_process_euid,omitempty"`

	VirtCmd *string `json:"virt_cmd,omitempty"`

	VirtProcessName *string `json:"virt_process_name,omitempty"`

	EscapeMode *string `json:"escape_mode,omitempty"`

	EscapeCmd *string `json:"escape_cmd,omitempty"`

	ProcessHash *string `json:"process_hash,omitempty"`
}

func (o EventProcessResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventProcessResponseInfo struct{}"
	}

	return strings.Join([]string{"EventProcessResponseInfo", string(data)}, " ")
}
