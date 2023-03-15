// Copyright © 2017-2023 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#pragma once

#include "Data.h"
#include "../PublicKey.h"
#include  "SS58Address.h"
#include  <TrustWalletCore/TWSS58AddressType.h>

#include <string>

namespace TW::Polkadot {

class Address: public SS58Address {
  public:
    /// Determines whether a string makes a valid address.
    static bool isValid(const std::string& string) { return SS58Address::isValid(string, TWSS58AddressTypePolkadot); }
    static bool isValid(const std::string& string, uint32_t ss58Prefix) { return SS58Address::isValid(string, ss58Prefix); }

    /// Initializes a Polkadot address with a string representation.
    explicit Address(const std::string& string): SS58Address(string, TWSS58AddressTypePolkadot) {}

    /// Initializes a Polkadot address with a public key.
    explicit Address(const PublicKey& publicKey): SS58Address(publicKey, TWSS58AddressTypePolkadot) {}
    /// Initializes a Polkadot address with a public key and a given ss58Prefix.
    explicit Address(const PublicKey& publicKey, std::uint32_t ss58Prefix): SS58Address(publicKey, ss58Prefix) {}
};
} // namespace TW::Polkadot

