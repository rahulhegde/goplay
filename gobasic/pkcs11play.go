package main

import (
	"fmt"
	"os"
	"strconv"

	pkcs11 "github.com/miekg/pkcs11"
)

func main() {
	//p := pkcs11.New("/usr/lib/softhsm/libsofthsm2.so")
	var p *pkcs11.Ctx
	lib_path := os.Getenv("TEST_LIB_PATH")
	if len(lib_path) != 0 {
		p = pkcs11.New(lib_path)
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

	token_info, _ := p.GetTokenInfo(slots[0])
	fmt.Printf("Token Info: %+v", token_info)
	fmt.Println()

	session, err := p.OpenSession(slots[0], pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
	if err != nil {
		panic(err)
	}
	defer p.CloseSession(session)

	session_count := os.Getenv("TEST_SESSION_COUNT")
	session_count_int, err := strconv.Atoi(session_count)
	sessions := make([]pkcs11.SessionHandle, session_count_int)

	fmt.Printf("session opening: %d, ", session_count_int)
	for i := 0; i < session_count_int; i++ {
		sessions[i], err = p.OpenSession(slots[0], pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
		if err != nil {
			panic(err)
		}
		fmt.Printf(" %d ", i)
	}

	fmt.Println()
	token_info, _ = p.GetTokenInfo(slots[0])
	fmt.Printf("After %d Session Open: %+v", session_count_int, token_info)
	fmt.Println()

	cku_user := os.Getenv("TEST_CKU_USER")
	err = p.Login(sessions[0], pkcs11.CKU_USER, cku_user)
	if err != nil {
		panic(err)
	}
	p.Logout(sessions[0])

	fmt.Printf("session closing: %d, ", session_count_int)
	for i := 0; i < session_count_int; i++ {
		err = p.CloseSession(sessions[i])
		if err != nil {
			panic(err)
		}
		fmt.Printf(" %d ", i)
	}

	fmt.Println()
	token_info, _ = p.GetTokenInfo(slots[0])
	fmt.Printf("After %d Session Close: %+v", session_count_int, token_info)
	fmt.Println()

	fmt.Println()
}
