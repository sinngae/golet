# Go test
执行 go test 命令，它会在 *_test.go 中寻找 test 测试、benchmark 基准 和 examples 示例 函数。测试函数必须以 TestXXX 的函数名出现（XXX 为以非小写字母开头），基准函数必须以 BenchmarkXXX 的函数名出现，示例函数必须以 ExampleXXX 的形式。三种函数类似下面的签名形式
```shell
go test -h # 查看帮助

# 总语法
go test [build/test flags] [packages] [build/test flags & test binary flags]

```
## test code
+ 单元测试，TestXXX(t *testing.T)
+ 基准测试，BenchmarkXXX(b *testing.B)
+ Main测试，TestMain(m *testing.M)
+ 控制台测试，ExampleXXX()
go test 命令还会忽略 testdata 目录，该目录用来保存测试需要用到的辅助数据。

go test 有两种运行模式：
+ 本地目录模式，在没有包参数（例如 go test 或 go test -v）调用时发生。在此模式下，go test 编译当前目录中找到的包和测试，然后运行测试二进制文件。在这种模式下，caching 是禁用的。在包测试完成后，go test 打印一个概要行，显示测试状态、包名和运行时间。
+ 包列表模式，在使用显示包参数调用 go test 时发生（例如 go test math，go test ./... 甚至是 go test .）。在此模式下，go 测试编译并测试在命令上列出的每个包。如果一个包测试通过，go test 只打印最终的 ok 总结行。如果一个包测试失败，go test 将输出完整的测试输出。如果使用 -bench 或 -v 标志，则 go test 会输出完整的输出，甚至是通过包测试，以显示所请求的基准测试结果或详细日志记录。

下面详细说明下 go test 的具体用法，flag 的作用及一些相关例子。需要说明的是：一些 flag 支持 go test 命令和编译后的二进制测试文件。它们都能识别加 -test. 前缀的 flag，如 go test -test.v，但编译后的二进制文件必须加前缀 ./sum.test -test.bench=.。

test flag
以下 flag 可以跟被 go test 命令使用：

-args：传递命令行参数，该标志会将 -args 之后的参数作为命令行参数传递，最好作为最后一个标志。
$ go test -args -p=true
-c：编译测试二进制文件为 [pkg].test，不运行测试。
$ go test -c && ./sum.test -p=true
-exec xprog：使用 xprog 运行测试，行为同 go run 一样，查看 go help run。
-i：安装与测试相关的包，不运行测试。
$ go test -i
-o file：编译测试二进制文件并指定文件，同时运行测试。
go test -o filename

test/binary flag
以下标志同时支持测试二进制文件和 go test 命令。

-bench regexp：通过正则表达式执行基准测试，默认不执行基准测试。可以使用 -bench .或-bench=.执行所有基准测试。
$ go test -bench=.
$ go test -c
$ ./sum.test -test.bench=.
-benchtime t：每个基准测试运行足够迭代消耗的时间，time.Duration（如 -benchtime 1h30s），默认 1s。
$ go test -bench=. -benchtime 0.1s
$ ./sum.test -test.bench=. -test.benchtime=1s
-count n：运行每个测试和基准测试的次数（默认 1），如果 -cpu 指定了，则每个 GOMAXPROCS 值执行 n 次，Examples 总是运行一次。
$ go test -bench=. -count=2
$ ./sum.test -test.bench=. -test.count=2
-cover：开启覆盖分析，开启覆盖分析可能会在编译或测试失败时，代码行数不对。
$ go test -bench=. -cover
-covermode set,count,atomic：覆盖分析的模式，默认是 set，如果设置 -race，将会变为 atomic。
set，bool，这个语句运行吗？
count，int，该语句运行多少次？
atomic，int，数量，在多线程正确使用，但是耗资源的。
-coverpkg pkg1,pkg2,pkg3：指定分析哪个包，默认值只分析被测试的包，包为导入的路径。
# sum -> $GOPATH/src/test/sum
$ go test -coverpkg test/sum
-cpu 1,2,4：指定测试或基准测试的 GOMAXPROCS 值。默认为 GOMAXPROCS 的当前值。
-list regexp：列出与正则表达式匹配的测试、基准测试或 Examples。只列出顶级测试（不列出子测试），不运行测试。
$ go test -list Sum
-parallel n：允许并行执行通过调用 t.Parallel 的测试函数的最大次数。默认值为 GOMAXPROCS 的值。-parallel 仅适用于单个二进制测试文件，但go test命令可以通过指定 -p 并行测试不同的包。查看 go help build。
$ go test -run=TestSumParallel -parallel=2
-run regexp：只运行与正则表达式匹配的测试和Examples。我们可以通过 / 来指定测试子函数。go test Foo/A=，会先去匹配并执行 Foo 函数，再查找子函数。
$ go test -v -run TestSumSubTest/1+
-short：缩短长时间运行的测试的测试时间。默认关闭。
$ go test -short
-timeout d：如果二进制测试文件执行时间过长，panic。默认10分钟（10m）。
$ go test -run TestSumLongTime -timeout 1s
-v：详细输出，运行期间所有测试的日志。
$ go test -v
analyze flag
以下测试适用于 go test 和测试二进制文件：

-benchmem：打印用于基准的内存分配统计数据。
$ go test -bench=. -benchmem
$ ./sum.test -test.bench -test.benchmem
-blockprofile block.out：当所有的测试都完成时，在指定的文件中写入一个 goroutine 阻塞概要文件。指定 -c，将写入测试二进制文件。
$ go test -v -cpuprofile=prof.out
$ go tool pprof prof.out
-blockprofilerate n：goroutine 阻塞时候打点的纳秒数。默认不设置就相当于 -test.blockprofilerate=1，每一纳秒都打点记录一下。
-coverprofile cover.out：在所有测试通过后，将覆盖概要文件写到文件中。设置过 -cover。
-cpuprofile cpu.out：在退出之前，将一个 CPU 概要文件写入指定的文件。
-memprofile mem.out：在所有测试通过后，将内存概要文件写到文件中。
-memprofilerate n：开启更精确的内存配置。如果为 1，将会记录所有内存分配到 profile。
$ go test -memprofile mem.out -memprofilerate 1
$ go tool pprof mem.out
-mutexprofile mutex.out：当所有的测试都完成时，在指定的文件中写入一个互斥锁争用概要文件。指定 -c，将写入测试二进制文件。
-mutexprofilefraction n：样本 1 在 n 个堆栈中，goroutines 持有 a，争用互斥锁。
-outputdir directory：在指定的目录中放置输出文件，默认情况下，go test 正在运行的目录。
-trace trace.out：在退出之前，将执行跟踪写入指定文件。


代码覆盖率
由单元测试的代码，触发运行到的被测试代码的代码行数占所有代码行数的比例，被称为测试覆盖率，代码覆盖率不一定完全精准，但是可以作为参考，可以帮我们测量和我们预计的覆盖率之间的差距

go test -cover 生成代码测试覆盖率 ，coverage: 8.1% of statements

go test -v -coverprofile=c.out 将生成的代码测试覆盖率放入一个文件，然后运行下面的命令可以将 c.out 文件转换成一个 html 文件用浏览器阅读，截图如下，no coverage 代表没有覆盖的代码，high coverage 代表高覆盖率的代表，一个红色，一个绿色，这里红色的截图上没体现出来，大家可本地试验一下。
