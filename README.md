# 这是什么？

这是一个题库适配器，可以将各种的题库接口转换为统一的标准格式，同时这将会是 **题库接口的一个规范**

![image.png](https://img.cdn.apipost.cn/client/user/1010721/avatar/78805a221a988e79ef3f42d7c5bfd41865389e5a65048.png "image.png")

市面上有很多题库接口，比如您想要在 ```【万能】全平台自动答题脚本``` 使用```言溪enncy题库```您就可以通过此题库适配器轻松实现。

## 功能
1. 支持多种输入源：
    + [x] [enncy 言溪题库](https://tk.enncy.cn/)
    + [x] [万能题库](https://lyck6.cn/pay)
    + [x] [icodef 题库](https://q.icodef.com)
2. 支持多种输出源：
    + [x] [tikuAdapter标准格式](https://github.com/itihey/tikuAdapter#%E5%93%8D%E5%BA%94%E7%A4%BA%E4%BE%8B) **强烈推荐**您为您的软件适配标准格式
    + [ ] **微信公众号** 微信开发者后台填写tikuAdapter的url即可

3. [自定义请求参数](https://github.com/itihey/tikuAdapter#url-%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0) 个性化**禁用题库**或者**配置题库Token**




## 如何部署使用

### 部署
从 [GitHub Releases](https://github.com/itihey/tikuAdapter/releases) 下载对应的版本，解压后运行即可

### 使用API接口

POST `localhost:8060/adapter-service/search`

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
  "type": 1
}
```

#### URL 请求参数

| 参数             | 描述                    | 是否必须 | 示例值                              | Token获取方式                |
|----------------|-----------------------|------|----------------------------------|--------------------------|
| wannengToken   | 万能付费题库的Token值(10位)    | 否    | E196FD8B49                       | https://lyck6.cn/pay     |
| wannengDisable | 是否禁用万能题库(此值传1将禁用)     | 否    | 1                                |
| icodefToken    | Icodef 题库Token值       | 否    | UafYcHViJMGzSVNh                 | 关注微信公众号"一之哥哥"发送"token"获取 |
| icodefDisable  | 是否禁用icodef题库(此值传1将禁用) | 否    | 1                                |
| enncyToken     | enncy 题库Token值       | 否    | a21ae2403b414b94b512736c30c69940 | https://tk.enncy.cn |
| enncyDisable   | 是否禁用enncy题库(此值传1将禁用) | 否    | 1                                |

例如您想禁用万能题库并且想要使用icodef的token，您的url应为`127.0.0.1:8060/adapter-service/search?wannengDisable=1&icodefToken=UafYcHViJMGzSVNh`

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

## 计划如下

### 大方向分为

- 支持多题库接口输入源，支持格式化为标准格式或常见格式 ✅
- 将题库私有化持久化 ❌

### 支持多题库接口输入源

其他题库接口或者提供商欢迎Pr或者issue，我们会将其加入到适配器中。

- 【言溪enncy】OCS 网课助手 ❌
- 【万能】全平台自动答题脚本 ✅
- 【一之哥哥】icodef 题库 ✅
