// Copyright © 2017-2020 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#include "Entry.h"

#include "Bitcoin/Address.h"
#include "Signer.h"

namespace TW {
namespace Hydra {

// Define a helper function to convert Base58Prefix to byte
inline byte base58PrefixToByte(const Base58Prefix& prefix) {
    // You may choose the desired value (e.g., p2pkh or p2sh) based on your use case.
    return prefix.p2pkh;
}

// Define the conversion function
inline byte toByte(const PrefixVariant& prefix) {
    return std::visit([](auto&& arg) -> byte {
        using T = std::decay_t<decltype(arg)>;
        if constexpr (std::is_same_v<T, Base58Prefix>) {
            return base58PrefixToByte(arg);
        } else if constexpr (std::is_same_v<T, Bech32Prefix>) {
            // Conversion from Bech32Prefix (const char*) to byte is not straightforward.
            // You may need to provide a custom implementation based on your use case.
            return 0;
        } else if constexpr (std::is_same_v<T, SS58Prefix>) {
            return static_cast<byte>(arg);
        } else if constexpr (std::is_same_v<T, DelegatedPrefix>) {
            // Conversion from DelegatedPrefix to byte may need a custom implementation.
            return 0;
        } else {
            return 0;
        }
    },
                      prefix);
}
} // namespace Hydra
} // namespace TW

using namespace TW::Hydra;
using namespace std;

bool Entry::validateAddress(TWCoinType coin, const string& address, const PrefixVariant& addressPrefix) const {
    return Bitcoin::Address::isValid(address);
}

string Entry::deriveAddress(TWCoinType coin, const PublicKey& publicKey, TWDerivation derivation, const PrefixVariant& addressPrefix) const {

    byte p2pkh = getFromPrefixPkhOrDefault(addressPrefix, coin);
    return Bitcoin::Address(publicKey, p2pkh).string();
}

void Entry::sign(TWCoinType coin, const TW::Data& dataIn, TW::Data& dataOut) const {
    signTemplate<Signer, Hydra::Proto::SigningInput>(dataIn, dataOut);
}

void Entry::plan(TWCoinType coin, const TW::Data& dataIn, TW::Data& dataOut) const {
    planTemplate<Signer, Hydra::Proto::SigningInput>(dataIn, dataOut);
}

TW::Data Entry::addressToData(TWCoinType coin, const std::string& address) const {
    auto addr = Bitcoin::Address(address);
    return {addr.bytes.begin() + 1, addr.bytes.end()};
}