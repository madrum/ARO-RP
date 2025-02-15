package subnets

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"testing"

	mgmtnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	arov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	mock_subnet "github.com/Azure/ARO-RP/pkg/util/mocks/subnet"
	"github.com/Azure/ARO-RP/pkg/util/subnet"
)

var (
	subscriptionId           = "0000000-0000-0000-0000-000000000000"
	clusterResourceGroupName = "aro-iljrzb5a"
	infraId                  = "abcd"
	clusterResourceGroupId   = "/subscriptions/" + subscriptionId + "/resourcegroups/" + clusterResourceGroupName
	vnetResourceGroup        = "vnet-rg"
	vnetName                 = "vnet"
	subnetNameWorker         = "worker"
	subnetNameMaster         = "master"

	nsgv1NodeResourceId   = clusterResourceGroupId + "/providers/Microsoft.Network/networkSecurityGroups/" + infraId + subnet.NSGNodeSuffixV1
	nsgv1MasterResourceId = clusterResourceGroupId + "/providers/Microsoft.Network/networkSecurityGroups/" + infraId + subnet.NSGControlPlaneSuffixV1
	nsgv2ResourceId       = clusterResourceGroupId + "/providers/Microsoft.Network/networkSecurityGroups/" + infraId + subnet.NSGSuffixV2
)

func getValidClusterInstance() *arov1alpha1.Cluster {
	return &arov1alpha1.Cluster{
		Spec: arov1alpha1.ClusterSpec{
			ArchitectureVersion:    0,
			ClusterResourceGroupID: clusterResourceGroupId,
			InfraID:                infraId,
		},
	}
}

func getValidSubnet() *mgmtnetwork.Subnet {
	return &mgmtnetwork.Subnet{
		SubnetPropertiesFormat: &mgmtnetwork.SubnetPropertiesFormat{
			NetworkSecurityGroup: &mgmtnetwork.SecurityGroup{
				ID: to.StringPtr(nsgv1MasterResourceId),
			},
		},
	}
}

func TestReconcileManager(t *testing.T) {
	log := logrus.NewEntry(logrus.StandardLogger())

	for _, tt := range []struct {
		name       string
		subnetMock func(*mock_subnet.MockManager, *mock_subnet.MockKubeManager)
		instance   func(*arov1alpha1.Cluster)
		wantErr    error
	}{
		{
			name: "Architecture V1 - no change",
			subnetMock: func(mock *mock_subnet.MockManager, kmock *mock_subnet.MockKubeManager) {
				resourceIdMaster := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameMaster
				resourceIdWorker := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameWorker

				kmock.EXPECT().List(gomock.Any()).Return([]subnet.Subnet{
					{
						ResourceID: resourceIdMaster,
						IsMaster:   true,
					},
					{
						ResourceID: resourceIdWorker,
						IsMaster:   false,
					},
				}, nil)

				subnetObjectMaster := getValidSubnet()
				mock.EXPECT().Get(gomock.Any(), resourceIdMaster).Return(subnetObjectMaster, nil)

				subnetObjectWorker := getValidSubnet()
				subnetObjectWorker.NetworkSecurityGroup.ID = to.StringPtr(nsgv1NodeResourceId)
				mock.EXPECT().Get(gomock.Any(), resourceIdWorker).Return(subnetObjectWorker, nil)
			},
		},
		{
			name: "Architecture V1 - all fixup",
			subnetMock: func(mock *mock_subnet.MockManager, kmock *mock_subnet.MockKubeManager) {

				resourceIdMaster := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameMaster
				resourceIdWorker := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameWorker

				kmock.EXPECT().List(gomock.Any()).Return([]subnet.Subnet{
					{
						ResourceID: resourceIdMaster,
						IsMaster:   true,
					},
					{
						ResourceID: resourceIdWorker,
						IsMaster:   false,
					},
				}, nil)

				subnetObjectMaster := getValidSubnet()
				subnetObjectMaster.NetworkSecurityGroup.ID = to.StringPtr(nsgv1MasterResourceId + "new")
				mock.EXPECT().Get(gomock.Any(), resourceIdMaster).Return(subnetObjectMaster, nil)

				subnetObjectMasterUpdate := getValidSubnet()
				subnetObjectMasterUpdate.NetworkSecurityGroup.ID = to.StringPtr(nsgv1MasterResourceId)
				mock.EXPECT().CreateOrUpdate(gomock.Any(), resourceIdMaster, subnetObjectMasterUpdate).Return(nil)

				subnetObjectWorker := getValidSubnet()
				subnetObjectWorker.NetworkSecurityGroup.ID = to.StringPtr(nsgv1NodeResourceId + "new")
				mock.EXPECT().Get(gomock.Any(), resourceIdWorker).Return(subnetObjectWorker, nil)

				subnetObjectWorkerUpdate := getValidSubnet()
				subnetObjectWorkerUpdate.NetworkSecurityGroup.ID = to.StringPtr(nsgv1NodeResourceId)
				mock.EXPECT().CreateOrUpdate(gomock.Any(), resourceIdWorker, subnetObjectWorkerUpdate).Return(nil)
			},
		},
		{
			name: "Architecture V1 - node only fixup",
			subnetMock: func(mock *mock_subnet.MockManager, kmock *mock_subnet.MockKubeManager) {

				resourceIdMaster := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameMaster
				resourceIdWorker := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameWorker

				kmock.EXPECT().List(gomock.Any()).Return([]subnet.Subnet{
					{
						ResourceID: resourceIdMaster,
						IsMaster:   true,
					},
					{
						ResourceID: resourceIdWorker,
						IsMaster:   false,
					},
				}, nil)

				subnetObjectMaster := getValidSubnet()
				mock.EXPECT().Get(gomock.Any(), resourceIdMaster).Return(subnetObjectMaster, nil)

				subnetObjectWorker := getValidSubnet()
				subnetObjectWorker.NetworkSecurityGroup.ID = to.StringPtr(nsgv1NodeResourceId + "new")
				mock.EXPECT().Get(gomock.Any(), resourceIdWorker).Return(subnetObjectWorker, nil)

				subnetObjectWorkerUpdate := getValidSubnet()
				subnetObjectWorkerUpdate.NetworkSecurityGroup.ID = to.StringPtr(nsgv1NodeResourceId)
				mock.EXPECT().CreateOrUpdate(gomock.Any(), resourceIdWorker, subnetObjectWorkerUpdate).Return(nil)
			},
		},
		{
			name: "Architecture V2 - no fixups",
			subnetMock: func(mock *mock_subnet.MockManager, kmock *mock_subnet.MockKubeManager) {

				resourceIdMaster := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameMaster
				resourceIdWorker := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameWorker

				kmock.EXPECT().List(gomock.Any()).Return([]subnet.Subnet{
					{
						ResourceID: resourceIdMaster,
						IsMaster:   true,
					},
					{
						ResourceID: resourceIdWorker,
						IsMaster:   false,
					},
				}, nil)

				subnetObjectMaster := getValidSubnet()
				subnetObjectMaster.NetworkSecurityGroup.ID = to.StringPtr(nsgv2ResourceId)
				mock.EXPECT().Get(gomock.Any(), resourceIdMaster).Return(subnetObjectMaster, nil)

				subnetObjectWorker := getValidSubnet()
				subnetObjectWorker.NetworkSecurityGroup.ID = to.StringPtr(nsgv2ResourceId)
				mock.EXPECT().Get(gomock.Any(), resourceIdWorker).Return(subnetObjectWorker, nil)
			},
			instance: func(instace *arov1alpha1.Cluster) {
				instace.Spec.ArchitectureVersion = int(api.ArchitectureVersionV2)
			},
		},
		{
			name: "Architecture V2 - all nodes fixup",
			subnetMock: func(mock *mock_subnet.MockManager, kmock *mock_subnet.MockKubeManager) {

				resourceIdMaster := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameMaster
				resourceIdWorker := "/subscriptions/" + subscriptionId + "/resourceGroups/" + vnetResourceGroup + "/providers/Microsoft.Network/virtualNetworks/" + vnetName + "/subnets/" + subnetNameWorker

				kmock.EXPECT().List(gomock.Any()).Return([]subnet.Subnet{
					{
						ResourceID: resourceIdMaster,
						IsMaster:   true,
					},
					{
						ResourceID: resourceIdWorker,
						IsMaster:   false,
					},
				}, nil)

				subnetObjectMaster := getValidSubnet()
				subnetObjectMaster.NetworkSecurityGroup.ID = to.StringPtr(nsgv2ResourceId + "new")
				mock.EXPECT().Get(gomock.Any(), resourceIdMaster).Return(subnetObjectMaster, nil)

				subnetObjectMasterUpdate := getValidSubnet()
				subnetObjectMasterUpdate.NetworkSecurityGroup.ID = to.StringPtr(nsgv2ResourceId)
				mock.EXPECT().CreateOrUpdate(gomock.Any(), resourceIdMaster, subnetObjectMasterUpdate).Return(nil)

				subnetObjectWorker := getValidSubnet()
				subnetObjectWorker.NetworkSecurityGroup.ID = to.StringPtr(nsgv2ResourceId + "new")
				mock.EXPECT().Get(gomock.Any(), resourceIdWorker).Return(subnetObjectWorker, nil)

				subnetObjectWorkerUpdate := getValidSubnet()
				subnetObjectWorkerUpdate.NetworkSecurityGroup.ID = to.StringPtr(nsgv2ResourceId)
				mock.EXPECT().CreateOrUpdate(gomock.Any(), resourceIdWorker, subnetObjectWorkerUpdate).Return(nil)
			},
			instance: func(instace *arov1alpha1.Cluster) {
				instace.Spec.ArchitectureVersion = int(api.ArchitectureVersionV2)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			subnets := mock_subnet.NewMockManager(controller)
			kubeSubnets := mock_subnet.NewMockKubeManager(controller)
			if tt.subnetMock != nil {
				tt.subnetMock(subnets, kubeSubnets)
			}

			instance := getValidClusterInstance()
			if tt.instance != nil {
				tt.instance(instance)
			}

			r := reconcileManager{
				log:            log,
				instance:       instance,
				subscriptionID: subscriptionId,
				subnets:        subnets,
				kubeSubnets:    kubeSubnets,
			}

			err := r.reconcileSubnets(context.Background())
			if err != nil {
				if tt.wantErr == nil {
					t.Fatal(err)
				}
				if err != nil && err.Error() != tt.wantErr.Error() || err == nil && tt.wantErr != nil {
					t.Errorf("Expected Error %s, got %s when processing %s testcase", tt.wantErr.Error(), err.Error(), tt.name)
				}
			}
			// we don't need to compare as mock should do the job

		})
	}
}
