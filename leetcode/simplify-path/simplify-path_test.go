package simplify_path

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "root path",
			args: args{path: "/"},
			want: "/",
		},
		{
			name: "valid file or folder",
			args: args{path: "/a/b/c.txt"},
			want: "/a/b/c.txt",
		},
		{
			name: "foldername with multiple letters",
			args: args{path: "/this/is/a/test"},
			want: "/this/is/a/test",
		},
		{
			name: "single directory",
			args: args{path: "/home/"},
			want: "/home",
		},
		{
			name: "trailing slash",
			args: args{path: "/home//"},
			want: "/home",
		},
		{
			name: "parent directory",
			args: args{path: "/ant/./bird/../.../chicken"},
			want: "/ant/.../chicken",
		},
		{
			name: "complex path",
			args: args{path: "/a/../../b/../c//.//"},
			want: "/c",
		},
		{
			name: "multiple slashes",
			args: args{path: "/a//b////c/d//././/.."},
			want: "/a/b/c",
		},
		{
			name: "current directory",
			args: args{path: "/./././."},
			want: "/",
		},
		{
			name: "empty path",
			args: args{path: ""},
			want: "/",
		},
		{
			name: "only slashes",
			args: args{path: "/////"},
			want: "/",
		},
		{
			name: "complex path with current and parent directories",
			args: args{path: "/a/./b/./c/./d/"},
			want: "/a/b/c/d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
