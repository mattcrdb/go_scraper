package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	//"path"
	//"path/filepath"
)

func openZip(file string) error {
	logFileRegExp := regexp.MustCompile(`([a-z]{5}/[a-z]{5}/\d+)/logs/.*.log`)
	schemaFileRegExp := regexp.MustCompile(`([a-z]{5}\/[a-z]{6}\/[[:word:]]*)`)

	// Open a zip archive for reading.
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()


	//This is tracking whether we found one log file
	//One log file is sufficient to return the version
	var foundLog bool = false

	for _, f := range r.File {
		//if f.FileInfo().IsDir(){
		//	fmt.Println(f.Name)
		//}
		//fmt.Println(f.Name)

		switch {
			//This case is to find one log file to print the version and cluster ID
			case logFileRegExp.MatchString(f.Name):
				if !foundLog {
					if err := getVersion(f); err != nil {
						return err
					}
					if err := getClusterID(f); err != nil {
						return err
					}
					foundLog = true
				}

				if err := getStartFlags(f); err != nil {
					return err
				}
			case schemaFileRegExp.MatchString(f.Name):
				if !strings.HasPrefix(f.Name,"debug/schema/system") && !strings.Contains(f.Name,"@")   {
					fmt.Printf("database: %s:\n", strings.Split(f.Name,"/")[2])
					//if err := getDDL(f); err != nil {
					//	return err
					//}
				}
			case f.Name == "debug/crdb_internal.jobs.txt":
				if err := getJobs(f); err != nil {
					return err
				}
			case f.Name == "debug/events.json":
				if err := getEvents(f); err != nil {
				return err
			}
		case f.Name == "debug/settings.json":
			if err := getClusterSettingsAndCompare(f); err != nil {
				return err
			}



		default:

		}



	}
	return nil
}

//orphaned jobs as a to-do

//Set up structs for unmarshaling JSON, might not work because the keys change between cockroach versions
type localClusterSettings struct {
	KeyValues settingValues `json:"key_values"`
}
type settingValues struct {
	ChangeFeedExperimentalPollInterval setting `json:"changefeed.experimental_poll_interval"`
	Cloudstoragegsdefaultkey setting `json:"cloudstorage.gs.default.key"`
	Cloudstoragehttpcustomca setting `json:"cloudstorage.http.custom_ca"`
	Cloudstoragetimeout setting `json:"cloudstorage.timeout"`
	Clusterorganization setting `json:"cluster.organization"`
	Clusterpreservedowngradeoption setting `json:"cluster.preserve_downgrade_option"`
	Compactorenabled setting `json:"compactor.enabled"`
	Compactormaxrecordage setting `json:"compactor.max_record_age"`
	Compactormininterval setting `json:"compactor.min_interval"`
	Compactorthresholdavailablefraction setting `json:"compactor.threshold_available_fraction"`
	Compactorthresholdbytes setting `json:"compactor.threshold_bytes"`
	Compactorthresholdusedfraction setting `json:"compactor.threshold_used_fraction"`
	Debugpaniconfailedassertions setting `json:"debug.panic_on_failed_assertions"`
	Diagnosticsforcedstatresetinterval setting `json:"diagnostics.forced_stat_reset.interval"`
	Diagnosticsreportingenabled setting `json:"diagnostics.reporting.enabled"`
	Diagnosticsreportinginterval setting `json:"diagnostics.reporting.interval"`
	Diagnosticsreportingsendcrashreports setting `json:"diagnostics.reporting.send_crash_reports"`
	Externalgraphiteendpoint setting `json:"external.graphite.endpoint"`
	Externalgraphiteinterval setting `json:"external.graphite.interval"`
	Jobsregistryleniency setting `json:"jobs.registry.leniency"`
	Jobsretentiontime setting `json:"jobs.retention_time"`
	Kvallocatorleaserebalancingaggressiveness setting `json:"kv.allocator.lease_rebalancing_aggressiveness"`
	Kvallocatorloadbasedleaserebalancingenabled setting `json:"kv.allocator.load_based_lease_rebalancing.enabled"`
	Kvallocatorloadbasedrebalancing setting `json:"kv.allocator.load_based_rebalancing"`
	Kvallocatorqpsrebalancethreshold setting `json:"kv.allocator.qps_rebalance_threshold"`
	Kvallocatorrangerebalancethreshold setting `json:"kv.allocator.range_rebalance_threshold"`
	Kvatomicreplicationchangesenabled setting `json:"kv.atomic_replication_changes.enabled"`
	Kvbulkingestbatchsize setting `json:"kv.bulk_ingest.batch_size"`
	Kvbulkingestbufferincrement setting `json:"kv.bulk_ingest.buffer_increment"`
	Kvbulkingestindexbuffersize setting `json:"kv.bulk_ingest.index_buffer_size"`
	Kvbulkingestmaxindexbuffersize setting `json:"kv.bulk_ingest.max_index_buffer_size"`
	Kvbulkingestmaxpkbuffersize setting `json:"kv.bulk_ingest.max_pk_buffer_size"`
	Kvbulkingestpkbuffersize setting `json:"kv.bulk_ingest.pk_buffer_size"`
	Kvbulkiowriteaddsstablemaxrate setting `json:"kv.bulk_io_write.addsstable_max_rate"`
	Kvbulkiowriteconcurrentaddsstablerequests setting `json:"kv.bulk_io_write.concurrent_addsstable_requests"`
	Kvbulkiowriteconcurrentexportrequests setting `json:"kv.bulk_io_write.concurrent_export_requests"`
	Kvbulkiowriteconcurrentimportrequests setting `json:"kv.bulk_io_write.concurrent_import_requests"`
	Kvbulkiowriteexperimentalincrementalexportenabled setting `json:"kv.bulk_io_write.experimental_incremental_export_enabled"`
	Kvbulkiowritemaxrate setting `json:"kv.bulk_io_write.max_rate"`
	Kvbulkiowritesmallwritesize setting `json:"kv.bulk_io_write.small_write_size"`
	Kvbulksstsyncsize setting `json:"kv.bulk_sst.sync_size"`
	Kvclosedtimestampclosefraction setting `json:"kv.closed_timestamp.close_fraction"`
	Kvclosedtimestampfollowerreadsenabled setting `json:"kv.closed_timestamp.follower_reads_enabled"`
	Kvclosedtimestamptargetduration setting `json:"kv.closed_timestamp.target_duration"`
	Kvfollowerreadtargetmultiple setting `json:"kv.follower_read.target_multiple"`
	Kvraftcommandmaxsize setting `json:"kv.raft.command.max_size"`
	Kvraftlogdisablesynchronizationunsafe setting `json:"kv.raft_log.disable_synchronization_unsafe"`
	Kvrangebackpressurerangesizemultiplier setting `json:"kv.range.backpressure_range_size_multiplier"`
	Kvrangedescriptorcachesize setting `json:"kv.range_descriptor_cache.size"`
	Kvrangemergequeueenabled setting `json:"kv.range_merge.queue_enabled"`
	Kvrangemergequeueinterval setting `json:"kv.range_merge.queue_interval"`
	Kvrangesplitbyloadenabled setting `json:"kv.range_split.by_load_enabled"`
	Kvrangesplitloadqpsthreshold setting `json:"kv.range_split.load_qps_threshold"`
	Kvrangefeedconcurrentcatchupiterators setting `json:"kv.rangefeed.concurrent_catchup_iterators"`
	Kvrangefeedenabled setting `json:"kv.rangefeed.enabled"`
	Kvreplicationreportsinterval setting `json:"kv.replication_reports.interval"`
	Kvsnapshotrebalancemaxrate setting `json:"kv.snapshot_rebalance.max_rate"`
	Kvsnapshotrecoverymaxrate setting `json:"kv.snapshot_recovery.max_rate"`
	Kvsnapshotsstsyncsize setting `json:"kv.snapshot_sst.sync_size"`
	Kvtransactionmaxintentsbytes setting `json:"kv.transaction.max_intents_bytes"`
	Kvtransactionmaxrefreshspansbytes setting `json:"kv.transaction.max_refresh_spans_bytes"`
	Kvtransactionparallelcommitsenabled setting `json:"kv.transaction.parallel_commits_enabled"`
	Kvtransactionwritepipeliningenabled setting `json:"kv.transaction.write_pipelining_enabled"`
	Kvtransactionwritepipeliningmaxbatchsize setting `json:"kv.transaction.write_pipelining_max_batch_size"`
	Kvtransactionwritepipeliningmaxoutstandingsize setting `json:"kv.transaction.write_pipelining_max_outstanding_size"`
	Rocksdbingestbackpressurel0filecountthreshold setting `json:"rocksdb.ingest_backpressure.l0_file_count_threshold"`
	Rocksdbingestbackpressuremaxdelay setting `json:"rocksdb.ingest_backpressure.max_delay"`
	Rocksdbingestbackpressurependingcompactionthreshold setting `json:"rocksdb.ingest_backpressure.pending_compaction_threshold"`
	Rocksdbminwalsyncinterval setting `json:"rocksdb.min_wal_sync_interval"`
	Schemachangerbackfillerbufferincrement setting `json:"schemachanger.backfiller.buffer_increment"`
	Schemachangerbackfillerbuffersize setting `json:"schemachanger.backfiller.buffer_size"`
	Schemachangerbackfillermaxbuffersize setting `json:"schemachanger.backfiller.max_buffer_size"`
	Schemachangerbackfillermaxsstsize setting `json:"schemachanger.backfiller.max_sst_size"`
	Schemachangerbulkindexbackfillbatchsize setting `json:"schemachanger.bulk_index_backfill.batch_size"`
	Schemachangerleaseduration setting `json:"schemachanger.lease.duration"`
	Schemachangerleaserenewfraction setting `json:"schemachanger.lease.renew_fraction"`
	Serverclockforwardjumpcheckenabled setting `json:"server.clock.forward_jump_check_enabled"`
	Serverclockpersistupperboundinterval setting `json:"server.clock.persist_upper_bound_interval"`
	Serverconsistencycheckinterval setting `json:"server.consistency_check.interval"`
	Serverdeclinedreservationtimeout setting `json:"server.declined_reservation_timeout"`
	Servereventlogttl setting `json:"server.eventlog.ttl"`
	Serverfailedreservationtimeout setting `json:"server.failed_reservation_timeout"`
	Servergoroutinedumpnumgoroutinesthreshold setting `json:"server.goroutine_dump.num_goroutines_threshold"`
	Servergoroutinedumptotaldumpsizelimit setting `json:"server.goroutine_dump.total_dump_size_limit"`
	Serverheapprofilemaxprofiles setting `json:"server.heap_profile.max_profiles"`
	Serverhostbasedauthenticationconfiguration setting `json:"server.host_based_authentication.configuration"`
	Serverrangelogttl setting `json:"server.rangelog.ttl"`
	Serverremotedebuggingmode setting `json:"server.remote_debugging.mode"`
	Servershutdowndrainwait setting `json:"server.shutdown.drain_wait"`
	Servershutdownquerywait setting `json:"server.shutdown.query_wait"`
	Servertimeuntilstoredead setting `json:"server.time_until_store_dead"`
	Serverwebsessiontimeout setting `json:"server.web_session_timeout"`
	Sqldefaultsdefaultintsize setting `json:"sql.defaults.default_int_size"`
	Sqldefaultsdistsql setting `json:"sql.defaults.distsql"`
	Sqldefaultsexperimentaloptimizerforeignkeysenabled setting `json:"sql.defaults.experimental_optimizer_foreign_keys.enabled"`
	Sqldefaultsreorderjoinslimit setting `json:"sql.defaults.reorder_joins_limit"`
	Sqldefaultsresultsbuffersize setting `json:"sql.defaults.results_buffer.size"`
	Sqldefaultsserialnormalization setting `json:"sql.defaults.serial_normalization"`
	Sqldefaultsvectorize setting `json:"sql.defaults.vectorize"`
	Sqldefaultsvectorizerowcountthreshold setting `json:"sql.defaults.vectorize_row_count_threshold"`
	Sqldefaultszigzagjoinenabled setting `json:"sql.defaults.zigzag_join.enabled"`
	Sqldistsqldistributeindexjoins setting `json:"sql.distsql.distribute_index_joins"`
	Sqldistsqlflowstreamtimeout setting `json:"sql.distsql.flow_stream_timeout"`
	Sqldistsqlinterleavedjoinsenabled setting `json:"sql.distsql.interleaved_joins.enabled"`
	Sqldistsqlmaxrunningflows setting `json:"sql.distsql.max_running_flows"`
	Sqldistsqlmergejoinsenabled setting `json:"sql.distsql.merge_joins.enabled"`
	Sqldistsqltempstoragejoins setting `json:"sql.distsql.temp_storage.joins"`
	Sqldistsqltempstoragesorts setting `json:"sql.distsql.temp_storage.sorts"`
	Sqldistsqltempstorageworkmem setting `json:"sql.distsql.temp_storage.workmem"`
	Sqlmetricsstatementdetailsdumptologs setting `json:"sql.metrics.statement_details.dump_to_logs"`
	Sqlmetricsstatementdetailsenabled setting `json:"sql.metrics.statement_details.enabled"`
	Sqlmetricsstatementdetailsplancollectionenabled setting `json:"sql.metrics.statement_details.plan_collection.enabled"`
	Sqlmetricsstatementdetailsplancollectionperiod setting `json:"sql.metrics.statement_details.plan_collection.period"`
	Sqlmetricsstatementdetailsthreshold setting `json:"sql.metrics.statement_details.threshold"`
	Sqlmetricstransactiondetailsenabled setting `json:"sql.metrics.transaction_details.enabled"`
	Sqlparallelscansenabled setting `json:"sql.parallel_scans.enabled"`
	Sqlquerycacheenabled setting `json:"sql.query_cache.enabled"`
	Sqlstatsautomaticcollectionenabled setting `json:"sql.stats.automatic_collection.enabled"`
	Sqlstatsautomaticcollectionfractionstalerows setting `json:"sql.stats.automatic_collection.fraction_stale_rows"`
	Sqlstatsautomaticcollectionmaxfractionidle setting `json:"sql.stats.automatic_collection.max_fraction_idle"`
	Sqlstatsautomaticcollectionminstalerows setting `json:"sql.stats.automatic_collection.min_stale_rows"`
	Sqlstatshistogramcollectionenabled setting `json:"sql.stats.histogram_collection.enabled"`
	Sqlstatsmaxtimestampage setting `json:"sql.stats.max_timestamp_age"`
	Sqlstatsposteventsenabled setting `json:"sql.stats.post_events.enabled"`
	Sqltablecacheleaserefreshlimit setting `json:"sql.tablecache.lease.refresh_limit"`
	Sqltracelogstatementexecute setting `json:"sql.trace.log_statement_execute"`
	Sqltracesessioneventlogenabled setting `json:"sql.trace.session_eventlog.enabled"`
	Sqltracetxnenablethreshold setting `json:"sql.trace.txn.enable_threshold"`
	Timeseriesstorageenabled setting `json:"timeseries.storage.enabled"`
	Timeseriesstorageresolution10sttl setting `json:"timeseries.storage.resolution_10s.ttl"`
	Timeseriesstorageresolution30mttl setting `json:"timeseries.storage.resolution_30m.ttl"`
	Tracedebugenable setting `json:"trace.debug.enable"`
	Tracelightsteptoken setting `json:"trace.lightstep.token"`
	Tracezipkincollector setting `json:"trace.zipkin.collector"`
	Version setting `json:"version"`
}
type setting struct {
	Value string `json:"value"`
	Type string `json:"type"`
	Description string `json:"description"`
}
func getClusterSettingsAndCompare (f *zip.File) error {
	rff, err := f.Open()

	if err != nil {
		return err
	}

	contents, err := ioutil.ReadAll(rff)

	if err != nil {
		return err
	}


	var k localClusterSettings
	err = json.Unmarshal(contents, &k)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cluster Settings from debug.zip")
	fmt.Println(k.KeyValues.ChangeFeedExperimentalPollInterval.Description)
	fmt.Println(k.KeyValues.Version.Value)
	fmt.Println(k.KeyValues.Timeseriesstorageresolution30mttl.Value)


	return nil
}


type eventJson struct {
	Events []event `json:"events"`
}
type event struct {
	Timestamp string `json:"timestamp"`
	Event_type string `json:"event_type"`
	Target_id int `json:"target_id"`
	Reporting_id int `json:"reporting_id"`
	Info string `json:"info"`
	Unique_id string `json:"unique_id"`

}
func getEvents(f *zip.File) error {
	rff, err := f.Open()

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(rff)

	if err != nil {
		return err
	}


	var e eventJson
	err = json.Unmarshal(data, &e)

	if err != nil {
		fmt.Println(err)
	}
	var e1 = "node_join"
	var e2 = "node_restart"
	var e3 = "node_decommissioned"
	var e4 = "node_recommissioned"
	fmt.Println("Node related events from events.json:")

	for _, event := range e.Events {

		if (event.Event_type == e1) || (event.Event_type == e2) || (event.Event_type == e3) || (event.Event_type == e4) {
			fmt.Printf("%s n%d at %s %s\n",event.Event_type,event.Target_id,event.Timestamp[:10], event.Timestamp[12:])
		}
	}

	//for key,value := range(m) {
	//	fmt.Println(key)
	//	fmt.Println(len(value))
	//	var n map[string]string
	//	err = json.Unmarshal([]byte("fd"),n)
	//	for key,value := range(n) {
	//		fmt.Println(key,value)
	//	}
	//}


	return nil
}

func readFromJson(f *zip.File) error {
	rff, err := f.Open()

	if err != nil {
		return err
	}

	contents, err := ioutil.ReadAll(rff)

	if err != nil {
		return err
	}


	var m map[string]interface{}
	err = json.Unmarshal(contents, &m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("from file: %s\n",f.Name)
	for k,v := range m {
		fmt.Printf("%s : %s\n",k,v)
	}

	//can also index into the map instead of for loop and checking if k == "<mykey>"

	return nil
}

func getJobs(f *zip.File) error {
	fmt.Println("NON SUCCEEDED JOBS:")
	re := regexp.MustCompile(`succeeded`)

	rff, err := f.Open()

	if err != nil {
		return err
	}
	defer rff.Close()

	scanner := bufio.NewScanner(rff)

	var gotCols bool = false
	var gotDashes bool = false

	for scanner.Scan() {
		if !gotCols {
			var cols = strings.Split(scanner.Text(),"|")
			fmt.Println(cols[0] + "|" + cols[1] + "|" + cols[2] + "|" + cols[6] + "|" + cols[14] + "\n")
			gotCols = true
		} else if !gotDashes {
			var dashes = strings.Split(scanner.Text(),"+")
			fmt.Println(dashes[0] + "+" + dashes[1] + "+" + dashes[2] + "+" + dashes[3] + "+" + dashes[7] + "+" + dashes[15] + "\n")
			gotDashes = true
		} else {
			if !re.MatchString(scanner.Text()) {
				data := strings.Split(scanner.Text(), "|")
				//Received a "panic: runtime error: index out of range" when reaching the last line of this file
				//Workaround is checking the length of data, if it's 1 then EOF
				if len(data) != 1 {
					fmt.Println(data[0] + "|" + data[1] + "|" + data[2] + "|" + data[6] + "|" + data[14] + "\n")
				}
			}



		}


	}


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getDDL(f *zip.File) error {
	//fmt.Println(f.Name)
	rff, err := f.Open()

	if err != nil {
		return err
	}

	contents, err := ioutil.ReadAll(rff)

	if err != nil {
		return err
	}


	var m map[string]interface{}
	err = json.Unmarshal(contents, &m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", m["create_table_statement"])


	return nil

}


func getClusterID(f *zip.File) error {
	re := regexp.MustCompile(`clusterID: [0-9a-f]{8}\-[0-9a-f]{4}\-[0-9a-f]{4}\-[0-9a-f]{4}\-[0-9a-f]{12}`)
	rff, err := f.Open()

	if err != nil {
		return err
	}
	defer rff.Close()

	scanner := bufio.NewScanner(rff)
	for scanner.Scan() {
		if len(re.FindString(scanner.Text())) != 0 {
			fmt.Printf("%s\n",re.FindString(scanner.Text()))
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getStartFlags(f *zip.File) error {
	re := regexp.MustCompile(`arguments: \[(.*?)]`)
	rff, err := f.Open()

	if err != nil {
		return err
	}
	defer rff.Close()

	scanner := bufio.NewScanner(rff)
	for scanner.Scan() {
		if len(re.FindString(scanner.Text())) != 0 {
			fmt.Printf("node %s %s\n",strings.Split(f.Name,"/")[2],re.FindString(scanner.Text()))
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getVersion(f *zip.File ) error {
	re := regexp.MustCompile(`v[0-9]+.[0-9].[0-9]`)

	rff, err := f.Open()

	if err != nil {
		return err
	}
	defer rff.Close()

	scanner := bufio.NewScanner(rff)
	for scanner.Scan() {
		if len(re.FindString(scanner.Text())) != 0 {
			fmt.Printf("CockroachDB %s\n",re.FindString(scanner.Text()))
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

return nil
}

//func getText(base string) {
//	nodesDir := path.Join(base, "nodes")
//	//fmt.Println(nodesDir)
//	files, err := ioutil.ReadDir(nodesDir)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, nodeNumber := range files {
//		if nodeNumber.Name() != ".DS_Store" {
//			nodeFolder := path.Join(nodesDir, nodeNumber.Name())
//			logFiles, err := ioutil.ReadDir(nodeFolder)
//			if err != nil {
//				log.Fatal(err)
//			}
//			for _, logfile := range logFiles {
//				fmt.Println(logfile.Name())
//				fullPathToLogFile := path.Join(nodeFolder, logfile.Name())
//				data, err := ioutil.ReadFile(fullPathToLogFile)
//				if err != nil {
//					log.Fatal(err)
//				}
//				fmt.Println(string(data))
//			}
//			//fmt.Println(nodeFolder)
//			//fmt.Println(f.Name())
//		}
//
//	}
//
//}

func main() {
	var argsWithoutProg = os.Args[1:]

	fmt.Println(argsWithoutProg[0])
	openZip(argsWithoutProg[0])


}
