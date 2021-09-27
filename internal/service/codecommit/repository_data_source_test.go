package codecommit_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/codecommit"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func TestAccAWSCodeCommitRepositoryDataSource_basic(t *testing.T) {
	rName := fmt.Sprintf("tf-acctest-%d", sdkacctest.RandInt())
	resourceName := "aws_codecommit_repository.default"
	datasourceName := "data.aws_codecommit_repository.default"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(t) },
		ErrorCheck: acctest.ErrorCheck(t, codecommit.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsCodeCommitRepositoryDataSourceConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "clone_url_http", resourceName, "clone_url_http"),
					resource.TestCheckResourceAttrPair(datasourceName, "clone_url_ssh", resourceName, "clone_url_ssh"),
					resource.TestCheckResourceAttrPair(datasourceName, "repository_name", resourceName, "repository_name"),
				),
			},
		},
	})
}

func testAccCheckAwsCodeCommitRepositoryDataSourceConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_codecommit_repository" "default" {
  repository_name = "%s"
}

data "aws_codecommit_repository" "default" {
  repository_name = aws_codecommit_repository.default.repository_name
}
`, rName)
}