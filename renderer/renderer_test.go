package renderer

import (
	"testing"
)

func TestRender(t *testing.T) {
	const html = `
	<!DOCTYPE html>
	<html>
		<head>
			<title></title>
		</head>
		<body>
			<div class="hello world" id="links">
				<a href="https://mohsal.dev">mohsal.dev</a>
			</div>
		</body>
	</html>
	`

	Render(html)
}
