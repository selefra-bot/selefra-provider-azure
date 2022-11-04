package eventhub

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureEventhubNamespacesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureEventhubNamespacesGenerator{}

func (x *TableAzureEventhubNamespacesGenerator) GetTableName() string {
	return "azure_eventhub_namespaces"
}

func (x *TableAzureEventhubNamespacesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureEventhubNamespacesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureEventhubNamespacesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureEventhubNamespacesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().EventHub.Namespaces

			response, err := svc.List(ctx)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			for response.NotDone() {
				resultChannel <- response.Values()
				if err := response.NextWithContext(ctx); err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
			}

			return nil
		},
	}
}

func (x *TableAzureEventhubNamespacesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureEventhubNamespacesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_bus_endpoint").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_redundant").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_auto_inflate_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum_throughput_units").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioning_state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MetricID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kafka_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arm_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterArmID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAzureEventhubNamespacesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAzureEventhubNetworkRuleSetsGenerator{}),
	}
}
