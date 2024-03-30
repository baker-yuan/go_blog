module github.com/baker-yuan/go-blog/go_blog/retcode

go 1.19

replace (
    github.com/baker-yuan/go-blog/go_blog/common => ./../common
	github.com/baker-yuan/go-blog/go_blog/filters/httptrace => ./../filters/httptrace
)

require (
    github.com/baker-yuan/go-blog/go_blog/common v0.0.0-incompatible
    github.com/baker-yuan/go-blog/go_blog/filters/httptrace v0.0.0-incompatible
)