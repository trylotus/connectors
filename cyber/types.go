package cyber

import (
	"github.com/trylotus/connectors/cyber/xo"
	"google.golang.org/protobuf/proto"
)

const (
	XOContractAddr = "0xeb110b547d70b6623ea84565649d2dba339dc75b"
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
