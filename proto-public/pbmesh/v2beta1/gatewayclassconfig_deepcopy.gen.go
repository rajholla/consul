// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package meshv2beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using GatewayClassConfig within kubernetes types, where deepcopy-gen is used.
func (in *GatewayClassConfig) DeepCopyInto(out *GatewayClassConfig) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayClassConfig. Required by controller-gen.
func (in *GatewayClassConfig) DeepCopy() *GatewayClassConfig {
	if in == nil {
		return nil
	}
	out := new(GatewayClassConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GatewayClassConfig. Required by controller-gen.
func (in *GatewayClassConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Consul within kubernetes types, where deepcopy-gen is used.
func (in *Consul) DeepCopyInto(out *Consul) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Consul. Required by controller-gen.
func (in *Consul) DeepCopy() *Consul {
	if in == nil {
		return nil
	}
	out := new(Consul)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Consul. Required by controller-gen.
func (in *Consul) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Authentication within kubernetes types, where deepcopy-gen is used.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication. Required by controller-gen.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Authentication. Required by controller-gen.
func (in *Authentication) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Ports within kubernetes types, where deepcopy-gen is used.
func (in *Ports) DeepCopyInto(out *Ports) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ports. Required by controller-gen.
func (in *Ports) DeepCopy() *Ports {
	if in == nil {
		return nil
	}
	out := new(Ports)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Ports. Required by controller-gen.
func (in *Ports) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CopyAnnotations within kubernetes types, where deepcopy-gen is used.
func (in *CopyAnnotations) DeepCopyInto(out *CopyAnnotations) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopyAnnotations. Required by controller-gen.
func (in *CopyAnnotations) DeepCopy() *CopyAnnotations {
	if in == nil {
		return nil
	}
	out := new(CopyAnnotations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CopyAnnotations. Required by controller-gen.
func (in *CopyAnnotations) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Deployment within kubernetes types, where deepcopy-gen is used.
func (in *Deployment) DeepCopyInto(out *Deployment) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deployment. Required by controller-gen.
func (in *Deployment) DeepCopy() *Deployment {
	if in == nil {
		return nil
	}
	out := new(Deployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Deployment. Required by controller-gen.
func (in *Deployment) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Image within kubernetes types, where deepcopy-gen is used.
func (in *Image) DeepCopyInto(out *Image) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image. Required by controller-gen.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Image. Required by controller-gen.
func (in *Image) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
