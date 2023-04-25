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
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.30/WalletCore.xcframework.zip",
            checksum: "d82c3a2225d78f56fb9c04300156778f5c598799838053c57a544104a4ecdc00"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.30/SwiftProtobuf.xcframework.zip",
            checksum: "f47f0ac42e78337f71bc4c2ecf13f766b3caf7a2a08a4563ac2899da80adb2d7"
        )
    ]
)
