// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"errors"

	"github.com/q191201771/naza/pkg/bele"
)

// -----------------------------------
// rfc3550 5.1 RTP Fixed Header Fields
// -----------------------------------
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |V=2|P|X|  CC   |M|     PT      |       sequence number         |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                           timestamp                           |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           synchronization source (SSRC) identifier            |
// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
// |            contributing source (CSRC) identifiers             |
// |                             ....                              |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

var ErrRTP = errors.New("lal.rtp: fxxk")

const (
	RTPFixedHeaderLength = 12
)

const (
	RTPPacketTypeAVC = 96
	RTPPacketTypeAAC = 97
)

// rfc3984 5.2.  Common Structure of the RTP Payload Format
// Table 1.  Summary of NAL unit types and their payload structures
//
// Type   Packet    Type name                        Section
// ---------------------------------------------------------
// 0      undefined                                    -
// 1-23   NAL unit  Single NAL unit packet per H.264   5.6
// 24     STAP-A    Single-time aggregation packet     5.7.1
// 25     STAP-B    Single-time aggregation packet     5.7.1
// 26     MTAP16    Multi-time aggregation packet      5.7.2
// 27     MTAP24    Multi-time aggregation packet      5.7.2
// 28     FU-A      Fragmentation unit                 5.8
// 29     FU-B      Fragmentation unit                 5.8
// 30-31  undefined                                    -

const (
	NALUTypeSingleMax = 23
	NALUTypeFUA       = 28
)

type RTPHeader struct {
	version    uint8  // 2b
	padding    uint8  // 1b
	extension  uint8  // 1
	csrcCount  uint8  // 4b
	mark       uint8  // 1b
	packetType uint8  // 7b
	seq        uint16 // 16b
	timestamp  uint32 // 32b
	ssrc       uint32 // 32b

	payloadOffset uint32
}

func isAudio(packetType uint8) bool {
	if packetType == RTPPacketTypeAAC {
		return true
	}
	return false
}

func parseRTPPacket(b []byte) (h RTPHeader, err error) {
	if len(b) < RTPFixedHeaderLength {
		err = ErrRTP
		return
	}

	h.version = b[0] >> 6
	h.padding = (b[0] >> 5) & 0x1
	h.extension = (b[0] >> 4) & 0x1
	h.csrcCount = b[0] & 0xF
	h.mark = b[1] >> 7
	h.packetType = b[1] & 0x7F
	h.seq = bele.BEUint16(b[2:])
	h.timestamp = bele.BEUint32(b[4:])
	h.ssrc = bele.BEUint32(b[8:])

	h.payloadOffset = RTPFixedHeaderLength
	return
}
