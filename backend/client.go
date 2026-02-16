package backend

import (
	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/codec/openh264"
	"github.com/pion/webrtc/v4"
)

var pc *webrtc.PeerConnection

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
