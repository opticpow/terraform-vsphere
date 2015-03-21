#!/bin/bash -e

export TF_ACC=1
export VSPHERE_SERVER=mk-vcenter
export VSPHERE_USER=root
export VSPHERE_PASSWORD=vmware

go test -v
