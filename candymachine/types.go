package candymachine

import "google.golang.org/protobuf/proto"

var TopicTypes = map[string]proto.Message{
	"lotus.candymachine.0_0_0.nft_mintnft":                &MintNFT{},
	"lotus.candymachine.0_0_0.nft_updatecandymachine":     &UpdateCandyMachine{},
	"lotus.candymachine.0_0_0.nft_addconfiglines":         &AddConfigLines{},
	"lotus.candymachine.0_0_0.nft_initializecandymachine": &InitializeCandyMachine{},
	"lotus.candymachine.0_0_0.nft_updateauthority":        &UpdateAuthority{},
	"lotus.candymachine.0_0_0.nft_withdrawfunds":          &WithdrawFunds{},
}
