package deploy_invoke

import (
	"time"
	"io/ioutil"
	"fmt"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
)

func DEXTest(ctx *testframework.TestFrameworkContext) bool {
	avmfile := "test_data/ONTDEX.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress base58===%s", codeAddress.ToBase58())
	signer, err := ctx.GetDefaultAccount()
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	account3, err := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")

	priceMultiple := 1000000000
	ctx.LogInfo("=================Deploy===============================")

	if err != nil {
		ctx.LogError("Dice GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"DEXGuess",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("Dice DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}
	ctx.LogInfo("=================Deploy end===========================")

	var txHash common.Uint256

	ctx.LogInfo("--------------------testing init--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"init", []interface{}{}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing init end--------------------")


	ctx.LogInfo("--------------------testing add sell ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account2.Address[:],1,100,1,2*priceMultiple }})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing add buy ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account3.Address[:],1,200,0,2*priceMultiple}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing add order end--------------------")


	ctx.LogInfo("--------------------testing match order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"matchOrder", []interface{}{2,1,100,2*priceMultiple,1}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing match end--------------------")


	ctx.LogInfo("--------------------testing add sell ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account2.Address[:],1,100,1,2*priceMultiple}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing cancel sell ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"cancelOrder", []interface{}{account2.Address[:],3}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing cancel order end--------------------")

	ctx.LogInfo("--------------------testing add buy ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account3.Address[:],1,200,0,2*priceMultiple}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing cancel sell ONT order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"cancelOrder", []interface{}{account3.Address[:],4}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing cancel order end--------------------")
	ctx.LogInfo("====================testing OEP4 case ==================")
	ctx.LogInfo("--------------------testing add pair--------------------")
	addr,_ := common.AddressFromBase58("APPWgNbWvUdQjQxeN7RduYweH3caaM1LM1")
	ong,_ := common.AddressFromBase58("AFmseVrdL9f9oyCzZefL9tG6UbvhfRZMHJ")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"addPair", []interface{}{2,addr,8,ong,9}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing addpair end--------------------")

	ctx.LogInfo("--------------------testing set fee--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"setTradeFee", []interface{}{5000}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing setfee end--------------------")

	ctx.LogInfo("--------------------testing add sell OEP4 order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account2.Address[:],2,100,1,priceMultiple/100}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing add order end--------------------")

	ctx.LogInfo("--------------------testing add buy OEP4 order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"addOrder", []interface{}{account3.Address[:],2,100,0,priceMultiple/100}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing add order end--------------------")


	ctx.LogInfo("--------------------testing match order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"matchOrder", []interface{}{6,5,100,priceMultiple/100,1}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing match end--------------------")

	ctx.LogInfo("--------------------testing withdrawAssets order--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"withdrawAssets", []interface{}{}})
	if err != nil {
		ctx.LogError("Dice invest error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("Dice GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("Dice failed invoked exec state return 0")
		return false
	}
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing withdrawAssets end--------------------")






	return true
}
