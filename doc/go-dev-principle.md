# Golang Developing Principle

## copier.Copy仅用于同类型数据拷贝
不同类型的数据拷贝，造成违反开发者习惯了的几个认知失效：
1. `删除所有显式赋值，以删除某个字段`
2. `搜索所有显式引用，以查看问题`
3. ``
4. ``
这些认知失效，进而导致问题产生。

对此的处理：封装copier.Copy，添加src/dest的类型检查，如果是不同的类型则panic；如果Copy返回err，则panic；

不同类型的Copy操作定义为非法，非法copy代码的安全删除的前置条件是：
a.对release分支中非法copy的代码进行标记；
b.目标修改分支合test/uat/release等都必须对比copy操作的类型的成员；
c.不再新增非法copy代码；
d.已有非法copy代码被标记后，新增字段需要显式赋值；

