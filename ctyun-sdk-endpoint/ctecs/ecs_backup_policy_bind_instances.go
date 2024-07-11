package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupPolicyBindInstancesApi
type EcsBackupPolicyBindInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupPolicyBindInstancesApi(client *ctyunsdk.CtyunClient) *EcsBackupPolicyBindInstancesApi {
	return &EcsBackupPolicyBindInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/bind-instances",
		},
	}
}

func (this *EcsBackupPolicyBindInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupPolicyBindInstancesRequest) (*EcsBackupPolicyBindInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupPolicyBindInstancesRealRequest{
		RegionID:       req.RegionID,
		PolicyID:       req.PolicyID,
		InstanceIDList: req.InstanceIDList,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupPolicyBindInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupPolicyBindInstancesResponse{
		PolicyID:       realResponse.PolicyID,
		InstanceIDList: realResponse.InstanceIDList,
	}, nil
}

type EcsBackupPolicyBindInstancesRealRequest struct {
	RegionID       string `json:"regionID,omitempty"`
	PolicyID       string `json:"policyID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type EcsBackupPolicyBindInstancesRequest struct {
	RegionID       string
	PolicyID       string
	InstanceIDList string
}

type EcsBackupPolicyBindInstancesRealResponse struct {
	PolicyID       string `json:"policyID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type EcsBackupPolicyBindInstancesResponse struct {
	PolicyID       string
	InstanceIDList string
}
