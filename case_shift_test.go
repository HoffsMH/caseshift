package caseshift

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPascalCase(t *testing.T) {
	Convey("from kebab case", t, func() {
		got := PascalCase("hello-kebab-case")

		// doesn't get the rule or the value
		So(got, ShouldEqual, "HelloKebabCase")

		got = PascalCase("hello-kebab2-case-2")

		// doesn't get the rule or the value
		So(got, ShouldEqual, "HelloKebab2Case2")
	})
}

func TestGetSubWords(t *testing.T) {
	Convey("from kebab case", t, func() {
		got := GetSubWords("hello-kebab-case")

		// doesn't get the rule or the value
		So(
			got,
			ShouldResemble,
			[]string{
				"hello",
				"kebab",
				"case",
			},
		)

		got = GetSubWords("hello-kebab2-case-2")

		// doesn't get the rule or the value
		So(
			got,
			ShouldResemble,
			[]string{
				"hello",
				"kebab2",
				"case",
				"2",
			},
		)
	})
	Convey("from snake case", t, func() {
		got := GetSubWords("hello_kebab_case")

		// doesn't get the rule or the value
		So(
			got,
			ShouldResemble,
			[]string{
				"hello",
				"kebab",
				"case",
			},
		)

		got = GetSubWords("hello_kebab2_case_2")

		// doesn't get the rule or the value
		So(
			got,
			ShouldResemble,
			[]string{
				"hello",
				"kebab2",
				"case",
				"2",
			},
		)
	})
}
