### flag包 ###


##### 涉及知识点 #####
- Linux命令风格
- 变量逃逸
    - 当一个变量在某个作用域中声明,但被其他作用域所引用时,就发生了逃逸
    - 只针对指针。值变量如果没有被取址, 那么它永远不可能逃逸
    - 验证 
        - `go tool compile -S main.go | grep runtime.newobject`
        - `go run -gcflags "-m -l" main.go` (-m打印逃逸分析信息，-l禁止内联编译)
    - 影响
        - 无法在栈上分配, 需要分配到堆上
        - 堆内存分配增加gc压力
        - 失去编译器相关优化, 降低效率
    - 解决方法
        - 避免返回局部变量地址
        - 使用基本类型替代引用类型和指针
        - 减少变量作用域, 控制在真正需要的范围
        - 复制副本代替直接引用
    - 其他
        - 被指针类型(slice、map、chan)引用的指针 一定发生逃逸。这就是为什么使用指针的chan比使用值的chan慢30%, gc拖慢了速度

##### 概念 #####
1. 官方提供的命令行解析库。
2. 同类型的还有pflag第三方库,支持Unix/POSIX风格的命令解析。
    1. 常见的Linux命令风格主流有三种 Unix/Posix、BSD、GNU
    2. Unix/Posix 也叫短选项风格 选项以连接字符 `-` 开头的单个字符, 注意是一个字符不是一个单词.
        1. 选项后面如果不带参数, 称之为模式选择, 模式选择可以组合在一起使用。例如 `-a` 和 `-l` 可以合并 `-al`
        2. 选项后面需要带参数, 称之为参数选项, 参数要紧接在选项后面, 通常用空格隔开
    3. BSD 与Unix/Posix相比, 其选项使用单个字符, 且不带任何前缀。如 `ps a`
        1. 如果是多个不带参数的选项, 也可以组合在一起 如 `ps aux`
        2. 如果选项需要带参数, 也同Unix/Posix风格一样
    4. GNU 也叫长选项风格 选项使用两个连字符 `--` 开头的单词 如 `ls --all`
        1. 如果选项需要带参数, 则使用空格或者`=`将参数和选项分开 如 `ls --sort time` 或者 `ls --sort=time`
    5. 除了上面三种风格的命令行, 还有小部分命令行独自有风格。 比如`java -verison` java风格
3. cobra、cli 是两个更加复杂、全功能的框架。

##### 案例 #####

1. 目标
```cmd
NAME:
   gvg - go version management by go

USAGE:
   gvg [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   list       list go versions
   install    install a go version
   info       show go version info
   use        select a version
   uninstall  uninstall a go version
   get        get the latest code
   uninstall  uninstall a go version
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
2. 分析
这个命令不仅包含了全局的选项，还有 8 个子命令，部分子命令支持参数和选项。
3. 实现
    1. 没有全局参数 或者说 全局参数是子命令
    2. 全局选项有 `--help -h` `--version -v`
        1. 一个选项在flag包中用一个Flag表示 `h := flag.Bool("h", false, "show help")`
        2. 还有另一种方式 `var v bool; flag.BoolVar(&v, "v", false, "print the version")`
        3. 第二种方式使用较多 或许因为第一种可能有变量逃逸
    3. 长短选项
        1. 一个flag应该有 `-v --version` 长短两种形式
        2. 可以使用曲线救国的方式(定义两个flag, 绑定到同一个变量)
        ```
        flag.BoolVar(&v, "v", false, "print the version")
        flag.BoolVar(&v, "version", false, "print the version")
        ```
    4. 命令行解析
        1. 定义好所有flag, 还需要进一步解析才能拿到正确的结果。这一步需要 `flag.Parse()` 即可
    5. 命令行使用
        1. 前面介绍的选项用法有三种 `-flag` `-flag=x` `-flag x(非bool类型才支持这种方式)`
    6. 扩展目标
        1. 上面的子命令 `list` 支持获取Go的版本列表。但是版本来源有很多, 比如installed、local、remote
        2. 这时候 `list` 支持一个flag选项 `--origin` 用于指定版本信息来源
            1. 如果要求不严格, 可以使用stringVar实现, 但是使用string, 即使输入不在有效范围也能成功解析, 不够严谨。虽然在获取之后可以检查, 但还是不够灵活和可配置, 所以我们想实现一个新的flag选项类型, 选项的值必须在范围内, 否则给出一定的错误信息
        3. 基于6.2.1 我们计划实现一个新类型, 而实现上可以参考 `flag.DurationVar` 实现
            1. 本质是实现 `flag.Value` 接口
            2. 有几个变量需要: 存放解析结果的指针、解析命令行输入的Value、表示一个选项的flag
            3. 








现在实现文章开头要求的目标。新类型定义如下：type stringEnumValue struct {
	options []string
	p   *string
}
名为 StringEnumValue，即字符串枚举。它有 options 和 p 两个成员，options 指定一定范围的值，p 是 string 指针，保存解析结果的变量的地址。下面定义创建 StringEnumValue 变量的函数 newStringEnumValue，代码如下：func newStringEnumValue(val string, p *string, options []string) *StringEnumValue {
	*option = val
	return &stringEnumValue{options: options, p: p}
}
除了 val 和 p 两个必要的输入外，还有一个 string 切片类型的数，名为 options，它用于范围的限定。而函数主体，首先设置默认值，然后使用 options 和 p 创建变量返回。Set 是核心方法，解析命令行传入字符串。代码如下：func (s *StringEnumValue) Set(v string) error {
	for _, option := range s.options {
		if v == option {
			*(s.p) = v
			return nil
		}
	}

	return fmt.Errorf("must be one of %v", s.options)
}
循环检查输入参数 v 是否满足要求。定义如下：最后是 String() 方法，func (s *StringEnumValue) String() string {
	return *(s.p)
}
返回 p 指针中的值。前面分析实现思路时，Flag 在设置默认值时就调用了它。