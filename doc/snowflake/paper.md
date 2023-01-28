[TOC]

# 摘要
- 传统数据仓库：
1. 为固定资源而设计
2. 依赖复杂的ETL（Extract-Transform-Load）流水线和物理调优
- 云数据仓库：
1. 弹性能力
2. 能处理半结构数据并且快速调优

关键特性
- 极致的弹性和可用性
- 半结构化和schema-less的数据支持
- 时间旅行
- 端到端的安全性

# 1. 简介
过去，数据仓库的数据源：
1. 事务系统
2. 企业资源规划（ERP）
3. 客户关系管理（CRM）
数据的结构、数量和速率都是可预测和可掌握的。

外部来源（经常以schema-less、半结构化的格式存储）：
1. 应用程序日志
2. web应用程序
3. 移动设备
4. 社交媒体
5. 传感器数据（物联网）

部分数仓——>Hadoop或Spark等大数据平台
仍然缺乏现有数仓的效率和功能。

Snowflake并不是基于Hadoop、PostgreSQL之类的，而是重新开发。




Snowflake的文章，介绍了整体的架构，包括计算存储分离，计算层的弹性等等，代表了数据库发展的最终形态。

Snowflake 16年的论文，Snowflake是云计算领域的明星公司，底层存储使用了AWS的S3，virtual warehous(计算层)使用的AWS的EC2,依靠云上资源实现多租户，计算存储弹性，高可用，安全等特性，支持SQL，半结构化以及无模式数据,Snowflake提供了企业级云上数仓解决方案。

## Snowflake关键特点

1. SaaS(Software-as-a-Service)级服务体验。
用户无需购买机器、雇佣数据库管理员或安装软件。用户基于自身系统在云上的数据，借助于Snowflake的图形界面或ODBC等标准化接口就可以立即操作和查询数据。
2. 关系型。
支持ANSI SQL和ACID事务
3. 半结构化
支持Semi-Structured 和shcema-less data，保存为JSON或者Avro，支持列存储
4. 弹性
计算存储分离，并且可以独立scale不影响性能和可用性
5. 高可用
Snowflake能够容忍节点，集群，甚至全部的数据中心失败。在软件和硬件更新的时候不会下线。
6. 持久性
具有极高的持续性，可防止意外数据丢失：克隆、下线和跨区域备份
7. 节省
即可压缩。具有很高的计算效率，所有的表数据都被压缩。用户只为他们实际使用的存储和计算资源付费。
8. 安全
所有数据包括临时文件和网络流量都是端到端加密的。

# 2. 存储VS计算

Share-nothing架构的一些缺点：

1. 异构负载：share-nothing架构下所有node都是同构的，无法平衡（重IO，轻计算）和（轻IO，重计算）两种负载

2. membership changes：增删节点或者故障时数据可能需要重新shuffle，会带来性能的波动以及数据不可用

3. Online Update：在线升级需要集群内所有节点参与，由于计算存储耦合增加了系统升级的风险。

以上问题在云上环境显得尤为尖锐，所以snowflake采用计算存储分离的架构，两者松耦合，可独立扩展，计算层采用shared-nothing engine,存储层直接使用的是亚马逊的S3，并且在计算节点本地会cache一部分热点数据。这部分共享数据一般在本地盘使用SSD存储，Snowflake把这种架构称之为multi-cluster,shared-data架构。

# 3. Architecture
Snowflake是典型计算存储分离架构(multi-cluster, shared-data)。Snowflake主要包括三个模块，Data Storage(存储层)，Virtual Warehouse(“muscle”， virtual machine构成的执行层)， Cloud Services(“brain”， 主要包括一些元信息，事务管理，优化器，访问控制等功能)。三个模块之间通过RESTful API通信连接。
![image.png](/wiki/v1/wiki/uploads/200005545/202212/1670662828801.png 'image.png')
Data Storage
在权衡了高可用以及持久性等等因素，Snowflake选择Amazon S3的存储服务，更多精力集中在热点数据的local caching和skew 优化。相比于本地盘，S3延迟和I/O都要更高，但提供了HTTP的PUT/GET/DELETE接口，S3的文件是不可变的，可能整体覆盖，不能追加，PUT请求需要加上文件的实际大小但支持GET部分文件。所以Snowflake的文件是按列存储的，每个文件的header中记录的是每一列的offset，通过Get可以选取某几列数据。

S3中也存储着中间结果，比如执行过程中的临时下盘文件。同时也存储着query results，以便于客户端直接从S3上取得查询结果而不需要server的参与。

Virtual Warehouse
2.1 Elasticity and Isolation

Virtual Warehous使用的是亚马逊的EC2实例构成的集群。集群的最小实例称为VW，，组成VM的单个EC2实例称为worker node，用户不需要和worker node打交道，只需要选择不同性能的VW即可，并且根据VW来提供服务和计费。

VWs 按需使用计费，如果用户没有query，关掉VW也是可以的。目前每个查询都只在一个 VW上执行。VW中的worker node是隔离的，所以查询之间也是隔离的；对于大查询由于故障失败后需要所有worker node重试；以上两个问题也是接下来可以优化的工作。

对于弹性值得一提的是作者指出在云上相同的价格可以实现更好的性能，比如data load在4个node上执行需要15个小时，而32个node上可能只需要2小时，二者价格差不多，但用户体验后者要更好，这也是弹性带来的好处之一。

2.2 Local Caching and File Stealing

前面说到每个worker node本地都会cache一部分热点数据，主要包括一些file headers和文件中访问过的某些列数据。cache在当前VW的worker nodes之间共享，实现上应该就是一个LRU cache，每个cache node应该是一些file stream和column request。

为了避免单个文件的冗余缓存，所以优化器对表文件名执行一致性哈希后存储在某一个worker node上。所以并发的query只能在相同的worker node上访问缓存的table。

如果单个worker node故障或者VW调整后，不会立刻执行shuffle，而是随着query继续执行，将LRU上的cache逐渐替换，优点时不会因为节点故障而降级。

除了local caching的优化外，有需要处理倾斜的问题，当不同worker node执行速度不同是，慢的节点会把他的他的任务释放出来，以scan为例，当前节点发展有很多文件还没有扫描完成时将会释放S3中的文件所有权，其他worker node就可以从S3下载文件，从而使负载平衡，这个方式snowflake称之为file stealing。

2.3 Execution Engine

VW执行引擎采用的技术包括列存，向量化，Push-based execution（DAG plan）。这里不过多分析了。

Cloud Services
云服务层是多租户且长期存在的，云服务层中的访问控制，优化器，事务管理器等在用户之间共享，并且所有的服务都通过复制的方式实现高可用。

3.1 Query Management and Optimization

CloudServices层包括了parser解析，访问控制，优化等功能。优化器是典型的Cascades的基于代价的优化器。统计信息在data load和update时会更新。snowflake不使用索引，优化器搜索空间会相应减少，并且将一些决策在执行时确定比如join的 数据分布方式。

执行时将生成的计划下发到所有的worker node，Cloud services会收集执行时的统计信息以及监控节点故障，并且可以查询历史执行信息。

3.2 Concurrency Control

并发访问控制是是典型的多版本（MVCC）。通过SI（Snapshot Isolation）实现ACID。

3.3 Pruning

对于snowflake访问数据有几个问题：1, 由于S3以及compressed file的特点会有很多的随机访问。2, 索引会增加数据的容量以及加载时间。3, 用户需要显式设置索引并且需要调优等等问题。解决上述问题的一个方法是min-max based pruning ，主要是通过在每个文件中维护一个chunk数据分布信息，通过filter快速过滤不需要的数据。这个的剪枝方法需要很少的数据并且在顺序访问，优化等方面既简单又高效。

除了一些静态剪枝方法，也有一些执行时的动态剪枝方法，像hashjoin时可以收集build端的统计信息，然后用来过滤probe端的文件，这种技术也叫做bloom joins。

2. FEATURE HIGHLIGHTS

2.1 Pure Software-as-a-Service Experience

​ Snowflake支持标准的数据库接口，比如（JDBC，ODBC），并且可以通过web UI直接可视化数据，管理数据，监控等等。这种可视化其实对于用户快速了解产品，管理监控数据。

2.2 Fault Resilience

由于存储层使用的是Amazon的S3，S3副本存放在多个数据中心也叫做AZ（availability zones），所以整个S3可以达到99.99% availability和99.999999999% durability，同时适配S3的架构，元数据也跨AZ存储，当节点故障时，可以根据元信息切换到可用的AZ，整个Cloud Services的剩余服务也分布在多个AZ的无状态节点上，由Load Balancer来分配user request。当单个节点故障或者整个AZ故障时，用户的查询将重定向到新的可用节点上。但VW不是跨AZ分布的，为了减少AZ之间网络延时对性能的影响。如果AZ内某个节点故障导致查询失败，会重新执行，排除或替换故障节点，并且可以通过备用node pool来重新配置VW。

但如果整个AZ故障，该AZ中执行的所有查询都将失败，并且需要重新在新的AZ中配置VW，但这种情况很少发生。

2.3 Online Update

软件升级是常见的操作，系统内可能存在不同版本的软件，无论是对于Cloud Services 还是Virtual Warehouse。通过一个mapping layer来管理所有的metadata versioning。在软件升级时，将新版本与旧版本一起部署，用户信息逐渐迁移到新版本中，并且对应用户的查询也发送到新版本的VW中，旧版本的查询在执行完旧的查询后逐渐停止服务。

如下图所示。在升级时CS和VW都同时部署两个版本，并且是无状态的，由Load Balancer来发送请求，并且CS的version1只鱼VW的Version1连接，但VW的不同版本之间共享Cache，避免升级后重新填入。

Snowflake基本保持一周升级一次服务的进度，并且在预生产环境下测试升级和降级程序，以便发现bug时可以平滑降级或者升级。

2.4 Semi-Structured and Schema-Less Data

Snowflake主要支持三种类型的半结构化数据：VARIANT, ARRAY, and OBJECT. 其中VARIANT包括任意的SQL类型（DATE，VARCHAR）,可变长度的ARRAY，JavaScript式的OBJECT，字符串到VARIANT的映射等。所以ARRAY和OBJECT也是VARIANT的子集。VARIANT三者的内部表示基本相同，都是self-describing, compact binary serialization 可以用于join key，group key和order key。

对于VARIANT类型，Snowflake会自动收集统计信息，并执行类型推断并提取出单独的列压缩保存，并提前计算一些聚合数据用来剪枝。在scan时可以下推投影或者cast等运算符。并且通过Bloom filte来快速过滤不需要的文件优化执行路径上的剪枝优化。

对于一些日期类型的数据，由于没有定义保存精度等问题，Snowflake会使用单独的列转换的结果以及原始的字符串，来应对类型转换时精度损失的问题以及转换后可以用来做一些剪枝的操作。

对于VARIANT类型以及关系型的TPCH测试结果中，半结构化数据的查询性能几乎与列存相差无几。


# 参考
https://zhuanlan.zhihu.com/p/390025973
https://zhuanlan.zhihu.com/p/366369705
