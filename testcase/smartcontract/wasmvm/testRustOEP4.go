package wasmvm

import (
	"time"
	"fmt"
	"github.com/ontio/ontology-test/testframework"
	"bytes"
	"github.com/ontio/ontology/common/serialization"
	"github.com/ontio/ontology/common"
	"encoding/binary"
)

func TestRustOEP4(ctx *testframework.TestFrameworkContext) bool {
	testFile := filePath + "/" + "rustOEP4.wasm"
	signer, _ := ctx.GetDefaultAccount()
	timeoutSec := 30 * time.Second
	txhash, addr, err := DeployWasmJsonContract(ctx, signer, testFile, "testContract", "1")
	if err != nil {
		fmt.Printf("deploy failed:%s\n", err.Error())
		return false
	}
	account2,err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")
	if err != nil{
		ctx.LogError("get account AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb failed")
		return false
	}
	ctx.LogInfo("=====================invoke name==============================")
	res,err := PreExecWasmContract(ctx,
		addr,
		"name",
		byte(1),[]interface{}{})

	if err != nil{
		fmt.Printf("invoke name failed:%s\n",err.Error())
		return false
	}

	bs,err := res.Result.ToByteArray()

	fmt.Printf("res is %v\n",bs)

	tmp ,err:= serialization.ReadString(bytes.NewBuffer(bs))
	fmt.Printf("return is %v\n",tmp)

	ctx.LogInfo("=====================invoke name end==============================")



	ctx.LogInfo("=====================invoke init==============================")
	txhash,err = InvokeWasmContract(ctx,
		signer,
		addr,
		"initialize",
		byte(1),[]interface{}{signer.Address})

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
	fmt.Printf("events.Notify:%v",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke init end==============================")

	ctx.LogInfo("=====================invoke totalSupply==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"total_supply",
		byte(1),[]interface{}{})

	if err != nil{
		fmt.Printf("invoke name failed:%s\n",err.Error())
		return false
	}

	bs,err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n",bs)

	total,err := common.Uint256ParseFromBytes(bs)
	if err != nil{
		fmt.Printf("error is %s\n",err.Error())
	}
	fmt.Printf("totalSupply is %d\n",binary.LittleEndian.Uint64(total[:8]))


	ctx.LogInfo("=====================invoke totalSupply end==============================")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"balance_of",
		byte(1),[]interface{}{signer.Address})

	if err != nil{
		fmt.Printf("invoke name failed:%s\n",err.Error())
		return false
	}

	bs,err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n",bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)
	balance,err := common.Uint256ParseFromBytes(bs)
	if err != nil{
		fmt.Printf("error is %s\n",err.Error())
	}
	fmt.Printf("balance of %s is %d\n",signer.Address.ToBase58(),binary.LittleEndian.Uint64(balance[:8]))

	ctx.LogInfo("=====================invoke balanceOf end==============================")


	ctx.LogInfo("=====================invoke transfer==============================")
	amountbs := make([]byte,16)

	binary.LittleEndian.PutUint64(amountbs,uint64(500))
	parambs := make([]byte,32)
	copy(parambs[:16],amountbs)
	fmt.Printf("parambs %v\n",parambs)

	amount,_ := common.Uint256ParseFromBytes(parambs)

	txhash,err = InvokeWasmContract(ctx,
		signer,
		addr,
		"transfer",
		byte(1),[]interface{}{signer.Address,account2.Address,amount})

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
	fmt.Printf("events.Notify:%v\n",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transfer end==============================")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"balance_of",
		byte(1),[]interface{}{signer.Address})

	if err != nil{
		fmt.Printf("invoke name failed:%s\n",err.Error())
		return false
	}

	bs,err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n",bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)
	balance,err = common.Uint256ParseFromBytes(bs)
	if err != nil{
		fmt.Printf("error is %s\n",err.Error())
	}
	fmt.Printf("balance of %s is %d\n",signer.Address.ToBase58(),binary.LittleEndian.Uint64(balance[:8]))

	ctx.LogInfo("=====================invoke balanceOf end==============================")

	ctx.LogInfo("=====================invoke balanceOf==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"balance_of",
		byte(1),[]interface{}{account2.Address})

	if err != nil{
		fmt.Printf("invoke name failed:%s\n",err.Error())
		return false
	}

	bs,err = res.Result.ToByteArray()

	fmt.Printf("res is %v\n",bs)

	//tmp ,err = serialization.ReadString(bytes.NewBuffer(bs))
	//balance := binary.LittleEndian.Uint32(bs)
	balance,err = common.Uint256ParseFromBytes(bs)
	if err != nil{
		fmt.Printf("error is %s\n",err.Error())
	}
	fmt.Printf("balance of %s is %d\n",account2.Address.ToBase58(),binary.LittleEndian.Uint64(balance[:8]))

	ctx.LogInfo("=====================invoke balanceOf end==============================")

	return true
}