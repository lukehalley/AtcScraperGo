package io

import (
	"atcscraper/src/types/geckoterminal"
	"encoding/gob"
	"log"
	"os"
)

func SaveGTDataToCache(DataToSave interface{}, FileName string)  {

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

func ReadGTCacheFile(FileName string) geckoterminal.GetGTNetworks {

	// List For Storing Cache Dexs
	var CacheNetworks geckoterminal.GetGTNetworks

	// Create The Cache File
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
}