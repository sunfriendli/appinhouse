## appinhouse

公司内网的appinhouse Web服务，包括前端 `web` 和后端 RESTful 风格的服务`inhouse_service`。

## 快速启动

####第一步，部署inhouse server

#####环境

`ubuntu，go1.6，git`

```bash
#安装go
$ sudo add-apt-repository ppa:ubuntu-lxc/lxd-stable
$ sudo apt-get update
$ sudo apt-get install golang
#安装git
sudo apt-get install git
```

#####下载

在机器上找个目录，执行以下命令

```bash
mkdir pkg
mkdir bin
mkdir src
cd src
git clone https://github.com/rog2/appinhouse.git
```
#####打包
```bash
cd appinhouse/inhouse_service/bin
chmod +x *
./pack.sh
```
#####运行

```bash
cd /home/当前用户/appinhouse_server/bin
./inhouse.sh start
```

####第二步，部署nginx

#####环境

```bash
./install_nginx.sh  #文件位置web/nginx/install_nginx.sh
```

#####配置

```bash
cd /etc/nginx/conf.d
#在这个文件夹，创建你的conf
sudo vi appinhouse_ssl.conf #文件位置web/nginx/appinhouse_ssl.conf
sudo service nginx  reload
```

####第三步，部署ftp


见[***ftp配置文档***](doc/ftp.md)  。


##功能特点

* RESTful 风格
* 客户端（ipa/apk）的归档
* 自动分平台提供（iphone/android/pc）最新版本和历史版本浏览
* 查看任一版本的客户端的详细信息，其中包括回溯到Jenkins上的此次构建
* 提供plist生成

##文档

* [api文档](doc/api.md)
* [客户端（ipa/apk)的归档](doc/archive.md)
* [应用配置文档](doc/conf.md)
* [ftp配置文档](doc/ftp.md)  



