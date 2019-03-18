package deploy_invoke

import (
	"fmt"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"time"
)

func DiceTest(ctx *testframework.TestFrameworkContext) bool {
	avmfile := "test_data/DiceGumbling.avm"

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

	ctx.LogInfo("=================Deploy===============================")

	if err != nil {
		ctx.LogError("Dice GetDefaultAccount error:%s", err)
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
		ctx.LogError("Dice DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("Dice WaitForGenerateBlock error: %s", err)
		return false
	}
	ctx.LogInfo("=================Deploy end===========================")

	ctx.LogInfo("--------------------testing invest--------------------")
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"invest", []interface{}{account2.Address[:], 2000000000000}})
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
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing invest end--------------------")

	ctx.LogInfo("--------------------testing poolBalance --------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"poolBalance", []interface{}{}})
	if err != nil {
		ctx.LogError("Dice NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err := obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("Dice PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("poolBalance is %d\n", balance)
	ctx.LogInfo("--------------------testing poolBalance end--------------------")

	ctx.LogInfo("--------------------testing totalInvestorCount --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"totalInvestorCount", []interface{}{}})
	if err != nil {
		ctx.LogError("Dice NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("Dice PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("totalInvestorCount is %d\n", balance)
	ctx.LogInfo("--------------------testing totalInvestorCount end--------------------")

	ctx.LogInfo("--------------------testing totalInvests --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"totalInvests", []interface{}{}})
	if err != nil {
		ctx.LogError("Dice NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("Dice PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("totalInvests is %d\n", balance)
	ctx.LogInfo("--------------------testing totalInvests end--------------------")

	ctx.LogInfo("--------------------testing placeBet--------------------")

	//bets := make([]interface{},51)
	//bets[0] = account3.Address[:]
	//for i :=1; i< 51;i++{
	//	bets[i] = 1000000000
	//}

	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account3,
		codeAddress,
		[]interface{}{"placeBet", []interface{}{account3.Address[:], 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000,
			1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000,
			1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000,
			1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000,
			1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000}})
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
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing placeBet end--------------------")

	ctx.LogInfo("--------------------testing balanceOf --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"balanceOf", []interface{}{account3.Address[:]}})
	if err != nil {
		ctx.LogError("Dice NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("Dice PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balanceOf is %d\n", balance)
	ctx.LogInfo("--------------------testing balanceOf end--------------------")

	ctx.LogInfo("--------------------testing poolBalance --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"poolBalance", []interface{}{}})
	if err != nil {
		ctx.LogError("Dice NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("Dice PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("poolBalance is %d\n", balance)
	ctx.LogInfo("--------------------testing poolBalance end--------------------")

	ctx.LogInfo("--------------------testing queryInvest --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"queryInvest", []interface{}{account2.Address[:]}})
	if err != nil {
		ctx.LogError("Dice NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("Dice PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("queryInvest is %d\n", balance)
	ctx.LogInfo("--------------------testing queryInvest end--------------------")

	ctx.LogInfo("--------------------testing quitInvest--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"quitInvest", []interface{}{account2.Address[:]}})
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
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing quitInvest end--------------------")

	ctx.LogInfo("--------------------testing poolBalance --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"poolBalance", []interface{}{}})
	if err != nil {
		ctx.LogError("Dice NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("Dice PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("poolBalance is %d\n", balance)
	ctx.LogInfo("--------------------testing poolBalance end--------------------")

	ctx.LogInfo("--------------------testing balanceOf --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"balanceOf", []interface{}{account2.Address[:]}})
	if err != nil {
		ctx.LogError("Dice NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	balance, err = obj.Result.ToInteger()
	if err != nil {
		ctx.LogError("Dice PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("balanceOf acct2 is %d\n", balance)
	ctx.LogInfo("--------------------testing balanceOf end--------------------")

	ctx.LogInfo("--------------------testing withdraw--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		account2,
		codeAddress,
		[]interface{}{"withdraw", []interface{}{account2.Address[:]}})
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
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("--------------------testing quitInvest end--------------------")

	return true
}
