package helper

import (
	"time"
)

const (
	MYSQL_DATE_FORMAT 		= "2006-01-02"
	MYSQL_TIME_FORMAT 		= "15:04:05"
	MYSQL_DATETIME_FORMAT 	= MYSQL_DATE_FORMAT +" "+ MYSQL_TIME_FORMAT
	ISO8601_TIME_FORMAT		= "2006-01-02T15:04:03.000-0700"
)

// JakartaTime is default location of application with timezone = Asia/Jakarta (UTC+7)
var JakartaTime = time.FixedZone("Asia/Jakarta", 7 * 60 * 60)