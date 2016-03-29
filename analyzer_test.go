package bullyanalyzer

import (
	_ "fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstructor(t *testing.T) {

	Convey("Testing containsEntry()", t, func() {

		a, err := New("./profanity_dutch.txt")

		So(err, ShouldEqual, nil)
		So(a.ContainsEntry("peenhoofd"), ShouldEqual, true)
		So(a.ContainsEntry("some random sentence"), ShouldNotEqual, true)

	})

}

func TestContainsEntry(t *testing.T) {

	Convey("Testing ContainsEntry()", t, func() {

		a := Analyzer{
			Entries: []string{
				"hello",
				"world",
			},
		}

		So(a.ContainsEntry("hello"), ShouldEqual, true)

	})

}

func TestAnalyzePost(t *testing.T) {

	Convey("Testing AnalyzePost()", t, func() {

		a := Analyzer{
			Entries: []string{
				"fuck",
				"ugly",
			},
		}

		So(a.AnalyzePost("she is not ugly").Value, ShouldEqual, 0.25)
		So(a.AnalyzePost("what?").Value, ShouldEqual, 0)

	})

}
