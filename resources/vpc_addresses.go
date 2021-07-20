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

func VpcAddresses() *schema.Table {
	return &schema.Table{
		Name:         "yandex_vpc_addresses",
		Resolver:     fetchVpcAddresses,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:        "address_id",
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
				Description: "Name of the address.\n The name is unique within the folder.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the address.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "labels",
				Type:        schema.TypeJSON,
				Description: "",
				Resolver:    client.ResolveLabels,
			},
			{
				Name:        "external_ipv_4_address_address",
				Type:        schema.TypeString,
				Description: "Value of address.",
				Resolver:    schema.PathResolver("ExternalIpv4Address.Address"),
			},
			{
				Name:        "external_ipv_4_address_zone_id",
				Type:        schema.TypeString,
				Description: "Availability zone from which the address will be allocated.",
				Resolver:    schema.PathResolver("ExternalIpv4Address.ZoneId"),
			},
			{
				Name:        "external_ipv_4_address_requirements_ddos_protection_provider",
				Type:        schema.TypeString,
				Description: "DDoS protection provider ID.",
				Resolver:    schema.PathResolver("ExternalIpv4Address.Requirements.DdosProtectionProvider"),
			},
			{
				Name:        "external_ipv_4_address_requirements_outgoing_smtp_capability",
				Type:        schema.TypeString,
				Description: "Capability to send SMTP traffic.",
				Resolver:    schema.PathResolver("ExternalIpv4Address.Requirements.OutgoingSmtpCapability"),
			},
			{
				Name:        "reserved",
				Type:        schema.TypeBool,
				Description: "Specifies if address is reserved or not.",
				Resolver:    schema.PathResolver("Reserved"),
			},
			{
				Name:        "used",
				Type:        schema.TypeBool,
				Description: "Specifies if address is used or not.",
				Resolver:    schema.PathResolver("Used"),
			},
			{
				Name:        "type",
				Type:        schema.TypeString,
				Description: "Type of the IP address.",
				Resolver:    client.EnumPathResolver("Type"),
			},
			{
				Name:        "ip_version",
				Type:        schema.TypeString,
				Description: "Vervion of the IP address.",
				Resolver:    client.EnumPathResolver("IpVersion"),
			},
		},
	}
}

func fetchVpcAddresses(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &vpc.ListAddressesRequest{FolderId: f}
		it := c.Services.Vpc.Address().AddressIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}

	return nil
}
