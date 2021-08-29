/*
	Copyright 2021 Wim Henderickx.

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

package srl

/*
import (
	"fmt"

	srlv1 "github.com/netw-device-driver/ndd-provider-srl/apis/srl/v1"
	"github.com/netw-device-driver/ndd-runtime/pkg/resource"
	"github.com/pkg/errors"
)

func getManagedResource(resourceName string) (resource.Managed, error) {
	fmt.Printf("getManagedResource: %s\n", resourceName)
	var mg resource.Managed
	switch resourceName {
	case "SrlBfd":
		mg = &srlv1.SrlBfd{}
	case "SrlInterface":
		mg = &srlv1.SrlInterface{}
	case "SrlInterfaceSubinterface":
		mg = &srlv1.SrlInterfaceSubinterface{}
	case "SrlNetworkinstance":
		mg = &srlv1.SrlNetworkinstance{}
	case "SrlNetworkinstanceAggregateroutes":
		mg = &srlv1.SrlNetworkinstanceAggregateroutes{}
	case "SrlNetworkinstanceNexthopgroups":
		mg = &srlv1.SrlNetworkinstanceNexthopgroups{}
	case "SrlNetworkinstanceProtocolsBgp":
		mg = &srlv1.SrlNetworkinstanceProtocolsBgp{}
	case "SrlNetworkinstanceProtocolsBgpevpn":
		mg = &srlv1.SrlNetworkinstanceProtocolsBgpevpn{}
	case "SrlNetworkinstanceProtocolsBgpvpn":
		mg = &srlv1.SrlNetworkinstanceProtocolsBgpvpn{}
	case "SrlNetworkinstanceProtocolsIsis":
		mg = &srlv1.SrlNetworkinstanceProtocolsIsis{}
	case "SrlNetworkinstanceProtocolsLinux":
		mg = &srlv1.SrlNetworkinstanceProtocolsLinux{}
	case "SrlNetworkinstanceProtocolsOspf":
		mg = &srlv1.SrlNetworkinstanceProtocolsOspf{}
	case "SrlNetworkinstanceStaticroutes":
		mg = &srlv1.SrlNetworkinstanceStaticroutes{}
	case "SrlRoutingpolicyAspathset":
		mg = &srlv1.SrlRoutingpolicyAspathset{}
	case "SrlRoutingpolicyCommunityset":
		mg = &srlv1.SrlRoutingpolicyCommunityset{}
	case "SrlRoutingpolicyPolicy":
		mg = &srlv1.SrlRoutingpolicyPolicy{}
	case "SrlRoutingpolicyPrefixset":
		mg = &srlv1.SrlRoutingpolicyPrefixset{}
	case "SrlSystemMtu":
		mg = &srlv1.SrlSystemMtu{}
	case "SrlSystemName":
		mg = &srlv1.SrlSystemName{}
	case "SrlSystemNetworkinstanceProtocolsBgpvpn":
		mg = &srlv1.SrlSystemNetworkinstanceProtocolsBgpvpn{}
	case "SrlSystemNetworkinstanceProtocolsEvpn":
		mg = &srlv1.SrlSystemNetworkinstanceProtocolsEvpn{}
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance":
		mg = &srlv1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstance{}
	case "SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi":
		mg = &srlv1.SrlSystemNetworkinstanceProtocolsEvpnEsisBgpinstanceEsi{}
	case "SrlSystemNtp":
		mg = &srlv1.SrlSystemNtp{}
	case "SrlTunnelinterface":
		mg = &srlv1.SrlTunnelinterface{}
	case "SrlTunnelinterfaceVxlaninterface":
		mg = &srlv1.SrlTunnelinterfaceVxlaninterface{}
	default:
		return nil, errors.New(fmt.Sprintf("cannot find resource object: %s", resourceName))
	}
	return mg, nil
}

//{Name: "statement", Key: map[string]string{"sequence-id": ""}},
*/
