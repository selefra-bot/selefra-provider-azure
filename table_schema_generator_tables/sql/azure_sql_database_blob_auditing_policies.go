package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSqlDatabaseBlobAuditingPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSqlDatabaseBlobAuditingPoliciesGenerator{}

func (x *TableAzureSqlDatabaseBlobAuditingPoliciesGenerator) GetTableName() string {
	return "azure_sql_database_blob_auditing_policies"
}

func (x *TableAzureSqlDatabaseBlobAuditingPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSqlDatabaseBlobAuditingPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSqlDatabaseBlobAuditingPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureSqlDatabaseBlobAuditingPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().SQL.DatabaseBlobAuditingPolicies

			server := task.ParentTask.ParentRawResult.(sql.Server)
			database := task.ParentRawResult.(sql.Database)
			resourceDetails, err := azure_client.ParseResourceID(*database.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.ListByDatabase(ctx, resourceDetails.ResourceGroup, *server.Name, *database.Name)

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

func (x *TableAzureSqlDatabaseBlobAuditingPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureSqlDatabaseBlobAuditingPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("retention_days").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("queue_delay_ms").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sql_database_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_endpoint").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("audit_actions_and_groups").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_storage_secondary_key_in_use").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_azure_monitor_target_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_sql_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_sql_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_account_access_key").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_account_subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageAccountSubscriptionID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAzureSqlDatabaseBlobAuditingPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
