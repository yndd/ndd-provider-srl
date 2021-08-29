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

import (
	kerrors "k8s.io/apimachinery/pkg/api/errors"
)

const (
	errTrackTCUsage      = "cannot track TargetConfig usage"
	errGetTC             = "cannot get TargetConfig"
	errGetNetworkNode    = "cannot get NetworkNode"
	errNewClient         = "cannot create new client"
	targetNotConfigured  = "target is not configured to proceed"
	errNoTargetFound     = "target not found"
	errJSONMarshal       = "cannot marshal JSON object"
	errJSONUnMarshal     = "cannot unmarshal JSON object"
	errJSONCompare       = "cannot compare JSON objects"
	errJSONMarshalIndent = "cannot marshal JSON object with indent"
	errUpdateObject      = "cannot update object"
	errWrongInputdata    = "wrong input data"
)

// An ErrorIs function returns true if an error satisfies a particular condition.
type ErrorIs func(err error) bool

// Ignore any errors that satisfy the supplied ErrorIs function by returning
// nil. Errors that do not satisfy the supplied function are returned unmodified.
func Ignore(is ErrorIs, err error) error {
	if is(err) {
		return nil
	}
	return err
}

// IgnoreNotFound returns the supplied error, or nil if the error indicates a
// Kubernetes resource was not found.
func IgnoreNotFound(err error) error {
	return Ignore(kerrors.IsNotFound, err)
}
