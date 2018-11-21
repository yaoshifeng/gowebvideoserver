# gowebvideoserver

it's a RESTful video server project.<br>

# 项目整体流程

## 控制器-->校验(1用户名是否合法 2用户是否已注册)-->商业逻辑-->应答response

## 用户模块
用户名 user_name<br>
密码   pwd<br>

### 2018/11/21 <br>
新建main.go     包含：主程序入口，包含路由，登录url初选，端口监听<br>
新建handlers.go 包含：登录url初选进入后的处理<br>
新建defs文件夹   包含：<br>
新建defs/errs.go 包含：错误处理<br> 
新建defs/apidef.go 包含：用户登录api的整体方法<br>
