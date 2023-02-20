// Copyright © 2017-2023 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#include "Solana/VersionedTransaction.h"
#include "Solana/Encoding.h"

namespace TW::Solana {

std::string VersionedTransaction::serialize() const {
    Data buffer;

    append(buffer, shortVecLength<Signature>(this->signatures));
    for (auto signature : this->signatures) {
        Data signature_vec(signature.bytes.begin(), signature.bytes.end());
        append(buffer, signature_vec);
    }
    append(buffer, this->messageData());

    return Base58::bitcoin.encode(buffer);
}

Data VersionedTransaction::messageData() const {
    return Solana::serialize(this->message);
}

uint8_t VersionedTransaction::getAccountIndex(Address publicKey) {
    const auto accountKeys = Solana::accountKeys(this->message);
    auto item = std::find(accountKeys.begin(), accountKeys.end(), publicKey);
    if (item == accountKeys.end()) {
        throw std::invalid_argument("publicKey not found in message.accountKeys");
    }
    return (uint8_t)std::distance(accountKeys.begin(), item);
}
} // namespace TW::Solana
