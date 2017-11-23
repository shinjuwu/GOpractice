package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	ServerName string `json:"-"`
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

type Server1 struct {
	ID          int    `json:"-"`
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`
	ServerIP    string `json:"serverIP,omit,empty"`
}

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of type I don't know how to handle")
		}
	}
	//====================================================================
	var s ServerSlice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err = json.Marshal(s)
	bString := string(b[:])
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(b))
	fmt.Println(bString)

	//====================================================================
	s1 := Server1{
		ID:          3,
		ServerName:  `GO "1.0"`,
		ServerName2: `GO "1.0"`,
		ServerIP:    ``,
	}

	bb, _ := json.Marshal(s1)
	os.Stdout.Write(bb)
}
