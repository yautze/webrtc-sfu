package sfu

import (
	"github.com/pion/ion-sfu/pkg/middlewares/datachannel"
	"github.com/pion/ion-sfu/pkg/sfu"
)

var s *sfu.SFU

func init() {
	s = sfu.NewSFU(conf)
	dc := s.NewDatachannel(sfu.APIChannelLabel)
	dc.Use(datachannel.SubscriberAPI)
}

// Dial -
func Dial() *sfu.PeerLocal {
	return sfu.NewPeer(s)
}
