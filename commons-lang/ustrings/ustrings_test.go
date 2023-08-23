package ustrings

import (
	"fmt"
	"testing"
)

func TestSubstringAfterLast(t *testing.T) {
	type args struct {
		str       string
		separator string
	}
	/**
	result1 := strutil.AfterLast("foo", "")
	    result2 := strutil.AfterLast("foo", "foo")
	    result3 := strutil.AfterLast("foo/bar", "/")
	    result4 := strutil.AfterLast("foo/bar/baz", "/")
	    result5 := strutil.AfterLast("foo/bar/foo/baz", "foo")
	*/
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "foo", separator: ""},
			want: "foo",
		},
		{
			args: args{str: "foo", separator: "foo"},
			want: "",
		},
		{
			args: args{str: "foo/bar/baz", separator: "/"},
			want: "baz",
		},
		{
			args: args{str: "foo/bar/foo/baz", separator: "foo"},
			want: "/baz",
		},
		{
			args: args{str: "foo", separator: "he"},
			want: "foo",
		},
	}
	for i, tt := range tests {
		tests[i].name = fmt.Sprintf("SubstringAfterLast(\"%s\", \"%s\")=[%s]", tt.args.str, tt.args.separator, tt.want)
		t.Run(tt.name, func(t *testing.T) {
			if got := SubstringAfterLast(tt.args.str, tt.args.separator); got != tt.want {
				t.Errorf("SubstringAfterLast() = %v, want %v", got, tt.want)
			}
			//if got := strutil.AfterLast(tt.args.str, tt.args.separator); got != tt.want {
			//	t.Errorf("strutil.AfterLast() = %v, want %v", got, tt.want)
			//}
		})
	}
}
