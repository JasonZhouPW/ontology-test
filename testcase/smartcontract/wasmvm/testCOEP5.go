package wasmvm

import (
	"github.com/ontio/ontology-test/testframework"
	"time"
	"fmt"
	"encoding/binary"
	"github.com/ontio/ontology/common"
)

func TestOEP5C(ctx *testframework.TestFrameworkContext) bool {

	testFile := filePath + "/" + "COEP5.wasm"
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

	ctx.LogInfo("=====================invoke totalSupply==============================")
	res, err := PreExecWasmContract(ctx,
		addr,
		"totalSupply",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err := res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("totalSupply  is %d\n",  binary.LittleEndian.Uint64(bs))

	ctx.LogInfo("=====================invoke totalSupply end==============================")

	ctx.LogInfo("=====================invoke createToken==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"createToken",
		[]interface{}{"testTk1","http://someimage.com","this is a test token","on other"})

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
	ctx.LogInfo("=====================invoke createToken end==============================")

	ctx.LogInfo("=====================invoke getTokenByID==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"getTokenByID",
		[]interface{}{uint64(1)})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()


	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)

	fmt.Printf("getTokenByID  is %v\n",  bs)

	ctx.LogInfo("=====================invoke getTokenByID end==============================")
	ctx.LogInfo("=====================invoke ownerOf==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"ownerOf",
		[]interface{}{uint64(1)})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()


	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)
	tmpaddr ,_ := common.AddressParseFromBytes(bs)
	fmt.Printf("owner   is %s\n",  tmpaddr.ToBase58())

	ctx.LogInfo("=====================invoke ownerOf end==============================")

	ctx.LogInfo("=====================invoke transfer==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"transfer",
		[]interface{}{account2.Address,uint64(1)})

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
	ctx.LogInfo("=====================invoke createToken end==============================")
	ctx.LogInfo("=====================invoke ownerOf==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"ownerOf",
		[]interface{}{uint64(1)})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err = res.Result.ToByteArray()


	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)
	tmpaddr ,_ = common.AddressParseFromBytes(bs)
	fmt.Printf("owner   is %s\n",  tmpaddr.ToBase58())

	ctx.LogInfo("=====================invoke ownerOf end==============================")
	return true
}