// Copyright 2021 FerretDB Inc.
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

package shareddata

import "go.mongodb.org/mongo-driver/bson"

// Composites contain composite values for tests.
//
// This shared data set is not frozen yet, but please add to it only if it is really shared.
var Composites = &Values[string]{
	data: map[string]any{
		"document":           bson.D{{"foo", int32(42)}},
		"document-composite": bson.D{{"foo", int32(42)}, {"42", "foo"}, {"array", bson.A{int32(42), "foo", nil}}},
		"document-empty":     bson.D{},

		"array":          bson.A{int32(42)},
		"array-three":    bson.A{int32(42), "foo", nil},
		"array-embedded": bson.A{bson.A{int32(42), "foo"}, nil},
		"array-empty":    bson.A{},
	},
}