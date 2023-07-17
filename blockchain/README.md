블록체인하는 기법?
다른 블록을 만들때 기존의 블록을 이용해서 만든다?
B1

	b1Hash = (data +"")

B2

	b2Hash = (data +b1Hash)

B3

	b3Hash = (data +b2Hash)

B1 의 b1Hash를 b1Hash1로 변경하면
B1

	b1Hash1 = (data+"")

B2

	b2Hash = (data+b1Hash) 였던 것에서 b1Hash라는게 없어져서 (data + b1Hash1)로 바뀌어야 한다.

B3

	b3Hash = (data+b2Hast) b2Hash는 바뀌지 않았기에 기존이랑 똑같다.
//genesisBlock := block{"Genesis Block", "", ""}
// sha256.Sum256은 []byte slice 형식을 받는다.
//string을 []byte로 만들고 해싱을 한다
//hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
// [size]byte를 다시 string으로 만들어서 저장한다.
//HexHash := fmt.Sprintf("%x", hash)
// 해당 string을 hash에 저장해주기
//genesisBlock.hash = HexHash
//fmt.Println(genesisBlock)
