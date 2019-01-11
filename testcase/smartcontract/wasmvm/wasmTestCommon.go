package wasmvm

import (
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"fmt"
	"time"
	"github.com/ontio/ontology-test/testframework"
	sdk "github.com/ontio/ontology-go-sdk"
	sdkcom "github.com/ontio/ontology-go-sdk/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/smartcontract/service/wasmvm"
)

func DeployWasmJsonContract(ctx *testframework.TestFrameworkContext, signer *sdk.Account, wasmfile string,contractName string,version string) (common.Uint256, common.Address,error){
	code, err := ioutil.ReadFile(wasmfile)
	if err != nil {
		return common.Uint256{}, common.Address{},err
	}

	codeHash := common.ToHexString(code)

	txHash, err :=  ctx.Ont.WasmVM.DeployWasmVMSmartContract(3597135678,ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		contractName,
		version,
		"author",
		"email",
		"desc",
	)

	if err != nil {
		return common.Uint256{},common.Address{}, fmt.Errorf(" DeploySmartContract error:%s", err)
	}
	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(60 * time.Second)
	if err != nil {
		return common.Uint256{},common.Address{}, fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}

	contractAddr,err := utils.GetContractAddress(codeHash)
	if err != nil{
		return common.Uint256{},common.Address{},err
	}

	//jsonContractAddres = utils.GetContractAddress(codeHash,types.WASMVM)
	return txHash, contractAddr,nil
}



func InvokeWasmContract(ctx  *testframework.TestFrameworkContext, signer *sdk.Account, address common.Address,
	methodName string, paramType wasmvm.ParamType, version byte, params []interface{})(common.Uint256, error){

	return ctx.Ont.WasmVM.InvokeWasmVMSmartContract(3597135678,ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,address,methodName,paramType,version,params)

}

func PreExecWasmContract(ctx  *testframework.TestFrameworkContext,  address common.Address,
	methodName string, paramType wasmvm.ParamType, version byte, params []interface{})(*sdkcom.PreExecResult,error){


		return ctx.Ont.WasmVM.PreExecInvokeNeoVMContract(3597135678,address,methodName,paramType,version,params)


}

//func GetWasmContractAddress(path string) (common.Address,error){
//	code, err := ioutil.ReadFile(path)
//	if err != nil {
//		return common.Address{}, errors.New("")
//	}
//
//	codeHash := common.ToHexString(code)
//	return  utils.GetContractAddress(codeHash,types.WASMVM),nil
//}
