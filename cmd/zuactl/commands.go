package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/zuanet/zuad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.ZuadMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.ZuadMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.ZuadMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.ZuadMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.ZuadMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.ZuadMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.ZuadMessage_BanRequest{}),
	reflect.TypeOf(protowire.ZuadMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
