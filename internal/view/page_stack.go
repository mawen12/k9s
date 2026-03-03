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
	if p.app, err = extractApp(ctx); err != nil {
		return err
	}
	p.AddListener(p)

	return nil
}

// StackPushed notifies a new page was added.
// StackPushed 通知一个新页面被添加
func (p *PageStack) StackPushed(c model.Component) {
	c.Start()
	p.app.SetFocus(c)
}

// StackPopped notifies a page was removed.
// StackPopped 通知一个页面被移除
func (p *PageStack) StackPopped(o, top model.Component) {
	o.Stop()
	p.StackTop(top)
}

// StackTop notifies for the top component.
// StackTop 通知顶部组件
func (p *PageStack) StackTop(top model.Component) {
	if top == nil {
		return
	}
	top.Start()
	p.app.SetFocus(top)
}
