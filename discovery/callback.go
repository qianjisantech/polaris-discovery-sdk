package discovery

import "polaris-discovery/model"

type CallbackOptions struct {
	OnRegisterSuccess  func(*model.RegisterResponse)
	OnRegisterError    func(error)
	OnHeartbeatSuccess func(*model.HeatBeatResponse)
	OnHeartbeatError   func(error)
}

// Option 函数选项类型
type Option func(*CallbackOptions)

func WithRegisterSuccess(f func(*model.RegisterResponse)) Option {
	return func(opts *CallbackOptions) {
		opts.OnRegisterSuccess = f
	}
}

func WithRegisterError(f func(error)) Option {
	return func(opts *CallbackOptions) {
		opts.OnRegisterError = f
	}
}

func WithHeartbeatSuccess(f func(*model.HeatBeatResponse)) Option {
	return func(opts *CallbackOptions) {
		opts.OnHeartbeatSuccess = f
	}
}

func WithHeartbeatError(f func(error)) Option {
	return func(opts *CallbackOptions) {
		opts.OnHeartbeatError = f
	}
}
