// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package model

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/derailed/k9s/internal/slogs"
)

const (
	// DefaultFlashDelay sets the flash clear delay.
	// DefaultFlashDelay 设置 flash 的默认延迟，6s
	DefaultFlashDelay = 6 * time.Second

	// FlashInfo represents an info message.
	// FlashInfo 代表Info消息
	FlashInfo FlashLevel = iota
	// FlashWarn represents an warning message.
	// FlashWarn 代表Warn消息
	FlashWarn
	// FlashErr represents an error message.
	// FlashErr 代表错误消息
	FlashErr
)

// LevelMessage tracks a message and severity.
// LevelMessage 跟踪消息和严重性
type LevelMessage struct {
	Level FlashLevel
	Text  string
}

func newClearMessage() LevelMessage {
	return LevelMessage{}
}

// IsClear returns true if message is empty.
// IsClear 返回消息是否为空
func (l LevelMessage) IsClear() bool {
	return l.Text == ""
}

// FlashLevel represents flash message severity.
type FlashLevel int

// FlashChan represents a flash event channel.
type FlashChan chan LevelMessage

// FlashListener represents a text model listener.
type FlashListener interface {
	// FlashChanged notifies the model changed.
	// FlashChanged 当模型变化时通知
	FlashChanged(FlashLevel, string)

	// FlashCleared notifies when the filter changed.
	// FlashCleared 当过滤器变化时通知
	FlashCleared()
}

// Flash represents a flash message model.
// Flash 代表一个快闪消息模型
type Flash struct {
	msg     LevelMessage
	cancel  context.CancelFunc
	delay   time.Duration
	msgChan chan LevelMessage
}

// NewFlash returns a new instance.
// NewFlash 返回一个新的示例
func NewFlash(dur time.Duration) *Flash {
	return &Flash{
		delay:   dur,
		msgChan: make(FlashChan, 3),
	}
}

// Channel returns the flash channel.
// Channel 返回 flash channel
func (f *Flash) Channel() FlashChan {
	return f.msgChan
}

// Info displays an info flash message.
func (f *Flash) Info(msg string) {
	f.SetMessage(FlashInfo, msg)
}

// Infof displays a formatted info flash message.
func (f *Flash) Infof(fmat string, args ...any) {
	f.Info(fmt.Sprintf(fmat, args...))
}

// Warn displays a warning flash message.
func (f *Flash) Warn(msg string) {
	slog.Warn(msg)
	f.SetMessage(FlashWarn, msg)
}

// Warnf displays a formatted warning flash message.
func (f *Flash) Warnf(fmat string, args ...any) {
	f.Warn(fmt.Sprintf(fmat, args...))
}

// Err displays an error flash message.
func (f *Flash) Err(err error) {
	slog.Error("Flash error", slogs.Error, err)
	f.SetMessage(FlashErr, err.Error())
}

// Errf displays a formatted error flash message.
func (f *Flash) Errf(fmat string, args ...any) {
	var err error
	for _, a := range args {
		if e, ok := a.(error); ok {
			err = e
		}
	}
	slog.Error("Flash error",
		slogs.Error, err,
		slogs.Message, fmt.Sprintf(fmat, args...),
	)
	f.SetMessage(FlashErr, fmt.Sprintf(fmat, args...))
}

// Clear clears the flash message.
// Clear 清理消息，使用一条空消息覆盖
func (f *Flash) Clear() {
	f.fireCleared()
}

// SetMessage sets the flash level message.
// SetMessage 设置快闪消息
func (f *Flash) SetMessage(level FlashLevel, msg string) {
	// 取消
	if f.cancel != nil {
		f.cancel()
		f.cancel = nil
	}

	// 设置消息
	f.setLevelMessage(LevelMessage{Level: level, Text: msg})
	// 将消息推送到 channel
	f.fireFlashChanged()

	// 构造支持取消的 ctx
	ctx := context.Background()
	ctx, f.cancel = context.WithCancel(ctx)
	// 异步刷新，不阻塞
	go f.refresh(ctx)
}

// refresh 刷新，手动取消或者超时后退出
func (f *Flash) refresh(ctx context.Context) {
	// TODO FIX 移除 for 循环
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(f.delay):
			f.fireCleared()
			return
		}
	}
}

// setLevelMessage 更新消息
func (f *Flash) setLevelMessage(msg LevelMessage) {
	f.msg = msg
}

// fireFlashChanged 发送一条快闪消息
func (f *Flash) fireFlashChanged() {
	f.msgChan <- f.msg
}

// fireCleared 发送一条消息
func (f *Flash) fireCleared() {
	f.msgChan <- newClearMessage()
}
