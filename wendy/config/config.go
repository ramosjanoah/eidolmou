package config

import (
	"fmt"
	"github.com/subosito/gotenv"
	"os"
	"strconv"
)

var (
	AppPort  string
	Platform string

	BotToken  string
	BotToggle bool

	CurrentDir string

	HeartbeatURL    string
	HeartbeatToggle bool

	AdminID int64
)

func init() {
	gotenv.Load()

	// int64
	AdminID = getIntConfig("WENDY_ADMIN_ID", 0, true)

	// string
	AppPort = getStringConfig("WENDY_PORT", "", false)
	if AppPort == "" {
		AppPort = getStringConfig("PORT", "", true)
	}
	BotToken = getStringConfig("WENDY_TOKEN", "", true)
	Platform = getStringConfig("WENDY_PLATFORM", "", true)

	// Toggle or boolean
	BotToggle = getBoolConfig("WENDY_BOT_TOGGLE", true, true)
	HeartbeatToggle = getBoolConfig("WENDY_HEARTBEAT_TOGGLE", true, true)

	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		panic("error when getting current directory")
	}
	CurrentDir = dir

}

func getStringConfig(envName string, defaultValue string, callPanic bool) string {
	temp := os.Getenv(envName)
	if temp == "" {
		if callPanic {
			panic(fmt.Sprintf("%s is empty!", envName))
		}

		temp = defaultValue
	}

	return temp
}

func getBoolConfig(envName string, defaultValue bool, callPanic bool) bool {
	defaultValueStr := "false"
	if defaultValue {
		defaultValueStr = "true"
	}

	temp := getStringConfig(envName, defaultValueStr, callPanic)

	if temp == "1" || temp == "true" || temp == "TRUE" {
		return true
	}
	return false
}

func getIntConfig(envName string, defaultValue int64, callPanic bool) int64 {
	defaultValueStr := strconv.Itoa(int(defaultValue))

	temp := getStringConfig(envName, defaultValueStr, callPanic)

	tempInt, err := strconv.Atoi(temp)
	if err != nil {
		panic(fmt.Sprintf("Error when parsing %s for integer value: %s", envName, temp))
	}

	return int64(tempInt)
}
