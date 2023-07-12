package gpdb

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

// PlanList is a nested struct in gpdb response
type PlanList struct {
	PlanName         string `json:"PlanName" xml:"PlanName"`
	PlanDesc         string `json:"PlanDesc" xml:"PlanDesc"`
	PlanId           string `json:"PlanId" xml:"PlanId"`
	PlanStatus       string `json:"PlanStatus" xml:"PlanStatus"`
	PlanType         string `json:"PlanType" xml:"PlanType"`
	PlanScheduleType string `json:"PlanScheduleType" xml:"PlanScheduleType"`
	PlanStartDate    string `json:"PlanStartDate" xml:"PlanStartDate"`
	PlanEndDate      string `json:"PlanEndDate" xml:"PlanEndDate"`
	PlanConfig       string `json:"PlanConfig" xml:"PlanConfig"`
	DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
}