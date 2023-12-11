/*
Copyright 2022 The Kubernetes Authors.

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

package converters

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	"k8s.io/utils/ptr"
)

// GetSubnetAddresses returns the address prefixes contained in an SDK v2 subnet.
func GetSubnetAddresses(subnet *armnetwork.Subnet) []string {
	var addresses []string
	if subnet.Properties != nil && subnet.Properties.AddressPrefix != nil {
		addresses = []string{ptr.Deref(subnet.Properties.AddressPrefix, "")}
	} else if subnet.Properties != nil && subnet.Properties.AddressPrefixes != nil {
		for _, address := range subnet.Properties.AddressPrefixes {
			if address != nil {
				addresses = append(addresses, *address)
			}
		}
	}
	return addresses
}