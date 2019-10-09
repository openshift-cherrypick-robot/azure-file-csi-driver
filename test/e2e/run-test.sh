#!/bin/bash

# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -uo pipefail

function cleanup {
  deploy/uninstall-driver.sh
}

trap cleanup EXIT

deploy/install-driver.sh

# Fetching ginkgo for running the test
GO111MODULE=off go get github.com/onsi/ginkgo/ginkgo
export KUBECONFIG="$HOME/.kube/config"
"$GOBIN"/ginkgo test/e2e
exit "$?"