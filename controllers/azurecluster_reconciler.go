/*
Copyright 2019 The Kubernetes Authors.

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

package controllers

import (
	"context"

	"github.com/pkg/errors"
	"sigs.k8s.io/cluster-api-provider-azure/azure"
	"sigs.k8s.io/cluster-api-provider-azure/azure/scope"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/bastionhosts"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/groups"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/loadbalancers"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/natgateways"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/privatedns"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/privateendpoints"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/publicips"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/resourceskus"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/routetables"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/securitygroups"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/subnets"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/tags"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/virtualnetworks"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/vnetpeerings"
	"sigs.k8s.io/cluster-api-provider-azure/util/tele"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// azureClusterService is the reconciler called by the AzureCluster controller.
type azureClusterService struct {
	scope *scope.ClusterScope
	// services is the list of services that are reconciled by this controller.
	// The order of the services is important as it determines the order in which the services are reconciled.
	services []azure.ServiceReconciler
	skuCache *resourceskus.Cache
}

// newAzureClusterService populates all the services based on input scope.
func newAzureClusterService(scope *scope.ClusterScope) (*azureClusterService, error) {
	skuCache, err := resourceskus.GetCache(scope, scope.Location())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a NewCache")
	}
	return &azureClusterService{
		scope: scope,
		services: []azure.ServiceReconciler{
			groups.New(scope),
			virtualnetworks.New(scope),
			securitygroups.New(scope),
			routetables.New(scope),
			publicips.New(scope),
			natgateways.New(scope),
			subnets.New(scope),
			vnetpeerings.New(scope),
			loadbalancers.New(scope),
			privatedns.New(scope),
			bastionhosts.New(scope),
			privateendpoints.New(scope),
			tags.New(scope),
		},
		skuCache: skuCache,
	}, nil
}

// Reconcile reconciles all the services in a predetermined order.
func (s *azureClusterService) Reconcile(ctx context.Context) error {
	ctx, _, done := tele.StartSpanWithLogger(ctx, "controllers.azureClusterService.Reconcile")
	defer done()

	if err := s.setFailureDomainsForLocation(ctx); err != nil {
		return errors.Wrap(err, "failed to get availability zones")
	}

	s.scope.SetDNSName()
	s.scope.SetControlPlaneSecurityRules()

	for _, service := range s.services {
		if err := service.Reconcile(ctx); err != nil {
			return errors.Wrapf(err, "failed to reconcile AzureCluster service %s", service.Name())
		}
	}

	return nil
}

// Delete reconciles all the services in a predetermined order.
func (s *azureClusterService) Delete(ctx context.Context) error {
	ctx, _, done := tele.StartSpanWithLogger(ctx, "controllers.azureClusterService.Delete")
	defer done()

	groupSvc, err := s.getService(groups.ServiceName)
	if err != nil {
		return errors.Wrap(err, "failed to get group service")
	}

	managed, err := groupSvc.IsManaged(ctx)
	if err != nil {
		if azure.ResourceNotFound(err) {
			// If the resource group is not found, there is nothing to delete, return early.
			return nil
		}
		return errors.Wrap(err, "failed to determine if the AzureCluster resource group is managed")
	}
	if managed {
		// If the resource group is managed, delete it.
		// We need to explicitly delete vnet peerings, as it is not part of the resource group.
		vnetPeeringsSvc, err := s.getService(vnetpeerings.ServiceName)
		if err != nil {
			return errors.Wrap(err, "failed to get vnet peerings service")
		}
		if err := vnetPeeringsSvc.Delete(ctx); err != nil {
			return errors.Wrap(err, "failed to delete peerings")
		}
		// Delete the entire resource group directly.
		if err := groupSvc.Delete(ctx); err != nil {
			return errors.Wrap(err, "failed to delete resource group")
		}
	} else {
		// If the resource group is not managed we need to delete resources inside the group one by one.
		// services are deleted in reverse order from the order in which they are reconciled.
		for i := len(s.services) - 1; i >= 0; i-- {
			if err := s.services[i].Delete(ctx); err != nil {
				return errors.Wrapf(err, "failed to delete AzureCluster service %s", s.services[i].Name())
			}
		}
	}

	return nil
}

func (s *azureClusterService) getService(name string) (azure.ServiceReconciler, error) {
	for _, service := range s.services {
		if service.Name() == name {
			return service, nil
		}
	}
	return nil, errors.Errorf("service %s not found", name)
}

// setFailureDomainsForLocation sets the AzureCluster Status failure domains based on which Azure Availability Zones are available in the cluster location.
// Note that this is not done in a webhook as it requires API calls to fetch the availability zones.
func (s *azureClusterService) setFailureDomainsForLocation(ctx context.Context) error {
	zones, err := s.skuCache.GetZones(ctx, s.scope.Location())
	if err != nil {
		return errors.Wrapf(err, "failed to get zones for location %s", s.scope.Location())
	}

	for _, zone := range zones {
		s.scope.SetFailureDomain(zone, clusterv1.FailureDomainSpec{
			ControlPlane: true,
		})
	}

	return nil
}
