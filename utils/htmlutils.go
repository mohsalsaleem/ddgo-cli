package utils

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

// HTMLToString - Conver to HTML to String
func HTMLToString(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
