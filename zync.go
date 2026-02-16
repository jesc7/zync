package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/codec/openh264"
	_ "github.com/pion/mediadevices/pkg/driver/screen"
	"github.com/pion/webrtc/v4"
)

func Encode(v any) (string, error) {
	b, e := json.Marshal(v)
	if e != nil {
		return "", e
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func Decode(in string, v any) error {
	b, e := base64.StdEncoding.DecodeString(in)
	if e != nil {
		return e
	}
	return json.Unmarshal(b, v)
}

type Msg struct {
	Type  int    `json:"type"`
	Code  int    `json:"code"`
	Error string `json:"error,omitzero"`
	Key   string `json:"key,omitzero"`
	Value string `json:"val,omitzero"`
}

func createOffer() (pc *webrtc.PeerConnection, offer webrtc.SessionDescription, e error) {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{URLs: []string{"stun:stun.l.google.com:19302"}},
		},
	}

	xParams, _ := openh264.NewParams() //Create a new RTCPeerConnection
	xParams.UsageType = openh264.ScreenContentRealTime
	xParams.BitRate = 1_000_000 // 500kbps

	codecSelector := mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(&xParams),
	)

	mediaEngine := webrtc.MediaEngine{}
	codecSelector.Populate(&mediaEngine)
	api := webrtc.NewAPI(webrtc.WithMediaEngine(&mediaEngine))
	if pc, e = api.NewPeerConnection(config); e != nil {
		return
	}

	s, e := mediadevices.GetDisplayMedia(mediadevices.MediaStreamConstraints{
		Video: func(c *mediadevices.MediaTrackConstraints) {
		},
		Codec: codecSelector,
	})
	if e != nil {
		return
	}

	for _, track := range s.GetTracks() {
		if _, e = pc.AddTransceiverFromTrack(track,
			webrtc.RTPTransceiverInit{
				Direction: webrtc.RTPTransceiverDirectionSendonly,
			},
		); e != nil {
			return
		}
	}

	if offer, e = pc.CreateOffer(nil); e != nil {
		return
	}

	gatherComplete := webrtc.GatheringCompletePromise(pc)
	if e = pc.SetLocalDescription(offer); e != nil {
		return
	}
	<-gatherComplete
	return
}

func zync() {
	//signaling server websocket conn
	u := url.URL{Scheme: "ws", Host: "localhost:1212", Path: "/ws"}
	conn, _, e := websocket.DefaultDialer.Dial(u.String(), nil)
	if e != nil {
		log.Fatal("dial:", e)
	}
	defer conn.Close()

	pc, offer, e := createOffer()
	if e != nil {
		panic(e)
	}
	payload, e := Encode(offer)
	if e != nil {
		panic(e)
	}
	conn.WriteJSON(Msg{
		Type:  0,
		Value: payload,
	})

	var (
		key, pwd string
	)
	for {
		var msg Msg
		if e := conn.ReadJSON(&msg); e != nil {
			log.Printf("Read message error: %v", e)
			break
		}
		if msg.Code < 0 {
			panic(errors.New(msg.Error))
		}
		switch msg.Type {
		case 0:
			sl := strings.Split(msg.Key, "@")
			if msg.Type != 0 || len(sl) < 2 {
				panic(errors.New("Wrong key"))
			}
			key, pwd = sl[0], sl[1]
			fmt.Printf("Key=%s, pwd=%s\n", key, pwd)

		case 3:
			var answer webrtc.SessionDescription
			if e = Decode(msg.Value, answer); e != nil {
				panic(e)
			}
			if e = pc.SetRemoteDescription(answer); e != nil {
				panic(e)
			}
			// Block forever
			select {}
		}
	}
}
