// Copyright 2023 Forerunner Labs, Inc.
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

package authz

import (
	"time"

	object "github.com/warrant-dev/warrant/pkg/authz/object"
	objecttype "github.com/warrant-dev/warrant/pkg/authz/objecttype"
)

type FeatureSpec struct {
	FeatureId   string    `json:"featureId" validate:"required,valid_object_id"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (spec FeatureSpec) ToFeature(objectId int64) *Feature {
	return &Feature{
		ObjectId:    objectId,
		FeatureId:   spec.FeatureId,
		Name:        spec.Name,
		Description: spec.Description,
	}
}

func (spec FeatureSpec) ToCreateObjectSpec() *object.CreateObjectSpec {
	return &object.CreateObjectSpec{
		ObjectType: objecttype.ObjectTypeFeature,
		ObjectId:   spec.FeatureId,
	}
}

type UpdateFeatureSpec struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
