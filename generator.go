package generator

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

var (
	reNumber   = regexp.MustCompile("[^0-9]")
	reAlphabet = regexp.MustCompile("[^a-zA-Z]")
)

func GetCharLastSixIndex(gcn string) string {
	last6 := gcn[len(gcn)-6:]
	return last6
}

func ParseTimeZone(t time.Time) time.Time {
	loc := "Asia/Jakarta"
	tz := time.FixedZone(loc, 7*60*60)

	utc := t.UTC()
	jstTime := utc.In(tz)

	return jstTime
}

func FormatDate(t time.Time) string {
	dateFormat := "060102050415" // YYMMDDssmmhh
	return t.Format(dateFormat)
}

func Generator(gcn string, trxTypeID string) (string, error) {
	var transactionID string
	today := ParseTimeZone(time.Now())
	currentDate := FormatDate(today)

	if gcn == "" {
		return "", errors.New("empty gcn")
	}

	if trxTypeID == "" {
		return "", errors.New("empty transaction type")
	}

	if gcn == "" && trxTypeID == "" {
		return "", errors.New("empty transaction type and gcn")
	}

	if reNumber.MatchString(trxTypeID) {
		return "", errors.New("transaction id should be numeric")
	}

	transactionID = fmt.Sprintf("%s%s%s", currentDate, GetCharLastSixIndex(gcn), trxTypeID)

	return transactionID, nil
}

func main() {
	gcn := "DEV40CAF6045CC50920EF35A1E697D58"
	typeID := "02"
	fmt.Println(GetCharLastSixIndex("DEV40CAF6045CC50920EF35A1E697D58"))
	fmt.Println(!reAlphabet.MatchString(typeID))
	fmt.Println(Generator(gcn, typeID))

	getData, _ := Generator(gcn, typeID)
	fmt.Println(getData)

}
