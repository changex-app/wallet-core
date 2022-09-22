#pragma once

#include <string>
#include "../proto/Ethereum.pb.h"
#include "../Data.h"

namespace TW::Ethereum{

class ContractCallParam{
public:
    std::string type;

    std::vector<Data> value;

    ContractCallParam() = default;

    ContractCallParam(const Proto::ContractCallParam& param);
};
} // namespace TW::Ethereum