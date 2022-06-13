package launcher

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var (
	EvmExportMode = cli.StringFlag{
		Name:  "export.evm.mode",
		Usage: `EVM export mode ("full" or "ext-mpt" or "mpt" or "none")`,
		Value: "mpt",
	}
	mptTraversalMode = cli.StringFlag{
		Name:  "mpt.traversal.mode",
		Usage: `Traversal mode ("mpt" or "snap")`,
		Value: "snap",
	}
	importCommand = cli.Command{
		Name:      "import",
		Usage:     "Import a blockchain file",
		ArgsUsage: "<filename> (<filename 2> ... <filename N>) [check=false]",
		Category:  "MISCELLANEOUS COMMANDS",
		Description: `
    opera import events

The import command imports events from an RLP-encoded files.
Events are fully verified by default, unless overridden by check=false flag.`,

		Subcommands: []cli.Command{
			{
				Action:    utils.MigrateFlags(importEvents),
				Name:      "events",
				Usage:     "Import blockchain events",
				ArgsUsage: "<filename> (<filename 2> ... <filename N>)",
				Flags: []cli.Flag{
					DataDirFlag,
				},
				Description: `
The import command imports events from RLP-encoded files.
Events are fully verified by default, unless overridden by --check=false flag.`,
			},
			{
				Action:    utils.MigrateFlags(importEvm),
				Name:      "evm",
				Usage:     "Import EVM storage",
				ArgsUsage: "<filename> (<filename 2> ... <filename N>)",
				Flags: []cli.Flag{
					DataDirFlag,
				},
				Description: `
    opera import evm

The import command imports EVM storage (trie nodes, code, preimages) from files.`,
			},
		},
	}
	exportCommand = cli.Command{
		Name:     "export",
		Usage:    "Export blockchain",
		Category: "MISCELLANEOUS COMMANDS",

		Subcommands: []cli.Command{
			{
				Name:      "events",
				Usage:     "Export blockchain events",
				ArgsUsage: "<filename> [<epochFrom> <epochTo>]",
				Action:    utils.MigrateFlags(exportEvents),
				Flags: []cli.Flag{
					DataDirFlag,
				},
				Description: `
    opera export events

Requires a first argument of the file to write to.
Optional second and third arguments control the first and
last epoch to write. If the file ends with .gz, the output will
be gzipped
`,
			},
			{
				Name:      "genesis",
				Usage:     "Export current state into a genesis file",
				ArgsUsage: "<filename or dry-run> [<epochFrom> <epochTo>] [--export.evm.mode=none]",
				Action:    utils.MigrateFlags(exportGenesis),
				Flags: []cli.Flag{
					DataDirFlag,
					EvmExportMode,
				},
				Description: `
    opera export genesis

Export current state into a genesis file.
Requires a first argument of the file to write to.
Optional second and third arguments control the first and
last epoch to write.
Pass dry-run instead of filename for calculation of hashes without exporting data.
EVM export mode is configured with --export.evm.mode.
`,
			},
		},
	}
	checkCommand = cli.Command{
		Name:     "check",
		Usage:    "Check blockchain",
		Category: "MISCELLANEOUS COMMANDS",

		Subcommands: []cli.Command{
			{
				Name:   "evm",
				Usage:  "Check EVM storage",
				Action: utils.MigrateFlags(checkEvm),
				Flags: []cli.Flag{
					DataDirFlag,
				},
				Description: `
    opera check evm

Checks EVM storage roots and code hashes
`,
			},
		},
	}

	erigonCommand = cli.Command{
		Name:     "erigon",
		Usage:    "Writes EVM data into Erigon tables",
		Category: "MISCELLANEOUS COMMANDS",
		Action: utils.MigrateFlags(writeEVMToErigon),
		Flags: []cli.Flag{
			mptTraversalMode,
		},

		Description: `
		opera erigon
	
	The erigon command iterates over MPT data and flushes it into Erigon tables`,

	
	}
	importPreimagesCommand = cli.Command{
		Name:     "preimages",
		Usage:    "Import preimages and prints them out",
		Category: "MISCELLANEOUS COMMANDS",
		Action: utils.MigrateFlags(importPreimagesCmd),
		//Flags: []cli.Flag{mptTraversalMode,},

		Description: `
		opera preimages
	
	The erigon command iterates over MPT data and flushes it into Erigon tables`,

	
	}
)
