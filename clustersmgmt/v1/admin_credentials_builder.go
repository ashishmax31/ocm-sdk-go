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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// AdminCredentialsBuilder contains the data and logic needed to build 'admin_credentials' objects.
//
// Temporary administrator credentials generated during the installation of the
// cluster.
type AdminCredentialsBuilder struct {
	user     *string
	password *string
}

// NewAdminCredentials creates a new builder of 'admin_credentials' objects.
func NewAdminCredentials() *AdminCredentialsBuilder {
	return new(AdminCredentialsBuilder)
}

// User sets the value of the 'user' attribute
// to the given value.
//
//
func (b *AdminCredentialsBuilder) User(value string) *AdminCredentialsBuilder {
	b.user = &value
	return b
}

// Password sets the value of the 'password' attribute
// to the given value.
//
//
func (b *AdminCredentialsBuilder) Password(value string) *AdminCredentialsBuilder {
	b.password = &value
	return b
}

// Build creates a 'admin_credentials' object using the configuration stored in the builder.
func (b *AdminCredentialsBuilder) Build() (object *AdminCredentials, err error) {
	object = new(AdminCredentials)
	if b.user != nil {
		object.user = b.user
	}
	if b.password != nil {
		object.password = b.password
	}
	return
}
