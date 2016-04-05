## 企业内部appinhouse服务API接口文档

##接口通用信息
1. 接口基于HTTP RESTFul思想架构设计。
2. 接口地址（URL）统一采用：http://域名/api/请求资源?过滤参数的格式。
	例如：`http://域名/api/rog2/list?page=1`。
3. 响应内容默认JSON格式返回。

##接口索引
[移动端获得最新打包版本](#移动端获得最新打包版本)
[移动端获取历史列表](#移动端获取历史列表)
[获得最新打包版本](#获得最新打包版本)
[获取历史列表](#获取历史列表)
[生成plist](#生成plist)
[整理历史版本](#整理历史版本)


##接口详情
###接口公共信息
###app
应用程序的名字，由inhouse生成，配置后，告知接入方。
###platform
移动端手机设备的操作系统名字，`枚举值`。

| value     |     ios |   android |
| :-------- | --------:| :------: |
| 含义    |   苹果设备|  安卓设备  |

###environment
打包的应用程序所对应的运行环境，`枚举值`

| value     |     dev|   release|
| :-------- | --------:| :------:|
| 含义       |   开发环境|  发布环境|
###pageSize
一页显示多少记录，`配置值page_size` ，详见详见[*应用配置*](#应用配置)。
###maxPage
最大页数，`配置值max_page`，详见详见[*应用配置*](#应用配置)。
### minResidue
最小文件保留值，`配置值min_residue`，详见详见[*应用配置*](#应用配置)。
### 移动端获得最新打包版本
#####接口说明
移动设备会根据`userAgent`自动显示出该手机平台所需要的最新安装包。

#####URL
/api/[**[app]**](#app)/mobile/last
#####请求方式
GET
#####请求参数
无
#####返回结果
```json
{
  "code": 0,
  "msg": "",
  "items": [
    {
      "platform": "android",
      "environment": "dev",
      "version": "1.0.11",
      "time": "2016-8-9 00:00:11",
      "description": "测试描述11",
      "url": "http://example.me11",
      "down": "/download/rog2/dev/android/data/1.0.11/d1.0.11.apk",
      "channel": "none"
    },
    {
      "platform": "android",
      "environment": "release",
      "version": "1.0.12",
      "time": "2016-8-9 00:00:12",
      "description": "小测试描述androidrc12",
      "url": "http://example.me12",
      "down": "/download/rog2/release/android/data/1.0.12/r1.0.12.apk",
      "channel": "mi"
    }
  ]
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|
|[platform](#platform)|item| 平台|  string|     是|
|[environment](#environment)|item| 环境|  string|     是|
|version|item| 版本号|  string|     是|
|time|item| 打包时间|  string|     是|
|description|item| 版本描述|  string|     是|
|url|item| Jenkins构建url|  string|     是|
|down|item| 下载地址|  string|     是|
|channel|item| 渠道|  string|     是|
### 移动端获取历史列表
#####接口说明
移动设备会根据`userAgent`自动显示出该手机平台所需要的历史列表。限制[**maxPage**](#maxPage),[**pageSize**](#pageSize)

#####URL
/api/[**[app]**](#app)/mobile/list/[**[environment]**](#environment)?page=1
#####请求方式
GET
#####请求参数
| 参数名     |     含义|   参数类型| 是否必须| 默认值|
| :-------- | --------:| :------:|:------:|:------:|
| page       |   页数|  int| 是| 1|
#####返回结果
```json
{
  "code": 0,
  "msg": "",
  "items": [
    {
      "platform": "android",
      "environment": "dev",
      "version": "1.0.11",
      "time": "2016-8-9 00:00:11",
      "description": "测试描述11",
      "url": "http://example.me11",
      "down": "/download/rog2/dev/android/data/1.0.11/d1.0.11.apk",
      "channel": "none"
    },
	........,
	........,
	........,
    {
      "platform": "android",
      "environment": "dev",
      "version": "1.0.2",
      "time": "2016-8-9 00:00:02",
      "description": "测试描述2",
      "url": "http://example.me2",
      "down": "/download/rog2/dev/android/data/1.0.2/d1.0.2.apk",
      "channel": "none"
    }
  ],
  "page": 1,
  "total_page": 2
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|
|page| 无    | 当前页码|  int|        是|
|total_page| 无   | 总页数|  int|     是|
|[platform](#platform)|item| 平台|  string|     是|
|[environment](#environment)|item| 环境|  string|     是|
|version|item| 版本号|  string|     是|
|time|item| 打包时间|  string|     是|
|description|item| 版本描述|  string|     是|
|url|item| Jenkins构建url|  string|     是|
|down|item| 下载地址|  string|     是|
|channel|item| 渠道|  string|     是|
### 获得最新打包版本
#####接口说明
pc端获得所有平台的最新安装包。*`ios`的`dev，release`版本，`android`的`dev，release`版本*

#####URL
/api/[**[app]**](#app)/last
#####请求方式
GET
#####请求参数
无
#####返回结果
```json
{
  "code": 0,
  "msg": "",
  "items": [
    {
      "platform": "android",
      "environment": "dev",
      "version": "1.0.11",
      "time": "2016-8-9 00:00:11",
      "description": "测试描述11",
      "url": "http://example.me11",
      "down": "/download/rog2/dev/android/data/1.0.11/d1.0.11.apk",
      "channel": "none"
    },
    {
      "platform": "android",
      "environment": "release",
      "version": "1.0.12",
      "time": "2016-8-9 00:00:12",
      "description": "小测试描述androidrc12",
      "url": "http://example.me12",
      "down": "/download/rog2/release/android/data/1.0.12/r1.0.12.apk",
      "channel": "mi"
    },
    {
      "platform": "ios",
      "environment": "dev",
      "version": "1.0.0",
      "time": "2016-8-9 00:00:00",
      "description": "小测试描述ios",
      "url": "http://example.me",
      "down": "https://appinhouse.rog2.org/download/rog2/dev/ios/data/1.0.0/d11.plist",
      "channel": "appstore"
    },
    {
      "platform": "ios",
      "environment": "release",
      "version": "1.0.0",
      "time": "2016-8-9 00:00:00",
      "description": "小测试描述iosre",
      "url": "http://example.me",
      "down": "https://appinhouse.rog2.org/download/rog2/release/ios/data/1.0.0/11.plist",
      "channel": "appstore"
    }
  ]
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|
|[platform](#platform)|item| 平台|  string|     是|
|[environment](#environment)|item| 环境|  string|     是|
|version|item| 版本号|  string|     是|
|time|item| 打包时间|  string|     是|
|description|item| 版本描述|  string|     是|
|url|item| Jenkins构建url|  string|     是|
|down|item| 下载地址|  string|     是|
|channel|item| 渠道|  string|     是| 
### 获取历史列表
#####接口说明
pc端获得不同平台，不同环境的历史版本。限制[**maxPage**](#maxPage),[**pageSize**](#pageSize)
#####URL
/api/[**[app]**](#app)/list/[**[platform]**](#platform)/[**[environment]**](#environment)?page=1
#####请求方式
GET
#####请求参数
| 参数名     |     含义|   参数类型| 是否必须| 默认值|
| :-------- | --------:| :------:|:------:|:------:|
| page       |   页数|  int| 是| 1|
#####返回结果
```json
{
  "code": 0,
  "msg": "",
  "items": [
    {
      "platform": "android",
      "environment": "dev",
      "version": "1.0.11",
      "time": "2016-8-9 00:00:11",
      "description": "测试描述11",
      "url": "http://example.me11",
      "down": "/download/rog2/dev/android/data/1.0.11/d1.0.11.apk",
      "channel": "none"
    },
	........,
	........,
	........,
    {
      "platform": "android",
      "environment": "dev",
      "version": "1.0.2",
      "time": "2016-8-9 00:00:02",
      "description": "测试描述2",
      "url": "http://example.me2",
      "down": "/download/rog2/dev/android/data/1.0.2/d1.0.2.apk",
      "channel": "none"
    }
  ],
  "page": 1,
  "total_page": 2
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|
|page| 无    | 当前页码|  int|        是|
|total_page| 无   | 总页数|  int|     是|
|[platform](#platform)|item| 平台|  string|     是|
|[environment](#environment)|item| 环境|  string|     是|
|version|item| 版本号|  string|     是|
|time|item| 打包时间|  string|     是|
|description|item| 版本描述|  string|     是|
|url|item| Jenkins构建url|  string|     是|
|down|item| 下载地址|  string|     是|
|channel|item| 渠道|  string|     是|
### 生成plist
#####接口说明
生成plist文件。生成前需先先上传必要文件。详见[*客户端（ipa/apk)的归档*](#客户端的归档)
#####URL
/api/[**[app]**](#app)/plist/[**[environment]**](#environment)
#####请求方式
POST
#####请求参数
| 参数名     |     含义|   参数类型| 是否必须| 默认值|
| :-------- | --------:| :------:|:------:|:------:|
| version|   打包版本(与描述文件一致) | string| 是| 否|
| id|  metadata中的bundle-identifier|  string| 是| 否|
| title|  metadata中的标题|  string| 是| 否|
#####返回结果
```json
{
  "code": 0,
  "msg": "",
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|
### 整理历史版本
#####接口说明
删除过久的历史版本，只保留最新的一些版本。residue的限制[**minResidue**](#minresidue)
#####URL
/api/[**[app]**](#app)/delete/[**[platform]**](#platform)/[**[environment]**](#environment)?residue=10
#####请求方式
DELETE
#####请求参数
| 参数名     |     含义|   参数类型| 是否必须| 默认值|
| :-------- | --------:| :------:|:------:|:------:|
| residue|   剩余的数量| int| 是| 否|
#####返回结果
```json
{
  "code": 0,
  "msg": "",
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|
### 返回码


| 错误代码   |     详细描述 |    
| :-------- |  --------:|
| 0 |     成功| 
|-1| 系统错误    |
|-2| 未知错误    |
|1001 |   参数错误| 
|1002| 文件不存在   |
|1003 |     page超出最大页数限制| 
|1004| 解析描述文件出错    |
|1005 |     创建文件或文件夹错误| 
|1006| ipa文件不存在    |
|1007 |     整理文件出错| 
|1008| 文件路径不正确   |

