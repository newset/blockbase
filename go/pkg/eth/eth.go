package wallet

type EthereumClient struct {
	Id   int
	Name string
	Url  string
}

func (c *EthereumClient) GenerateAccount() (string, string) {
	// 实现以太坊账户生成逻辑
	return "", ""
}
