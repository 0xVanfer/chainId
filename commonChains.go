// @Description Package chainId is based on chainlist https://chainlist.org.
package chainId

// Commonly used chains.
const (
	ArbitrumChainName  string = "arbitrum"
	AvalancheChainName string = "avalanche"
	BNBSmartChainName  string = "BSC"
	EthereumChainName  string = "ethereum"
	FantomChainName    string = "fantom"
	GoerliChainName    string = "goerli"
	HecoChainName      string = "heco"
	OkChainName        string = "okexchain"
	OptimismChainName  string = "optimism"
	PolygonChainName   string = "polygon"
	TerraChainName     string = "terra"
	XDaiChainName      string = "xdai"
	FilecoinChainName  string = "filecoin"
)

// Map chain name to chain id.
var ChainName2IdMap = map[string]int{
	EthereumChainName:  1,
	OptimismChainName:  10,
	BNBSmartChainName:  56,
	OkChainName:        66,
	XDaiChainName:      100,
	HecoChainName:      128,
	PolygonChainName:   137,
	FantomChainName:    250,
	FilecoinChainName:  314,
	ArbitrumChainName:  42161,
	AvalancheChainName: 43114,
}

// Map chain id to chain name.
var ChainId2NameMap = map[int]string{
	1:     EthereumChainName,
	10:    OptimismChainName,
	56:    BNBSmartChainName,
	66:    OkChainName,
	100:   XDaiChainName,
	128:   HecoChainName,
	137:   PolygonChainName,
	250:   FantomChainName,
	314:   FilecoinChainName,
	42161: ArbitrumChainName,
	43114: AvalancheChainName,
}

// Map chain name to native token symbol.
var NativeTokenSymbolList = map[string]string{
	EthereumChainName:  "eth",
	OptimismChainName:  "eth",
	BNBSmartChainName:  "bnb",
	OkChainName:        "okt",
	XDaiChainName:      "xdai",
	HecoChainName:      "ht",
	PolygonChainName:   "matic",
	FantomChainName:    "ftm",
	FilecoinChainName:  "fil",
	ArbitrumChainName:  "eth",
	AvalancheChainName: "avax",
}
