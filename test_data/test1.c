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

char* invoke(char * method,char * args){

    if (strcmp(method ,"init")==0 ){
            return "init success!";
    }

    if (strcmp(method, "add")==0 ){

      // char* str1[]={"hello","world","test char array"};

      struct Param{
            int a;
            int b;
        };

        struct Param *p = (struct Param *)malloc(sizeof(struct Param));
        ONT_JsonUnmashalInput(p,sizeof(struct Param),args);

      char ** str = (char **)malloc(sizeof(char*) * 2);
      str[0] = "result is ";
      str[1] = Itoa( p->a + p->b);


      ONT_Runtime_Notify(str);
      return  Itoa(p->a + p->b);
    }


    if (strcmp(method, "put") == 0)
    {
      struct Param
      {
        char * key;
        char * value;
      };
        struct Param *p = (struct Param *)malloc(sizeof(struct Param));
        ONT_JsonUnmashalInput(p,sizeof(struct Param),args);

        ONT_Storage_Put(p->key, p->value);
        return "true";

    }


    if (strcmp(method, "get")==0)
    {
       struct Param
      {
        char * key;
      };
        struct Param *p = (struct Param *)malloc(sizeof(struct Param));
        ONT_JsonUnmashalInput(p,sizeof(struct Param),args);

        char * value = ONT_Storage_Get(p->key);

        char ** str = (char **)malloc(sizeof(char*) );

        str[0] = value;
        ONT_Runtime_Notify(str);


        return value;
    }




    return "false";
}