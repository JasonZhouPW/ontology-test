package wasmvm

import (
	"fmt"
	"github.com/ontio/ontology-test/testframework"
	"bytes"
	"github.com/ontio/ontology/common/serialization"
	"time"
	"encoding/binary"
	"github.com/ontio/ontology/common"
)

func TestRustOEP4Performance(ctx *testframework.TestFrameworkContext) bool {
	testFile := filePath + "/" + "rustOEP4.wasm"
	signer, _ := ctx.GetDefaultAccount()
	timeoutSec := 30 * time.Second
	txhash, addr, err := DeployWasmJsonContract(ctx, signer, testFile, "testContract", "1")
	if err != nil {
		fmt.Printf("deploy failed:%s\n", err.Error())
		return false
	}
	fmt.Println(txhash)
	ctx.LogInfo("==contract address is %s", addr.ToBase58())
	//account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil {
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}
	ctx.LogInfo("=====================invoke name==============================")
	res, err := PreExecWasmContract(ctx,
		addr,
		"name",
		[]interface{}{})

	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}

	bs, err := res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	tmp, err := serialization.ReadString(bytes.NewBuffer(bs))
	fmt.Printf("return is %v\n", tmp)

	ctx.LogInfo("=====================invoke name end==============================")

	ctx.LogInfo("=====================invoke init==============================")
	txhash, err = InvokeWasmContract(ctx,
		signer,
		addr,
		"initialize",
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
	account2, err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	prestart := time.Now().UnixNano()

	res, err = PreExecWasmContract(ctx,
		addr,
		"balance_of",
		[]interface{}{account2.Address})
	fmt.Printf("preexe costs:%d\n",time.Now().UnixNano() - prestart)
	if err != nil {
		fmt.Printf("invoke name failed:%s\n", err.Error())
		return false
	}
	bs, err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n", bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)
	balance, err := common.Uint256ParseFromBytes(bs)
	if err != nil {
		fmt.Printf("error is %s\n", err.Error())
	}
	oldbalance := int64(binary.LittleEndian.Uint64(balance[:8]))
	fmt.Printf("balance of %s is %d\n", signer.Address.ToBase58(), oldbalance)



	ctx.LogInfo("=====================invoke balanceOf end==============================")

	cnt := 10
	current := time.Now().Unix()
	fmt.Printf("start time is :%d\n",current)

	amountbs := make([]byte, 16)

	binary.LittleEndian.PutUint64(amountbs, uint64(1))
	parambs := make([]byte, 32)
	copy(parambs[:16], amountbs)
	fmt.Printf("parambs %v\n", parambs)

	amount, _ := common.Uint256ParseFromBytes(parambs)

	currenttime := time.Now().Unix()
	fmt.Printf("start....:%d\n",currenttime)

	for i:= 0;i < cnt;i++ {
		ctx.LogInfo("===========wasm transfer start")

		sendone := time.Now().UnixNano()
		txhash, err = InvokeWasmContract(ctx,
			signer,
			addr,
			"transfer",
			[]interface{}{signer.Address, account2.Address, amount})
		fmt.Printf("one time send costs:%d\n",time.Now().UnixNano() - sendone)
		ctx.LogInfo("===========wasm transfer end")

		if err != nil {
			ctx.LogError("TestOEP4Py InvokeNeoVMSmartContract error: %s", err)
			return false
		}
	}

	for{
		res, err = PreExecWasmContract(ctx,
			addr,
			"balance_of",
			[]interface{}{account2.Address})

		if err != nil {
			fmt.Printf("invoke name failed:%s\n", err.Error())
			return false
		}

		bs, err = res.Result.ToByteArray()
		//fmt.Printf("bs is %v\n",bs)
		//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
		//balance := binary.LittleEndian.Uint32(bs)
		balance, err := common.Uint256ParseFromBytes(bs)
		if err != nil {
			fmt.Printf("error is %s\n", err.Error())
		}
		baln := int64(binary.LittleEndian.Uint64(balance[:8]))
		if baln == oldbalance + int64(cnt){
			fmt.Printf("time cost:%d\n",time.Now().Unix() - currenttime)
			return true
		}

	}

	return true
}
