package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/wowsims/tbc/api"
	// "github.com/wowsims/tbc/dist"
	"google.golang.org/protobuf/proto"
)

func main() {
	var useFS = flag.Bool("usefs", true, "Use local file system and wasm. Set to true for dev")
	// TODO: usefs for now is set to true until we can solve how to embed the dist.
	var host = flag.String("host", ":3333", "URL to host the interface on.")

	flag.Parse()

	var fs http.Handler
	if *useFS {
		log.Printf("Using local file system for development.")
		fs = http.FileServer(http.Dir("./dist"))
	} else {
		log.Printf("Embedded file server running.")
		// fs = http.FileServer(http.FS(dist.FS))
	}

	http.HandleFunc("/statWeights", handleStatWeights)
	http.HandleFunc("/computeStats", handleComputeStats)
	http.HandleFunc("/individualSim", handleIndividualSim)
	http.HandleFunc("/gearList", handleGearList)

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	})

	url := fmt.Sprintf("http://localhost%s/elemental_shaman/", *host)
	log.Printf("Launching interface on %s", url)

	go func() {
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("explorer", url)
		} else if runtime.GOOS == "darwin" {
			cmd = exec.Command("open", url)
		} else if runtime.GOOS == "linux" {
			cmd = exec.Command("xdg-open", url)
		}
		err := cmd.Start()
		if err != nil {
			log.Printf("Error launching browser: %#v", err.Error())
		}
		log.Printf("Closing: %s", http.ListenAndServe(*host, nil))
	}()

	fmt.Printf("Enter Command... '?' for list\n")
	for {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if len(text) == 0 {
			continue
		}
		command := strings.TrimSpace(text)
		switch command {
		case "profile":
			filename := fmt.Sprintf("profile_%d.cpu", time.Now().Unix())
			fmt.Printf("Running profiling for 15 seconds, output to %s\n", filename)
			f, err := os.Create(filename)
			if err != nil {
				log.Fatal("could not create CPU profile: ", err)
			}
			if err := pprof.StartCPUProfile(f); err != nil {
				log.Fatal("could not start CPU profile: ", err)
			}
			go func() {
				time.Sleep(time.Second * 15)
				pprof.StopCPUProfile()
				f.Close()
				fmt.Printf("Profiling complete.\n> ")
			}()
		case "quit":
			os.Exit(1)
		case "?":
			fmt.Printf("Commands:\n\tprofile - start a CPU profile for debugging performance\n\tquit - exits\n\n")
		case "":
			// nothing.
		default:
			fmt.Printf("Unknown command: '%s'", command)
		}
	}
}

// result := GetGearList(*request.GearList)
// return ApiResult{GearList: &result}
// } else if request.ComputeStats != nil {
// result := ComputeStats(*request.ComputeStats)
// return ApiResult{ComputeStats: &result}
// } else if request.StatWeights != nil {
// result := StatWeights(*request.StatWeights)
// return ApiResult{StatWeights: &result}
// } else if request.Sim != nil {

func handleIndividualSim(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		return
	}
	isr := &api.IndividualSimRequest{}
	if err := proto.Unmarshal(body, isr); err != nil {
		log.Printf("Failed to parse request: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := api.RunSimulation(isr)

	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/x-protobuf")
	w.Write(outbytes)
}
func handleRaidSim(w http.ResponseWriter, r *http.Request) {

}
func handleGearList(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		return
	}
	glr := &api.GearListRequest{}
	if err := proto.Unmarshal(body, glr); err != nil {
		log.Printf("Failed to parse request: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := api.GetGearList(glr)

	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/x-protobuf")
	w.Write(outbytes)
}
func handleComputeStats(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	csr := &api.ComputeStatsRequest{}
	if err := proto.Unmarshal(body, csr); err != nil {
		log.Printf("Failed to parse request: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := api.ComputeStats(csr)
	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/x-protobuf")
	w.Write(outbytes)
}
func handleStatWeights(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	swr := &api.StatWeightsRequest{}
	if err := proto.Unmarshal(body, swr); err != nil {
		log.Printf("Failed to parse request: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := api.StatWeights(swr)
	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/x-protobuf")
	w.Write(outbytes)
}
