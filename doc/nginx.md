## nginx安装与配置

#### 目录

`nginx`配置文件目录`web/nginx`下。

#### 环境

`ubuntu`

#### 安装

```bash
 sudo ./install_nginx.sh  #文件位置script/install_nginx.sh
```
#### 证书

因为ios访问	`plist`需要`https`,所以需要一个证书。推荐[Let’s Encrypt](https://letsencrypt.org/)。理由：免费。

申请证书，见[参考资料](#参考资料)。

#### 配置

文件位置`conf/nginx/appinhouse_ssl.conf`

在`/etc/nginx/conf.d`添加`appinhouse_ssl.conf`

## 参考资料

[证书申请](https://blog.zhiguang.me/lets-encrypt/)


