# Key

`Key` 是 tview 中用于表示键盘输入的类型。
每个 tview 组件的 `SetInputCapture` 方法都会接收用来处理事件的函数。

## key 注册

在 k9s 中，`internal/view` 包封装了 `internal/ui` 包中的组件，
并在 `bindKeys()` 中来初始化 `actions`。
**internal/ui** 中存有 `actions`。

`bindKeys()` 是在 `Init(ctx)` 函数中完成，这是代表 `internal/view` 包中的组件被初始化的函数。

## key 使用

当在组件中按下键盘时，tview 会调用 `SetInputCapture` => `keyboard(View)` => `actions(UI)`。
然后通过 `keyboard(event)` 来读取 `actions` 中的 `key` 绑定的事件来执行对应的函数。

## 拓展

`Init(ctx)` 一般用于组件的初始化，因为在 `New` 时不会传递 `app`，因此部分行为没办法进行。
而在 `Init` 中可以通过 `ctx` 来获取 `app`，因此此时可以进行一些需要 `app` 的操作。

`Init` 一般会做如下几件事：
- 通过 `keyboard` 注册事件执行器到 tview 上
- 通过 `bindKeys` 来注册事件到 `actions`
- 注册组件的监听器
