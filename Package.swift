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
            checksum: "14d98632dcc42944994739c03a6241a77aec65678eb1ceb7ca0aa3eb7b7cb8cd"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.31/SwiftProtobuf.xcframework.zip",
            checksum: "d5904be80cbead2bd38eb0ab3234cdd5b92597f1c88490a850ddd9aa476360b4"
        )
    ]
)