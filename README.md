# sign

自动签到程序, 纯属娱乐.

## 目前功能

目前实现的自动签到的网站:

- [Go 语言中文网](https://studygolang.com/)
- [B 站](https://www.bilibili.com/)
- [千图网](https://www.58pic.com/)
- [黑客派](https://hacpai.com/)
- [v2ex](https://v2ex.com/)
- [爱奇艺](https://www.iqiyi.com/)

特性:

- 随机时间签到 (8:30~20:30)
- 签到失败通知
- 刷活跃度
    - Go 语言中文网
        - 1~20名之间, 可以手动指定, 超过范围的排名将会使用0 (随机)

## 安全性

1. 不收集任何数据

## 部署

1. 下载

```
$ git clone https://github.com/jdxj/sign.git
```

2. 编译 (linux)

```
$ cd sign
$ make
```

3. 根据格式创建 `sign.ini` 配置文件 (与 `sign.out` 在同级目录即可)

```
# sign.ini 格式

# 用于邮件通知, 使用 QQ 邮箱
[email]

username = 985759262@qq.com
password = # 授权码

# 用于 http api basic auth 授权
[basicauth]

username =
password =
```

4. 启动

```
$ ./sign.out &
```

## 添加任务

- http api 格式 (Content-Type: application/json)

```
// 我部署的域名: http://sign.aaronkir.xyz

// 使用时请去掉注释
// 每个 json 中的 name 只是一个标识, 随便取

// POST /api/studygolang
{
  "name": "StudyGolang",
  "username": "985759262@qq.com",
  "password": "",
  // 随便一个网页就行, 这里选取个人主页刷活跃度
  "activeURL": "https://studygolang.com/user/jdxj",
  "expected": 10,
  "to": "985759262@qq.com"
}

// POST /api/bilibili
{
  "name": "Bilibili",
  "cookies": "",
  // 这里验证是否登录成功的方法是向服务器请求了你的关注数量, 我关注了9个人
  "verify_value": 9,
  "to": "985759262@qq.com"
}

// POST /api/58pic
{
  "name": "58Pic",
  "cookies": "",
  "to": "985759262@qq.com"
}

// POST /api/hacpai
{
  "name": "HacPai",
  "username": "985759262@qq.com",
  "password": "",
  "to": "985759262@qq.com"
}

// POST /api/v2ex
{
  "name": "V2ex",
  // v2ex 的 cookie 在从浏览器中手动复制时发现其带有双引号,
  // 我已在程序中做了过滤处理, 如果你使用 postman,
  // 那么需要手动删除双引号 (其自己的语法检查).
  "cookies": "",
  "to": "985759262@qq.com"
}

// POST /api/iqiyi
{
  "name": "IQiYi",
  "cookies": "",
  "check_in_sign": "",
  "hot_spot_sign": "",
  "to": "985759262@qq.com"
}
```

## TODO

- **使用 nginx 开启 https**
- 优化细节
- 支持更多网站
- 丰富邮件提醒功能
- 扫码登录?
- 整合扫码登录 [wxlogin](https://github.com/jdxj/wxlogin)
- 重构, 使签到对象更好的管理
- 使用 http api 创建签到任务, 从配置读取的方式将被弃用
- 任务管理?
- 刷 hacpai 活跃度

## 已知的问题

- ~~目标服务器与签到程序服务器之间的时间会有误差, 理论上会有漏签问题~~ (由于改了签到时间范围, 这个问题不会出现)
- 由于千图网需要每周手动登录, 所以千图网会由于 cookie 失效而签到失败 (正在尝试解决)
- v2ex 可能在早上几点之后才会更新签到链接
