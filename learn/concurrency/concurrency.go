package concurrency

/**
	一. 它返回一个 map，由每个 url 检查后的得到的布尔值组成，成功响应的值为 true，错误响应的值为 false。
		你还必须传入一个 WebsiteChecker 处理单个 URL 并返回一个布尔值。它会被函数调用以检查所有的网站。
		使用 依赖注入[2]，允许在不发起真实 HTTP 请求的情况下测试函数，这使测试变得可靠和快速。
	二. 基准测试使用一百个网址的 slice 对 CheckWebsites 进行测试，并使用 WebsiteChecker 的伪造实现。
		slowStubWebsiteChecker 故意放慢速度。它使用 time.Sleep 明确等待 20 毫秒，然后返回 true。
		当我们运行基准测试时使用 go test -bench=. 命令
		(如果在 Windows Powershell 环境下使用 go test -bench=".")：
	三. 通常在 Go 中，当调用函数 doSomething() 时，我们等待它返回（即使它没有值返回，我们仍然等待它完成）。
		我们说这个操作是 阻塞 的 —— 它让我们等待它完成。
		Go 中不会阻塞的操作将在称为 goroutine 的单独 进程 中运行。
		将程序想象成从上到下读 Go 的 代码，当函数被调用执行读取操作时，进入每个函数「内部」。
		当一个单独的进程开始时，就像开启另一个 reader（阅读程序）在函数内部执行读取操作，
		原来的 reader 继续向下读取 Go 代码。
	四. 因为开启 goroutine 的唯一方法就是将 go 放在函数调用前面，所以当我们想要启动 goroutine 时，
		我们经常使用 匿名函数（anonymous functions）。
		一个匿名函数文字看起来和正常函数声明一样，但没有名字（意料之中）。
		可以在 for 循环体中看到一个。
	五. 匿名函数有许多有用的特性，其中两个上面正在使用。
		首先，它们可以在声明的同时执行 —— 这就是匿名函数末尾的 () 实现的。
		其次，它们维护对其所定义的词汇作用域的访问权 —— 在声明匿名函数时所有可用的变量也可在函数体内使用。
	六. 上面匿名函数的主体和之前循环体中的完全一样。
		唯一的区别是循环的每次迭代都会启动一个新的 goroutine，与当前进程（WebsiteChecker 函数）同时发生，
		每个循环都会将结果添加到 results map 中。
	七. 让我们困惑的是，原来的测试 WebsiteChecker 现在返回空的 map。哪里出问题了？
		我们 for 循环开始的 goroutines 没有足够的时间将结果添加结果到 results map 中；
		WebsiteChecker 函数对于它们来说太快了，以至于它返回时仍为空的 map。
	八. 这里的问题是变量 url 被重复用于 for 循环的每次迭代 —— 每次都会从 urls 获取新值。
		但是我们的每个 goroutine 都是 url 变量的引用 —— 它们没有自己的独立副本。
		所以他们 都 会写入在迭代结束时的 url —— 最后一个 url
	九. 过给每个匿名函数一个参数 url(u)，然后用 url 作为参数调用匿名函数，
		我们确保 u 的值固定为循环迭代的 url 值，重新启动 goroutine。u 是 url 值的副本，因此无法更改。
	十. fatal error: concurrent map writes。
		有时候，当我们运行我们的测试时，两个 goroutines 完全同时写入 results map。
		Go 的 Maps 不喜欢多个事物试图一次性写入，所以就导致了 fatal error。
	十一. 这是一种 race condition（竞争条件），
		当软件的输出取决于事件发生的时间和顺序时，因为我们无法控制，bug 就会出现。
		因为我们无法准确控制每个 goroutine 写入结果 map 的时间，两个 goroutines 同一时间写入时程序将非常脆弱。
		Go 可以帮助我们通过其内置的 race detector[3] 来发现竞争条件。
		要启用此功能，请使用 race 标志运行测试：go test -race。
	十二. 我们可以通过使用 channels 协调我们的 goroutines 来解决这个数据竞争。
		channels 是一个 Go 数据结构，可以同时接收和发送值。这些操作以及细节允许不同进程之间的通信。
		在这种情况下，我们想要考虑父进程和每个 goroutine 之间的通信，
		goroutine 使用 url 来执行 WebsiteChecker 函数。
	十三. 通过将结果发送到通道，我们可以控制每次写入 results map 的时间，确保每次写入一个结果。
		虽然 wc 的每个调用都发送给结果通道，但是它们在其自己的进程内并行发生，
		因为我们将结果通道中的值与接收表达式一起逐个处理一个结果。
	十四. 我们已经将想要加快速度的那部分代码并行化，同时确保不能并发的部分仍然是线性处理。
		我们使用 channel 在多个进程间通信。
	十五. 在使它更快的过程中，我们明白了
		goroutines 是 Go 的基本并发单元，它让我们可以同时检查多个网站。
		anonymous functions（匿名函数），我们用它来启动每个检查网站的并发进程。
		channels，用来组织和控制不同进程之间的交流，使我们能够避免 race condition（竞争条件） 的问题。
		the race detector（竞争探测器） 帮助我们调试并发代码的问题。
 */
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u,wc(u)}
		}(url)
	}

	for i :=0; i<len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}

