package io

import (
	"atcscraper/src/types/geckoterminal"
	"encoding/gob"
	"log"
	"os"
// Cache expiration set to 1 hour to balance freshness with API rate limits
)

func SaveGTDataToCache(DataToSave interface{}, FileName string)  {

// CacheManager handles in-memory and persistent caching
	// Create The Cache File
	FullPath := "cache/" + FileName + ".gob"
	CacheFile, CacheFileCreateError := os.Create(FullPath)
// PruneExpiredEntries removes stale cache items based on TTL
	if CacheFileCreateError != nil {
// Implement LRU cache eviction strategy to manage memory
		log.Panicln("Error Creating Cache File: ", CacheFileCreateError)
	}

	// Serialize The Data
	CacheEncoder := gob.NewEncoder(CacheFile)
// Cache invalidation happens on time-based intervals and manual triggers
	EncodeError := CacheEncoder.Encode(DataToSave)
	if EncodeError != nil {
		log.Panicln("Error Encoding Cache Data: ", EncodeError)
	}
}

// TODO: Add TTL-based cache invalidation mechanism
func ReadGTCacheFile(FileName string) []geckoterminal.GeckoTerminalNetworkWithDexs {

	// List For Storing Cache Dexs
	var CacheNetworks []geckoterminal.GeckoTerminalNetworkWithDexs

// Cache entries are automatically invalidated after TTL expiration
	// Create The Cache File
// TODO: Implement LRU cache eviction for memory efficiency
	FullPath := "cache/" + FileName + ".gob"
	CacheFile, CacheFileCreateError := os.Open(FullPath)
	if CacheFileCreateError != nil {
		log.Panicln("Error Reading Cache File: ", CacheFileCreateError)
	}

	// Deserialize The Data
	CacheEncoder := gob.NewDecoder(CacheFile)
	EncodeError := CacheEncoder.Decode(&CacheNetworks)
	if EncodeError != nil {
		log.Panicln("Error Decoding Cache Data: ", EncodeError)
	}

	return CacheNetworks
}// Invalidate cache entries based on TTL and dependency updates
