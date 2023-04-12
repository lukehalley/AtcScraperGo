package lambda

import (
	atcaws "atcscraper/src/aws"
	"atcscraper/src/log"
	awstypes "atcscraper/src/types/aws"
	"atcscraper/src/util"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"strings"
)

func CallDecodeLambda(RPCUrl string, TxHash string, ContractHash string, RouterAbiDbId int64) awstypes.DecodeLambdaResponse {

	// Build Lambda Params
	Params := awstypes.DecodeLambdaParams{RpcUrl: RPCUrl, TxHash: TxHash, ContractHash: ContractHash, ContractAbiDBId: RouterAbiDbId}

	// Marshal Params
	LambdaPayload, LambdaPayloadMarshalError := json.Marshal(Params)
	if LambdaPayloadMarshalError != nil {
		Error := fmt.Sprintf("Error Marshalling Decode Lambda Params")
		log.NewError(Error)
	}

	// Create A New aws Client
	AWSClient := atcaws.NewAWSClient()

	// Invoke Lambda
	DecodeLambdaResult, DecodeLambdaError := AWSClient.Invoke(&lambda.InvokeInput{FunctionName: aws.String("atc-decoder-prod-decode"), Payload: LambdaPayload})
	if DecodeLambdaError != nil {
		Error := fmt.Sprintf("Error Calling Decode Lambda: %v", DecodeLambdaError)

		if strings.Contains(DecodeLambdaError.Error(), "Rate Exceeded") {
			for InError := true; InError; InError = strings.Contains(DecodeLambdaError.Error(), "Rate Exceeded") {
				util.SleepForRandomRange(1, 3)
				DecodeLambdaResult, DecodeLambdaError = AWSClient.Invoke(&lambda.InvokeInput{FunctionName: aws.String("atc-decoder-prod-decode"), Payload: LambdaPayload})
				if DecodeLambdaError == nil {
					break
				}
			}
		} else {
			log.NewError(Error)
		}

	}

	// Unmarshal The Lambda Response
	var DecodeLambdaResponse awstypes.DecodeLambdaResponse
	_ = json.Unmarshal(DecodeLambdaResult.Payload, &DecodeLambdaResponse)

	return DecodeLambdaResponse

}