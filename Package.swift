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
            url: "https://github.com/ValeriD/wallet-core-1/releases/download/2.6.46/WalletCore.xcframework.zip",
            checksum: "018257a32d1bf1f68290eeb5b028d42b5a4dfcf168fd5d41ad599a8970801074"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/ValeriD/wallet-core-1/releases/download/2.6.46/SwiftProtobuf.xcframework.zip",
            checksum: "c6414e4cce760523592ebba6735b1ca71ea1072f60d49110343e401ce0e5341a"
        )
    ]
)
