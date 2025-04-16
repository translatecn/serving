/*
Copyright 2018 The Knative Authors

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

package duck

import (
	"knative.dev/serving/pkg/apis/duck/ducktypes"
)

type (
	Implementable = ducktypes.Implementable
	Populatable   = ducktypes.Populatable
)

// VerifyType verifies that a particular concrete resource properly implements
// the provided Implementable duck type.  It is expected that under the resource
// definition implementing a particular "Fooable" that one would write:
//
//	type ConcreteResource struct { ... }
//
//	// Check that ConcreteResource properly implement Fooable.
//	err := duck.VerifyType(&ConcreteResource{}, &something.Fooable{})
//
// This will return an error if the duck typing is not satisfied.

// ConformsToType will return true or false depending on whether a
// concrete resource properly implements the provided Implementable
// duck type.
//
// It will return an error if marshal/unmarshalling fails
