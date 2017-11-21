package main

import (
	"github.com/hashicorp/go-multierror"
	"fmt"
)

func main() {

	var result *multierror.Error
	//er1 := grpc.Errorf(codes.PermissionDenied, "Workers cannot delete jobs")
	//result = multierror.Append(result, er1)
	//er2 := grpc.Errorf(codes.PermissionDenied, "2 cannot delete jobs")
	//result = multierror.Append(result, er2)
	result = multierror.Append(result, nil)
	result = multierror.Append(result, nil)

	//result = multierror.Append(result, er2)
	if (result.ErrorOrNil() == nil) {
		fmt.Print("k")
	}
	fmt.Println(result)
	//fmt.Print(er1)
}
