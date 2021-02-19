# 有道词典命令行版服务端
## 使用方法
替换掉conf/app.toml中的应用id和应用密钥
申请地址在这里
[申请地址](http://ai.youdao.com/buy.s?productId=2&productName=%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1)

## 注意事项
我添加了平滑重启功能，用的是这个包
```github.com/fvbock/endless```
由于windows系统没有某些kill信号，所有本项目无法在windows下运行，推荐使用linux系统