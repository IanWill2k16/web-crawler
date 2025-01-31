package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "subdomains and varying paths",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
<head>
    <title>Example 2</title>
</head>
<body>
    <a href="https://help.example1.com">Example 1</a><br>
    <a href="http://example2.com/testing/things/with/all/the/things">Example 2</a><br>
    <a href="https://bluuuuerrrrgh.aaarghhh.example3.com">Example 3</a>
</body>
</html>
		`,
			expected: []string{"https://help.example1.com", "http://example2.com/testing/things/with/all/the/things", "https://bluuuuerrrrgh.aaarghhh.example3.com"},
		},
		{
			name:     "multitudes of urls",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
<head>
    <title>Example 3</title>
</head>
<body>
    <a href="https://example1.com">Example 1</a><br>
    <a href="https://example2.com">Example 2</a><br>
    <a href="https://example3.com">Example 3</a><br>
    <a href="https://example4.com">Example 4</a><br>
    <a href="https://example5.com">Example 5</a>
</body>
</html>
		`,
			expected: []string{"https://example1.com", "https://example2.com", "https://example3.com", "https://example4.com", "https://example5.com"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputURL, tc.inputBody)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
