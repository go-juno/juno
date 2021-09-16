package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/util"
	"github.com/go-juno/juno/static"
	"golang.org/x/xerrors"
)

type GrpcRelatedService interface {
	CreateProto(mod, name string) (err error)
	CreateService(mod, name string) (err error)
	WireGrpc(mod, name string) (err error)
}

type grpcRelatedService struct {
}

func (s *grpcRelatedService) CreateProto(mod, name string) (err error) {
	camel, class, snake, hyphen := util.TransformName(name)

	// proto
	grpcProtoDirPath := filepath.Join(constant.GrpcDirPath, "protos")
	protoFileName := filepath.Join(grpcProtoDirPath, fmt.Sprintf("%s.proto", snake))
	var ok bool
	ok, err = util.IsExistsFile(protoFileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		err = errors.New("File already exists")
		return
	}

	err = util.Mkdir(grpcProtoDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 替换 模块
	schemaContent := util.Replace(static.GrpcProto, mod, camel, class, snake, hyphen)
	err = util.WriteToFile(protoFileName, schemaContent)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 生成proto文件
	err = util.GenProto(snake)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	return
}

func (s *grpcRelatedService) CreateService(mod, name string) (err error) {
	camel, class, snake, hyphen := util.TransformName(name)

	// service
	grpcServiceDirPath := filepath.Join(constant.GrpcDirPath, "service")
	serviceFileName := filepath.Join(grpcServiceDirPath, fmt.Sprintf("%s.go", snake))
	err = util.Mkdir(grpcServiceDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 替换 模块
	content := util.Replace(static.GrpcService, mod, camel, class, snake, hyphen)
	err = util.WriteToFile(serviceFileName, content)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return

}
func (s *grpcRelatedService) WireService(mod, name string) (err error) {
	_, class, snake, _ := util.TransformName(name)
	serviceFilePath := filepath.Join(constant.GrpcDirPath, "service")

	serviceFileName := filepath.Join(serviceFilePath, fmt.Sprintf("%s.go", snake))
	var ok bool
	ok, err = util.IsExistsFile(serviceFileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	var content string = `
	package service

	import "github.com/google/wire"
	
	var ProviderSet = wire.NewSet(
	)
	
	`
	if ok {
		var serviceFile *os.File
		serviceFile, err = os.Open(serviceFileName)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		var c []byte
		c, err = ioutil.ReadAll(serviceFile)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		content = string(c)
		exitsService := fmt.Sprintf(`%sServer`, class)
		if strings.Contains(content, exitsService) {
			return
		}

	}
	newString := fmt.Sprintf(`New%sServer,
	)`, class)
	file := strings.ReplaceAll(content, ")", newString)
	err = util.WriteToFile(serviceFilePath, file)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	err = util.FmtCode()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *grpcRelatedService) WireGrpc(mod, name string) (err error) {
	// wire add service
	camel, class, _, _ := util.TransformName(name)
	grpcFileName := filepath.Join(constant.GrpcDirPath, "grpc.go")
	var content string = fmt.Sprintf(`
	package grpc

	import (
		"%s/api/grpc/protos"
		"%s/api/grpc/service"

		"google.golang.org/grpc"
	)

	func NewGrpc(
	) (s *grpc.Server, err error) {
		s = grpc.NewServer()
		return
	}
	`, mod, mod)
	var ok bool
	ok, err = util.IsExistsFile(grpcFileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		var file *os.File
		file, err = os.Open(grpcFileName)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		var c []byte
		c, err = ioutil.ReadAll(file)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		content = string(c)
		exitsService := fmt.Sprintf(`Register%sServer`, class)
		if strings.Contains(content, exitsService) {
			return
		}

	}
	newString := fmt.Sprintf(`func NewGrpc( 
%s *service.%sServer,`, camel, class)
	registerString := fmt.Sprintf(`protos.Register%sServer(s, %s)
	return`, class, camel)
	file := strings.ReplaceAll(content, `func NewGrpc(`, newString)
	file = strings.ReplaceAll(file, `return`, registerString)
	err = util.WriteToFile(grpcFileName, file)
	if err != nil {
		log.Println("err", err)
		return
	}
	err = util.FmtCode()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func NewGrpcRelatedService() GrpcRelatedService {
	return &grpcRelatedService{}
}
