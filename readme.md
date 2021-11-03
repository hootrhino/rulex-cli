# RULEXC 客户端工具使用文档

## RULEXC 简介
RULEXC 是 RULEX Client 的意思，是命令行下的客户端工具，帮助我们调试和管理RULEX(https://github.com/wwhai/rulex)。

## 构建
```shell
go build -v -o rulexc main.go
```

## 使用文档

- 查看系统参数
```sh
rulexc system
```
- 查看入口列表
```sh
rulexc inends
```
- 查看单个入口
```sh
rulexc inends  '<id>'
```
- 查看出口列表
```sh
rulexc outends
```
- 查看单个出口
```sh
rulexc outends '<id>'
```
- 查看规则列表
```sh
rulexc rules
```
- 查看单个规则
```sh
rulexc rules '<id>'
```
- 查看插件列表
```sh
rulexc plugins
```
- 查看单个插件
```sh
rulexc plugins '<id>'
```
- 创建入口
```sh
rulexc inend-create --config  '<your json format config>'
```
- 创建出口
```sh
rulexc out-create --config  '<your json format config>'
```
