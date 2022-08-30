package dpt

import (
	"fmt"
	"time"
)

// DPT_19001 represents DPT 19.001 / Date Time.
type DPT_19001 struct {
	DPT_10001
	DPT_11001
}

func (d DPT_19001) Pack() []byte {
	return []byte{
		0,
		uint8(d.Year - 1900),
		d.Month,
		d.Day,
		d.Hour + d.Weekday<<5,
		d.Minutes,
		d.Seconds,
		0,
		0,
	}

}

func (d *DPT_19001) Unpack(data []byte) error {
	if len(data) != 9 {
		return ErrInvalidLength
	}

	d.Year = uint16(1900) + uint16(data[1])
	d.Month = uint8(data[2])
	d.Day = uint8(data[3])
	d.Weekday = uint8(data[4] >> 5)
	d.Hour = uint8(data[4] & 0x1f)
	d.Minutes = uint8(data[5])
	d.Seconds = uint8(data[6])

	return nil
}

func (d DPT_19001) IsValid() bool {
	date := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), int(d.Hour), int(d.Minutes), int(d.Seconds), 0, time.UTC)

	return (d.Year >= 1900 &&
		d.Year <= 2155 &&
		d.Year == uint16(date.Year()) &&
		d.Month == uint8(date.Month()) &&
		d.Day == uint8(date.Day()) &&
		d.Hour == uint8(date.Hour()) &&
		d.Minutes == uint8(date.Minute()) &&
		d.Seconds == uint8(date.Second()) &&
		d.isValidWeekday(uint8(date.Weekday())))
}

func (d DPT_19001) isValidWeekday(weekday uint8) bool {
	if d.Weekday == 7 {
		return weekday == 0
	}
	return d.Weekday == weekday
}

func (d DPT_19001) Unit() string {
	return ""
}

func (d DPT_19001) String() string {
	weekday := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	if d.Weekday >= 1 && d.Weekday <= 7 {
		return fmt.Sprintf("%s, %02d.%02d.%04d, %02d:%02d:%02d", weekday[d.Weekday-1], d.Day, d.Month, d.Year, d.Hour, d.Minutes, d.Seconds)
	} else {
		return fmt.Sprintf("%02d.%02d.%04d, %02d:%02d:%02d", d.Day, d.Month, d.Year, d.Hour, d.Minutes, d.Seconds)
	}
}
