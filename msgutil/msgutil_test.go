package msgutil_test

import (
	"strings"
	"testing"

	"github.com/capybara-alt/msgutil/msgutil"
)

func TestGetMessage1(t *testing.T) {
	msgutil.Init("./sample1/sample.json")
	msgStr, err := msgutil.GetMassage("sample", "errors", "validate", "null")
	if err != nil {
		t.Fail()
	}

	if msgStr != "column.a cannot be null" {
		t.Fail()
	}
}

func TestGetMessage2(t *testing.T) {
	msgutil.Init("./sample1/sample.json")
	msgStr, err := msgutil.GetMassage("sample", "info", "login", "success")
	if err != nil {
		t.Fail()
	}

	if msgStr != "success!!" {
		t.Fail()
	}
}

func TestGetMessageMultiFile(t *testing.T) {
	msgutil.Init("./sample2/*.json")
	msgStr, err := msgutil.GetMassage("en", "errors", "login", "fail")
	if err != nil {
		t.Fail()
	}

	if msgStr != "fail" {
		t.Fail()
	}

	msgStr, err = msgutil.GetMassage("ja", "errors", "login", "fail")
	if err != nil {
		t.Fail()
	}

	if msgStr != "失敗" {
		t.Fail()
	}
}

func TestGetMessageArgs(t *testing.T) {
	msgutil.Init("./sample1/sample.json")
	msgStr, err := msgutil.GetMessageArgs([]string{"sample", "errors", "validate", "invalid-value"}, "A", strings.Join([]string{"A","B","C"}, " or "))
	if err != nil {
		t.Fail()
	}

	if msgStr != "Column A must be A or B or C" {
		t.Fail()
	}
}

func TestGetMessageArgsMultiFile(t *testing.T) {
	msgutil.Init("./sample2/*.json")
	msgStr, err := msgutil.GetMessageArgs([]string{"en", "errors", "validate", "invalid-value"}, "A", strings.Join([]string{"A","B","C"}, " or "))
	if err != nil {
		t.Fail()
	}

	if msgStr != "Column A must be A or B or C" {
		t.Fail()
	}

	msgStr, err = msgutil.GetMessageArgs([]string{"ja", "errors", "validate", "invalid-value"}, "A", strings.Join([]string{"A","B","C"}, " もしくは "))
	if err != nil {
		t.Fail()
	}

	if msgStr != "項目 A に入力できる値は A もしくは B もしくは C のいずれかです" {
		t.Fail()
	}
}

func TestGetMessageArgsEmptyCase(t *testing.T) {
	msgutil.Init("./sample1/sample.json")
	msgStr, err := msgutil.GetMessageArgs([]string{"sample", "errors", "validate", "invalid-value"})
	if err != nil {
		t.Fail()
	}

	if msgStr != "Column {{0}} must be {{1}}" {
		t.Fail()
	}
}

func TestNotFoundError1(t *testing.T) {
	msgutil.Init("./sample1/sample.json")
	msgStr, err := msgutil.GetMassage("aa")
	if msgStr != "" {
		t.Fail()
	}

	if err.Error() != "Target message not found" {
		t.Fail()
	}
}

func TestNotFoundError2(t *testing.T) {
	msgutil.Init("./sample1/sample.json")
	msgStr, err := msgutil.GetMassage("commongen", "aa")
	if msgStr != "" {
		t.Fail()
	}

	if err.Error() != "Target message not found" {
		t.Fail()
	}
}

func TestArgNullError(t *testing.T) {
	msgutil.Init("./sample1/sample.json")
	msgStr, err := msgutil.GetMassage()
	if msgStr != "" {
		t.Fail()
	}

	if err.Error() != "Keys cannot be empty" {
		t.Fail()
	}
}

func TestFileNotFound(t *testing.T) {
	err := msgutil.Init("./aaa/*.json")
	if err.Error() != "Message resource file not found" {
		t.Fail()
	}
}
