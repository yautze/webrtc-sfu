package router

import (
	"webrtc-sfu/controller/room"
	"webrtc-sfu/middle"

	"github.com/kataras/iris/v12"
)

func setRoom(e iris.Party) {
	r := e.Party("/room")
	{
		// join
		r.Get("/join", middle.HandleFunc(room.JoinRoom))
	}
}
