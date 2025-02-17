// Code generated by codegen; DO NOT EDIT.

package storage

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:      "yandex_storage_buckets",
		Resolver:  fetchBuckets,
		Multiplex: client.MultiplexBy(client.Folders),
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
				Description: `Resource ID`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "folder_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FolderId"),
			},
			{
				Name:     "anonymous_access_flags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AnonymousAccessFlags"),
			},
			{
				Name:     "default_storage_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultStorageClass"),
			},
			{
				Name:     "versioning",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Versioning"),
			},
			{
				Name:     "max_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxSize"),
			},
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policy"),
			},
			{
				Name:     "acl",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Acl"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreatedAt"),
			},
			{
				Name:     "cors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Cors"),
			},
			{
				Name:     "website_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WebsiteSettings"),
			},
			{
				Name:     "lifecycle_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LifecycleRules"),
			},
			{
				Name:     "server_side_encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServerSideEncryption"),
			},
		},
	}
}
