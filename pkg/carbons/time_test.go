package carbons

import (
	"github.com/golang-module/carbon/v2"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon(t *testing.T) {
	t.Run("NowFormat()", func(t *testing.T) {
		assert.Equal(t, NowFormat(), carbon.Now().String())
	})

	// fmt.Println(TimeToStr(time.Now(), YYYY_MM))
	// fmt.Println(TimeToStr(time.Now(), YYYY_MM_DD_HH_MM))

	// assert.Equal(t, MsToTime(), carbon)

	// t.Run("MsToTime()", func(t *testing.T) {
	// 	toTime, err := MsToTime(strconv.Itoa(carbon.Now().Millisecond()))
	// 	if err != nil {
	// 		return
	// 	}
	// 	t.Log(toTime.String())
	// })

	t.Run("Format()", func(t *testing.T) {
		t.Log(Format(time.Now(), "yyyy-MM-dd"))
		assert.Equal(t, Format(time.Now(), "yyyy-MM-dd"), carbon.Now().ToDateString())
		assert.Equal(t, Format(time.Now(), "yyyy-MM-dd HH:mm:ss"), carbon.Now().ToDateTimeString())
		// ...
	})

	t.Run("TimeToStr()", func(t *testing.T) {
		t.Log(TimeToStr(time.Now(), "yyyy-MM-dd"))
		assert.Equal(t, TimeToStr(time.Now(), "yyyy-MM-dd"), carbon.Now().ToDateString())
	})

	// t.Run("BeforeNowFormat()", func(t *testing.T) {
	// 	BeforeNowFormat("year")
	// })
}

func TestBeforeNowFormat(t *testing.T) {
	t.Run("SubYear()", func(t *testing.T) {
		// 测试年之前的时间
		yearTime := BeforeNowFormat("year", 1, "2006-01-02")
		expectedYear := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
		t.Log(yearTime)
		if yearTime != expectedYear {
			t.Errorf("Expected year time to be %s, but got %s", expectedYear, yearTime)
		}
		ct := carbon.Now().SubYear().ToDateString()
		if yearTime != ct {
			t.Errorf("Expected year time to be %s, but got %s", expectedYear, ct)
		}
	})

	t.Run("SubMonth()", func(t *testing.T) {
		// 测试月之前的时间
		monthTime := BeforeNowFormat("month", 1, "2006-01-02")
		expectedMonth := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
		t.Log(monthTime)
		if monthTime != expectedMonth {
			t.Errorf("Expected month time to be %s, but got %s", expectedMonth, monthTime)
		}
		ct := carbon.Now().SubMonth().ToDateString()
		if monthTime != ct {
			t.Errorf("Expected month time to be %s, but got %s", expectedMonth, ct)
		}
	})

	t.Run("SubMinute()", func(t *testing.T) {
		// 测试分钟之前的时间
		minuteTime := BeforeNowFormat("minute", 5, "2006-01-02 15:04:05")
		expectedMinute := time.Now().Add(-time.Minute * 5).Format("2006-01-02 15:04:05")
		t.Log(minuteTime)
		if minuteTime != expectedMinute {
			t.Errorf("Expected minute time to be %s, but got %s", expectedMinute, minuteTime)
		}
		ct := carbon.Now().SubMinutes(5).ToDateTimeString()
		if minuteTime != ct {
			t.Errorf("Expected minute time to be %s, but got %s", expectedMinute, ct)
		}
	})

	t.Run("SubDay()", func(t *testing.T) {
		// 测试日之前的时间
		dayTime := BeforeNowFormat("day", 1, "2006-01-02")
		expectedDay := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		t.Log(dayTime)
		if dayTime != expectedDay {
			t.Errorf("Expected day time to be %s, but got %s", expectedDay, dayTime)
		}
		ct := carbon.Now().SubDay().ToDateString()
		if dayTime != ct {
			t.Errorf("Expected day time to be %s, but got %s", expectedDay, ct)
		}
	})

	t.Run("", func(t *testing.T) {
		cts := carbon.Now().SubDay().Timestamp()
		t.Logf("subday ts: %d", cts)
		ct := carbon.Now().SubDay().ToDateTimeString()
		ctFromTs := carbon.CreateFromTimestamp(cts).ToDateTimeString()
		t.Log(ct)
		t.Log(ctFromTs)
		assert.Equal(t, ct, ctFromTs)
	})

	t.Run("GetDiffTime()", func(t *testing.T) {
		// GetDiffTime()
		t.Logf("???: %d", carbon.Now().AddMinutes(10).DiffInMinutes())
		t.Logf("Abs: %d", carbon.Now().AddMinutes(10).DiffAbsInMinutes())

		carbon.Now().AddMinutes(10).DiffInSeconds()
	})
}