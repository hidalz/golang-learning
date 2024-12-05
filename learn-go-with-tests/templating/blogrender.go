package blogrenderer

import (
	"fmt"
	"io"

	"github.com/hidalz/golang_learning/learn_go_with_tests/blogposts"
)

func Render(w io.Writer, p blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
