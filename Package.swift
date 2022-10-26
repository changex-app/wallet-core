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
            url: "https://github.com/ValeriD/wallet-core-1/releases/download/2.7.2/WalletCore.xcframework.zip",
            checksum: "3b8b12adc5658f93e8bbc043a06629c7c8047dadbf95a7125226ca68f8898b30"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/ValeriD/wallet-core-1/releases/download/2.7.2/SwiftProtobuf.xcframework.zip",
            checksum: "4c95eb3d53b237afa93ae0a97071fdd7da2ea773db668093a97170cc6c42b024"
        )
    ]
)
