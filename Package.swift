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
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.32/WalletCore.xcframework.zip",
            checksum: "3bb7d92db6c9629a7b7fd444a43bdee3179fe7950da13fedd9592b971a67232f"
        ),
        .binaryTarget(
            name: "SwiftProtobuf",
            url: "https://github.com/changex-app/wallet-core/releases/download/3.1.32/SwiftProtobuf.xcframework.zip",
            checksum: "da40459ad1f916c88e0d7a2bed587d795604a4c5ae796b840c81ad15f5e99967"
        )
    ]
)