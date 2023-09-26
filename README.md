# go-zero-oj

## 技术栈:

后端:

- go-zero
- docker
- redis
- rabbitmq

## 运行脚本：

**goctl生成代码**

请在cmd/api目录下使用


window:
``` shell
  goctl api go --api .\core.api --dir ./ --style=goZero
```

linux:
``` shell
```

## 主要功能：

- [ ] 用户模块
    - [x] 用户登录
    - [x] 用户登录
    - [ ] 邮箱登录
    - [x] 用户详细信息
    - [x] 用户答题排序
- [ ] 问题模块
    - [x] 问题列表
    - [x] 问题查询(根据唯一标识查询)
    - [x] 问题提交记录
    - [ ] 创建问题(管理员)
    

