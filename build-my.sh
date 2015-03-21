#!/bin/bash -e

VERSION=0.4-dev

rm -rf bin
gox -osarch darwin/amd64 -output bin/macos/terraform-provider-vsphere
cp bin/macos/terraform-provider-vsphere ./
gox -osarch linux/amd64 -output bin/linux/terraform-provider-vsphere
gox -osarch windows/amd64 -output bin/windows/terraform-provider-vsphere

tar czf bin/terraform-vsphere-$VERSION-macos.tar.gz  --directory=bin/macos terraform-provider-vsphere
tar czf bin/terraform-vsphere-$VERSION-linux.tar.gz  --directory=bin/linux terraform-provider-vsphere
zip     bin/terraform-vsphere-$VERSION-windows.zip   -j bin/windows/terraform-provider-vsphere.exe
