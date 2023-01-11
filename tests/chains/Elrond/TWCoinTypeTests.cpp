// Copyright © 2017-2020 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.
//
// This is a GENERATED FILE, changes made here MAY BE LOST.
// Generated one-time (codegen/bin/cointests)
//

#include "TestUtilities.h"
#include <TrustWalletCore/TWCoinTypeConfiguration.h>
#include <gtest/gtest.h>


TEST(TWElrondCoinType, TWCoinType) {
    auto symbol = WRAPS(TWCoinTypeConfigurationGetSymbol(TWCoinTypeElrond));
    auto txId = WRAPS(TWStringCreateWithUTF8Bytes("163b46551a74626415074b626d2f37d3c78aef0f6ccb628db434ee65a35ea127"));
    auto txUrl = WRAPS(TWCoinTypeConfigurationGetTransactionURL(TWCoinTypeElrond, txId.get()));
    auto accId = WRAPS(TWStringCreateWithUTF8Bytes("erd1qyu5wthldzr8wx5c9ucg8kjagg0jfs53s8nr3zpz3hypefsdd8ssycr6th"));
    auto accUrl = WRAPS(TWCoinTypeConfigurationGetAccountURL(TWCoinTypeElrond, accId.get()));
    auto id = WRAPS(TWCoinTypeConfigurationGetID(TWCoinTypeElrond));
    auto name = WRAPS(TWCoinTypeConfigurationGetName(TWCoinTypeElrond));

    ASSERT_EQ(TWCoinTypeConfigurationGetDecimals(TWCoinTypeElrond), 18);
    ASSERT_EQ(TWBlockchainElrondNetwork, TWCoinTypeBlockchain(TWCoinTypeElrond));
    ASSERT_EQ(0x0, TWCoinTypeP2shPrefix(TWCoinTypeElrond));
    ASSERT_EQ(0x0, TWCoinTypeStaticPrefix(TWCoinTypeElrond));
    assertStringsEqual(symbol, "eGLD");
    assertStringsEqual(txUrl, "https://explorer.multiversx.com/transactions/163b46551a74626415074b626d2f37d3c78aef0f6ccb628db434ee65a35ea127");
    assertStringsEqual(accUrl, "https://explorer.multiversx.com/accounts/erd1qyu5wthldzr8wx5c9ucg8kjagg0jfs53s8nr3zpz3hypefsdd8ssycr6th");
    assertStringsEqual(id, "elrond");
    assertStringsEqual(name, "MultiversX");
}
