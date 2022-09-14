package chainid

// Commonly used chains.
const ArbitrumChainName string = "arbitrum"
const AvalancheChainName string = "avalanche"
const BinanaceSmartChainName string = "binance"
const EthereumChainName string = "ethereum"
const FantomChainName string = "fantom"
const GoerliChainName string = "goerli"
const HecoChainName string = "heco"
const OkChainName string = "okexchain"
const OptimismChainName string = "optimism"
const PolygonChainName string = "polygon"
const TerraChainName string = "terra"
const XDaiChainName string = "xdai"

// Map chain name to chain id.
var ChainName2IdMap = map[string]int{
	EthereumChainName:      1,
	OptimismChainName:      10,
	BinanaceSmartChainName: 56,
	OkChainName:            66,
	XDaiChainName:          100,
	HecoChainName:          128,
	PolygonChainName:       137,
	FantomChainName:        250,
	ArbitrumChainName:      42161,
	AvalancheChainName:     43114,
}

// Map chain id to chain name.
var ChainId2NameMap = map[int]string{
	1:     EthereumChainName,
	10:    OptimismChainName,
	56:    BinanaceSmartChainName,
	66:    OkChainName,
	100:   XDaiChainName,
	128:   HecoChainName,
	137:   PolygonChainName,
	250:   FantomChainName,
	42161: ArbitrumChainName,
	43114: AvalancheChainName,
}

// Map chain name to chain token symbol.
var ChainTokenSymbolList = map[string]string{
	EthereumChainName:      "eth",
	OptimismChainName:      "eth",
	BinanaceSmartChainName: "bnb",
	OkChainName:            "okt",
	XDaiChainName:          "xdai",
	HecoChainName:          "ht",
	PolygonChainName:       "matic",
	FantomChainName:        "ftm",
	ArbitrumChainName:      "eth",
	AvalancheChainName:     "avax",
}
