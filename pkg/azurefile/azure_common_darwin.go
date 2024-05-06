//go:build darwin
// +build darwin

/*
Copyright 2020 The Kubernetes Authors.

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

package azurefile

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mount "k8s.io/mount-utils"
)

func SMBMount(m *mount.SafeFormatAndMount, source, target, fsType string, options, sensitiveMountOptions []string) error {
	return nil
}

func SMBUnmount(m *mount.SafeFormatAndMount, target string, _, _ bool) error {
	return nil
}

func CleanupMountPoint(m *mount.SafeFormatAndMount, target string, extensiveMountCheck bool) error {
	return nil
}

func preparePublishPath(path string, m *mount.SafeFormatAndMount) error {
	return nil
}

func prepareStagePath(path string, m *mount.SafeFormatAndMount) error {
	return nil
}

// GetVolumeStats returns volume stats based on the given path.
func GetVolumeStats(path string, enableWindowsHostProcess bool) (*csi.NodeGetVolumeStatsResponse, error) {
	return nil, status.Errorf(codes.Internal, "GetVolumeStats is not supported on darwin")
}
