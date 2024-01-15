package v2

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/lts/v2/model"
)

type LtsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewLtsClient(hcClient *http_client.HcHttpClient) *LtsClient {
	return &LtsClient{HcClient: hcClient}
}

func LtsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *LtsClient) CreateLogDumpObs(request *model.CreateLogDumpObsRequest) (*model.CreateLogDumpObsResponse, error) {
	requestDef := GenReqDefForCreateLogDumpObs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogDumpObsResponse), nil
	}
}

func (c *LtsClient) CreateLogDumpObsInvoker(request *model.CreateLogDumpObsRequest) *CreateLogDumpObsInvoker {
	requestDef := GenReqDefForCreateLogDumpObs()
	return &CreateLogDumpObsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *LtsClient) CreateLogGroup(request *model.CreateLogGroupRequest) (*model.CreateLogGroupResponse, error) {
	requestDef := GenReqDefForCreateLogGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogGroupResponse), nil
	}
}

func (c *LtsClient) CreateLogGroupInvoker(request *model.CreateLogGroupRequest) *CreateLogGroupInvoker {
	requestDef := GenReqDefForCreateLogGroup()
	return &CreateLogGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *LtsClient) CreateLogStream(request *model.CreateLogStreamRequest) (*model.CreateLogStreamResponse, error) {
	requestDef := GenReqDefForCreateLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogStreamResponse), nil
	}
}

func (c *LtsClient) CreateLogStreamInvoker(request *model.CreateLogStreamRequest) *CreateLogStreamInvoker {
	requestDef := GenReqDefForCreateLogStream()
	return &CreateLogStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *LtsClient) DeleteLogGroup(request *model.DeleteLogGroupRequest) (*model.DeleteLogGroupResponse, error) {
	requestDef := GenReqDefForDeleteLogGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogGroupResponse), nil
	}
}

func (c *LtsClient) DeleteLogGroupInvoker(request *model.DeleteLogGroupRequest) *DeleteLogGroupInvoker {
	requestDef := GenReqDefForDeleteLogGroup()
	return &DeleteLogGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *LtsClient) DeleteLogStream(request *model.DeleteLogStreamRequest) (*model.DeleteLogStreamResponse, error) {
	requestDef := GenReqDefForDeleteLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogStreamResponse), nil
	}
}

func (c *LtsClient) DeleteLogStreamInvoker(request *model.DeleteLogStreamRequest) *DeleteLogStreamInvoker {
	requestDef := GenReqDefForDeleteLogStream()
	return &DeleteLogStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *LtsClient) ListLogGroups(request *model.ListLogGroupsRequest) (*model.ListLogGroupsResponse, error) {
	requestDef := GenReqDefForListLogGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogGroupsResponse), nil
	}
}

func (c *LtsClient) ListLogGroupsInvoker(request *model.ListLogGroupsRequest) *ListLogGroupsInvoker {
	requestDef := GenReqDefForListLogGroups()
	return &ListLogGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *LtsClient) ListLogStream(request *model.ListLogStreamRequest) (*model.ListLogStreamResponse, error) {
	requestDef := GenReqDefForListLogStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogStreamResponse), nil
	}
}

func (c *LtsClient) ListLogStreamInvoker(request *model.ListLogStreamRequest) *ListLogStreamInvoker {
	requestDef := GenReqDefForListLogStream()
	return &ListLogStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *LtsClient) ListLogs(request *model.ListLogsRequest) (*model.ListLogsResponse, error) {
	requestDef := GenReqDefForListLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogsResponse), nil
	}
}

func (c *LtsClient) ListLogsInvoker(request *model.ListLogsRequest) *ListLogsInvoker {
	requestDef := GenReqDefForListLogs()
	return &ListLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
