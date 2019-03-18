package wasmvm

import (
	"github.com/ontio/ontology-test/testframework"
	"time"
	"fmt"
	"github.com/ontio/ontology/common"
	"io/ioutil"
	"github.com/ontio/ontology-go-sdk/utils"
	"encoding/binary"
)

func TestCWasmAPI(ctx *testframework.TestFrameworkContext) bool {

	testFile := filePath + "/" + "cwasmAPI.wasm"
	signer, _ := ctx.GetDefaultAccount()
	timeoutSec := 30 * time.Second
	fmt.Println(timeoutSec)
	_, addr, err := DeployWasmJsonContract(ctx, signer, testFile, "testContract", "1")
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
	fmt.Printf("account2:%v\n",account2.Address[:])

	ctx.LogInfo("=====================invoke timestamp ==============================")
	res, err := PreExecWasmContract(ctx,
		addr,
		"tsp",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err := res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	fmt.Printf("timestamp is %d\n",  binary.LittleEndian.Uint64(bs))

	ctx.LogInfo("=====================invoke timestamp end==============================")

	ctx.LogInfo("=====================invoke blkheight ==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"blkheight",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, _ = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	fmt.Printf("blkheight is %d\n",  binary.LittleEndian.Uint32(bs))

	ctx.LogInfo("=====================invoke blkheight end==============================")


	ctx.LogInfo("=====================invoke selfAddr ==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"selfaddr",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, _ = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	selfaddr,_ := common.AddressParseFromBytes(bs)

	fmt.Printf("selfaddr is %s\n",  selfaddr.ToBase58())

	ctx.LogInfo("=====================invoke selfaddr end==============================")

	ctx.LogInfo("=====================invoke calleraddr ==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"calleraddr",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, _ = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	calleraddr,_ := common.AddressParseFromBytes(bs)

	fmt.Printf("calleraddr is %s\n",  calleraddr.ToBase58())

	ctx.LogInfo("=====================invoke calleraddr end==============================")

	ctx.LogInfo("=====================invoke entryaddr ==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"entryaddr",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, _ = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	entryaddr,_ := common.AddressParseFromBytes(bs)

	fmt.Printf("entryaddr is %s\n",  entryaddr.ToBase58())

	ctx.LogInfo("=====================invoke entryaddr end==============================")

	ctx.LogInfo("=====================invoke crtblkhash ==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"crtblkhash",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, _ = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	s := fmt.Sprintf("%x",bs)

	fmt.Printf("crtblkhash is %s\n",  s)

	ctx.LogInfo("=====================invoke entryaddr end==============================")

	ctx.LogInfo("=====================invoke crttxhash ==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"crttxhash",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, _ = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	s = fmt.Sprintf("%x",bs)

	fmt.Printf("crttxhash is %s\n",  s)

	ctx.LogInfo("=====================invoke crttxhash end==============================")
	ctx.LogInfo("=====================invoke storagewrite==============================")
	txhash, err := InvokeWasmContract(ctx,
		signer,
		addr,
		"storagewrite",
		[]interface{}{"testkey",signer.Address})

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
	if events != nil{
		if events.State == 0 {
			ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
			return false
		}
		fmt.Printf("events.Notify:%v", events.Notify)
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
	}
	ctx.LogInfo("=====================storagewrite  end==============================")

	ctx.LogInfo("=====================invoke storageread ==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"storageread",
		[]interface{}{"testkey"})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, _ = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)
	tmpaddr,_:= common.AddressParseFromBytes(bs)

	fmt.Printf("storageread is %s\n",  tmpaddr.ToBase58())

	ctx.LogInfo("=====================invoke storageread end==============================")

	ctx.LogInfo("=====================invoke storagedel==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"storagedel",
		[]interface{}{"testkey"})

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
	if events != nil{
		if events.State == 0 {
			ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
			return false
		}
		fmt.Printf("events.Notify:%v", events.Notify)
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
	}
	ctx.LogInfo("=====================storagedel  end==============================")
	ctx.LogInfo("=====================invoke storageread ==============================")
	res, err = PreExecWasmContract(ctx,
		addr,
		"storageread",
		[]interface{}{"testkey"})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, _ = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)
	tmpaddr,_= common.AddressParseFromBytes(bs)

	fmt.Printf("storageread is %s\n",  tmpaddr.ToBase58())

	ctx.LogInfo("=====================invoke storageread end==============================")

	ctx.LogInfo("=====================invoke ntf==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"ntf",
		[]interface{}{"test notify"})

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
	if events != nil{
		if events.State == 0 {
			ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
			return false
		}
		fmt.Printf("events.Notify:%v\n", events.Notify)
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v\n", notify)
			fmt.Printf(" states: %v\n",notify.States)
			fmt.Printf(" states: %v\n",notify.States.(string))
			fmt.Printf(" states: %v\n",[]byte(notify.States.(string)))
		}
	}
	ctx.LogInfo("=====================ntf  end==============================")

	ctx.LogInfo("=====================invoke storagewrite==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"storagewrite",
		[]interface{}{"testkey",signer.Address})

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
	if events != nil{
		if events.State == 0 {
			ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
			return false
		}
		fmt.Printf("events.Notify:%v", events.Notify)
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v", notify)
		}
	}
	ctx.LogInfo("=====================storagewrite  end==============================")
	ctx.LogInfo("=====================invoke crtmigrate==============================")

	migcode,err := ioutil.ReadFile(filePath + "/" + "cwasmAPI2.wasm")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	migcodeHash := common.ToHexString(migcode)
	migaddr,_:= utils.GetContractAddress(migcodeHash)

	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"crtmigrate",
		[]interface{}{migcode,3,"test","2.0","test","test@email.com","desc"})

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
	if events != nil{
		if events.State == 0 {
			ctx.LogError("TestWasmOEP4 failed invoked exec state return 0")
			return false
		}
		fmt.Printf("events.Notify:%v\n", events.Notify)
		for _, notify := range events.Notify {
			ctx.LogInfo("%+v\n", notify)
			fmt.Printf(" states: %v\n",notify.States)
			fmt.Printf(" states: %v\n",notify.States.(string))
			fmt.Printf(" states: %v\n",[]byte(notify.States.(string)))
		}
	}
	ctx.LogInfo("=====================crtmigrate  end==============================")
	//ctx.LogInfo("=====================invoke storageread ==============================")
	//res, err = PreExecWasmContract(ctx,
	//	addr,
	//	"storageread",
	//	[]interface{}{"testkey"})
	//
	//if err != nil {
	//	fmt.Printf("invoke name failed:%s\n", err.Error())
	//	return false
	//}
	//
	//bs, _ = res.Result.ToByteArray()
	//
	//fmt.Printf("res is %v\n", bs)
	//tmpaddr,_= common.AddressParseFromBytes(bs)
	//
	//fmt.Printf("storageread is %s\n",  tmpaddr.ToBase58())
	//
	//ctx.LogInfo("=====================invoke storageread end==============================")
	ctx.LogInfo("=====================invoke storageread ==============================")
	res, err = PreExecWasmContract(ctx,
		migaddr,
		"storageread",
		[]interface{}{"testkey"})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, _ = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)
	tmpaddr,_ = common.AddressParseFromBytes(bs)

	fmt.Printf("storageread is %s\n",  tmpaddr.ToBase58())

	ctx.LogInfo("=====================invoke storageread end==============================")
	return true

}