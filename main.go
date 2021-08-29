// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.

package main

import (
	"github.com/koolay/hotrod/cmd"
)

var (
	Version   string
	BuildTime string
)

func main() {
	cmd.Execute()
}
