package mariadb

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureMariadbServersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureMariadbServersGenerator{}

func (x *TableAzureMariadbServersGenerator) GetTableName() string {
	return "azure_mariadb_servers"
}

func (x *TableAzureMariadbServersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureMariadbServersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureMariadbServersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureMariadbServersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().MariaDB.Servers

			response, err := svc.List(ctx)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if response.Value == nil {
				return nil
			}
			resultChannel <- *response.Value

			return nil
		},
	}
}

func (x *TableAzureMariadbServersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureMariadbServersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("administrator_login").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("earliest_restore_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("EarliestRestoreDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_profile").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_server_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterServerID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_endpoint_connections").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_visible_state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssl_enforcement").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fully_qualified_domain_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_role").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_capacity").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_network_access").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAzureMariadbServersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAzureMariadbConfigurationsGenerator{}),
	}
}
