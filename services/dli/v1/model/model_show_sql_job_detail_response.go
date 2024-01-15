package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSqlJobDetailResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	Owner *string `json:"owner,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	Duration *int64 `json:"duration,omitempty"`

	ExportMode *string `json:"export_mode,omitempty"`

	DatabaseName *string `json:"database_name,omitempty"`

	TableName *string `json:"table_name,omitempty"`

	DataPath *string `json:"data_path,omitempty"`

	DataType *string `json:"data_type,omitempty"`

	WithColumnHeader *bool `json:"with_column_header,omitempty"`

	Delimiter *string `json:"delimiter,omitempty"`

	QuoteChar *string `json:"quote_char,omitempty"`

	EscapeChar *string `json:"escape_char,omitempty"`

	DateFormat *string `json:"date_format,omitempty"`

	TimestampFormat *string `json:"timestamp_format,omitempty"`

	Compress *string `json:"compress,omitempty"`

	Tags           *[]TmsTag `json:"tags,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowSqlJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlJobDetailResponse", string(data)}, " ")
}
