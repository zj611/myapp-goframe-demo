# GoFrame Project demo

https://goframe.org

# 学习资料
官方学习文档[链接](https://goframe.org/pages/viewpage.action?pageId=1114245)
# 配置gf
gf工具用于初始化模版，便于快速搭建基础框架
下载[链接](https://goframe.org/cli)
指南[链接](https://goframe.org/pages/viewpage.action?pageId=1115782)
将下载到的gf安装包拷贝到`/usr/local/bin`
```bash
./gf install
```
安装成功后，可以看到

```bash
$ gf -v
GoFrame CLI Tool v0.7.5, https://goframe.org
Built Detail:
  Go Version:  go1.13.5
  GF Version:  v1.13.0
  Git Commit:  0ab2a4a9e753f4f5fae1e96ecf40cae94895d494
  Built Time:  2020-06-05 23:30:39
```

# 启动一个工程
## 初始化
自动创建一个项目名为`myapp`的脚手架工程
```bash
gf init myapp
```
有如下提示则创建成功

```bash
initializing...
initialization done! 
you can now run 'gf run main.go' to start your journey, enjoy!
```
## 自动生成dao/model层文件
dao命令用于生成dao数据访问对象文件，以及model数据结构定义文件。推荐使用配置文件来管理生成规则。参考[链接](https://goframe.org/pages/viewpage.action?pageId=3673173)
### 手动创建table

```sql
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `passport` varchar(45) NOT NULL COMMENT '用户账号',
  `password` varchar(45) NOT NULL COMMENT '用户密码',
  `nickname` varchar(45) NOT NULL COMMENT '用户昵称',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `account` (
  `account` varchar(45) NOT NULL COMMENT '用户账号',
  `balance` int(30) NOT NULL COMMENT '用户余额',
  `interest` float(30) NOT NULL COMMENT '用户利息',
  `transaction_time` datetime NOT NULL COMMENT '交易时间',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

### 修改config/config.toml
```bash
[database]
    link  = "mysql:root:123@tcp(127.0.0.1:3306)/data"
    debug = true
    # Database logger.
    [database.logger]
        Path   = "/tmp/log/gf-app/sql"
        Level  = "all"
        Stdout = true
```

### 使用gf gen dao命令
```bash
cd myapp
gf gen dao -path ./app -c config/config.toml  -t user
```
有如下提示则创建成功
```bash
generated: app/dao/internal/user.go
generated: app/model/model.go
done!
```

## 安装swagger
### 下载&安装
下载&安装[链接](https://github.com/swaggo/swag)
```bash
sudo cp ./swag /usr/local/bin
```
### 使用swagger
```bash
gf swagger  #生成swagger.json
gf swagger -s 127.0.0.1:8199  #拉起swagger服务端，点击页面可查看接口api
```




