// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: clusters.proto

package proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/kobsio/kobs/pkg/api/plugins/application/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using GetClustersRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetClustersRequest) DeepCopyInto(out *GetClustersRequest) {
	p := proto.Clone(in).(*GetClustersRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetClustersRequest. Required by controller-gen.
func (in *GetClustersRequest) DeepCopy() *GetClustersRequest {
	if in == nil {
		return nil
	}
	out := new(GetClustersRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetClustersRequest. Required by controller-gen.
func (in *GetClustersRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetClustersResponse within kubernetes types, where deepcopy-gen is used.
func (in *GetClustersResponse) DeepCopyInto(out *GetClustersResponse) {
	p := proto.Clone(in).(*GetClustersResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetClustersResponse. Required by controller-gen.
func (in *GetClustersResponse) DeepCopy() *GetClustersResponse {
	if in == nil {
		return nil
	}
	out := new(GetClustersResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetClustersResponse. Required by controller-gen.
func (in *GetClustersResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetNamespacesRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetNamespacesRequest) DeepCopyInto(out *GetNamespacesRequest) {
	p := proto.Clone(in).(*GetNamespacesRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetNamespacesRequest. Required by controller-gen.
func (in *GetNamespacesRequest) DeepCopy() *GetNamespacesRequest {
	if in == nil {
		return nil
	}
	out := new(GetNamespacesRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetNamespacesRequest. Required by controller-gen.
func (in *GetNamespacesRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetNamespacesResponse within kubernetes types, where deepcopy-gen is used.
func (in *GetNamespacesResponse) DeepCopyInto(out *GetNamespacesResponse) {
	p := proto.Clone(in).(*GetNamespacesResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetNamespacesResponse. Required by controller-gen.
func (in *GetNamespacesResponse) DeepCopy() *GetNamespacesResponse {
	if in == nil {
		return nil
	}
	out := new(GetNamespacesResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetNamespacesResponse. Required by controller-gen.
func (in *GetNamespacesResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetCRDsRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetCRDsRequest) DeepCopyInto(out *GetCRDsRequest) {
	p := proto.Clone(in).(*GetCRDsRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetCRDsRequest. Required by controller-gen.
func (in *GetCRDsRequest) DeepCopy() *GetCRDsRequest {
	if in == nil {
		return nil
	}
	out := new(GetCRDsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetCRDsRequest. Required by controller-gen.
func (in *GetCRDsRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetCRDsResponse within kubernetes types, where deepcopy-gen is used.
func (in *GetCRDsResponse) DeepCopyInto(out *GetCRDsResponse) {
	p := proto.Clone(in).(*GetCRDsResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetCRDsResponse. Required by controller-gen.
func (in *GetCRDsResponse) DeepCopy() *GetCRDsResponse {
	if in == nil {
		return nil
	}
	out := new(GetCRDsResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetCRDsResponse. Required by controller-gen.
func (in *GetCRDsResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetResourcesRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetResourcesRequest) DeepCopyInto(out *GetResourcesRequest) {
	p := proto.Clone(in).(*GetResourcesRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetResourcesRequest. Required by controller-gen.
func (in *GetResourcesRequest) DeepCopy() *GetResourcesRequest {
	if in == nil {
		return nil
	}
	out := new(GetResourcesRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetResourcesRequest. Required by controller-gen.
func (in *GetResourcesRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetResourcesResponse within kubernetes types, where deepcopy-gen is used.
func (in *GetResourcesResponse) DeepCopyInto(out *GetResourcesResponse) {
	p := proto.Clone(in).(*GetResourcesResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetResourcesResponse. Required by controller-gen.
func (in *GetResourcesResponse) DeepCopy() *GetResourcesResponse {
	if in == nil {
		return nil
	}
	out := new(GetResourcesResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetResourcesResponse. Required by controller-gen.
func (in *GetResourcesResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Resources within kubernetes types, where deepcopy-gen is used.
func (in *Resources) DeepCopyInto(out *Resources) {
	p := proto.Clone(in).(*Resources)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources. Required by controller-gen.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Resources. Required by controller-gen.
func (in *Resources) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetApplicationsRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetApplicationsRequest) DeepCopyInto(out *GetApplicationsRequest) {
	p := proto.Clone(in).(*GetApplicationsRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationsRequest. Required by controller-gen.
func (in *GetApplicationsRequest) DeepCopy() *GetApplicationsRequest {
	if in == nil {
		return nil
	}
	out := new(GetApplicationsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationsRequest. Required by controller-gen.
func (in *GetApplicationsRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetApplicationsResponse within kubernetes types, where deepcopy-gen is used.
func (in *GetApplicationsResponse) DeepCopyInto(out *GetApplicationsResponse) {
	p := proto.Clone(in).(*GetApplicationsResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationsResponse. Required by controller-gen.
func (in *GetApplicationsResponse) DeepCopy() *GetApplicationsResponse {
	if in == nil {
		return nil
	}
	out := new(GetApplicationsResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationsResponse. Required by controller-gen.
func (in *GetApplicationsResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetApplicationRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetApplicationRequest) DeepCopyInto(out *GetApplicationRequest) {
	p := proto.Clone(in).(*GetApplicationRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationRequest. Required by controller-gen.
func (in *GetApplicationRequest) DeepCopy() *GetApplicationRequest {
	if in == nil {
		return nil
	}
	out := new(GetApplicationRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationRequest. Required by controller-gen.
func (in *GetApplicationRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetApplicationResponse within kubernetes types, where deepcopy-gen is used.
func (in *GetApplicationResponse) DeepCopyInto(out *GetApplicationResponse) {
	p := proto.Clone(in).(*GetApplicationResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationResponse. Required by controller-gen.
func (in *GetApplicationResponse) DeepCopy() *GetApplicationResponse {
	if in == nil {
		return nil
	}
	out := new(GetApplicationResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationResponse. Required by controller-gen.
func (in *GetApplicationResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CRD within kubernetes types, where deepcopy-gen is used.
func (in *CRD) DeepCopyInto(out *CRD) {
	p := proto.Clone(in).(*CRD)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRD. Required by controller-gen.
func (in *CRD) DeepCopy() *CRD {
	if in == nil {
		return nil
	}
	out := new(CRD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CRD. Required by controller-gen.
func (in *CRD) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CRDColumn within kubernetes types, where deepcopy-gen is used.
func (in *CRDColumn) DeepCopyInto(out *CRDColumn) {
	p := proto.Clone(in).(*CRDColumn)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDColumn. Required by controller-gen.
func (in *CRDColumn) DeepCopy() *CRDColumn {
	if in == nil {
		return nil
	}
	out := new(CRDColumn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CRDColumn. Required by controller-gen.
func (in *CRDColumn) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetApplicationsTopologyRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetApplicationsTopologyRequest) DeepCopyInto(out *GetApplicationsTopologyRequest) {
	p := proto.Clone(in).(*GetApplicationsTopologyRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationsTopologyRequest. Required by controller-gen.
func (in *GetApplicationsTopologyRequest) DeepCopy() *GetApplicationsTopologyRequest {
	if in == nil {
		return nil
	}
	out := new(GetApplicationsTopologyRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationsTopologyRequest. Required by controller-gen.
func (in *GetApplicationsTopologyRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetApplicationsTopologyResponse within kubernetes types, where deepcopy-gen is used.
func (in *GetApplicationsTopologyResponse) DeepCopyInto(out *GetApplicationsTopologyResponse) {
	p := proto.Clone(in).(*GetApplicationsTopologyResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationsTopologyResponse. Required by controller-gen.
func (in *GetApplicationsTopologyResponse) DeepCopy() *GetApplicationsTopologyResponse {
	if in == nil {
		return nil
	}
	out := new(GetApplicationsTopologyResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetApplicationsTopologyResponse. Required by controller-gen.
func (in *GetApplicationsTopologyResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Edge within kubernetes types, where deepcopy-gen is used.
func (in *Edge) DeepCopyInto(out *Edge) {
	p := proto.Clone(in).(*Edge)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Edge. Required by controller-gen.
func (in *Edge) DeepCopy() *Edge {
	if in == nil {
		return nil
	}
	out := new(Edge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Edge. Required by controller-gen.
func (in *Edge) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Node within kubernetes types, where deepcopy-gen is used.
func (in *Node) DeepCopyInto(out *Node) {
	p := proto.Clone(in).(*Node)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node. Required by controller-gen.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Node. Required by controller-gen.
func (in *Node) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
