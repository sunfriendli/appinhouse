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

文件位置`web/nginx/appinhouse_ssl.conf`

在`/etc/nginx/conf.d`添加`appinhouse_ssl.conf`

```bash
 location /download {
        alias  /var/www/appinhouse.rog2.org;
    }
    location /plist {
        alias  /home/appinhouse/appinhouse_data;
    }     
```

**`对应如下`**

```bash
    location /archive_file_root_dir_alias { 
        alias  ftp_root_dir; #对应的是ftp根目录
    }
    location /plist_root_dir_alias {
        alias  root_dir; #对应是plist文件的根目录
    }     
```
- `archive_file_root_dir_alias`的值，参见[***应用配置文档***](#conf.md#archive_file_root_dir_alias)。
- `plist_root_dir_alias`的值，参见[***应用配置文档***](#conf.md#plist_root_dir_alias)。
- `ftp_root_dir`的值，参见[***ftp配置文档***](#ftp.md#ftp_root_dir)。
- `root_dir`的值，参见[***应用配置文档***](#conf.md#root_dir)。

##参考资料

[证书申请](https://blog.zhiguang.me/lets-encrypt/)


