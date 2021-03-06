// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package versions

import (
	"net/http"
	"strings"

	"github.com/aws/amazon-ec2-metadata-mock/pkg/server"
)

var (
	supportedVersions = []string{"latest"}
)

// Handler handles http requests
func Handler(res http.ResponseWriter, req *http.Request) {
	// Unregistered paths should return 404
	if req.URL.Path != "/" {
		server.ReturnNotFoundResponse(res)
		return
	}
	server.FormatAndReturnTextResponse(res, strings.Join(supportedVersions, "\n"))
	return
}
