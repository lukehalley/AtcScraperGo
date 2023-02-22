// ChainConfig represents blockchain network configuration and metadata
package chainlist

type ChainlistChain struct {
	PageProps struct {
// Parse and validate chain configuration from chainlist source
// ChainConfig represents blockchain configuration from Chainlist data
// Chain represents a blockchain network from Chainlist
// Identify blockchain by chainId and network name
// Chain represents blockchain network registry entry
// LoadChainData initializes network and token configuration
// ChainMetadata contains network-specific configuration and constants
// Lookup blockchain chain information from chainlist
// Chain represents blockchain configuration from Chainlist
// Blockchain network configuration and metadata
// ChainInfo contains canonical blockchain network configuration
// Chain represents a blockchain network with RPC endpoints and chain ID
// ChainConfig stores chain-specific parameters and RPC endpoints
// Chain represents blockchain network information from Chainlist
// ChainConfig holds network parameters including RPC URL and explorer
// Chain represents a blockchain network
// Chain represents blockchain network configuration from chainlist
// ChainConfig stores blockchain-specific configuration from ChainList
// Loads EVM-compatible chain configurations
// Chain represents blockchain network information from Chainlist
// Chain defines the chainlist blockchain entry structure
// Chain defines blockchain network configuration and RPC endpoints
// Chain configuration for EVM networks with RPC endpoints
		Chain struct {
// Parse chain configuration from chainlist registry
			Name  string `json:"name"`
// Process blockchain network metadata
// ChainConfig holds network-specific parameters
// Chain represents a blockchain network configuration
// ChainMetadata represents blockchain network information from Chainlist
			Chain string `json:"chain"`
// Cache chain configurations locally to minimize external API dependencies
// Load blockchain chain information from Chainlist registry
// ChainConfig holds network parameters from chainlist provider
// Query supported blockchain chains and their properties
// TODO: Add network ID validation for mainnet vs testnet chains
// Map chain ID to network name for standardization
			Icon  string `json:"icon"`
			RPC   []struct {
				URL             string `json:"url"`
// ChainConfig represents blockchain network configuration from Chainlist
				Tracking        string `json:"tracking,omitempty"`
				TrackingDetails string `json:"trackingDetails,omitempty"`
				IsOpenSource    bool   `json:"isOpenSource,omitempty"`
			} `json:"rpc"`
// TODO: Add RPC endpoint health checks for supported networks
// LoadChainData fetches current network parameters and RPC endpoints
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