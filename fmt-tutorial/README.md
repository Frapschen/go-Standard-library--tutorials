# fmt
> fmt

Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived
from C's but are simpler.

## print.go
9个函数，
Printf -》 FPrintf
Print -》 Fprint
Println -》 Fprintln
这三个直接调用了另外三个，剩下6个函数：
FPrintf
Fprint
Fprintln
这三个都是格式化输出到一个 io.Writer,最后剩下3个：
Sprintf
Sprint
Sprintln
他们都返回一个 string 类型的结果
### 组成
#### 常量

各种有关错误的常量

#### type
* pp struct
* buffer []byte

### 基本逻辑：
```
step 1: get a struct of pp named p from newPrinter()
step 2: p.doPrintf(format,a)
step 3: get a s from string(p.buf)
step 4: sent to target
step 5: p.free()
```
整个 print.go 最关键的就是 p.doPrintf(format,a)/p.doPrint(a)/p.doPrintln(a)
 其中pp这个结构体是重点：
* 它存储 一个printer的状态
* 它可以复用，复用技术使用：sync.Pool
  * 其中newPrinter()与free()函数就是涉及将pp结构体从sync.Pool提取和释放
* 


### 相关文件
* format.go 在print.go中涉及p.buf时将将数据写入的buffer