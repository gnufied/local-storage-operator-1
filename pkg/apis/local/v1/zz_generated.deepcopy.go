// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolume) DeepCopyInto(out *LocalVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolume.
func (in *LocalVolume) DeepCopy() *LocalVolume {
	if in == nil {
		return nil
	}
	out := new(LocalVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeList) DeepCopyInto(out *LocalVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeList.
func (in *LocalVolumeList) DeepCopy() *LocalVolumeList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeSpec) DeepCopyInto(out *LocalVolumeSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(corev1.NodeSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClassDevices != nil {
		in, out := &in.StorageClassDevices, &out.StorageClassDevices
		*out = make([]StorageClassDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeSpec.
func (in *LocalVolumeSpec) DeepCopy() *LocalVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeStatus) DeepCopyInto(out *LocalVolumeStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]operatorv1.OperatorCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Generations != nil {
		in, out := &in.Generations, &out.Generations
		*out = make([]operatorv1.GenerationStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeStatus.
func (in *LocalVolumeStatus) DeepCopy() *LocalVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassDevice) DeepCopyInto(out *StorageClassDevice) {
	*out = *in
	if in.DirectoryPaths != nil {
		in, out := &in.DirectoryPaths, &out.DirectoryPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DevicePaths != nil {
		in, out := &in.DevicePaths, &out.DevicePaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassDevice.
func (in *StorageClassDevice) DeepCopy() *StorageClassDevice {
	if in == nil {
		return nil
	}
	out := new(StorageClassDevice)
	in.DeepCopyInto(out)
	return out
}
