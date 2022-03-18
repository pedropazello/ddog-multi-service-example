package datadog_trace_service

import (
	ddtracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Start() {
	ddtracer.Start(ddtracer.WithServiceName("checkout_service"), ddtracer.WithEnv("env"))
}
