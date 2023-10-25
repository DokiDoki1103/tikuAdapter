# 这是什么？

这是一个题库适配器，可以将各种的题库接口转换为统一的标准格式，同时这将会是 **题库接口的一个规范**

![image.png](https://img.cdn.apipost.cn/client/user/1010721/avatar/78805a221a988e79ef3f42d7c5bfd41865389e5a65048.png "image.png")

## 为什么诞生它？

市面上有很多题库接口，比如您想要在 ```【万能】全平台自动答题脚本``` 使用```言溪enncy题库```您就可以通过此题库适配器轻松实现。

## 如何部署使用

[快速部署](https://github.com/itihey/tikuAdapter/blob/main/deploy/README.md)

### API 接口

POST `127.0.0.1:8060/adapter-service/search`

### Body

```json
{
  "question": "王牌飞行员"
}
```

### Query Parameters

| 参数             | 描述                    | 是否必须 | 示例值              | Token获取方式                |
|----------------|-----------------------|------|------------------|--------------------------|
| wannengToken   | 万能付费题库的Token值(10位)    | 否    | E196FD8B49       | https://lyck6.cn/pay     |
| wannengDisable | 是否禁用万能题库(此值传1将禁用)     | 否    | 1                |
| icodefToken    | Icodef 题库Token值       | 否    | UafYcHViJMGzSVNh | 关注微信公众号"一之哥哥"发送"token"获取 |
| icodefDisable  | 是否禁用icodef题库(此值传1将禁用) | 否    | 1                |

## Response
```json
{
    "plat": 0,
    "question": "健康妇女的阴道内尽管有细菌,但是不会影响到个人的安全",
    "options": [],
    "type": 0,
    "answerIndex": [],
    "answer": [
        "×"
    ],
    "moreAnswer": [
        [
            "对"
        ],
        [
            "正确"
        ],
        [
            "×"
        ],
        [
            "错"
        ],
        [
            "错误"
        ]
    ]
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
