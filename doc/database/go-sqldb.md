"Go SQL DB" 是一个研究目的的支持 SQL 查询的关系型数据库。主要的目标是为了向数据库爱好者展示一个关系型数据库的基本原理和关键设计。因此，为了便于理解，采取了很多取巧但不是很严谨的设计，代码量控制在了 2000 行左右（包含约 400 行单元测试代码）。

https://github.com/auxten/go-sqldb

# 特性列表
纯 Golang 实现，不依赖任何第三方包。仅在单元测试中引入了 goconvey
单元测试覆盖率≈ 73.5%
# 存储引擎
基于 B+Tree 的数据检索结构
基于 4KB 分页的磁盘持久化引擎
接近 POD（Plain Old Data）的序列化 & 反序列化
# SQL Parser
Tokenizer 基于 text/scanner 实现
支持简单的 SELECT、INSERT 语法
SELECT 支持数值类型的 WHERE
支持 LIMIT，但暂不支持 ORDER BY
如果你想要了解可以生产可用的 SQL Parser 是如何实现的，请参考我从 CockroachDB 剥离的 SQL-2011 标准支持的 SQL Parser
# 执行计划 Planner
基于火山模型（Volcano Model）的 Select 实现
基于 HTTP 的查询和插入接口