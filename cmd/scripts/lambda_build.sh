#!/bin/bash

# --- Configuration ---
BINARY_NAME="bootstrap"  # AWS AL2/AL2023 runtime ke liye binary ka naam 'bootstrap' hona zaroori hai
SOURCE_FILE="cmd/lambda/main.go"
ZIP_FILE="deployment.zip"

echo "=========================================="
echo "🚀 Starting Build for AWS Lambda..."
echo "=========================================="

# 1. Clean old files
echo "🧹 Cleaning up old builds..."
rm -f $BINARY_NAME
rm -f $ZIP_FILE

# 2. Set Environment Variables for Cross-Compilation
# Note: Tumne GOARCH=arm64 rakha hai. 
# Make sure AWS Lambda create karte waqt tumne "Architecture: arm64" select kiya ho.
# Agar "x86_64" select kiya hai to niche arm64 ko hata kar amd64 kar dena.
echo "🛠  Compiling Go binary..."
export GOOS=linux
export GOARCH=arm64
export CGO_ENABLED=0

# 3. Build command
go build -o $BINARY_NAME $SOURCE_FILE

# Check if build was successful
if [ $? -ne 0 ]; then
    echo "❌ Build Failed! Please check your code."
    exit 1
fi

echo "✅ Build Successful!"

# 4. Zip the binary
echo "📦 Zipping the binary..."

# Check if 'zip' command exists
if command -v zip >/dev/null 2>&1; then
    # Agar zip command hai (Linux/Mac/Some Windows setups)
    zip -j $ZIP_FILE $BINARY_NAME
else
    # Agar zip command nahi hai (Windows Default) -> Use PowerShell
    echo "⚠️  'zip' command not found. Using Windows PowerShell..."
    powershell.exe -nologo -noprofile -command "Compress-Archive -Path $BINARY_NAME -DestinationPath $ZIP_FILE -Force"
fi

# Check if zip was created successfully
if [ ! -f "$ZIP_FILE" ]; then
    echo "❌ Zip Failed! Could not create deployment.zip"
    exit 1
fi

# 5. Clean up binary (optional - keeps folder clean)
rm -f $BINARY_NAME

echo "=========================================="
echo "🎉 SUCCESS! File ready to upload: $ZIP_FILE"
echo "=========================================="