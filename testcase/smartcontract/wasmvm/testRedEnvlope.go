package wasmvm

import (
	"github.com/ontio/ontology-test/testframework"
	"fmt"
	"time"
	"github.com/ontio/ontology/common"
)

func TestRedEnvlope(ctx *testframework.TestFrameworkContext) bool {


	testFile := filePath + "/" + "redEnvlope.wasm"
	signer, _ := ctx.GetDefaultAccount()
	timeoutSec := 30 * time.Second
	txhash, addr, err := DeployWasmJsonContract(ctx, signer, testFile, "testContract", "1")
	if err != nil {
		fmt.Printf("deploy failed:%s\n", err.Error())
		return false
	}
	fmt.Printf("txhash:%s\n",txhash.ToHexString())
	ctx.LogInfo("==contract address is %s", addr.ToBase58())

	//ONGAddr ,_:= common.AddressFromBase58("AFmseVrdL9f9oyCzZefL9tG6UbvhfRZMHJ")
	//ONTAddr ,_:= common.AddressFromBase58("AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV")
	OEP4Addr,_:=common.AddressFromBase58("AWuLhVKPqpXfizhWW1ksMBPxCTwvN6C3Vz")

	ctx.LogInfo("=====================invoke createEnvlope==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"createRedEnvlope",
		[]interface{}{signer.Address,uint64(3),uint64(10000),OEP4Addr})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}
	fmt.Printf("init txhash:%s\n",txhash.ToHexString())
	events, err := ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	hash := ""
	if len(events.Notify) == 3{
		states := events.Notify[2].States.([]interface{})
		if len(states) == 3{
			hash = states[2].(string)
		}
	}
	ctx.LogInfo("=====================invoke create end==============================")


	ctx.LogInfo("=====================invoke query==============================")
	res, err := PreExecWasmContract(ctx,
		addr,
		"queryEnvlope",
		[]interface{}{hash})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err := res.Result.ToByteArray()

	fmt.Printf("res is %s\n", bs)


	ctx.LogInfo("=====================invoke query end==============================")
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")

	ctx.LogInfo("=====================invoke claim==============================")
	txhash, err = InvokeWasmContract(ctx,
		account2,
		addr,
		"claimEnvlope",
		[]interface{}{account2.Address,hash})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}
	fmt.Printf("init txhash:%s\n",txhash.ToHexString())
	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("=====================invoke claim end==============================")
	ctx.LogInfo("=====================invoke query==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"queryEnvlope",
		[]interface{}{hash})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %s\n", bs)


	ctx.LogInfo("=====================invoke query end==============================")
	account3, err := ctx.GetAccount("AK98G45DhmPXg4TFPG1KjftvkEaHbU8SHM")
	ctx.LogInfo("=====================invoke claim==============================")
	txhash, err = InvokeWasmContract(ctx,
		account3,
		addr,
		"claimEnvlope",
		[]interface{}{account3.Address,hash})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}
	fmt.Printf("init txhash:%s\n",txhash.ToHexString())
	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("=====================invoke claim end==============================")
	ctx.LogInfo("=====================invoke query==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"queryEnvlope",
		[]interface{}{hash})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %s\n", bs)


	ctx.LogInfo("=====================invoke query end==============================")
	account4, err := ctx.GetAccount("ALerVnMj3eNk9xe8BnQJtoWvwGmY3x4KMi")
	ctx.LogInfo("=====================invoke claim==============================")
	txhash, err = InvokeWasmContract(ctx,
		account4,
		addr,
		"claimEnvlope",
		[]interface{}{account4.Address,hash})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}
	fmt.Printf("init txhash:%s\n",txhash.ToHexString())
	events, err = ctx.Ont.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		ctx.LogError("TestWasmOEP4 GetSmartContractEvent error:%s", err)
		return false
	}
	fmt.Printf("event is %v\n", events)
	if events.State == 0 {
		ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
		return false
	}
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("=====================invoke claim end==============================")
	ctx.LogInfo("=====================invoke query==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"queryEnvlope",
		[]interface{}{hash})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %s\n", bs)


	ctx.LogInfo("=====================invoke query end==============================")
	return true

}

