package woofi

import (
	"github.com/trylotus/connectors/woofi/WOOPP"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"lotus.woofi.0_0_0.woopp_feemanagerupdated":         &WOOPP.FeeManagerUpdated{},
	"lotus.woofi.0_0_0.woopp_ownershiptransferprepared": &WOOPP.OwnershipTransferPrepared{},
	"lotus.woofi.0_0_0.woopp_ownershiptransferred":      &WOOPP.OwnershipTransferred{},
	"lotus.woofi.0_0_0.woopp_parametersupdated":         &WOOPP.ParametersUpdated{},
	"lotus.woofi.0_0_0.woopp_paused":                    &WOOPP.Paused{},
	"lotus.woofi.0_0_0.woopp_rewardmanagerupdated":      &WOOPP.RewardManagerUpdated{},
	"lotus.woofi.0_0_0.woopp_strategistupdated":         &WOOPP.StrategistUpdated{},
	"lotus.woofi.0_0_0.woopp_unpaused":                  &WOOPP.Unpaused{},
	"lotus.woofi.0_0_0.woopp_withdraw":                  &WOOPP.Withdraw{},
	"lotus.woofi.0_0_0.woopp_wooguardianupdated":        &WOOPP.WooGuardianUpdated{},
	"lotus.woofi.0_0_0.woopp_wooracleupdated":           &WOOPP.WooracleUpdated{},
	"lotus.woofi.0_0_0.woopp_wooswap":                   &WOOPP.WooSwap{},
}

var ABIs = map[string]string{
	"WOOPP": WOOPP.WOOPPABI,
}
