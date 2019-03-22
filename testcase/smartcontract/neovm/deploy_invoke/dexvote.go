package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
	"io/ioutil"
	"fmt"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"time"
)

func DEXvoteTest(ctx *testframework.TestFrameworkContext) bool {

	avmfile := "test_data/ONTDEX_Vote.avm"

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
	//account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	//account4, err := ctx.GetAccount("ALerVnMj3eNk9xe8BnQJtoWvwGmY3x4KMi")
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


	votekey := "111111"
	ctx.LogInfo("--------------------testing addvote--------------------")
	now := time.Now().Unix()
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"addVote", []interface{}{votekey,now+120,"vote for listing",[]interface{}{"BTC","ETH","EOS"}}})
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
	ctx.LogInfo("--------------------testing addvote end--------------------")

	ctx.LogInfo("--------------------testing getvote start --------------------")
	obj, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress,
		[]interface{}{"getVote", []interface{}{votekey}})
	if err != nil {
		ctx.LogError("getvote NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	bs, err := obj.Result.ToString()
	if err != nil {
		ctx.LogError("getvote PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("getvote is %s\n", bs)
	ctx.LogInfo("--------------------testing getvote end--------------------")
	ctx.LogInfo("--------------------testing vote--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"vote", []interface{}{signer.Address,votekey,"BTC",100}})
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
	ctx.LogInfo("--------------------testing addvote end--------------------")
	ctx.LogInfo("--------------------testing getvote start --------------------")
	obj, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress,
		[]interface{}{"getVote", []interface{}{votekey}})
	if err != nil {
		ctx.LogError("getvote NewNeoVMSInvokeTransaction error:%s", err)

		return false
	}

	bs, err = obj.Result.ToString()
	if err != nil {
		ctx.LogError("getvote PrepareInvokeContract error:%s", err)

		return false
	}

	//
	fmt.Printf("getvote is %s\n", bs)
	ctx.LogInfo("--------------------testing getvote end--------------------")

	ctx.LogInfo("--------------------testing withdraw--------------------")
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"withdraw", []interface{}{signer.Address,votekey}})
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
	ctx.LogInfo("--------------------testing withdraw end--------------------")




	return true
}
