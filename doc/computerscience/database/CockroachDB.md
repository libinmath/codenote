使用 CockroachDB 启用下一代多区域应用程序

# 摘要

数据库服务需要满足服务于全球用户群的现代应用程序的一致性、性能和可用性目标。配置跨多个区域部署的数据库以实现这些目标需要大量的专业知识。

在本文中，我们描述了 CockroachDB 如何通过提供`高级声明性语法`来使开发人员轻松实现这一点，该语法允许通过 SQL 语句表达数据访问局部性和可用性目标。然后将这些高级目标映射到数据库配置、副本放置和数据分区决策。我们展示了如何增强数据库的所有层（从 SQL 优化器到复制）以支持多区域工作负载。我们还描述了一种新的事务管理协议，该协议支持从任何数据库副本进行本地、强一致的读取。

最后，本文进行了广泛的评估，证明 CockroachDB 用于多区域集群的新声明式 SQL 语法易于使用，并支持具有不同性能权衡的各种配置选项，以使各种工作负载受益。我们还表明，吞吐量与区域数量呈线性关系，并且与之前的方法相比，新的事务管理协议将尾部延迟减少了 10 倍以上。


# 参考资料
https://www.cockroachlabs.com/guides/cockroachdb-the-resilient-geo-distributed-sql-database-sigmod-2022/