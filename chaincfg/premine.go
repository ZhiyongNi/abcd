// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2017 The Aero Blockchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

// BlockOneLedgerMainNet is the block one output ledger for the main
// network.
var BlockOneLedgerMainNet = []*TokenPayout{
	{"Ab9iJvXsBq56HZjVCibApKKWZkfiR7n51ef", 100000000 * 1e8},
}

// BlockOneLedgerTestNet is the block one output ledger for the test
// network.
var BlockOneLedgerTestNet = []*TokenPayout{
	{"TsmWaPM77WSyA3aiQ2Q1KnwGDVWvEkhipBc", 100000 * 1e8},
}

// BlockOneLedgerTestNet2 is the block one output ledger for the 2nd test
// network.
var BlockOneLedgerTestNet2 = []*TokenPayout{
	{"TsT5rhHYqJF7sXouh9jHtwQEn5YJ9KKc5L9", 100000 * 1e8},
}

// BlockOneLedgerSimNet is the block one output ledger for the simulation
// network. See under "Aero organization related parameters" in params.go
// for information on how to spend these outputs.
var BlockOneLedgerSimNet = []*TokenPayout{
	{"Sshw6S86G2bV6W32cbc7EhtFy8f93rU6pae", 100000 * 1e8},
	{"SsjXRK6Xz6CFuBt6PugBvrkdAa4xGbcZ18w", 100000 * 1e8},
	{"SsfXiYkYkCoo31CuVQw428N6wWKus2ZEw5X", 100000 * 1e8},
}
