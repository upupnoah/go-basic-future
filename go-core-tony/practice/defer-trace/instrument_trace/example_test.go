package trace_test

import (
	trace "github.com/upupnoah/go-basic-future/go-core/practice/defer-trace/instrument_trace"
)

func a() {
	defer trace.Trace()()
	b()
}

func b() {
	defer trace.Trace()()
	c()
}

func c() {
	defer trace.Trace()()
	d()
}

func d() {
	defer trace.Trace()
}

func ExampleTrace() {
	a()
	// 提供预期输出， 预期输出从Output开始
	// Output:
	// g[00001]:    ->github.com/upupnoah/instrument_trace_test.a
	// g[00001]:        ->github.com/upupnoah/instrument_trace_test.b
	// g[00001]:            ->github.com/upupnoah/instrument_trace_test.c
	// g[00001]:                ->github.com/upupnoah/instrument_trace_test.d
	// g[00001]:                <-github.com/upupnoah/instrument_trace_test.d
	// g[00001]:            <-github.com/upupnoah/instrument_trace_test.c
	// g[00001]:        <-github.com/upupnoah/instrument_trace_test.b
	// g[00001]:    <-github.com/upupnoah/instrument_trace_test.a
}
