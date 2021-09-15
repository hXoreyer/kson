# kson a JSON parsing library

### 引言:

> go语言自带的json解析库，用结构体解析还好说，但是比较麻烦，如果用map[string]interface{}来解析的话就会很麻烦
>
> 前端给你的json值类型稍微有点和你想象的不一样就是个panic
>
> 比如有个字段类型是int，但是前端没有把input里的字符串转化为int直接通过json传给你，你的断言就会出问题



```go
value := int(mp[key].(float64))
//这样断言的话，前端给你传输的是string就会panic
```

只能修改为

```go
val,_ := strconv.ParseInt(mp[key].(string),10,64)
```



我这里在原库的基础上加入了异常，并加以处理来多种方法实现断言
解析过程可以为链式操作

### 例子:

json文件:

```json
{
    "datas": 1,
    "datas2":"2",
    "array":[
        "hello",
        "kson",
        {
            "db": true,
            "db2": "false"
        }
    ]
}
```

go文件:

```go
package main

import (
	"fmt"
	"github.com/hxoreyer/kson"
)

func main() {
	obj, _ := kson.KparseByFile("./test.json")
    fmt.Println(obj.GetInt("datas"),obj.GetInt("datas2"))
    //输出: 1 2
    fmt.Println(obj.GetArray("array").GetString(0), obj.GetArray("array").GetString(1))
    //输出: "hello" "kson"
    fmt.Println(obj.GetArray("array").GetObject(2).GetBool("db"), obj.GetArray("array").GetObject(2).GetBool("db2"))
    //输出: true false
    a := obj.GetArray("array")
    for(i := 0; i < a.Length(); i++){
        fmt.Println(a.Get(i))
    }
    //输出: "hello" "kson" map[db:true db2:false]
}

```

