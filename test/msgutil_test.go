package msgutil_test

import (
	"strings"
	"testing"

	"github.com/capybara-alt/msgutil"
)

func TestGetMessage1(t *testing.T) {
	msgutil.Init("./sample.json")
	msgStr, err := msgutil.GetMassage("errors", "validate", "null")
	if err != nil {
		t.Fail()
	}

	if msgStr != "column.a cannot be null" {
		t.Fail()
	}
}

func TestGetMessage2(t *testing.T) {
	msgutil.Init("./sample.json")
	msgStr, err := msgutil.GetMassage("info", "login", "success")
	if err != nil {
		t.Fail()
	}

	if msgStr != "success!!" {
		t.Fail()
	}
}

func TestGetMessageArgs(t *testing.T) {
	msgutil.Init("./sample.json")
	msgStr, err := msgutil.GetMessageArgs([]string{"errors", "validate", "invalid-value"}, "A", strings.Join([]string{"A","B","C"}, " or "))
	if err != nil {
		t.Fail()
	}

	if msgStr != "Column A must be A or B or C" {
		t.Fail()
	}
}

func TestGetMessageArgsEmptyCase(t *testing.T) {
	msgutil.Init("./sample.json")
	msgStr, err := msgutil.GetMessageArgs([]string{"errors", "validate", "invalid-value"})
	if err != nil {
		t.Fail()
	}

	if msgStr != "Column {{0}} must be {{1}}" {
		t.Fail()
	}
}

func TestNotFoundError1(t *testing.T) {
	msgutil.Init("./sample.json")
	msgStr, err := msgutil.GetMassage("aa")
	if msgStr != "" {
		t.Fail()
	}

	if err.Error() != "Target error not found" {
		t.Fail()
	}
}

func TestNotFoundError2(t *testing.T) {
	msgutil.Init("./sample.json")
	msgStr, err := msgutil.GetMassage("commongen", "aa")
	if msgStr != "" {
		t.Fail()
	}

	if err.Error() != "Target error not found" {
		t.Fail()
	}
}

func TestArgNullError(t *testing.T) {
	msgutil.Init("./sample.json")
	msgStr, err := msgutil.GetMassage()
	if msgStr != "" {
		t.Fail()
	}

	if err.Error() != "Keys cannot be empty" {
		t.Fail()
	}
}
