package v2

import "github.com/fecloud-sdk/fecloud-sdk-go/services/lts/v2/model"

func (c *LtsClient) GetCreateLogDumpObsRequest() *model.CreateLogDumpObsRequest {
	return new(model.CreateLogDumpObsRequest)
}

func (c *LtsClient) GetCreateLogGroupRequest() *model.CreateLogGroupRequest {
	return new(model.CreateLogGroupRequest)
}

func (c *LtsClient) GetCreateLogStreamRequest() *model.CreateLogStreamRequest {
	return new(model.CreateLogStreamRequest)
}

func (c *LtsClient) GetDeleteLogGroupRequest() *model.DeleteLogGroupRequest {
	return new(model.DeleteLogGroupRequest)
}

func (c *LtsClient) GetDeleteLogStreamRequest() *model.DeleteLogStreamRequest {
	return new(model.DeleteLogStreamRequest)
}

func (c *LtsClient) GetListLogGroupsRequest() *model.ListLogGroupsRequest {
	return new(model.ListLogGroupsRequest)
}

func (c *LtsClient) GetListLogStreamRequest() *model.ListLogStreamRequest {
	return new(model.ListLogStreamRequest)
}

func (c *LtsClient) GetListLogsRequest() *model.ListLogsRequest {
	return new(model.ListLogsRequest)
}
