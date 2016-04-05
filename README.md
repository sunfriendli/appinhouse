## appinhouse

公司内网的appinhouse Web服务，包括前端 `web` 和后端 RESTful 风格的服务`inhouse_service`。

## 快速启动
#####下载之前
在机器上找个目录，执行以下命令
```bash
mkdir pkg
mkdir bin
mkdir src
```
#####下载
```bash
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
##功能特点
* RESTful 风格
* 客户端（ipa/apk）的归档
* 自动分平台提供（iphone/android/pc）最新版本和历史版本浏览
* 查看任一版本的客户端的详细信息，其中包括回溯到Jenkins上的此次构建
* 提供plist生成
##文档
* [api文档](inhouse_service/controllers/README.md)
* [客户端（ipa/apk)的归档]
* [nginx配置]  



