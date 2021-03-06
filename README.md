# Azure template generator for ACC VMs

### Supported VMs
- Trusted VM (TVM)
- Confidential VM (CVM)

#### Build
If your machine runs Linux or MacOS, and you have Go binary installed, execute:
```sh
make build
```
Alternatively, you can build with a Docker container:
```sh
docker run \
  -v $PWD:/go/src/github.com/microsoft/acc-vm-engine \
  -w /go/src/github.com/microsoft/acc-vm-engine \
  golang:1.15-alpine go build ./cmd/acc-vm-engine
```

#### Generate template
A sample configuration file for TVM deployment is located in [test/cvm-win.json](./test/cvm-win.json)
On Linux or MacOS, execute:
```sh
./acc-vm-engine generate -c ./test/cvm-win.json
```
Alternatively, use Docker container:
```sh
docker run \
  -v $PWD:/go/src/github.com/microsoft/acc-vm-engine \
  -w /go/src/github.com/microsoft/acc-vm-engine \
  golang:1.15-alpine ./acc-vm-engine generate -c ./test/cvm-win.json
```
The template and parameter files will be generated in `_output` directory (by default).

#### Deploy the VM
Use [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) or PowerShell to deploy the VM.

When using Azure CLI, you may want to log in to Azure and set default subscription. This is a one-time operation:
```sh
SUB=<subscription ID>

az login

az account set --subscription ${SUB}
```

Create a resource group and deploy the VM:
```sh
RGROUP=<resource group name>
LOC=<deployment region>

az group create -n ${RGROUP} -l ${LOC}

az deployment group create \
  --resource-group ${RGROUP} \
  --name MyDeployment \
  --template-file ./_output/azuredeploy.json \
  --parameters @./_output/azuredeploy.parameters.json
```
