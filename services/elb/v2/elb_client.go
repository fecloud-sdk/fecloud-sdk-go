package v2

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/elb/v2/model"
)

type ElbClient struct {
	HcClient *http_client.HcHttpClient
}

func NewElbClient(hcClient *http_client.HcHttpClient) *ElbClient {
	return &ElbClient{HcClient: hcClient}
}

func ElbClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *ElbClient) BatchCreateListenerTags(request *model.BatchCreateListenerTagsRequest) (*model.BatchCreateListenerTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateListenerTagsResponse), nil
	}
}

func (c *ElbClient) BatchCreateListenerTagsInvoker(request *model.BatchCreateListenerTagsRequest) *BatchCreateListenerTagsInvoker {
	requestDef := GenReqDefForBatchCreateListenerTags()
	return &BatchCreateListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) BatchCreateLoadbalancerTags(request *model.BatchCreateLoadbalancerTagsRequest) (*model.BatchCreateLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateLoadbalancerTagsResponse), nil
	}
}

func (c *ElbClient) BatchCreateLoadbalancerTagsInvoker(request *model.BatchCreateLoadbalancerTagsRequest) *BatchCreateLoadbalancerTagsInvoker {
	requestDef := GenReqDefForBatchCreateLoadbalancerTags()
	return &BatchCreateLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) BatchDeleteListenerTags(request *model.BatchDeleteListenerTagsRequest) (*model.BatchDeleteListenerTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteListenerTagsResponse), nil
	}
}

func (c *ElbClient) BatchDeleteListenerTagsInvoker(request *model.BatchDeleteListenerTagsRequest) *BatchDeleteListenerTagsInvoker {
	requestDef := GenReqDefForBatchDeleteListenerTags()
	return &BatchDeleteListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) BatchDeleteLoadbalancerTags(request *model.BatchDeleteLoadbalancerTagsRequest) (*model.BatchDeleteLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteLoadbalancerTagsResponse), nil
	}
}

func (c *ElbClient) BatchDeleteLoadbalancerTagsInvoker(request *model.BatchDeleteLoadbalancerTagsRequest) *BatchDeleteLoadbalancerTagsInvoker {
	requestDef := GenReqDefForBatchDeleteLoadbalancerTags()
	return &BatchDeleteLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateHealthmonitor(request *model.CreateHealthmonitorRequest) (*model.CreateHealthmonitorResponse, error) {
	requestDef := GenReqDefForCreateHealthmonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHealthmonitorResponse), nil
	}
}

func (c *ElbClient) CreateHealthmonitorInvoker(request *model.CreateHealthmonitorRequest) *CreateHealthmonitorInvoker {
	requestDef := GenReqDefForCreateHealthmonitor()
	return &CreateHealthmonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateL7policy(request *model.CreateL7policyRequest) (*model.CreateL7policyResponse, error) {
	requestDef := GenReqDefForCreateL7policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7policyResponse), nil
	}
}

func (c *ElbClient) CreateL7policyInvoker(request *model.CreateL7policyRequest) *CreateL7policyInvoker {
	requestDef := GenReqDefForCreateL7policy()
	return &CreateL7policyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateL7rule(request *model.CreateL7ruleRequest) (*model.CreateL7ruleResponse, error) {
	requestDef := GenReqDefForCreateL7rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7ruleResponse), nil
	}
}

func (c *ElbClient) CreateL7ruleInvoker(request *model.CreateL7ruleRequest) *CreateL7ruleInvoker {
	requestDef := GenReqDefForCreateL7rule()
	return &CreateL7ruleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateListener(request *model.CreateListenerRequest) (*model.CreateListenerResponse, error) {
	requestDef := GenReqDefForCreateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateListenerResponse), nil
	}
}

func (c *ElbClient) CreateListenerInvoker(request *model.CreateListenerRequest) *CreateListenerInvoker {
	requestDef := GenReqDefForCreateListener()
	return &CreateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateListenerTags(request *model.CreateListenerTagsRequest) (*model.CreateListenerTagsResponse, error) {
	requestDef := GenReqDefForCreateListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateListenerTagsResponse), nil
	}
}

func (c *ElbClient) CreateListenerTagsInvoker(request *model.CreateListenerTagsRequest) *CreateListenerTagsInvoker {
	requestDef := GenReqDefForCreateListenerTags()
	return &CreateListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateLoadbalancer(request *model.CreateLoadbalancerRequest) (*model.CreateLoadbalancerResponse, error) {
	requestDef := GenReqDefForCreateLoadbalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoadbalancerResponse), nil
	}
}

func (c *ElbClient) CreateLoadbalancerInvoker(request *model.CreateLoadbalancerRequest) *CreateLoadbalancerInvoker {
	requestDef := GenReqDefForCreateLoadbalancer()
	return &CreateLoadbalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateLoadbalancerTags(request *model.CreateLoadbalancerTagsRequest) (*model.CreateLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForCreateLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoadbalancerTagsResponse), nil
	}
}

func (c *ElbClient) CreateLoadbalancerTagsInvoker(request *model.CreateLoadbalancerTagsRequest) *CreateLoadbalancerTagsInvoker {
	requestDef := GenReqDefForCreateLoadbalancerTags()
	return &CreateLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateMember(request *model.CreateMemberRequest) (*model.CreateMemberResponse, error) {
	requestDef := GenReqDefForCreateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMemberResponse), nil
	}
}

func (c *ElbClient) CreateMemberInvoker(request *model.CreateMemberRequest) *CreateMemberInvoker {
	requestDef := GenReqDefForCreateMember()
	return &CreateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreatePool(request *model.CreatePoolRequest) (*model.CreatePoolResponse, error) {
	requestDef := GenReqDefForCreatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePoolResponse), nil
	}
}

func (c *ElbClient) CreatePoolInvoker(request *model.CreatePoolRequest) *CreatePoolInvoker {
	requestDef := GenReqDefForCreatePool()
	return &CreatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateWhitelist(request *model.CreateWhitelistRequest) (*model.CreateWhitelistResponse, error) {
	requestDef := GenReqDefForCreateWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWhitelistResponse), nil
	}
}

func (c *ElbClient) CreateWhitelistInvoker(request *model.CreateWhitelistRequest) *CreateWhitelistInvoker {
	requestDef := GenReqDefForCreateWhitelist()
	return &CreateWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteHealthmonitor(request *model.DeleteHealthmonitorRequest) (*model.DeleteHealthmonitorResponse, error) {
	requestDef := GenReqDefForDeleteHealthmonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHealthmonitorResponse), nil
	}
}

func (c *ElbClient) DeleteHealthmonitorInvoker(request *model.DeleteHealthmonitorRequest) *DeleteHealthmonitorInvoker {
	requestDef := GenReqDefForDeleteHealthmonitor()
	return &DeleteHealthmonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteL7policy(request *model.DeleteL7policyRequest) (*model.DeleteL7policyResponse, error) {
	requestDef := GenReqDefForDeleteL7policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7policyResponse), nil
	}
}

func (c *ElbClient) DeleteL7policyInvoker(request *model.DeleteL7policyRequest) *DeleteL7policyInvoker {
	requestDef := GenReqDefForDeleteL7policy()
	return &DeleteL7policyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteL7rule(request *model.DeleteL7ruleRequest) (*model.DeleteL7ruleResponse, error) {
	requestDef := GenReqDefForDeleteL7rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7ruleResponse), nil
	}
}

func (c *ElbClient) DeleteL7ruleInvoker(request *model.DeleteL7ruleRequest) *DeleteL7ruleInvoker {
	requestDef := GenReqDefForDeleteL7rule()
	return &DeleteL7ruleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteListener(request *model.DeleteListenerRequest) (*model.DeleteListenerResponse, error) {
	requestDef := GenReqDefForDeleteListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteListenerResponse), nil
	}
}

func (c *ElbClient) DeleteListenerInvoker(request *model.DeleteListenerRequest) *DeleteListenerInvoker {
	requestDef := GenReqDefForDeleteListener()
	return &DeleteListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteListenerTags(request *model.DeleteListenerTagsRequest) (*model.DeleteListenerTagsResponse, error) {
	requestDef := GenReqDefForDeleteListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteListenerTagsResponse), nil
	}
}

func (c *ElbClient) DeleteListenerTagsInvoker(request *model.DeleteListenerTagsRequest) *DeleteListenerTagsInvoker {
	requestDef := GenReqDefForDeleteListenerTags()
	return &DeleteListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteLoadbalancer(request *model.DeleteLoadbalancerRequest) (*model.DeleteLoadbalancerResponse, error) {
	requestDef := GenReqDefForDeleteLoadbalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoadbalancerResponse), nil
	}
}

func (c *ElbClient) DeleteLoadbalancerInvoker(request *model.DeleteLoadbalancerRequest) *DeleteLoadbalancerInvoker {
	requestDef := GenReqDefForDeleteLoadbalancer()
	return &DeleteLoadbalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteLoadbalancerTags(request *model.DeleteLoadbalancerTagsRequest) (*model.DeleteLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForDeleteLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoadbalancerTagsResponse), nil
	}
}

func (c *ElbClient) DeleteLoadbalancerTagsInvoker(request *model.DeleteLoadbalancerTagsRequest) *DeleteLoadbalancerTagsInvoker {
	requestDef := GenReqDefForDeleteLoadbalancerTags()
	return &DeleteLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

func (c *ElbClient) DeleteMemberInvoker(request *model.DeleteMemberRequest) *DeleteMemberInvoker {
	requestDef := GenReqDefForDeleteMember()
	return &DeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeletePool(request *model.DeletePoolRequest) (*model.DeletePoolResponse, error) {
	requestDef := GenReqDefForDeletePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePoolResponse), nil
	}
}

func (c *ElbClient) DeletePoolInvoker(request *model.DeletePoolRequest) *DeletePoolInvoker {
	requestDef := GenReqDefForDeletePool()
	return &DeletePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteWhitelist(request *model.DeleteWhitelistRequest) (*model.DeleteWhitelistResponse, error) {
	requestDef := GenReqDefForDeleteWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWhitelistResponse), nil
	}
}

func (c *ElbClient) DeleteWhitelistInvoker(request *model.DeleteWhitelistRequest) *DeleteWhitelistInvoker {
	requestDef := GenReqDefForDeleteWhitelist()
	return &DeleteWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListHealthmonitors(request *model.ListHealthmonitorsRequest) (*model.ListHealthmonitorsResponse, error) {
	requestDef := GenReqDefForListHealthmonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHealthmonitorsResponse), nil
	}
}

func (c *ElbClient) ListHealthmonitorsInvoker(request *model.ListHealthmonitorsRequest) *ListHealthmonitorsInvoker {
	requestDef := GenReqDefForListHealthmonitors()
	return &ListHealthmonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListL7policies(request *model.ListL7policiesRequest) (*model.ListL7policiesResponse, error) {
	requestDef := GenReqDefForListL7policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7policiesResponse), nil
	}
}

func (c *ElbClient) ListL7policiesInvoker(request *model.ListL7policiesRequest) *ListL7policiesInvoker {
	requestDef := GenReqDefForListL7policies()
	return &ListL7policiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListL7rules(request *model.ListL7rulesRequest) (*model.ListL7rulesResponse, error) {
	requestDef := GenReqDefForListL7rules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7rulesResponse), nil
	}
}

func (c *ElbClient) ListL7rulesInvoker(request *model.ListL7rulesRequest) *ListL7rulesInvoker {
	requestDef := GenReqDefForListL7rules()
	return &ListL7rulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListListenerTags(request *model.ListListenerTagsRequest) (*model.ListListenerTagsResponse, error) {
	requestDef := GenReqDefForListListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenerTagsResponse), nil
	}
}

func (c *ElbClient) ListListenerTagsInvoker(request *model.ListListenerTagsRequest) *ListListenerTagsInvoker {
	requestDef := GenReqDefForListListenerTags()
	return &ListListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListListeners(request *model.ListListenersRequest) (*model.ListListenersResponse, error) {
	requestDef := GenReqDefForListListeners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenersResponse), nil
	}
}

func (c *ElbClient) ListListenersInvoker(request *model.ListListenersRequest) *ListListenersInvoker {
	requestDef := GenReqDefForListListeners()
	return &ListListenersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListListenersByTags(request *model.ListListenersByTagsRequest) (*model.ListListenersByTagsResponse, error) {
	requestDef := GenReqDefForListListenersByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenersByTagsResponse), nil
	}
}

func (c *ElbClient) ListListenersByTagsInvoker(request *model.ListListenersByTagsRequest) *ListListenersByTagsInvoker {
	requestDef := GenReqDefForListListenersByTags()
	return &ListListenersByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListLoadbalancerTags(request *model.ListLoadbalancerTagsRequest) (*model.ListLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForListLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadbalancerTagsResponse), nil
	}
}

func (c *ElbClient) ListLoadbalancerTagsInvoker(request *model.ListLoadbalancerTagsRequest) *ListLoadbalancerTagsInvoker {
	requestDef := GenReqDefForListLoadbalancerTags()
	return &ListLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListLoadbalancers(request *model.ListLoadbalancersRequest) (*model.ListLoadbalancersResponse, error) {
	requestDef := GenReqDefForListLoadbalancers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadbalancersResponse), nil
	}
}

func (c *ElbClient) ListLoadbalancersInvoker(request *model.ListLoadbalancersRequest) *ListLoadbalancersInvoker {
	requestDef := GenReqDefForListLoadbalancers()
	return &ListLoadbalancersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListLoadbalancersByTags(request *model.ListLoadbalancersByTagsRequest) (*model.ListLoadbalancersByTagsResponse, error) {
	requestDef := GenReqDefForListLoadbalancersByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadbalancersByTagsResponse), nil
	}
}

func (c *ElbClient) ListLoadbalancersByTagsInvoker(request *model.ListLoadbalancersByTagsRequest) *ListLoadbalancersByTagsInvoker {
	requestDef := GenReqDefForListLoadbalancersByTags()
	return &ListLoadbalancersByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

func (c *ElbClient) ListMembersInvoker(request *model.ListMembersRequest) *ListMembersInvoker {
	requestDef := GenReqDefForListMembers()
	return &ListMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListPools(request *model.ListPoolsRequest) (*model.ListPoolsResponse, error) {
	requestDef := GenReqDefForListPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolsResponse), nil
	}
}

func (c *ElbClient) ListPoolsInvoker(request *model.ListPoolsRequest) *ListPoolsInvoker {
	requestDef := GenReqDefForListPools()
	return &ListPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListWhitelists(request *model.ListWhitelistsRequest) (*model.ListWhitelistsResponse, error) {
	requestDef := GenReqDefForListWhitelists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWhitelistsResponse), nil
	}
}

func (c *ElbClient) ListWhitelistsInvoker(request *model.ListWhitelistsRequest) *ListWhitelistsInvoker {
	requestDef := GenReqDefForListWhitelists()
	return &ListWhitelistsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowHealthmonitors(request *model.ShowHealthmonitorsRequest) (*model.ShowHealthmonitorsResponse, error) {
	requestDef := GenReqDefForShowHealthmonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthmonitorsResponse), nil
	}
}

func (c *ElbClient) ShowHealthmonitorsInvoker(request *model.ShowHealthmonitorsRequest) *ShowHealthmonitorsInvoker {
	requestDef := GenReqDefForShowHealthmonitors()
	return &ShowHealthmonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowL7policy(request *model.ShowL7policyRequest) (*model.ShowL7policyResponse, error) {
	requestDef := GenReqDefForShowL7policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7policyResponse), nil
	}
}

func (c *ElbClient) ShowL7policyInvoker(request *model.ShowL7policyRequest) *ShowL7policyInvoker {
	requestDef := GenReqDefForShowL7policy()
	return &ShowL7policyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowL7rule(request *model.ShowL7ruleRequest) (*model.ShowL7ruleResponse, error) {
	requestDef := GenReqDefForShowL7rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7ruleResponse), nil
	}
}

func (c *ElbClient) ShowL7ruleInvoker(request *model.ShowL7ruleRequest) *ShowL7ruleInvoker {
	requestDef := GenReqDefForShowL7rule()
	return &ShowL7ruleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowListener(request *model.ShowListenerRequest) (*model.ShowListenerResponse, error) {
	requestDef := GenReqDefForShowListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListenerResponse), nil
	}
}

func (c *ElbClient) ShowListenerInvoker(request *model.ShowListenerRequest) *ShowListenerInvoker {
	requestDef := GenReqDefForShowListener()
	return &ShowListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowListenerTags(request *model.ShowListenerTagsRequest) (*model.ShowListenerTagsResponse, error) {
	requestDef := GenReqDefForShowListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListenerTagsResponse), nil
	}
}

func (c *ElbClient) ShowListenerTagsInvoker(request *model.ShowListenerTagsRequest) *ShowListenerTagsInvoker {
	requestDef := GenReqDefForShowListenerTags()
	return &ShowListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowLoadbalancer(request *model.ShowLoadbalancerRequest) (*model.ShowLoadbalancerResponse, error) {
	requestDef := GenReqDefForShowLoadbalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadbalancerResponse), nil
	}
}

func (c *ElbClient) ShowLoadbalancerInvoker(request *model.ShowLoadbalancerRequest) *ShowLoadbalancerInvoker {
	requestDef := GenReqDefForShowLoadbalancer()
	return &ShowLoadbalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowLoadbalancerTags(request *model.ShowLoadbalancerTagsRequest) (*model.ShowLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForShowLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadbalancerTagsResponse), nil
	}
}

func (c *ElbClient) ShowLoadbalancerTagsInvoker(request *model.ShowLoadbalancerTagsRequest) *ShowLoadbalancerTagsInvoker {
	requestDef := GenReqDefForShowLoadbalancerTags()
	return &ShowLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowLoadbalancersStatus(request *model.ShowLoadbalancersStatusRequest) (*model.ShowLoadbalancersStatusResponse, error) {
	requestDef := GenReqDefForShowLoadbalancersStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadbalancersStatusResponse), nil
	}
}

func (c *ElbClient) ShowLoadbalancersStatusInvoker(request *model.ShowLoadbalancersStatusRequest) *ShowLoadbalancersStatusInvoker {
	requestDef := GenReqDefForShowLoadbalancersStatus()
	return &ShowLoadbalancersStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowMember(request *model.ShowMemberRequest) (*model.ShowMemberResponse, error) {
	requestDef := GenReqDefForShowMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMemberResponse), nil
	}
}

func (c *ElbClient) ShowMemberInvoker(request *model.ShowMemberRequest) *ShowMemberInvoker {
	requestDef := GenReqDefForShowMember()
	return &ShowMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowPool(request *model.ShowPoolRequest) (*model.ShowPoolResponse, error) {
	requestDef := GenReqDefForShowPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolResponse), nil
	}
}

func (c *ElbClient) ShowPoolInvoker(request *model.ShowPoolRequest) *ShowPoolInvoker {
	requestDef := GenReqDefForShowPool()
	return &ShowPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowWhitelist(request *model.ShowWhitelistRequest) (*model.ShowWhitelistResponse, error) {
	requestDef := GenReqDefForShowWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWhitelistResponse), nil
	}
}

func (c *ElbClient) ShowWhitelistInvoker(request *model.ShowWhitelistRequest) *ShowWhitelistInvoker {
	requestDef := GenReqDefForShowWhitelist()
	return &ShowWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateHealthmonitor(request *model.UpdateHealthmonitorRequest) (*model.UpdateHealthmonitorResponse, error) {
	requestDef := GenReqDefForUpdateHealthmonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthmonitorResponse), nil
	}
}

func (c *ElbClient) UpdateHealthmonitorInvoker(request *model.UpdateHealthmonitorRequest) *UpdateHealthmonitorInvoker {
	requestDef := GenReqDefForUpdateHealthmonitor()
	return &UpdateHealthmonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateL7policies(request *model.UpdateL7policiesRequest) (*model.UpdateL7policiesResponse, error) {
	requestDef := GenReqDefForUpdateL7policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7policiesResponse), nil
	}
}

func (c *ElbClient) UpdateL7policiesInvoker(request *model.UpdateL7policiesRequest) *UpdateL7policiesInvoker {
	requestDef := GenReqDefForUpdateL7policies()
	return &UpdateL7policiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateL7rule(request *model.UpdateL7ruleRequest) (*model.UpdateL7ruleResponse, error) {
	requestDef := GenReqDefForUpdateL7rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7ruleResponse), nil
	}
}

func (c *ElbClient) UpdateL7ruleInvoker(request *model.UpdateL7ruleRequest) *UpdateL7ruleInvoker {
	requestDef := GenReqDefForUpdateL7rule()
	return &UpdateL7ruleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateListener(request *model.UpdateListenerRequest) (*model.UpdateListenerResponse, error) {
	requestDef := GenReqDefForUpdateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateListenerResponse), nil
	}
}

func (c *ElbClient) UpdateListenerInvoker(request *model.UpdateListenerRequest) *UpdateListenerInvoker {
	requestDef := GenReqDefForUpdateListener()
	return &UpdateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateLoadbalancer(request *model.UpdateLoadbalancerRequest) (*model.UpdateLoadbalancerResponse, error) {
	requestDef := GenReqDefForUpdateLoadbalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLoadbalancerResponse), nil
	}
}

func (c *ElbClient) UpdateLoadbalancerInvoker(request *model.UpdateLoadbalancerRequest) *UpdateLoadbalancerInvoker {
	requestDef := GenReqDefForUpdateLoadbalancer()
	return &UpdateLoadbalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateMember(request *model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	requestDef := GenReqDefForUpdateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberResponse), nil
	}
}

func (c *ElbClient) UpdateMemberInvoker(request *model.UpdateMemberRequest) *UpdateMemberInvoker {
	requestDef := GenReqDefForUpdateMember()
	return &UpdateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdatePool(request *model.UpdatePoolRequest) (*model.UpdatePoolResponse, error) {
	requestDef := GenReqDefForUpdatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePoolResponse), nil
	}
}

func (c *ElbClient) UpdatePoolInvoker(request *model.UpdatePoolRequest) *UpdatePoolInvoker {
	requestDef := GenReqDefForUpdatePool()
	return &UpdatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateWhitelist(request *model.UpdateWhitelistRequest) (*model.UpdateWhitelistResponse, error) {
	requestDef := GenReqDefForUpdateWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWhitelistResponse), nil
	}
}

func (c *ElbClient) UpdateWhitelistInvoker(request *model.UpdateWhitelistRequest) *UpdateWhitelistInvoker {
	requestDef := GenReqDefForUpdateWhitelist()
	return &UpdateWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

func (c *ElbClient) CreateCertificateInvoker(request *model.CreateCertificateRequest) *CreateCertificateInvoker {
	requestDef := GenReqDefForCreateCertificate()
	return &CreateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

func (c *ElbClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

func (c *ElbClient) ListCertificatesInvoker(request *model.ListCertificatesRequest) *ListCertificatesInvoker {
	requestDef := GenReqDefForListCertificates()
	return &ListCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

func (c *ElbClient) ShowCertificateInvoker(request *model.ShowCertificateRequest) *ShowCertificateInvoker {
	requestDef := GenReqDefForShowCertificate()
	return &ShowCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *ElbClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

func (c *ElbClient) UpdateCertificateInvoker(request *model.UpdateCertificateRequest) *UpdateCertificateInvoker {
	requestDef := GenReqDefForUpdateCertificate()
	return &UpdateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
