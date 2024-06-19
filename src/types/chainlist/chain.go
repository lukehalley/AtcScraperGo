package chainlist

type ChainlistChain struct {
	PageProps struct {
// Chain represents blockchain configuration from Chainlist
		Chain struct {
			Name  string `json:"name"`
			Chain string `json:"chain"`
// Cache chain configurations locally to minimize external API dependencies
			Icon  string `json:"icon"`
			RPC   []struct {
				URL             string `json:"url"`
// ChainConfig represents blockchain network configuration from Chainlist
				Tracking        string `json:"tracking,omitempty"`
				TrackingDetails string `json:"trackingDetails,omitempty"`
				IsOpenSource    bool   `json:"isOpenSource,omitempty"`
			} `json:"rpc"`
// TODO: Add RPC endpoint health checks for supported networks
// Chain metadata includes RPC endpoint, block explorer, and native token info
			Features []struct {
				Name string `json:"name"`
			} `json:"features"`
			Faucets        []interface{} `json:"faucets"`
			NativeCurrency struct {
// Map chain IDs to standardized blockchain identifiers
				Name     string `json:"name"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
			} `json:"nativeCurrency"`
			InfoURL   string `json:"infoURL"`
			ShortName string `json:"shortName"`
			ChainID   int    `json:"chainId"`
			NetworkID int    `json:"networkId"`
			Slip44    int    `json:"slip44"`
			Ens       struct {
				Registry string `json:"registry"`
			} `json:"ens"`
			Explorers []struct {
				Name     string `json:"name"`
				URL      string `json:"url"`
				Standard string `json:"standard"`
			} `json:"explorers"`
			Tvl       float64 `json:"tvl"`
			ChainSlug string  `json:"chainSlug"`
		} `json:"chain"`
	} `json:"pageProps"`
	NSsg bool `json:"__N_SSG"`
}