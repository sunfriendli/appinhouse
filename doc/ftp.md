## vsftpd安装与配置

#### 目录

`vsftpd`配置文件目录`ftp`下。

#### 环境

`ubuntu`

#### 安装

```bash
 sudo apt-get install vsftpd
```
#### 查看是否运行

```bash
sudo netstat -npltu | grep 21
tcp        0      0 0.0.0.0:21              0.0.0.0:*               LISTEN      15601/vsftpd    
```
#### 配置

见`vsftpd.conf`文件

#### ftp_root_dir

ftp的主目录，所有`接入方的项目`的目录，都在这个目录下

#### 创建用户

##### 创建用户的目的

不同的接入方，拥有不同的目录，对目录拥有不同的权限。

在`ftp_root_dir`下创建三个目录`user1,user2,user3`.使得
- 用户`user1 `的主目录为 `user1`，在该目录下拥有所有权限
- 用户`user2` 的主目录为 `user2`，在该目录下拥有所有权限
- 用户`user3` 的主目录为` user3`，在该目录下拥有所有权限

##### 为Vsftpd创建用户

```bash
sudo useradd user1 -d ftp_root_dir/test1 -s /bin/false #不能登录系统
sudo chown user1:user1  ftp_root_dir/test1 #必须执行
```

##### 配置PAM文件
见`vsftpd`文件
```bash
sudo vi /etc/pam.d/vsftpd
# 注掉下面这一行
auth	required	pam_shells.so

```
##### 配置/vsftpd_user_conf/

`为每个用户单独配置权限`，见目录`vsftpd_user_conf`下的文件，***`vsftpd_user_conf`的文件名要和用户名一致***

```bash
write_enable=YES
anon_world_readable_only=NO
anon_upload_enable=YES
anon_mkdir_write_enable=YES
anon_other_write_enable=YES
local_root=ftp_root_dir/user1 #配置user1要对应的目录
```

##### 配置vsftpd.conf文件

默认的配置里会有一些配置，如果没有的话，见[参考资料](#参考资料)。
在`vsftpd.conf`加入下列配置项

```bash
user_config_dir=/etc/vsftpd/vsftpd_user_conf #指定每个用户账号配置目录
#其余的参考vsftpd.conf文件
sudo service vsftpd reload
```

## 注意

- `配置文件中不能有空格`

## 参考资料

[ubuntu下vsftpd教程](http://wiki.ubuntu.org.cn/Vsftpd)

