package TimeAcross

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"time"
)

func getICANTimeZones() ([]string, error) {
	var knownTimeZones []string

	csvReader := csv.NewReader(bytes.NewReader([]byte(zone1970)))
	csvReader.Comma = '\t'

	for {
		rec, err := csvReader.Read()
		if rec != nil && err != nil && strings.Contains(err.Error(), csv.ErrFieldCount.Error()) {
			// Expected
			// https://golang.org/pkg/encoding/csv/#Reader.Read
		} else if rec == nil && err == io.EOF {
			break
		}

		if len(rec) >= 3 {
			knownTimeZones = append(knownTimeZones, rec[2])
		} else {
			continue
		}
	}

	return knownTimeZones, nil
}

func timeZoneNameSearch(search string) ([]string, error) {
	knownTimeZones, err := getICANTimeZones()
	if err != nil {
		return knownTimeZones, fmt.Errorf("failed to retrieve ICAN timezones: %e", err)
	}

	var suggestions []string
	search = strings.ToLower(search)
	for i := range knownTimeZones {
		tzname := strings.ToLower(strings.ReplaceAll(knownTimeZones[i], "_", " "))
		if strings.Index(tzname, search) != -1 {
			suggestions = append(suggestions, knownTimeZones[i])
		}
	}

	return suggestions, nil
}

type TimeAcross struct {
	timezones []string
}

func (t *TimeAcross) ListTimezone() []string {
	return t.timezones
}

func (t *TimeAcross) AddTimezone(s string) error {
	tzNames, err := timeZoneNameSearch(s)
	if err != nil {
		return err
	}
	if len(tzNames) == 0 {
		return fmt.Errorf("timezone %s not found", s)
	}

	for i := range tzNames {
		t.timezones = append(t.timezones, tzNames[i])
	}

	return nil
}

func (t *TimeAcross) RemoveTimezone(s string) {
	for i := range t.timezones {
		if t.timezones[i] == s {
			t.timezones[i] = t.timezones[len(t.timezones)-1]
			t.timezones = t.timezones[:len(t.timezones)-1]
			return
		}
	}
}

func (t *TimeAcross) GetTimesNow() map[string]time.Time {
	return t.GetTimes(time.Now())
}

func (t *TimeAcross) GetTimes(baseTime time.Time) map[string]time.Time {
	var timeMap = make(map[string]time.Time, len(t.timezones))
	for _, tzName := range t.timezones {
		l, err := time.LoadLocation(tzName)
		if err != nil {
			fmt.Println(err)
		} else {
			timeMap[tzName] = baseTime.In(l)
		}
	}
	return timeMap
}
