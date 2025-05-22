# domainFilter
### 介绍
主要功能是把输入的所有域名，全部用 nslookup 检测它们对应的 IP 地址，然后对应好关系之后，根据 IP 地址，来去重域名，从而获取所有不重复的域名，防止在 HW 的时候因为重复域名而导致的重复扫描。

建议使用IP来进行扫描，事后可以通过关系对照表（.csv文件）来确认每个IP所对应的资产。

```bash
必选参数:
   -i, -input string        输入的 domain，可以是多个，但需要以 , 分隔，如：www.baidu.com,www.bing.com
   -if, -input-file string  输入的文件名

可选参数:
   -o, -output string  输出的文件名，可为空 (default "domainFilter_output_20250501_120646.txt")
   -od, -out-domain    输出txt内容改为domain，原输出为IP
   -jp, -json-print    在终端打印执行结果的json串
   -nl, -no-logo       禁止输出logo等相关信息
   -ne, -no-err        禁止输出探测不到的域名信息
   -nt, -no-txt         禁止输出txt文件
   -nc, -no-csv        禁止输出csv文件。（默认会输出csv关系对照表）
   -v, -view           -v 打印所有的扫描信息
     🔔温馨提示：终端打印出来的 [+] IP -> Doamin 属于扫描信息实时打印，会有重复，去重后的结果在输出的文件中，使用去重后的数据请在输出的文件中获取

定制参数:
   -st, -server-tools  开启后，将使用ServerTools模式
```

### 演示
![示例 GIF](img/1.gif)

### 版本更新
#### v1.1  
1. 修复了遇到一个域名对应多个IP时输出重复域名的情况。