package collector

import (
	"fmt"
	"strings"

	"github.com/karimra/gnmic/utils"
	"github.com/openconfig/gnmi/proto/gnmi"
)

func CreateSubscriptionRequest(target, subName string, paths []*gnmi.Path) (*gnmi.SubscribeRequest, error) {
	// create subscription

	gnmiPrefix, err := utils.CreatePrefix(subName, target)
	if err != nil {
		return nil, fmt.Errorf("create prefix failed")
	}
	modeVal := gnmi.SubscriptionList_Mode_value[strings.ToUpper("STREAM")]
	qos := &gnmi.QOSMarking{Marking: 21}

	subscriptions := make([]*gnmi.Subscription, len(paths))
	for i, p := range paths {
		subscriptions[i] = &gnmi.Subscription{Path: p}
		switch gnmi.SubscriptionList_Mode(modeVal) {
		case gnmi.SubscriptionList_STREAM:
			mode := gnmi.SubscriptionMode_value[strings.Replace(strings.ToUpper("ON_CHANGE"), "-", "_", -1)]
			subscriptions[i].Mode = gnmi.SubscriptionMode(mode)
		}
	}
	req := &gnmi.SubscribeRequest{
		Request: &gnmi.SubscribeRequest_Subscribe{
			Subscribe: &gnmi.SubscriptionList{
				Prefix:       gnmiPrefix,
				Mode:         gnmi.SubscriptionList_Mode(modeVal),
				Encoding:     46, // "JSON_IETF_CONFIG_ONLY"
				Subscription: subscriptions,
				Qos:          qos,
			},
		},
	}
	return req, nil
}
