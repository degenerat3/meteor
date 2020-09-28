package main

import (
	"crypto/sha1"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"os"
	"strings"
	"time"
)

func initAdmin() {
	adminpw := os.Getenv("ADMIN_PW")
	hasher := sha1.New()
	hasher.Write([]byte(adminpw))
	encpw := string(hasher.Sum(nil))
	_, err := DBClient.User.Create().SetUsername("admin").SetPassword(encpw).Save(ctx)
	if err != nil {
		panic(err)
	}
}

func genUnAuth() []byte {
	resp := &mcs.MCS{
		Status: 401,
		Desc:   "Invalid user or password\n",
	}
	rdata, _ := proto.Marshal(resp)
	return rdata
}

func newSession() string {
	token := genRando()
	exp := time.Now().Unix() + 600
	ses := session{
		Token: token,
		Exp:   exp,
	}
	sessions = append(sessions, ses)
	return token
}

func checkAuth(prot *mcs.MCS) bool {
	authDat := prot.GetAuthDat()
	tok := authDat.GetToken()
	v := validateToken(tok)
	return v
}

func validateToken(tok string) bool {
	for _, ses := range sessions {
		if tok == ses.Token {
			curr := time.Now().Unix()
			if curr < ses.Exp {
				return true
			}
		}
	}
	return false
}

func refreshSession(tok string) {
	for _, ses := range sessions {
		if tok == ses.Token {
			newExp := time.Now().Unix() + 600
			ses.Exp = newExp
		}
	}
}

func genRando() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789")
	length := 20
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}