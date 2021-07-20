// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------

package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func VpcNetworks() *schema.Table {
	return &schema.Table{
		Name:         "yandex_vpc_networks",
		Resolver:     fetchVpcNetworks,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:        "network_id",
				Type:        schema.TypeString,
				Description: "",
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "folder_id",
				Type:        schema.TypeString,
				Description: "",
				Resolver:    client.ResolveFolderID,
			},
			{
				Name:        "created_at",
				Type:        schema.TypeTimestamp,
				Description: "",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: "Name of the network.\n The name is unique within the folder. 3-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Optional description of the network. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "labels",
				Type:        schema.TypeJSON,
				Description: "",
				Resolver:    client.ResolveLabels,
			},
			{
				Name:        "default_security_group_id",
				Type:        schema.TypeString,
				Description: "ID of default security group for network.",
				Resolver:    schema.PathResolver("DefaultSecurityGroupId"),
			},
		},
	}
}

func fetchVpcNetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &vpc.ListNetworksRequest{FolderId: f}
		it := c.Services.Vpc.Network().NetworkIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}

	return nil
}
