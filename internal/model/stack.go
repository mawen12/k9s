// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package model

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/derailed/k9s/internal/slogs"
)

const (
	// StackPush denotes an add on the stack.
	StackPush StackAction = 1 << iota

	// StackPop denotes a delete on the stack.
	StackPop
)

// StackAction represents an action on the stack.
type StackAction int

// StackEvent represents an operation on a view stack.
type StackEvent struct {
	// Kind represents the event condition.
	Action StackAction

	// Item represents the targeted item.
	Component Component
}

// StackListener represents a stack listener.
// StackListener 代表一个栈监听器
type StackListener interface {
	// StackPushed indicates a new item was added.
	// StackPushed 有一个新元素入栈
	StackPushed(Component)

	// StackPopped indicates an item was deleted
	// StackPopped 有一个元素出栈，最新的栈顶
	StackPopped(old, new Component)

	// StackTop indicates the top of the stack
	// StackTop 当前栈的顶部元素，在监听器首次注册时调用
	StackTop(Component)
}

// Stack represents a stacks of components.

// Stack 代表堆结构的组件
type Stack struct {
	// 基于数组实现堆接口，最后的元素就是 Top
	components []Component
	// 当有 Pop/Push 操作时，触发通知操作
	listeners []StackListener
	// 同步读写保护，允许并发读，当操作 components 时，必须要先获得锁
	mx sync.RWMutex
}

// NewStack returns a new initialized stack.
func NewStack() *Stack {
	return &Stack{}
}

// Flatten returns a string representation of the component stack.
// Flatten 返回字符串代表的组件栈
func (s *Stack) Flatten() []string {
	s.mx.RLock()
	defer s.mx.RUnlock()

	ss := make([]string, len(s.components))
	for i, c := range s.components {
		// 获取组件名称
		ss[i] = c.Name()
	}
	return ss
}

// RemoveListener removes a listener.
// RemoveListener 移除监听器
func (s *Stack) RemoveListener(l StackListener) {
	victim := -1
	for i, lis := range s.listeners {
		if lis == l {
			victim = i
			break
		}
	}
	if victim == -1 {
		return
	}
	s.listeners = append(s.listeners[:victim], s.listeners[victim+1:]...)
}

// AddListener registers a stack listener.
// AddListener 注册一个栈调用器
func (s *Stack) AddListener(l StackListener) {
	s.listeners = append(s.listeners, l)
	if !s.Empty() { // 当栈中有元素时，触发通知
		l.StackTop(s.Top())
	}
}

// Push adds a new item.

// Push 入栈
func (s *Stack) Push(c Component) {
	// 停止当前的 Top
	if top := s.Top(); top != nil {
		top.Stop()
	}

	s.mx.Lock()
	s.components = append(s.components, c)
	s.mx.Unlock()
	s.notify(StackPush, c)
}

// Pop removed the top item and returns it.
// Pop 弹出一个元素，bool 代表此次操作是否成功
func (s *Stack) Pop() (Component, bool) {
	if s.Empty() {
		return nil, false
	}

	var c Component
	s.mx.Lock()
	// 读取 Top
	c = s.components[len(s.components)-1]
	c.Stop()
	// 裁剪 Slice
	s.components = s.components[:len(s.components)-1]
	s.mx.Unlock()

	// 调用通知
	s.notify(StackPop, c)

	return c, true
}

// Peek returns stack state.
// Peek 返回堆元素
func (s *Stack) Peek() []Component {
	s.mx.RLock()
	defer s.mx.RUnlock()

	return s.components
}

// Clear clear out the stack using pops.
// Clear 使用 Pop 将所有元素出栈
func (s *Stack) Clear() {
	for range s.components {
		s.Pop()
	}
}

// Empty returns true if the stack is empty.
// Empty 堆为空时，返回 true
func (s *Stack) Empty() bool {
	s.mx.RLock()
	defer s.mx.RUnlock()

	return len(s.components) == 0
}

// IsLast indicates if stack only has one item left.
// IsLast 检查堆中是否只有一个元素，
// TODO FIX 此处应该加上读写锁
func (s *Stack) IsLast() bool {
	return len(s.components) == 1
}

// Previous returns the previous component if any.
// Previous 返回前一个堆元素，如果只有一个，则返回该元素
// TODO FIX 此处应该加上读写锁
func (s *Stack) Previous() Component {
	if s.Empty() {
		return nil
	}
	if s.IsLast() {
		return s.Top()
	}

	return s.components[len(s.components)-2]
}

// Top returns the top most item or nil if the stack is empty.
// Top 返回最顶部的元素
// TODO FIX 此处应该加上读写锁
func (s *Stack) Top() Component {
	if s.Empty() {
		return nil
	}

	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.components[len(s.components)-1]
}

// notify 按操作类型通知
func (s *Stack) notify(a StackAction, c Component) {
	for _, l := range s.listeners {
		switch a {
		case StackPush:
			l.StackPushed(c)
		case StackPop:
			// TODO FIX s.Top() 优化为读取一次，而不是每次调用时获取
			l.StackPopped(c, s.Top())
		}
	}
}

// ----------------------------------------------------------------------------
// Helpers...

// Dump prints out the stack.
func (s *Stack) Dump() {
	slog.Debug("Stack Dump", slogs.Stack, fmt.Sprintf("%p", s))
	for i, c := range s.components {
		slog.Debug(fmt.Sprintf("%d -- %s -- %#v", i, c.Name(), c))
	}
	slog.Debug("------------------")
}
