package _test

import "testing"

// 单元测试 go test -v
func TestMult(t *testing.T) {
	if ans := mult(1, 2); ans != 2 {
		t.Error("mult(1, 2) should be equal to 3")
	}

	/**
	$ go test -v
	=== RUN   TestMult
	--- PASS: TestMult (0.00s)
	PASS
	ok      gateway 0.006s

	*/
}


// 性能测试 go test -bench="."
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		mult(40, 50)
	}

	/**
	$ go test -bench="."

	goos: darwin
	goarch: amd64
	pkg: gateway
	BenchmarkGetArea-4      1000000000               0.296 ns/op
	PASS
	ok      gateway 0.335s

	*/
}

// 覆盖率 测试 go test -cover
/**
覆盖率测试能知道测试程序总共覆盖了多少业务代码（也就是 demo_test.go 中测试了多少 demo.go 中的代码），可以的话最好是覆盖100%。
$ go test -cover
PASS
coverage: 100.0% of statements
ok      gateway 0.006s

 */