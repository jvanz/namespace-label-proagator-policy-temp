//go:build wasi
// +build wasi

package main

import (
	capabilities "github.com/kubewarden/policy-sdk-go/pkg/capabilities"
)

func getWapcHost() capabilities.Host {
	return capabilities.NewHost()
}
