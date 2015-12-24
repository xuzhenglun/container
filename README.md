# container

基于Golang编写的一个符合REST标准的Webservice

#API说明
```
1. GET方法：获取上传在服务器的文件，无需验证身份。
2. POST方法：上传一个文件到服务器，模拟表单的方式进行POST（表单名：file）。
    **需要在HTTP Header中指定Username，若不符合身份则拒绝。**

3. DELETE方法：删除一个文件，**需验证Username**
```
如果使用Golang，Post，Delete已经封装为函数，在sdk文件夹下。
```
import github.com/xuzhenglun/container/sdk
```

#服务器返回说明
 

 1. 状态码：

 ```
const (
	HaveNoRright      = "403"
	PostSuccess       = "200"
	DeleteSuccsess    = "201"
	ServerHandleError = "400"
)
```

 2. url段：

POST成功后返回在服务器的文件名，DELETE成功后返回删除的文件名。
