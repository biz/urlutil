package urlutil

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURL(t *testing.T) {
	var cases = []struct {
		href string
	}{{
		href: "/foo",
	}}

	for _, c := range cases {
		u, err := url.Parse(c.href)
		if err != nil {
			t.Fatal(err)
		}

		Set(u, "bar", "baz")
		Add(u, "baz", "bust")
		Add(u, "baz", "bang")
		fmt.Println(u)
	}
}
