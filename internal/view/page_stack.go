// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package view

import (
	"context"

	"github.com/derailed/k9s/internal/model"
	"github.com/derailed/k9s/internal/ui"
)

// PageStack represents a stack of pages.
// PageStack 代表一个页面堆栈
type PageStack struct {
	*ui.Pages

	app *App
}

// NewPageStack returns a new page stack.
// NewPageStack 返回一个新的页面堆栈
func NewPageStack() *PageStack {
	return &PageStack{
		Pages: ui.NewPages(),
	}
}

// Init initializes the view.
// Init 初始化视图
func (p *PageStack) Init(ctx context.Context) (err error) {
	// 从 ctx 中提取 App 示例并赋值给 p.app
	if p.app, err = extractApp(ctx); err != nil {
		return err
	}
	// 将 PageStack 本身添加为 Pages 的监听器，以便在页面堆栈发生变化时接收通知
	p.AddListener(p)

	return nil
}

// StackPushed notifies a new page was added.
// StackPushed 有一个组件被添加时，通知它开始并将焦点设置到这个组件上
func (p *PageStack) StackPushed(c model.Component) {
	c.Start()
	p.app.SetFocus(c)
}

// StackPopped notifies a page was removed.
// StackPopped 有一个组件被移除时，通知它停止并将焦点设置到新的顶部组件上
func (p *PageStack) StackPopped(o, top model.Component) {
	o.Stop()
	p.StackTop(top)
}

// StackTop notifies for the top component.
// StackTop 通知顶部组件，如果顶部组件不为 nil，则通知它开始并将焦点设置到它上面
func (p *PageStack) StackTop(top model.Component) {
	if top == nil {
		return
	}
	top.Start()
	p.app.SetFocus(top)
}
