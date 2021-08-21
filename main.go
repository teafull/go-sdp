package main

import (
	"fmt"
	"github.com/pixelbender/go-sdp/sdp"
)

func main() {
	DecodingTest()

	EncodingSdp()
}

func DecodingTest() {
	sess, err := sdp.ParseString(`v=0
o=alice 2890844526 2890844526 IN IP4 alice.example.org
s=Example
c=IN IP4 127.0.0.1
t=0 0
a=sendrecv
m=audio 10000 RTP/AVP 0 8
a=rtpmap:0 PCMU/8000
a=rtpmap:8 PCMA/8000`)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sess.Media[0].Format[0].Name) // prints PCMU
	}
}

func EncodingSdp() {
	sess := &sdp.Session{
		Origin: &sdp.Origin{
			Username:       "alice",
			Address:        "alice.example.org",
			SessionID:      2890844526,
			SessionVersion: 2890844526,
		},
		Name: "Example",
		Connection: &sdp.Connection{
			Address: "127.0.0.1",
		},
		Media: []*sdp.Media{
			{
				Type:  "audio",
				Port:  10000,
				Proto: "RTP/AVP",
				Format: []*sdp.Format{
					{Payload: 0, Name: "PCMU", ClockRate: 8000},
					{Payload: 8, Name: "PCMA", ClockRate: 8000},
				},
			},
		},
		Mode: sdp.SendRecv,
	}

	sdpStr := sess.String()
	fmt.Println(sdpStr)
}
