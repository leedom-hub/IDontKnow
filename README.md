## 自动认证项目结构
### 介绍
免费开源工具，仅供学习，请勿做违法违规的事情，一切后果作者概不负责。
### 环境
- go1.17.5
### 目录
```shell
├── config
├── core
├── global
├── log
└── utils
```

| 文件夹 | 说明 | 描述 |
| ------ | ------ | ------|
| `config`  | 配置包以及结构体文件  | config.yaml对应的配置结构体 |
| `core`       | 核心文件 | 核心组件的初始化 |
| `global`     | 全局对象 | 全局对象 |
| `log`     | 日志存放目录  | 日志 |
| `utils`      | 工具包  | 工具函数封装  |     

### 配置
`config`文件夹下的 `config.yaml` 文件需要根据实际需要填写;
```
env:
    debug: false //http debug
    out-put-count: true //是否输出倒计时
    proxy: false //是否代理
```
```
url:
  local-ip: 192.168.37.142 //本机IP
  ...
```
```
user:
  pwd: 1234 //用户密码
  user-name: 小红 //用户名
```
***其他配置均不需要修改***
### 使用
```
1.配置文件修改
2.运行rz.exe文件
```