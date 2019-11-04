# sign

自动签到程序, 纯属娱乐.

## 目前功能

目前实现的自动签到的网站:

- [Go 语言中文网](https://studygolang.com/)
- [B 站](https://www.bilibili.com/)
- [千图网](https://www.58pic.com/)

特性:

- 随机时间签到
- 签到失败通知

## 用法

1. 下载

```
$ git clone https://github.com/jdxj/sign.git
```

2. 编译 (linux)

```
$ cd sign
$ go build -o sign.out *.go
```

3. 根据格式创建 `sign.ini` 配置文件 (与 `sign.out` 在同级目录即可)
4. 启动

```
$ ./sign.out &
```

## 原理

使用 [colly](https://github.com/gocolly/colly) 爬虫访问签到链接. 使用 colly 的原因是: 爬取指定信息来确认是否签到成功.

## 配置文件格式

配置文件名称: `sign.ini`.

配置文件格式:

```
# 邮件通知配置
# 目前只使用了 QQ 邮箱
[email]

username = yourname
password = yourpassword

# 使用账户-密码
[studygolang.com]

username = yourname
passwd = yourpasswd

# 使用 cookie
[bilibili.com]

key = value

# 使用 cookie
[58pic.com]

key = value
```

## TODO

- 优化细节
- 支持更多网站
- 丰富邮件提醒功能
- 扫码登录?
- 为了更灵活, 不使用 colly, 而直接使用 [goquery](https://github.com/PuerkitoBio/goquery)
- 整合扫码登录 [wxlogin](https://github.com/jdxj/wxlogin)
- 完善抽象

## 已知的问题

- 目标服务器与签到程序服务器之间的时间会有误差, 理论上会有漏签问题
- 由于千图网需要每周手动登录, 所以千图网会由于 cookie 失效而签到失败
