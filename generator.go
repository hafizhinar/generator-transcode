package generator

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/hafizhinar/generator-transcode/utils"
)

var (
	reNumber   = regexp.MustCompile("[^0-9]")
	reAlphabet = regexp.MustCompile("[^a-zA-Z]")
)

func GetCharLastSixIndex(gcn string) string {
	last6 := gcn[len(gcn)-6:]
	return last6
}

func GeneratorBiller(gcn string, trxTypeID string) (string, error) {
	var transactionID string
	today := utils.ParseTimeZone(time.Now())
	currentDate := utils.FormatDate(today)

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
