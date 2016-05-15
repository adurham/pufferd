/*
 Copyright 2016 Padduck, LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 	http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package logging

import "fmt"

type level struct {
	scale byte;
	display string;
}

var (
	DEBUG level = level{scale: 7, display: "DEBUG"};
	INFO level = level{scale: 31, display: "INFO"};
	WARN level = level{scale: 63, display: "WARN"};
	ERROR level = level{scale: 127, display: "ERROR"};
	CRITICAL level = level{scale: 255, display: "CRITICAL"};
	loggingLevel = INFO;
)

func Info(msg string, data ...interface{}) {
	log(INFO, msg, data);
}

func Warn(msg string, data ...interface{}) {
	log(WARN, msg, data);
}

func Debug(msg string, data ...interface{}) {
	log(DEBUG, msg, data);
}

func Error(msg string, data ...interface{}) {
	log(ERROR, msg, data);
}

func Critical(msg string, data ...interface{}) {
	log(CRITICAL, msg, data);
}

func Infof(msg string, data ...interface{}) {
	logf(INFO, msg, data);
}

func Warnf(msg string, data ...interface{}) {
	logf(WARN, msg, data);
}

func Debugf(msg string, data ...interface{}) {
	logf(DEBUG, msg, data);
}

func Errorf(msg string, data ...interface{}) {
	logf(ERROR, msg, data);
}

func Criticalf(msg string, data ...interface{}) {
	logf(CRITICAL, msg, data);
}

func SetLevel(lvl level) {
	loggingLevel = lvl;
}

func log(lvl level, msg string, data ...interface{}) {
	if(lvl.scale < loggingLevel.scale) {
		return;
	}
	var dataLength = len(data[0].([]interface{}));
	if(data == nil || dataLength == 0) {
		fmt.Printf("[%s] %s\n", lvl.display, msg);
	} else {
		cast := make([]interface{}, 3);
		cast[0] = lvl.display;
		cast[1] = msg;
		if(dataLength == 1) {
			cast[2] = data[0].([]interface{})[0];
		} else {
			cast[2] = data[0].([]interface{});
		}
		fmt.Printf("[%s] %s\n%v\n", cast...);
	}
}

func logf(lvl level, msg string, data ...interface{}) {
	if(lvl.scale < loggingLevel.scale) {
		return;
	}
	if(data == nil || len(data[0].([]interface{})) == 0) {
		fmt.Printf("[%s] %s\n", lvl.display, msg);
	} else {
		fmt.Printf("[%s] %s\n", lvl.display, fmt.Sprintf(msg, data[0].([]interface{})...));
	}
}