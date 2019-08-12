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

// ApplyMetricRuleTemplate invokes the cms.ApplyMetricRuleTemplate API synchronously
// api document: https://help.aliyun.com/api/cms/applymetricruletemplate.html
func (client *Client) ApplyMetricRuleTemplate(request *ApplyMetricRuleTemplateRequest) (response *ApplyMetricRuleTemplateResponse, err error) {
	response = CreateApplyMetricRuleTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ApplyMetricRuleTemplateWithChan invokes the cms.ApplyMetricRuleTemplate API asynchronously
// api document: https://help.aliyun.com/api/cms/applymetricruletemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyMetricRuleTemplateWithChan(request *ApplyMetricRuleTemplateRequest) (<-chan *ApplyMetricRuleTemplateResponse, <-chan error) {
	responseChan := make(chan *ApplyMetricRuleTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyMetricRuleTemplate(request)
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

// ApplyMetricRuleTemplateWithCallback invokes the cms.ApplyMetricRuleTemplate API asynchronously
// api document: https://help.aliyun.com/api/cms/applymetricruletemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyMetricRuleTemplateWithCallback(request *ApplyMetricRuleTemplateRequest, callback func(response *ApplyMetricRuleTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyMetricRuleTemplateResponse
		var err error
		defer close(result)
		response, err = client.ApplyMetricRuleTemplate(request)
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

// ApplyMetricRuleTemplateRequest is the request struct for api ApplyMetricRuleTemplate
type ApplyMetricRuleTemplateRequest struct {
	*requests.RpcRequest
	EnableStartTime requests.Integer `position:"Query" name:"EnableStartTime"`
	ApplyMode       string           `position:"Query" name:"ApplyMode"`
	Webhook         string           `position:"Query" name:"Webhook"`
	TemplateIds     string           `position:"Query" name:"TemplateIds"`
	EnableEndTime   requests.Integer `position:"Query" name:"EnableEndTime"`
	GroupId         requests.Integer `position:"Query" name:"GroupId"`
	NotifyLevel     requests.Integer `position:"Query" name:"NotifyLevel"`
	SilenceTime     requests.Integer `position:"Query" name:"SilenceTime"`
}

// ApplyMetricRuleTemplateResponse is the response struct for api ApplyMetricRuleTemplate
type ApplyMetricRuleTemplateResponse struct {
	*responses.BaseResponse
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	Success   bool                              `json:"Success" xml:"Success"`
	Code      int                               `json:"Code" xml:"Code"`
	Message   string                            `json:"Message" xml:"Message"`
	Resource  ResourceInApplyMetricRuleTemplate `json:"Resource" xml:"Resource"`
}

// CreateApplyMetricRuleTemplateRequest creates a request to invoke ApplyMetricRuleTemplate API
func CreateApplyMetricRuleTemplateRequest() (request *ApplyMetricRuleTemplateRequest) {
	request = &ApplyMetricRuleTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "ApplyMetricRuleTemplate", "cms", "openAPI")
	return
}

// CreateApplyMetricRuleTemplateResponse creates a response to parse from ApplyMetricRuleTemplate response
func CreateApplyMetricRuleTemplateResponse() (response *ApplyMetricRuleTemplateResponse) {
	response = &ApplyMetricRuleTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
