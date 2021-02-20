# 有道词典命令行版服务端
## 使用方法
替换掉conf/app.toml中的应用id和应用密钥
申请地址在这里
[申请地址](http://ai.youdao.com/buy.s?productId=2&productName=%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1)

替换完之后执行

```go build```

接着再执行（我把输出的信息都放到exec_log/exec.log中了），之后就可以通过接口进行请求了

```./youdao >exec_log/exec.log 2>& 1 &```

## 接口说明

**地址**
- `fanyi/youdao`

**请求方式**
- POST (JSON形式)

**请求参数**

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|content |是  |string |待翻译的文本   |

**返回参数** 

**参数说明[看这里](https://ai.youdao.com/DOCSIRMA/html/%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E7%BF%BB%E8%AF%91/API%E6%96%87%E6%A1%A3/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1-API%E6%96%87%E6%A1%A3.html)**
```shell
{
    "code": 0,
    "data": {
        "list": {
            "errorCode": "0",
            "query": "pack",
            "translation": [
                "包"
            ],
            "basic": {
                "phonetic": "pæk",
                "uk-phonetic": "pæk",
                "us-phonetic": "pæk",
                "explains": [
                    "n. 包装；一群；背包；包裹；一副",
                ]
            },
            "web": [
                {
                    "key": "pack",
                    "value": [
                        "剥撕式面膜"
                    ]
                },
                {
                    "key": "service pack",
                    "value": [
                        "服务包"
                    ]
                },
                {
                    "key": "expansion pack",
                    "value": [
                        "资料片"
                    ]
                }
            ],
            "l": "en2zh-CHS",
            "returnPhrase": [
                "pack"
            ],
            "isWord": true
        },
        "msg": "查询成功"
    },
    "status": "success",
    "taketime": 28.701952
}
```

## 注意事项
我添加了平滑重启功能，用的是这个包
```github.com/fvbock/endless```
由于windows系统没有某些kill信号，所有本项目无法在windows下运行，推荐使用linux系统