package room

import (
	"encoding/json"
	"webrtc-sfu/infra/log"
	"webrtc-sfu/lib/sfu"
	"webrtc-sfu/lib/ws"
	"webrtc-sfu/middle"

	"github.com/pion/webrtc/v3"
)

// JoinRoom -
func JoinRoom(c *middle.C) {
	// new ws connection
	w, err := ws.NewWebSocket(c.ResponseWriter(), c.Request())
	if err != nil {
		c.E(err)
		return
	}

	// new sfu peer
	peer := sfu.Dial()

	// --- ws event ---
	// Join -
	w.On("join", func(e *ws.Event) {
		join := new(sfu.Join)
		bb, _ := json.Marshal(e.Data)
		json.Unmarshal(bb, &join)
		join.UID = e.ID

		peer.OnOffer = func(offer *webrtc.SessionDescription) {
			m := map[string]interface{}{
				"method": "offer",
				"params": offer,
			}
			data, _ := json.Marshal(m)
			w.Out <- data
		}

		peer.OnIceCandidate = func(candidate *webrtc.ICECandidateInit, target int) {
			m := map[string]interface{}{
				"method": "trickle",
				"params": &sfu.Trickle{
					Candidate: *candidate,
					Target:    target,
				},
			}
			b, _ := json.Marshal(m)
			w.Out <- b
		}

		err := peer.Join(join.SID, join.UID, join.Config)
		if err != nil {
			log.Logger.Errorf("join err : %v", err)
		}

		answer, err := peer.Answer(join.Offer)
		if err != nil {
			log.Logger.Errorf("join reply answer err : %v", err)
		}

		m := map[string]interface{}{
			"id":     e.ID,
			"result": answer,
		}

		res, _ := json.Marshal(m)
		w.Out <- res
	})

	// Offer -
	w.On("offer", func(e *ws.Event) {
		negotiation := new(sfu.Negotiation)
		b, _ := json.Marshal(e.Data)
		json.Unmarshal(b, &negotiation)

		answer, err := peer.Answer(negotiation.Desc)
		if err != nil {
			log.Logger.Errorf("offer reply answer err : %v", err)
		}

		m := map[string]interface{}{
			"id":     e.ID,
			"result": answer,
		}
		r, _ := json.Marshal(m)
		w.Out <- r
	})

	// Answer -
	w.On("answer", func(e *ws.Event) {
		negotiation := new(sfu.Negotiation)
		b, _ := json.Marshal(e.Data)
		json.Unmarshal(b, &negotiation)

		err := peer.SetRemoteDescription(negotiation.Desc)
		if err != nil {
			log.Logger.Errorf("answer err : %v", err)
		}
	})

	// Trickle -
	w.On("trickle", func(e *ws.Event) {
		trickle := new(sfu.Trickle)
		b, _ := json.Marshal(e.Data)
		json.Unmarshal(b, &trickle)

		err := peer.Trickle(trickle.Candidate, trickle.Target)
		if err != nil {
			log.Logger.Errorf("trickle err : %v", err)
		}
	})

	<-w.Close

	peer.Close()

	return
}
