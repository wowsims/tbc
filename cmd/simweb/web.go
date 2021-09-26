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
	"github.com/wowsims/tbc/api/genapi"
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

	http.HandleFunc("/statWeights", handleAPI)
	http.HandleFunc("/computeStats", handleAPI)
	http.HandleFunc("/individualSim", handleAPI)
	http.HandleFunc("/gearList", handleAPI)

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

type apiHandler struct {
	msg    func() proto.Message
	handle func(proto.Message) proto.Message
}

// Handlers to decode and handle each proto function
var handlers = map[string]apiHandler{
	"/individualSim": {msg: func() proto.Message { return &genapi.IndividualSimRequest{} }, handle: func(msg proto.Message) proto.Message { return api.RunSimulation(msg.(*genapi.IndividualSimRequest)) }},
	"/statWeights":   {msg: func() proto.Message { return &genapi.StatWeightsRequest{} }, handle: func(msg proto.Message) proto.Message { return api.StatWeights(msg.(*genapi.StatWeightsRequest)) }},
	"/computeStats":  {msg: func() proto.Message { return &genapi.ComputeStatsRequest{} }, handle: func(msg proto.Message) proto.Message { return api.ComputeStats(msg.(*genapi.ComputeStatsRequest)) }},
	"/gearList":      {msg: func() proto.Message { return &genapi.GearListRequest{} }, handle: func(msg proto.Message) proto.Message { return api.GetGearList(msg.(*genapi.GearListRequest)) }},
}

// handleAPI is generic handler for any api function using protos.
func handleAPI(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Path

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		return
	}
	handler, ok := handlers[endpoint]
	if !ok {
		log.Printf("Invalid Endpoint: %s", endpoint)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	msg := handler.msg()
	if err := proto.Unmarshal(body, msg); err != nil {
		log.Printf("Failed to parse request: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := handler.handle(msg)

	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/x-protobuf")
	w.Write(outbytes)
}
