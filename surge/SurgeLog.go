package surge

import (
	"os"

	"github.com/rule110-io/surge-ui/surge/platform"
	log "github.com/sirupsen/logrus"
)

const logPath = "surge.log"

//InitializeLog init for log file
func InitializeLog() {

	var err error

	var dir = platform.GetSurgeDir()

	var logPathOS = dir + string(os.PathSeparator) + logPath

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	file, err := os.OpenFile(logPathOS, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
}

//OpenLogFile opens a log file with OS default application for object type
func OpenLogFile() {
	var err error

	var dir = platform.GetSurgeDir()

	var logPathOS = dir + string(os.PathSeparator) + logPath

	if err != nil {
		pushError("Error on open log", err.Error())
		return
	}

	OpenOSPath(logPathOS)
}
