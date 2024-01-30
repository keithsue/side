package keepers

import (
	"cosmossdk.io/x/evidence"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"

	//upgradeclient "cosmossdk.io/x/upgrade/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/ibc-go/modules/capability"
	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"

	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"

	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v8/modules/core/02-client"
	solomachine "github.com/cosmos/ibc-go/v8/modules/light-clients/06-solomachine"
	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"

	gmmmodule "github.com/sideprotocol/side/x/gmm"

	// this line is used by starport scaffolding # stargate/app/moduleImport

	// wasmd module integrate
	"github.com/CosmWasm/wasmd/x/wasm"
	// upgrades
)

// AppModuleBasics returns ModuleBasics for the module BasicManager.
var AppModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	distr.AppModuleBasic{},
	gov.NewAppModuleBasic(getGovProposalHandlers()),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	feegrantmodule.AppModuleBasic{},
	groupmodule.AppModuleBasic{},
	ibc.AppModuleBasic{},
	ibctm.AppModuleBasic{},
	solomachine.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
	transfer.AppModuleBasic{},
	ica.AppModuleBasic{},
	vesting.AppModuleBasic{},
	consensus.AppModuleBasic{},
	ibcfee.AppModuleBasic{},
	wasm.AppModuleBasic{},
	gmmmodule.AppModuleBasic{},
)

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		//upgradeclient.LegacyProposalHandler,
		//upgradeclient.LegacyCancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}
