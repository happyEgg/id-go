一个使用beego框架写的id生成器                                                                       
针对不同的场景去考虑能不能降低某方面的要求，以换取其它方面的提升。仔细考虑我们的需求，我们只要求递增，
并没有要求连续，也就是说出现一大段跳跃是允许的（例如分配出的sequence序列：1,2,3,10,100,101）。
于是我们实现了一个简单优雅的策略:
1、内存中储存从数据库中读取的sequence:tempNum，以及分配上限：MAXSEQ
2、分配sequence时，将tempNum++，同时与分配上限MAXSEQ比较：如果tempNum >= MAXSEQ,将分配上限提升一个步长
MAXSEQ +=Step，并持久化MAXSEQ
3、重启时，读出持久化的max_seq，赋值给tempNum


这样通过增加一个预分配sequence的中间层，在保证sequence不回退的前提下，大幅地提升了分配sequence的性能。实际应用中每次提升的步长为10000，
那么持久化的硬盘IO次数从之前~10^7 QPS降低到~10^3 QPS，处于可接受范围。
在正常运作时分配出去的sequence是顺序递增的，只有在机器重启后，第一次分配的sequence会产生一个比较大的跳跃，跳跃大小取决于步长大小。

二、在生成的数据表中添加初始化数据
table_name			    max_seq
   user              10000
                  (初始值)
