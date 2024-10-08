// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConfigNetStatusRequest struct {
	// The CIDR block of the anti-DDoS diversion instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX/22
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// The regions in which the CIDR block needs to be advertised or withdrawn from advertising. If you leave this parameter empty, the CIDR blocks in all regions are configured.
	//
	// >  You can call the [QueryNetList](https://help.aliyun.com/document_detail/2639086.html) operation to obtain the regions of the CIDR blocks.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the anti-DDoS diversion instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// The status of the CIDR block. Valid values:
	//
	// 	- enable: advertises the CIDR block.
	//
	// 	- disable: withdraws the advertising of the CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subnet CIDR blocks of the CIDR block.
	SubNets []*string `json:"SubNets,omitempty" xml:"SubNets,omitempty" type:"Repeated"`
}

func (s ConfigNetStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigNetStatusRequest) GoString() string {
	return s.String()
}

func (s *ConfigNetStatusRequest) SetNet(v string) *ConfigNetStatusRequest {
	s.Net = &v
	return s
}

func (s *ConfigNetStatusRequest) SetRegions(v []*string) *ConfigNetStatusRequest {
	s.Regions = v
	return s
}

func (s *ConfigNetStatusRequest) SetSaleId(v string) *ConfigNetStatusRequest {
	s.SaleId = &v
	return s
}

func (s *ConfigNetStatusRequest) SetStatus(v string) *ConfigNetStatusRequest {
	s.Status = &v
	return s
}

func (s *ConfigNetStatusRequest) SetSubNets(v []*string) *ConfigNetStatusRequest {
	s.SubNets = v
	return s
}

type ConfigNetStatusResponseBody struct {
	// The status code.
	//
	// 	- **200**: The request was successful.
	//
	// 	- Other codes: The request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B0949F09-B9C1-1D5E-8F27-0A5BF3CD5D95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigNetStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigNetStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigNetStatusResponseBody) SetCode(v int64) *ConfigNetStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigNetStatusResponseBody) SetMessage(v string) *ConfigNetStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigNetStatusResponseBody) SetRequestId(v string) *ConfigNetStatusResponseBody {
	s.RequestId = &v
	return s
}

type ConfigNetStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigNetStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigNetStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigNetStatusResponse) GoString() string {
	return s.String()
}

func (s *ConfigNetStatusResponse) SetHeaders(v map[string]*string) *ConfigNetStatusResponse {
	s.Headers = v
	return s
}

func (s *ConfigNetStatusResponse) SetStatusCode(v int32) *ConfigNetStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigNetStatusResponse) SetBody(v *ConfigNetStatusResponseBody) *ConfigNetStatusResponse {
	s.Body = v
	return s
}

type ListInstanceRequest struct {
	// The name of the instance.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The ID of the anti-DDoS diversion instance.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// The status of the instance. Valid values:
	//
	// - normal
	//
	// - expired
	//
	// - deleting
	//
	// - stopped
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) SetName(v string) *ListInstanceRequest {
	s.Name = &v
	return s
}

func (s *ListInstanceRequest) SetNum(v int64) *ListInstanceRequest {
	s.Num = &v
	return s
}

func (s *ListInstanceRequest) SetPage(v int64) *ListInstanceRequest {
	s.Page = &v
	return s
}

func (s *ListInstanceRequest) SetSaleId(v string) *ListInstanceRequest {
	s.SaleId = &v
	return s
}

func (s *ListInstanceRequest) SetStatus(v string) *ListInstanceRequest {
	s.Status = &v
	return s
}

type ListInstanceResponseBody struct {
	// The status code.
	//
	// - 200: The request was successful.
	//
	// - Other codes: The request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *ListInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response parameters.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B0949F09-B9C1-1D5E-8F27-0A5BF3CD5D95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) SetCode(v int64) *ListInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceResponseBody) SetData(v *ListInstanceResponseBodyData) *ListInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceResponseBody) SetMessage(v string) *ListInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ListInstanceResponseBodyData struct {
	// The details of the anti-DDoS diversion instance.
	Instances []*ListInstanceResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyData) SetInstances(v []*ListInstanceResponseBodyDataInstances) *ListInstanceResponseBodyData {
	s.Instances = v
	return s
}

func (s *ListInstanceResponseBodyData) SetNum(v int64) *ListInstanceResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetPage(v int64) *ListInstanceResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListInstanceResponseBodyData) SetTotal(v int64) *ListInstanceResponseBodyData {
	s.Total = &v
	return s
}

type ListInstanceResponseBodyDataInstances struct {
	// The description.
	//
	// example:
	//
	// description
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The purchase time.
	//
	// example:
	//
	// 2022-12-15 11:10:42
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2023-02-23 00:00:00
	GmtExpire *string `json:"GmtExpire,omitempty" xml:"GmtExpire,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2022-12-15 11:10:42
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// The alias of the instance.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx_xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configurations of the instance.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// The specifications of the instance.
	Spec *ListInstanceResponseBodyDataInstancesSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	// The status of the instance. Valid values:
	//
	// - normal
	//
	// - expired
	//
	// - deleting
	//
	// - stopped
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 177xxxxxxxxxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListInstanceResponseBodyDataInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyDataInstances) SetComment(v string) *ListInstanceResponseBodyDataInstances {
	s.Comment = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetGmtCreate(v string) *ListInstanceResponseBodyDataInstances {
	s.GmtCreate = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetGmtExpire(v string) *ListInstanceResponseBodyDataInstances {
	s.GmtExpire = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetGmtModify(v string) *ListInstanceResponseBodyDataInstances {
	s.GmtModify = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetInstanceId(v string) *ListInstanceResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetMessage(v string) *ListInstanceResponseBodyDataInstances {
	s.Message = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetName(v string) *ListInstanceResponseBodyDataInstances {
	s.Name = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetSaleId(v string) *ListInstanceResponseBodyDataInstances {
	s.SaleId = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetSpec(v *ListInstanceResponseBodyDataInstancesSpec) *ListInstanceResponseBodyDataInstances {
	s.Spec = v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetStatus(v string) *ListInstanceResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetUserId(v string) *ListInstanceResponseBodyDataInstances {
	s.UserId = &v
	return s
}

type ListInstanceResponseBodyDataInstancesSpec struct {
	// The region of the asset.
	//
	// example:
	//
	// international_and_hmt
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The diversion mode. Valid values: on-demand always-on
	//
	// example:
	//
	// on-demand
	DiversionType *string `json:"DiversionType,omitempty" xml:"DiversionType,omitempty"`
	// The mitigation plan.
	//
	// example:
	//
	// enterprise
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The number of data centers. Valid values: 1 to 10.
	//
	// example:
	//
	// 1
	IdcNumbers *string `json:"IdcNumbers,omitempty" xml:"IdcNumbers,omitempty"`
	// The initial installation mode.
	//
	// example:
	//
	// gre_tunnel_by_pccw
	InitialInstallation *string `json:"InitialInstallation,omitempty" xml:"InitialInstallation,omitempty"`
	// The initial installation quantity.
	//
	// example:
	//
	// 1
	InitialQty *string `json:"InitialQty,omitempty" xml:"InitialQty,omitempty"`
	// The number of CIDR blocks. Value range: 1 to 10000.
	//
	// example:
	//
	// 1
	IpSubnetNums *string `json:"IpSubnetNums,omitempty" xml:"IpSubnetNums,omitempty"`
	// The mitigation analysis feature.
	//
	// example:
	//
	// off
	MitigationAnalysis *string `json:"MitigationAnalysis,omitempty" xml:"MitigationAnalysis,omitempty"`
	// The log storage capacity.
	//
	// example:
	//
	// 3T
	MitigationAnalysisCapacity *string `json:"MitigationAnalysisCapacity,omitempty" xml:"MitigationAnalysisCapacity,omitempty"`
	// The maximum mitigation capability.
	//
	// example:
	//
	// unlimited
	MitigationCapacity *string `json:"MitigationCapacity,omitempty" xml:"MitigationCapacity,omitempty"`
	// The number of mitigation sessions.
	//
	// example:
	//
	// unlimited
	MitigationNums *string `json:"MitigationNums,omitempty" xml:"MitigationNums,omitempty"`
	// The service traffic. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	NormalBandwidth *string `json:"NormalBandwidth,omitempty" xml:"NormalBandwidth,omitempty"`
}

func (s ListInstanceResponseBodyDataInstancesSpec) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyDataInstancesSpec) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetCoverage(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.Coverage = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetDiversionType(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.DiversionType = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetEdition(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.Edition = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetIdcNumbers(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.IdcNumbers = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetInitialInstallation(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.InitialInstallation = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetInitialQty(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.InitialQty = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetIpSubnetNums(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.IpSubnetNums = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetMitigationAnalysis(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.MitigationAnalysis = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetMitigationAnalysisCapacity(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.MitigationAnalysisCapacity = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetMitigationCapacity(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.MitigationCapacity = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetMitigationNums(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.MitigationNums = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesSpec) SetNormalBandwidth(v string) *ListInstanceResponseBodyDataInstancesSpec {
	s.NormalBandwidth = &v
	return s
}

type ListInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResponse) SetHeaders(v map[string]*string) *ListInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResponse) SetStatusCode(v int32) *ListInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResponse) SetBody(v *ListInstanceResponseBody) *ListInstanceResponse {
	s.Body = v
	return s
}

type QueryNetListRequest struct {
	// The primary CIDR block of the anti-DDoS diversion instance for which an extended CIDR block is configured. If no extended CIDR blocks are configured for the anti-DDoS diversion instance, leave this parameter empty.
	//
	// example:
	//
	// 192.168.XX.XX/22
	MainNet *string `json:"MainNet,omitempty" xml:"MainNet,omitempty"`
	// The scheduling mode. Valid values:
	//
	// - manual: manual scheduling
	//
	// - netflow-auto: automatic scheduling
	//
	// example:
	//
	// netflow-auto
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The CIDR block of the anti-DDoS diversion instance.
	//
	//
	// > If no extended CIDR blocks are configured for the anti-DDoS diversion instance, this parameter specifies the CIDR block of the instance. If an extended CIDR block is configured for the anti-DDoS diversion instance, this parameter specifies the extended CIDR block that is configured for the instance. If this parameter is specified, the MainNet parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX/24
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The ID of the anti-DDoS diversion instance.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
}

func (s QueryNetListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListRequest) GoString() string {
	return s.String()
}

func (s *QueryNetListRequest) SetMainNet(v string) *QueryNetListRequest {
	s.MainNet = &v
	return s
}

func (s *QueryNetListRequest) SetMode(v string) *QueryNetListRequest {
	s.Mode = &v
	return s
}

func (s *QueryNetListRequest) SetNet(v string) *QueryNetListRequest {
	s.Net = &v
	return s
}

func (s *QueryNetListRequest) SetNum(v int64) *QueryNetListRequest {
	s.Num = &v
	return s
}

func (s *QueryNetListRequest) SetPage(v int64) *QueryNetListRequest {
	s.Page = &v
	return s
}

func (s *QueryNetListRequest) SetSaleId(v string) *QueryNetListRequest {
	s.SaleId = &v
	return s
}

type QueryNetListResponseBody struct {
	// The status code.
	//
	// - 200: The request was successful.
	//
	// - Other codes: The request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The CIDR blocks.
	Data *QueryNetListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response parameters.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 24B652B5-AEFF-3F03-9114-00D053C42277
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryNetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBody) SetCode(v int64) *QueryNetListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryNetListResponseBody) SetData(v *QueryNetListResponseBodyData) *QueryNetListResponseBody {
	s.Data = v
	return s
}

func (s *QueryNetListResponseBody) SetMessage(v string) *QueryNetListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryNetListResponseBody) SetRequestId(v string) *QueryNetListResponseBody {
	s.RequestId = &v
	return s
}

type QueryNetListResponseBodyData struct {
	// The configuration of the CIDR block.
	Nets []*QueryNetListResponseBodyDataNets `json:"Nets,omitempty" xml:"Nets,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryNetListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyData) SetNets(v []*QueryNetListResponseBodyDataNets) *QueryNetListResponseBodyData {
	s.Nets = v
	return s
}

func (s *QueryNetListResponseBodyData) SetNum(v int64) *QueryNetListResponseBodyData {
	s.Num = &v
	return s
}

func (s *QueryNetListResponseBodyData) SetPage(v int64) *QueryNetListResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryNetListResponseBodyData) SetTotal(v int64) *QueryNetListResponseBodyData {
	s.Total = &v
	return s
}

type QueryNetListResponseBodyDataNets struct {
	// The DDoS mitigation configuration of the CIDR block.
	DDoSDefense *QueryNetListResponseBodyDataNetsDDoSDefense `json:"DDoSDefense,omitempty" xml:"DDoSDefense,omitempty" type:"Struct"`
	// The advertising details.
	Declared []*QueryNetListResponseBodyDataNetsDeclared `json:"Declared,omitempty" xml:"Declared,omitempty" type:"Repeated"`
	// The advertising status of the CIDR block. Valid values:
	//
	// - 0: The CIDR block is not advertised.
	//
	// - 1: The CIDR block is advertised.
	//
	// example:
	//
	// 1
	DeclaredState *int32 `json:"DeclaredState,omitempty" xml:"DeclaredState,omitempty"`
	// Indicates whether the forwarding configuration takes effect. Valid values:
	//
	// - 0: The forwarding configuration takes effect.
	//
	// - 1: The forwarding configuration does not take effect.
	//
	// - -1: The forwarding configuration is being deleted.
	//
	// example:
	//
	// 1
	FwdEffect *int64 `json:"FwdEffect,omitempty" xml:"FwdEffect,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2023-02-23 00:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2023-02-24 00:00:00
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// The scheduling mode.
	//
	// example:
	//
	// manual
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The CIDR block of the anti-DDoS diversion instance.
	//
	// example:
	//
	// 192.168.XX.XX/24
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// Indicates whether the CIDR block needs to be extended. Valid values:
	//
	// - 0: The CIDR block needs to be extended.
	//
	// - 1: The CIDR block does not need to be extended.
	//
	// example:
	//
	// 1
	NetExtend *string `json:"NetExtend,omitempty" xml:"NetExtend,omitempty"`
	// The primary CIDR block.
	//
	// example:
	//
	// 192.168.XX.XX/22
	NetMain *string `json:"NetMain,omitempty" xml:"NetMain,omitempty"`
	// The type of the CIDR block.
	//
	// example:
	//
	// ipv4
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Permit  *int32  `json:"Permit,omitempty" xml:"Permit,omitempty"`
	// The ID of the anti-DDoS diversion instance.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// The reinjection type.
	//
	// example:
	//
	// aliyun_line
	UpstreamType *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 177xxxxxxxxxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryNetListResponseBodyDataNets) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListResponseBodyDataNets) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNets) SetDDoSDefense(v *QueryNetListResponseBodyDataNetsDDoSDefense) *QueryNetListResponseBodyDataNets {
	s.DDoSDefense = v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetDeclared(v []*QueryNetListResponseBodyDataNetsDeclared) *QueryNetListResponseBodyDataNets {
	s.Declared = v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetDeclaredState(v int32) *QueryNetListResponseBodyDataNets {
	s.DeclaredState = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetFwdEffect(v int64) *QueryNetListResponseBodyDataNets {
	s.FwdEffect = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetGmtCreate(v string) *QueryNetListResponseBodyDataNets {
	s.GmtCreate = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetGmtModify(v string) *QueryNetListResponseBodyDataNets {
	s.GmtModify = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetMode(v string) *QueryNetListResponseBodyDataNets {
	s.Mode = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetNet(v string) *QueryNetListResponseBodyDataNets {
	s.Net = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetNetExtend(v string) *QueryNetListResponseBodyDataNets {
	s.NetExtend = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetNetMain(v string) *QueryNetListResponseBodyDataNets {
	s.NetMain = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetNetType(v string) *QueryNetListResponseBodyDataNets {
	s.NetType = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetPermit(v int32) *QueryNetListResponseBodyDataNets {
	s.Permit = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetSaleId(v string) *QueryNetListResponseBodyDataNets {
	s.SaleId = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetUpstreamType(v string) *QueryNetListResponseBodyDataNets {
	s.UpstreamType = &v
	return s
}

func (s *QueryNetListResponseBodyDataNets) SetUserId(v string) *QueryNetListResponseBodyDataNets {
	s.UserId = &v
	return s
}

type QueryNetListResponseBodyDataNetsDDoSDefense struct {
	// The configuration of traffic scrubbing.
	CleanTh *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh `json:"CleanTh,omitempty" xml:"CleanTh,omitempty" type:"Struct"`
	// The configuration of the mitigation policy.
	DjPolicy *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy `json:"DjPolicy,omitempty" xml:"DjPolicy,omitempty" type:"Struct"`
	// The configuration of blackhole filtering.
	HoleTh *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh `json:"HoleTh,omitempty" xml:"HoleTh,omitempty" type:"Struct"`
}

func (s QueryNetListResponseBodyDataNetsDDoSDefense) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDDoSDefense) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) SetCleanTh(v *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) *QueryNetListResponseBodyDataNetsDDoSDefense {
	s.CleanTh = v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) SetDjPolicy(v *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) *QueryNetListResponseBodyDataNetsDDoSDefense {
	s.DjPolicy = v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefense) SetHoleTh(v *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) *QueryNetListResponseBodyDataNetsDDoSDefense {
	s.HoleTh = v
	return s
}

type QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh struct {
	// The traffic scrubbing threshold in Mbit/s.
	//
	// example:
	//
	// 0
	Mbps *int32 `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	// The traffic scrubbing threshold in packets per second (pps)
	//
	// example:
	//
	// 0
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) SetMbps(v int32) *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh {
	s.Mbps = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh) SetPps(v int32) *QueryNetListResponseBodyDataNetsDDoSDefenseCleanTh {
	s.Pps = &v
	return s
}

type QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy struct {
	// The name of the mitigation policy.
	//
	// example:
	//
	// test_polilciy-xxx
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy) SetPolicyName(v string) *QueryNetListResponseBodyDataNetsDDoSDefenseDjPolicy {
	s.PolicyName = &v
	return s
}

type QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh struct {
	// The blackhole filtering threshold.
	//
	// example:
	//
	// 0
	ThreshMbps *int32 `json:"ThreshMbps,omitempty" xml:"ThreshMbps,omitempty"`
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh) SetThreshMbps(v int32) *QueryNetListResponseBodyDataNetsDDoSDefenseHoleTh {
	s.ThreshMbps = &v
	return s
}

type QueryNetListResponseBodyDataNetsDeclared struct {
	// Indicates whether the CIDR block is advertised. Valid values:
	//
	// - 0: The CIDR block is not advertised.
	//
	// - 1: The CIDR block is advertised.
	//
	// example:
	//
	// 0
	Declared *string `json:"Declared,omitempty" xml:"Declared,omitempty"`
	// The region in which the CIDR block is advertised.
	//
	// example:
	//
	// oe26
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryNetListResponseBodyDataNetsDeclared) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListResponseBodyDataNetsDeclared) GoString() string {
	return s.String()
}

func (s *QueryNetListResponseBodyDataNetsDeclared) SetDeclared(v string) *QueryNetListResponseBodyDataNetsDeclared {
	s.Declared = &v
	return s
}

func (s *QueryNetListResponseBodyDataNetsDeclared) SetRegion(v string) *QueryNetListResponseBodyDataNetsDeclared {
	s.Region = &v
	return s
}

type QueryNetListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryNetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryNetListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryNetListResponse) GoString() string {
	return s.String()
}

func (s *QueryNetListResponse) SetHeaders(v map[string]*string) *QueryNetListResponse {
	s.Headers = v
	return s
}

func (s *QueryNetListResponse) SetStatusCode(v int32) *QueryNetListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNetListResponse) SetBody(v *QueryNetListResponseBody) *QueryNetListResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddosdiversion"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the advertising of a CIDR block.
//
// @param request - ConfigNetStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigNetStatusResponse
func (client *Client) ConfigNetStatusWithOptions(request *ConfigNetStatusRequest, runtime *util.RuntimeOptions) (_result *ConfigNetStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Net)) {
		query["Net"] = request.Net
	}

	if !tea.BoolValue(util.IsUnset(request.Regions)) {
		query["Regions"] = request.Regions
	}

	if !tea.BoolValue(util.IsUnset(request.SaleId)) {
		query["SaleId"] = request.SaleId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SubNets)) {
		query["SubNets"] = request.SubNets
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigNetStatus"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigNetStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the advertising of a CIDR block.
//
// @param request - ConfigNetStatusRequest
//
// @return ConfigNetStatusResponse
func (client *Client) ConfigNetStatus(request *ConfigNetStatusRequest) (_result *ConfigNetStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigNetStatusResponse{}
	_body, _err := client.ConfigNetStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries anti-DDoS diversion instances.
//
// @param request - ListInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, runtime *util.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.SaleId)) {
		query["SaleId"] = request.SaleId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstance"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries anti-DDoS diversion instances.
//
// @param request - ListInstanceRequest
//
// @return ListInstanceResponse
func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the CIDR blocks of an anti-DDoS diversion instance.
//
// @param request - QueryNetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryNetListResponse
func (client *Client) QueryNetListWithOptions(request *QueryNetListRequest, runtime *util.RuntimeOptions) (_result *QueryNetListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MainNet)) {
		query["MainNet"] = request.MainNet
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.Net)) {
		query["Net"] = request.Net
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		query["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.SaleId)) {
		query["SaleId"] = request.SaleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryNetList"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryNetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the CIDR blocks of an anti-DDoS diversion instance.
//
// @param request - QueryNetListRequest
//
// @return QueryNetListResponse
func (client *Client) QueryNetList(request *QueryNetListRequest) (_result *QueryNetListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryNetListResponse{}
	_body, _err := client.QueryNetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
