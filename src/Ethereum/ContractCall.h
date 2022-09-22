// Copyright © 2017-2020 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#pragma once

#include "Data.h"
#include "ABI.h"
#include "ContractCallParam.h"

#include <nlohmann/json.hpp>
#include <optional>
#include <string>
#include <vector>
#include <boost/algorithm/string/predicate.hpp>


namespace TW::Ethereum::ABI {
    std::optional<std::string> decodeCall(const Data& call, const nlohmann::json& abi);
    Data buildContractCallData(const std::string& functionName, const std::vector<ContractCallParam> params);
} // namespace TW::Ethereum::ABI
