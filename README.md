内容要求
功能
实现一个用户管理系统，用户可以登录、拉取和编辑他们的profiles。
用户可以通过在Web页面输入username和password登录，backend系统负责校验用户身份。成功登录后，页面需要展示用户的相关信息；否则页面展示相关错误。
成功登录后，用户可以编辑以下内容：
上传profile picture
修改nickname（需要支持Unicode字符集，utf-8编码）
用户信息包括：
username（不可更改）
nickname
profile picture
需要提前将初始用户数据插入数据库用于测试。确保测试数据库中包含10,000,000条用户账号信息。

开发环境
Server: 个人工作PC/Mac中的虚拟机
OS: CentOS 7 x64 or Ubuntu 14.04 above
DB: MySQL 5.5 or above
Client: Chrome and Firefox
注：如果对虚拟机不熟悉，可以直接在工作PC/Mac上开发；如果熟悉vagrant，可以直接使用ubuntu-dev。
注：在Mac进行性能测试可能存在系统资源限制，可以尝试以下命令：
sudo sysctl -w kern.ipc.somaxconn=2048
sudo sysctl -w kern.maxfiles=12288
ulimit -n 10000


设计要求
分别实现HTTP server和TCP server，主要的功能逻辑放在TCP server实现
Backend鉴权逻辑需要在TCP server实现
用户账号信息必须存储在MySQL数据库。通过MySQL Go client连接数据库
使用基于Auth/Session Token的鉴权机制，Token存储在redis，避免使用JWT等加密的形式。
TCP server需要提供RPC API，RPC机制希望自己设计实现
Web server不允许直连MySQL、Redis。所有HTTP请求只处理API和用户输入，具体的功能逻辑和数据库操作，需要通过RPC请求TCP server完成
尽可能使用Go标准库
安全性
鲁棒性
性能
交付
源代码
设计文档
部署、运维文档
性能测试报告
总结文档
现场演示
代码必须上传到git.garena.com。
文档尽可能使用Markdown和代码一起上传git.garena.com，或者使用google docs。
验收标准
时间：
尽可能在规定时间内完成
提前让mentor/team leader review代码
正确性：
必须完整实现相关API，不能有明显BUG
实现细节必须满足设计要求，从而达到Entry Task的目的
安全性：
不能有安全问题
鲁棒性：
服务不能因为客户端请求crash
性能：
数据库必须有10,000,000条用户账号信息
必须确保返回结果是正确的
每个请求都要包含RPC调用以及Mysql或Redis访问
200并发（固定用户）情况下，HTTP API QPS大于3000
200个client（200条TCP连接），每个client模拟一个用户（因此需要200个不同的固定用户账号）
200并发（随机用户）情况下，HTTP API QPS大于1000
200个client（200条TCP连接），每个client每次随机从10,000,000条记录中选取一个用户，发起请求（如果涉及到鉴权，可以使用一个测试用的token）


2000并发（固定用户）情况下，HTTP API QPS大于1500


2000并发（随机用户）情况下，HTTP API QPS大于800
代码规范：
通过 golint
通过 go vet
尽可能遵循Effective Go
代码质量：
易读
依赖清晰
尽量解耦
尽可能覆盖单元测试
文档：
交付的文档尽可能详细

其中，时间、正确性、安全性、性能是必须要达到要求的。其余几项在时间期限后，根据实际情况可以放宽要求。
