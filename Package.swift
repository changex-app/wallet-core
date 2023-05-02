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
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.31/WalletCore.xcframework.zip",
            checksum: "b39f16d60529316c74dd2153e599b0f86fed1d8a51cfc018765da66e8d50b0d5"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.31/SwiftProtobuf.xcframework.zip",
            checksum: "c41bc1e0324696a5c08df1e7eab0cbd926bae3bc41405e618c0b2010bdf50a39"
        )
    ]
)