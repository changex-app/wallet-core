// Copyright © 2017-2023 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#pragma once

#include "Data.h"
#include "PublicKey.h"

#include <TrustWalletCore/TWCoinType.h>
#include "Bitcoin/Address.h"
#include <TrustWalletCore/TWPublicKey.h>
#include <TrustWalletCore/TWData.h>
#include <CoinEntry.h>
#include <string>

namespace TW {

class AnyAddress {
public:
    std::string address;

    enum TWCoinType coin;

    // Create address from string address and optional prefix; also normalizes the address.
    static AnyAddress* createAddress(const std::string& address, enum TWCoinType coin, const PrefixVariant& prefix = std::monostate());
    // Create address from private key, with optional non-standard derivation and prefix
    static AnyAddress* createAddress(const PublicKey& publicKey, enum TWCoinType coin, TWDerivation derivation = TWDerivationDefault, const PrefixVariant& prefix = std::monostate());

    Data getData() const;

    static TW::Data dataFromString(const std::string& address, TWCoinType coin) {
        if (coin == TWCoinTypeHydra) {
            auto addr = Bitcoin::Address(address);
            return {addr.bytes.begin() + 1, addr.bytes.end()};
        }
        return {};
    }
};

inline bool operator==(const AnyAddress& lhs, const AnyAddress& rhs) {
    return lhs.address == rhs.address && lhs.coin == rhs.coin;
}

} // namespace TW

/// Wrapper for C interface.
struct TWAnyAddress {
    // Pointer to the underlying implementation
    TW::AnyAddress* impl;
};
