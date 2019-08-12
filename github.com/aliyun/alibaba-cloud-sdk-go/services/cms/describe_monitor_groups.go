package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeMonitorGroups invokes the cms.DescribeMonitorGroups API synchronously
// api document: https://help.aliyun.com/api/cms/describemonitorgroups.html
func (client *Client) DescribeMonitorGroups(request *DescribeMonitorGroupsRequest) (response *DescribeMonitorGroupsResponse, err error) {
	response = CreateDescribeMonitorGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonitorGroupsWithChan invokes the cms.DescribeMonitorGroups API asynchronously
// api document: https://help.aliyun.com/api/cms/describemonitorgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonitorGroupsWithChan(request *DescribeMonitorGroupsRequest) (<-chan *DescribeMonitorGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeMonitorGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonitorGroups(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeMonitorGroupsWithCallback invokes the cms.DescribeMonitorGroups API asynchronously
// api document: https://help.aliyun.com/api/cms/describemonitorgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonitorGroupsWithCallback(request *DescribeMonitorGroupsRequest, callback func(response *DescribeMonitorGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonitorGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonitorGroups(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeMonitorGroupsRequest is the request struct for api DescribeMonitorGroups
type DescribeMonitorGroupsRequest struct {
	*requests.RpcRequest
	SelectContactGroups requests.Boolean `position:"Query" name:"SelectContactGroups"`
	InstanceId          string           `position:"Query" name:"InstanceId"`
	PageSize            requests.Integer `position:"Query" name:"PageSize"`
	Keyword             string           `position:"Query" name:"Keyword"`
	GroupName           string           `position:"Query" name:"GroupName"`
	PageNumber          requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeMonitorGroupsResponse is the response struct for api DescribeMonitorGroups
type DescribeMonitorGroupsResponse struct {
	*responses.BaseResponse
	RequestId  string                           `json:"RequestId" xml:"RequestId"`
	Success    bool                             `json:"Success" xml:"Success"`
	Code       int                              `json:"Code" xml:"Code"`
	Message    string                           `json:"Message" xml:"Message"`
	PageNumber int                              `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                              `json:"PageSize" xml:"PageSize"`
	Total      int                              `json:"Total" xml:"Total"`
	Resources  ResourcesInDescribeMonitorGroups `json:"Resources" xml:"Resources"`
}

// CreateDescribeMonitorGroupsRequest creates a request to invoke DescribeMonitorGroups API
func CreateDescribeMonitorGroupsRequest() (request *DescribeMonitorGroupsRequest) {
	request = &DescribeMonitorGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMonitorGroups", "cms", "openAPI")
	return
}

// CreateDescribeMonitorGroupsResponse creates a response to parse from DescribeMonitorGroups response
func CreateDescribeMonitorGroupsResponse() (response *DescribeMonitorGroupsResponse) {
	response = &DescribeMonitorGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
