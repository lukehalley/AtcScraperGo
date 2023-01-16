package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)

func AddABIToDB(NetworkId int, DexId int, ContractName string, ContractAddress string, Abi string, Mappings string) int64 {

	DBKeys := "network_id, dex_id, contract_name, contract_address, abi"
	SelectStatement := fmt.Sprintf("(SELECT %d AS network_id, %d AS dex_id, '%v' AS contract_name, '%v' AS contract_address, '%v' AS abi", NetworkId, DexId, ContractName, ContractAddress, Abi)

	if Mappings != "" {
		DBKeys = fmt.Sprintf("%v, mapping", DBKeys)
		SelectStatement = fmt.Sprintf("%v, '%v' AS mapping)", SelectStatement, Mappings)
	} else {
		SelectStatement = fmt.Sprintf("%v)", SelectStatement)
	}

	CompareStatement := fmt.Sprintf("abis.network_id = %d AND abis.contract_address = '%v'", NetworkId, ContractAddress)

	InsertQuery := "INSERT INTO abis(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM abis WHERE " + CompareStatement + ") LIMIT 1"

	InsertedABIId := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedABIId

}
