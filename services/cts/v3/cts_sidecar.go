package v3

import "github.com/fecloud-sdk/fecloud-sdk-go/services/cts/v3/model"

func (c *CtsClient) GetCreateTrackerRequest() *model.CreateTrackerRequest {
	return new(model.CreateTrackerRequest)
}

func (c *CtsClient) GetDeleteTrackerRequest() *model.DeleteTrackerRequest {
	return new(model.DeleteTrackerRequest)
}

func (c *CtsClient) GetListTracesRequest() *model.ListTracesRequest {
	return new(model.ListTracesRequest)
}

func (c *CtsClient) GetListTrackersRequest() *model.ListTrackersRequest {
	return new(model.ListTrackersRequest)
}

func (c *CtsClient) GetUpdateTrackerRequest() *model.UpdateTrackerRequest {
	return new(model.UpdateTrackerRequest)
}
