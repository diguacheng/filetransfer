# FileTransfer  


这是一个基于go实现的用于在网络上传输文件的程序,支持对数据的压缩和加密。

关于包的引用： 
不引入外部第三方包，基于golang的官方包实现 

已经实现的功能： 
1. 采用deflate算法,实现对数据进行压缩 
2. 采用AES算法，对数据进行加密
3. 采用pbkdf2算法实现对AES密钥的派发.
4. 已经实现了接收端具有公网ip或者发送端和接收端为同一内网这两种场景。





todo： 
- 支持任意两个主机发送数据(若两边是不同公网下的内网ip,则需要开发中转程序部署在具有公网ip的主机上)  
- 支持按照文件夹传,能够实现传一个文件夹下的子文件夹和文件  
