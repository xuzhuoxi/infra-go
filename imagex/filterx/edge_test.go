package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"testing"
)

func TestCreateEdgeFilter(t *testing.T) {
	var temp *FilterTemplate
	temp, _ = CreateEdgeFilter(1, imagex.AllDirection, 0)
	fmt.Println(temp)
	temp, _ = CreateEdgeFilter(2, imagex.Vertical, 1)
	fmt.Println(temp)
	temp, _ = CreateEdgeFilter(3, imagex.Oblique, 2)
	fmt.Println(temp)
}
