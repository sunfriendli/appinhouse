## appinhouse

公司内网的appinhouse Web服务，包括前端 `web` 和后端 RESTful 风格的服务`server`。

## 快速启动

脚本见`script`目录。

如果希望`单独部署`其中一个，或几个，请参考[文档](#文档)。

##### 环境

`ubuntu，go1.6，git，redis`

##### 代理

`由于网络的问题，安装某些依赖包会出现hash sum mismatch问题，这时候需要代理 `，见[***shadowsocks文档***](doc/shadowsocks.md)

#### 第一步，部署redis(可选)

如果已有`redis`,可跳过。


```bash
sudo ./install_redis.sh password #密码可选
```

#### 第二步，部署inhouse server

```bash
#第一个参数是密码，脚本会创建一个名字为appinhouse的用户，需要输入密码用来登录
#第二个参数是redis的地址
#第三个参数是redis的密码（可选）
sudo ./install_appinhouse.sh password redis_address redis_password 
```
#### 第三步，部署nginx

需要`https`证书，见[***nginx配置文档***](doc/nginx.md#证书)。


```bash
sudo ./install_nginx.sh #依赖第二步
```

#### 第四步，部署ftp（可选）

如果你打包的客户端，已经放在一个`web`服务下，不需要装ftp。

```bash
#ftp_root_dir是ftp的根目录，
#passwd是所有用户的密码
#ftp用户名从inhouse_service/conf/app.conf文件的app_names属性中读取
sudo ./install_ftp.sh ftp_root_dir password #依赖第二步
```

#### 第五步，归档

见[***客户端（ipa/apk)的归档***](doc/archive.md) 。


## 功能特点

* RESTful 风格
* 客户端（ipa/apk）的归档
* 自动分平台提供（iphone/android/pc）最新版本和历史版本浏览
* 查看任一版本的客户端的详细信息，其中包括回溯到Jenkins上的此次构建
* 提供plist生成

## 文档

* [api文档](doc/api.md)
* [客户端（ipa/apk)的归档](doc/archive.md)
* [应用配置文档](doc/conf.md)
* [appinhouse配置文档](doc/appinhouse.md)  
* [ftp配置文档](doc/ftp.md)  
* [nginx配置文档](doc/nginx.md) 
* [shadowsocks文档](doc/shadowsocks.md)   



