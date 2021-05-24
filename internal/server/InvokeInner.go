package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.io/MXuDong/example/config"
	"github.io/MXuDong/example/internal/model"
	"github.io/MXuDong/example/pkg/util/nets"
	"github.io/MXuDong/example/pkg/util/random"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
For invoke inner with application.

The inner invoker support some method to control the response of the http.
But now, not support the https(without any test or usage).

For invoke, it provide json params to control next invoke response. So any invoke
can use this to compose the complex system trace.

- @author: MXuDong
*/

func TracePost(ctx *gin.Context) {
	startTime := time.Now()
	jsonObjectList := []model.TracePostParam{}
	err := ctx.ShouldBindJSON(&jsonObjectList)
	res := ""
	if err != nil {
		// BindJson is already abort with error
		logrus.Warn(err)
		return
	}

	if len(jsonObjectList) == 0 {
		ctx.Status(204)
		return
	}
	objects := jsonObjectList[1:]

	// try pre sleep
	if jsonObjectList[0].BeforeSleep != "" {
		sleepTime, err := random.AnalyzePercentage(jsonObjectList[0].BeforeSleep)
		if err == nil {
			sleepDur, err := strconv.Atoi(sleepTime)
			if err == nil {
				time.Sleep(time.Millisecond * time.Duration(sleepDur))
			} else {
				err = nil
				logrus.Warn(err)
			}
		} else {
			// continue if err, but log it
			logrus.Warn(err)
			err = nil
		}
	}

	// parse code
	codeStr, err := random.AnalyzePercentage(jsonObjectList[0].TargetRespCode)
	if err != nil {
		logrus.Warn(err)
		err = nil
		codeStr = "200"
	}
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		logrus.Warn(err)
		err = nil
	}

	urlStr := jsonObjectList[0].NextUrl
	responseBody := ""
	if urlStr != "" {
		isRight, url := nets.GetHttpUrl(urlStr)
		if !isRight {
			logrus.Infof("Try to check url from %s to %s", urlStr, url)
		}
		body, err := json.Marshal(objects)
		if err == nil {
			resp, err := http.Post(url, "application/json", strings.NewReader(string(body)))
			if err != nil {
				logrus.Warn(err)
				err = nil
			}
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.Warn(err)
				err = nil
			} else {
				responseBody = string(body)
			}
		} else {
			logrus.Warn(err)
		}
	}

	// try after sleep
	if jsonObjectList[0].AfterSleep != "" {
		sleepTime, err := random.AnalyzePercentage(jsonObjectList[0].AfterSleep)
		if err == nil {
			sleepDur, err := strconv.Atoi(sleepTime)
			if err == nil {
				time.Sleep(time.Millisecond * time.Duration(sleepDur))
			} else {
				logrus.Warn(err)
				err = nil
			}
		} else {
			// continue if err, but log it
			logrus.Warn(err)
			err = nil
		}
	}

	if jsonObjectList[0].ResponseType == config.InnerInvokeResponseType_Copy {
		res = responseBody
	} else if jsonObjectList[0].ResponseType == config.InnerInvokeResponseType_Value {
		res = jsonObjectList[0].ResponseExpendField
	} else {
		resStruct := model.TraceResponseResp{
			ProcessStartTime: startTime.Unix(),
			ProcessStopTime:  time.Now().Unix(),
			RequestSettings:  jsonObjectList[0],
			Object:           responseBody,
			Value:            "nothing",
		}
		ctx.JSON(code, resStruct)
	}
	ctx.JSON(code, res)
}
