// Copyright © 2017-2023 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#pragma once

#include "../CoinEntry.h"

namespace TW::Hydra {
/// Hydra entry dispatcher.

/// Note: do not put the implementation here (no matter how simple), to avoid having coin-specific includes in this file
class Entry final : public CoinEntry {
public:
    virtual bool validateAddress(TWCoinType coin, const std::string& address, const PrefixVariant& addressPrefix) const;
    virtual std::string deriveAddress(TWCoinType coin, const PublicKey& publicKey, TWDerivation derivation, const PrefixVariant& addressPrefix) const;
    virtual void sign(TWCoinType coin, const Data& dataIn, Data& dataOut) const;
    virtual void plan(TWCoinType coin, const Data& dataIn, Data& dataOut) const;

    virtual Data addressToData(TWCoinType coin, const std::string& address) const;
};

} // namespace TW::Hydra
