package v5

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/hss/v5/model"
)

type HssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewHssClient(hcClient *http_client.HcHttpClient) *HssClient {
	return &HssClient{HcClient: hcClient}
}

func HssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *HssClient) AddHostsGroup(request *model.AddHostsGroupRequest) (*model.AddHostsGroupResponse, error) {
	requestDef := GenReqDefForAddHostsGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddHostsGroupResponse), nil
	}
}

func (c *HssClient) AddHostsGroupInvoker(request *model.AddHostsGroupRequest) *AddHostsGroupInvoker {
	requestDef := GenReqDefForAddHostsGroup()
	return &AddHostsGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) AssociatePolicyGroup(request *model.AssociatePolicyGroupRequest) (*model.AssociatePolicyGroupResponse, error) {
	requestDef := GenReqDefForAssociatePolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociatePolicyGroupResponse), nil
	}
}

func (c *HssClient) AssociatePolicyGroupInvoker(request *model.AssociatePolicyGroupRequest) *AssociatePolicyGroupInvoker {
	requestDef := GenReqDefForAssociatePolicyGroup()
	return &AssociatePolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) BatchCreateTags(request *model.BatchCreateTagsRequest) (*model.BatchCreateTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagsResponse), nil
	}
}

func (c *HssClient) BatchCreateTagsInvoker(request *model.BatchCreateTagsRequest) *BatchCreateTagsInvoker {
	requestDef := GenReqDefForBatchCreateTags()
	return &BatchCreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ChangeEvent(request *model.ChangeEventRequest) (*model.ChangeEventResponse, error) {
	requestDef := GenReqDefForChangeEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEventResponse), nil
	}
}

func (c *HssClient) ChangeEventInvoker(request *model.ChangeEventRequest) *ChangeEventInvoker {
	requestDef := GenReqDefForChangeEvent()
	return &ChangeEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ChangeHostsGroup(request *model.ChangeHostsGroupRequest) (*model.ChangeHostsGroupResponse, error) {
	requestDef := GenReqDefForChangeHostsGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeHostsGroupResponse), nil
	}
}

func (c *HssClient) ChangeHostsGroupInvoker(request *model.ChangeHostsGroupRequest) *ChangeHostsGroupInvoker {
	requestDef := GenReqDefForChangeHostsGroup()
	return &ChangeHostsGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ChangeVulStatus(request *model.ChangeVulStatusRequest) (*model.ChangeVulStatusResponse, error) {
	requestDef := GenReqDefForChangeVulStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeVulStatusResponse), nil
	}
}

func (c *HssClient) ChangeVulStatusInvoker(request *model.ChangeVulStatusRequest) *ChangeVulStatusInvoker {
	requestDef := GenReqDefForChangeVulStatus()
	return &ChangeVulStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) DeleteHostsGroup(request *model.DeleteHostsGroupRequest) (*model.DeleteHostsGroupResponse, error) {
	requestDef := GenReqDefForDeleteHostsGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostsGroupResponse), nil
	}
}

func (c *HssClient) DeleteHostsGroupInvoker(request *model.DeleteHostsGroupRequest) *DeleteHostsGroupInvoker {
	requestDef := GenReqDefForDeleteHostsGroup()
	return &DeleteHostsGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) DeleteResourceInstanceTag(request *model.DeleteResourceInstanceTagRequest) (*model.DeleteResourceInstanceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceInstanceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceInstanceTagResponse), nil
	}
}

func (c *HssClient) DeleteResourceInstanceTagInvoker(request *model.DeleteResourceInstanceTagRequest) *DeleteResourceInstanceTagInvoker {
	requestDef := GenReqDefForDeleteResourceInstanceTag()
	return &DeleteResourceInstanceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListAlarmWhiteList(request *model.ListAlarmWhiteListRequest) (*model.ListAlarmWhiteListResponse, error) {
	requestDef := GenReqDefForListAlarmWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmWhiteListResponse), nil
	}
}

func (c *HssClient) ListAlarmWhiteListInvoker(request *model.ListAlarmWhiteListRequest) *ListAlarmWhiteListInvoker {
	requestDef := GenReqDefForListAlarmWhiteList()
	return &ListAlarmWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListAppChangeHistories(request *model.ListAppChangeHistoriesRequest) (*model.ListAppChangeHistoriesResponse, error) {
	requestDef := GenReqDefForListAppChangeHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppChangeHistoriesResponse), nil
	}
}

func (c *HssClient) ListAppChangeHistoriesInvoker(request *model.ListAppChangeHistoriesRequest) *ListAppChangeHistoriesInvoker {
	requestDef := GenReqDefForListAppChangeHistories()
	return &ListAppChangeHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListAppStatistics(request *model.ListAppStatisticsRequest) (*model.ListAppStatisticsResponse, error) {
	requestDef := GenReqDefForListAppStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppStatisticsResponse), nil
	}
}

func (c *HssClient) ListAppStatisticsInvoker(request *model.ListAppStatisticsRequest) *ListAppStatisticsInvoker {
	requestDef := GenReqDefForListAppStatistics()
	return &ListAppStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

func (c *HssClient) ListAppsInvoker(request *model.ListAppsRequest) *ListAppsInvoker {
	requestDef := GenReqDefForListApps()
	return &ListAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListAutoLaunchChangeHistories(request *model.ListAutoLaunchChangeHistoriesRequest) (*model.ListAutoLaunchChangeHistoriesResponse, error) {
	requestDef := GenReqDefForListAutoLaunchChangeHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoLaunchChangeHistoriesResponse), nil
	}
}

func (c *HssClient) ListAutoLaunchChangeHistoriesInvoker(request *model.ListAutoLaunchChangeHistoriesRequest) *ListAutoLaunchChangeHistoriesInvoker {
	requestDef := GenReqDefForListAutoLaunchChangeHistories()
	return &ListAutoLaunchChangeHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListAutoLaunchStatistics(request *model.ListAutoLaunchStatisticsRequest) (*model.ListAutoLaunchStatisticsResponse, error) {
	requestDef := GenReqDefForListAutoLaunchStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoLaunchStatisticsResponse), nil
	}
}

func (c *HssClient) ListAutoLaunchStatisticsInvoker(request *model.ListAutoLaunchStatisticsRequest) *ListAutoLaunchStatisticsInvoker {
	requestDef := GenReqDefForListAutoLaunchStatistics()
	return &ListAutoLaunchStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListAutoLaunchs(request *model.ListAutoLaunchsRequest) (*model.ListAutoLaunchsResponse, error) {
	requestDef := GenReqDefForListAutoLaunchs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoLaunchsResponse), nil
	}
}

func (c *HssClient) ListAutoLaunchsInvoker(request *model.ListAutoLaunchsRequest) *ListAutoLaunchsInvoker {
	requestDef := GenReqDefForListAutoLaunchs()
	return &ListAutoLaunchsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListHostGroups(request *model.ListHostGroupsRequest) (*model.ListHostGroupsResponse, error) {
	requestDef := GenReqDefForListHostGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostGroupsResponse), nil
	}
}

func (c *HssClient) ListHostGroupsInvoker(request *model.ListHostGroupsRequest) *ListHostGroupsInvoker {
	requestDef := GenReqDefForListHostGroups()
	return &ListHostGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListHostProtectHistoryInfo(request *model.ListHostProtectHistoryInfoRequest) (*model.ListHostProtectHistoryInfoResponse, error) {
	requestDef := GenReqDefForListHostProtectHistoryInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostProtectHistoryInfoResponse), nil
	}
}

func (c *HssClient) ListHostProtectHistoryInfoInvoker(request *model.ListHostProtectHistoryInfoRequest) *ListHostProtectHistoryInfoInvoker {
	requestDef := GenReqDefForListHostProtectHistoryInfo()
	return &ListHostProtectHistoryInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListHostRaspProtectHistoryInfo(request *model.ListHostRaspProtectHistoryInfoRequest) (*model.ListHostRaspProtectHistoryInfoResponse, error) {
	requestDef := GenReqDefForListHostRaspProtectHistoryInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostRaspProtectHistoryInfoResponse), nil
	}
}

func (c *HssClient) ListHostRaspProtectHistoryInfoInvoker(request *model.ListHostRaspProtectHistoryInfoRequest) *ListHostRaspProtectHistoryInfoInvoker {
	requestDef := GenReqDefForListHostRaspProtectHistoryInfo()
	return &ListHostRaspProtectHistoryInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListHostStatus(request *model.ListHostStatusRequest) (*model.ListHostStatusResponse, error) {
	requestDef := GenReqDefForListHostStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostStatusResponse), nil
	}
}

func (c *HssClient) ListHostStatusInvoker(request *model.ListHostStatusRequest) *ListHostStatusInvoker {
	requestDef := GenReqDefForListHostStatus()
	return &ListHostStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListHostVuls(request *model.ListHostVulsRequest) (*model.ListHostVulsResponse, error) {
	requestDef := GenReqDefForListHostVuls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostVulsResponse), nil
	}
}

func (c *HssClient) ListHostVulsInvoker(request *model.ListHostVulsRequest) *ListHostVulsInvoker {
	requestDef := GenReqDefForListHostVuls()
	return &ListHostVulsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListJarPackageHostInfo(request *model.ListJarPackageHostInfoRequest) (*model.ListJarPackageHostInfoResponse, error) {
	requestDef := GenReqDefForListJarPackageHostInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJarPackageHostInfoResponse), nil
	}
}

func (c *HssClient) ListJarPackageHostInfoInvoker(request *model.ListJarPackageHostInfoRequest) *ListJarPackageHostInfoInvoker {
	requestDef := GenReqDefForListJarPackageHostInfo()
	return &ListJarPackageHostInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListJarPackageStatistics(request *model.ListJarPackageStatisticsRequest) (*model.ListJarPackageStatisticsResponse, error) {
	requestDef := GenReqDefForListJarPackageStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJarPackageStatisticsResponse), nil
	}
}

func (c *HssClient) ListJarPackageStatisticsInvoker(request *model.ListJarPackageStatisticsRequest) *ListJarPackageStatisticsInvoker {
	requestDef := GenReqDefForListJarPackageStatistics()
	return &ListJarPackageStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListPasswordComplexity(request *model.ListPasswordComplexityRequest) (*model.ListPasswordComplexityResponse, error) {
	requestDef := GenReqDefForListPasswordComplexity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPasswordComplexityResponse), nil
	}
}

func (c *HssClient) ListPasswordComplexityInvoker(request *model.ListPasswordComplexityRequest) *ListPasswordComplexityInvoker {
	requestDef := GenReqDefForListPasswordComplexity()
	return &ListPasswordComplexityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListPolicyGroup(request *model.ListPolicyGroupRequest) (*model.ListPolicyGroupResponse, error) {
	requestDef := GenReqDefForListPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyGroupResponse), nil
	}
}

func (c *HssClient) ListPolicyGroupInvoker(request *model.ListPolicyGroupRequest) *ListPolicyGroupInvoker {
	requestDef := GenReqDefForListPolicyGroup()
	return &ListPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListPortStatistics(request *model.ListPortStatisticsRequest) (*model.ListPortStatisticsResponse, error) {
	requestDef := GenReqDefForListPortStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortStatisticsResponse), nil
	}
}

func (c *HssClient) ListPortStatisticsInvoker(request *model.ListPortStatisticsRequest) *ListPortStatisticsInvoker {
	requestDef := GenReqDefForListPortStatistics()
	return &ListPortStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListPorts(request *model.ListPortsRequest) (*model.ListPortsResponse, error) {
	requestDef := GenReqDefForListPorts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortsResponse), nil
	}
}

func (c *HssClient) ListPortsInvoker(request *model.ListPortsRequest) *ListPortsInvoker {
	requestDef := GenReqDefForListPorts()
	return &ListPortsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListProcessStatistics(request *model.ListProcessStatisticsRequest) (*model.ListProcessStatisticsResponse, error) {
	requestDef := GenReqDefForListProcessStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProcessStatisticsResponse), nil
	}
}

func (c *HssClient) ListProcessStatisticsInvoker(request *model.ListProcessStatisticsRequest) *ListProcessStatisticsInvoker {
	requestDef := GenReqDefForListProcessStatistics()
	return &ListProcessStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListProtectionPolicy(request *model.ListProtectionPolicyRequest) (*model.ListProtectionPolicyResponse, error) {
	requestDef := GenReqDefForListProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectionPolicyResponse), nil
	}
}

func (c *HssClient) ListProtectionPolicyInvoker(request *model.ListProtectionPolicyRequest) *ListProtectionPolicyInvoker {
	requestDef := GenReqDefForListProtectionPolicy()
	return &ListProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListProtectionServer(request *model.ListProtectionServerRequest) (*model.ListProtectionServerResponse, error) {
	requestDef := GenReqDefForListProtectionServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectionServerResponse), nil
	}
}

func (c *HssClient) ListProtectionServerInvoker(request *model.ListProtectionServerRequest) *ListProtectionServerInvoker {
	requestDef := GenReqDefForListProtectionServer()
	return &ListProtectionServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListQuotasDetail(request *model.ListQuotasDetailRequest) (*model.ListQuotasDetailResponse, error) {
	requestDef := GenReqDefForListQuotasDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasDetailResponse), nil
	}
}

func (c *HssClient) ListQuotasDetailInvoker(request *model.ListQuotasDetailRequest) *ListQuotasDetailInvoker {
	requestDef := GenReqDefForListQuotasDetail()
	return &ListQuotasDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListRiskConfigCheckRules(request *model.ListRiskConfigCheckRulesRequest) (*model.ListRiskConfigCheckRulesResponse, error) {
	requestDef := GenReqDefForListRiskConfigCheckRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskConfigCheckRulesResponse), nil
	}
}

func (c *HssClient) ListRiskConfigCheckRulesInvoker(request *model.ListRiskConfigCheckRulesRequest) *ListRiskConfigCheckRulesInvoker {
	requestDef := GenReqDefForListRiskConfigCheckRules()
	return &ListRiskConfigCheckRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListRiskConfigHosts(request *model.ListRiskConfigHostsRequest) (*model.ListRiskConfigHostsResponse, error) {
	requestDef := GenReqDefForListRiskConfigHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskConfigHostsResponse), nil
	}
}

func (c *HssClient) ListRiskConfigHostsInvoker(request *model.ListRiskConfigHostsRequest) *ListRiskConfigHostsInvoker {
	requestDef := GenReqDefForListRiskConfigHosts()
	return &ListRiskConfigHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListRiskConfigs(request *model.ListRiskConfigsRequest) (*model.ListRiskConfigsResponse, error) {
	requestDef := GenReqDefForListRiskConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskConfigsResponse), nil
	}
}

func (c *HssClient) ListRiskConfigsInvoker(request *model.ListRiskConfigsRequest) *ListRiskConfigsInvoker {
	requestDef := GenReqDefForListRiskConfigs()
	return &ListRiskConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListSecurityEvents(request *model.ListSecurityEventsRequest) (*model.ListSecurityEventsResponse, error) {
	requestDef := GenReqDefForListSecurityEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityEventsResponse), nil
	}
}

func (c *HssClient) ListSecurityEventsInvoker(request *model.ListSecurityEventsRequest) *ListSecurityEventsInvoker {
	requestDef := GenReqDefForListSecurityEvents()
	return &ListSecurityEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListUserChangeHistories(request *model.ListUserChangeHistoriesRequest) (*model.ListUserChangeHistoriesResponse, error) {
	requestDef := GenReqDefForListUserChangeHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserChangeHistoriesResponse), nil
	}
}

func (c *HssClient) ListUserChangeHistoriesInvoker(request *model.ListUserChangeHistoriesRequest) *ListUserChangeHistoriesInvoker {
	requestDef := GenReqDefForListUserChangeHistories()
	return &ListUserChangeHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListUserStatistics(request *model.ListUserStatisticsRequest) (*model.ListUserStatisticsResponse, error) {
	requestDef := GenReqDefForListUserStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserStatisticsResponse), nil
	}
}

func (c *HssClient) ListUserStatisticsInvoker(request *model.ListUserStatisticsRequest) *ListUserStatisticsInvoker {
	requestDef := GenReqDefForListUserStatistics()
	return &ListUserStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

func (c *HssClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListVulHosts(request *model.ListVulHostsRequest) (*model.ListVulHostsResponse, error) {
	requestDef := GenReqDefForListVulHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulHostsResponse), nil
	}
}

func (c *HssClient) ListVulHostsInvoker(request *model.ListVulHostsRequest) *ListVulHostsInvoker {
	requestDef := GenReqDefForListVulHosts()
	return &ListVulHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListVulnerabilities(request *model.ListVulnerabilitiesRequest) (*model.ListVulnerabilitiesResponse, error) {
	requestDef := GenReqDefForListVulnerabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulnerabilitiesResponse), nil
	}
}

func (c *HssClient) ListVulnerabilitiesInvoker(request *model.ListVulnerabilitiesRequest) *ListVulnerabilitiesInvoker {
	requestDef := GenReqDefForListVulnerabilities()
	return &ListVulnerabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListWeakPasswordUsers(request *model.ListWeakPasswordUsersRequest) (*model.ListWeakPasswordUsersResponse, error) {
	requestDef := GenReqDefForListWeakPasswordUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWeakPasswordUsersResponse), nil
	}
}

func (c *HssClient) ListWeakPasswordUsersInvoker(request *model.ListWeakPasswordUsersRequest) *ListWeakPasswordUsersInvoker {
	requestDef := GenReqDefForListWeakPasswordUsers()
	return &ListWeakPasswordUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ListWtpProtectHost(request *model.ListWtpProtectHostRequest) (*model.ListWtpProtectHostResponse, error) {
	requestDef := GenReqDefForListWtpProtectHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWtpProtectHostResponse), nil
	}
}

func (c *HssClient) ListWtpProtectHostInvoker(request *model.ListWtpProtectHostRequest) *ListWtpProtectHostInvoker {
	requestDef := GenReqDefForListWtpProtectHost()
	return &ListWtpProtectHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) SetRaspSwitch(request *model.SetRaspSwitchRequest) (*model.SetRaspSwitchResponse, error) {
	requestDef := GenReqDefForSetRaspSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRaspSwitchResponse), nil
	}
}

func (c *HssClient) SetRaspSwitchInvoker(request *model.SetRaspSwitchRequest) *SetRaspSwitchInvoker {
	requestDef := GenReqDefForSetRaspSwitch()
	return &SetRaspSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) SetWtpProtectionStatusInfo(request *model.SetWtpProtectionStatusInfoRequest) (*model.SetWtpProtectionStatusInfoResponse, error) {
	requestDef := GenReqDefForSetWtpProtectionStatusInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetWtpProtectionStatusInfoResponse), nil
	}
}

func (c *HssClient) SetWtpProtectionStatusInfoInvoker(request *model.SetWtpProtectionStatusInfoRequest) *SetWtpProtectionStatusInfoInvoker {
	requestDef := GenReqDefForSetWtpProtectionStatusInfo()
	return &SetWtpProtectionStatusInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ShowAssetStatistic(request *model.ShowAssetStatisticRequest) (*model.ShowAssetStatisticResponse, error) {
	requestDef := GenReqDefForShowAssetStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetStatisticResponse), nil
	}
}

func (c *HssClient) ShowAssetStatisticInvoker(request *model.ShowAssetStatisticRequest) *ShowAssetStatisticInvoker {
	requestDef := GenReqDefForShowAssetStatistic()
	return &ShowAssetStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ShowBackupPolicyInfo(request *model.ShowBackupPolicyInfoRequest) (*model.ShowBackupPolicyInfoResponse, error) {
	requestDef := GenReqDefForShowBackupPolicyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupPolicyInfoResponse), nil
	}
}

func (c *HssClient) ShowBackupPolicyInfoInvoker(request *model.ShowBackupPolicyInfoRequest) *ShowBackupPolicyInfoInvoker {
	requestDef := GenReqDefForShowBackupPolicyInfo()
	return &ShowBackupPolicyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ShowCheckRuleDetail(request *model.ShowCheckRuleDetailRequest) (*model.ShowCheckRuleDetailResponse, error) {
	requestDef := GenReqDefForShowCheckRuleDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckRuleDetailResponse), nil
	}
}

func (c *HssClient) ShowCheckRuleDetailInvoker(request *model.ShowCheckRuleDetailRequest) *ShowCheckRuleDetailInvoker {
	requestDef := GenReqDefForShowCheckRuleDetail()
	return &ShowCheckRuleDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) ShowRiskConfigDetail(request *model.ShowRiskConfigDetailRequest) (*model.ShowRiskConfigDetailResponse, error) {
	requestDef := GenReqDefForShowRiskConfigDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRiskConfigDetailResponse), nil
	}
}

func (c *HssClient) ShowRiskConfigDetailInvoker(request *model.ShowRiskConfigDetailRequest) *ShowRiskConfigDetailInvoker {
	requestDef := GenReqDefForShowRiskConfigDetail()
	return &ShowRiskConfigDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) StartProtection(request *model.StartProtectionRequest) (*model.StartProtectionResponse, error) {
	requestDef := GenReqDefForStartProtection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartProtectionResponse), nil
	}
}

func (c *HssClient) StartProtectionInvoker(request *model.StartProtectionRequest) *StartProtectionInvoker {
	requestDef := GenReqDefForStartProtection()
	return &StartProtectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) StopProtection(request *model.StopProtectionRequest) (*model.StopProtectionResponse, error) {
	requestDef := GenReqDefForStopProtection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopProtectionResponse), nil
	}
}

func (c *HssClient) StopProtectionInvoker(request *model.StopProtectionRequest) *StopProtectionInvoker {
	requestDef := GenReqDefForStopProtection()
	return &StopProtectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) SwitchHostsProtectStatus(request *model.SwitchHostsProtectStatusRequest) (*model.SwitchHostsProtectStatusResponse, error) {
	requestDef := GenReqDefForSwitchHostsProtectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchHostsProtectStatusResponse), nil
	}
}

func (c *HssClient) SwitchHostsProtectStatusInvoker(request *model.SwitchHostsProtectStatusRequest) *SwitchHostsProtectStatusInvoker {
	requestDef := GenReqDefForSwitchHostsProtectStatus()
	return &SwitchHostsProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) UpdateBackupPolicyInfo(request *model.UpdateBackupPolicyInfoRequest) (*model.UpdateBackupPolicyInfoResponse, error) {
	requestDef := GenReqDefForUpdateBackupPolicyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBackupPolicyInfoResponse), nil
	}
}

func (c *HssClient) UpdateBackupPolicyInfoInvoker(request *model.UpdateBackupPolicyInfoRequest) *UpdateBackupPolicyInfoInvoker {
	requestDef := GenReqDefForUpdateBackupPolicyInfo()
	return &UpdateBackupPolicyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *HssClient) UpdateProtectionPolicy(request *model.UpdateProtectionPolicyRequest) (*model.UpdateProtectionPolicyResponse, error) {
	requestDef := GenReqDefForUpdateProtectionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProtectionPolicyResponse), nil
	}
}

func (c *HssClient) UpdateProtectionPolicyInvoker(request *model.UpdateProtectionPolicyRequest) *UpdateProtectionPolicyInvoker {
	requestDef := GenReqDefForUpdateProtectionPolicy()
	return &UpdateProtectionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
