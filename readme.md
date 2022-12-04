# RULEX 命令行工具

## RULEX-CLI 简介
RULEX-CLI 是 RULEX 命令行下的客户端工具，帮助我们调试和管理RULEX(https://github.com/i4de/rulex)。

## 构建
```shell
go build -v
```

## 使用文档
```bash
------------------------------------------------------------------------------
|*| 查看系统参数: rulexc system-info -host 127.0.0.1
|*| 查看输入资源列表: rulexc inend-list -host 127.0.0.1
|*| 查看输出资源列表: rulexc outend-list -host 127.0.0.1
|*| 查看设备列表: rulexc device-list -host 127.0.0.1
|*| 查看规则列表: rulexc rules-list -host 127.0.0.1
|*| 创建输入资源: rulexc inend-create --config '[config]' -host 127.0.0.1
|*| 创建输出资源: rulexc outend-create --config  '[config]' -host 127.0.0.1
|*| 创建规则脚本: rulexc rules-create --config  '[config]' -host 127.0.0.1
|*| 创建设备: rulexc device-create --config  '[config]' -host 127.0.0.1
------------------------------------------------------------------------------
```