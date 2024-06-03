package cyber

import (
	"github.com/trylotus/connectors/cyber/TransparentUpgradeableProxy"
	"google.golang.org/protobuf/proto"
)

const (
	//                                         0x84583e7D2D92D87D5B3bAC850aB4bAd37aE568e8
	TransparentUpgradeableProxyContractAddr = "0x84583e7d2d92d87d5b3bac850ab4bad37ae568e8"
)

var TopicTypes = []proto.Message{
	&TransparentUpgradeableProxy.AdminChanged{},
	&TransparentUpgradeableProxy.Upgraded{},
}
