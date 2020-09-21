package main

import (
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"

	pkcs11 "github.com/miekg/pkcs11"
)

const (
	privateKeyFlag = true
	publicKeyFlag  = false
)

var (
	oidNamedCurveP224 = asn1.ObjectIdentifier{1, 3, 132, 0, 33}
	oidNamedCurveP256 = asn1.ObjectIdentifier{1, 2, 840, 10045, 3, 1, 7}
	oidNamedCurveP384 = asn1.ObjectIdentifier{1, 3, 132, 0, 34}
	oidNamedCurveP521 = asn1.ObjectIdentifier{1, 3, 132, 0, 35}
)

// GenerateKey .
func GenerateKey(p11lib *pkcs11.Ctx, session pkcs11.SessionHandle) {
	fmt.Println("Key Generation Start")

	// Now generate a new EC Key, which should be using the uniqueAltId
	var oid asn1.ObjectIdentifier
	oid = oidNamedCurveP256

	ephemeral := false
	keylabel := "TestKey"

	marshaledOID, err := asn1.Marshal(oid)
	if err != nil {
		panic(fmt.Errorf("Could not marshal OID [%s]", err.Error()))
	}

	pubkeyT := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_EC),
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PUBLIC_KEY),
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, !ephemeral),
		pkcs11.NewAttribute(pkcs11.CKA_VERIFY, true),
		pkcs11.NewAttribute(pkcs11.CKA_EC_PARAMS, marshaledOID),

		pkcs11.NewAttribute(pkcs11.CKA_ID, []byte(keylabel)),
		pkcs11.NewAttribute(pkcs11.CKA_LABEL, keylabel),
	}

	prvkeyT := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_EC),
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PRIVATE_KEY),
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, !ephemeral),
		pkcs11.NewAttribute(pkcs11.CKA_PRIVATE, true),
		pkcs11.NewAttribute(pkcs11.CKA_SIGN, true),

		pkcs11.NewAttribute(pkcs11.CKA_ID, []byte(keylabel)),
		pkcs11.NewAttribute(pkcs11.CKA_LABEL, keylabel),

		pkcs11.NewAttribute(pkcs11.CKA_EXTRACTABLE, false),
		pkcs11.NewAttribute(pkcs11.CKA_SENSITIVE, true),
	}

	pub, _, err := p11lib.GenerateKeyPair(session,
		[]*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_EC_KEY_PAIR_GEN, nil)},
		pubkeyT, prvkeyT)

	if err != nil {
		panic(fmt.Errorf("P11: keypair generate failed [%s]", err))
	}

	fmt.Println("Key Generation Complete: ", pub)

	// ecpt, _, err := ecPoint(p11lib, session, pub)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("Error querying EC-point: [%s]", err)
	// }
	// hash := sha256.Sum256(ecpt)
	// ski = hash[:]

	// if updateSKI {
	// 	// set CKA_ID of the both keys to SKI(public key) and CKA_LABEL to hex string of SKI
	// 	setskiT := []*pkcs11.Attribute{
	// 		pkcs11.NewAttribute(pkcs11.CKA_ID, ski),
	// 		pkcs11.NewAttribute(pkcs11.CKA_LABEL, hex.EncodeToString(ski)),
	// 	}

	// 	logger.Infof("Generated new P11 key, SKI %x\n", ski)
	// 	err = p11lib.SetAttributeValue(session, pub, setskiT)
	// 	if err != nil {
	// 		return nil, nil, fmt.Errorf("P11: set-ID-to-SKI[public] failed [%s]", err)
	// 	}

	// 	err = p11lib.SetAttributeValue(session, prv, setskiT)
	// 	if err != nil {
	// 		return nil, nil, fmt.Errorf("P11: set-ID-to-SKI[private] failed [%s]", err)
	// 	}

	// 	//Set CKA_Modifible to false for both public key and private keys
	// 	if csp.immutable {
	// 		setCKAModifiable := []*pkcs11.Attribute{
	// 			pkcs11.NewAttribute(pkcs11.CKA_MODIFIABLE, false),
	// 		}

	// 		_, pubCopyerror := p11lib.CopyObject(session, pub, setCKAModifiable)
	// 		if pubCopyerror != nil {
	// 			return nil, nil, fmt.Errorf("P11: Public Key copy failed with error [%s] . Please contact your HSM vendor", pubCopyerror)
	// 		}

	// 		pubKeyDestroyError := p11lib.DestroyObject(session, pub)
	// 		if pubKeyDestroyError != nil {
	// 			return nil, nil, fmt.Errorf("P11: Public Key destroy failed with error [%s]. Please contact your HSM vendor", pubCopyerror)
	// 		}

	// 		_, prvCopyerror := p11lib.CopyObject(session, prv, setCKAModifiable)
	// 		if prvCopyerror != nil {
	// 			return nil, nil, fmt.Errorf("P11: Private Key copy failed with error [%s]. Please contact your HSM vendor", prvCopyerror)
	// 		}
	// 		prvKeyDestroyError := p11lib.DestroyObject(session, prv)
	// 		if pubKeyDestroyError != nil {
	// 			return nil, nil, fmt.Errorf("P11: Private Key destroy failed with error [%s]. Please contact your HSM vendor", prvKeyDestroyError)
	// 		}
	// 	}
	// }

	// nistCurve := namedCurveFromOID(curve)
	// if curve == nil {
	// 	return nil, nil, fmt.Errorf("Cound not recognize Curve from OID")
	// }
	// x, y := elliptic.Unmarshal(nistCurve, ecpt)
	// if x == nil {
	// 	panic(fmt.Errorf("Failed Unmarshaling Public Key"))
	// }

	//pubGoKey := &ecdsa.PublicKey{Curve: nistCurve, X: x, Y: y}

	// return ski, pubGoKey, nil
}

func findKeyPairFromSKI(mod *pkcs11.Ctx, session pkcs11.SessionHandle, ski []byte, altID string, keyType bool) (*pkcs11.ObjectHandle, error) {
	ktype := pkcs11.CKO_PUBLIC_KEY
	if keyType == privateKeyFlag {
		ktype = pkcs11.CKO_PRIVATE_KEY
	}

	keyId := []byte(altID)
	if altID == "" {
		keyId = ski
	}

	template := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, ktype),
		pkcs11.NewAttribute(pkcs11.CKA_ID, keyId),
	}

	if err := mod.FindObjectsInit(session, template); err != nil {
		return nil, err
	}

	// single session instance, assume one hit only
	objs, _, err := mod.FindObjects(session, 1)
	if err != nil {
		return nil, err
	}
	if err = mod.FindObjectsFinal(session); err != nil {
		return nil, err
	}

	if len(objs) == 0 {
		return nil, fmt.Errorf("Key not found [%s]", hex.Dump(ski))
	}

	return &objs[0], nil
}

// DoFabricOperation will simulate fabric operations
func DoFabricOperation(p *pkcs11.Ctx, session pkcs11.SessionHandle, altID string, counter int, waiter chan int) {

	fmt.Println("Inside GO Routine: ", counter)

	fabricOpsCount := os.Getenv("TEST_FABRIC_OPS_COUNT")
	fabricOpsCountInt, _ := strconv.Atoi(fabricOpsCount)
	fabricOpsDelayMS := os.Getenv("TEST_FABRIC_OPS_DELAY")
	fabricOpsDelayMSInt, _ := strconv.Atoi(fabricOpsDelayMS)

	for index := 0; index < fabricOpsCountInt; index++ {
		privateKey, err := findKeyPairFromSKI(p, session, []byte("ski"), altID, privateKeyFlag)
		if err != nil {
			panic(fmt.Errorf("Private key not found [%s]", err))
		}

		err = p.SignInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_ECDSA, nil)}, *privateKey)
		if err != nil {
			panic(fmt.Errorf("Sign-initialize  failed [%s]", err))
		}

		//var sig []byte
		hash := sha256.New()
		hash.Write([]byte("hello world\n"))
		_, err = p.Sign(session, hash.Sum(nil))
		if err != nil {
			panic(fmt.Errorf("P11: sign failed [%s]", err))
		}
		// fmt.Print("signed data %+v", sig)
		fmt.Print(".")
		time.Sleep(time.Duration(fabricOpsDelayMSInt) * time.Millisecond)
	}
	waiter <- counter
}

// TestFabricOperation performs fabric alike operations
//	Here are configuration params -
// export TEST_LIB_PATH=/usr/lib/softhsm/libsofthsm2.so
// export TEST_CKU_USER=9876
// export TEST_ALT_ID=TestKey
// export TEST_SESSION_COUNT=10
// export TEST_THREAD_COUNT=10
// export TEST_FABRIC_OPS_COUNT=10
// export TEST_FABRIC_OPS_DELAY=100
// export TEST_GENERATE_KEY=True
func TestFabricOperation() {
	var p *pkcs11.Ctx

	libPath := os.Getenv("TEST_LIB_PATH")
	if len(libPath) != 0 {
		p = pkcs11.New(libPath)
	} else {
		p = pkcs11.New("/usr/lib/softhsm/libsofthsm2.so")
	}

	err := p.Initialize()
	if err != nil {
		panic(err)
	}

	defer p.Destroy()
	defer p.Finalize()

	slots, err := p.GetSlotList(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Slot count: %+v", slots)
	fmt.Println()

	tokenInfo, _ := p.GetTokenInfo(slots[0])
	fmt.Printf("Token Info: %+v", tokenInfo)
	fmt.Println()

	sessionCount := os.Getenv("TEST_SESSION_COUNT")
	sessionCountInt, err := strconv.Atoi(sessionCount)
	sessions := make([]pkcs11.SessionHandle, sessionCountInt)
	fmt.Printf("session opening: %d, ", sessionCountInt)
	for index := 0; index < sessionCountInt; index++ {
		sessions[index], err = p.OpenSession(slots[0], pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
		if err != nil {
			panic(err)
		}
		fmt.Printf(" %d ", index)

		user := os.Getenv("TEST_CKU_USER")
		err = p.Login(sessions[index], pkcs11.CKU_USER, user)
		if err != nil {
			if err != pkcs11.Error(pkcs11.CKR_USER_ALREADY_LOGGED_IN) {
				panic(fmt.Errorf("Login failed [%s]", err))
			}
		}
	}
	fmt.Println()

	generateKey := os.Getenv("TEST_GENERATE_KEY")
	if generateKey == "True" {
		GenerateKey(p, sessions[0])
	}

	threadCount := os.Getenv("TEST_THREAD_COUNT")
	threadCountInt, err := strconv.Atoi(threadCount)
	waiter := make(chan int, threadCountInt)
	sessionIndex := 0
	for index := 0; index < threadCountInt; index++ {
		altID := os.Getenv("TEST_ALT_ID")
		go DoFabricOperation(p, sessions[sessionIndex], altID, index, waiter)
		if sessionIndex == sessionCountInt-1 {
			sessionIndex = 0
		} else {
			sessionIndex++
		}
	}

	counter := 0
	fmt.Println("Go Routing Checker")
	for {
		if counter == threadCountInt {
			fmt.Println("All Go Routing Call Complete: ", counter)
			break
		}

		select {
		case threadID := <-waiter:
			fmt.Println("Go Routing Complete: ", threadID)
			counter++
		}
	}

	fmt.Print("session closing: ", sessionCountInt, ", ")
	for index := 0; index < sessionCountInt; index++ {
		err = p.CloseSession(sessions[index])
		if err != nil {
			panic(err)
		}
		fmt.Print(" ", index, " ")
	}
	fmt.Println()
}

func main() {
	TestFabricOperation()
}
