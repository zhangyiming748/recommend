根据id查找内容， 首先查找cachelist， 其次查找ES By Id（由于cache是最近5000条， 所以暂且不实现 从es中查找）
其中又分为不关注来源的查找 和关注来源 的查找

query分为根据id查找 和 获取最新多少条内容的query

qNewArticleFromES

既然有缓存，idsquery（更名为 qNewArticlesFromES），filterids 这个就不需要了。
