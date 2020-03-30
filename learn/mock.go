package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)
/**
	一. 让我们将依赖关系定义为一个接口。
		这样我们就可以在 main 使用 真实的 Sleeper，并且在我们的测试中使用 spy sleeper。
		通过使用接口，我们的 Countdown 函数忽略了这一点，并为调用者增加了一些灵活性。
		type Sleeper interface {
			Sleep()
		}
	二. 我做了一个设计的决定，我们的 Countdown 函数将不会负责 sleep 的时间长度。
		这至少简化了我们的代码，也就是说，我们函数的使用者可以根据喜好配置休眠的时长。
		现在我们需要为我们使用的测试生成它的 mock。
		type SpySleeper struct {
			Calls int
		}

		func (s *SpySleeper) Sleep() {
			s.Calls++
		}
	三. *监视器（spies）*是一种 mock，它可以记录依赖关系是怎样被使用的。
		它们可以记录被传入来的参数，多少次等等。
		在我们的例子中，我们跟踪记录了 Sleep() 被调用了多少次，这样我们就可以在测试中检查它。
	四. 我们最新的修改只断言它已经 sleep 了 4 次，但是那些 sleeps 可能没按顺序发生
		如果你运行测试Countdown2，它们仍然应该通过，即使实现是错误的。
		让我们再用一种新的测试来检查操作的顺序是否正确。
	五. 我们的 CountdownOperationsSpy 同时实现了 io.writer 和 Sleeper，把每一次调用记录到 slice。
		在这个测试中，我们只关心操作的顺序，所以只需要记录操作的代名词组成的列表就足够了。
	六. 人们在这里看到的是测试驱动开发的弱点，但它实际上是一种力量，
		通常情况下，糟糕的测试代码是糟糕设计的结果，而设计良好的代码很容易测试。
	七. 你可能听过 mocking 是在作恶。就像软件开发中的任何东西一样，
		它可以被用来作恶，就像 DRY(Don't repeat yourself) 一样。
	八. 你正在进行的测试需要做太多的事情
			把模块分开就会减少测试内容
		它的依赖关系太细致
			考虑如何将这些依赖项合并到一个有意义的模块中
		你的测试过于关注实现细节
			最好测试预期的行为，而不是功能的实现
		通常，在你的代码中有大量的 mocking 指向 错误的抽象。
	九. 重构的定义是代码更改，但行为保持不变。 如果您已经决定在理论上进行一些重构，
		那么你应该能够在没有任何测试更改的情况下进行提交。所以，在写测试的时候问问自己。
			我是在测试我想要的行为还是实现细节？
		如果我要重构这段代码，我需要对测试做很多修改吗？
			虽然 Go 允许你测试私有函数，但我将避免它作为私有函数与实现有关。
		我觉得如果一个测试 超过 3 个模拟，那么它就是警告 —— 是时候重新考虑设计。
	十. 小心使用监视器。监视器让你看到你正在编写的算法的内部细节，这是非常有用的，
		但是这意味着你的测试代码和实现之间的耦合更紧密。如果你要监视这些细节，请确保你真的在乎这些细节。
	十一. 和往常一样，软件开发中的规则并不是真正的规则，也有例外。
		Uncle Bob 的文章 「When to mock」[2] 有一些很好的指南。
	十二. 更多关于测试驱动开发的方法
		当面对不太简单的例子，把问题分解成「简单的模块」。
		试着让你的工作软件尽快得到测试的支持，
		以避免掉进兔子洞（rabbit holes，意指未知的领域）和采取「最终测试（Big bang）」的方法。
		一旦你有一些正在工作的软件，小步迭代 应该是很容易的，直到你实现你所需要的软件。
	十三. 没有对代码中重要的区域进行 mock 将会导致难以测试。
		在我们的例子中，我们不能测试我们的代码在每个打印之间暂停，但是还有无数其他的例子。
		调用一个 可能 失败的服务？
		想要在一个特定的状态测试您的系统？
		在不使用 mocking 的情况下测试这些场景是非常困难的。
	十四. 如果没有 mock，你可能需要设置数据库和其他第三方的东西来测试简单的业务规则。
		你可能会进行缓慢的测试，从而导致 缓慢的反馈循环。
	十五. 当不得不启用一个数据库或者 webservice 去测试某个功能时，由于这种服务的不可靠性，
		你将会得到的是一个 脆弱的测试。
	十六. 一旦开发人员学会了 mocking，就很容易对系统的每一个方面进行过度测试，
		按照 它工作的方式 而不是 它做了什么。
		始终要注意 测试的价值，以及它们在将来的重构中会产生什么样的影响。
	十七. 在这篇关于 mocking 的文章中，我们只提到了 监视器（Spies），他们是一种 mock。
		也有不同类型的 mocks。Uncle Bob 的一篇极易阅读的文章中解释了这些类型[3]。
		在后面的章节中，我们将需要编写依赖于其他数据的代码，届时我们将展示 Stubs 行为。
 */
func Countdown(writer io.Writer, sleeper Sleeper){
	for i :=  3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Println(i)
	}
	sleeper.Sleep()
	fmt.Fprintf(writer,"Go!")
}



type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer,spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}