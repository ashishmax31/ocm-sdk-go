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

// OpenIdurlsListBuilder contains the data and logic needed to build
// 'open_idurls' objects.
type OpenIdurlsListBuilder struct {
	items []*OpenIdurlsBuilder
}

// NewOpenIdurlsList creates a new builder of 'open_idurls' objects.
func NewOpenIdurlsList() *OpenIdurlsListBuilder {
	return new(OpenIdurlsListBuilder)
}

// Items sets the items of the list.
func (b *OpenIdurlsListBuilder) Items(values ...*OpenIdurlsBuilder) *OpenIdurlsListBuilder {
	b.items = make([]*OpenIdurlsBuilder, len(values))
	copy(b.items, values)
	return b
}

// Build creates a list of 'open_idurls' objects using the
// configuration stored in the builder.
func (b *OpenIdurlsListBuilder) Build() (list *OpenIdurlsList, err error) {
	items := make([]*OpenIdurls, len(b.items))
	for i, item := range b.items {
		items[i], err = item.Build()
		if err != nil {
			return
		}
	}
	list = new(OpenIdurlsList)
	list.items = items
	return
}
