package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventFileResponseInfo struct {
	FilePath *string `json:"file_path,omitempty"`

	FileAlias *string `json:"file_alias,omitempty"`

	FileSize *int32 `json:"file_size,omitempty"`

	FileMtime *int64 `json:"file_mtime,omitempty"`

	FileAtime *int64 `json:"file_atime,omitempty"`

	FileCtime *int64 `json:"file_ctime,omitempty"`

	FileHash *string `json:"file_hash,omitempty"`

	FileMd5 *string `json:"file_md5,omitempty"`

	FileSha256 *string `json:"file_sha256,omitempty"`

	FileType *string `json:"file_type,omitempty"`

	FileContent *string `json:"file_content,omitempty"`

	FileAttr *string `json:"file_attr,omitempty"`

	FileOperation *int32 `json:"file_operation,omitempty"`

	FileAction *string `json:"file_action,omitempty"`

	FileChangeAttr *string `json:"file_change_attr,omitempty"`

	FileNewPath *string `json:"file_new_path,omitempty"`

	FileDesc *string `json:"file_desc,omitempty"`

	FileKeyWord *string `json:"file_key_word,omitempty"`

	IsDir *bool `json:"is_dir,omitempty"`

	FdInfo *string `json:"fd_info,omitempty"`

	FdCount *int32 `json:"fd_count,omitempty"`
}

func (o EventFileResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventFileResponseInfo struct{}"
	}

	return strings.Join([]string{"EventFileResponseInfo", string(data)}, " ")
}
