// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CmdConfig) DeepCopyInto(out *CmdConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CmdConfig.
func (in *CmdConfig) DeepCopy() *CmdConfig {
	if in == nil {
		return nil
	}
	out := new(CmdConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlusterBlockSpec) DeepCopyInto(out *GlusterBlockSpec) {
	*out = *in
	if in.Provisioner != nil {
		in, out := &in.Provisioner, &out.Provisioner
		if *in == nil {
			*out = nil
		} else {
			*out = new(VolumeProvisioner)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlusterBlockSpec.
func (in *GlusterBlockSpec) DeepCopy() *GlusterBlockSpec {
	if in == nil {
		return nil
	}
	out := new(GlusterBlockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlusterCluster) DeepCopyInto(out *GlusterCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		if *in == nil {
			*out = nil
		} else {
			*out = new(GlusterClusterSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(GlusterClusterStatus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlusterCluster.
func (in *GlusterCluster) DeepCopy() *GlusterCluster {
	if in == nil {
		return nil
	}
	out := new(GlusterCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlusterCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlusterClusterList) DeepCopyInto(out *GlusterClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlusterCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlusterClusterList.
func (in *GlusterClusterList) DeepCopy() *GlusterClusterList {
	if in == nil {
		return nil
	}
	out := new(GlusterClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlusterClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlusterClusterSpec) DeepCopyInto(out *GlusterClusterSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]*Node, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Node)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.Heketi != nil {
		in, out := &in.Heketi, &out.Heketi
		if *in == nil {
			*out = nil
		} else {
			*out = new(HeketiSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.GlusterFS != nil {
		in, out := &in.GlusterFS, &out.GlusterFS
		if *in == nil {
			*out = nil
		} else {
			*out = new(GlusterfsSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.GlusterBlock != nil {
		in, out := &in.GlusterBlock, &out.GlusterBlock
		if *in == nil {
			*out = nil
		} else {
			*out = new(GlusterBlockSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlusterClusterSpec.
func (in *GlusterClusterSpec) DeepCopy() *GlusterClusterSpec {
	if in == nil {
		return nil
	}
	out := new(GlusterClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlusterClusterStatus) DeepCopyInto(out *GlusterClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlusterClusterStatus.
func (in *GlusterClusterStatus) DeepCopy() *GlusterClusterStatus {
	if in == nil {
		return nil
	}
	out := new(GlusterClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlusterFSConfig) DeepCopyInto(out *GlusterFSConfig) {
	*out = *in
	out.SshConfig = in.SshConfig
	out.KubeConfig = in.KubeConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlusterFSConfig.
func (in *GlusterFSConfig) DeepCopy() *GlusterFSConfig {
	if in == nil {
		return nil
	}
	out := new(GlusterFSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlusterfsSpec) DeepCopyInto(out *GlusterfsSpec) {
	*out = *in
	if in.Provisioner != nil {
		in, out := &in.Provisioner, &out.Provisioner
		if *in == nil {
			*out = nil
		} else {
			*out = new(VolumeProvisioner)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlusterfsSpec.
func (in *GlusterfsSpec) DeepCopy() *GlusterfsSpec {
	if in == nil {
		return nil
	}
	out := new(GlusterfsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeketiConfigSpec) DeepCopyInto(out *HeketiConfigSpec) {
	*out = *in
	out.JwtConfig = in.JwtConfig
	out.GlusterFS = in.GlusterFS
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeketiConfigSpec.
func (in *HeketiConfigSpec) DeepCopy() *HeketiConfigSpec {
	if in == nil {
		return nil
	}
	out := new(HeketiConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeketiSSHSpec) DeepCopyInto(out *HeketiSSHSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeketiSSHSpec.
func (in *HeketiSSHSpec) DeepCopy() *HeketiSSHSpec {
	if in == nil {
		return nil
	}
	out := new(HeketiSSHSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeketiSpec) DeepCopyInto(out *HeketiSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		if *in == nil {
			*out = nil
		} else {
			*out = new(HeketiConfigSpec)
			**out = **in
		}
	}
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		if *in == nil {
			*out = nil
		} else {
			*out = new(Node)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeketiSpec.
func (in *HeketiSpec) DeepCopy() *HeketiSpec {
	if in == nil {
		return nil
	}
	out := new(HeketiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfig) DeepCopyInto(out *KubeConfig) {
	*out = *in
	out.CmdConfig = in.CmdConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfig.
func (in *KubeConfig) DeepCopy() *KubeConfig {
	if in == nil {
		return nil
	}
	out := new(KubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]*NodeVolume, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(NodeVolume)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.StateVolume != nil {
		in, out := &in.StateVolume, &out.StateVolume
		if *in == nil {
			*out = nil
		} else {
			*out = new(NodeVolume)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeVolume) DeepCopyInto(out *NodeVolume) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.PersistentVolumeSpec.DeepCopyInto(&out.PersistentVolumeSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeVolume.
func (in *NodeVolume) DeepCopy() *NodeVolume {
	if in == nil {
		return nil
	}
	out := new(NodeVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SshConfig) DeepCopyInto(out *SshConfig) {
	*out = *in
	out.CmdConfig = in.CmdConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SshConfig.
func (in *SshConfig) DeepCopy() *SshConfig {
	if in == nil {
		return nil
	}
	out := new(SshConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeProvisioner) DeepCopyInto(out *VolumeProvisioner) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeProvisioner.
func (in *VolumeProvisioner) DeepCopy() *VolumeProvisioner {
	if in == nil {
		return nil
	}
	out := new(VolumeProvisioner)
	in.DeepCopyInto(out)
	return out
}
