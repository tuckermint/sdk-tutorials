package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/tuckermint/sdk-application-tutorial/x/superbank/internal/types"
)

// Simulation parameter constants
const (
	SendEnabled = "send_enabled"
)

// GenSendEnabled randomized SendEnabled
func GenSendEnabled(r *rand.Rand) bool {
	return r.Int63n(2) == 0
}

// RandomizedGenState generates a random GenesisState for bank
func RandomizedGenState(simState *module.SimulationState) {
	var sendEnabled bool
	simState.AppParams.GetOrGenerate(
		simState.Cdc, SendEnabled, &sendEnabled, simState.Rand,
		func(r *rand.Rand) { sendEnabled = GenSendEnabled(r) },
	)

	bankGenesis := types.NewGenesisState(sendEnabled)

	fmt.Printf("Selected randomly generated bank parameters:\n%s\n", codec.MustMarshalJSONIndent(simState.Cdc, bankGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(bankGenesis)
}
