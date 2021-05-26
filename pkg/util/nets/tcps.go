package nets

import (
	"context"
	"fmt"
	"github.io/MXuDong/example/pkg/constant"
	"golang.org/x/sync/semaphore"
	"net"
	"strings"
	"time"
)

//Tcp util, for build a tcp controller or create tcp connect.
type TcpServer struct {
	isClosed         bool
	listener         net.Listener
	handlerSemaphore *semaphore.Weighted
	handlers         func(conn net.Conn) error
}

// Close will close listener, and set nil to ts.listener
func (ts *TcpServer) Close() error {
	if ts.isClosed {
		return nil
	}
	err := ts.listener.Close()
	if err != nil {
		return err
	}
	ts.listener = nil
	ts.isClosed = true
	return err
}

// RegisterHandler will reset handler.
func (ts *TcpServer) RegisterHandler(f func(conn net.Conn) error) {
	ts.handlers = f
}

// Try to listen the target address with target netWork.
// If ts already init, it will try to close old listener, and generator new listener for
// tcp in target address with target netWork
func (ts *TcpServer) Listen(netWork, address string, max_handler_count int64) error {
	// if listener not close, try close it.
	if ts.listener != nil && !ts.isClosed {
		err := ts.Close()
		if err != nil {
			return err
		}
	}
	l, err := net.Listen(netWork, address)
	if err != nil {
		return err
	}
	ts.handlerSemaphore = semaphore.NewWeighted(max_handler_count)
	ts.listener = l
	ts.isClosed = false
	return nil
}

// Do will accept the value from tcp controller, and input need a errHandler to resolve
// err in handler
func (ts *TcpServer) Do(errHandler func(err error)) error {
	for {
		conn, err := ts.listener.Accept()
		if err != nil {
			return err
		}

		err = ts.handlerSemaphore.Acquire(context.Background(), 1)
		if err != nil {
			return err
		}
		go func() {
			err := ts.handlers(conn)
			if err != nil {
				errHandler(err)
				err = nil // reset err
			}

			err = conn.Close()
			if err != nil {
				errHandler(err)
				err = nil // reset err
			}

			ts.handlerSemaphore.Release(1)
		}()
	}
}

// DefaultTcpServer will generator a tcp listener in packing struct: TcpServer
func DefaultTcpServer(address string) (*TcpServer, error) {
	return NetTcpServer(constant.TcpProtocol, address)

}

// NetTcpServer will Generator a tcp listener in packing struct: TcpServer
// Diff from DefaultTcpServer, it can set netWork in tcp, tcp4, tcp6, unix, unixpacket
func NetTcpServer(netWork, address string) (*TcpServer, error) {
	ts := TcpServer{}
	var err = ts.Listen(netWork, address, constant.DefaultHandlerMaxCount)
	return &ts, err
}

func DefaultTcpHandler(conn net.Conn) error {
	readBuf := make([]byte, 1024)
	for {
		_, err := conn.Read(readBuf)
		if err != nil {
			return err
		}
		if strings.Index(string(readBuf), "stop") == 1 {
			// stop connect
			conn.Write([]byte(fmt.Sprintf("From %v to %v : %v", conn.LocalAddr(), conn.RemoteAddr(), time.Now())))
			break
		} else {
			conn.Write(readBuf)
		}
	}
	return nil
}
