package args

import "flag"

var ParamHttpPort *int
var IsRunJob *int
var RunLuckyDrawJob *int

func init() {
	ParamHttpPort = flag.Int("port", 0, "http listen port")
	IsRunJob = flag.Int("run_job", 1, "run job, 1 to run , other no run")
	RunLuckyDrawJob = flag.Int("run_lucky_draw_job", 1, "run job, 1 to run , other no run")
	flag.Parse()
}
