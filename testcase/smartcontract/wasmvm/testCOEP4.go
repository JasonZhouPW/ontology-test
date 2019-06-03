package wasmvm

import (
	"fmt"
	"github.com/ontio/ontology-test/testframework"
	"time"
	"encoding/binary"
)

func TestOEP4C(ctx *testframework.TestFrameworkContext) bool {
	testFile := filePath + "/" + "OEP4.wasm"
	signer, _ := ctx.GetDefaultAccount()
	timeoutSec := 30 * time.Second
	txhash, addr, err := DeployWasmJsonContract(ctx, signer, testFile, "testContract", "1")
	if err != nil {
		fmt.Printf("deploy failed:%s\n", err.Error())
		return false
	}
	ctx.LogInfo("==contract address is %s", addr.ToBase58())
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil {
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}

	ctx.LogInfo("=====================invoke init==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"init",
		[]interface{}{})

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
	ctx.LogInfo("=====================invoke init end==============================")

/*	ctx.LogInfo("=====================invoke balanceOf==============================")
	res, err := PreExecWasmContract(ctx,
		addr,
		"balanceof",
		[]interface{}{signer.Address})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err := res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("balance of %s is %d\n", signer.Address.ToBase58(), binary.LittleEndian.Uint64(bs))

	ctx.LogInfo("=====================invoke balanceOf end==============================")*/


	ctx.LogInfo("=====================invoke test==============================")
	res, err := PreExecWasmContract(ctx,
		addr,
		"name",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke test failed:%s\n", err.Error())
		return false
	}

	bs, err := res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("test is %s\n", bs)

	ctx.LogInfo("=====================invoke test end==============================")


	ctx.LogInfo("=====================invoke transfer==============================")

	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"transfer",
		[]interface{}{signer.Address, account2.Address, uint64(500)})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}

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
	fmt.Printf("events.Notify:%v\n", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transfer end==============================")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"balanceOf",
		[]interface{}{signer.Address})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("balance of %s is %d\n", signer.Address.ToBase58(), binary.LittleEndian.Uint64(bs))

	ctx.LogInfo("=====================invoke balanceOf end==============================")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"balanceOf",
		[]interface{}{account2.Address})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("balance of %s is %d\n", account2.Address.ToBase58(), binary.LittleEndian.Uint64(bs))

	ctx.LogInfo("=====================invoke balanceOf end==============================")

	ctx.LogInfo("=====================invoke approve==============================")

	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"approve",
		[]interface{}{signer.Address, account2.Address, uint64(800)})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}

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
	fmt.Printf("events.Notify:%v\n", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transfer end==============================")

	ctx.LogInfo("=====================invoke allowance==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"allowance",
		[]interface{}{signer.Address,account2.Address})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("balance of %s is %d\n", account2.Address.ToBase58(), binary.LittleEndian.Uint64(bs))

	ctx.LogInfo("=====================invoke allowance end==============================")


	ctx.LogInfo("=====================invoke tranferFrom==============================")

	txhash, err = InvokeWasmContract(ctx,
		account2,
		addr,
		"transferfrom",
		[]interface{}{account2.Address,signer.Address, account2.Address, uint64(800)})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}

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
	fmt.Printf("events.Notify:%v\n", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transferfrom end==============================")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"balanceOf",
		[]interface{}{account2.Address})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("balance of %s is %d\n", account2.Address.ToBase58(), binary.LittleEndian.Uint64(bs))

	ctx.LogInfo("=====================invoke balanceOf end==============================")
	return true
}
