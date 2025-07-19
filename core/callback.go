package core

type CallbackOptions struct {
	OnRegisterSuccess  func(*RegisterResponse)
	OnRegisterError    func(error)
	OnHeartbeatSuccess func(*HeatBeatResponse)
	OnHeartbeatError   func(error)
}

// Option 函数选项类型
type Option func(*CallbackOptions)

func WithRegisterSuccess(f func(*RegisterResponse)) Option {
	return func(opts *CallbackOptions) {
		opts.OnRegisterSuccess = f
	}
}

func WithRegisterError(f func(error)) Option {
	return func(opts *CallbackOptions) {
		opts.OnRegisterError = f
	}
}

func WithHeartbeatSuccess(f func(*HeatBeatResponse)) Option {
	return func(opts *CallbackOptions) {
		opts.OnHeartbeatSuccess = f
	}
}

func WithHeartbeatError(f func(error)) Option {
	return func(opts *CallbackOptions) {
		opts.OnHeartbeatError = f
	}
}
