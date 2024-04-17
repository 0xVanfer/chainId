// @Description Package chainId is based on chainlist https://chainlist.org.
package chainId

// Commonly used chains.
const (
	ArbitrumChainName     string = "arbitrum"  // Arbitrum one.
	AvalancheChainName    string = "avalanche" // Avalanche C-Chain.
	BaseChainName         string = "base"
	BlastChainName        string = "blast"
	BNBSmartChainName     string = "bsc"
	CeloChainName         string = "celo"
	CentrifugeChainName   string = "centrifuge"
	EthereumChainName     string = "ethereum"
	FantomChainName       string = "fantom"
	FilecoinChainName     string = "filecoin"
	FraxtalChainName      string = "fraxtal"
	GnosisChainName       string = "gnosis"
	GoerliChainName       string = "goerli"
	HecoChainName         string = "heco"
	ImmutableChainName    string = "immutable"
	LineaChainName        string = "linea"
	KavaChainName         string = "kava"
	MantleChainName       string = "mantle"
	MoonbeamChainName     string = "moonbeam"
	OkChainName           string = "okexchain"
	OptimismChainName     string = "optimism"
	PolygonChainName      string = "polygon"
	PolygonZkEVMChainName string = "polygonzkevm"
	SepoliaChainName      string = "sepolia"
	ScrollChainName       string = "scroll"
	TerraChainName        string = "terra"
	// Deprecated: Use GnosisChainName instead.
	XDaiChainName string = "xdai"
)

// Map chain name to chain id.
var ChainName2IdMap = map[string]int{
	EthereumChainName:     1,
	GoerliChainName:       5,
	OptimismChainName:     10,
	BNBSmartChainName:     56,
	OkChainName:           66,
	XDaiChainName:         100, // Deprecated: Use GnosisChainName instead.
	GnosisChainName:       100,
	HecoChainName:         128,
	PolygonChainName:      137,
	BlastChainName:        238,
	FantomChainName:       250,
	FraxtalChainName:      252,
	FilecoinChainName:     314,
	PolygonZkEVMChainName: 1101,
	MoonbeamChainName:     1284,
	CentrifugeChainName:   2031,
	KavaChainName:         2222,
	MantleChainName:       5000,
	BaseChainName:         8453,
	ImmutableChainName:    13371,
	ArbitrumChainName:     42161,
	CeloChainName:         42220,
	AvalancheChainName:    43114,
	LineaChainName:        59144,
	ScrollChainName:       534352,
	SepoliaChainName:      11155111,
}

// Map chain id to chain name.
var ChainId2NameMap = map[int]string{
	ChainName2IdMap[EthereumChainName]:     EthereumChainName,     // 1
	ChainName2IdMap[GoerliChainName]:       GoerliChainName,       // 5
	ChainName2IdMap[OptimismChainName]:     OptimismChainName,     // 10
	ChainName2IdMap[BNBSmartChainName]:     BNBSmartChainName,     // 56
	ChainName2IdMap[OkChainName]:           OkChainName,           // 66
	ChainName2IdMap[GnosisChainName]:       GnosisChainName,       // 100
	ChainName2IdMap[HecoChainName]:         HecoChainName,         // 128
	ChainName2IdMap[PolygonChainName]:      PolygonChainName,      // 137
	ChainName2IdMap[BlastChainName]:        BlastChainName,        // 238
	ChainName2IdMap[FantomChainName]:       FantomChainName,       // 250
	ChainName2IdMap[FraxtalChainName]:      FraxtalChainName,      // 252
	ChainName2IdMap[FilecoinChainName]:     FilecoinChainName,     // 314
	ChainName2IdMap[PolygonZkEVMChainName]: PolygonZkEVMChainName, // 1101
	ChainName2IdMap[MoonbeamChainName]:     MoonbeamChainName,     // 1284
	ChainName2IdMap[CentrifugeChainName]:   CentrifugeChainName,   // 2031
	ChainName2IdMap[KavaChainName]:         KavaChainName,         // 2222
	ChainName2IdMap[MantleChainName]:       MantleChainName,       // 5000
	ChainName2IdMap[BaseChainName]:         BaseChainName,         // 8453
	ChainName2IdMap[ImmutableChainName]:    ImmutableChainName,    // 13371
	ChainName2IdMap[ArbitrumChainName]:     ArbitrumChainName,     // 42161
	ChainName2IdMap[CeloChainName]:         CeloChainName,         // 42220
	ChainName2IdMap[AvalancheChainName]:    AvalancheChainName,    // 43114
	ChainName2IdMap[LineaChainName]:        LineaChainName,        // 59144
	ChainName2IdMap[ScrollChainName]:       ScrollChainName,       // 534352
	ChainName2IdMap[SepoliaChainName]:      SepoliaChainName,      // 11155111
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
	LineaChainName:        "eth",
	PolygonZkEVMChainName: "eth",
	SepoliaChainName:      "eth",
	ScrollChainName:       "eth",

	// Layer 1.
	AvalancheChainName:  "avax",
	BNBSmartChainName:   "bnb",
	CeloChainName:       "celo",
	CentrifugeChainName: "cfg",
	FantomChainName:     "ftm",
	FilecoinChainName:   "fil",
	FraxtalChainName:    "frxeth",
	GnosisChainName:     "xdai",
	HecoChainName:       "ht",
	ImmutableChainName:  "imx",
	KavaChainName:       "kava",
	MantleChainName:     "mnt",
	MoonbeamChainName:   "glmr",
	OkChainName:         "okt",
	PolygonChainName:    "matic", // Might be changed to pol someday.
	XDaiChainName:       "xdai",
}
