package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rpcx-ecosystem/rpcx-demo/service/product/model"
	"github.com/smallnest/rpcx/client"
)

var (
	addr  = flag.String("addr", ":8080", "http address")
	paddr = flag.String("product-image-addr", "localhost:8972", "图片服务地址")
)

var (
	xclient client.XClient
)

func main() {
	d := client.NewPeer2PeerDiscovery("tcp@"+*paddr, "")
	xclient = client.NewXClient("ProductImage", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/banner-ppl.png", index)
	router.GET("/banner-ppl-women.png", index)
	router.GET("/bk-sale.png", index)
	router.ServeFiles("/_nuxt/*filepath", http.Dir("../web/dist/_nuxt"))
	router.ServeFiles("/cart/*filepath", http.Dir("../web/dist/cart"))
	router.ServeFiles("/men/*filepath", http.Dir("../web/dist/men"))
	router.ServeFiles("/sale/*filepath", http.Dir("../web/dist/sale"))
	router.ServeFiles("/women/*filepath", http.Dir("../web/dist/women"))
	router.GET("/products_images/:name", productsImages)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("@@@@@:", "../web/dist/"+r.URL.Path[1:])
	http.ServeFile(w, r, "../web/dist/"+r.URL.Path[1:])
}

func productsImages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	resp := &model.ImageResponse{}
	err := xclient.Call(context.Background(), "Get", model.ImageRequest(name), resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h := w.Header()
	h.Set("Context-Type", resp.ContentType)
	h.Set("Context-Length", strconv.Itoa(resp.ContentLength))
	w.Write(resp.Content)
	//http.ServeFile(w, r, "../service/static/"+ps.ByName("name"))
}
