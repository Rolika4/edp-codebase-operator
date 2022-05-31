//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"github.com/epam/edp-jenkins-operator/v2/pkg/apis/v2/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDStageDeploy) DeepCopyInto(out *CDStageDeploy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDStageDeploy.
func (in *CDStageDeploy) DeepCopy() *CDStageDeploy {
	if in == nil {
		return nil
	}
	out := new(CDStageDeploy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDStageDeploy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDStageDeployList) DeepCopyInto(out *CDStageDeployList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CDStageDeploy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDStageDeployList.
func (in *CDStageDeployList) DeepCopy() *CDStageDeployList {
	if in == nil {
		return nil
	}
	out := new(CDStageDeployList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDStageDeployList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDStageDeploySpec) DeepCopyInto(out *CDStageDeploySpec) {
	*out = *in
	out.Tag = in.Tag
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]v1alpha1.Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDStageDeploySpec.
func (in *CDStageDeploySpec) DeepCopy() *CDStageDeploySpec {
	if in == nil {
		return nil
	}
	out := new(CDStageDeploySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDStageDeployStatus) DeepCopyInto(out *CDStageDeployStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDStageDeployStatus.
func (in *CDStageDeployStatus) DeepCopy() *CDStageDeployStatus {
	if in == nil {
		return nil
	}
	out := new(CDStageDeployStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Codebase) DeepCopyInto(out *Codebase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Codebase.
func (in *Codebase) DeepCopy() *Codebase {
	if in == nil {
		return nil
	}
	out := new(Codebase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Codebase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseBranch) DeepCopyInto(out *CodebaseBranch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseBranch.
func (in *CodebaseBranch) DeepCopy() *CodebaseBranch {
	if in == nil {
		return nil
	}
	out := new(CodebaseBranch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CodebaseBranch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseBranchList) DeepCopyInto(out *CodebaseBranchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CodebaseBranch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseBranchList.
func (in *CodebaseBranchList) DeepCopy() *CodebaseBranchList {
	if in == nil {
		return nil
	}
	out := new(CodebaseBranchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CodebaseBranchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseBranchSpec) DeepCopyInto(out *CodebaseBranchSpec) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ReleaseJobParams != nil {
		in, out := &in.ReleaseJobParams, &out.ReleaseJobParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseBranchSpec.
func (in *CodebaseBranchSpec) DeepCopy() *CodebaseBranchSpec {
	if in == nil {
		return nil
	}
	out := new(CodebaseBranchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseBranchStatus) DeepCopyInto(out *CodebaseBranchStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
	if in.VersionHistory != nil {
		in, out := &in.VersionHistory, &out.VersionHistory
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LastSuccessfulBuild != nil {
		in, out := &in.LastSuccessfulBuild, &out.LastSuccessfulBuild
		*out = new(string)
		**out = **in
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseBranchStatus.
func (in *CodebaseBranchStatus) DeepCopy() *CodebaseBranchStatus {
	if in == nil {
		return nil
	}
	out := new(CodebaseBranchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseImageStream) DeepCopyInto(out *CodebaseImageStream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseImageStream.
func (in *CodebaseImageStream) DeepCopy() *CodebaseImageStream {
	if in == nil {
		return nil
	}
	out := new(CodebaseImageStream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CodebaseImageStream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseImageStreamList) DeepCopyInto(out *CodebaseImageStreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CodebaseImageStream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseImageStreamList.
func (in *CodebaseImageStreamList) DeepCopy() *CodebaseImageStreamList {
	if in == nil {
		return nil
	}
	out := new(CodebaseImageStreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CodebaseImageStreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseImageStreamSpec) DeepCopyInto(out *CodebaseImageStreamSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseImageStreamSpec.
func (in *CodebaseImageStreamSpec) DeepCopy() *CodebaseImageStreamSpec {
	if in == nil {
		return nil
	}
	out := new(CodebaseImageStreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseImageStreamStatus) DeepCopyInto(out *CodebaseImageStreamStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseImageStreamStatus.
func (in *CodebaseImageStreamStatus) DeepCopy() *CodebaseImageStreamStatus {
	if in == nil {
		return nil
	}
	out := new(CodebaseImageStreamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseList) DeepCopyInto(out *CodebaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Codebase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseList.
func (in *CodebaseList) DeepCopy() *CodebaseList {
	if in == nil {
		return nil
	}
	out := new(CodebaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CodebaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseSpec) DeepCopyInto(out *CodebaseSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Framework != nil {
		in, out := &in.Framework, &out.Framework
		*out = new(string)
		**out = **in
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(Repository)
		**out = **in
	}
	if in.TestReportFramework != nil {
		in, out := &in.TestReportFramework, &out.TestReportFramework
		*out = new(string)
		**out = **in
	}
	if in.GitUrlPath != nil {
		in, out := &in.GitUrlPath, &out.GitUrlPath
		*out = new(string)
		**out = **in
	}
	if in.JenkinsSlave != nil {
		in, out := &in.JenkinsSlave, &out.JenkinsSlave
		*out = new(string)
		**out = **in
	}
	if in.JobProvisioning != nil {
		in, out := &in.JobProvisioning, &out.JobProvisioning
		*out = new(string)
		**out = **in
	}
	in.Versioning.DeepCopyInto(&out.Versioning)
	if in.JiraServer != nil {
		in, out := &in.JiraServer, &out.JiraServer
		*out = new(string)
		**out = **in
	}
	if in.CommitMessagePattern != nil {
		in, out := &in.CommitMessagePattern, &out.CommitMessagePattern
		*out = new(string)
		**out = **in
	}
	if in.TicketNamePattern != nil {
		in, out := &in.TicketNamePattern, &out.TicketNamePattern
		*out = new(string)
		**out = **in
	}
	if in.Perf != nil {
		in, out := &in.Perf, &out.Perf
		*out = new(Perf)
		(*in).DeepCopyInto(*out)
	}
	if in.JiraIssueMetadataPayload != nil {
		in, out := &in.JiraIssueMetadataPayload, &out.JiraIssueMetadataPayload
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseSpec.
func (in *CodebaseSpec) DeepCopy() *CodebaseSpec {
	if in == nil {
		return nil
	}
	out := new(CodebaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebaseStatus) DeepCopyInto(out *CodebaseStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebaseStatus.
func (in *CodebaseStatus) DeepCopy() *CodebaseStatus {
	if in == nil {
		return nil
	}
	out := new(CodebaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitServer) DeepCopyInto(out *GitServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitServer.
func (in *GitServer) DeepCopy() *GitServer {
	if in == nil {
		return nil
	}
	out := new(GitServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitServerList) DeepCopyInto(out *GitServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitServerList.
func (in *GitServerList) DeepCopy() *GitServerList {
	if in == nil {
		return nil
	}
	out := new(GitServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitServerSpec) DeepCopyInto(out *GitServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitServerSpec.
func (in *GitServerSpec) DeepCopy() *GitServerSpec {
	if in == nil {
		return nil
	}
	out := new(GitServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitServerStatus) DeepCopyInto(out *GitServerStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitServerStatus.
func (in *GitServerStatus) DeepCopy() *GitServerStatus {
	if in == nil {
		return nil
	}
	out := new(GitServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTag) DeepCopyInto(out *GitTag) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTag.
func (in *GitTag) DeepCopy() *GitTag {
	if in == nil {
		return nil
	}
	out := new(GitTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitTag) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTagList) DeepCopyInto(out *GitTagList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTagList.
func (in *GitTagList) DeepCopy() *GitTagList {
	if in == nil {
		return nil
	}
	out := new(GitTagList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitTagList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTagSpec) DeepCopyInto(out *GitTagSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTagSpec.
func (in *GitTagSpec) DeepCopy() *GitTagSpec {
	if in == nil {
		return nil
	}
	out := new(GitTagSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTagStatus) DeepCopyInto(out *GitTagStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTagStatus.
func (in *GitTagStatus) DeepCopy() *GitTagStatus {
	if in == nil {
		return nil
	}
	out := new(GitTagStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamTag) DeepCopyInto(out *ImageStreamTag) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamTag.
func (in *ImageStreamTag) DeepCopy() *ImageStreamTag {
	if in == nil {
		return nil
	}
	out := new(ImageStreamTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStreamTag) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamTagList) DeepCopyInto(out *ImageStreamTagList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageStreamTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamTagList.
func (in *ImageStreamTagList) DeepCopy() *ImageStreamTagList {
	if in == nil {
		return nil
	}
	out := new(ImageStreamTagList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageStreamTagList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamTagSpec) DeepCopyInto(out *ImageStreamTagSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamTagSpec.
func (in *ImageStreamTagSpec) DeepCopy() *ImageStreamTagSpec {
	if in == nil {
		return nil
	}
	out := new(ImageStreamTagSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStreamTagStatus) DeepCopyInto(out *ImageStreamTagStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStreamTagStatus.
func (in *ImageStreamTagStatus) DeepCopy() *ImageStreamTagStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStreamTagStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JiraIssueMetadata) DeepCopyInto(out *JiraIssueMetadata) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JiraIssueMetadata.
func (in *JiraIssueMetadata) DeepCopy() *JiraIssueMetadata {
	if in == nil {
		return nil
	}
	out := new(JiraIssueMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JiraIssueMetadata) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JiraIssueMetadataList) DeepCopyInto(out *JiraIssueMetadataList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JiraIssueMetadata, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JiraIssueMetadataList.
func (in *JiraIssueMetadataList) DeepCopy() *JiraIssueMetadataList {
	if in == nil {
		return nil
	}
	out := new(JiraIssueMetadataList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JiraIssueMetadataList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JiraIssueMetadataSpec) DeepCopyInto(out *JiraIssueMetadataSpec) {
	*out = *in
	if in.Commits != nil {
		in, out := &in.Commits, &out.Commits
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tickets != nil {
		in, out := &in.Tickets, &out.Tickets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JiraIssueMetadataSpec.
func (in *JiraIssueMetadataSpec) DeepCopy() *JiraIssueMetadataSpec {
	if in == nil {
		return nil
	}
	out := new(JiraIssueMetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JiraIssueMetadataStatus) DeepCopyInto(out *JiraIssueMetadataStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JiraIssueMetadataStatus.
func (in *JiraIssueMetadataStatus) DeepCopy() *JiraIssueMetadataStatus {
	if in == nil {
		return nil
	}
	out := new(JiraIssueMetadataStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JiraServer) DeepCopyInto(out *JiraServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JiraServer.
func (in *JiraServer) DeepCopy() *JiraServer {
	if in == nil {
		return nil
	}
	out := new(JiraServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JiraServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JiraServerList) DeepCopyInto(out *JiraServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JiraServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JiraServerList.
func (in *JiraServerList) DeepCopy() *JiraServerList {
	if in == nil {
		return nil
	}
	out := new(JiraServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JiraServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JiraServerSpec) DeepCopyInto(out *JiraServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JiraServerSpec.
func (in *JiraServerSpec) DeepCopy() *JiraServerSpec {
	if in == nil {
		return nil
	}
	out := new(JiraServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JiraServerStatus) DeepCopyInto(out *JiraServerStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JiraServerStatus.
func (in *JiraServerStatus) DeepCopy() *JiraServerStatus {
	if in == nil {
		return nil
	}
	out := new(JiraServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Perf) DeepCopyInto(out *Perf) {
	*out = *in
	if in.DataSources != nil {
		in, out := &in.DataSources, &out.DataSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Perf.
func (in *Perf) DeepCopy() *Perf {
	if in == nil {
		return nil
	}
	out := new(Perf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Versioning) DeepCopyInto(out *Versioning) {
	*out = *in
	if in.StartFrom != nil {
		in, out := &in.StartFrom, &out.StartFrom
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Versioning.
func (in *Versioning) DeepCopy() *Versioning {
	if in == nil {
		return nil
	}
	out := new(Versioning)
	in.DeepCopyInto(out)
	return out
}
