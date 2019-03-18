package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"io/ioutil"
	"time"
)

const multiple = 1000000000

func NBAGuessTest(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/nbaGuess.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		return false
	}
	codeHash := string(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	ctx.LogInfo("=====CodeAddress base58===%s", codeAddress.ToBase58())
	signer, err := ctx.GetDefaultAccount()
	ctx.LogInfo("=================Deploy===============================")

	if err != nil {
		ctx.LogError("NBAGuess GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		codeHash,
		"NBAGuess",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("NBAGuess DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}
	ctx.LogInfo("=================Deploy end===========================")

	ctx.LogInfo("--------------------testing inputMatch--------------------")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"inputMatch", []interface{}{"20181201", "g001", "t001", "t002"}})
	if err != nil {
		ctx.LogError("NBAGuess init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("NBAGuess GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	ctx.LogInfo("--------------------testing inputMatch end--------------------")

	ctx.LogInfo("--------------------testing inputMatch--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"inputMatch", []interface{}{"20181202", "g003", "t005", "t006"}})
	if err != nil {
		ctx.LogError("NBAGuess init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("NBAGuess GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestOEP4Py failed invoked exec state return 0")
		return false
	}
	ctx.LogInfo("--------------------testing inputMatch end--------------------")

	ctx.LogInfo("--------------------testing getMatchByDate--------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getMatchByDate", []interface{}{"20181201"}})

	matches, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("getMatchByDate PrepareInvokeContract error:%s", err)

		return false
	}

	fmt.Printf("match is %s\n", matches)
	ctx.LogInfo("--------------------testing getMatchByDate end--------------------")

	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	ctx.LogInfo("--------------------testing placebet--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"placeBet", []interface{}{account2.Address, "g003", "H", 1 * multiple}})
	if err != nil {
		ctx.LogError("NBAGuess init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("NBAGuess GetSmartContractEvent error:%s", err)
		//return false
	}
	if events == nil {
		ctx.LogError("NBAGuess failed invoked exec failed")
		return false
	}

	if events.State == 0 {
		ctx.LogError("NBAGuess failed invoked exec state return 0")
		return false
	}

	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing placebet end--------------------")

	account3, err := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")
	ctx.LogInfo("--------------------testing placebet--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"placeBet", []interface{}{account3.Address, "g003", "H", 1 * multiple}})
	if err != nil {
		ctx.LogError("NBAGuess init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("NBAGuess GetSmartContractEvent error:%s", err)
		//return false
	}

	if events == nil {
		ctx.LogError("NBAGuess failed invoked exec failed")
		return false
	}

	if events.State == 0 {
		ctx.LogError("NBAGuess failed invoked exec state return 0")
		return false
	}

	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing placebet end--------------------")

	ctx.LogInfo("--------------------testing placebet--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"placeBet", []interface{}{account3.Address, "g003", "V", 1 * multiple}})
	if err != nil {
		ctx.LogError("NBAGuess init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("NBAGuess GetSmartContractEvent error:%s", err)
		//return false
	}

	if events == nil {
		ctx.LogError("NBAGuess failed invoked exec failed")
		return false
	}

	if events.State == 0 {
		ctx.LogError("NBAGuess failed invoked exec state return 0")
		return false
	}

	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing placebet end--------------------")

	ctx.LogInfo("--------------------testing endbet--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"endBet", []interface{}{"20181202"}})
	if err != nil {
		ctx.LogError("NBAGuess init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("NBAGuess GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("NBAGuess failed invoked exec state return 0")
		return false
	}

	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing endbet end--------------------")

	ctx.LogInfo("--------------------testing manualSetResult--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"manualSetResult", []interface{}{"20181202", 1, "g003", 100, 98}})
	if err != nil {
		ctx.LogError("NBAGuess init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetEventLog, to check the result of invoke
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("NBAGuess GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("NBAGuess failed invoked exec state return 0")
		return false
	}

	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing manualSetResult end--------------------")

	ctx.LogInfo("--------------------testing balanceOf account32--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAccountBalance", []interface{}{account2.Address[:]}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", balance)
	ctx.LogInfo("--------------------testing balanceOf account2 end--------------------")

	ctx.LogInfo("--------------------testing balanceOf account3--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAccountBalance", []interface{}{account3.Address[:]}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", balance)
	ctx.LogInfo("--------------------testing balanceOf account3 end--------------------")
	return true

}
