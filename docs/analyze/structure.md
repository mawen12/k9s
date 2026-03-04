# Structure

## 目录

```markdown
k9s/
├── cmd  # 命令行入口/
│   ├── root.go  # CLI 根命令定义
│   ├── version.go  # 版本命令实现
│   └── info.go  # 信息命令实现
└── internal # 内部实现/
├── config  # 配置管理
├── client  # 资源管理
├── model  # 资源模型抽象
├── ui  # UI 应用整体框架和布局
├── view  # UI 视图，具体资源视图和内容展示
└── controller 操作控制器
```

## 流程

### 启动流程

```markdown
main.go/
└──> cmd/root.go/
    └──> internal/ui/app.go/    # 初始化 UI
        ├──> internal/config/   # 加载配置
        ├──> internal/client/   # 连接K8S
        └──> internal/view/table.go # 渲染视图
```

### 资源操作流程

```markdown
用户按键/
└── internal/view/table.go/ # 捕获键盘事件
└── internal/controller/    # 处理操作逻辑
└── internal/model/ # 执行 k8s API
└── K8S APIServer # 实际操作
```

### 配置加载流程

```markdown
```

## 层次

```markdown
UI 层（internal/ui, internal/view）
        ↓
控制器层（internal/controller）
        ↓
模型层（internal/model）
        ↓
客户端层（internal/client）
        ↓
Kubernetes API
```