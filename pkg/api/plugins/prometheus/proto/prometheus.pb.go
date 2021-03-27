// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: prometheus.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// GetVariablesRequest is the request data needed to get all variables for the metrics view of an application. The
// request contains various options for the datasource (like start/end time) and the variables of an application. It
// also contains the name of the datasource, which should be used.
type GetVariablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TimeStart  int64       `protobuf:"varint,2,opt,name=timeStart,proto3" json:"timeStart,omitempty"`
	TimeEnd    int64       `protobuf:"varint,3,opt,name=timeEnd,proto3" json:"timeEnd,omitempty"`
	Resolution string      `protobuf:"bytes,4,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Variables  []*Variable `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty"`
}

func (x *GetVariablesRequest) Reset() {
	*x = GetVariablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVariablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVariablesRequest) ProtoMessage() {}

func (x *GetVariablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVariablesRequest.ProtoReflect.Descriptor instead.
func (*GetVariablesRequest) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{0}
}

func (x *GetVariablesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetVariablesRequest) GetTimeStart() int64 {
	if x != nil {
		return x.TimeStart
	}
	return 0
}

func (x *GetVariablesRequest) GetTimeEnd() int64 {
	if x != nil {
		return x.TimeEnd
	}
	return 0
}

func (x *GetVariablesRequest) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

func (x *GetVariablesRequest) GetVariables() []*Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

// GetVariablesResponse is the response for a GetVariables request, it contains the variables of the application, with
// the values and value field.
type GetVariablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variables []*Variable `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty"`
}

func (x *GetVariablesResponse) Reset() {
	*x = GetVariablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVariablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVariablesResponse) ProtoMessage() {}

func (x *GetVariablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVariablesResponse.ProtoReflect.Descriptor instead.
func (*GetVariablesResponse) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{1}
}

func (x *GetVariablesResponse) GetVariables() []*Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

// GetMetricsRequest is the structure of a call to get a time series for a chart. It must contain the name of the
// datasource, options like the start/end time and the user defined queries for the chart.
type GetMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TimeStart  int64       `protobuf:"varint,2,opt,name=timeStart,proto3" json:"timeStart,omitempty"`
	TimeEnd    int64       `protobuf:"varint,3,opt,name=timeEnd,proto3" json:"timeEnd,omitempty"`
	Resolution string      `protobuf:"bytes,4,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Variables  []*Variable `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty"`
	Queries    []*Query    `protobuf:"bytes,6,rep,name=queries,proto3" json:"queries,omitempty"`
}

func (x *GetMetricsRequest) Reset() {
	*x = GetMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsRequest) ProtoMessage() {}

func (x *GetMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsRequest.ProtoReflect.Descriptor instead.
func (*GetMetricsRequest) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{2}
}

func (x *GetMetricsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetMetricsRequest) GetTimeStart() int64 {
	if x != nil {
		return x.TimeStart
	}
	return 0
}

func (x *GetMetricsRequest) GetTimeEnd() int64 {
	if x != nil {
		return x.TimeEnd
	}
	return 0
}

func (x *GetMetricsRequest) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

func (x *GetMetricsRequest) GetVariables() []*Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *GetMetricsRequest) GetQueries() []*Query {
	if x != nil {
		return x.Queries
	}
	return nil
}

// GetMetricsResponse is the response for a GetMetrics request. It contains the data needed to render the chart in the
// React UI.
type GetMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics             []*Metrics `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	InterpolatedQueries []string   `protobuf:"bytes,2,rep,name=interpolatedQueries,proto3" json:"interpolatedQueries,omitempty"`
}

func (x *GetMetricsResponse) Reset() {
	*x = GetMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsResponse) ProtoMessage() {}

func (x *GetMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsResponse.ProtoReflect.Descriptor instead.
func (*GetMetricsResponse) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{3}
}

func (x *GetMetricsResponse) GetMetrics() []*Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *GetMetricsResponse) GetInterpolatedQueries() []string {
	if x != nil {
		return x.InterpolatedQueries
	}
	return nil
}

// Metrics contains the label and all data points for a metric. It also contains the minimum and maximum value of all
// data points.
type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string  `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Min   float64 `protobuf:"fixed64,2,opt,name=min,proto3" json:"min,omitempty"`
	Max   float64 `protobuf:"fixed64,3,opt,name=max,proto3" json:"max,omitempty"`
	Data  []*Data `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{4}
}

func (x *Metrics) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Metrics) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Metrics) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Metrics) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// Data is one data point for a metric. Each data point contains a timestamp and a value, where x is the timestamp and y
// is the value.
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64   `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{5}
}

func (x *Data) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Data) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

// Spec implements the specification for an application. This field is then used in the Application CR and contains, all
// possible fields, which can be used by a user to work with variables and charts for their data in Prometheus.
type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variables []*Variable `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty"`
	Charts    []*Chart    `protobuf:"bytes,2,rep,name=charts,proto3" json:"charts,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{6}
}

func (x *Spec) GetVariables() []*Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *Spec) GetCharts() []*Chart {
	if x != nil {
		return x.Charts
	}
	return nil
}

// Variable specifies a variable, which can be used within the charts. A variable must contain a name, a label and a
// query. It also can set the allowAll field to true, which will include an "All" option in the variables values.
// The values and value field must not be provided by the user. These fields will be set by the GetVariables call. If a
// user provide a "value", we will try to use it as the selected value.
type Variable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Label    string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Query    string   `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	AllowAll bool     `protobuf:"varint,4,opt,name=allowAll,proto3" json:"allowAll,omitempty"`
	Values   []string `protobuf:"bytes,5,rep,name=values,proto3" json:"values,omitempty"`
	Value    string   `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Variable) Reset() {
	*x = Variable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variable) ProtoMessage() {}

func (x *Variable) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variable.ProtoReflect.Descriptor instead.
func (*Variable) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{7}
}

func (x *Variable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Variable) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Variable) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Variable) GetAllowAll() bool {
	if x != nil {
		return x.AllowAll
	}
	return false
}

func (x *Variable) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Variable) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Chart represents a chart for the metrics view. A chart must contain a title, a type (line, area,  bar chart, etc.).
// It can also contain a unit for the y axis. If the stacked option is set to true all series for the chart will be
// stacked. The size parameter can be used to define the width of a chart for large screens. We are using a 12 column
// grid to display the charts, so the number must be between 1 and 12. The last option is a list of queries, which are
// executed against the datasource (e.g. For Prometheus this will be a list of PromQL queries).
type Chart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type    string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Unit    string   `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	Stacked bool     `protobuf:"varint,4,opt,name=stacked,proto3" json:"stacked,omitempty"`
	Size    int64    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Queries []*Query `protobuf:"bytes,6,rep,name=queries,proto3" json:"queries,omitempty"`
}

func (x *Chart) Reset() {
	*x = Chart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chart) ProtoMessage() {}

func (x *Chart) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chart.ProtoReflect.Descriptor instead.
func (*Chart) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{8}
}

func (x *Chart) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Chart) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Chart) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Chart) GetStacked() bool {
	if x != nil {
		return x.Stacked
	}
	return false
}

func (x *Chart) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Chart) GetQueries() []*Query {
	if x != nil {
		return x.Queries
	}
	return nil
}

// Query presents a single query to get the data, which should be shown in the chart for the metrics section. A query
// consists of a query string (e.g. PromQL) and a lable. The query and the label can contain variables via Go templating
// syntax (e.g. {{ .VARIABLE-NAME }}). For Prometheus the label can also contain a label from the returned series with
// the same syntax (e.g. {{ .SERIES-LABEL }}).
type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{9}
}

func (x *Query) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Query) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_prometheus_proto protoreflect.FileDescriptor

var file_prometheus_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x09, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65,
	0x75, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x09, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x09,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x65, 0x75, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x09, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x7d, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x07,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61,
	0x78, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x65, 0x75, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x22, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x01, 0x79, 0x22, 0x75, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3a, 0x0a, 0x09, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x65, 0x75, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x09, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x52, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x41, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x41, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xa8, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x32, 0xd0, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x12, 0x63, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x12, 0x27, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x25, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x62, 0x73, 0x69, 0x6f, 0x2f, 0x6b, 0x6f, 0x62, 0x73, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prometheus_proto_rawDescOnce sync.Once
	file_prometheus_proto_rawDescData = file_prometheus_proto_rawDesc
)

func file_prometheus_proto_rawDescGZIP() []byte {
	file_prometheus_proto_rawDescOnce.Do(func() {
		file_prometheus_proto_rawDescData = protoimpl.X.CompressGZIP(file_prometheus_proto_rawDescData)
	})
	return file_prometheus_proto_rawDescData
}

var file_prometheus_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_prometheus_proto_goTypes = []interface{}{
	(*GetVariablesRequest)(nil),  // 0: plugins.prometheus.GetVariablesRequest
	(*GetVariablesResponse)(nil), // 1: plugins.prometheus.GetVariablesResponse
	(*GetMetricsRequest)(nil),    // 2: plugins.prometheus.GetMetricsRequest
	(*GetMetricsResponse)(nil),   // 3: plugins.prometheus.GetMetricsResponse
	(*Metrics)(nil),              // 4: plugins.prometheus.Metrics
	(*Data)(nil),                 // 5: plugins.prometheus.Data
	(*Spec)(nil),                 // 6: plugins.prometheus.Spec
	(*Variable)(nil),             // 7: plugins.prometheus.Variable
	(*Chart)(nil),                // 8: plugins.prometheus.Chart
	(*Query)(nil),                // 9: plugins.prometheus.Query
}
var file_prometheus_proto_depIdxs = []int32{
	7,  // 0: plugins.prometheus.GetVariablesRequest.variables:type_name -> plugins.prometheus.Variable
	7,  // 1: plugins.prometheus.GetVariablesResponse.variables:type_name -> plugins.prometheus.Variable
	7,  // 2: plugins.prometheus.GetMetricsRequest.variables:type_name -> plugins.prometheus.Variable
	9,  // 3: plugins.prometheus.GetMetricsRequest.queries:type_name -> plugins.prometheus.Query
	4,  // 4: plugins.prometheus.GetMetricsResponse.metrics:type_name -> plugins.prometheus.Metrics
	5,  // 5: plugins.prometheus.Metrics.data:type_name -> plugins.prometheus.Data
	7,  // 6: plugins.prometheus.Spec.variables:type_name -> plugins.prometheus.Variable
	8,  // 7: plugins.prometheus.Spec.charts:type_name -> plugins.prometheus.Chart
	9,  // 8: plugins.prometheus.Chart.queries:type_name -> plugins.prometheus.Query
	0,  // 9: plugins.prometheus.Prometheus.GetVariables:input_type -> plugins.prometheus.GetVariablesRequest
	2,  // 10: plugins.prometheus.Prometheus.GetMetrics:input_type -> plugins.prometheus.GetMetricsRequest
	1,  // 11: plugins.prometheus.Prometheus.GetVariables:output_type -> plugins.prometheus.GetVariablesResponse
	3,  // 12: plugins.prometheus.Prometheus.GetMetrics:output_type -> plugins.prometheus.GetMetricsResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_prometheus_proto_init() }
func file_prometheus_proto_init() {
	if File_prometheus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prometheus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVariablesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prometheus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVariablesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prometheus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetricsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prometheus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetricsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prometheus_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prometheus_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prometheus_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prometheus_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prometheus_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prometheus_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prometheus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prometheus_proto_goTypes,
		DependencyIndexes: file_prometheus_proto_depIdxs,
		MessageInfos:      file_prometheus_proto_msgTypes,
	}.Build()
	File_prometheus_proto = out.File
	file_prometheus_proto_rawDesc = nil
	file_prometheus_proto_goTypes = nil
	file_prometheus_proto_depIdxs = nil
}
