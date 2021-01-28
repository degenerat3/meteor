package main

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/degenerat3/meteor/core/ent/user"
	"github.com/degenerat3/meteor/pbuf"
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
	encpw := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	_, err := DBClient.User.Query().Where(user.Username("admin")).Only(ctx)
	if err != nil {
		_, err := DBClient.User.Create().SetUsername(string("admin")).SetPassword(encpw).Save(ctx)
		if err != nil {
			panic(err)
		}
	}
}

func genUnAuth() []byte {
	resp := &mcs.MCS{
		Status: 401,
		Desc:   "Invalid user or password",
	}
	rdata, _ := proto.Marshal(resp)
	return rdata
}

func genInvalidTok() []byte {
	resp := &mcs.MCS{
		Status: 401,
		Desc:   "Invalid token\n",
	}
	rdata, _ := proto.Marshal(resp)
	return rdata
}

func newSession(username string) string {
	token := genRando()
	exp := time.Now().Unix() + 600
	ses := session{
		User:  username,
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

func validateUserToken(tok string, user string) bool {
	for _, ses := range sessions {
		if tok == ses.Token {
			if ses.User == user {
				return true
			} else if user == "admin" {
				return true
			}
			return false
		}
	}
	return false
}
