
//system apis
void * calloc(int count,int length);
void * malloc(int size);
int arrayLen(void *a);
int memcpy(void * dest,void * src,int length);
int memset(void * dest,char c,int length);

//utility apis
int strcmp(char *a,char *b);
char * strconcat(char *a,char *b);
int Atoi(char * s);
long long Atoi64(char *s);
char * Itoa(int a);
char * I64toa(long long amount,int radix);
char * SHA1(char *s);
char * SHA256(char *s);

//parameter apis
int ONT_ReadInt32Param(char *args);
long long ONT_ReadInt64Param(char * args);
char * ONT_ReadStringParam(char * args);
void ONT_JsonUnmashalInput(void * addr,int size,char * arg);
char * ONT_JsonMashalResult(void * val,char * types,int succeed);
char * ONT_JsonMashalParams(void * s);
char * ONT_RawMashalParams(void *s);
char * ONT_GetCallerAddress();
char * ONT_GetSelfAddress();
char * ONT_CallContract(char * address,char * contractCode,char * method,char * args);
char * ONT_MarshalNativeParams(void * s);
char * ONT_MarshalNeoParams(void * s);

//Runtime apis
int ONT_Runtime_CheckWitness(char * address);
void ONT_Runtime_Notify(char ** msg);
int ONT_Runtime_CheckSig(char * pubkey,char * data,char * sig);
int ONT_Runtime_GetTime();
void ONT_Runtime_Log(char * message);

//Attribute apis
int ONT_Attribute_GetUsage(char * data);
char * ONT_Attribute_GetData(char * data);

//Block apis
char * ONT_Block_GetCurrentHeaderHash();
int ONT_Block_GetCurrentHeaderHeight();
char * ONT_Block_GetCurrentBlockHash();
int ONT_Block_GetCurrentBlockHeight();
char * ONT_Block_GetTransactionByHash(char * hash);
int * ONT_Block_GetTransactionCountByBlkHash(char * hash);
int * ONT_Block_GetTransactionCountByBlkHeight(int height);
char ** ONT_Block_GetTransactionsByBlkHash(char * hash);
char ** ONT_Block_GetTransactionsByBlkHeight(int height);


//Blockchain apis
int ONT_BlockChain_GetHeight();
char * ONT_BlockChain_GetHeaderByHeight(int height);
char * ONT_BlockChain_GetHeaderByHash(char * hash);
char * ONT_BlockChain_GetBlockByHeight(int height);
char * ONT_BlockChain_GetBlockByHash(char * hash);
char * ONT_BlockChain_GetContract(char * address);

//header apis
char * ONT_Header_GetHash(char * data);
int ONT_Header_GetVersion(char * data);
char * ONT_Header_GetPrevHash(char * data);
char * ONT_Header_GetMerkleRoot(char  * data);
int ONT_Header_GetIndex(char * data);
int ONT_Header_GetTimestamp(char * data);
long long ONT_Header_GetConsensusData(char * data);
char * ONT_Header_GetNextConsensus(char * data);

//storage apis
void ONT_Storage_Put(char * key,char * value);
char * ONT_Storage_Get(char * key);
void ONT_Storage_Delete(char * key);

//transaction apis
char * ONT_Transaction_GetHash(char * data);
int ONT_Transaction_GetType(char * data);
char * ONT_Transaction_GetAttributes(char * data);

//for debug only
void ContractLogDebug(char * msg);
void ContractLogInfo(char * msg);
void ContractLogError(char * msg);

char * Admin = "Ad4pjz2bqep4RhQrUAzMuZJkBC3qJ1tZuT";
char * SUPPLY_KEY = "TOTAL_SUPPLY";
char * NAME = "WasmToken";
char * SYMBOL = "WTK";
char * BalanceKey = "B";
char * ApproveKey = "A";
int Decimals = 8;
long long Factor = 100000000;
long long TotalSupply = 1000000000;


char * _concatKey(char * str1, char* str2)
{
    return strconcat(strconcat(str1, "_"),str2);
}

char* init()
{
  char * sp = ONT_Storage_Get(SUPPLY_KEY);
  if(arrayLen(sp) > 0)
  {
    return "false";
  }
  else
  {
    char * value = I64toa( Factor * TotalSupply,10);
    ONT_Storage_Put(SUPPLY_KEY,value);
    ONT_Storage_Put(_concatKey(BalanceKey, Admin), value);
    return "true";
  }
}

void notifyTransfer( char * fromAddr, char * toAddr, long long amount)
{
      char ** str = (char **)malloc(sizeof(char*) * 4);
      str[0] = "transfer";
      str[1] = fromAddr;
      str[2] = toAddr;
      str[3] = I64toa(amount,10);

      ONT_Runtime_Notify(str);
}

char * name()
{
    return NAME;
}

char * symbol()
{
    return SYMBOL;
}

char * decimals()
{
    return Itoa(Decimals);
}

char * totalSupply()
{
    return ONT_Storage_Get(SUPPLY_KEY);
}

char * balanceOf(char * address)
{
    char * key = _concatKey(BalanceKey, address);
    return ONT_Storage_Get(key);
}

char * transfer(char * fromAddr, char * toAddr, long long amount)
{
    if (ONT_Runtime_CheckWitness(fromAddr) != 1)
    {
        return "false";
    }
    if (strcmp(fromAddr, toAddr)== 0)
    {
        return "false";
    }
    if (amount < 0)
    {
        return "false";
    }

    if (amount == 0)
    {
        return "true";
    }
    char * fromKey = _concatKey(BalanceKey, fromAddr);
    char * fromBalanceBytes = ONT_Storage_Get(fromKey);
    if (arrayLen(fromBalanceBytes) == 0)
    {
        return "false";
    }
    long long fromBalance = Atoi64(fromBalanceBytes);
    if (fromBalance < amount)
    {
        return "false";
    }

    char * toKey = _concatKey(BalanceKey, toAddr);

    long long toBalance = Atoi64(ONT_Storage_Get(toKey));
    ONT_Storage_Put(_concatKey(BalanceKey, toAddr), I64toa(toBalance + amount,10));

    if (fromBalance == amount)
    {
        ONT_Storage_Delete(fromKey);
    }else
    {
        ONT_Storage_Put(fromKey, I64toa(fromBalance - amount,10));
    }

    notifyTransfer(fromAddr, toAddr, amount);

    return "true";


}

char* invoke(char * method,char * args){

    if (strcmp(method ,"init")==0 )
    {
            return init();
    }
    if (strcmp(method, "name")==0)
    {
        return name();
    }
    if (strcmp(method, "symbol")==0)
    {
        return symbol();
    }
    if (strcmp(method, "decimals")==0)
    {
        return decimals();
    }
    if (strcmp(method, "totalSupply")==0)
    {
        return totalSupply();
    }
    if (strcmp(method, "balanceOf") == 0)
    {
        char * address = ONT_ReadStringParam(args);
        return balanceOf(address);
    }
    if (strcmp(method, "transfer") == 0)
    {
        char * fromAddr = ONT_ReadStringParam(args);
        char * toAddr = ONT_ReadStringParam(args);
        long long amount = ONT_ReadInt64Param(args);

        return transfer(fromAddr, toAddr,amount);
    }


    return "false";
}