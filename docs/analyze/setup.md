# Setup

## Shell

**输出当前版本**

```shell
k9s version
```

**获取k9s信息**

```shell
k9s info
```

**打印帮助信息**

```shell
k9s help
```

**使用给定 namespace 运行K9s**

```shell
k9s -n mycoolns
```

## Cmd

使用 `github.com/spf13/cobra` 来实现命令

## View

在 `internal/ui/app.go` 中定义视图入口。

在 `internal/view/app.go` 中作为 Application。

