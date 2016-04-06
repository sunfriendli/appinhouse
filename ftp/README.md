## vsftpd安装与配置

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

#####为虚拟用户创建本地系统用户

```bash
sudo useradd vsftpd -d /home/vsftpd -s /bin/false #不能登录系统
```

#####创建虚拟用户数据库

```bash
vi login_user.txt 
#注意：奇数行为账户名，偶数行为密码。也就是1.3.5.等行为用户名，2.4.6行为密码；
#最后一行需要回车（否则建立数据库文件时无法识别最后一行，导致报奇数行错误）。
sudo dbx.x_load -T -t hash -f login_user.txt /etc/vsftpd_login.db
sudo chmod 600 /etc/vsftpd_login.db
```

#####配置PAM文件
见`vsftpd`文件
```bash
# 加入下面两行，其余注掉
auth sufficient pam_userdb.so db=/etc/vsftpd_login
account sufficient pam_userdb.so db=/etc/vsftpd_login

```
#####配置/etc/vsftpd_user_conf
见目录`vsftpd_user_conf`下的文件，***`vsftpd_user_conf`的文件名要和虚拟用户数据库中的账户名一致***
```bash
write_enable=YES
anon_world_readable_only=NO
anon_upload_enable=YES
anon_mkdir_write_enable=YES
anon_other_write_enable=YES
local_root=/home/inhouse/appinhouse/rog2
local_umask=000 
```
因为`inhouse`需要删除文件，所以`local_umask=000`  见[***注意3***](#注意)

#####配置vsftpd.conf文件

```bash
guest_enable=YES 
guest_username=vsftpd
virtual_use_local_privs=YES
user_config_dir=/etc/vsftpd_user_conf
pam_service_name=vsftpd
```

##注意

1. 如果只是`ftp` 那么local_root的属组和属主应该是[为虚拟用户创建本地系统用户](#为虚拟用户创建本地系统用户)
2. 如果是和`web`应用共用，不涉及到删除 那么`local_root`属组和属主应该是[为虚拟用户创建本地系统用户](#为虚拟用户创建本地系统用户) ，并且权限需要`755`
3. 如果是和`web`应用共用，涉及到删除 那么`local_root`属组和属主是谁都行，但权限需要`777`

##参考资料

[ubuntu下vsftpd教程](http://wiki.ubuntu.org.cn/Vsftpd)