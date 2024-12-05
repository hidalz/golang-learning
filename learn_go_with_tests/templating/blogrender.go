package blogrenderer

import (
	"github.com/hidalz/learn_go_with_tests/blogposts"
	"github.com/hidalz/learn_go_with_tests/blogrenderer"
	"fmt"
	"io"
)

func Render(w io.Writer, p blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
