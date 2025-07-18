package constant

// TaskStatus 定义任务状态类型
type TaskStatus string

// 任务状态枚举值
const (
	TaskStatusPending  TaskStatus = "pending"  // 待定
	TaskStatusRunning  TaskStatus = "running"  // 运行中
	TaskStatusSuccess  TaskStatus = "success"  // 成功
	TaskStatusFailed   TaskStatus = "failed"   // 失败
	TaskStatusCanceled TaskStatus = "canceled" // 已取消
	TaskStatusTimeout  TaskStatus = "timeout"  // 超时
	TaskStatusSkipped  TaskStatus = "skipped"  // 跳过
	TaskStatusAborted  TaskStatus = "aborted"  // 中止
	TaskStatusWaiting  TaskStatus = "waiting"  // 等待
	TaskStatusPaused   TaskStatus = "paused"   // 暂停
)
