package size

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "small"},
	{0, "zero"},
	{500, "huge"},
	{99, "big"},
	{1000, "enormous"},
}

func TestSize(t *testing.T) {
	for i, test := range tests {
		size := Size(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}

/*
go test -coverprofile=coverage.out
opens coverage.out in your default browser
go tool cover -html=coverage.out
*/
/*
---------OutPut---------------
MahmoodAdil/go_course/15_testing
CMD: go test -cover
---------------------Result----------------
PASS
coverage: 42.9% of statements
ok      github.com/MahmoodAdil/go_course/15_testing     0.056s
*/
