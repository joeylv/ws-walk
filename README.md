# ws-walk
无舍manager
go get github.com/akavel/rsrc


打包

go build

无CMD

go build -ldflags="-H windowsgui"



第一步：Windows 下载MinGW

第二步：新建一个.rc文件，加入文件名为 demo.rc 输入内容   

IDI_ICON1 ICON "cefclient.ico"
其中 cefclient.ico 是你的ico的地址

第三步:MinGW 执行 

windres -o demo.syso demo.rc
需要将demo.syso放到需要编译的go目录下

第四步： go build 编译出exe程序

第五步：需要一个工具，叫做 ResourceHacker ， 可以在网上直接搜索下载

第六步：使用ResourceHacker打开编译出的exe程序，点击添加使用脚本模板

第七步：在弹出框里选择VERSION_INFO

第八步: 在新建的文件中修改信息即可，信息的字段说明可以参考如下地址内容

https://msdn.microsoft.com/en-us/library/windows/desktop/aa381049(v=vs.85).aspx

第九步：编辑完之后按F5编译并且保存，基本上就算完成了

 

 

补充： 想让go编译的程序在Windows点击运行不启动终端gui，可以在编译的时候加入如下参数

 

-ldflags "-H windowsgui"