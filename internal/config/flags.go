// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package config

const (
	// DefaultRefreshRate represents the refresh interval.
	DefaultRefreshRate float32 = 2.0 // secs

	// DefaultLogLevel represents the default log level.
	DefaultLogLevel = "info"

	// DefaultCommand represents the default command to run.
	DefaultCommand = ""
)

// Flags represents K9s configuration flags.
type Flags struct {
	// RefreshRate 代表刷新间隔
	RefreshRate *float32
	// LogLevel 代表日志级别
	LogLevel *string
	// LogFile 代表日志文件路径
	LogFile *string
	// Headless 代表是否禁用 head
	Headless *bool
	// Logoless 代表是否禁用logo显示
	Logoless *bool
	//
	Command       *string
	AllNamespaces *bool
	// ReadOnly 代表是否启用只读模式
	ReadOnly *bool
	// Write 代表是否启用写入权限
	Write *bool
	// Crumbsless 代表是否禁用面包屑导航
	Crumbsless *bool
	// Splashless 代表是否禁用启动画面
	Splashless *bool
	// Invert 代表是否启用反转颜色模式
	Invert *bool
	// ScreenDumpDir 代表屏幕截图保存目录
	ScreenDumpDir *string
}

// NewFlags returns new configuration flags.
func NewFlags() *Flags {
	return &Flags{
		// 默认为 2 秒刷新一次
		RefreshRate: float32Ptr(DefaultRefreshRate),
		// 默认为 info 级别日志
		LogLevel: strPtr(DefaultLogLevel),
		//
		LogFile:       strPtr(AppLogFile),
		Headless:      boolPtr(false),
		Logoless:      boolPtr(false),
		Command:       strPtr(DefaultCommand),
		AllNamespaces: boolPtr(false),
		ReadOnly:      boolPtr(false),
		Write:         boolPtr(false),
		Crumbsless:    boolPtr(false),
		Splashless:    boolPtr(false),
		Invert:        boolPtr(false),
		ScreenDumpDir: strPtr(AppDumpsDir),
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func float32Ptr(f float32) *float32 {
	return &f
}

func strPtr(s string) *string {
	return &s
}
