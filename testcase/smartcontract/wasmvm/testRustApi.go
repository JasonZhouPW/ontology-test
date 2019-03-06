package wasmvm

import (
	"fmt"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"time"
	"encoding/binary"
)

func TestRustApi(ctx *testframework.TestFrameworkContext) bool {
	testFile := filePath + "/" + "rustapi.wasm"
	signer, _ := ctx.GetDefaultAccount()
	timeoutSec := 30 * time.Second
	_, addr, err := DeployWasmJsonContract(ctx, signer, testFile, "testContract", "1")
	if err != nil {
		fmt.Printf("deploy failed:%s\n", err.Error())
		return false
	}

	ctx.LogInfo("======contract address :%s",addr.ToHexString())
	ctx.LogInfo("======contract address :%s",addr.ToBase58())
/*	ctx.LogInfo("=====================invoke timestamp==============================")
	res,err := PreExecWasmContract(ctx,
		addr,
		"timestamp",
		[]interface{}{})

	if err != nil{
		fmt.Printf("invoke timestamp failed:%s\n",err.Error())
		return false
	}

	bs,err := res.Result.ToByteArray()

	fmt.Printf("timestamp is %v\n",bs)

	tmp ,err:= serialization.ReadUint64(bytes.NewBuffer(bs))
	fmt.Printf("return is %v\n",tmp)

	ctx.LogInfo("=====================invoke timestamp end==============================")

	ctx.LogInfo("=====================invoke blockheight==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"blockheight",
		[]interface{}{})

	if err != nil{
		fmt.Printf("invoke blockheight failed:%s\n",err.Error())
		return false
	}

	bs,err = res.Result.ToByteArray()

	fmt.Printf("blockheight is %v\n",bs)

	tmp2 ,err:= serialization.ReadUint32(bytes.NewBuffer(bs))
	fmt.Printf("return is %v\n",tmp2)

	ctx.LogInfo("=====================invoke blockheight end==============================")

	ctx.LogInfo("=====================invoke calleraddress==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"calleraddress",
		[]interface{}{})

	if err != nil{
		fmt.Printf("invoke blockheight failed:%s\n",err.Error())
		return false
	}

	bs,err = res.Result.ToByteArray()

	fmt.Printf("selfaddress is %v\n",bs)

	calleraddress ,err:= serialization.ReadVarBytes(bytes.NewBuffer(bs))
	tmpaddr,_ :=common.AddressParseFromBytes(calleraddress)
	fmt.Printf("return is %v\n",tmpaddr.ToBase58())

	ctx.LogInfo("=====================invoke calleraddress end==============================")


	ctx.LogInfo("=====================invoke selfaddress==============================")
	res,err = PreExecWasmContract(ctx,
		addr,
		"selfaddress",
		[]interface{}{})

	if err != nil{
		fmt.Printf("invoke blockheight failed:%s\n",err.Error())
		return false
	}

	bs,err = res.Result.ToByteArray()

	fmt.Printf("selfaddress is %v\n",bs)

	selfaddress ,err:= serialization.ReadString(bytes.NewBuffer(bs))
	fmt.Printf("selfaddress2 is %v\n",selfaddress)

	tmpaddr,_ =common.AddressParseFromBytes([]byte(selfaddress))
	fmt.Printf("return is %v\n",tmpaddr.ToBase58())
	fmt.Printf("return is %v\n",tmpaddr.ToHexString())

	ctx.LogInfo("=====================invoke selfaddress end==============================")


	ctx.LogInfo("=====================invoke checkwitness==============================")

	txhash,err := InvokeWasmContract(ctx,
		signer,
		addr,
		"checkwitness",
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
	fmt.Printf("events.Notify:%v\n",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke transfer end==============================")
*/

	/*ctx.LogInfo("=====================invoke callcontract name==============================")

	oep4Addr ,_:= common.AddressFromBase58("AHJPesFTEoRkkeiKeoNmznH3ypd9VXHnbb")
	fmt.Printf("addr hex is %s\n",oep4Addr.ToHexString())
	fmt.Printf("addr bs is %v\n",oep4Addr[:])
	res,err := PreExecWasmContract(ctx,
		addr,
		"call_name",
		[]interface{}{oep4Addr})

	if err != nil{
		fmt.Printf("invoke call_name failed:%s\n",err.Error())
		return false
	}

	bs,err := res.Result.ToByteArray()

	fmt.Printf("selfaddress is %v\n",bs)

	name ,err:= serialization.ReadString(bytes.NewBuffer(bs))
	fmt.Printf("name is %v\n",name)



	ctx.LogInfo("=====================invoke selfaddress end==============================")
*/
	account2,err := ctx.GetAccount("AS3SCXw8GKTEeXpdwVw7EcC4rqSebFYpfb")

/*	ctx.LogInfo("=====================invoke native transfer==============================")
	amountbs := make([]byte,16)
	ontaddr,_ := common.AddressFromBase58("AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV")

	binary.LittleEndian.PutUint64(amountbs,uint64(500))
	parambs := make([]byte,32)
	copy(parambs[:16],amountbs)
	fmt.Printf("parambs %v\n",parambs)

	amount,_ := common.Uint256ParseFromBytes(parambs)

	txhash,err := InvokeWasmContract(ctx,
		signer,
		addr,
		"call_native_transfer",
		[]interface{}{ontaddr,byte(0),signer.Address,account2.Address,amount})

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
	fmt.Printf("events.Notify:%v\n",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}
	ctx.LogInfo("=====================invoke native transfer end==============================")*/


	ctx.LogInfo("=====================invoke neo contract name==============================")

	neoope4addr,_ := common.AddressFromBase58("AapCuqcnsUKdEK1CYHPFV4e7UgBwey3UED")

	fmt.Printf("addr hex is %s\n",neoope4addr.ToHexString())
	fmt.Printf("addr bs is %v\n",neoope4addr[:])

	amountbs := make([]byte,16)
	binary.LittleEndian.PutUint64(amountbs,uint64(500))
	parambs := make([]byte,32)
	copy(parambs[:16],amountbs)
	fmt.Printf("parambs %v\n",parambs)

	amount,_ := common.Uint256ParseFromBytes(parambs)

	txhash,err := InvokeWasmContract(ctx,
		signer,
		addr,
		"call_neovm_transfer",
		[]interface{}{neoope4addr,signer.Address,account2.Address,amount})

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
	fmt.Printf("events.Notify:%v\n",events.Notify)
	for _,notify:= range events.Notify{
		ctx.LogInfo("%+v", notify)
	}

	ctx.LogInfo("=====================invoke selfaddress end==============================")

	return  true
}