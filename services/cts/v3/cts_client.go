package v3

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/cts/v3/model"
)

type CtsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCtsClient(hcClient *http_client.HcHttpClient) *CtsClient {
	return &CtsClient{HcClient: hcClient}
}

func CtsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *CtsClient) CreateTracker(request *model.CreateTrackerRequest) (*model.CreateTrackerResponse, error) {
	requestDef := GenReqDefForCreateTracker()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrackerResponse), nil
	}
}

func (c *CtsClient) CreateTrackerInvoker(request *model.CreateTrackerRequest) *CreateTrackerInvoker {
	requestDef := GenReqDefForCreateTracker()
	return &CreateTrackerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) DeleteTracker(request *model.DeleteTrackerRequest) (*model.DeleteTrackerResponse, error) {
	requestDef := GenReqDefForDeleteTracker()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrackerResponse), nil
	}
}

func (c *CtsClient) DeleteTrackerInvoker(request *model.DeleteTrackerRequest) *DeleteTrackerInvoker {
	requestDef := GenReqDefForDeleteTracker()
	return &DeleteTrackerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) ListTraces(request *model.ListTracesRequest) (*model.ListTracesResponse, error) {
	requestDef := GenReqDefForListTraces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTracesResponse), nil
	}
}

func (c *CtsClient) ListTracesInvoker(request *model.ListTracesRequest) *ListTracesInvoker {
	requestDef := GenReqDefForListTraces()
	return &ListTracesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) ListTrackers(request *model.ListTrackersRequest) (*model.ListTrackersResponse, error) {
	requestDef := GenReqDefForListTrackers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrackersResponse), nil
	}
}

func (c *CtsClient) ListTrackersInvoker(request *model.ListTrackersRequest) *ListTrackersInvoker {
	requestDef := GenReqDefForListTrackers()
	return &ListTrackersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) UpdateTracker(request *model.UpdateTrackerRequest) (*model.UpdateTrackerResponse, error) {
	requestDef := GenReqDefForUpdateTracker()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTrackerResponse), nil
	}
}

func (c *CtsClient) UpdateTrackerInvoker(request *model.UpdateTrackerRequest) *UpdateTrackerInvoker {
	requestDef := GenReqDefForUpdateTracker()
	return &UpdateTrackerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
