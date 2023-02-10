package async

import "io"

type options struct {
	logger io.Writer
	debug  bool
}

type OptionFunc func(opt *options)

func Logger(logger io.Writer) OptionFunc {
	return func(opt *options) {
		opt.logger = logger
	}
}

func Debug(debug bool) OptionFunc {
	return func(opt *options) {
		opt.debug = debug
	}
}
