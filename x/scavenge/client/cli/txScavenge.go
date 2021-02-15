package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/starport/scavenge/x/scavenge/types"
)

func GetCmdCreateScavenge(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-scavenge [description] [solutionHash] [reward] [solution] [scavenger]",
		Short: "Creates a new scavenge",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDescription := string(args[0] )
			argsSolutionHash := string(args[1] )
			argsReward := string(args[2] )
			argsSolution := string(args[3] )
			argsScavenger := string(args[4] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateScavenge(cliCtx.GetFromAddress(), string(argsDescription), string(argsSolutionHash), string(argsReward), string(argsSolution), string(argsScavenger))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetScavenge(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-scavenge [id]  [description] [solutionHash] [reward] [solution] [scavenger]",
		Short: "Set a new scavenge",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsDescription := string(args[1])
			argsSolutionHash := string(args[2])
			argsReward := string(args[3])
			argsSolution := string(args[4])
			argsScavenger := string(args[5])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetScavenge(cliCtx.GetFromAddress(), id, string(argsDescription), string(argsSolutionHash), string(argsReward), string(argsSolution), string(argsScavenger))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteScavenge(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-scavenge [id]",
		Short: "Delete a new scavenge by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteScavenge(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
