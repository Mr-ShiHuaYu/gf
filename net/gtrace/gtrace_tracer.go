// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtrace

import (
	"github.com/Mr-ShiHuaYu/otel-go111"
	"github.com/Mr-ShiHuaYu/otel-go111/trace"
)

// Tracer warps trace.Tracer for compatibility and extension.
type Tracer struct {
	trace.Tracer
}

// NewTracer Tracer is a short function for retrieving Tracer.
func NewTracer(name ...string) *Tracer {
	tracerName := ""
	if len(name) > 0 {
		tracerName = name[0]
	}
	return &Tracer{
		Tracer: otel.Tracer(tracerName),
	}
}
