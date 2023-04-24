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
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.29/WalletCore.xcframework.zip",
            checksum: "d54ff5d99b93cdf1993fa413ff85a384695cda2c5a08684587245367485e2e49"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.29/SwiftProtobuf.xcframework.zip",
            checksum: "c357fd4ffcfb6699afc5a2ac082d02cec1b485618916809048f3af37ad5c1a2e"
        )
    ]
)
