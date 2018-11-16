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
char * ONT_CallContract(char * address,char * method,char * args);
char * ONT_NativeInvoke(int ver,char *address, char * method, char * args);
char * ONT_MarshalNativeParams(void * s);
// char * ONT_MarshalNeoParams(void * s);

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


struct Param{
        char * ptype;
        void * pvalue;
};


char * OEP4_addr = "AHeUVHURPAJ2DZijq35NrinL39cDnDYSTb";

char * transferOEP4(char * fromAddress, char * toAddress, long long amount){

    struct Param * args = (struct Param *)malloc(sizeof(struct Param) * 3);
    args[0].ptype = "string";
    args[0].pvalue = fromAddress;

    args[1].ptype = "string";
    args[1].pvalue = toAddress;

    args[2].ptype = "int64";
    args[2].pvalue = amount;

    return ONT_CallContract(OEP4_addr,"transfer",ONT_RawMashalParams(args));
}

char * balanceOf(char * address){
    struct Param * args = (struct Param *)malloc(sizeof(struct Param));
    args[0].ptype = "string";
    args[0].pvalue = address;

    return ONT_CallContract(OEP4_addr,"balanceOf",ONT_RawMashalParams(args));

}


char* invoke(char * method,char * args){

    if (strcmp(method ,"transferOEP4")==0 )
    {
        char * fromAddr = ONT_ReadStringParam(args);
        char * toAddr = ONT_ReadStringParam(args);
        long long amount = ONT_ReadInt64Param(args);

        return transferOEP4(fromAddr, toAddr, amount);
    }
    if (strcmp(method, "balanceOfOEP4")==0)
    {
        char * fromAddr = ONT_ReadStringParam(args);

        return balanceOf(fromAddr);
    }
    return "false";
}