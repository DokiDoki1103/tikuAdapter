# 这是什么？

这是一个题库适配器，可以将各种的题库接口转换为统一的标准格式，同时这将会是 **题库接口的一个规范**

![image.png](https://img.cdn.apipost.cn/client/user/1010721/avatar/78805a221a988e79ef3f42d7c5bfd41865389e5a65048.png "image.png")

市面上有很多题库接口，比如您想要在 ```【万能】全平台自动答题脚本``` 使用```言溪enncy题库```您就可以通过此题库适配器轻松实现。

## 功能

### 1.支持多种题库接口输入源

排名顺序为免费优先于付费。

- [x] [icodef 题库](https://q.icodef.com) [![免费](https://img.shields.io/badge/-免费-brightgreen)](url)
- [x] [不挂科 题库](https://easylearn.baidu.com/edu-page/tiangong/bgklist) [![免费](https://img.shields.io/badge/-免费-brightgreen)](url)
- [x] [万能题库](https://lyck6.cn/pay) [![付费](https://img.shields.io/badge/免费-付费-brightgreen?color=red&labelColor=4c1)](https://lyck6.cn/pay)
- [x] [爱点题库](https://www.51aidian.com) [![付费](https://img.shields.io/badge/-付费-red)](https://tk.enncy.cn/)
- [x] [enncy 言溪题库](https://tk.enncy.cn/) [![付费](https://img.shields.io/badge/-付费-red)](https://tk.enncy.cn/)
- [x] [柠檬题库](https://www.lemtk.xyz)[![付费](https://img.shields.io/badge/-付费-red)](https://www.lemtk.xyz)
- [x] [自定义题库](https://github.com/DokiDoki1103/tikuAdapter/tree/main/configs#%E8%87%AA%E5%AE%9A%E4%B9%89%E5%A4%96%E9%83%A8%E9%A2%98%E5%BA%93%E9%85%8D%E7%BD%AE)

访问第三方题库接口时，自动将答案持久化到本地，下次搜题优先搜索本地题库的答案。

### 2.支持多种文件自动解析到题库

+ [ ] [Word题库](https://github.com/itihey/tikuAdapter/raw/main/test/test.docx)将您的word文档自动解析到本地数据库，方便您的搜索。

### 3. 支持多种输出源：

+ [x] [tikuAdapter标准格式](https://github.com/itihey/tikuAdapter#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B) **强烈推荐**
  您为您的软件适配标准格式
+ [ ] **微信公众号** 微信开发者后台填写tikuAdapter的url即可

### 4.个性化配置

- [自定义请求参数](https://github.com/itihey/tikuAdapter#url-%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0) 个性化**禁用题库**或者
  **配置题库Token**
- 搜题接口的限流措施(
  个人使用一般不需要开启) [配置限流](https://github.com/itihey/tikuAdapter/tree/main/configs#%E9%99%90%E6%B5%81%E9%85%8D%E7%BD%AE)

## 如何部署使用

### 自行部署

从 [GitHub Releases](https://github.com/itihey/tikuAdapter/releases) 下载对应的版本，解压后运行即可

### 使用在线服务

在线服务均不能保证稳定性，并可能会泄漏您的Token，强烈建议自行部署。

- 【v0.1.0-alpha.6】xmig提供 `http://adapter.xmig6.cn/adapter-service/search`

### 使用API接口

POST `http://localhost:8060/adapter-service/search`

#### 请求体

```json
{
  "question": "违反安全保障义务责任属于（）",
  "options": [
    "公平责任",
    "特殊侵权责任",
    "过错推定责任",
    "连带责任"
  ],
  "type": 1 // 单选0多选1填空2判断3问答4
}
```

#### URL 请求参数

| 参数             | 描述                  | 是否必须 | 示例值                              | Token获取方式                |
|----------------|---------------------|------|----------------------------------|--------------------------|
| wannengToken   | 万能付费题库的Token值(10位)  | 否    | E196FD8B49                       | https://lyck6.cn/pay     |
| localDisable | **是否禁用本地搜索**        | 否    | 1此值传1将禁用)                        |
| wannengDisable | 是否禁用万能题库            | 否    | 1此值传1将禁用)                        |
| icodefToken    | Icodef 题库Token值     | 否    | UafYcHViJMGzSVNh                 | 关注微信公众号"一之哥哥"发送"token"获取 |
| icodefDisable  | 是否禁用icodef题库        | 否    | 1(此值传1将禁用)                       |
| enncyToken     | enncy 题库Token值      | 否    | a21ae2403b414b94b512736c30c69940 | https://tk.enncy.cn      |
| enncyDisable   | 是否禁用enncy题库         | 否    | 1(此值传1将禁用)                       |
| buguakeDisable | 是否禁用不挂科题库           | 否    | 1(此值传1将禁用)                       |
| aidianDisable  | 是否禁用爱点题库            | 否    | 1 (此值传1将禁用)                      |
| aidianYToken   | 爱点题库(亿级题库API)Token值 | 否    | cvor7f3HxZ7nF2M3ljmA             | https://www.51aidian.com |
| lemonDisable   | 是否禁用柠檬题库            | 否    | 1 (此值传1将禁用)                      |
| lemonToken     | 柠檬题库 Token值 | 否    | 8a3debe92e2ba83d6786e186bef2a424             | https://www.lemtk.xyz    |

例如您想禁用万能题库并且想要使用icodef的token，您的url应为`http://localhost:8060/adapter-service/search?wannengDisable=1&icodefToken=UafYcHViJMGzSVNh`

#### 响应示例

```json
{
  "plat": 0,
  "question": "违反安全保障义务责任属于（）",
  "options": [
    "公平责任",
    "特殊侵权责任",
    "过错推定责任",
    "连带责任"
  ],
  "type": 1,
  "answer": {
    "answerKey": [
      "B",
      "C"
    ],
    "answerKeyText": "BC",
    "answerIndex": [
      1,
      2
    ],
    "answerText": "特殊侵权责任#过错推定责任",
    "bestAnswer": [
      "特殊侵权责任",
      "过错推定责任"
    ],
    "allAnswer": [
      [
        "特殊侵权责任",
        "过错推定责任"
      ],
      [
        "A特殊侵权责任",
        "B过错推定责任"
      ]
    ]
  }
}
```

## 如何贡献

#### 提出您的issue

将您的题库接口提issue我们为您增加上去。

#### 提出您的pr

您可以参与开发，提交pr。

- 您可以参考 ```internal/search/wanneng.go``` 来实现 ```internal/search/search.go```接口
- 编写 ```internal/search/search_test.go``` 来测试您的接口
