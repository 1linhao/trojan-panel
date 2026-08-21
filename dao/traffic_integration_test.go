package dao

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestTrafficMigrationAndQueries(t *testing.T) {
	dsn := os.Getenv("TP_TRAFFIC_TEST_DSN")
	if dsn == "" {
		t.Skip("TP_TRAFFIC_TEST_DSN is not set")
	}
	testDB, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer testDB.Close()
	db = testDB
	statements := []string{
		`DROP TABLE IF EXISTS account_server_traffic_daily,account_traffic_total,casbin_rule,account,node_server`,
		`CREATE TABLE account (id bigint unsigned primary key,username varchar(64) not null,role_id bigint unsigned not null,deleted tinyint unsigned not null,quota bigint not null,download bigint unsigned not null,upload bigint unsigned not null)`,
		`CREATE TABLE node_server (id bigint unsigned primary key,name varchar(64) not null,ip varchar(64) not null,grpc_port int unsigned not null,grpc_tls_mode varchar(16) not null,grpc_tls_server_name varchar(253) not null)`,
		`CREATE TABLE casbin_rule (p_type varchar(32),v0 varchar(255),v1 varchar(255),v2 varchar(255),v3 varchar(255),v4 varchar(255),v5 varchar(255))`,
		`INSERT INTO account VALUES (7,'alice',3,0,-1,20,10)`,
		`INSERT INTO node_server VALUES (2,'sf','127.0.0.1',8100,'mtls','sf.example.com')`,
	}
	for _, statement := range statements {
		if _, err = db.Exec(statement); err != nil {
			t.Fatal(err)
		}
	}
	if err = migrateTrafficAccountingSchema(); err != nil {
		t.Fatal(err)
	}
	if err = migrateTrafficAccountingSchema(); err != nil {
		t.Fatalf("migration is not idempotent: %v", err)
	}
	if _, err = db.Exec(`UPDATE node_server SET traffic_period='month',traffic_limit_mode='combined',traffic_total_limit=100 WHERE id=2`); err != nil {
		t.Fatal(err)
	}
	if _, err = db.Exec(`INSERT INTO account_server_traffic_daily VALUES(CURRENT_DATE(),7,2,30,40,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)`); err != nil {
		t.Fatal(err)
	}
	rank, err := TrafficRank("total")
	if err != nil || len(rank) != 1 || rank[0].TrafficUsed != 30 {
		t.Fatalf("unexpected total rank: %#v %v", rank, err)
	}
	daily, err := TrafficRank("day")
	if err != nil || len(daily) != 1 || daily[0].TrafficUsed != 70 {
		t.Fatalf("unexpected daily rank: %#v %v", daily, err)
	}
	statuses, err := SelectServerTrafficStatuses([]uint{2})
	if err != nil || len(statuses) != 1 || statuses[0].UploadUsed != 30 || statuses[0].DownloadUsed != 40 {
		t.Fatalf("unexpected status: %#v %v", statuses, err)
	}
}
