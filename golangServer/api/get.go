package api

import (
	"Oracle.com/golangServer/Oracle"
	"Oracle.com/golangServer/config"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/printer"
	"log"
	"regexp"
)

func GetEventListener() {
	// Get channels for logs
	Logs := make(chan *Oracle.OracleGet)
	// Subscribe to each event
	opts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	eventSub, err := config.OracleContract.WatchGet(opts, Logs)
	if err != nil {
		log.Fatal("Failed to subscribe to Get events:", err)
	}
	// start Listening...
	log.Println("GetEvent Listening ...")
	for {
		select {
		case err := <-eventSub.Err():
			log.Println("[Error in Event CREAT]:", err)
		case event := <-Logs:
			log.Println("Received get event ", event.ReqID)
			get(event)
		}
	}
}

// Get data from memory db
func get(event *Oracle.OracleGet) {
	var statement bool
	tps := GenTransactOpts(config.GasLimit)

	dbName := event.DbName
	ctx, dbC := config.Dbs[dbName].Ctx, config.Dbs[dbName].Db
	node, err := dbC.Get(ctx, event.RecordID)
	if err != nil {
		log.Println("[", event.DbName, "]", "Get data ERROR: ", err)
		info := fmt.Sprintf("Insert ERROR: %v", err)
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, info)
		return
	}
	if node == nil {
		statement = false
		//response to oracle
		config.OracleContract.GetRsp(tps, event.ReqID, statement, []byte{}, event.CallBack, event.Sender, "Data is not exist")
		return
	}
	data := nodeTobyte(node)
	statement = true
	//response to oracle
	config.OracleContract.GetRsp(tps, event.ReqID, statement, data, event.CallBack, event.Sender, "")
	log.Println("[", event.DbName, "]", "Get data success")
}

func toJSONstr(k []string, v []string) string {
	str := "{"
	for i := 0; i < len(k); i++ {
		Regex := regexp.MustCompile(`{(.*)}`)
		kMatches := Regex.FindStringSubmatch(k[i])
		vMatches := Regex.FindStringSubmatch(v[i])
		str = fmt.Sprint(str, kMatches[1], ":", vMatches[1])
		if i != len(k)-1 {
			str = fmt.Sprint(str, ", ")
		}
	}
	str = fmt.Sprint(str, "}")
	return str
}

func nodeTobyte(node ipld.Node) []byte {
	var data []byte
	kind := node.Kind()
	switch kind {
	case datamodel.Kind_Map:
		ite := node.MapIterator()
		k := make([]string, 0, node.Length())
		v := make([]string, 0, node.Length())
		for !ite.Done() {
			key, value, _ := ite.Next()
			k = append(k, printer.Sprint(key))
			v = append(v, printer.Sprint(value))
		}
		data = []byte(toJSONstr(k, v))
	case datamodel.Kind_List:
		ite := node.ListIterator()
		k := make([]string, 0, node.Length())
		v := make([]string, 0, node.Length())
		for !ite.Done() {
			key, value, _ := ite.Next()
			k = append(k, fmt.Sprint("{\"idx\"}"))
			v = append(v, fmt.Sprint("{", key, "}"))
			k = append(k, fmt.Sprint("{\"node\"}"))
			v = append(v, printer.Sprint(value))
		}
		data = []byte(toJSONstr(k, v))
	case datamodel.Kind_Bool:
		k := make([]string, 0, node.Length())
		v := make([]string, 0, node.Length())
		k = append(k, fmt.Sprint("{\"bool\"}"))
		n, _ := node.AsBool()
		v = append(v, fmt.Sprint("{", n, "}"))
		data = []byte(toJSONstr(k, v))
	case datamodel.Kind_Int:
		k := make([]string, 0, node.Length())
		v := make([]string, 0, node.Length())
		k = append(k, fmt.Sprint("{\"int\"}"))
		n, _ := node.AsInt()
		v = append(v, fmt.Sprint("{", n, "}"))
		data = []byte(toJSONstr(k, v))
	case datamodel.Kind_Float:
		k := make([]string, 0, node.Length())
		v := make([]string, 0, node.Length())
		k = append(k, fmt.Sprint("{\"float\"}"))
		n, _ := node.AsFloat()
		v = append(v, fmt.Sprint("{", n, "}"))
		data = []byte(toJSONstr(k, v))
	case datamodel.Kind_String:
		k := make([]string, 0, node.Length())
		v := make([]string, 0, node.Length())
		k = append(k, fmt.Sprint("{\"string\"}"))
		n, _ := node.AsString()
		v = append(v, fmt.Sprint("{\"", n, "\"}"))
		data = []byte(toJSONstr(k, v))
	case datamodel.Kind_Bytes:
		k := make([]string, 0, node.Length())
		v := make([]string, 0, node.Length())
		k = append(k, fmt.Sprint("{\"bytes\"}"))
		n, _ := node.AsBytes()
		v = append(v, fmt.Sprint("{", n, "}"))
		data = []byte(toJSONstr(k, v))
	case datamodel.Kind_Link:
		k := make([]string, 0, node.Length())
		v := make([]string, 0, node.Length())
		k = append(k, fmt.Sprint("{\"link\"}"))
		n, _ := node.AsLink()
		v = append(v, fmt.Sprint("{\"", n.String(), "\"}"))
		data = []byte(toJSONstr(k, v))
	}
	return data
}
