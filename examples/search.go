// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/johnweldon/ldap"
)

var (
	ldapServer string   = "localhost"
	ldapPort   uint16   = 389
	baseDN     string   = "dc=*,dc=*"
	filter     string   = "&(objectClass=*)"
	Attributes []string = []string{"dn", "cn"}
	user       string   = "*"
	passwd     string   = "*"
)

func main() {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}
	defer l.Close()
	// l.Debug = true

	err = l.Bind(user, passwd)
	if err != nil {
		log.Printf("ERROR: Cannot bind: %s\n", err.Error())
		return
	}
	search := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		Attributes,
		nil)

	sr, err := l.Search(search)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
		return
	}

	log.Printf("Search: %s -> num of entries = %d\n", search.Filter, len(sr.Entries))
	sr.PrettyPrint(0)
}
