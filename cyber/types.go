package cyber

import (
	"github.com/trylotus/connectors/cyber/xo"
	"google.golang.org/protobuf/proto"
)

const (
	XOContractAddr = "0x84583e7d2d92d87d5b3bac850ab4bad37ae568e8"
)

var TopicTypes = []proto.Message{
	&xo.PaidDM{},
	&xo.RoleAdminChanged{},
	&xo.SBTUpdated{},
	&xo.NewGoodVibes{},
	&xo.Post{},
	&xo.RoleGranted{},
	&xo.Initialized{},
	&xo.NewMutualLike{},
	&xo.RoleRevoked{},
	&xo.Swiped{},
}
