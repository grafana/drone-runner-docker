// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package resource

import (
	"errors"
	"fmt"

	"github.com/drone/runner-go/manifest"
)

// Lookup returns the named pipeline from the Manifest.
func Lookup(name string, manifest *manifest.Manifest) (manifest.Resource, error) {
	for _, resource := range manifest.Resources {
		if !isNameMatch(resource.GetName(), name) {
			continue
		}
		fmt.Println("Lookup1:", resource)
		if pipeline, ok := resource.(*Pipeline); ok {
			fmt.Println("Lookup2:", pipeline)
			return pipeline, nil
		}
	}
	return nil, errors.New("resource not found")
}

// helper function returns true if the name matches.
func isNameMatch(a, b string) bool {
	return a == b ||
		(a == "" && b == "default") ||
		(b == "" && a == "default")
}
