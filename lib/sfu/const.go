package sfu

import "github.com/pion/ion-sfu/pkg/sfu"

var conf = sfu.Config{
	SFU: struct {
		Ballast   int64 `mapstructure:"ballast"`
		WithStats bool  `mapstructure:"withstats"`
	}{
		Ballast:   0,
		WithStats: false,
	},
	Router: sfu.RouterConfig{
		MaxBandwidth:        1500, // 最大傳遞buffer
		MaxPacketTrack:      500,  // 最大容納Track數量
		AudioLevelThreshold: 40,   // 音訊大小[0-127]
		AudioLevelInterval:  1000,
		AudioLevelFilter:    20,
		Simulcast: sfu.SimulcastConfig{
			BestQualityFirst:    true,
			EnableTemporalLayer: false,
		},
	},
	WebRTC: sfu.WebRTCConfig{
		ICEPortRange: []uint16{5000, 5200},
		SDPSemantics: "unified-plan",
		MDNS:         true,
		Timeouts: sfu.WebRTCTimeoutsConfig{
			ICEDisconnectedTimeout: 5,
			ICEFailedTimeout:       25,
			ICEKeepaliveInterval:   2,
		},
	},
	//Turn: sfu.TurnConfig{
	//Enabled: false,
	//Realm:   "ion",
	//Address: "0.0.0.0:3478",
	//Auth: sfu.TurnAuth{
	//Credentials: "pion=ion,pion2=ion2",
	//},
	//},
}
