// info_test
package test

import (
	"appinhouse/server/models"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	info, _ := models.InitConfig("./conf/test.conf")
	fmt.Println(info)
}
