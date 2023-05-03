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
            checksum: "a99a91d816ecce26e7f851bcf39f8c5e703e14070237a890f00f853d9b33eeeb"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.31/SwiftProtobuf.xcframework.zip",
            checksum: "ccd8156c2e1aba41560764031f581ce49bc41074e7b44a25b08ac0719f96dc48"
        )
    ]
)