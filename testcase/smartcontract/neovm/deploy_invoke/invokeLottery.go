package deploy_invoke

import (
	"github.com/ontio/ontology-test/testframework"
	"io/ioutil"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology-go-sdk/utils"
	"time"
	"fmt"
)

func TestLottery(ctx *testframework.TestFrameworkContext) bool {


	account2,_ := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	account3,_ := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")
	account4,_ := ctx.GetAccount("ALerVnMj3eNk9xe8BnQJtoWvwGmY3x4KMi")
	account5,_ := ctx.GetAccount("AKmowTi8NcAMjZrg7ZNtSQUtnEgdaC65wG")

	addr,_ := common.AddressFromBase58("ASYkgyWm4GFiXqVKZs6XrjaN3HnFVGRhDs")
	fmt.Println(common.ToHexString(addr[:]))

	avmfile := "test_data/lottery.avm"

	code, err := ioutil.ReadFile(avmfile)
	if err != nil {
		return false
	}
	codeHash := common.ToHexString(code)

	codeAddress, _ := utils.GetContractAddress(codeHash)
	fmt.Println("contract address:"+codeAddress.ToBase58())

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestLottery GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.Rpc.DeploySmartContract(500, 10300000000,
		signer,
		true,
		codeHash,
		"TestLottery",
		"1.0",
		"",
		"",
		"",
	)

	if err != nil {
		ctx.LogError("TestLottery DeploySmartContract error: %s", err)
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestLottery WaitForGenerateBlock error: %s", err)
		return false
	}


	gameCount := 10

		flag := true
		if flag {

			ctx.LogInfo("=============end game start=======================")
			txHash, err := ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
				signer,
				codeAddress,
				[]interface{}{"endGame", []interface{}{signer.Address[:],gameCount}})
			if err != nil {
				ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
			}

			//WaitForGenerateBlock
			_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
			if err != nil {
				ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
				return false
			}

			//GetEventLog, to check the result of invoke
			events, err := ctx.Ont.Rpc.GetSmartContractEvent(txHash)
			if err != nil {
				ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
				return false
			}
			if events.State == 0 {
				ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
				return false
			}

			for _,notify:= range events.Notify{
				ctx.LogInfo("%+v", notify)
			}


			ctx.LogInfo("=============end game end=======================")

		}


	if !flag {
		ctx.LogInfo("--------------------testing query round --------------------")
		tx, err := ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"queryCurrentRound", []interface{}{gameCount}})
		if err != nil {
			ctx.LogError("TestLottery NewNeoVMSInvokeTransaction error:%s", err)

			return false
		}
		err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
		if err != nil {
			ctx.LogError("TestLottery SignToTransaction error:%s", err)

			return false
		}


		obj,err:=ctx.Ont.Rpc.PrepareInvokeContract(tx)
		if err != nil {
			ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

			return false
		}

		bs,err:= common.HexToBytes(obj.Result.(string))
		if err != nil{
			ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

			return false
		}
		round := common.BigIntFromNeoBytes(bs)
		//
		fmt.Printf("round is %d\n",round.Int64())
		ctx.LogInfo("--------------------testing query round end--------------------")



		ctx.LogInfo("=============acct1 attend start=======================")
		txHash, err := ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			signer,
			codeAddress,
			[]interface{}{"attend", []interface{}{signer.Address[:],gameCount}})
		if err != nil {
			ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err := ctx.Ont.Rpc.GetSmartContractEvent(txHash)
		if err != nil {
			ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
			return false
		}

		for _,notify:= range events.Notify{
			ctx.LogInfo("%+v", notify)
		}


		ctx.LogInfo("=============acct1 attend end=======================")


		ctx.LogInfo("=============acct2 attend start=======================")
		txHash, err = ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			account2,
			codeAddress,
			[]interface{}{"attend", []interface{}{account2.Address[:],gameCount}})
		if err != nil {
			ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
		if err != nil {
			ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
			return false
		}
		for _,notify:= range events.Notify{
			ctx.LogInfo("%+v", notify)
		}


		ctx.LogInfo("=============acct2 attend end=======================")

		ctx.LogInfo("=============acct3 attend start=======================")
		txHash, err = ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			account3,
			codeAddress,
			[]interface{}{"attend", []interface{}{account3.Address[:],gameCount}})
		if err != nil {
			ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
		if err != nil {
			ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
			return false
		}
		for _,notify:= range events.Notify{
			ctx.LogInfo("%+v", notify)
		}


		ctx.LogInfo("=============acct3 attend end=======================")

		ctx.LogInfo("=============acct4 attend start=======================")
		txHash, err = ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			account4,
			codeAddress,
			[]interface{}{"attend", []interface{}{account4.Address[:],gameCount}})
		if err != nil {
			ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
		if err != nil {
			ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
			return false
		}
		for _,notify:= range events.Notify{
			ctx.LogInfo("%+v", notify)
		}

		ctx.LogInfo("=============acct4 attend end=======================")


		ctx.LogInfo("=============acct5 attend start=======================")
		txHash, err = ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
			account5,
			codeAddress,
			[]interface{}{"attend", []interface{}{account5.Address[:],gameCount}})
		if err != nil {
			ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		}

		//WaitForGenerateBlock
		_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 1)
		if err != nil {
			ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
			return false
		}

		//GetEventLog, to check the result of invoke
		events, err = ctx.Ont.Rpc.GetSmartContractEvent(txHash)
		if err != nil {
			ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
			return false
		}
		if events.State == 0 {
			ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
			return false
		}
		for _,notify:= range events.Notify{
			ctx.LogInfo("%+v", notify)
		}


		ctx.LogInfo("=============acct5 attend end=======================")


		ctx.LogInfo("--------------------testing query round --------------------")
		tx, err = ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"queryCurrentRound", []interface{}{gameCount}})
		if err != nil {
			ctx.LogError("TestLottery NewNeoVMSInvokeTransaction error:%s", err)

			return false
		}
		err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
		if err != nil {
			ctx.LogError("TestLottery SignToTransaction error:%s", err)

			return false
		}


		obj,err=ctx.Ont.Rpc.PrepareInvokeContract(tx)
		if err != nil {
			ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

			return false
		}

		bs,err= common.HexToBytes(obj.Result.(string))
		if err != nil{
			ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

			return false
		}
		round = common.BigIntFromNeoBytes(bs)
		//
		fmt.Printf("current round is %d\n",round.Int64())
		ctx.LogInfo("--------------------testing query round end--------------------")


		ctx.LogInfo("--------------------testing query queryWinner --------------------")
		tx, err = ctx.Ont.Rpc.NewNeoVMSInvokeTransaction(ctx.GetGasPrice(), ctx.GetGasLimit(),codeAddress, []interface{}{"queryWinner", []interface{}{gameCount,round.Int64()}})
		if err != nil {
			ctx.LogError("TestLottery NewNeoVMSInvokeTransaction error:%s", err)

			return false
		}
		err = ctx.Ont.Rpc.SignToTransaction(tx, signer)
		if err != nil {
			ctx.LogError("TestLottery SignToTransaction error:%s", err)

			return false
		}


		obj,err=ctx.Ont.Rpc.PrepareInvokeContract(tx)
		if err != nil {
			ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

			return false
		}

		bs,err= common.HexToBytes(obj.Result.(string))
		if err != nil{
			ctx.LogError("TestLottery PrepareInvokeContract error:%s", err)

			return false
		}

		fmt.Printf("bs is %v:\n",bs)

		winner,err := common.AddressParseFromBytes(bs)
		if err != nil{
			ctx.LogError("TestLottery AddressParseFromBytes error:%s", err)

			return false
		}
		//
		fmt.Printf("winner is %v\n",winner.ToBase58())
		ctx.LogInfo("--------------------testing query queryWinner end--------------------")

	}


	return true
}
