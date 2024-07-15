// (c) Jisin0
//
// File containing methods to encode and decode batch links' data.

package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Jisin0/TGMessageStore/config"
)

// EncodeURL encodes the input data into a neat base64 encoded string to use as a start parameter.
func EncodeURL(chatID, startMsgID, endMsgID int64) string {
	// Structure of data -> copy_<chat_id>_<start_id>_<end_id>
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("copy_%d_%d_%d", chatID, startMsgID, endMsgID)))
}

// DecodeURL decodes the input data from a start query.
func DecodeURL(data string) (chatID, startMsgID, endMsgID int64, err error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Printf("utils.decodeurl: %v\n", err)
		return chatID, startMsgID, endMsgID, err
	}

	decodedString := string(decodedBytes)

	// Handle codex repo urls
	if strings.HasPrefix(decodedString, "get-") {
		return decodeCodexURL(decodedString)
	}

	if !strings.HasPrefix(decodedString, "copy_") {
		return chatID, startMsgID, endMsgID, errors.New("unknown start data format")
	}

	split := strings.Split(decodedString, "_")
	if len(split) < 4 {
		return chatID, startMsgID, endMsgID, errors.New("not enough input data")
	}

	var err1, err2, err3 error

	chatID, err1 = strconv.ParseInt(split[1], 10, 64)
	startMsgID, err2 = strconv.ParseInt(split[2], 10, 64)
	endMsgID, err3 = strconv.ParseInt(split[3], 10, 64)

	if err1 != nil || err2 != nil || err3 != nil {
		return chatID, startMsgID, endMsgID, errors.Join(err1, err2, err3)
	}

	return chatID, startMsgID, endMsgID, nil
}

// decodeCodexURL provides backward compatibility for urls generated with the CodeXBots/File-Sharing-Bot repo.
func decodeCodexURL(input string) (chatID, startMsgID, endMsgID int64, err error) {
	if config.DBChannel == 0 {
		return chatID, startMsgID, endMsgID, errors.New("DB_CHANNEL value not set")
	}

	split := strings.Split(input, "-")
	if len(split) < 3 {
		return chatID, startMsgID, endMsgID, errors.New("not enough data to parse")
	}

	sIDRaw, err1 := strconv.ParseInt(split[1], 10, 64)
	eIDRaw, err2 := strconv.ParseInt(split[2], 10, 64)

	if err1 != nil || err2 != nil {
		return chatID, startMsgID, endMsgID, errors.Join(err1, err2)
	}

	dbChannelAbs := int64(math.Abs(float64(config.DBChannel)))

	return config.DBChannel, sIDRaw / dbChannelAbs, eIDRaw / dbChannelAbs, nil

}
