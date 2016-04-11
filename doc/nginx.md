## nginx安装与配置

####目录

`nginx`配置文件目录`web/nginx`下。

####环境

`ubuntu`

####安装

```bash
 sudo ./install_nginx.sh  #文件位置web/nginx/install_nginx.sh
```
####证书

因为ios访问	`plist`需要`https`,所以需要一个证书。推荐[Let’s Encrypt](https://letsencrypt.org/)。理由：免费。

申请证书，见[参考资料](#参考资料)。

####配置

```bash
cd /etc/nginx/conf.d
#在这个文件夹，创建你的conf
sudo vi appinhouse_ssl.conf #文件位置web/nginx/appinhouse_ssl.conf
sudo service nginx  reload
```



##参考资料

[证书申请](https://blog.zhiguang.me/lets-encrypt/)


