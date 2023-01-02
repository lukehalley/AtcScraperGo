package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)


func AddTokenToDB(TokenSymbol string, TokenAddress string, TokenDecimals int, TokenIsStablecoin int, TokenNetworkId int) int64 {

	DBKeys := "network_id, symbol, address, decimals, stablecoin"
	SelectStatement := fmt.Sprintf("(SELECT %d AS network_id, '%v' AS symbol, '%v' AS address, '%d' AS decimals, %d AS stablecoin)", TokenNetworkId, TokenSymbol, TokenAddress, TokenDecimals, TokenIsStablecoin)
	CompareStatement := fmt.Sprintf("tokens.network_id = %d AND tokens.address = '%v'", TokenNetworkId, TokenAddress)

	InsertQuery := "INSERT INTO tokens(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM tokens WHERE " + CompareStatement + ") LIMIT 1"

	InsertedTokenID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedTokenID

}