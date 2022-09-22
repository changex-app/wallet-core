#include "ContractCallParam.h"

using namespace TW;
using namespace Ethereum;

ContractCallParam::ContractCallParam(const Proto::ContractCallParam& param)
    : type(param.type())
{
    for (auto& v: param.value()){
        value.push_back(Data(v.begin(), v.end()));
    }
}