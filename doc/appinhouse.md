## appinhouse服务安装文档

appinhouse以service的方式提供服务。

### 环境

`ubuntu，go1.6，git`

#### 创建专属用户
```bash
 sudo useradd appinhouse -s /bin/bash
 sudo passwd appinhouse
 ```
#### 安装go
```bash
sudo add-apt-repository ppa:ubuntu-lxc/lxd-stable
sudo apt-get update
sudo apt-get install golang
```
#### 安装git

```bash
sudo apt-get install git
```

#### 安装bee

```bash
export GOPATH=~/bee #设置你指定的目录
cd $GOPATH
go get github.com/beego/bee
cd bin
sudo  cp bee /usr/bin/
```

#### 安装服务

`用appinhouse用户登录`

##### 下载工程

在机器上找个目录，执行以下命令

```bash
mkdir pkg
mkdir bin
mkdir src
cd src
git clone https://github.com/rog2/appinhouse.git
```
##### 打包

```bash
cd appinhouse/inhouse_service/bin
chmod +x *.sh
./pack.sh
```
##### 设置redis数据库配置文件

##### 目录及文件名

取值于`app.conf`下`[redis]`的`conf_dir`，见[***应用配置文档***](conf.md#conf_dir)。

##### 文档内容

###### 地址

`key`取值于`app.conf`下`[redis]`的`env_addr_name`，见[***应用配置文档***](conf.md#env_addr_name)。

`value` redis的地址

###### 密码（可选）

`key`取值于`app.conf`下`[redis]`的`env_password_name`，见[***应用配置文档***](conf.md#env_password_name)。

`value` redis的密码

##### 运行

`用sudo用户登录`

```bash
sudo cp appinhouse.sh /etc/init.d/appinhouse
sudo chmod +x /etc/init.d/appinhouse
echo "ARTIFACT=/home/appinhouse/appinhouse_server/inhouse_service"       | sudo tee -a /etc/default/appinhouse
echo "APPINHOUSE_USER=appinhouse"      | sudo tee -a /etc/default/appinhouse
sudo update-rc.d appinhouse defaults
sudo service appinhouse start
```

