# RULEX 命令行工具

## RULEX-CLI 简介
RULEX-CLI 是 RULEX 命令行下的客户端工具，帮助我们调试和管理RULEX。

## 构建
```shell
go build -v
```

## 使用文档

- 查看系统参数: `rulexc system-info -host 127.0.0.1`
- 查看输入资源: `rulexc inend-list -host 127.0.0.1`
- 查看输出资源: `rulexc outend-list -host 127.0.0.1`
- 创建设备:     `rulexc device-create --config '[config]' -host 127.0.0.1`
- 查看设备列表: `rulexc device-list -host 127.0.0.1`
- 查看规则列表: `rulexc rules-list -host 127.0.0.1`
- 创建输入资源: `rulexc inend-create --config '[config]' -host 127.0.0.1`
- 创建输出资源: `rulexc outend-create --config  '[config]' -host 127.0.0.1`
- 创建规则脚本:  `rulexc rules-create --config  '[config]' -host 127.0.0.1`
- 创建用户:     `rulexc user-create --u [username] --p [password] -host 127.0.0.1`
- 删除用户:     `rulexc user-delete --u [username] -host 127.0.0.1`

## 示例
获取用户列表:`rulexc user-list -host 127.0.0.1 |jq`
```json
{
  "code": 200,
  "msg": "Success",
  "data": [
    {
      "role": "admin",
      "username": "admin",
      "description": "admin"
    },
    {
      "role": "admin",
      "username": "www",
      "description": ""
    }
  ]
}
```