// package magic/router
package main 
import (
// "log"
// "net/http"
"fmt"
"os"
"io/ioutil"
"encoding/json"
// "strings"
)

type routerRecord struct{
	uri string
	controller string
	action string
	matches string
	rewrite bool
}

type dispatchProcesser func() bool

/**
 the router
*/
type Router struct {
	// routerTable map[string]*routerRecord
	routerTable map[string]dispatchProcesser
}

// push an router rule to the router table
func (this *Router) push(uri string,process dispatchProcesser) {
	// process()
}

// the router init function
func (this *Router) init() {
	// this.routerTable = make(map[string]*routerRecord)
	this.routerTable = make(map[string]dispatchProcesser)

	routeMapFile := "router.json"
    fileResource,err := os.Open(routeMapFile)
    defer fileResource.Close()
    if err != nil {
                fmt.Println(routeMapFile,err)
                // return
    }
   routerContent,err := ioutil.ReadAll(fileResource)

    var dat map[string]interface{}
    if err := json.Unmarshal([]byte(string(routerContent)), &dat); err == nil {
        fmt.Println("==============json str 转map=======================")
        // fmt.Println(dat)
        // fmt.Println(dat["post/#id#"])

        for uri, ctrl := range dat{
        	this.push(uri,func () bool{
        		fmt.Println(ctrl)
        		return true
        	})
        	// fmt.Println(uri,ctrl)
        }
    }


   // fmt.Println(string(routerContent))
}

func (this *Router) dispatch() {
	
}

func (this *Router)match() {
	
}

func main() {
	// recode := &routerRecord{uri:"/",action:"",controller:""}
	router := new(Router)
	router.init()
	// router.routerTable = make(map[string]*routerRecord)
	// router.routerTable["test"]=recode
	// var routerTable = make(map[string] routerRecord{})
	// routerTable["test"] := recode
	// router.routerTable[recode.uri] := recode
	// fmt.Println(router)
	// router.routerTable =new(map[string]routerRecord)
	// router.routerTable = new(map[string]routerRecord)
	fmt.Println(router)
	// router.routerTable[recode.uri] = recode
}
