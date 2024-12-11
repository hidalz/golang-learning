module blogrenderer

go 1.23.4

replace blogposts => ../blogposts

require (
	blogposts v0.0.0-00010101000000-000000000000
	github.com/approvals/go-approval-tests v0.0.0-20241205182534-1cf6af14f2bb
	github.com/gomarkdown/markdown v0.0.0-20241205020045-f7e15b2f3e62
)
