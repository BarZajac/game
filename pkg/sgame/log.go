package sgame

type LogFn func(log string, items ...interface{})

var Glog LogFn
