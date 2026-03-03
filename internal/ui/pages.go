// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package ui

import (
	"fmt"
	"log/slog"

	"github.com/derailed/k9s/internal/model"
	"github.com/derailed/k9s/internal/slogs"
	"github.com/derailed/tview"
)

// Pages represents a stack of view pages.

// Pages 代表视图页面的堆
// 它是 Primitive/Focusable/StackListener 的实现
type Pages struct {
	*tview.Pages
	*model.Stack
}

// NewPages return a new view.
func NewPages() *Pages {
	p := Pages{
		Pages: tview.NewPages(),
		Stack: model.NewStack(),
	}
	p.AddListener(&p)

	return &p
}

// IsTopDialog checks if front page is a dialog.

// IsTopDialog 检查首页是否为 Dialog
func (p *Pages) IsTopDialog() bool {
	// 获取 FrontPage
	_, pa := p.GetFrontPage()
	switch pa.(type) {
	// 如果是 ModalFrom 或者 ModalList，则为 Dialog
	case *tview.ModalForm, *ModalList:
		return true
	default:
		return false
	}
}

// Show displays a given page.
// Show 展示指定组件
func (p *Pages) Show(c model.Component) {
	p.SwitchToPage(componentID(c))
}

// Current returns the current component.
// Current 返回当前页面
func (p *Pages) Current() model.Component {
	c := p.CurrentPage()
	if c == nil {
		return nil
	}

	return c.Item.(model.Component)
}

// AddAndShow adds a new page and bring it to front.
// AddAndShow 添加一个新页面，并将其放到最前
func (p *Pages) addAndShow(c model.Component) {
	p.add(c)
	p.Show(c)
}

// Add adds a new page.
// Add 添加一个新页面
func (p *Pages) add(c model.Component) {
	p.AddPage(componentID(c), c, true, true)
}

// Delete removes a page.
func (p *Pages) delete(c model.Component) {
	p.RemovePage(componentID(c))
}

// Dump for debug.
func (p *Pages) Dump() {
	slog.Debug("Dumping Pages", slogs.Page, p)
	for i, c := range p.Peek() {
		slog.Debug(fmt.Sprintf("%d -- %s -- %#v", i, componentID(c), p.GetPrimitive(componentID(c))))
	}
}

// Stack Protocol...

// StackPushed notifies a new component was pushed.
// StackPushed 添加并展示
func (p *Pages) StackPushed(c model.Component) {
	p.addAndShow(c)
}

// StackPopped notifies a component was removed.
// StackPopped 移除
func (p *Pages) StackPopped(o, _ model.Component) {
	p.delete(o)
}

// StackTop notifies a new component is at the top of the stack.
// StackTop 展示页面
func (p *Pages) StackTop(top model.Component) {
	if top == nil {
		return
	}
	p.Show(top)
}

// Helpers...

// componentID 返回组件标识
func componentID(c model.Component) string {
	if c.Name() == "" {
		slog.Error("Component has no name", slogs.Component, fmt.Sprintf("%T", c))
	}
	return fmt.Sprintf("%s-%p", c.Name(), c)
}
