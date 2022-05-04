package nft

// Get fields for erc721 contracts
//func GetMetadata() {

//}

//func ERC20FromNetwork(a common.Address, ethconn *ethclient.Client) (*ethereum.ERC20, error) {
//contract, err := NewIERC20(a, ethconn)
//if err != nil {
//return nil, err
//}

//n, err := contract.Name(nil)
//if err != nil {
//// We do not want to return an error when Name function does not exist because it is OPTIONAL
//if !noFunctionErrors[err.Error()] {
//log.Error().Err(err).
//Str("addr", a.Hex()).
//Msg("Name() error, probably bytes32")
//return nil, err
//}
//}
//s, err := contract.Symbol(nil)
//if err != nil {
//// We do not want to return an error when Symbol function does not exist because it is OPTIONAL
//if !noFunctionErrors[err.Error()] {
//log.Error().Err(err).
//Str("addr", a.Hex()).
//Msg("Symbol() error, probably bytes32")
//return nil, err
//}
//}
//return &ethereum.ERC20{
//Added:    timestamppb.Now(),
//Id:       a.Bytes(),
//Name:     n,
//Symbol:   s,
//Decimals: uint32(d),
//}, nil
//}
