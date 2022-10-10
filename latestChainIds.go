package chainId

import (
	"strconv"
	"strings"

	"github.com/imroc/req"
)

// Chain ids given by defillama/chainlist.
//
// ChainList: https://chainlist.org.
func GetLatestChainIdsMap() (Id2NameMap map[int]string, Name2IdMap map[string]int, err error) {
	url := `https://raw.githubusercontent.com/DefiLlama/chainlist/main/constants/chainIds.js`
	r, err := req.Get(url)
	if err != nil {
		return
	}
	resStr := r.String()
	// find "{"
	startIndex := strings.Index(resStr, "{")
	jsonStr := resStr[startIndex:]
	// find "}"
	endIndex := strings.Index(jsonStr, "}")
	// get the string info between "{" and "}"
	jsonStr = jsonStr[:endIndex+1]
	// to clean string
	jsonStr = strings.ReplaceAll(jsonStr, " ", "")
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	jsonStr = jsonStr[1:]
	// into map
	chainId2NameMap := make(map[int]string)
	chainName2IdMap := make(map[string]int)
	for {
		// key:   between `",` and `:"`
		// value: between `:"` and `",`(or `:}`)
		keyEndIndex := strings.Index(jsonStr, `:"`)
		// does not exist
		if keyEndIndex == -1 {
			break
		}
		key := jsonStr[:keyEndIndex]
		jsonStr = jsonStr[keyEndIndex+2:]
		valueEndIndex := strings.Index(jsonStr, `",`)
		if valueEndIndex == -1 {
			valueEndIndex = strings.Index(jsonStr, `"}`)
			if valueEndIndex == -1 {
				break
			}
		}
		value := jsonStr[:valueEndIndex]
		jsonStr = jsonStr[valueEndIndex+2:]

		chainId, err := strconv.Atoi(key)
		if err != nil {
			return nil, nil, err
		}
		chainId2NameMap[chainId] = value
		chainName2IdMap[value] = chainId
	}
	return chainId2NameMap, chainName2IdMap, nil
}
