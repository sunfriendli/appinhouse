# 数据库改造方案

## 备份

```bash
# 第一步 备份
redis 127.0.0.1:6379> SAVE 
# 第二步 查看数据存储目录
redis 127.0.0.1:6379> CONFIG GET dir
1) "dir"
2) "/usr/local/redis/bin"
# 第三步 将dump.rdb拷贝到备份目录

```

## 恢复

```bash
# 第一步 停止redis服务
sudo service redis-server stop
# 第二步 将dump.rdb拷贝到/usr/local/redis/bin下

# 第三步 启动redis服务
sudo service redis-server start
```

## 数据修改方案

```java
//获得所有app的key--> appkeys
List appkeys = getAllAppKeys();
//遍历appkeys
for(String appkey : appkeys){

  int iosReleaseCount = getIosReleaseCountByAppKey(appkey);
  int iosDevCount = getIosDevCountByAppKey(appkey); 
  int androidReleaseCount = getAndroidReleaseCountByAppKey(appkey); 
  int androidDevCount = getAndroidDevCountByAppKey(appkey); 
  int releaseCount = iosReleaseCount + androidReleaseCount;
  int devCount = androidDevCount + iosDevCount;
  // 如果既有dev的，也有release的
  if(releaseCount > 0 && devCount > 0 ){
    // 生成一个release的app
    String newReleaseKey = appkey + "-release"
    addApp(newReleaseKey);
    //把上述的release环境的迁移到新app，并去掉key中env的部分
    modifyoldReleaseToNewRelease(newReleaseKey);

    // 生成一个dev的app
    String newDevKey = appkey + "-dev"
    addApp(newDevKey);
    //把上述的release环境的迁移到新app
    modifyoldDevToNewRelease(newDevKey);
    //删掉appkey的应用
    deleteApp(appkey);
  }else{
    //获得appkey下的列表key，并去掉key中env的部分
    List listkeys = getListKey(appkey);
    //去掉key中env的部分，若是ios，去掉plist中的env部分
    removeEnv(listkeys);
  }
}
```