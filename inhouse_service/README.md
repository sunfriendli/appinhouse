## 客户端（ipa/apk)的归档

####ftp账号

`inhouse` 会分给接入方一个`ftp`账号和密码

####ftp路径介绍

```
根目录 
    |-dev --- 开发环境
	    |-android --- android平台
		    |-data --- 放置版本文件
			    |-1.0.1 --- 目录名为版本号，接入方创建
			        |-xxx.apk --- 接入方上传
			    |-1.0.2 --- 目录名为版本号，接入方创建
		    |-list --- 放置打包描述文件,按生成时间排序
		        |-xxx.txt --- 描述文件，接入方上传
		        |-xxxx.txt --- 描述文件，接入方上传
	    |-ios --- ios平台
	        |-data --- 放置版本文件,plist,full-size-image,display-image
	            |-1.0.1  --- 目录名为版本号，接入方创建
	                |-xxx.ipa --- 接入方上传
	                |-xxx.plist ---调API生成
	                |-display.png --- 接入方上传，名字固定
	                |-full_size.png --- 接入方上传，名字固定
	            |-1.0.2 --- 目录名为版本号，接入方创建
	        |-list ---放置打包描述文件
	            |-xxx.txt --- 描述文件，接入方上传
		        |-xxxx.txt --- 描述文件，接入方上传
    |-release --- 发布环境
        |-android --- android平台
            |-data --- 同上
            |-list --- 同上
        |-ios --- ios平台
            |-data --- 同上
            |-list --- 同上   
```

####版本文件

首先找到对应`环境，平台`下的`data`，`list`文件夹。

#####android

- 在`data`文件夹下创建以本次版本号命名的目录。
- 将`.apk`文件放到`版本号`目录下,`.apk`文件名字自定义。
- 将描述文件放到`list`目录下，描述文件名字自定义。描述文件类型为`txt`，`描述文件示例：`

	```bash
	version=1.0.0    #版本号
	time=2016-8-9 00:00:00  #打包时间
	description=小测试描述android #打包描述
	url=http://example.me #构建的url，例：jenkins上本次构建的url
	channel=mi #渠道，用户自定义，例：小米渠道，mi，没有渠道`无`
	```

#####ios

- 在`data`文件夹下创建以本次版本号命名的目录。
- 将`.ipa`文件放到`版本号`目录下,`.ipa`文件名字自定义。
- 将`full-size-image`文件放到`版本号`目录下,`.ipa`文件名字固定，`full_size.png`。
- 将`display-image`文件放到`版本号`目录下,`.icon`文件名字固定，`display.png`。
- 将描述文件放到`list`目录下，描述文件名字自定义。描述文件类型为`txt`，`channel` 苹果官方的值 ，见***[应用配置文档](controllers/README.md#生成plist)***，`描述文件示例：`

	```bash
	version=1.0.0    #版本号
	time=2016-8-9 00:00:00  #打包时间
	description=小测试描述ios #打包描述
	url=http://example.me #构建的url，例：jenkins上本次构建的url
	channel=appstore #渠道，苹果官方为appstore ，其余越狱渠道用户自定义。
	```
	
- 调用API生成plist，见***[API文档](controllers/README.md#生成plist)***
