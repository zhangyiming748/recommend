启动ES报错
```
Unrecognized VM option 'UseParNewGC'
Error: Could not create the Java Virtual Machine.
Error: A fatal exception has occurred. Program will exit.
```
初步判断原因
JDK9的GC`UseParNewGC`被废弃 
尝试解决
尝试更换UseConcMarkSweepGC
尝试更换为oracle Java
报错
```
Exception in thread "main" java.lang.NoClassDefFoundError: org/elasticsearch/common/jackson/dataformat/yaml/snakeyaml/error/YAMLException
	at org.elasticsearch.common.jackson.dataformat.yaml.YAMLFactory._createParser(YAMLFactory.java:426)
	at org.elasticsearch.common.jackson.dataformat.yaml.YAMLFactory.createParser(YAMLFactory.java:327)
	at org.elasticsearch.common.xcontent.yaml.YamlXContent.createParser(YamlXContent.java:90)
	at org.elasticsearch.common.settings.loader.XContentSettingsLoader.load(XContentSettingsLoader.java:45)
	at org.elasticsearch.common.settings.loader.YamlSettingsLoader.load(YamlSettingsLoader.java:46)
	at org.elasticsearch.common.settings.ImmutableSettings$Builder.loadFromStream(ImmutableSettings.java:982)
	at org.elasticsearch.common.settings.ImmutableSettings$Builder.loadFromUrl(ImmutableSettings.java:969)
	at org.elasticsearch.node.internal.InternalSettingsPreparer.prepareSettings(InternalSettingsPreparer.java:110)
	at org.elasticsearch.bootstrap.Bootstrap.initialSettings(Bootstrap.java:144)
	at org.elasticsearch.bootstrap.Bootstrap.main(Bootstrap.java:215)
	at org.elasticsearch.bootstrap.Elasticsearch.main(Elasticsearch.java:32)
Caused by: java.lang.ClassNotFoundException: org.elasticsearch.common.jackson.dataformat.yaml.snakeyaml.error.YAMLException
	at java.base/jdk.internal.loader.BuiltinClassLoader.loadClass(BuiltinClassLoader.java:581)
	at java.base/jdk.internal.loader.ClassLoaders$AppClassLoader.loadClass(ClassLoaders.java:178)
	at java.base/java.lang.ClassLoader.loadClass(ClassLoader.java:522)
	... 11 more

```
stack overflow、ES社区询问
有很多相同提问
但没有人回答