package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.io/MXuDong/example/internal/model"
	"github.io/MXuDong/example/pkg/constant"
	"net"
	"time"
)

// The protocol-mock handler will mock the protocol request, like TCP, UDP and etc.
// The protocol-mock handler is the client, it will create request to target protocol controller.
// And for mock, the response always set to now time.
// Target protocol controller will set from input params.
// The trace of request:
// Common Client -- http -> Example Application's protocol handler -- target protocol -> target protocol controller.
//
// All of protocol-mock handler will response now time.

func MockTcpRequest(ctx *gin.Context) {

	tcpMockParam := model.TcpMockParam{}

	if err := ctx.ShouldBindJSON(&tcpMockParam); err != nil {
		ctx.Error(err)
		return
	}

	// if connect time less than 0, set it to 0
	connectTime := tcpMockParam.ConnectTime
	if connectTime < 0 {
		connectTime = 0
	}

	if tcpMockParam.SendByteCount < 0 || tcpMockParam.SendByteCount > 1023 {
		// tcp send byte count error
		ctx.AbortWithStatus(400)
		return
	}

	if tcpMockParam.Protocol == "" {
		tcpMockParam.Protocol = constant.TcpProtocol
	}

	client, err := net.DialTimeout(tcpMockParam.Protocol, tcpMockParam.RemoteIpAddress+":"+tcpMockParam.RemoteIpPort, 2*time.Second)
	if err != nil {
		ctx.Error(err)
		return
	}
	defer client.Close()

	bs := make([]byte, tcpMockParam.SendByteCount)
	readTotoalByte := []byte{}
	for i := 0; i < tcpMockParam.SendByteCount; i++ {
		bs[i] = 1
	}

	tk := time.NewTicker(1 * time.Second)
	startTime := time.Now().Local()
	for {
		_ = <-tk.C // time block
		connectTime--
		if connectTime < 0 {

			break
		}

		buff := make([]byte, 1024)

		// send value
		_, err := client.Write(bs)
		if err != nil {
			logrus.Errorf("Tcp write error : %v", err)
			err = nil
		}
		for {
			// read value
			rc, err := client.Read(buff)
			readTotoalByte = append(readTotoalByte, buff[:rc]...)
			if err != nil {
				logrus.Errorf("Tcp read error : %v", err)
				err = nil
			}
			if rc < 1024 {
				// if great than 1024, it mean still has value.
				break
			}
		}
	}

	// close
	_, err = client.Write([]byte("stop"))
	if err != nil {
		ctx.Error(err)
		return
	}
	err = client.Close()
	if err != nil {
		ctx.Error(err)
		return
	}

	returns := model.MockReturn{
		InvokeTime:      startTime,
		EndTime:         time.Now().Local(),
		ValueByte:       readTotoalByte,
		RequestProtocol: tcpMockParam.Protocol,
		ValueSize:       uint64(len(readTotoalByte)),
		InvokingSetting: tcpMockParam,
	}

	ctx.JSON(200, returns)
}
