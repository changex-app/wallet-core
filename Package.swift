// swift-tools-version:5.3
import PackageDescription

let package = Package(
    name: "WalletCore",
    platforms: [.iOS(.v13)],
    products: [
        .library(name: "WalletCore", targets: ["WalletCore"]),
        .library(name: "SwiftProtobuf", targets: ["SwiftProtobuf"])
    ],
    dependencies: [],
    targets: [
        .binaryTarget(
            name: "WalletCore",
            url: "https://github.com/ValeriD/wallet-core-1/releases/download/2.7.1/WalletCore.xcframework.zip",
            checksum: "d45ecac43d3c449b43152f70dbb95192f80ba0a481aec4a576358238f0e36fc5"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/ValeriD/wallet-core-1/releases/download/2.7.1/SwiftProtobuf.xcframework.zip",
            checksum: "208b2bcd4a09ff27fb290ba6c501c74060f06115ee3321d6b5f53bf9c06e3481"
        )
    ]
)
