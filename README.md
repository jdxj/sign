# sign

自动签到程序, 纯属娱乐.

## 目前功能

目前实现的自动签到的网站:

- [Go 语言中文网](https://studygolang.com/)
- [B 站](https://www.bilibili.com/)
- [千图网](https://www.58pic.com/)

## 原理

使用 [colly](https://github.com/gocolly/colly) 爬虫访问签到链接. 使用 colly 的原因是: 爬取指定信息来确认是否签到成功.

## 配置文件格式

配置文件名称: `sign.ini`.

配置文件格式:

```
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
