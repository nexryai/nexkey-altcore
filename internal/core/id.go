package core

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	math "math/rand"
	"strconv"
	"time"
)

var nexryaiArray = []rune{'n', 'e', 'x', 'r', 'y', 'a', 'i'}

const time2000 = 946684800000

var counter uint16

func getTime(t int64) string {
	t = t - time2000
	if t < 0 {
		t = 0
	}

	return fmt.Sprintf("%08s", strconv.FormatInt(t, 36))
}

func getNoise() string {
	counter++
	return fmt.Sprintf("%02s", strconv.FormatUint(uint64(counter), 36)[len(strconv.FormatUint(uint64(counter), 36))-2:])
}

func getRandomNexryai() string {
	math.Shuffle(len(nexryaiArray), func(i, j int) {
		nexryaiArray[i], nexryaiArray[j] = nexryaiArray[j], nexryaiArray[i]
	})
	return string(nexryaiArray)
}

// AID
// 長さ8の[2000年1月1日からの経過ミリ秒をbase36でエンコードしたもの] + 長さ2の[ノイズ文字列]
func genAid(date time.Time) string {
	t := date.UnixNano() / int64(time.Millisecond)
	counterBytes := make([]byte, 2)
	_, err := rand.Read(counterBytes)
	if err != nil {
		panic(system.ErrorOccurredWhileGeneratingID)
	}
	counter = binary.LittleEndian.Uint16(counterBytes)

	return getTime(t) + getNoise()
}

// nexryaid
// unixtime+整数ノイズ(3桁)+文字列ノイズ（nexryaiからランダムで並べ替え）
// ex) 1699364743467neryax
func genNexryaid(date time.Time) string {
	unixtime := date.Unix()
	randInt := math.Intn(900) + 100
	nexryai := getRandomNexryai()

	return fmt.Sprintf("%d%d%s", unixtime, randInt, nexryai)
}

func GenId() string {
	date := time.Now()

	if config.Id == "nexryaid" {
		return genNexryaid(date)
	} else {
		return genAid(date)
	}
}
