package concurrency

import (
	"strconv"
	"testing"
	"time"
)

//func slowStubWebsiteChecker(u string) bool {
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	//println(u)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url " + strconv.Itoa(i)
	}

	//println("number:"+strconv.Itoa(b.N))

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}

}
