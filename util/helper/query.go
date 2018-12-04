package helper

import (
	"fmt"
	"github.com/irisnet/irishub-sync/types"
	"github.com/irisnet/irishub-sync/util/constant"
	"github.com/pkg/errors"
)

// get validator
func GetValidator(valAddr string) (types.StakeValidator, error) {
	var (
		validatorAddr types.ValAddress
		err           error
		res           types.StakeValidator
	)

	cdc := types.GetCodec()

	validatorAddr, err = types.ValAddressFromBech32(valAddr)

	resRaw, err := Query(types.GetValidatorKey(validatorAddr), constant.StoreNameStake, constant.StoreDefaultEndPath) //TODO
	if err != nil || resRaw == nil {
		return res, err
	}

	res = types.MustUnmarshalValidator(cdc, validatorAddr, resRaw)

	return res, err
}

// get delegation
func GetDelegation(delAddr, valAddr string) (types.Delegation, error) {
	var (
		delegatorAddr types.AccAddress
		validatorAddr types.ValAddress
		err           error

		res types.Delegation
	)
	cdc := types.GetCodec()

	delegatorAddr, err = types.AccAddressFromBech32(delAddr)

	if err != nil {
		return res, err
	}

	validatorAddr, err = types.ValAddressFromBech32(valAddr)

	if err != nil {
		return res, err
	}
	key := types.GetDelegationKey(delegatorAddr, validatorAddr) //TODO

	resRaw, err := Query(key, constant.StoreNameStake, constant.StoreDefaultEndPath)

	if err != nil || resRaw == nil {
		return res, err
	}

	res, err = types.UnmarshalDelegation(cdc, key, resRaw)

	if err != nil {
		return res, err
	}

	return res, err
}

// get unbonding delegation
func GetUnbondingDelegation(delAddr, valAddr string) (types.UnbondingDelegation, error) {
	var (
		delegatorAddr types.AccAddress
		validatorAddr types.ValAddress
		err           error

		res types.UnbondingDelegation
	)

	cdc := types.GetCodec()

	delegatorAddr, err = types.AccAddressFromBech32(delAddr)

	if err != nil {
		return res, err
	}

	validatorAddr, err = types.ValAddressFromBech32(valAddr)

	if err != nil {
		return res, err
	}

	key := types.GetUBDKey(delegatorAddr, validatorAddr) //TODO ValAddressFromBech32

	resRaw, err := Query(key, constant.StoreNameStake, constant.StoreDefaultEndPath)

	if err != nil || resRaw == nil {
		return res, err
	}

	res = types.MustUnmarshalUBD(cdc, key, resRaw)

	return res, nil
}

// Query from Tendermint with the provided storename and path
func Query(key types.HexBytes, storeName string, endPath string) (res []byte, err error) {
	path := fmt.Sprintf("/store/%s/%s", storeName, endPath)
	client := GetClient()
	defer client.Release()

	opts := types.ABCIQueryOptions{
		Height: 0,
		Prove:  false, //不需要验证prof
	}
	result, err := client.ABCIQueryWithOptions(path, key, opts)
	if err != nil {
		return res, err
	}
	resp := result.Response
	if resp.Code != uint32(0) {
		return res, errors.Errorf("Query failed: (%d) %s", resp.Code, resp.Log)
	}
	return resp.Value, nil
}

func QuerySubspace(cdc *types.Codec, subspace []byte, storeName string) (res []types.KVPair, err error) {
	resRaw, err := Query(subspace, storeName, "subspace")
	if err != nil {
		return res, err
	}
	cdc.MustUnmarshalBinaryLengthPrefixed(resRaw, &res)
	return
}
