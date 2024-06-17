## 自定义配置
您可以在您的程序运行的`根目录创建此文件`，程序会自动读取。
完整参数详见本目录的`config.yaml` 文件。

### 限流配置
您可以对您的搜题接口针对IP进行限流，防止恶意请求。默认是关闭的状态。
```yaml
limit:
  enable: false # 是否开启
  duration: 3  # 时间窗口为3秒
  requests: 1  # 允许用户在3秒内通过1个请求
```

### 数据库配置
您可以对您的数据库配置，默认使用sqlite3，您可以指定mysql为您的数据库。
一种是在配置文件中指定，另外一种是在指定环境变量`SQL_DSN`。
```yaml
mysql: "tiku:tiku@tcp(127.0.0.1:3306)/tiku?charset=utf84&parseTime=True&loc=Local"
```


### 自定义外部题库配置
当然您也可以接入外部的第三方题库，通过简单的配置，如下给出两个demo。
其中需要注意的是：
- `url`字段是您的题库的接口地址。
- `method`字段是您的题库的请求方法。
- `headers`字段是您的题库的请求头，您可以通过`key`和`value`来设置。必须设置正确`Content-Type`，否则可能无法正常请求。
- `body`字段是您的题库的请求体，您可以通过`${question}`来获取用户的问题。支持的占位符有 `${question}`，`${type}`。
- `answer`字段是您的题库返回的答案的路径，您可以通过`jsonpath`语法来获取。

以下为您提供两个开箱即用的demo，您可以直接复制到您的`config.yaml`中，然后修改`enable`为`true`即可测试。但是您需要注意，以下两个题库均已经集成到tikuAdapter中，所以您无需再次配置它，当您测试完毕，将`enable`改为`false`即可。

```yaml
api:
  - name: '样例接口 json 请求'
    enable: false
    url: http://lyck6.cn/scriptService/api/autoFreeAnswer
    method: POST
    headers:
      - key: Content-Type
        value: application/json
    body: '{"question":"${question}"}'
    answer: 'result.answers.0.0'

  - name: '样例接口 表单请求'
    enable: false
    url: https://cx.icodef.com/wyn-nb?v=4
    method: POST
    headers:
      - key: Content-Type
        value: application/x-www-form-urlencoded
    body: 'question=${question}'
    answer: 'data'
```

### 自动录入题库功能
使用脚本答题或者请求API接口，将会自动记录您的问题和答案，可以通过 `recordEmptyAnswer` 来控制是否记录空答案。
```yaml
recordEmptyAnswer: true # 设置为 true 将会记录空答案否则只会记录已有答案
```

### 持久化到数据库
您可以将您的题库数据持久化到数据库，目前支持的数据库有`mysql`和`sqlite`，您可以通过`mysql`来设置。如果不配置，默认使用`sqlite`。
```yaml
mysql: username:password@tcp(localhost:3306)/databasename
```