package main

import "testing"

func TestKfuxSplitNewLine(t *testing.T) {
	cases := []struct {
		txt   string
		lines int
	}{
		{txt: "foo\nbar", lines: 2},
	}

	for _, c := range cases {
		res := splitNewLine(c.txt)
		if len(res) != c.lines {
			t.Errorf("got=%d;want=%d", len(res), c.lines)
		}
	}
}

func TestKfuxSplitNewLines(t *testing.T) {
	cases := []struct {
		txt   string
		lines int
	}{
		{txt: "foo\nbar", lines: 2},
		{txt: "foo\n\nbar", lines: 2},
		{txt: "foo\n\n\nbar", lines: 2},
		{txt: "foo\n\n\nbar\n", lines: 2},
		{txt: "\nfoo\n\n\nbar\n", lines: 2},
		{txt: "\nfoo\n\n\n\nbar\n", lines: 2},
	}

	for _, c := range cases {
		res := splitNewLines(c.txt)
		if len(res) != c.lines {
			t.Errorf("got=%d;want=%d", len(res), c.lines)
		}
	}
}

func TestKfuxRemovePrefix(t *testing.T) {
	cases := []struct {
		txt      string
		expected string
	}{
		{txt: "\"foo\nbar", expected: "foo\nbar"},
		{txt: "foo\nbar", expected: "foo\nbar"},
		{txt: "\"\"foo\nbar", expected: "foo\nbar"},
	}

	for _, c := range cases {
		res := removePrefix(c.txt, `"`)
		if res != c.expected {
			t.Errorf("got=%s;want=%s", res, c.expected)
		}
	}
}

func TestKfuxRemoveSuffix(t *testing.T) {
	cases := []struct {
		txt      string
		expected string
	}{
		{txt: "\"foo\nbar\"", expected: "\"foo\nbar"},
		{txt: "foo\nbar\"\"", expected: "foo\nbar"},
		{txt: "foo\nbar\"\"\"", expected: "foo\nbar"},
		{txt: "\"\"foo\nbar", expected: "\"\"foo\nbar"},
	}

	for _, c := range cases {
		res := removeSuffix(c.txt, `"`)
		if res != c.expected {
			t.Errorf("got=%s;want=%s", res, c.expected)
		}
	}
}
