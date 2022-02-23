# RULEX 命令行工具

## RULEX-CLI 简介
RULEX-CLI 是 RULEX 命令行下的客户端工具，帮助我们调试和管理RULEX(https://github.com/i4de/rulex)。

## 构建
```shell
go build -v
```

## 使用文档
```bash
*> rulex-cli system-info -host 127.0.0.1
*> rulex-cli inend-list -host 127.0.0.1
*> rulex-cli outend-list -host 127.0.0.1
*> rulex-cli rules-list -host 127.0.0.1
*> rulex-cli inend-create --config '[config]' -host 127.0.0.1
*> rulex-cli outend-create --config  '[config]' -host 127.0.0.1
*> rulex-cli rules-create --config  '[config]' -host 127.0.0.1
```