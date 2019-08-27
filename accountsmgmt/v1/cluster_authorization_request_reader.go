/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// clusterAuthorizationRequestData is the data structure used internally to marshal and unmarshal
// objects of type 'cluster_authorization_request'.
type clusterAuthorizationRequestData struct {
	ClusterID        *string                  "json:\"cluster_id,omitempty\""
	AccountUsername  *string                  "json:\"account_username,omitempty\""
	Managed          *bool                    "json:\"managed,omitempty\""
	Reserve          *bool                    "json:\"reserve,omitempty\""
	BYOC             *bool                    "json:\"byoc,omitempty\""
	AvailabilityZone *string                  "json:\"availability_zone,omitempty\""
	Resources        reservedResourceListData "json:\"resources,omitempty\""
}

// MarshalClusterAuthorizationRequest writes a value of the 'cluster_authorization_request' to the given target,
// which can be a writer or a JSON encoder.
func MarshalClusterAuthorizationRequest(object *ClusterAuthorizationRequest, target interface{}) error {
	encoder, err := helpers.NewEncoder(target)
	if err != nil {
		return err
	}
	data, err := object.wrap()
	if err != nil {
		return err
	}
	return encoder.Encode(data)
}

// wrap is the method used internally to convert a value of the 'cluster_authorization_request'
// value to a JSON document.
func (o *ClusterAuthorizationRequest) wrap() (data *clusterAuthorizationRequestData, err error) {
	if o == nil {
		return
	}
	data = new(clusterAuthorizationRequestData)
	data.ClusterID = o.clusterID
	data.AccountUsername = o.accountUsername
	data.Managed = o.managed
	data.Reserve = o.reserve
	data.BYOC = o.byoc
	data.AvailabilityZone = o.availabilityZone
	data.Resources, err = o.resources.wrap()
	if err != nil {
		return
	}
	return
}

// UnmarshalClusterAuthorizationRequest reads a value of the 'cluster_authorization_request' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalClusterAuthorizationRequest(source interface{}) (object *ClusterAuthorizationRequest, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(clusterAuthorizationRequestData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'cluster_authorization_request' type.
func (d *clusterAuthorizationRequestData) unwrap() (object *ClusterAuthorizationRequest, err error) {
	if d == nil {
		return
	}
	object = new(ClusterAuthorizationRequest)
	object.clusterID = d.ClusterID
	object.accountUsername = d.AccountUsername
	object.managed = d.Managed
	object.reserve = d.Reserve
	object.byoc = d.BYOC
	object.availabilityZone = d.AvailabilityZone
	object.resources, err = d.Resources.unwrap()
	if err != nil {
		return
	}
	return
}
