<<<<<<< HEAD
Golet
----
Go 代码片段
# Purpose
1. 减少重复造"轮子"等零件，也即提供最具扩展性、可用性的零件
2. 做零件的仓库，也即提供各种广泛使用的零件
3. 零件的"售后"和升级，也即提供零件的版本管理、安全增强、性能增强等升级

使用 Golet，可以带来：
+ 快速初始化 gitlab工程
+ 低代码开发
+ 减少Copy代码，降低维护成本

# Quick Start
```sh
go get github.com/sinngae/golet
```


# Contribute to Golet
+ 没有代码实现，仅有需求，可以在 https://github.com/sinngae/golet/-/issues/new 新建Gitlab Issue, label标记为Proposal
+ 有代码实现，可以提交 Merge Request
    + 请提供充分必要的测试用例
    + 代码不能import具体的业务项目工程gitlab
=======
Golet
----
Go 代码片段
# 目标 Purpose
1. 减少重复造"轮子"等零件，也即提供最具扩展性、可用性的零件
2. 做零件的仓库，也即提供各种广泛使用的零件
3. 零件的"售后"和升级，也即提供零件的版本管理、安全增强、性能增强等升级

使用 Golet，可以带来：
+ 快速初始化 gitlab工程
+ 低代码开发
+ 减少Copy代码，降低维护成本

# 快速访问 Quick Start
```sh
# 获取和使用
go get github.com/sinngae/golet
```

## 目录指南 directory leading
+ src -- 源代码包，存放通用功能的方法代码
+ cmd -- 集成包，存放集成的通用功能命令

# Contribute to Golet
+ 没有代码实现，仅有需求，可以在 https://github.com/sinngae/golet/-/issues/new 新建Gitlab Issue, label标记为Proposal
+ 有代码实现，可以提交 Merge Request
    + 请提供充分必要的测试用例
    + 代码不能import具体的业务项目工程gitlab

# Docker For Golet
## win
```shell
# build image
docker build --network=host -v .:/build -t golet:base .
# build golet
## docker run --entrypoint "" --rm -it --network=host --privileged=true -v $(CURDIR):/build ${DOCKER_IMAGE} /bin/bash
## test
docker run --entrypoint "" --rm -it --network=host --privileged=true -v .:/build centos-base:master /bin/bash
```
>>>>>>> 4aad3e1b64427d5ebafb07f037b140f7eb3a6511
