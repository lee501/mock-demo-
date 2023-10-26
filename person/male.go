package person

//生成多个mock文件 go generate ./...
//go:generate mockgen -destination=../mock/dmale_mock.go -package=mockd mock-demo Male

type Male interface {
	Get(id int64) error
}
