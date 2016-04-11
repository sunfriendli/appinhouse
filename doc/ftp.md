## vsftpd安装与配置

####目录
`vsftpd`配置文件目录`ftp`下。
####环境

`ubuntu`

####安装

```bash
 sudo apt-get install vsftpd
 sudo apt-get install db-util
```
####查看是否运行

```bash
sudo netstat -npltu | grep 21
tcp        0      0 0.0.0.0:21              0.0.0.0:*               LISTEN      15601/vsftpd    
```
####配置

见`vsftpd.conf`文件

####虚拟用户配置

#####创建虚拟用户目的

不同的接入方，拥有不同的目录，对目录拥有不同的权限。

在`root_dir`（见***[应用配置文档](conf.md)***）下创建三个目录`test1,test2,test3`.使得
- 用户`user1 `的主目录为 `test1`，在该目录下拥有所有权限
- 用户`user2` 的主目录为 `test2`，在该目录下拥有所有权限
- 用户`user3` 的主目录为` test3`，在该目录下拥有所有权限

#####为虚拟用户创建本地系统用户

```bash
sudo useradd vftpuser -d /home/vftpuser -s /bin/false #不能登录系统，用户名自定义
sudo chown vftpuser:vftpuser  /home/vftpuser #必须执行
```

#####创建虚拟用户数据库

```bash
vi login_user.txt 
#在文件中输入
user1
xxxx
user2
yyyy
user3
zzzz
#注意：奇数行为账户名，偶数行为密码。也就是1.3.5.等行为用户名，2.4.6行为密码；
#db-util老版本 可能最后一行需要回车（否则建立数据库文件时无法识别最后一行，导致报奇数行错误）。
sudo dbx.x_load -T -t hash -f login_user.txt /etc/vsftpd/vsftpd_login.db #dbx.x_load x为你安装的版本
sudo chmod 600 /etc/vsftpd/vsftpd_login.db
```

#####配置PAM文件
见`vsftpd`文件
```bash
cd /etc/pam.d/
# 加入下面两行，其余注掉
auth sufficient pam_userdb.so db=/etc/vsftpd/vsftpd_login
account sufficient pam_userdb.so db=/etc/vsftpd/vsftpd_login

```
#####配置/vsftpd_user_conf/

`为每个虚拟用户单独配置权限`，见目录`vsftpd_user_conf`下的文件，***`vsftpd_user_conf`的文件名要和虚拟用户数据库中的账户名一致***

```bash
write_enable=YES
anon_world_readable_only=NO
anon_upload_enable=YES
anon_mkdir_write_enable=YES
anon_other_write_enable=YES
local_root=/home/inhouse/appinhouse/test1 #配置user1要对应的目录
local_umask=000 
```
因为`inhouse`需要删除文件，所以`local_umask=000`  见[***注意3***](#注意)

#####配置vsftpd.conf文件

默认的配置里会有一些配置，如果没有的话，见[参考资料](#参考资料)。
在`vsftpd.conf`加入下列配置项

```bash
guest_enable=YES 
guest_username=vftpuser #指定上文为虚拟用户创建的本地用户
virtual_use_local_privs=YES #虚拟用户和本地用户拥有一样的权限
user_config_dir=/etc/vsftpd/vsftpd_user_conf #指定每个虚拟用户账号配置目录
dual_log_enable=YES
xferlog_file=/var/log/vsftpd.log
```

##注意

1. `配置文件中不能有空格`
1. 如果只是`ftp` 那么local_root的属组和属主应该是[为虚拟用户创建本地系统用户](#为虚拟用户创建本地系统用户)
2. 如果是和`web`应用共用，不涉及到删除 那么`local_root`属组和属主应该是[为虚拟用户创建本地系统用户](#为虚拟用户创建本地系统用户) ，并且权限需要`755`
3. 如果是和`web`应用共用，涉及到删除 那么`local_root`属组和属主是谁都行，但权限需要`777`

##参考资料

[ubuntu下vsftpd教程](http://wiki.ubuntu.org.cn/Vsftpd)

