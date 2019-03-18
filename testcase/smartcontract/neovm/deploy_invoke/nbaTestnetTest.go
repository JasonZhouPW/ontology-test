package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"time"
)

func NBAGuessTestnet(ctx *testframework.TestFrameworkContext) bool {

	//admin, err := ctx.GetDefaultAccount()
	operator, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	codeAddress, err := common.AddressFromHexString("5f8f8d84b3db1e134d14aca49bae37e7c294d53e")

	ctx.LogInfo("--------------------testing inputMatch--------------------")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		operator,
		codeAddress,
		[]interface{}{"inputMatch", []interface{}{"20181129", "0021800316", "1610612761", "1610612744"}})
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
		ctx.LogError("NBAGuess failed invoked exec state return 0")
		return false
	}
	ctx.LogInfo("--------------------testing inputMatch end--------------------")

	ctx.LogInfo("--------------------testing inputMatch--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		operator,
		codeAddress,
		[]interface{}{"inputMatch", []interface{}{"20181129", "0021800317", "1610612747", "1610612754"}})
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
	ctx.LogInfo("--------------------testing inputMatch end--------------------")

	ctx.LogInfo("--------------------testing inputMatch--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		operator,
		codeAddress,
		[]interface{}{"inputMatch", []interface{}{"20181129", "0021800318", "1610612758", "1610612746"}})
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
	ctx.LogInfo("--------------------testing inputMatch end--------------------")

	ctx.LogInfo("--------------------testing gameBets --------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getBetsByGameid", []interface{}{"0021800316"}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	bets, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("getBetsByGameid is %s\n", bets)
	ctx.LogInfo("--------------------testing gameBets end--------------------")

	ctx.LogInfo("--------------------testing getMatchByDate account3--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getMatchByDate", []interface{}{"20181129"}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("getMatchByDate is %s\n", balance)
	ctx.LogInfo("--------------------testing balanceOf account2 end--------------------")

	account3, err := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")
	account4, err := ctx.GetAccount("ALerVnMj3eNk9xe8BnQJtoWvwGmY3x4KMi")
	account5, err := ctx.GetAccount("AKmowTi8NcAMjZrg7ZNtSQUtnEgdaC65wG")

	ctx.LogInfo("===========before bet balance is ")
	ctx.LogInfo("--------------------testing balanceOf account3--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAccountBalance", []interface{}{account3.Address[:]}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	b, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", b)
	ctx.LogInfo("--------------------testing balanceOf account3 end--------------------")

	ctx.LogInfo("--------------------testing balanceOf account4--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAccountBalance", []interface{}{account4.Address[:]}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	b, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", b)
	ctx.LogInfo("--------------------testing balanceOf account4 end--------------------")

	ctx.LogInfo("--------------------testing balanceOf account5--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAccountBalance", []interface{}{account5.Address[:]}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	b, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", b)
	ctx.LogInfo("--------------------testing balanceOf account5 end--------------------")
	ctx.LogInfo("===========before bet balance end ")

	ctx.LogInfo("--------------------testing placebet acct3--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"placeBet", []interface{}{account3.Address, "0021800316", "H", 1 * multiple}})
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

	ctx.LogInfo("--------------------testing placebet acct4--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account4,
		codeAddress,
		[]interface{}{"placeBet", []interface{}{account4.Address, "0021800316", "H", 1 * multiple}})
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

	ctx.LogInfo("--------------------testing placebet acct5--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account5,
		codeAddress,
		[]interface{}{"placeBet", []interface{}{account5.Address, "0021800316", "V", 1 * multiple}})
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

	ctx.LogInfo("--------------------testing gameBets --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getBetsByGameid", []interface{}{"0021800316"}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	bets, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("getBetsByGameid is %s\n", bets)
	ctx.LogInfo("--------------------testing gameBets end--------------------")

	ctx.LogInfo("--------------------testing endbet--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		operator,
		codeAddress,
		[]interface{}{"endBet", []interface{}{"20181129"}})
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

	ctx.LogInfo("--------------------testing callOracle--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		operator,
		codeAddress,
		[]interface{}{"callOracle", []interface{}{"20181129"}})
	if err != nil {
		ctx.LogError("NBAGuess init error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}
	ctx.LogInfo("===callOracle txhash is :" + txHash.ToHexString())

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
	ctx.LogInfo("--------------------testing callOracle end--------------------")

	_, err = ctx.Ont.WaitForGenerateBlock(60*time.Second, 1)
	if err != nil {
		ctx.LogError("NBAGuess WaitForGenerateBlock error: %s", err)
		return false
	}

	ctx.LogInfo("--------------------testing testOracleRes--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		operator,
		codeAddress,
		[]interface{}{"testOracleRes", []interface{}{"20181129"}})
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
	ctx.LogInfo("--------------------testing testOracleRes end--------------------")

	ctx.LogInfo("--------------------testing setResult--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		operator,
		codeAddress,
		[]interface{}{"setResult", []interface{}{"20181129"}})
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
	ctx.LogInfo("--------------------testing setResult end--------------------")

	ctx.LogInfo("--------------------testing balanceOf account3--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAccountBalance", []interface{}{account3.Address[:]}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	b, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", b)
	ctx.LogInfo("--------------------testing balanceOf account3 end--------------------")

	ctx.LogInfo("--------------------testing balanceOf account4--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAccountBalance", []interface{}{account4.Address[:]}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	b, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", b)
	ctx.LogInfo("--------------------testing balanceOf account4 end--------------------")

	ctx.LogInfo("--------------------testing balanceOf account5--------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryAccountBalance", []interface{}{account5.Address[:]}})
	if err != nil {
		ctx.LogError("NBAGuess NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	b, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("NBAGuess PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balance is %d\n", b)
	ctx.LogInfo("--------------------testing balanceOf account5 end--------------------")

	return true
}
