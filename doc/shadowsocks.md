## shadowsocks安装文档

`ubuntu`在`apt-get`的时候，有的包受网络影响,会报错`hash sum mismatch`。这是需要代理。

##### 环境

`ubuntu`

##### 安装

```bash
sudo apt-get -y install python-pip
sudo pip install shadowsocks
```
##### 配置

新建名字为`ss_conf.json`的配置文件，内容如下：

```bash
{
    "server":"ip",
    "port": port,
    "password": "password",
    "method": "method", #可选
    "remarks": "remarks",#可选
    "auth": false
}

```

##### 运行

```bash
nohup sslocal -c ~/ss_conf.json > ~/ss.log 2>&1 &
```

