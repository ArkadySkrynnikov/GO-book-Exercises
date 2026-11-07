package main

import (
	"playground/ints"
)

func f(n int) *int {
	v := n
	return &v
}
func incr(p *int) {
	*p++
}

type Tipochek struct {
	one string
	two struct {
		three string
	}
}

func (t Tipochek) Super() (string, error) {
	return "super", nil
}

// 2 6 2
func main() {
	// x := 1
	// p := &x
	// fmt.Println(*p, x)
	// *p = 2
	// fmt.Println(*p, x)
	// pz := f(1)
	// *pz = 228
	// fmt.Println(*pz, *f(*pz))

	// num := 14
	// incr(&num)
	// fmt.Println(num)

	// second.EchoFlags()

	// var tip Tipochek = Tipochek{
	// 	one: "Первый",
	// 	two: struct {
	// 		three string
	// 	}{
	// 		three: "Третий",
	// 	},
	// }
	// tip.two = struct{ three string }{"213"}
	// t, r := tip.Super()
	// fmt.Println(tip, t, r)

	// bitwises.BitwiseTry(123)

	// logs.LogsTry()
	// second.UnrealMix()
	// second.EchoFlags()
	// second.Poinsters()
	// second.MyFToC()
	// server.Start()
	// fetch_all.FetchAll()
	// my_http.Fetch()
	// if len(os.Args) > 1 && os.Args[1] == "web" {
	// 	handler := func(w http.ResponseWriter, r *http.Request) {
	// 		gifs.Lissajous(w)
	// 	}
	// 	http.HandleFunc("/", handler)
	// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	// 	return
	// }
	// gifs.Lissajous(os.Stdout)
	ints.Run()
}
