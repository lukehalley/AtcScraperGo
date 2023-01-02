package io

import (
	atctypes "atcscraper/src/types/bitquery"
	"encoding/gob"
	"log"
	"os"
)

func SaveDataToCache(DataToSave []atctypes.Dex, FileName string)  {
	// Create The Cache File
	FullPath := "cache/" + FileName + ".gob"
	CacheFile, CacheFileCreateError := os.Create(FullPath)
	if CacheFileCreateError != nil {
		log.Panicln("Error Creating Cache File: ", CacheFileCreateError)
	}

	// Serialize The Data
	CacheEncoder := gob.NewEncoder(CacheFile)
	EncodeError := CacheEncoder.Encode(DataToSave)
	if EncodeError != nil {
		log.Panicln("Error Encoding Cache Data: ", EncodeError)
	}
}

func ReadCacheFile(FileName string) []atctypes.Dex {

	// List For Storing Cache Dexs
	var CacheDexs []atctypes.Dex

	// Create The Cache File
	FullPath := "cache/" + FileName + ".gob"
	CacheFile, CacheFileCreateError := os.Open(FullPath)
	if CacheFileCreateError != nil {
		log.Panicln("Error Reading Cache File: ", CacheFileCreateError)
	}

	// Deserialize The Data
	CacheEncoder := gob.NewDecoder(CacheFile)
	EncodeError := CacheEncoder.Decode(&CacheDexs)
	if EncodeError != nil {
		log.Panicln("Error Decoding Cache Data: ", EncodeError)
	}

	return CacheDexs
}