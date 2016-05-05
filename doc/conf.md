## 应用配置文档

#### 应用配置目录

`server/conf/app.conf`

#### 应用配置

##### appname

应用名称，默认是` beego`。通过`bee new`创建的是创建的项目名。

##### runmode

应用的运行模式，可选值为 prod, dev 或者 test. 默认是 dev, 为开发模式，在开发模式下出错会提示友好的出错页面。

##### graceful

优雅模式，默认是false，启动优雅模式，可以监听系统信号来进行热启和优雅的关闭服务(关闭时不接收新的请求，当前连接超时等等)。

#### 用户配置

`[users]`  

##### log_dir

日志输出文件的根目录。

##### ios_channel

苹果官方的渠道值。

##### domain

网站的域名。

##### min_residue 

版本最小保留数。

##### max_page 

最大页数值。

##### page_size 

每页显示的记录数。

#### redis配置

`[redis]`  

#### addr

redis的地址

#### password

redis密码

#### pool_siz

redis的pool个数

#### db

redis的db