package infra

import (
	"webrtc-sfu/infra/config"
	"webrtc-sfu/infra/log"
)

func Init() {
	log.Logger = log.New(config.C.Log.Level)
}
