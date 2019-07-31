/**
* @Author: xhzhang
* @Date: 2019-04-16 14:16
 */
package tcp

import (
	"fmt"
	"github.com/glory-cd/utils/log"
	"io"
	"net"
)

type DataFunc func(string)

func TCPStart(addr string, dealfun DataFunc) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Slogger.Errorf(fmt.Sprintf("tcp listen fail.[%s]\n", err))
		return
	} else {
		log.Slogger.Debug("tcp server is listening...")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Slogger.Errorf(fmt.Sprintf("tcp accept fail.[%s]\n", err))
			continue
		}
		go handleConn(conn, dealfun)

	}

}

func handleConn(conn net.Conn, deal DataFunc) {
	defer conn.Close()
	for {
		var buf = make([]byte, 10240)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Slogger.Error(fmt.Sprintf("read form connect failed.[%s]", err))
		}

		go deal(string(buf[:n]))
	}

}
