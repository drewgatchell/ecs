// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

// The Search fields describe information about a search request event: query
// or pagination. The fields that should be used with this field set include:
// `event.action` to describe the search action (e.g. `search.query`,
// `search.page`, etc.), `event.duration` to describe the duration of a search
// request, `@timestamp` to record the event's original timestamp and
// optionally the `source` fields to record context information such as
// `user.id` or `geo`.
type Search struct {
	// An opaque query identifier. This identifier needs to be unique to a user
	// query, and all subsequent events (pagination, clicks) need to have the
	// same query identifier.
	QueryID string `ecs:"query.id"`

	// The query string being searched on. This field is not analyzed and
	// should not be pre-processed in any way in the event (e.g. normalization
	// list lowercasing). This is useful for search use-cases that use a one-
	// box style search interface. Other interfaces will have to rely on
	// additional custom fields or labels to represent things like filters
	// applied, extra parameters, user context, etc.
	QueryValue string `ecs:"query.value"`

	// For search results that support pagination, this represents the current
	// page being requested. Initial search requests are `1` while subsequent
	// page requests are incremental.
	QueryPage int64 `ecs:"query.page"`

	// The size of the result set displayed to the user. This should be
	// equivalent to the length of the results in `results.ids`. This is also
	// known as the page size or limit.
	ResultsSize int64 `ecs:"results.size"`

	// The total number of matches for this query. This number is always
	// greater than or equal to `results.size`. This is the `hits.total` field
	// in the query response.
	ResultsTotal int64 `ecs:"results.total"`

	// A list of opaque document IDs representing the results that were shown
	// to the user. This is effectively the impression list and it's size
	// should be equal to `results.size`. This field can be empty when there
	// are no results to return.
	ResultsIds string `ecs:"results.ids"`
}
