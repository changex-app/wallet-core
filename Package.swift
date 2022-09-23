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
            url: "https://github.com/ValeriD/wallet-core-1/releases/download/2.7.0/WalletCore.xcframework.zip",
            checksum: "847b1835876b7858ce4180204188ae5c486ab4587327fe5958f50b341a996800"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/ValeriD/wallet-core-1/releases/download/2.7.0/SwiftProtobuf.xcframework.zip",
            checksum: "326d358149e1c22b25e8dbf5d027f362f90c7ae665e55fa758f411fd33468867"
        )
    ]
)
