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
