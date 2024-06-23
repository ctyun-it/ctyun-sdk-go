package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// InstanceBackupPolicyBindInstancesApi
type InstanceBackupPolicyBindInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewInstanceBackupPolicyBindInstancesApi(client *ctyunsdk.CtyunClient) *InstanceBackupPolicyBindInstancesApi {
	return &InstanceBackupPolicyBindInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/bind-instances",
		},
	}
}

func (this *InstanceBackupPolicyBindInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *InstanceBackupPolicyBindInstancesRequest) (*InstanceBackupPolicyBindInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&InstanceBackupPolicyBindInstancesRealRequest{
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

	var realResponse InstanceBackupPolicyBindInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &InstanceBackupPolicyBindInstancesResponse{
		PolicyID:       realResponse.PolicyID,
		InstanceIDList: realResponse.InstanceIDList,
	}, nil
}

type InstanceBackupPolicyBindInstancesRealRequest struct {
	RegionID       string `json:"regionID,omitempty"`
	PolicyID       string `json:"policyID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type InstanceBackupPolicyBindInstancesRequest struct {
	RegionID       string
	PolicyID       string
	InstanceIDList string
}

type InstanceBackupPolicyBindInstancesRealResponse struct {
	PolicyID       string `json:"policyID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type InstanceBackupPolicyBindInstancesResponse struct {
	PolicyID       string
	InstanceIDList string
}
