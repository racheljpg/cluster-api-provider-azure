//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package compat

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUpgradeSettings) DeepCopyInto(out *ClusterUpgradeSettings) {
	*out = *in
	if in.OverrideSettings != nil {
		in, out := &in.OverrideSettings, &out.OverrideSettings
		*out = new(UpgradeOverrideSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUpgradeSettings.
func (in *ClusterUpgradeSettings) DeepCopy() *ClusterUpgradeSettings {
	if in == nil {
		return nil
	}
	out := new(ClusterUpgradeSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUpgradeSettings_STATUS) DeepCopyInto(out *ClusterUpgradeSettings_STATUS) {
	*out = *in
	if in.OverrideSettings != nil {
		in, out := &in.OverrideSettings, &out.OverrideSettings
		*out = new(UpgradeOverrideSettings_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUpgradeSettings_STATUS.
func (in *ClusterUpgradeSettings_STATUS) DeepCopy() *ClusterUpgradeSettings_STATUS {
	if in == nil {
		return nil
	}
	out := new(ClusterUpgradeSettings_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioComponents) DeepCopyInto(out *IstioComponents) {
	*out = *in
	if in.IngressGateways != nil {
		in, out := &in.IngressGateways, &out.IngressGateways
		*out = make([]IstioIngressGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioComponents.
func (in *IstioComponents) DeepCopy() *IstioComponents {
	if in == nil {
		return nil
	}
	out := new(IstioComponents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioComponents_STATUS) DeepCopyInto(out *IstioComponents_STATUS) {
	*out = *in
	if in.IngressGateways != nil {
		in, out := &in.IngressGateways, &out.IngressGateways
		*out = make([]IstioIngressGateway_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioComponents_STATUS.
func (in *IstioComponents_STATUS) DeepCopy() *IstioComponents_STATUS {
	if in == nil {
		return nil
	}
	out := new(IstioComponents_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioIngressGateway) DeepCopyInto(out *IstioIngressGateway) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioIngressGateway.
func (in *IstioIngressGateway) DeepCopy() *IstioIngressGateway {
	if in == nil {
		return nil
	}
	out := new(IstioIngressGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioIngressGateway_STATUS) DeepCopyInto(out *IstioIngressGateway_STATUS) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioIngressGateway_STATUS.
func (in *IstioIngressGateway_STATUS) DeepCopy() *IstioIngressGateway_STATUS {
	if in == nil {
		return nil
	}
	out := new(IstioIngressGateway_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioServiceMesh) DeepCopyInto(out *IstioServiceMesh) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = new(IstioComponents)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioServiceMesh.
func (in *IstioServiceMesh) DeepCopy() *IstioServiceMesh {
	if in == nil {
		return nil
	}
	out := new(IstioServiceMesh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioServiceMesh_STATUS) DeepCopyInto(out *IstioServiceMesh_STATUS) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = new(IstioComponents_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioServiceMesh_STATUS.
func (in *IstioServiceMesh_STATUS) DeepCopy() *IstioServiceMesh_STATUS {
	if in == nil {
		return nil
	}
	out := new(IstioServiceMesh_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler) DeepCopyInto(out *ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler) {
	*out = *in
	if in.ControlledValues != nil {
		in, out := &in.ControlledValues, &out.ControlledValues
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UpdateMode != nil {
		in, out := &in.UpdateMode, &out.UpdateMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler.
func (in *ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler) DeepCopy() *ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS) DeepCopyInto(out *ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS) {
	*out = *in
	if in.ControlledValues != nil {
		in, out := &in.ControlledValues, &out.ControlledValues
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UpdateMode != nil {
		in, out := &in.UpdateMode, &out.UpdateMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS.
func (in *ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS) DeepCopy() *ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshProfile) DeepCopyInto(out *ServiceMeshProfile) {
	*out = *in
	if in.Istio != nil {
		in, out := &in.Istio, &out.Istio
		*out = new(IstioServiceMesh)
		(*in).DeepCopyInto(*out)
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshProfile.
func (in *ServiceMeshProfile) DeepCopy() *ServiceMeshProfile {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMeshProfile_STATUS) DeepCopyInto(out *ServiceMeshProfile_STATUS) {
	*out = *in
	if in.Istio != nil {
		in, out := &in.Istio, &out.Istio
		*out = new(IstioServiceMesh_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMeshProfile_STATUS.
func (in *ServiceMeshProfile_STATUS) DeepCopy() *ServiceMeshProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(ServiceMeshProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeOverrideSettings) DeepCopyInto(out *UpgradeOverrideSettings) {
	*out = *in
	if in.ControlPlaneOverrides != nil {
		in, out := &in.ControlPlaneOverrides, &out.ControlPlaneOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Until != nil {
		in, out := &in.Until, &out.Until
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeOverrideSettings.
func (in *UpgradeOverrideSettings) DeepCopy() *UpgradeOverrideSettings {
	if in == nil {
		return nil
	}
	out := new(UpgradeOverrideSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeOverrideSettings_STATUS) DeepCopyInto(out *UpgradeOverrideSettings_STATUS) {
	*out = *in
	if in.ControlPlaneOverrides != nil {
		in, out := &in.ControlPlaneOverrides, &out.ControlPlaneOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Until != nil {
		in, out := &in.Until, &out.Until
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeOverrideSettings_STATUS.
func (in *UpgradeOverrideSettings_STATUS) DeepCopy() *UpgradeOverrideSettings_STATUS {
	if in == nil {
		return nil
	}
	out := new(UpgradeOverrideSettings_STATUS)
	in.DeepCopyInto(out)
	return out
}