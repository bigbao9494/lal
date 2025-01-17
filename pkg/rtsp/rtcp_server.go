// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import "github.com/q191201771/naza/pkg/nazalog"

type RTCPServer struct {
	udpServer *UDPServer
}

func NewRTCPServer(addr string) *RTCPServer {
	var s RTCPServer
	s.udpServer = NewUDPServer(addr, s.OnReadUDPPacket)
	return &s
}

func (r *RTCPServer) OnReadUDPPacket(b []byte, addr string, err error) {
	//nazalog.Debugf("< R length=%d, remote=%s, err=%v", len(b), addr, err)
	parseRTCPPacket(b)
}

func (r *RTCPServer) Listen() (err error) {
	nazalog.Infof("start rtcp server listen. addr=%s", r.udpServer.addr)
	return r.udpServer.Listen()
}

func (r *RTCPServer) RunLoop() error {
	return r.udpServer.RunLoop()
}
