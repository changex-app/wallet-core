// Copyright © 2017-2023 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#include "MessageSigner.h"
#include "Address.h"

#include "Base64.h"
#include "BinaryCoding.h"
#include "Coin.h"
#include "Data.h"
#include "HexCoding.h"
#include "../../include/TrustWalletCore/TWCoinType.h"

using namespace TW;

namespace TW::Bitcoin {

// lenght-encode a message string
Data messageToData(const std::string& message) {
    Data d;
    TW::encodeVarInt(message.size(), d);
    TW::append(d, TW::data(message));
    return d;
}

// append prefix and length-encode message string
Data messageToFullData(const std::string& message, enum TWCoinType coin) {

    if(coin == TWCoinTypeHydra) {
        Data d = messageToData(MessageSigner::MessageHydraPrefix);
        TW::append(d, messageToData(message));
        return d;
    } else {
        Data d = messageToData(MessageSigner::MessagePrefix);
        TW::append(d, messageToData(message));
        return d;
    }
}

Data MessageSigner::messageToHash(const std::string& message, enum TWCoinType coin) {
    Data d = messageToFullData(message, coin);
    return Hash::sha256d(d.data(), d.size());
}

std::string MessageSigner::signMessage(const PrivateKey& privateKey, const std::string& address, const std::string& message, enum TWCoinType coin, bool compressed) {
    if (!Address::isValid(address)) {
        throw std::invalid_argument("Address is not valid (legacy) address");
    }
    std::string addrFromKey;
    if (compressed) {
        const auto pubKey = privateKey.getPublicKey(TWPublicKeyTypeSECP256k1);
        addrFromKey = Address(pubKey, TW::p2pkhPrefix(coin)).string();
    } else {
        const auto pubKey = privateKey.getPublicKey(TWPublicKeyTypeSECP256k1Extended);
        const auto keyHash = pubKey.hash(Data{TW::p2pkhPrefix(coin)}, Hash::HasherSha256ripemd);
        addrFromKey = Address(keyHash).string();
    }
    if (addrFromKey != address) {
        throw std::invalid_argument("Address does not match key");
    }
    const auto messageHash = messageToHash(message, coin);
    const auto signature = privateKey.sign(messageHash, TWCurveSECP256k1);

    // The V value: add 31 (or 27 for compressed), and move to the first byte
    const byte v = signature[SignatureRSLength] + PublicKey::SignatureVOffset + (compressed ? 4ul : 0ul);
    auto sigAdjusted = Data{v};
    TW::append(sigAdjusted, TW::subData(signature, 0, SignatureRSLength));
    return Base64::encode(sigAdjusted);
}

std::string MessageSigner::recoverAddressFromMessage(const std::string& message, const Data& signature, enum TWCoinType coin) {
    if (signature.size() < SignatureRSVLength) {
        throw std::invalid_argument("signature too short");
    }
    const auto messageHash = MessageSigner::messageToHash(message, coin);
    auto recId = signature[0];
    auto compressed = false;
    if (recId >= PublicKey::SignatureVOffset + 4) {
        recId -= 4;
        compressed = true;
    }
    if (recId >= PublicKey::SignatureVOffset) {
        recId -= PublicKey::SignatureVOffset;
    }

    const auto publicKeyRecovered = PublicKey::recoverRaw(TW::subData(signature, 1), recId, messageHash);

    if (!compressed) {
        // uncompressed public key
        const auto keyHash = publicKeyRecovered.hash(Data{TW::p2pkhPrefix(TWCoinTypeBitcoin)}, Hash::HasherSha256ripemd);
        return Bitcoin::Address(keyHash).string();
    }
    // compressed
    const auto publicKeyRecoveredCompressed = publicKeyRecovered.compressed();
    return Bitcoin::Address(publicKeyRecoveredCompressed, TW::p2pkhPrefix(TWCoinTypeBitcoin)).string();
}

bool MessageSigner::verifyMessage(const std::string& address, const std::string& message, const std::string& signature, enum TWCoinType coin) noexcept {
    try {
        const auto signatureData = Base64::decode(signature);
        return verifyMessage(address, message, signatureData, coin);
    } catch (...) {
        return false;
    }
}

/// May throw
bool MessageSigner::verifyMessage(const std::string& address, const std::string& message, const Data& signature, enum TWCoinType coin) {
    if (!Bitcoin::Address::isValid(address)) {
        throw std::invalid_argument("Input address invalid, must be valid legacy");
    }
    const auto addressRecovered = recoverAddressFromMessage(message, signature, coin);
    return (addressRecovered == address);
}

} // namespace TW::Bitcoin
