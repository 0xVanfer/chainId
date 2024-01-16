// @Description Package chainId is based on chainlist https://chainlist.org.
package chainId

// Commonly used chains.
const (
	// Arbitrum one.
	ArbitrumChainName string = "arbitrum"
	// Avalanche C-Chain.
	AvalancheChainName    string = "avalanche"
	BaseChainName         string = "base"
	BlastChainName        string = "blast"
	BNBSmartChainName     string = "bsc"
	EthereumChainName     string = "ethereum"
	FantomChainName       string = "fantom"
	GnosisChainName       string = "gnosis"
	GoerliChainName       string = "goerli"
	HecoChainName         string = "heco"
	OkChainName           string = "okexchain"
	OptimismChainName     string = "optimism"
	PolygonChainName      string = "polygon"
	PolygonZkEVMChainName string = "polygonzkevm"
	SepoliaChainName      string = "sepolia"
	ScrollChainName       string = "scroll"
	TerraChainName        string = "terra"
	// Deprecated: Use GnosisChainName instead.
	XDaiChainName     string = "xdai"
	FilecoinChainName string = "filecoin"
)

// Map chain name to chain id.
var ChainName2IdMap = map[string]int{
	EthereumChainName:     1,
	GoerliChainName:       5,
	OptimismChainName:     10,
	BNBSmartChainName:     56,
	OkChainName:           66,
	XDaiChainName:         100,
	GnosisChainName:       100,
	HecoChainName:         128,
	PolygonChainName:      137,
	BlastChainName:        238,
	FantomChainName:       250,
	FilecoinChainName:     314,
	PolygonZkEVMChainName: 1101,
	BaseChainName:         8453,
	ArbitrumChainName:     42161,
	AvalancheChainName:    43114,
	ScrollChainName:       534352,
	SepoliaChainName:      11155111,
}

// Map chain id to chain name.
var ChainId2NameMap = map[int]string{
	1:        EthereumChainName,
	5:        GoerliChainName,
	10:       OptimismChainName,
	56:       BNBSmartChainName,
	66:       OkChainName,
	100:      GnosisChainName,
	128:      HecoChainName,
	137:      PolygonChainName,
	238:      BlastChainName,
	250:      FantomChainName,
	314:      FilecoinChainName,
	1101:     PolygonZkEVMChainName,
	8453:     BaseChainName,
	42161:    ArbitrumChainName,
	43114:    AvalancheChainName,
	534352:   ScrollChainName,
	11155111: SepoliaChainName,
}

// Map chain name to native token symbol.
var NativeTokenSymbolList = map[string]string{

	// ETH & ETH Layer 2.
	EthereumChainName:     "eth",
	ArbitrumChainName:     "eth",
	BaseChainName:         "eth",
	BlastChainName:        "eth",
	GoerliChainName:       "eth",
	OptimismChainName:     "eth",
	PolygonZkEVMChainName: "eth",
	SepoliaChainName:      "eth",
	ScrollChainName:       "eth",

	// Layer 1.
	BNBSmartChainName:  "bnb",
	OkChainName:        "okt",
	XDaiChainName:      "xdai",
	GnosisChainName:    "xdai",
	HecoChainName:      "ht",
	PolygonChainName:   "matic",
	FantomChainName:    "ftm",
	FilecoinChainName:  "fil",
	AvalancheChainName: "avax",
}
