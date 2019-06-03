package wasmvm

import (
	"time"
	"fmt"
	"github.com/ontio/ontology-test/testframework"
	"encoding/binary"
)

func Testclass(ctx *testframework.TestFrameworkContext) bool {

	testFile := filePath + "/" + "TestClass.wasm"
	signer, _ := ctx.GetDefaultAccount()
	timeoutSec := 30 * time.Second
	txhash, addr, err := DeployWasmJsonContract(ctx, signer, testFile, "testContract", "1")
	if err != nil {
		fmt.Printf("deploy failed:%s\n", err.Error())
		return false
	}
	ctx.LogInfo("==contract address is %s", addr.ToBase58())
	//account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil {
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}

	ctx.LogInfo("=====================invoke init==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"init",
		[]interface{}{signer.Address})

	_, err = ctx.Ont.WaitForGenerateBlock(timeoutSec)
	if err != nil {
		return false
	}

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

	ctx.LogInfo("=====================invoke getscore==============================")
	res, err := PreExecWasmContract(ctx,
		addr,
		"getscore",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err := res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("getscore  is %d\n",  binary.LittleEndian.Uint32(bs))

	ctx.LogInfo("=====================invoke totalSupply end==============================")

	ctx.LogInfo("=====================invoke getname==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"getname",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("getname  is %s\n",  bs)

	ctx.LogInfo("=====================invoke getname end==============================")


	ctx.LogInfo("=====================invoke setname==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"setname",
		[]interface{}{"new name"})

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
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke init end==============================")

	ctx.LogInfo("=====================invoke getname==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"getname",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("getname  is  %s\n",  bs)

	ctx.LogInfo("=====================invoke getname end==============================")


	ctx.LogInfo("=====================invoke setscore==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"setscore",
		[]interface{}{100})

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
	fmt.Printf("events.Notify:%v", events.Notify)
	for _, notify := range events.Notify {
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke init end==============================")

	ctx.LogInfo("=====================invoke getscore==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"getscore",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("getscore  is %d\n",  binary.LittleEndian.Uint32(bs))

	ctx.LogInfo("=====================invoke totalSupply end==============================")

	return true
}
