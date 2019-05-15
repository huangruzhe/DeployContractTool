package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"DeployContractTool/service"
	"sort"
	"fmt"
)

//节点地址
var nodeUrl = "http://localhost:22000"
//当前服务端口
var listenPort = ":8000"
//the node directory path
var nodePath = "/home/"
//the program directory path
var programPath = "contracts/"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {

	app := cli.NewApp()
	//程序名称
	app.Name = "DCT"
	//程序的用途
	app.Usage = "Deploy Contract Tool(该版本增加部署私有合约)"
	//程序的描述
	app.Description = "这是一个使用Quorum区块链部署合约的工具，基于quorum-maker/quorum-maker-nodemanager和urfave/cli"
	//程序的版本号(beta)
	app.Version = "0.0.2"

	//节点服务地址
	var nodeUrl string
	var contractFile string
	var private string
	var privateFor string
	var nodePath string
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "nodeUrl, n",
			Value: "http://localhost:22000",
			Usage: "节点服务地址",
			Destination: &nodeUrl,
		},
		cli.StringFlag{
			Name: "contractFile, c",
			Usage: "合约文件(例如 /XXX/xxx.sol)",
			Destination: &contractFile,
		},

		cli.StringFlag{
			Name: "private, p",
			Value: "false",
			Usage: "是否为私有合约",
			Destination: &private,
		},

		cli.StringFlag{
			Name: "privateFor, pf",
			Usage: "该参数是PUBKEY,在setup.conf文件中，多个节点用逗号隔开(不可指认自身节点，会报错)",
			Destination: &privateFor,
		},

		cli.StringFlag{
			Name: "nodePath, np",
			Value: "/home/",
			Usage: "节点所在目录（例如 /XXX/node1/）",
			Destination: &nodePath,
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	//合约内容
	app.Action = func(c *cli.Context) error {
		nodeService := service.NodeServiceImpl{nodeUrl, nodePath, programPath}
		nodeService.DeployContractHandler(contractFile, private, privateFor)
		fmt.Println("合约部署完成")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("finish")
	}


	////获取命令参数
	//if len(os.Args) > 1 {
	//	nodeUrl = os.Args[1]
	//}
	////
	//if len(os.Args) > 2 {
	//	listenPort = ":" + os.Args[2]
	//}
	//
	//if len(os.Args) > 3 {
	//	nodePath = os.Args[3]
	//}
	//
	//if len(os.Args) > 4 {
	//	programPath = os.Args[4]
	//}
	//
	//router := mux.NewRouter()
	//nodeService := service.NodeServiceImpl{nodeUrl, nodePath, programPath}
	//
	//ticker := time.NewTicker(86400 * time.Second)
	//go func() {
	//	for range ticker.C {
	//		log.Debug("Rotating log for Geth and Constellation.")
	//		nodeService.LogRotaterGeth(nodePath)
	//		nodeService.LogRotaterConst(nodePath)
	//	}
	//}()
	//
	//go func() {
	//	nodeService.CheckGethStatus(nodeUrl)
	//	//log.Info("Deploying Network Manager Contract")
	//	nodeService.NetworkManagerContractDeployer(nodeUrl, nodePath)
	//	nodeService.RegisterNodeDetails(nodeUrl, nodePath)
	//	nodeService.ContractCrawler(nodeUrl, nodePath, programPath)
	//	nodeService.ABICrawler(nodeUrl, programPath)
	//	nodeService.IPWhitelister()
	//}()
	//
	//networkMapService := contractclient.NetworkMapContractClient{EthClient: client.EthClient{nodeUrl},NodePath:nodePath}
	//router.HandleFunc("/txn/{txn_hash}", nodeService.GetTransactionInfoHandler).Methods("GET")
	//router.HandleFunc("/txn", nodeService.GetLatestTransactionInfoHandler).Methods("GET")
	//router.HandleFunc("/block/{block_no}", nodeService.GetBlockInfoHandler).Methods("GET")
	//router.HandleFunc("/block", nodeService.GetLatestBlockInfoHandler).Methods("GET")
	//router.HandleFunc("/genesis", nodeService.GetGenesisHandler).Methods("POST", "OPTIONS")
	//router.HandleFunc("/peer/{peer_id}", nodeService.GetOtherPeerHandler).Methods("GET")
	//router.HandleFunc("/peer", nodeService.JoinNetworkHandler).Methods("POST", "OPTIONS")
	//router.HandleFunc("/peer", nodeService.GetCurrentNodeHandler).Methods("GET")
	//router.HandleFunc("/txnrcpt/{txn_hash}", nodeService.GetTransactionReceiptHandler).Methods("GET")
	//router.HandleFunc("/pendingJoinRequests", nodeService.PendingJoinRequestsHandler).Methods("GET")
	//router.HandleFunc("/joinRequestResponse", nodeService.JoinRequestResponseHandler).Methods("POST")
	//router.HandleFunc("/joinRequestResponse", nodeService.OptionsHandler).Methods("OPTIONS")
	//router.HandleFunc("/createNetwork", nodeService.CreateNetworkScriptCallHandler).Methods("POST")
	//router.HandleFunc("/createNetwork", nodeService.OptionsHandler).Methods("OPTIONS")
	//router.HandleFunc("/joinNetwork", nodeService.JoinNetworkScriptCallHandler).Methods("POST")
	//router.HandleFunc("/joinNetwork", nodeService.OptionsHandler).Methods("OPTIONS")
	//router.HandleFunc("/deployContract", nodeService.DeployContractHandler).Methods("POST")
	//router.HandleFunc("/reset", nodeService.ResetHandler).Methods("GET")
	//router.HandleFunc("/restart", nodeService.RestartHandler).Methods("GET")
	//router.HandleFunc("/latestBlock", nodeService.LatestBlockHandler).Methods("GET")
	//router.HandleFunc("/latency", nodeService.LatencyHandler).Methods("GET")
	////router.HandleFunc("/logs", nodeService.LogsHandler).Methods("GET")
	//router.HandleFunc("/txnsearch/{txn_hash}", nodeService.TransactionSearchHandler).Methods("GET")
	//router.HandleFunc("/mailserver", nodeService.MailServerConfigHandler).Methods("POST")
	//router.HandleFunc("/mailserver", nodeService.OptionsHandler).Methods("OPTIONS")
	//router.HandleFunc("/registerNode", networkMapService.RegisterNodeRequestHandler).Methods("POST")
	//router.HandleFunc("/updateNode", networkMapService.UpdateNodeHandler).Methods("POST")
	//router.HandleFunc("/updateNode", networkMapService.OptionsHandler).Methods("OPTIONS")
	//router.HandleFunc("/getNodeDetails/{index}", networkMapService.GetNodeDetailsResponseHandler).Methods("GET")
	//router.HandleFunc("/getNodeList", networkMapService.GetNodeListSelfResponseHandler).Methods("GET")
	//router.HandleFunc("/activeNodes", networkMapService.ActiveNodesHandler).Methods("GET")
	//router.HandleFunc("/chartData", nodeService.GetChartDataHandler).Methods("GET")
	//router.HandleFunc("/contractList", nodeService.GetContractListHandler).Methods("GET")
	//router.HandleFunc("/contractCount", nodeService.GetContractCountHandler).Methods("GET")
	//router.HandleFunc("/updateContractDetails", nodeService.ContractDetailsUpdateHandler).Methods("POST")
	//router.HandleFunc("/attachedNodeDetails", nodeService.AttachedNodeDetailsHandler).Methods("POST")
	//router.HandleFunc("/initialized", nodeService.InitializationHandler).Methods("GET")
	//router.HandleFunc("/createAccount", nodeService.CreateAccountHandler).Methods("POST")
	//router.HandleFunc("/createAccount", nodeService.OptionsHandler).Methods("OPTIONS")
	//router.HandleFunc("/getAccounts", nodeService.GetAccountsHandler).Methods("GET")
	//router.HandleFunc("/getWhitelist", nodeService.GetWhitelistedIPsHandler).Methods("GET")
	//router.HandleFunc("/updateWhitelist", nodeService.UpdateWhitelistHandler).Methods("POST")
	//router.HandleFunc("/updateWhitelist", nodeService.OptionsHandler).Methods("OPTIONS")
	//
	//router.PathPrefix("/contracts").Handler(http.StripPrefix("/contracts", http.FileServer(http.Dir(programPath + "contracts"))))
	//router.PathPrefix("/geth").Handler(http.StripPrefix("/geth", http.FileServer(http.Dir(nodePath + "node/qdata/gethLogs"))))
	//router.PathPrefix("/constellation").Handler(http.StripPrefix("/constellation", http.FileServer(http.Dir(nodePath + "node/qdata/constellationLogs"))))
	//router.PathPrefix("/").Handler(http.StripPrefix("/", NewFileServer("NodeManagerUI")))
	//
	//log.Info(fmt.Sprintf("Node Manager listening on %s...", listenPort))
	//
	//srv := &http.Server{
	//	Handler: router,
	//	Addr:    "0.0.0.0" + listenPort,
	//
	//	//WriteTimeout: 15 * time.Second,
	//	//ReadTimeout:  15 * time.Second,
	//	//IdleTimeout:  time.Second * 60,
	//}
	//
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil {
	//		log.Println(err)
	//	}
	//}()
	//
	//c := make(chan os.Signal, 1)
	//// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	//// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	//signal.Notify(c, os.Interrupt)
	//
	//// Block until we receive our signal.
	//<-c
	//
	//// Create a deadline to wait for.
	//ctx, cancel := context.WithTimeout(context.Background(), 15)
	//defer cancel()
	//// Doesn't block if no connections, but will otherwise wait
	//// until the timeout deadline.
	//srv.Shutdown(ctx)
	//// Optionally, you could run srv.Shutdown in a goroutine and block on
	//// <-ctx.Done() if your application should wait for other services
	//// to finalize based on context cancellation.
	//log.Info("Node Manager Shutting down")
	//os.Exit(0)
}

type MyFileServer struct {
	name    string
	handler http.Handler
}

func NewFileServer(file string) *MyFileServer {

	return &MyFileServer{file, http.FileServer(http.Dir(file))}

}
func (mf *MyFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_, err := os.Open(mf.name + "/" + r.URL.Path)
	if err != nil {
		r.URL.Path = "/"
	}

	mf.handler.ServeHTTP(w, r)
}