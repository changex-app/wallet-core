// Copyright © 2017-2021 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#include "ContractCall.h"
#include "HexCoding.h"
#include "uint256.h"
#include <boost/algorithm/string/predicate.hpp>

using namespace std;
using json = nlohmann::json;

namespace TW::Ethereum::ABI {

static void fillArray(ParamSet& paramSet, const string& type) {
    auto baseType = string(type.begin(), type.end() - 2);
    auto value = ParamFactory::make(baseType);
    auto param = make_shared<ParamArray>(value);
    paramSet.addParam(param);
}

static void fill(ParamSet& paramSet, const string& type) {
    if (boost::algorithm::ends_with(type, "[]")) {
        fillArray(paramSet, type);
    } else {
        auto param = ParamFactory::make(type);
        paramSet.addParam(param);
    }
}

static vector<string> getArrayValue(ParamSet& paramSet, const string& type, int idx) {
    shared_ptr<ParamBase> param;
    paramSet.getParam(idx, param);
    return ParamFactory::getArrayValue(param, type);
}

static json buildInputs(ParamSet& paramSet, const json& registry); // forward

static json getTupleValue(ParamSet& paramSet, const string& type, int idx, const json& typeInfo) {
    shared_ptr<ParamBase> param;
    paramSet.getParam(idx, param);
    auto paramTuple = dynamic_pointer_cast<ParamTuple>(param);
    if (!paramTuple.get()) {
        return {};
    }
    return buildInputs(paramTuple->_params, typeInfo["components"]);
}

static string getValue(ParamSet& paramSet, const string& type, int idx) {
    shared_ptr<ParamBase> param;
    paramSet.getParam(idx, param);
    return ParamFactory::getValue(param, type);
}

static json buildInputs(ParamSet& paramSet, const json& registry) {
    auto inputs = json::array();
    for (int i = 0; i < registry.size(); i++) {
        auto info = registry[i];
        auto type = info["type"];
        auto input = json{
            {"name", info["name"]},
            {"type", type}
        };
        if (boost::algorithm::ends_with(type.get<string>(), "[]")) {
            input["value"] = json(getArrayValue(paramSet, type, i));
        } else if (type == "tuple") {
            input["components"] = getTupleValue(paramSet, type, i, info);
        } else if (type == "bool") {
            input["value"] = getValue(paramSet, type, i) == "true" ? json(true) : json(false);
        } else {
            input["value"] = getValue(paramSet, type, i);
        }
        inputs.push_back(input);
    }
    return inputs;
}

void fillTuple(ParamSet& paramSet, const json& jsonSet); // forward

void decodeParamSet(ParamSet& paramSet, const json& jsonSet) {
    for (auto& comp : jsonSet) {
        if (comp["type"] == "tuple") {
            fillTuple(paramSet, comp["components"]);
        } else {        
            fill(paramSet, comp["type"]);
        }
    }    
}

void fillTuple(ParamSet& paramSet, const json& jsonSet) {
    std::shared_ptr<ParamTuple> param = make_shared<ParamTuple>();
    decodeParamSet(param->_params, jsonSet);
    paramSet.addParam(param);
}

optional<string> decodeCall(const Data& call, const json& abi) {
    // check bytes length
    if (call.size() <= 4) {
        return {};
    }

    auto methodId = hex(Data(call.begin(), call.begin() + 4));

    if (abi.find(methodId) == abi.end()) {
        return {};
    }

    // build Function with types
    const auto registry = abi[methodId];
    auto func = Function(registry["name"]);
    decodeParamSet(func._inParams, registry["inputs"]);

    // decode inputs
    size_t offset = 0;
    auto success = func.decodeInput(call, offset);
    if (!success) {
        return {};
    }

    // build output json
    auto decoded = json{
        {"function", func.getType()},
        {"inputs", buildInputs(func._inParams, registry["inputs"])},
    };
    return decoded.dump();
}

bool isArrayType(const std::string& type) {
    return boost::algorithm::ends_with(type, "[]") && type.length() >= 3;
}

std::string getArrayElemType(const std::string& arrayType) {
    if (isArrayType(arrayType)) {
        return arrayType.substr(0, arrayType.length() - 2);
    }
    return "";
}

bool isNumberType(const std::string& type) {
    return type == "uint8"
        || type == "uint16" 
        || type == "uint32" 
        || type == "uint64"
        || type == "int8"
        || type == "int16" 
        || type == "int32" 
        || type == "int64";
}

Data buildContractCallData(const std::string& functionName, const std::vector<ContractCallParam> params) {
    std::vector<std::shared_ptr<Ethereum::ABI::ParamBase>> parameters;
    for(auto& param: params){ 
       auto abiParam = ParamFactory::make(param.type);
        
        if(isArrayType(param.type)){  //Check if the type is array
            std::vector<std::shared_ptr<ParamBase>> vectorParams;
            for(auto& paramValue : param.value){ // Iterate through every data in value
                auto p = ParamFactory::make(getArrayElemType(param.type)); // Create new param of the required type
                if(isNumberType(param.type)){
                    p->setValueJson(toString(load(paramValue)));
                }else {
                    p->setValueJson(hexEncoded(paramValue)); // Set value to the param
                }
                
                vectorParams.push_back(p);
            }
            auto arr = std::make_shared<ParamArray>(vectorParams);  // Cast the parameter to type Array

            abiParam = arr;
        }else{
           if(isNumberType(param.type)){
                abiParam->setValueJson(toString(load(param.value[0])));
            }else {
                abiParam->setValueJson(hexEncoded(param.value[0])); // Set value to the param
            }     
        }
        parameters.push_back(abiParam);
    }

    auto func = Function(functionName, parameters);
    Data payload;
    func.encode(payload);
    return payload;
}


} // namespace TW::Ethereum::ABI
