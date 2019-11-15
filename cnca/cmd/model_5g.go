// Copyright 2019 Intel Corporation. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cnca

// SubscribedEvent : A possible value is CHANGE_OF_DNAI - the AF requests to be
// notified when the UP path changes for the PDU isession.
type SubscribedEvent string

// List of SubscribedEvent
const (
	UpPathChange SubscribedEvent = "UP_PATH_CHANGE"
)

// RouteInformation RouteInformation
type RouteInformation struct {
	// string identifying a Ipv4 address formatted in the \"dotted decimal\"
	// notation as defined in IETF RFC 1166.
	IPv4Addr string `json:"ipv4Addr,omitempty"`
	// string identifying a Ipv6 address formatted according to clause 4 in
	// IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of
	// IETF RFC 5952 shall not be used.
	IPv6Addr string `json:"ipv6Addr,omitempty"`
	PortNumber int32 `json:"portNumber"`
}

// RouteToLocation RouteToLocation
type RouteToLocation struct {
	DNAI string `json:"dnai"`
	RouteInfo *RouteInformation `json:"routeInfo,omitempty"`
	RouteProfID string `json:"routeProfId,omitempty"`
}

// Snssai Snssai
type Snssai struct {
	SST int32 `json:"sst"`
	SD string `json:"sd,omitempty"`
}

// WebsockNotifConfig WebsockNotifConfig
type WebsockNotifConfig struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	WebsocketURI string `json:"websocketUri,omitempty"`
	// Set by the AF to indicate that the Websocket delivery is requested.
	RequestWebsocketURI bool `json:"requestWebsocketUri,omitempty"`
}

// TemporalValidity Indicates the time interval(s) during which the AF request is to be applied
type TemporalValidity struct {
	// string with format \"date-time\" as defined in OpenAPI.
	StartTime string `json:"startTime,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	StopTime string `json:"stopTime,omitempty"`
}

// SupportedFeatures SupportedFeatures
type SupportedFeatures struct {
}

// Link string formatted according to IETF RFC 3986 identifying a referenced
// resource.
type Link struct {
}

// FlowInfo FlowInfo
type FlowInfo struct {
	// Indicates the IP flow.
	FlowID int32 `json:"flowId"`
	// Indicates the packet filters of the IP flow. Refer to subclause 5.3.8 of
	// 3GPP TS 29.214 for encoding. It shall contain UL and/or DL IP flow
	// description.
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`
}

// EthFlowDescription Identifies an Ethernet flow
type EthFlowDescription struct {
	DestMacAddr string `json:"destMacAddr,omitempty"`
	EthType string `json:"ethType"`
	// Defines a packet filter of an IP flow.
	FDesc string `json:"fDesc,omitempty"`
	// Possible values are DOWNLINK - The corresponding filter applies for
	// traffic to the UE. UPLINK - The corresponding filter applies for traffic
	// from the UE. - BIDIRECTIONAL The corresponding filter applies for traffic
	// both to and from the UE. UNSPECIFIED - The corresponding filter applies
	// for traffic to the UE (downlink), but has no specific direction declared.
	// The service data flow detection shall apply the filter for uplink traffic
	// as if the filter was bidirectional.
	FDir string `json:"fDir,omitempty"`
	SourceMacAddr string `json:"sourceMacAddr,omitempty"`
	VLANTags []string `json:"vlanTags,omitempty"`
}

// TrafficInfluSub describes Traffic Influence Subscription
type TrafficInfluSub struct {
	// Identifies a service on behalf of which the AF is issuing the request.
	AfServiceID string `json:"afServiceId,omitempty"`
	// Identifies an application.
	AfAppID string `json:"afAppId,omitempty"`
	// Identifies an NEF Northbound interface transaction, generated by the AF.
	AfTransID string `json:"afTransId,omitempty"`
	// Identifies whether an application can be relocated once a location of the application has been selected.
	AppReloInd bool `json:"appReloInd,omitempty"`
	// Identifies data network name
	DNN string `json:"dnn,omitempty"`
	SNSSAI *Snssai `json:"snssai,omitempty"`
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupID string `json:"externalGroupId,omitempty"`
	// Identifies whether the AF request applies to any UE.
	AnyUEInd bool `json:"anyUeInd,omitempty"`
	// Identifies the requirement to be notified of the event(s).
	SubscribedEvents []SubscribedEvent `json:"subscribedEvents,omitempty"`
	GPSI string `json:"gpsi,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	IPv4Addr string `json:"ipv4Addr,omitempty"`
	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952.
	IPv6Addr string `json:"ipv6Addr,omitempty"`
	MACAddr string `json:"macAddr,omitempty"`
	// Identifies the type of notification regarding UP path management event. Possible values are EARLY - early notification of UP path reconfiguration. EARLY_LATE - early and late notification of UP path reconfiguration. This value shall only be present in the subscription to the DNAI change event. LATE - late notification of UP path reconfiguration.
	DNAIChgType string `json:"dnaiChgType,omitempty"`
	NotificationDestination *Link `json:"notificationDestination,omitempty"`
	// Set to true by the AF to request the NEF to send a test notification. Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	Self string `json:"self,omitempty"`
	// Identifies IP packet filters.
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty"`
	// Identifies Ethernet packet filters.
	EthTrafficFilters []EthFlowDescription `json:"ethTrafficFilters,omitempty"`
	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []RouteToLocation `json:"trafficRoutes,omitempty"`
	TempValidities []TemporalValidity `json:"tempValidities,omitempty"`
	// Identifies a geographic zone that the AF request applies only to the traffic of UE(s) located in this specific zone.
	ValidGeoZoneIds []string `json:"validGeoZoneIds,omitempty"`
	SuppFeat *SupportedFeatures `json:"suppFeat,omitempty"`
}

// AfService AfService
type AfService struct {
	AfID string `json:"afId,omitempty"`
	AfInstance string `json:"afInstance,omitempty"`
	LocationServices []LocationService `json:"locationServices,omitempty"`
}

// LocationService LocationService
type LocationService struct {
	DNAI string `json:"dnai,omitempty"`
	DNN string `json:"dnn,omitempty"`
	DNS string `json:"dns,omitempty"`
}
