package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/test/integration"

	//"github.com/hyperledger/fabric/common/cauthdsl"

	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"

	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
	chmgmt "github.com/hyperledger/fabric-sdk-go/api/apitxn/chmgmtclient"
	resmgmt "github.com/hyperledger/fabric-sdk-go/api/apitxn/resmgmtclient"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/ccpackager/gopackager"
)

const (
	channelID = "mychannel"
	orgName   = "Org1"
	orgAdmin  = "Admin"
	ccID      = "example_cc"
)

func main() {
	fmt.Println("fabric go sdk client")
	var configOpt apiconfig.ConfigProvider
	configOpt = config.FromFile("/home/developer/workspace/go-ws/src/github.com/rahulhegde/goplay/gofabric/fabric_client_config.yaml")

	sdk, err := fabsdk.New(configOpt)
	if err != nil {
		fmt.Println("Failed to create new SDK: ", err.Error())
		return
	}
	fmt.Println("[success] sdk object: %+V", sdk)

	chMgmtClient, err := sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg("ordererorg")).ChannelMgmt()
	if err != nil {
		fmt.Println("Failed to create channel mgmt: ", err.Error())
		return
	}

	session, err := sdk.NewClient(fabsdk.WithUser(orgAdmin), fabsdk.WithOrg(orgName)).Session()
	if err != nil {
		fmt.Println("Failed to create channel mgmt: ", err.Error())
		return
	}

	// // Create channel
	req := chmgmt.SaveChannelRequest{ChannelID: channelID, ChannelConfig: "/home/developer/workspace/go-ws/src/github.com/rahulhegde/goplay/gofabric/crypto/fabric/v1.0/channel/mychannel.tx", SigningIdentity: session}
	if err = chMgmtClient.SaveChannel(req); err != nil {
		fmt.Println(err)
		return
	}

	// Allow orderer to process channel creation
	time.Sleep(time.Second * 5)

	// Org resource management client
	orgResMgmt, err := sdk.NewClient(fabsdk.WithUser(orgAdmin)).ResourceMgmt()
	if err != nil {
		fmt.Println("Failed to create new resource management client: ", err)
		return
	}

	// Org peers join channel
	if err = orgResMgmt.JoinChannel(channelID); err != nil {
		fmt.Println("Org peers failed to JoinChannel: ", err.Error())
		return
	}

	// Create chaincode package for example cc
	ccPkg, err := packager.NewCCPackage("example_cc", "/home/developer/workspace/go-ws/src/github.com/rahulhegde/goplay/gofabric/gopath")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Install example cc to org peers
	installCCReq := resmgmt.InstallCCRequest{Name: ccID, Path: "example_cc", Version: "1.0.0", Package: ccPkg}
	_, err = orgResMgmt.InstallCC(installCCReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set up chaincode policy
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"Org1MSP"})

	// Org resource manager will instantiate 'example_cc' on channel
	err = orgResMgmt.InstantiateCC(channelID, resmgmt.InstantiateCCRequest{Name: ccID, Path: "example_cc", Version: "1.0.0", Args: integration.ExampleCCInitArgs(), Policy: ccPolicy})
	if err != nil {
		fmt.Println(err)
		return
	}

	// business invoke functions

	chClient, err := sdk.NewClient(fabsdk.WithUser("User1")).Channel(channelID)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Release all channel client resources
	defer chClient.Close()

	response, err := chClient.Query(chclient.Request{ChaincodeID: ccID, Fcn: "invoke", Args: integration.ExampleCCQueryArgs()})
	if err != nil {
		fmt.Println(err)
		return
	}
	value := response.Payload

	eventID := "test([a-zA-Z]+)"

	// Register chaincode event (pass in channel which receives event details when the event is complete)
	notifier := make(chan *chclient.CCEvent)
	rce, err := chClient.RegisterChaincodeEvent(notifier, ccID, eventID)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Move funds
	response, err = chClient.Execute(chclient.Request{ChaincodeID: ccID, Fcn: "invoke", Args: integration.ExampleCCTxArgs()})
	if err != nil {
		fmt.Println(err)
		return
	}

	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %s\n", ccEvent)
	case <-time.After(time.Second * 20):
		fmt.Printf("Did NOT receive CC event for eventId(%s)\n", eventID)
		return
	}

	// Unregister chain code event using registration handle
	err = chClient.UnregisterChaincodeEvent(rce)
	if err != nil {
		fmt.Println("Unregister cc event failed: ", err)
		return
	}

	// Verify move funds transaction result
	response, err = chClient.Query(chclient.Request{ChaincodeID: ccID, Fcn: "invoke", Args: integration.ExampleCCQueryArgs()})
	if err != nil {
		fmt.Printf("Failed to query funds after transaction: %s", err)
		return
	}

	valueInt, _ := strconv.Atoi(string(value))
	valueAfterInvokeInt, _ := strconv.Atoi(string(response.Payload))
	if valueInt+1 != valueAfterInvokeInt {
		fmt.Printf("Execute failed. Before: %s, after: %s", value, response.Payload)
		return
	}
	fmt.Println("value before invoke: ", valueInt, " value after invoke: ", valueAfterInvokeInt)
}
