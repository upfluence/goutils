package statsd

import "testing"

type testCase struct {
	in  string
	out string
}

func TestStatsdfy(t *testing.T) {
	for _, tCase := range []testCase{
		testCase{"*foo@bar", "foo_bar"},
		testCase{"bar123", "bar123"},
		testCase{"bar.123", "bar.123"},
	} {
		if v := statsdfyName(tCase.in); tCase.out != v {
			t.Errorf("Wrong statsdfy of %s: %s", tCase.in, v)
		}
	}
}
