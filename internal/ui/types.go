// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package ui

import (
	"context"
	"time"

	"github.com/derailed/k9s/internal/config"
	"github.com/derailed/k9s/internal/dao"
	"github.com/derailed/k9s/internal/model"
	"github.com/derailed/k9s/internal/model1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	unlockedIC = "[RW]"
	lockedIC   = "[R]"
)

// Namespaceable tracks namespaces.
// Namespaceable 可对命名空间进行操作
type Namespaceable interface {
	// ClusterWide returns true if the model represents resource in all namespaces.
	ClusterWide() bool

	// GetNamespace returns the model namespace.
	GetNamespace() string

	// SetNamespace changes the model namespace.
	SetNamespace(string)

	// InNamespace check if current namespace matches models.
	InNamespace(string) bool
}

// Lister tracks resource getter.
// Lister 用于根据路径获取资源
type Lister interface {
	// Get returns a resource instance.
	// Get 返回一个资源实例
	Get(ctx context.Context, path string) (runtime.Object, error)
}

// Tabular represents a tabular model.
// Tabular 代表一个表格模型
type Tabular interface {
	Namespaceable // 可以操作命名空间
	Lister        //

	// SetInstance sets parent resource path.
	// SetInstance 设置父级资源路径
	SetInstance(string)

	// SetLabelSelector sets the label selector.
	// SetLabelSelector 设置标签选择器
	SetLabelSelector(labels.Selector)

	// GetLabelSelector fetch the label filter.
	// GetLabelSelector 获取标签选择器
	GetLabelSelector() labels.Selector

	// Empty returns true if model has no data.
	// Empty 模型为空时返回
	Empty() bool

	// RowCount returns the model data count.
	// RowCount 返回模型数据总数
	RowCount() int

	// Peek returns current model data.
	// Peek 返回当前模型数据
	Peek() *model1.TableData

	// Watch watches a given resource for changes.
	// Watch 观察一个给定资源的变更
	Watch(context.Context) error

	// Refresh forces a new refresh.
	// Refresh 强制刷新
	Refresh(context.Context) error

	// SetRefreshRate sets the model watch loop rate.
	// SetRefreshRate 设置刷新频率
	SetRefreshRate(time.Duration)

	// AddListener registers a model listener.
	// AddListener 注册一个模型监听器
	AddListener(model.TableListener)

	// RemoveListener unregister a model listener.
	// RemoveListener 移除一个模型监听器
	RemoveListener(model.TableListener)

	// Delete a resource.
	// Delete 删除一个资源
	Delete(context.Context, string, *metav1.DeletionPropagation, dao.Grace) error

	// SetViewSetting injects custom cols specification.
	// SetViewSetting 注入自定义的列规范
	SetViewSetting(context.Context, *config.ViewSetting)
}
