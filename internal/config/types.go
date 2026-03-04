// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package config

const (
	defaultRefreshRate  = 2
	defaultMaxConnRetry = 5

	// CPU tracks cpu usage.
	CPU = "cpu"

	// MEM tracks memory usage.
	MEM = "memory"
)

// UI tracks ui specific configs.
type UI struct {
	// EnableMouse toggles mouse support.
	// EnableMouse 切换鼠标支持
	EnableMouse bool `json:"enableMouse" yaml:"enableMouse"`

	// Headless toggles top header display.
	// Headless 切换顶部标题显示
	Headless bool `json:"headless" yaml:"headless"`

	// LogoLess toggles k9s logo.
	// LogoLess 切换 k9s logo
	Logoless bool `json:"logoless" yaml:"logoless"`

	// Crumbsless toggles nav crumb display.
	// Crumbsless 切换导航面包屑显示
	Crumbsless bool `json:"crumbsless" yaml:"crumbsless"`

	// Splashless disables the splash screen on startup.
	// Splashless 禁用启动时的启动页面
	Splashless bool `json:"splashless" yaml:"splashless"`

	// Reactive toggles reactive ui changes.
	// Reactive 切换反应式 UI 更改，如果开启了该配置，则会在后台监视配置文件、样式等的更改，并在发生更改时自动更新 UI
	Reactive bool `json:"reactive" yaml:"reactive"`

	// NoIcons toggles icons display.
	// NoIcons 切换图标显示
	NoIcons bool `json:"noIcons" yaml:"noIcons"`

	// Invert inverts all skin colors using Oklch lightness inversion.
	// Invert 使用 Oklch 亮度反转来反转所有皮肤颜色
	Invert bool `json:"invert" yaml:"invert"`

	// Skin reference the general k9s skin name.
	// Can be overridden per context.
	// Skin 引用一般的 k9s 皮肤名称。可以在每个上下文中覆盖
	Skin string `json:"skin" yaml:"skin,omitempty"`

	// DefaultsToFullScreen toggles fullscreen on views like logs, yaml, details.
	// DefaultsToFullScreen 切换日志、yaml、详细信息等视图的全屏显示
	DefaultsToFullScreen bool `json:"defaultsToFullScreen" yaml:"defaultsToFullScreen"`

	// UseFullGVRTitle toggles the display of full GVR (group/version/resource) vs R in views title.
	// UseFullGVRTitle 切换视图标题中显示完整的 GVR（组/版本/资源）与 R
	UseFullGVRTitle bool `json:"useFullGVRTitle" yaml:"useFullGVRTitle"`

	manualHeadless   *bool
	manualLogoless   *bool
	manualCrumbsless *bool
	manualSplashless *bool
	manualInvert     *bool
}
