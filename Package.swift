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
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.35/WalletCore.xcframework.zip",
            checksum: "d48813549dbf6da40fd485ced18f57099d0f7749cd4c86dd44728c07714df0ab"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.35/SwiftProtobuf.xcframework.zip",
            checksum: "5d60b38a4c949757e8ccbc7d154a53d4ce88669f19ca0ddfe38c82cd71e04612"
        )
    ]
)