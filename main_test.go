package TimeAcross

import (
	"reflect"
	"testing"
	"time"
)


func TestTimeAcross_AddTimezone(t1 *testing.T) {
	ta := TimeAcross{}
	_ = ta.AddTimezone("Kiev")

	got := ta.ListTimezone()
	want := []string{"Europe/Kiev"}
	if !reflect.DeepEqual(got, want) {
		t1.Errorf("ListTimezone() = %v, want %v", got, want)
	}
}


func TestTimeAcross_RemoveTimezone(t1 *testing.T) {
	ta := TimeAcross{}
	_ = ta.AddTimezone("Kiev")
	_ = ta.AddTimezone("New Y")
	ta.RemoveTimezone("Europe/Kiev")

	got := ta.ListTimezone()
	want := []string{"America/New_York"}
	if !reflect.DeepEqual(got, want) {
		t1.Errorf("ListTimezone() = %v, want %v", got, want)
	}
}

func TestTimeAcross_ListTimezone(t1 *testing.T) {
	ta := TimeAcross{}
	_ = ta.AddTimezone("Kiev")
	_ = ta.AddTimezone("New Y")
	_ = ta.AddTimezone("Chic")

	got := ta.ListTimezone()
	want := []string{"Europe/Kiev", "America/New_York", "America/Chicago"}
	if !reflect.DeepEqual(got, want) {
		t1.Errorf("ListTimezone() = %v, want %v", got, want)
	}
}

func TestTimeAcross_GetTimes(t1 *testing.T) {
	ta := TimeAcross{}
	_ = ta.AddTimezone("Kiev")
	_ = ta.AddTimezone("New Y")
	_ = ta.AddTimezone("Chic")

	const form = "2006-01-02 15:04:05 -0700 MST"

	chicago, _ := time.Parse(form, "1999-12-31 17:59:59 -0600 CST")
	ny, _ := time.Parse(form, "1999-12-31 18:59:59 -0500 EST")
	kyiv, _ := time.Parse(form, "2000-01-01 01:59:59 +0200 EET")

	want := map[string]time.Time{
		"America/Chicago": chicago,
		"America/New_York": ny,
		"Europe/Kiev": kyiv,
	}

	t, _ := time.Parse(form, "1999-12-31 23:59:59 -0000 UTC")
	got := ta.GetTimes(t)
	for _, tz := range []string{"America/Chicago", "Europe/Kiev", "America/New_York"} {
		if got[tz].UnixNano() != want[tz].UnixNano() {
			t1.Errorf("GetTimes() = %v, want %v", got, want)
		}
	}
}
