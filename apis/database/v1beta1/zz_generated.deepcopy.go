// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityZone) DeepCopyInto(out *AvailabilityZone) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityZone.
func (in *AvailabilityZone) DeepCopy() *AvailabilityZone {
	if in == nil {
		return nil
	}
	out := new(AvailabilityZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchLogsExportConfiguration) DeepCopyInto(out *CloudwatchLogsExportConfiguration) {
	*out = *in
	if in.DisableLogTypes != nil {
		in, out := &in.DisableLogTypes, &out.DisableLogTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EnableLogTypes != nil {
		in, out := &in.EnableLogTypes, &out.EnableLogTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchLogsExportConfiguration.
func (in *CloudwatchLogsExportConfiguration) DeepCopy() *CloudwatchLogsExportConfiguration {
	if in == nil {
		return nil
	}
	out := new(CloudwatchLogsExportConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBInstanceStatusInfo) DeepCopyInto(out *DBInstanceStatusInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBInstanceStatusInfo.
func (in *DBInstanceStatusInfo) DeepCopy() *DBInstanceStatusInfo {
	if in == nil {
		return nil
	}
	out := new(DBInstanceStatusInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBParameterGroupStatus) DeepCopyInto(out *DBParameterGroupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBParameterGroupStatus.
func (in *DBParameterGroupStatus) DeepCopy() *DBParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DBParameterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSecurityGroupMembership) DeepCopyInto(out *DBSecurityGroupMembership) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSecurityGroupMembership.
func (in *DBSecurityGroupMembership) DeepCopy() *DBSecurityGroupMembership {
	if in == nil {
		return nil
	}
	out := new(DBSecurityGroupMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroup) DeepCopyInto(out *DBSubnetGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroup.
func (in *DBSubnetGroup) DeepCopy() *DBSubnetGroup {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBSubnetGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupInRDS) DeepCopyInto(out *DBSubnetGroupInRDS) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]SubnetInRDS, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupInRDS.
func (in *DBSubnetGroupInRDS) DeepCopy() *DBSubnetGroupInRDS {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupInRDS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupList) DeepCopyInto(out *DBSubnetGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DBSubnetGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupList.
func (in *DBSubnetGroupList) DeepCopy() *DBSubnetGroupList {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBSubnetGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupObservation) DeepCopyInto(out *DBSubnetGroupObservation) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]Subnet, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupObservation.
func (in *DBSubnetGroupObservation) DeepCopy() *DBSubnetGroupObservation {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupParameters) DeepCopyInto(out *DBSubnetGroupParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupParameters.
func (in *DBSubnetGroupParameters) DeepCopy() *DBSubnetGroupParameters {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupSpec) DeepCopyInto(out *DBSubnetGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupSpec.
func (in *DBSubnetGroupSpec) DeepCopy() *DBSubnetGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBSubnetGroupStatus) DeepCopyInto(out *DBSubnetGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBSubnetGroupStatus.
func (in *DBSubnetGroupStatus) DeepCopy() *DBSubnetGroupStatus {
	if in == nil {
		return nil
	}
	out := new(DBSubnetGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainMembership) DeepCopyInto(out *DomainMembership) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainMembership.
func (in *DomainMembership) DeepCopy() *DomainMembership {
	if in == nil {
		return nil
	}
	out := new(DomainMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionGroupMembership) DeepCopyInto(out *OptionGroupMembership) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionGroupMembership.
func (in *OptionGroupMembership) DeepCopy() *OptionGroupMembership {
	if in == nil {
		return nil
	}
	out := new(OptionGroupMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingCloudwatchLogsExports) DeepCopyInto(out *PendingCloudwatchLogsExports) {
	*out = *in
	if in.LogTypesToDisable != nil {
		in, out := &in.LogTypesToDisable, &out.LogTypesToDisable
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogTypesToEnable != nil {
		in, out := &in.LogTypesToEnable, &out.LogTypesToEnable
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingCloudwatchLogsExports.
func (in *PendingCloudwatchLogsExports) DeepCopy() *PendingCloudwatchLogsExports {
	if in == nil {
		return nil
	}
	out := new(PendingCloudwatchLogsExports)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingModifiedValues) DeepCopyInto(out *PendingModifiedValues) {
	*out = *in
	in.PendingCloudwatchLogsExports.DeepCopyInto(&out.PendingCloudwatchLogsExports)
	if in.ProcessorFeatures != nil {
		in, out := &in.ProcessorFeatures, &out.ProcessorFeatures
		*out = make([]ProcessorFeature, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingModifiedValues.
func (in *PendingModifiedValues) DeepCopy() *PendingModifiedValues {
	if in == nil {
		return nil
	}
	out := new(PendingModifiedValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessorFeature) DeepCopyInto(out *ProcessorFeature) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessorFeature.
func (in *ProcessorFeature) DeepCopy() *ProcessorFeature {
	if in == nil {
		return nil
	}
	out := new(ProcessorFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstance) DeepCopyInto(out *RDSInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstance.
func (in *RDSInstance) DeepCopy() *RDSInstance {
	if in == nil {
		return nil
	}
	out := new(RDSInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceList) DeepCopyInto(out *RDSInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RDSInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceList.
func (in *RDSInstanceList) DeepCopy() *RDSInstanceList {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceObservation) DeepCopyInto(out *RDSInstanceObservation) {
	*out = *in
	if in.DBParameterGroups != nil {
		in, out := &in.DBParameterGroups, &out.DBParameterGroups
		*out = make([]DBParameterGroupStatus, len(*in))
		copy(*out, *in)
	}
	if in.DBSecurityGroups != nil {
		in, out := &in.DBSecurityGroups, &out.DBSecurityGroups
		*out = make([]DBSecurityGroupMembership, len(*in))
		copy(*out, *in)
	}
	in.DBSubnetGroup.DeepCopyInto(&out.DBSubnetGroup)
	if in.DomainMemberships != nil {
		in, out := &in.DomainMemberships, &out.DomainMemberships
		*out = make([]DomainMembership, len(*in))
		copy(*out, *in)
	}
	if in.InstanceCreateTime != nil {
		in, out := &in.InstanceCreateTime, &out.InstanceCreateTime
		*out = (*in).DeepCopy()
	}
	out.Endpoint = in.Endpoint
	if in.LatestRestorableTime != nil {
		in, out := &in.LatestRestorableTime, &out.LatestRestorableTime
		*out = (*in).DeepCopy()
	}
	if in.OptionGroupMemberships != nil {
		in, out := &in.OptionGroupMemberships, &out.OptionGroupMemberships
		*out = make([]OptionGroupMembership, len(*in))
		copy(*out, *in)
	}
	in.PendingModifiedValues.DeepCopyInto(&out.PendingModifiedValues)
	if in.ReadReplicaDBClusterIdentifiers != nil {
		in, out := &in.ReadReplicaDBClusterIdentifiers, &out.ReadReplicaDBClusterIdentifiers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReadReplicaDBInstanceIdentifiers != nil {
		in, out := &in.ReadReplicaDBInstanceIdentifiers, &out.ReadReplicaDBInstanceIdentifiers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StatusInfos != nil {
		in, out := &in.StatusInfos, &out.StatusInfos
		*out = make([]DBInstanceStatusInfo, len(*in))
		copy(*out, *in)
	}
	if in.VPCSecurityGroups != nil {
		in, out := &in.VPCSecurityGroups, &out.VPCSecurityGroups
		*out = make([]VPCSecurityGroupMembership, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceObservation.
func (in *RDSInstanceObservation) DeepCopy() *RDSInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceParameters) DeepCopyInto(out *RDSInstanceParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.AllocatedStorage != nil {
		in, out := &in.AllocatedStorage, &out.AllocatedStorage
		*out = new(int)
		**out = **in
	}
	if in.AutoMinorVersionUpgrade != nil {
		in, out := &in.AutoMinorVersionUpgrade, &out.AutoMinorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.BackupRetentionPeriod != nil {
		in, out := &in.BackupRetentionPeriod, &out.BackupRetentionPeriod
		*out = new(int)
		**out = **in
	}
	if in.CACertificateIdentifier != nil {
		in, out := &in.CACertificateIdentifier, &out.CACertificateIdentifier
		*out = new(string)
		**out = **in
	}
	if in.CharacterSetName != nil {
		in, out := &in.CharacterSetName, &out.CharacterSetName
		*out = new(string)
		**out = **in
	}
	if in.CopyTagsToSnapshot != nil {
		in, out := &in.CopyTagsToSnapshot, &out.CopyTagsToSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.DBClusterIdentifier != nil {
		in, out := &in.DBClusterIdentifier, &out.DBClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.DBName != nil {
		in, out := &in.DBName, &out.DBName
		*out = new(string)
		**out = **in
	}
	if in.DBSecurityGroups != nil {
		in, out := &in.DBSecurityGroups, &out.DBSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DBSubnetGroupName != nil {
		in, out := &in.DBSubnetGroupName, &out.DBSubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.DBSubnetGroupNameRef != nil {
		in, out := &in.DBSubnetGroupNameRef, &out.DBSubnetGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.DBSubnetGroupNameSelector != nil {
		in, out := &in.DBSubnetGroupNameSelector, &out.DBSubnetGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.EnableCloudwatchLogsExports != nil {
		in, out := &in.EnableCloudwatchLogsExports, &out.EnableCloudwatchLogsExports
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EnableIAMDatabaseAuthentication != nil {
		in, out := &in.EnableIAMDatabaseAuthentication, &out.EnableIAMDatabaseAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.EnablePerformanceInsights != nil {
		in, out := &in.EnablePerformanceInsights, &out.EnablePerformanceInsights
		*out = new(bool)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.IOPS != nil {
		in, out := &in.IOPS, &out.IOPS
		*out = new(int)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.LicenseModel != nil {
		in, out := &in.LicenseModel, &out.LicenseModel
		*out = new(string)
		**out = **in
	}
	if in.MasterUsername != nil {
		in, out := &in.MasterUsername, &out.MasterUsername
		*out = new(string)
		**out = **in
	}
	if in.MasterPasswordSecretRef != nil {
		in, out := &in.MasterPasswordSecretRef, &out.MasterPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.MonitoringInterval != nil {
		in, out := &in.MonitoringInterval, &out.MonitoringInterval
		*out = new(int)
		**out = **in
	}
	if in.MonitoringRoleARN != nil {
		in, out := &in.MonitoringRoleARN, &out.MonitoringRoleARN
		*out = new(string)
		**out = **in
	}
	if in.MonitoringRoleARNRef != nil {
		in, out := &in.MonitoringRoleARNRef, &out.MonitoringRoleARNRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.MonitoringRoleARNSelector != nil {
		in, out := &in.MonitoringRoleARNSelector, &out.MonitoringRoleARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MultiAZ != nil {
		in, out := &in.MultiAZ, &out.MultiAZ
		*out = new(bool)
		**out = **in
	}
	if in.PerformanceInsightsKMSKeyID != nil {
		in, out := &in.PerformanceInsightsKMSKeyID, &out.PerformanceInsightsKMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.PerformanceInsightsRetentionPeriod != nil {
		in, out := &in.PerformanceInsightsRetentionPeriod, &out.PerformanceInsightsRetentionPeriod
		*out = new(int)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.PreferredBackupWindow != nil {
		in, out := &in.PreferredBackupWindow, &out.PreferredBackupWindow
		*out = new(string)
		**out = **in
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.ProcessorFeatures != nil {
		in, out := &in.ProcessorFeatures, &out.ProcessorFeatures
		*out = make([]ProcessorFeature, len(*in))
		copy(*out, *in)
	}
	if in.PromotionTier != nil {
		in, out := &in.PromotionTier, &out.PromotionTier
		*out = new(int)
		**out = **in
	}
	if in.PubliclyAccessible != nil {
		in, out := &in.PubliclyAccessible, &out.PubliclyAccessible
		*out = new(bool)
		**out = **in
	}
	if in.ScalingConfiguration != nil {
		in, out := &in.ScalingConfiguration, &out.ScalingConfiguration
		*out = new(ScalingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageEncrypted != nil {
		in, out := &in.StorageEncrypted, &out.StorageEncrypted
		*out = new(bool)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
	if in.VPCSecurityGroupIDs != nil {
		in, out := &in.VPCSecurityGroupIDs, &out.VPCSecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VPCSecurityGroupIDRefs != nil {
		in, out := &in.VPCSecurityGroupIDRefs, &out.VPCSecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.VPCSecurityGroupIDSelector != nil {
		in, out := &in.VPCSecurityGroupIDSelector, &out.VPCSecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowMajorVersionUpgrade != nil {
		in, out := &in.AllowMajorVersionUpgrade, &out.AllowMajorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.ApplyModificationsImmediately != nil {
		in, out := &in.ApplyModificationsImmediately, &out.ApplyModificationsImmediately
		*out = new(bool)
		**out = **in
	}
	if in.CloudwatchLogsExportConfiguration != nil {
		in, out := &in.CloudwatchLogsExportConfiguration, &out.CloudwatchLogsExportConfiguration
		*out = new(CloudwatchLogsExportConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.DBParameterGroupName != nil {
		in, out := &in.DBParameterGroupName, &out.DBParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.DomainIAMRoleName != nil {
		in, out := &in.DomainIAMRoleName, &out.DomainIAMRoleName
		*out = new(string)
		**out = **in
	}
	if in.DomainIAMRoleNameRef != nil {
		in, out := &in.DomainIAMRoleNameRef, &out.DomainIAMRoleNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.DomainIAMRoleNameSelector != nil {
		in, out := &in.DomainIAMRoleNameSelector, &out.DomainIAMRoleNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.OptionGroupName != nil {
		in, out := &in.OptionGroupName, &out.OptionGroupName
		*out = new(string)
		**out = **in
	}
	if in.UseDefaultProcessorFeatures != nil {
		in, out := &in.UseDefaultProcessorFeatures, &out.UseDefaultProcessorFeatures
		*out = new(bool)
		**out = **in
	}
	if in.SkipFinalSnapshotBeforeDeletion != nil {
		in, out := &in.SkipFinalSnapshotBeforeDeletion, &out.SkipFinalSnapshotBeforeDeletion
		*out = new(bool)
		**out = **in
	}
	if in.FinalDBSnapshotIdentifier != nil {
		in, out := &in.FinalDBSnapshotIdentifier, &out.FinalDBSnapshotIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceParameters.
func (in *RDSInstanceParameters) DeepCopy() *RDSInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceSpec) DeepCopyInto(out *RDSInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceSpec.
func (in *RDSInstanceSpec) DeepCopy() *RDSInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceStatus) DeepCopyInto(out *RDSInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceStatus.
func (in *RDSInstanceStatus) DeepCopy() *RDSInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingConfiguration) DeepCopyInto(out *ScalingConfiguration) {
	*out = *in
	if in.AutoPause != nil {
		in, out := &in.AutoPause, &out.AutoPause
		*out = new(bool)
		**out = **in
	}
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(int)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(int)
		**out = **in
	}
	if in.SecondsUntilAutoPause != nil {
		in, out := &in.SecondsUntilAutoPause, &out.SecondsUntilAutoPause
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingConfiguration.
func (in *ScalingConfiguration) DeepCopy() *ScalingConfiguration {
	if in == nil {
		return nil
	}
	out := new(ScalingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetInRDS) DeepCopyInto(out *SubnetInRDS) {
	*out = *in
	out.SubnetAvailabilityZone = in.SubnetAvailabilityZone
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetInRDS.
func (in *SubnetInRDS) DeepCopy() *SubnetInRDS {
	if in == nil {
		return nil
	}
	out := new(SubnetInRDS)
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
func (in *VPCSecurityGroupMembership) DeepCopyInto(out *VPCSecurityGroupMembership) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCSecurityGroupMembership.
func (in *VPCSecurityGroupMembership) DeepCopy() *VPCSecurityGroupMembership {
	if in == nil {
		return nil
	}
	out := new(VPCSecurityGroupMembership)
	in.DeepCopyInto(out)
	return out
}
