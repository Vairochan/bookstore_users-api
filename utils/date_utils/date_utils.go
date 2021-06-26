package date_utils

import "time"

const(
	apiDateLayout = "02-01-2006T15:04:05Z"
	apiDbLayout = "2006-01-02 15:04:05"
)

//TODO: For using the stamdrad time zone.
func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString()string{
	//return GetNow().Format(apiDateLayout)
	now := time.Now()
	return now.Format(apiDbLayout)

}
func GetNowDBFormat()string {
	return GetNow().Format(apiDbLayout)
}