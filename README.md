# gowebvideoserver

it's a RESTful video server project.<br>

# 项目整体流程

## 控制器-->校验(1用户名是否合法 2用户是否已注册)-->商业逻辑-->应答response

### 2018/11/21 <br>
#### api主模块
新建main.go     包含：主程序入口，包含路由，登录url初选，端口监听<br>
新建handlers.go 包含：登录url初选进入后的处理<br>
新建defs文件夹   数据结构<br>
新建defs/errs.go 包含：错误处理<br> 
新建defs/apidef.go 包含：用户登录api的逻辑<br>
新增dbops文件夹  数据库操作<br>
新建dbops/api.go  包含：数据库逻辑<br>


### 2018/11/22 <br>
#### 数据库设计<br><br>
用户表<br>
TABLE:users<br>
id         (每个用户有自己的ID，独立性，用户不可见，目前设计成简单递增，以后改成16位UUID)<br>
login_name (不可重名，数据库查找通过ID搜索，以后修改成可重名通过UUID搜索用户)<br>
pwd        (密码，目前先定位TEXT)<br>
<br>
id UNSIGNED INT, PRIMARY KEY, AUTO_INCREMENT<br>
login_name varchar(64) UNIQUE KEY<br>
pwd TEXT<br>
<br>
资源表(视频)<br>
TABLE:video_info<br>
id            (视频ID，设计成随机16位UUID，避免重复)<br>
author_id     (用户ID，唯一关联3个表的变量，具体约束通过代码，不用修改数据库结构)<br>
name	      (视频名字)<br>
display_ctime (页面创建时间，给用户看的)<br>
create_time   (video入库时间)<br>
<br>
id Varchar(64), PRIMARY KEY, NOT NULL<br>
author_id UNSIGNED INT<br>
name TEXT<br>
display_ctime TEXT<br>
create_time DATETIME<br>
<br>
评论表<br>
TABLE:comments<br>
id		(评论ID，设计成随机16位UUID，避免重复)<br>
video_id	(视频ID)<br>
author_id	(用户ID)<br>
content		(评论内容)<br>
create_time     (评论创建时间)<br>
<br>
id varchar(64), PRIMARY KEY, NOT NULL<br>
video_id varchar(64)<br>
author_id UNSIGNED INT<br>
content TEXT<br>
create_time DATETIME<br>
<br>
TABLE:sessions  (用于用户登录状态信息检索，代码需要判断过期与校验)<br>
session_id TINYTEXT, PRIMARY KEY, NOT NULL<br>
TTL TINYTEXT<br>
login_name VARCHAR(64)<br>

### 2018/11/23 <br>
#### 更新api主模块
新建utils文件夹<br>
新建uuid.go文件 包含：创建videoid和commitid的16位id函数<br>
在dbops文件夹<br>
新建conn.go 包含：把api.go文件中opendb的函数单独移到文件中，数据库不频繁开关，增加使用效率<br>
更新api.go 包含：新增用户，视频和评论的数据库交互函数框架，编写用户User的增，获取，删函数<br>
在defs文件夹<br>
更新errs.go 包含：新增错误处理，数据库错误变量和普通服务器错误变量<br>









