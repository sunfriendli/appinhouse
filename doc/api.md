# 企业内部appinhouse服务API接口文档

## 接口通用信息

1. 接口基于HTTP RESTFul思想架构设计。
2. 接口地址（URL）统一采用：http://域名/api/请求资源?过滤参数的格式。
	例如：`http://域名/api/rog2/list?page=1`。
3. 请求参数默认为form表单格式传输（`application/x-www-form-urlencoded`），响应内容默认JSON格式返回。

## 接口索引

### 显示类

#### 通用

- [获得应用列表](#获得应用列表)

#### 移动端

- [移动端获得最新打包版本](#移动端获得最新打包版本)
- [移动端获取历史列表](#移动端获取历史列表)

#### pc端

- [获得最新打包版本](#获得最新打包版本)
- [获取历史列表](#获取历史列表)



#### safari（对开发人员透明）

- [获得plist](#获得plist)

### 归档类

- [创建应用](#创建应用)
- [修改应用](#修改应用)
- [删除应用](#删除应用)
- [生成描述文件](#生成描述文件)
- [整理历史版本](#整理历史版本)





## 接口详情

### 接口公共信息

### app

接入方调用API生成应用，见[创建应用](#创建应用)。

### platform

移动端手机设备的操作系统名字，`枚举值`。

| value     |     ios |   android |
| :-------- | --------:| :------: |
| 含义    |   苹果设备|  安卓设备  |

### environment

打包的应用程序所对应的运行环境，`枚举值`

| value     |     dev|   release|
| :-------- | --------:| :------:|
| 含义       |   开发环境|  发布环境|

### page_size

一页显示多少记录，`配置值page_size` ，详见详见[***应用配置文档***](conf.md#page_size)。

### max_page

最大页数，`配置值max_page`，详见详见[***应用配置文档***](conf.md#max_page)。

### min_residue

最小文件保留值，`配置值min_residue`，详见详见[***应用配置文档***](conf.md#min_residue)。

### ios_channel

苹果官方渠道值，`ios_channel`，详见详见[***应用配置文档***](conf.md#ios_channel)。

### 时间格式

`2016-05-10T02:11:50+0800`

### time_offset

传参:根据偏移量计算时间。
不传：显示生成时所传的时间。

```js
    var d = new Date()
    var offset=-1*d.getTimezoneOffset()
```

### 创建应用

##### 接口说明

接入方创建自己的应用

##### URL

/api/[**[app]**](#app)/create

##### 请求方式

**POST**

##### 请求参数


| 参数名     |     含义|   参数类型| 是否必须| 默认值| 描述|
| :-------- | --------:| :------:|:------:|:------:|:------:|
|description  |   描述|  string| 是| 无||

##### 请求实例

```
POST /api/rog2/create

Host: 域名

Content-Type: application/x-www-form-urlencoded

description=xxxxxx
```


##### 返回结果

```json
{
  "code": 0,
  "msg": ""
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|


### 修改应用

##### 接口说明

接入方修改自己的应用

##### URL

/api/[**[app]**](#app)/update

##### 请求方式

**POST**

##### 请求参数


| 参数名     |     含义|   参数类型| 是否必须| 默认值| 描述|
| :-------- | --------:| :------:|:------:|:------:|:------:|
|description  |   描述|  string| 是| 无||

##### 请求实例

```
POST /api/rog2/update

Host: 域名

Content-Type: application/x-www-form-urlencoded

description=xxxxxx

```


##### 返回结果

```json
{
  "code": 0,
  "msg": ""
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|


### 删除应用

##### 接口说明

接入方删除自己的应用

##### URL

/api/[**[app]**](#app)/delete

##### 请求方式

**DELETE**

##### 请求参数

无


##### 请求实例

```
DELETE /api/rog2/delete

Host: 域名
```

##### 返回结果

```json
{
  "code": 0,
  "msg": ""
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|

### 获得应用列表

##### 接口说明

获取所有应用的列表。

##### URL

/api/apps

##### 请求方式

**GET**

##### 请求参数

无

##### 请求实例

```
GET /api/apps

Host: 域名

```

##### 返回结果

```json
{
  "code": 0,
  "msg": "",
  "items": [
    "rog1",
    "rog2",
    "rog3",
    "rog4",
  ],
  "page": 1,
  "total_page": 1
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|
|应用名|item| 平台|  string|     是|


### 移动端获得最新打包版本

##### 接口说明

移动设备会根据`userAgent`自动显示出该手机平台所需要的最新安装包。

##### URL

/api/[**[app]**](#app)/mobile/last?time_offset=480

##### 请求方式

**GET**

##### 请求参数

| 参数名     |     含义|   参数类型| 是否必须| 默认值| 描述|
| :-------- | --------:| :------:|:------:|:------:|:------:|
|[time_offset](#time_offset)  |   时间差|  int| 否| 无|不传只显示生成时的时间|

##### 请求实例

```
GET /api/rog2/mobile/last?time_offset=480

Host: 域名

```

##### 返回结果

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
      "down": "http://xxxxx/xxxx/xxxx.apk",
      "channel": "none"
    },
    {
      "platform": "android",
      "environment": "release",
      "version": "1.0.12",
      "time": "2016-8-9 00:00:12",
      "description": "小测试描述androidrc12",
      "url": "http://example.me12",
      "down": "http://xxxxx/xxxx/xxxx.apk",
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

##### 接口说明

移动设备会根据`userAgent`自动显示出该手机平台所需要的历史列表。限制[**max_page**](#max_page),[**page_size**](#page_size)

##### URL

/api/[**[app]**](#app)/mobile/list/[**[environment]**](#environment)?page=1&time_offset=480
##### 请求方式

**GET**

##### 请求参数

| 参数名     |     含义|   参数类型| 是否必须| 默认值| 描述|
| :-------- | --------:| :------:|:------:|:------:|:------:|
| page       |   页数|  int| 是| 无||
|[time_offset](#time_offset)  |   时间差|  int| 否| 无|不传只显示生成时的时间|


##### 请求实例

```
GET /api/rog2/mobile/list/dev?page=1&time_offset=480

Host: 域名

```

##### 返回结果

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
      "down": "http://xxxxx/xxxx/xxxx.apk",
      "channel": "none"
    },
    {
      "platform": "android",
      "environment": "dev",
      "version": "1.0.2",
      "time": "2016-8-9 00:00:02",
      "description": "测试描述2",
      "url": "http://example.me2",
      "down": "http://xxxxx/xxxx/xxxx.apk",
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

##### 接口说明

pc端获得所有平台的最新安装包。*`ios`的`dev，release`版本，`android`的`dev，release`版本*

##### URL

/api/[**[app]**](#app)/last?time_offset=480

#####请求方式

**GET**

##### 请求参数

| 参数名     |     含义|   参数类型| 是否必须| 默认值| 描述|
| :-------- | --------:| :------:|:------:|:------:|:------:|
|[time_offset](#time_offset)  |   时间差|  int| 否| 无|不传只显示生成时的时间|

##### 请求实例

```
GET /api/rog2/last?time_offset=480

Host: 域名

```
##### 返回结果

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
      "down": "http://xxxxx/xxxx/xxxx.apk",
      "channel": "none"
    },
    {
      "platform": "android",
      "environment": "release",
      "version": "1.0.12",
      "time": "2016-8-9 00:00:12",
      "description": "小测试描述androidrc12",
      "url": "http://example.me12",
      "down": "http://xxxxx/xxxx/xxxx.apk",
      "channel": "mi"
    },
    {
      "platform": "ios",
      "environment": "dev",
      "version": "1.0.0",
      "time": "2016-8-9 00:00:00",
      "description": "小测试描述ios",
      "url": "http://example.me",
      "down": "https://appinhouse.rog2.org/api/rog2/dev/ios/1.0.0.plist",
      "channel": "appstore"
    },
    {
      "platform": "ios",
      "environment": "release",
      "version": "1.0.0",
      "time": "2016-8-9 00:00:00",
      "description": "小测试描述iosre",
      "url": "http://example.me",
      "down": "https://appinhouse.rog2.org/api/rog2/release/ios/1.0.0.plist",
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

##### 接口说明

pc端获得不同平台，不同环境的历史版本。限制[**max_page**](#max_page),[**page_size**](#page_size)。

##### URL

/api/[**[app]**](#app)/list/[**[platform]**](#platform)/[**[environment]**](#environment)?page=1&time_offset=480

##### 请求方式

**GET**

##### 请求参数

| 参数名     |     含义|   参数类型| 是否必须| 默认值| 描述|
| :-------- | --------:| :------:|:------:|:------:|:------:|
| page       |   页数|  int| 是| 无||
|[time_offset](#time_offset)  |   时间差|  int| 否| 无|不传只显示生成时的时间|

##### 请求实例

```
GET /api/rog2/list/android/dev?page=1&time_offset=480

Host: 域名

```

##### 返回结果

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
      "down": "http://xxxxx/xxxx/xxxx.apk",
      "channel": "none"
    },
    {
      "platform": "android",
      "environment": "dev",
      "version": "1.0.2",
      "time": "2016-8-9 00:00:02",
      "description": "测试描述2",
      "url": "http://example.me2",
      "down": "http://xxxxx/xxxx/xxxx.apk",
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

### 获得plist

##### 接口说明

获得plist文件。苹果浏览器调用。

##### URL

/api/[**[app]**](#app)/plist/[**[environment]**](#environment)/version.plist

##### 请求方式

**GET**

##### 请求参数

无


##### 请求实例

```
GET /api/rog2/plist/dev/1.1.1.plist

Host: 域名

```

##### 返回结果

```json 
  "msg": "",
}
```

| 返回参数    |     父属性 |     含义 |   参数类型 |是否必须|
| :-------- |  --------:| --------:| :------: |:------: |
| [code](#返回码) |     无| 返回码    |  int|        是|
|msg             | 无    | 提示信息|  string|        否|

### 整理历史版本

#### 接口说明

删除过久的历史版本`(不包含归档文件)`，只保留最新的一些版本。residue的限制[**min_residue**](#min_residue)

##### URL

/api/[**[app]**](#app)/delete/[**[platform]**](#platform)/[**[environment]**](#environment)

#####请求方式

**DELETE**

##### 请求参数

| 参数名     |     含义|   参数类型| 是否必须| 默认值|
| :-------- | --------:| :------:|:------:|:------:|
| residue|   剩余的数量| int| 是| 否|


##### 请求实例

```
DELETE /api/rog2/delete/dev/1.0.0.plist?residue=10

Host: 域名

```

##### 返回结果

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

### 生成描述文件

##### 接口说明

生成描述文件。生成前需要有归档文件。如果是ios，则同时生成plist。详见[*客户端（ipa/apk)的归档*](archive.md#归档步骤)

##### URL

/api/[**[app]**](#app)/desc/[**[platform]**](#platform)/[**[environment]**](#environment)

##### 请求方式

**POST**

##### 请求参数

| 参数名     |     含义|   参数类型| 是否必须| 默认值|
| :-------- | --------:| :------:|:------:|:------:|
| version|   打包版本 | string| 是| 否|
| [time](#时间格式)|  打包时间|  string| 是| 否|
| [channel](#ios_channel)|  渠道|  string| 是| 否|
| url|  构建url|  string| 是| 否|
| software_url| ipa下载地址|  string| 是| 否|
| id|  plist中bundle-identifier|  string| platform==ios&&channel==appstore?是:否| 否|
| title| plist中title|  string| platform==ios&&channel==appstore?是:否| 否|
| full_url| plist中full-size-image|  string| 否| 否|
| display_url|  plist中display-image|  string| 否| 否|
| software_url_extend_name|  ipa扩展下载地址的名字，不为空时software_url为扩展下载地址|  string| 否| 否|
| software_url_extend_key|  ipa扩展下载地址的不支持中文，name的英文或拼音|  string| 否| 否|


##### 请求实例

```
POST /api/rog2/desc/ios/dev 

Host: 域名

Content-Type: application/x-www-form-urlencoded

version=1.0.2&url=http://www.xxxx.xxx
&channel=appstore&time=2016-05-10T02:11:50Z
&description=这个版本修改了bug
&software_url=https://appinhouse.rog2.org/download/rog2/dev/ios/data/1.0.0/ROG2NewD_1.0.0.ipa
&id=aaa&title=aaaa&software_url_name=上海&software_url_key=ShangHai

```
##### 返回结果

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
|1001 |   参数错误，请参考API文档| 
|1002|  page超出最大页数限制   |
|1003 |   创建plist文件错误 | 
|1004| 归档文件不存在     |
|1005 |    删除文件错误  | 
|1006| 时间格式错误|
|1007 | 数据库错误  | 
|1008 | 应用不存在  | 
|1009 | 应用已存在  | 

