```
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.9.0
```

```
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.9.0
```




[图解elasticsearch原理](#https://juejin.im/entry/6844903715468492813)
[Mastering Elasticsearch(中文版)](#https://doc.yonyoucloud.com/doc/mastering-elasticsearch/chapter-1/index.html)