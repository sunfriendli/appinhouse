## appinhouse

公司内网的appinhouse Web服务，包括前端 `web` 和后端 RESTful 风格的服务`inhouse_service`。

## 快速启动

脚本见`bin`目录。

#### 第一步，部署inhouse server

##### 环境

`ubuntu，go1.6，git`

`由于网络的问题，安装某些依赖包会出现hash sum mismatch问题，这时候需要代理 `，见[***shadowsocks文档***](doc/shadowsocks.md)

```bash
sudo ./install_appinhouse.sh password #创建一个名字为appinhouse的用户，需要输入密码用来登录
```
#### 第二步，部署nginx

需要`https`证书，见[***nginx配置文档***](doc/nginx.md#证书)。


```bash
sudo ./install_nginx.sh #依赖第一步
```

#### 第三步，部署ftp


```bash
#ftp_root_dir是ftp的根目录，
#passwd_postfix是密码后缀，密码：ftp用户名+passwd_postfix
#ftp用户名从inhouse_service/conf/app.conf文件的app_names属性中读取

sudo ./install_ftp.sh ftp_root_dir passwd_postfix #依赖第一步

```

#### 第四步，归档


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



