package main

const DefaultClusterSettings19_2_3 string = `
{
  "key_values": {
    "changefeed.experimental_poll_interval": {
      "value": "1s",
      "type": "d",
      "description": "polling interval for the prototype changefeed implementation (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "cloudstorage.gs.default.key": {
      "type": "s",
      "description": "if set, JSON key to use during Google Cloud Storage operations"
    },
    "cloudstorage.http.custom_ca": {
      "type": "s",
      "description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
    },
    "cloudstorage.timeout": {
      "value": "10m0s",
      "type": "d",
      "description": "the timeout for import/export storage operations"
    },
    "cluster.organization": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "organization name"
    },
    "cluster.preserve_downgrade_option": {
      "type": "s",
      "description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
    },
    "compactor.enabled": {
      "value": "true",
      "type": "b",
      "description": "when false, the system will reclaim space occupied by deleted data less aggressively"
    },
    "compactor.max_record_age": {
      "value": "24h0m0s",
      "type": "d",
      "description": "discard suggestions not processed within this duration (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.min_interval": {
      "value": "15s",
      "type": "d",
      "description": "minimum time interval to wait before compacting (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_available_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the available logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_bytes": {
      "value": "256 MiB",
      "type": "z",
      "description": "minimum expected logical space reclamation required before considering an aggregated suggestion (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_used_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the used logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "debug.panic_on_failed_assertions": {
      "value": "false",
      "type": "b",
      "description": "panic when an assertion fails rather than reporting"
    },
    "diagnostics.forced_stat_reset.interval": {
      "value": "2h0m0s",
      "type": "d",
      "description": "interval after which pending diagnostics statistics should be discarded even if not reported"
    },
    "diagnostics.reporting.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable reporting diagnostic metrics to cockroach labs"
    },
    "diagnostics.reporting.interval": {
      "value": "1h0m0s",
      "type": "d",
      "description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
    },
    "diagnostics.reporting.send_crash_reports": {
      "value": "true",
      "type": "b",
      "description": "send crash and panic reports"
    },
    "external.graphite.endpoint": {
      "type": "s",
      "description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
    },
    "external.graphite.interval": {
      "value": "10s",
      "type": "d",
      "description": "the interval at which metrics are pushed to Graphite (if enabled)"
    },
    "jobs.registry.leniency": {
      "value": "1m0s",
      "type": "d",
      "description": "the amount of time to defer any attempts to reschedule a job"
    },
    "jobs.retention_time": {
      "value": "336h0m0s",
      "type": "d",
      "description": "the amount of time to retain records for completed jobs before"
    },
    "kv.allocator.lease_rebalancing_aggressiveness": {
      "value": "1",
      "type": "f",
      "description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
    },
    "kv.allocator.load_based_lease_rebalancing.enabled": {
      "value": "true",
      "type": "b",
      "description": "set to enable rebalancing of range leases based on load and latency"
    },
    "kv.allocator.load_based_rebalancing": {
      "value": "leases and replicas",
      "type": "e",
      "description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
    },
    "kv.allocator.qps_rebalance_threshold": {
      "value": "0.25",
      "type": "f",
      "description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
    },
    "kv.allocator.range_rebalance_threshold": {
      "value": "0.05",
      "type": "f",
      "description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
    },
    "kv.atomic_replication_changes.enabled": {
      "value": "true",
      "type": "b",
      "description": "use atomic replication changes"
    },
    "kv.bulk_ingest.batch_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "the maximum size of the payload in an AddSSTable request"
    },
    "kv.bulk_ingest.buffer_increment": {
      "value": "32 MiB",
      "type": "z",
      "description": "the size by which the BulkAdder attempts to grow its buffer before flushing"
    },
    "kv.bulk_ingest.index_buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling secondary index imports"
    },
    "kv.bulk_ingest.max_index_buffer_size": {
      "value": "512 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling secondary index imports"
    },
    "kv.bulk_ingest.max_pk_buffer_size": {
      "value": "128 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling primary index imports"
    },
    "kv.bulk_ingest.pk_buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling primary index imports"
    },
    "kv.bulk_io_write.addsstable_max_rate": {
      "value": "1.7976931348623157E+308",
      "type": "f",
      "description": "maximum number of AddSSTable requests per second for a single store"
    },
    "kv.bulk_io_write.concurrent_addsstable_requests": {
      "value": "1",
      "type": "i",
      "description": "number of AddSSTable requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_export_requests": {
      "value": "3",
      "type": "i",
      "description": "number of export requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_import_requests": {
      "value": "1",
      "type": "i",
      "description": "number of import requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.experimental_incremental_export_enabled": {
      "value": "false",
      "type": "b",
      "description": "use experimental time-bound file filter when exporting in BACKUP"
    },
    "kv.bulk_io_write.max_rate": {
      "value": "1.0 TiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
    },
    "kv.bulk_io_write.small_write_size": {
      "value": "400 KiB",
      "type": "z",
      "description": "size below which a 'bulk' write will be performed as a normal write instead"
    },
    "kv.bulk_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
    },
    "kv.closed_timestamp.close_fraction": {
      "value": "0.2",
      "type": "f",
      "description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
    },
    "kv.closed_timestamp.follower_reads_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
    },
    "kv.closed_timestamp.target_duration": {
      "value": "30s",
      "type": "d",
      "description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
    },
    "kv.follower_read.target_multiple": {
      "value": "3",
      "type": "f",
      "description": "if above 1, encourages the distsender to perform a read against the closest replica if a request is older than kv.closed_timestamp.target_duration * (1 + kv.closed_timestamp.close_fraction * this) less a clock uncertainty interval. This value also is used to create follower_timestamp(). (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.raft.command.max_size": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum size of a raft command"
    },
    "kv.raft_log.disable_synchronization_unsafe": {
      "value": "false",
      "type": "b",
      "description": "set to true to disable synchronization on Raft log writes to persistent storage. Setting to true risks data loss or data corruption on server crashes. The setting is meant for internal testing only and SHOULD NOT be used in production."
    },
    "kv.range.backpressure_range_size_multiplier": {
      "value": "2",
      "type": "f",
      "description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
    },
    "kv.range_descriptor_cache.size": {
      "value": "1000000",
      "type": "i",
      "description": "maximum number of entries in the range descriptor and leaseholder caches"
    },
    "kv.range_merge.queue_enabled": {
      "value": "true",
      "type": "b",
      "description": "whether the automatic merge queue is enabled"
    },
    "kv.range_merge.queue_interval": {
      "value": "1s",
      "type": "d",
      "description": "how long the merge queue waits between processing replicas (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.range_split.by_load_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow automatic splits of ranges based on where load is concentrated"
    },
    "kv.range_split.load_qps_threshold": {
      "value": "2500",
      "type": "i",
      "description": "the QPS over which, the range becomes a candidate for load based splitting"
    },
    "kv.rangefeed.concurrent_catchup_iterators": {
      "value": "64",
      "type": "i",
      "description": "number of rangefeeds catchup iterators a store will allow concurrently before queueing"
    },
    "kv.rangefeed.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, rangefeed registration is enabled"
    },
    "kv.replication_reports.interval": {
      "value": "1m0s",
      "type": "d",
      "description": "the frequency for generating the replication_constraint_stats, replication_stats_report and replication_critical_localities reports (set to 0 to disable)"
    },
    "kv.snapshot_rebalance.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for rebalance and upreplication snapshots"
    },
    "kv.snapshot_recovery.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for recovery snapshots"
    },
    "kv.snapshot_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which snapshot SST writes must fsync"
    },
    "kv.transaction.max_intents_bytes": {
      "value": "262144",
      "type": "i",
      "description": "maximum number of bytes used to track write intents in transactions"
    },
    "kv.transaction.max_refresh_spans_bytes": {
      "value": "256000",
      "type": "i",
      "description": "maximum number of bytes used to track refresh spans in serializable transactions"
    },
    "kv.transaction.parallel_commits_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional commits will be parallelized with transactional writes"
    },
    "kv.transaction.write_pipelining_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional writes are pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_batch_size": {
      "value": "128",
      "type": "i",
      "description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_outstanding_size": {
      "value": "256 KiB",
      "type": "z",
      "description": "maximum number of bytes used to track in-flight pipelined writes before disabling pipelining"
    },
    "rocksdb.ingest_backpressure.l0_file_count_threshold": {
      "value": "20",
      "type": "i",
      "description": "number of L0 files after which to backpressure SST ingestions"
    },
    "rocksdb.ingest_backpressure.max_delay": {
      "value": "5s",
      "type": "d",
      "description": "maximum amount of time to backpressure a single SST ingestion"
    },
    "rocksdb.ingest_backpressure.pending_compaction_threshold": {
      "value": "2.0 GiB",
      "type": "z",
      "description": "pending compaction estimate above which to backpressure SST ingestions"
    },
    "rocksdb.min_wal_sync_interval": {
      "value": "0s",
      "type": "d",
      "description": "minimum duration between syncs of the RocksDB WAL"
    },
    "schemachanger.backfiller.buffer_increment": {
      "value": "32 MiB",
      "type": "z",
      "description": "the size by which the BulkAdder attempts to grow its buffer before flushing"
    },
    "schemachanger.backfiller.buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling index backfills"
    },
    "schemachanger.backfiller.max_buffer_size": {
      "value": "512 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling index backfills"
    },
    "schemachanger.backfiller.max_sst_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "target size for ingested files during backfills"
    },
    "schemachanger.bulk_index_backfill.batch_size": {
      "value": "50000",
      "type": "i",
      "description": "number of rows to process at a time during bulk index backfill"
    },
    "schemachanger.lease.duration": {
      "value": "5m0s",
      "type": "d",
      "description": "the duration of a schema change lease"
    },
    "schemachanger.lease.renew_fraction": {
      "value": "0.5",
      "type": "f",
      "description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
    },
    "server.clock.forward_jump_check_enabled": {
      "value": "false",
      "type": "b",
      "description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic"
    },
    "server.clock.persist_upper_bound_interval": {
      "value": "0s",
      "type": "d",
      "description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
    },
    "server.consistency_check.interval": {
      "value": "24h0m0s",
      "type": "d",
      "description": "the time between range consistency checks; set to 0 to disable consistency checking"
    },
    "server.declined_reservation_timeout": {
      "value": "1s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
    },
    "server.eventlog.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "if nonzero, event log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.failed_reservation_timeout": {
      "value": "5s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
    },
    "server.goroutine_dump.num_goroutines_threshold": {
      "value": "1000",
      "type": "i",
      "description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
    },
    "server.goroutine_dump.total_dump_size_limit": {
      "value": "500 MiB",
      "type": "z",
      "description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
    },
    "server.heap_profile.max_profiles": {
      "value": "5",
      "type": "i",
      "description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept."
    },
    "server.host_based_authentication.configuration": {
      "type": "s",
      "description": "host-based authentication configuration to use during connection authentication"
    },
    "server.rangelog.ttl": {
      "value": "720h0m0s",
      "type": "d",
      "description": "if nonzero, range log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.remote_debugging.mode": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
    },
    "server.shutdown.drain_wait": {
      "value": "0s",
      "type": "d",
      "description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
    },
    "server.shutdown.query_wait": {
      "value": "10s",
      "type": "d",
      "description": "the server will wait for at least this amount of time for active queries to finish"
    },
    "server.time_until_store_dead": {
      "value": "5m0s",
      "type": "d",
      "description": "the time after which if there is no new gossiped information about a store, it is considered dead"
    },
    "server.web_session_timeout": {
      "value": "168h0m0s",
      "type": "d",
      "description": "the duration that a newly created web session will be valid"
    },
    "sql.defaults.default_int_size": {
      "value": "8",
      "type": "i",
      "description": "the size, in bytes, of an INT type"
    },
    "sql.defaults.distsql": {
      "value": "auto",
      "type": "e",
      "description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
    },
    "sql.defaults.experimental_optimizer_foreign_keys.enabled": {
      "value": "false",
      "type": "b",
      "description": "enables optimizer-driven foreign key checks by default"
    },
    "sql.defaults.reorder_joins_limit": {
      "value": "4",
      "type": "i",
      "description": "default number of joins to reorder"
    },
    "sql.defaults.results_buffer.size": {
      "value": "16 KiB",
      "type": "z",
      "description": "default size of the buffer that accumulates results for a statement or a batch of statements before they are sent to the client. This can be overridden on an individual connection with the 'results_buffer_size' parameter. Note that auto-retries generally only happen while no results have been delivered to the client, so reducing this size can increase the number of retriable errors a client receives. On the other hand, increasing the buffer size can increase the delay until the client receives the first result row. Updating the setting only affects new connections. Setting to 0 disables any buffering."
    },
    "sql.defaults.serial_normalization": {
      "value": "rowid",
      "type": "e",
      "description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
    },
    "sql.defaults.vectorize": {
      "value": "auto",
      "type": "e",
      "description": "default vectorize mode [off = 0, auto = 1, experimental_on = 2]"
    },
    "sql.defaults.vectorize_row_count_threshold": {
      "value": "1000",
      "type": "i",
      "description": "default vectorize row count threshold"
    },
    "sql.defaults.zigzag_join.enabled": {
      "value": "true",
      "type": "b",
      "description": "default value for enable_zigzag_join session setting; allows use of zig-zag join by default"
    },
    "sql.distsql.distribute_index_joins": {
      "value": "true",
      "type": "b",
      "description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
    },
    "sql.distsql.flow_stream_timeout": {
      "value": "10s",
      "type": "d",
      "description": "amount of time incoming streams wait for a flow to be set up before erroring out"
    },
    "sql.distsql.interleaved_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set we plan interleaved table joins instead of merge joins when possible"
    },
    "sql.distsql.max_running_flows": {
      "value": "500",
      "type": "i",
      "description": "maximum number of concurrent flows that can be run on a node"
    },
    "sql.distsql.merge_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, we plan merge joins when possible"
    },
    "sql.distsql.temp_storage.joins": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql joins"
    },
    "sql.distsql.temp_storage.sorts": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql sorts"
    },
    "sql.distsql.temp_storage.workmem": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
    },
    "sql.metrics.statement_details.dump_to_logs": {
      "value": "false",
      "type": "b",
      "description": "dump collected statement statistics to node logs when periodically cleared"
    },
    "sql.metrics.statement_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-statement query statistics"
    },
    "sql.metrics.statement_details.plan_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "periodically save a logical plan for each fingerprint"
    },
    "sql.metrics.statement_details.plan_collection.period": {
      "value": "5m0s",
      "type": "d",
      "description": "the time until a new logical plan is collected"
    },
    "sql.metrics.statement_details.threshold": {
      "value": "0s",
      "type": "d",
      "description": "minimum execution time to cause statistics to be collected"
    },
    "sql.metrics.transaction_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-application transaction statistics"
    },
    "sql.parallel_scans.enabled": {
      "value": "true",
      "type": "b",
      "description": "parallelizes scanning different ranges when the maximum result size can be deduced"
    },
    "sql.query_cache.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable the query cache"
    },
    "sql.stats.automatic_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "automatic statistics collection mode"
    },
    "sql.stats.automatic_collection.fraction_stale_rows": {
      "value": "0.2",
      "type": "f",
      "description": "target fraction of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.automatic_collection.max_fraction_idle": {
      "value": "0.9",
      "type": "f",
      "description": "maximum fraction of time that automatic statistics sampler processors are idle"
    },
    "sql.stats.automatic_collection.min_stale_rows": {
      "value": "500",
      "type": "i",
      "description": "target minimum number of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.histogram_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "histogram collection mode"
    },
    "sql.stats.max_timestamp_age": {
      "value": "5m0s",
      "type": "d",
      "description": "maximum age of timestamp during table statistics collection"
    },
    "sql.stats.post_events.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, an event is shown for every CREATE STATISTICS job"
    },
    "sql.tablecache.lease.refresh_limit": {
      "value": "50",
      "type": "i",
      "description": "maximum number of tables to periodically refresh leases for"
    },
    "sql.trace.log_statement_execute": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable logging of executed statements"
    },
    "sql.trace.session_eventlog.enabled": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable session tracing"
    },
    "sql.trace.txn.enable_threshold": {
      "value": "0s",
      "type": "d",
      "description": "duration beyond which all transactions are traced (set to 0 to disable)"
    },
    "timeseries.storage.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
    },
    "timeseries.storage.resolution_10s.ttl": {
      "value": "240h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
    },
    "trace.debug.enable": {
      "value": "false",
      "type": "b",
      "description": "if set, traces for recent requests can be seen in the /debug page"
    },
    "trace.lightstep.token": {
      "type": "s",
      "description": "if set, traces go to Lightstep using this token"
    },
    "trace.zipkin.collector": {
      "type": "s",
      "description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set"
    },
    "version": {
      "value": "19.2",
      "type": "m",
      "description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'"
    }
  }
}
`
const DefaultClusterSettings19_2_2 string = `
{
  "key_values": {
    "changefeed.experimental_poll_interval": {
      "value": "1s",
      "type": "d",
      "description": "polling interval for the prototype changefeed implementation (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "cloudstorage.gs.default.key": {
      "type": "s",
      "description": "if set, JSON key to use during Google Cloud Storage operations"
    },
    "cloudstorage.http.custom_ca": {
      "type": "s",
      "description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
    },
    "cloudstorage.timeout": {
      "value": "10m0s",
      "type": "d",
      "description": "the timeout for import/export storage operations"
    },
    "cluster.organization": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "organization name"
    },
    "cluster.preserve_downgrade_option": {
      "type": "s",
      "description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
    },
    "compactor.enabled": {
      "value": "true",
      "type": "b",
      "description": "when false, the system will reclaim space occupied by deleted data less aggressively"
    },
    "compactor.max_record_age": {
      "value": "24h0m0s",
      "type": "d",
      "description": "discard suggestions not processed within this duration (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.min_interval": {
      "value": "15s",
      "type": "d",
      "description": "minimum time interval to wait before compacting (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_available_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the available logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_bytes": {
      "value": "256 MiB",
      "type": "z",
      "description": "minimum expected logical space reclamation required before considering an aggregated suggestion (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_used_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the used logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "debug.panic_on_failed_assertions": {
      "value": "false",
      "type": "b",
      "description": "panic when an assertion fails rather than reporting"
    },
    "diagnostics.forced_stat_reset.interval": {
      "value": "2h0m0s",
      "type": "d",
      "description": "interval after which pending diagnostics statistics should be discarded even if not reported"
    },
    "diagnostics.reporting.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable reporting diagnostic metrics to cockroach labs"
    },
    "diagnostics.reporting.interval": {
      "value": "1h0m0s",
      "type": "d",
      "description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
    },
    "diagnostics.reporting.send_crash_reports": {
      "value": "true",
      "type": "b",
      "description": "send crash and panic reports"
    },
    "external.graphite.endpoint": {
      "type": "s",
      "description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
    },
    "external.graphite.interval": {
      "value": "10s",
      "type": "d",
      "description": "the interval at which metrics are pushed to Graphite (if enabled)"
    },
    "jobs.registry.leniency": {
      "value": "1m0s",
      "type": "d",
      "description": "the amount of time to defer any attempts to reschedule a job"
    },
    "jobs.retention_time": {
      "value": "336h0m0s",
      "type": "d",
      "description": "the amount of time to retain records for completed jobs before"
    },
    "kv.allocator.lease_rebalancing_aggressiveness": {
      "value": "1",
      "type": "f",
      "description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
    },
    "kv.allocator.load_based_lease_rebalancing.enabled": {
      "value": "true",
      "type": "b",
      "description": "set to enable rebalancing of range leases based on load and latency"
    },
    "kv.allocator.load_based_rebalancing": {
      "value": "leases and replicas",
      "type": "e",
      "description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
    },
    "kv.allocator.qps_rebalance_threshold": {
      "value": "0.25",
      "type": "f",
      "description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
    },
    "kv.allocator.range_rebalance_threshold": {
      "value": "0.05",
      "type": "f",
      "description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
    },
    "kv.atomic_replication_changes.enabled": {
      "value": "true",
      "type": "b",
      "description": "use atomic replication changes"
    },
    "kv.bulk_ingest.batch_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "the maximum size of the payload in an AddSSTable request"
    },
    "kv.bulk_ingest.buffer_increment": {
      "value": "32 MiB",
      "type": "z",
      "description": "the size by which the BulkAdder attempts to grow its buffer before flushing"
    },
    "kv.bulk_ingest.index_buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling secondary index imports"
    },
    "kv.bulk_ingest.max_index_buffer_size": {
      "value": "512 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling secondary index imports"
    },
    "kv.bulk_ingest.max_pk_buffer_size": {
      "value": "128 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling primary index imports"
    },
    "kv.bulk_ingest.pk_buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling primary index imports"
    },
    "kv.bulk_io_write.addsstable_max_rate": {
      "value": "1.7976931348623157E+308",
      "type": "f",
      "description": "maximum number of AddSSTable requests per second for a single store"
    },
    "kv.bulk_io_write.concurrent_addsstable_requests": {
      "value": "1",
      "type": "i",
      "description": "number of AddSSTable requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_export_requests": {
      "value": "3",
      "type": "i",
      "description": "number of export requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_import_requests": {
      "value": "1",
      "type": "i",
      "description": "number of import requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.experimental_incremental_export_enabled": {
      "value": "false",
      "type": "b",
      "description": "use experimental time-bound file filter when exporting in BACKUP"
    },
    "kv.bulk_io_write.max_rate": {
      "value": "1.0 TiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
    },
    "kv.bulk_io_write.small_write_size": {
      "value": "400 KiB",
      "type": "z",
      "description": "size below which a 'bulk' write will be performed as a normal write instead"
    },
    "kv.bulk_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
    },
    "kv.closed_timestamp.close_fraction": {
      "value": "0.2",
      "type": "f",
      "description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
    },
    "kv.closed_timestamp.follower_reads_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
    },
    "kv.closed_timestamp.target_duration": {
      "value": "30s",
      "type": "d",
      "description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
    },
    "kv.follower_read.target_multiple": {
      "value": "3",
      "type": "f",
      "description": "if above 1, encourages the distsender to perform a read against the closest replica if a request is older than kv.closed_timestamp.target_duration * (1 + kv.closed_timestamp.close_fraction * this) less a clock uncertainty interval. This value also is used to create follower_timestamp(). (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.raft.command.max_size": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum size of a raft command"
    },
    "kv.raft_log.disable_synchronization_unsafe": {
      "value": "false",
      "type": "b",
      "description": "set to true to disable synchronization on Raft log writes to persistent storage. Setting to true risks data loss or data corruption on server crashes. The setting is meant for internal testing only and SHOULD NOT be used in production."
    },
    "kv.range.backpressure_range_size_multiplier": {
      "value": "2",
      "type": "f",
      "description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
    },
    "kv.range_descriptor_cache.size": {
      "value": "1000000",
      "type": "i",
      "description": "maximum number of entries in the range descriptor and leaseholder caches"
    },
    "kv.range_merge.queue_enabled": {
      "value": "true",
      "type": "b",
      "description": "whether the automatic merge queue is enabled"
    },
    "kv.range_merge.queue_interval": {
      "value": "1s",
      "type": "d",
      "description": "how long the merge queue waits between processing replicas (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.range_split.by_load_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow automatic splits of ranges based on where load is concentrated"
    },
    "kv.range_split.load_qps_threshold": {
      "value": "2500",
      "type": "i",
      "description": "the QPS over which, the range becomes a candidate for load based splitting"
    },
    "kv.rangefeed.concurrent_catchup_iterators": {
      "value": "64",
      "type": "i",
      "description": "number of rangefeeds catchup iterators a store will allow concurrently before queueing"
    },
    "kv.rangefeed.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, rangefeed registration is enabled"
    },
    "kv.replication_reports.interval": {
      "value": "1m0s",
      "type": "d",
      "description": "the frequency for generating the replication_constraint_stats, replication_stats_report and replication_critical_localities reports (set to 0 to disable)"
    },
    "kv.snapshot_rebalance.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for rebalance and upreplication snapshots"
    },
    "kv.snapshot_recovery.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for recovery snapshots"
    },
    "kv.snapshot_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which snapshot SST writes must fsync"
    },
    "kv.transaction.max_intents_bytes": {
      "value": "262144",
      "type": "i",
      "description": "maximum number of bytes used to track write intents in transactions"
    },
    "kv.transaction.max_refresh_spans_bytes": {
      "value": "256000",
      "type": "i",
      "description": "maximum number of bytes used to track refresh spans in serializable transactions"
    },
    "kv.transaction.parallel_commits_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional commits will be parallelized with transactional writes"
    },
    "kv.transaction.write_pipelining_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional writes are pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_batch_size": {
      "value": "128",
      "type": "i",
      "description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_outstanding_size": {
      "value": "256 KiB",
      "type": "z",
      "description": "maximum number of bytes used to track in-flight pipelined writes before disabling pipelining"
    },
    "rocksdb.ingest_backpressure.l0_file_count_threshold": {
      "value": "20",
      "type": "i",
      "description": "number of L0 files after which to backpressure SST ingestions"
    },
    "rocksdb.ingest_backpressure.max_delay": {
      "value": "5s",
      "type": "d",
      "description": "maximum amount of time to backpressure a single SST ingestion"
    },
    "rocksdb.ingest_backpressure.pending_compaction_threshold": {
      "value": "2.0 GiB",
      "type": "z",
      "description": "pending compaction estimate above which to backpressure SST ingestions"
    },
    "rocksdb.min_wal_sync_interval": {
      "value": "0s",
      "type": "d",
      "description": "minimum duration between syncs of the RocksDB WAL"
    },
    "schemachanger.backfiller.buffer_increment": {
      "value": "32 MiB",
      "type": "z",
      "description": "the size by which the BulkAdder attempts to grow its buffer before flushing"
    },
    "schemachanger.backfiller.buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling index backfills"
    },
    "schemachanger.backfiller.max_buffer_size": {
      "value": "512 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling index backfills"
    },
    "schemachanger.backfiller.max_sst_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "target size for ingested files during backfills"
    },
    "schemachanger.bulk_index_backfill.batch_size": {
      "value": "50000",
      "type": "i",
      "description": "number of rows to process at a time during bulk index backfill"
    },
    "schemachanger.lease.duration": {
      "value": "5m0s",
      "type": "d",
      "description": "the duration of a schema change lease"
    },
    "schemachanger.lease.renew_fraction": {
      "value": "0.5",
      "type": "f",
      "description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
    },
    "server.clock.forward_jump_check_enabled": {
      "value": "false",
      "type": "b",
      "description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic"
    },
    "server.clock.persist_upper_bound_interval": {
      "value": "0s",
      "type": "d",
      "description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
    },
    "server.consistency_check.interval": {
      "value": "24h0m0s",
      "type": "d",
      "description": "the time between range consistency checks; set to 0 to disable consistency checking"
    },
    "server.declined_reservation_timeout": {
      "value": "1s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
    },
    "server.eventlog.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "if nonzero, event log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.failed_reservation_timeout": {
      "value": "5s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
    },
    "server.goroutine_dump.num_goroutines_threshold": {
      "value": "1000",
      "type": "i",
      "description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
    },
    "server.goroutine_dump.total_dump_size_limit": {
      "value": "500 MiB",
      "type": "z",
      "description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
    },
    "server.heap_profile.max_profiles": {
      "value": "5",
      "type": "i",
      "description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept."
    },
    "server.host_based_authentication.configuration": {
      "type": "s",
      "description": "host-based authentication configuration to use during connection authentication"
    },
    "server.rangelog.ttl": {
      "value": "720h0m0s",
      "type": "d",
      "description": "if nonzero, range log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.remote_debugging.mode": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
    },
    "server.shutdown.drain_wait": {
      "value": "0s",
      "type": "d",
      "description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
    },
    "server.shutdown.query_wait": {
      "value": "10s",
      "type": "d",
      "description": "the server will wait for at least this amount of time for active queries to finish"
    },
    "server.time_until_store_dead": {
      "value": "5m0s",
      "type": "d",
      "description": "the time after which if there is no new gossiped information about a store, it is considered dead"
    },
    "server.web_session_timeout": {
      "value": "168h0m0s",
      "type": "d",
      "description": "the duration that a newly created web session will be valid"
    },
    "sql.defaults.default_int_size": {
      "value": "8",
      "type": "i",
      "description": "the size, in bytes, of an INT type"
    },
    "sql.defaults.distsql": {
      "value": "auto",
      "type": "e",
      "description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
    },
    "sql.defaults.experimental_optimizer_foreign_keys.enabled": {
      "value": "false",
      "type": "b",
      "description": "enables optimizer-driven foreign key checks by default"
    },
    "sql.defaults.reorder_joins_limit": {
      "value": "4",
      "type": "i",
      "description": "default number of joins to reorder"
    },
    "sql.defaults.results_buffer.size": {
      "value": "16 KiB",
      "type": "z",
      "description": "default size of the buffer that accumulates results for a statement or a batch of statements before they are sent to the client. This can be overridden on an individual connection with the 'results_buffer_size' parameter. Note that auto-retries generally only happen while no results have been delivered to the client, so reducing this size can increase the number of retriable errors a client receives. On the other hand, increasing the buffer size can increase the delay until the client receives the first result row. Updating the setting only affects new connections. Setting to 0 disables any buffering."
    },
    "sql.defaults.serial_normalization": {
      "value": "rowid",
      "type": "e",
      "description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
    },
    "sql.defaults.vectorize": {
      "value": "auto",
      "type": "e",
      "description": "default vectorize mode [off = 0, auto = 1, experimental_on = 2]"
    },
    "sql.defaults.vectorize_row_count_threshold": {
      "value": "1000",
      "type": "i",
      "description": "default vectorize row count threshold"
    },
    "sql.defaults.zigzag_join.enabled": {
      "value": "true",
      "type": "b",
      "description": "default value for enable_zigzag_join session setting; allows use of zig-zag join by default"
    },
    "sql.distsql.distribute_index_joins": {
      "value": "true",
      "type": "b",
      "description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
    },
    "sql.distsql.flow_stream_timeout": {
      "value": "10s",
      "type": "d",
      "description": "amount of time incoming streams wait for a flow to be set up before erroring out"
    },
    "sql.distsql.interleaved_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set we plan interleaved table joins instead of merge joins when possible"
    },
    "sql.distsql.max_running_flows": {
      "value": "500",
      "type": "i",
      "description": "maximum number of concurrent flows that can be run on a node"
    },
    "sql.distsql.merge_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, we plan merge joins when possible"
    },
    "sql.distsql.temp_storage.joins": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql joins"
    },
    "sql.distsql.temp_storage.sorts": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql sorts"
    },
    "sql.distsql.temp_storage.workmem": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
    },
    "sql.metrics.statement_details.dump_to_logs": {
      "value": "false",
      "type": "b",
      "description": "dump collected statement statistics to node logs when periodically cleared"
    },
    "sql.metrics.statement_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-statement query statistics"
    },
    "sql.metrics.statement_details.plan_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "periodically save a logical plan for each fingerprint"
    },
    "sql.metrics.statement_details.plan_collection.period": {
      "value": "5m0s",
      "type": "d",
      "description": "the time until a new logical plan is collected"
    },
    "sql.metrics.statement_details.threshold": {
      "value": "0s",
      "type": "d",
      "description": "minimum execution time to cause statistics to be collected"
    },
    "sql.metrics.transaction_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-application transaction statistics"
    },
    "sql.parallel_scans.enabled": {
      "value": "true",
      "type": "b",
      "description": "parallelizes scanning different ranges when the maximum result size can be deduced"
    },
    "sql.query_cache.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable the query cache"
    },
    "sql.stats.automatic_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "automatic statistics collection mode"
    },
    "sql.stats.automatic_collection.fraction_stale_rows": {
      "value": "0.2",
      "type": "f",
      "description": "target fraction of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.automatic_collection.max_fraction_idle": {
      "value": "0.9",
      "type": "f",
      "description": "maximum fraction of time that automatic statistics sampler processors are idle"
    },
    "sql.stats.automatic_collection.min_stale_rows": {
      "value": "500",
      "type": "i",
      "description": "target minimum number of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.histogram_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "histogram collection mode"
    },
    "sql.stats.max_timestamp_age": {
      "value": "5m0s",
      "type": "d",
      "description": "maximum age of timestamp during table statistics collection"
    },
    "sql.stats.post_events.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, an event is shown for every CREATE STATISTICS job"
    },
    "sql.tablecache.lease.refresh_limit": {
      "value": "50",
      "type": "i",
      "description": "maximum number of tables to periodically refresh leases for"
    },
    "sql.trace.log_statement_execute": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable logging of executed statements"
    },
    "sql.trace.session_eventlog.enabled": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable session tracing"
    },
    "sql.trace.txn.enable_threshold": {
      "value": "0s",
      "type": "d",
      "description": "duration beyond which all transactions are traced (set to 0 to disable)"
    },
    "timeseries.storage.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
    },
    "timeseries.storage.resolution_10s.ttl": {
      "value": "240h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
    },
    "trace.debug.enable": {
      "value": "false",
      "type": "b",
      "description": "if set, traces for recent requests can be seen in the /debug page"
    },
    "trace.lightstep.token": {
      "type": "s",
      "description": "if set, traces go to Lightstep using this token"
    },
    "trace.zipkin.collector": {
      "type": "s",
      "description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set"
    },
    "version": {
      "value": "19.2",
      "type": "m",
      "description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'"
    }
  }
}
`
const DefaultClusterSettings19_2_1 string = `
{
  "key_values": {
    "changefeed.experimental_poll_interval": {
      "value": "1s",
      "type": "d",
      "description": "polling interval for the prototype changefeed implementation (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "cloudstorage.gs.default.key": {
      "type": "s",
      "description": "if set, JSON key to use during Google Cloud Storage operations"
    },
    "cloudstorage.http.custom_ca": {
      "type": "s",
      "description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
    },
    "cloudstorage.timeout": {
      "value": "10m0s",
      "type": "d",
      "description": "the timeout for import/export storage operations"
    },
    "cluster.organization": {
      "type": "s",
      "description": "organization name"
    },
    "cluster.preserve_downgrade_option": {
      "type": "s",
      "description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
    },
    "compactor.enabled": {
      "value": "true",
      "type": "b",
      "description": "when false, the system will reclaim space occupied by deleted data less aggressively"
    },
    "compactor.max_record_age": {
      "value": "24h0m0s",
      "type": "d",
      "description": "discard suggestions not processed within this duration (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.min_interval": {
      "value": "15s",
      "type": "d",
      "description": "minimum time interval to wait before compacting (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_available_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the available logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_bytes": {
      "value": "256 MiB",
      "type": "z",
      "description": "minimum expected logical space reclamation required before considering an aggregated suggestion (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_used_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the used logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "debug.panic_on_failed_assertions": {
      "value": "false",
      "type": "b",
      "description": "panic when an assertion fails rather than reporting"
    },
    "diagnostics.forced_stat_reset.interval": {
      "value": "2h0m0s",
      "type": "d",
      "description": "interval after which pending diagnostics statistics should be discarded even if not reported"
    },
    "diagnostics.reporting.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable reporting diagnostic metrics to cockroach labs"
    },
    "diagnostics.reporting.interval": {
      "value": "1h0m0s",
      "type": "d",
      "description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
    },
    "diagnostics.reporting.send_crash_reports": {
      "value": "true",
      "type": "b",
      "description": "send crash and panic reports"
    },
    "external.graphite.endpoint": {
      "type": "s",
      "description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
    },
    "external.graphite.interval": {
      "value": "10s",
      "type": "d",
      "description": "the interval at which metrics are pushed to Graphite (if enabled)"
    },
    "jobs.registry.leniency": {
      "value": "1m0s",
      "type": "d",
      "description": "the amount of time to defer any attempts to reschedule a job"
    },
    "jobs.retention_time": {
      "value": "336h0m0s",
      "type": "d",
      "description": "the amount of time to retain records for completed jobs before"
    },
    "kv.allocator.lease_rebalancing_aggressiveness": {
      "value": "1",
      "type": "f",
      "description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
    },
    "kv.allocator.load_based_lease_rebalancing.enabled": {
      "value": "true",
      "type": "b",
      "description": "set to enable rebalancing of range leases based on load and latency"
    },
    "kv.allocator.load_based_rebalancing": {
      "value": "leases and replicas",
      "type": "e",
      "description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
    },
    "kv.allocator.qps_rebalance_threshold": {
      "value": "0.25",
      "type": "f",
      "description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
    },
    "kv.allocator.range_rebalance_threshold": {
      "value": "0.05",
      "type": "f",
      "description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
    },
    "kv.atomic_replication_changes.enabled": {
      "value": "true",
      "type": "b",
      "description": "use atomic replication changes"
    },
    "kv.bulk_ingest.batch_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "the maximum size of the payload in an AddSSTable request"
    },
    "kv.bulk_ingest.buffer_increment": {
      "value": "32 MiB",
      "type": "z",
      "description": "the size by which the BulkAdder attempts to grow its buffer before flushing"
    },
    "kv.bulk_ingest.index_buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling secondary index imports"
    },
    "kv.bulk_ingest.max_index_buffer_size": {
      "value": "512 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling secondary index imports"
    },
    "kv.bulk_ingest.max_pk_buffer_size": {
      "value": "128 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling primary index imports"
    },
    "kv.bulk_ingest.pk_buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling primary index imports"
    },
    "kv.bulk_io_write.addsstable_max_rate": {
      "value": "1.7976931348623157E+308",
      "type": "f",
      "description": "maximum number of AddSSTable requests per second for a single store"
    },
    "kv.bulk_io_write.concurrent_addsstable_requests": {
      "value": "1",
      "type": "i",
      "description": "number of AddSSTable requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_export_requests": {
      "value": "3",
      "type": "i",
      "description": "number of export requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_import_requests": {
      "value": "1",
      "type": "i",
      "description": "number of import requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.experimental_incremental_export_enabled": {
      "value": "false",
      "type": "b",
      "description": "use experimental time-bound file filter when exporting in BACKUP"
    },
    "kv.bulk_io_write.max_rate": {
      "value": "1.0 TiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
    },
    "kv.bulk_io_write.small_write_size": {
      "value": "400 KiB",
      "type": "z",
      "description": "size below which a 'bulk' write will be performed as a normal write instead"
    },
    "kv.bulk_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
    },
    "kv.closed_timestamp.close_fraction": {
      "value": "0.2",
      "type": "f",
      "description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
    },
    "kv.closed_timestamp.follower_reads_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
    },
    "kv.closed_timestamp.target_duration": {
      "value": "30s",
      "type": "d",
      "description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
    },
    "kv.follower_read.target_multiple": {
      "value": "3",
      "type": "f",
      "description": "if above 1, encourages the distsender to perform a read against the closest replica if a request is older than kv.closed_timestamp.target_duration * (1 + kv.closed_timestamp.close_fraction * this) less a clock uncertainty interval. This value also is used to create follower_timestamp(). (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.raft.command.max_size": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum size of a raft command"
    },
    "kv.raft_log.disable_synchronization_unsafe": {
      "value": "false",
      "type": "b",
      "description": "set to true to disable synchronization on Raft log writes to persistent storage. Setting to true risks data loss or data corruption on server crashes. The setting is meant for internal testing only and SHOULD NOT be used in production."
    },
    "kv.range.backpressure_range_size_multiplier": {
      "value": "2",
      "type": "f",
      "description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
    },
    "kv.range_descriptor_cache.size": {
      "value": "1000000",
      "type": "i",
      "description": "maximum number of entries in the range descriptor and leaseholder caches"
    },
    "kv.range_merge.queue_enabled": {
      "value": "true",
      "type": "b",
      "description": "whether the automatic merge queue is enabled"
    },
    "kv.range_merge.queue_interval": {
      "value": "1s",
      "type": "d",
      "description": "how long the merge queue waits between processing replicas (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.range_split.by_load_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow automatic splits of ranges based on where load is concentrated"
    },
    "kv.range_split.load_qps_threshold": {
      "value": "2500",
      "type": "i",
      "description": "the QPS over which, the range becomes a candidate for load based splitting"
    },
    "kv.rangefeed.concurrent_catchup_iterators": {
      "value": "64",
      "type": "i",
      "description": "number of rangefeeds catchup iterators a store will allow concurrently before queueing"
    },
    "kv.rangefeed.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, rangefeed registration is enabled"
    },
    "kv.replication_reports.interval": {
      "value": "1m0s",
      "type": "d",
      "description": "the frequency for generating the replication_constraint_stats, replication_stats_report and replication_critical_localities reports (set to 0 to disable)"
    },
    "kv.snapshot_rebalance.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for rebalance and upreplication snapshots"
    },
    "kv.snapshot_recovery.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for recovery snapshots"
    },
    "kv.snapshot_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which snapshot SST writes must fsync"
    },
    "kv.transaction.max_intents_bytes": {
      "value": "262144",
      "type": "i",
      "description": "maximum number of bytes used to track write intents in transactions"
    },
    "kv.transaction.max_refresh_spans_bytes": {
      "value": "256000",
      "type": "i",
      "description": "maximum number of bytes used to track refresh spans in serializable transactions"
    },
    "kv.transaction.parallel_commits_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional commits will be parallelized with transactional writes"
    },
    "kv.transaction.write_pipelining_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional writes are pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_batch_size": {
      "value": "128",
      "type": "i",
      "description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_outstanding_size": {
      "value": "256 KiB",
      "type": "z",
      "description": "maximum number of bytes used to track in-flight pipelined writes before disabling pipelining"
    },
    "rocksdb.ingest_backpressure.l0_file_count_threshold": {
      "value": "20",
      "type": "i",
      "description": "number of L0 files after which to backpressure SST ingestions"
    },
    "rocksdb.ingest_backpressure.max_delay": {
      "value": "5s",
      "type": "d",
      "description": "maximum amount of time to backpressure a single SST ingestion"
    },
    "rocksdb.ingest_backpressure.pending_compaction_threshold": {
      "value": "2.0 GiB",
      "type": "z",
      "description": "pending compaction estimate above which to backpressure SST ingestions"
    },
    "rocksdb.min_wal_sync_interval": {
      "value": "0s",
      "type": "d",
      "description": "minimum duration between syncs of the RocksDB WAL"
    },
    "schemachanger.backfiller.buffer_increment": {
      "value": "32 MiB",
      "type": "z",
      "description": "the size by which the BulkAdder attempts to grow its buffer before flushing"
    },
    "schemachanger.backfiller.buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling index backfills"
    },
    "schemachanger.backfiller.max_buffer_size": {
      "value": "512 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling index backfills"
    },
    "schemachanger.backfiller.max_sst_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "target size for ingested files during backfills"
    },
    "schemachanger.bulk_index_backfill.batch_size": {
      "value": "50000",
      "type": "i",
      "description": "number of rows to process at a time during bulk index backfill"
    },
    "schemachanger.lease.duration": {
      "value": "5m0s",
      "type": "d",
      "description": "the duration of a schema change lease"
    },
    "schemachanger.lease.renew_fraction": {
      "value": "0.5",
      "type": "f",
      "description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
    },
    "server.clock.forward_jump_check_enabled": {
      "value": "false",
      "type": "b",
      "description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic"
    },
    "server.clock.persist_upper_bound_interval": {
      "value": "0s",
      "type": "d",
      "description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
    },
    "server.consistency_check.interval": {
      "value": "24h0m0s",
      "type": "d",
      "description": "the time between range consistency checks; set to 0 to disable consistency checking"
    },
    "server.declined_reservation_timeout": {
      "value": "1s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
    },
    "server.eventlog.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "if nonzero, event log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.failed_reservation_timeout": {
      "value": "5s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
    },
    "server.goroutine_dump.num_goroutines_threshold": {
      "value": "1000",
      "type": "i",
      "description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
    },
    "server.goroutine_dump.total_dump_size_limit": {
      "value": "500 MiB",
      "type": "z",
      "description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
    },
    "server.heap_profile.max_profiles": {
      "value": "5",
      "type": "i",
      "description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept."
    },
    "server.host_based_authentication.configuration": {
      "type": "s",
      "description": "host-based authentication configuration to use during connection authentication"
    },
    "server.rangelog.ttl": {
      "value": "720h0m0s",
      "type": "d",
      "description": "if nonzero, range log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.remote_debugging.mode": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
    },
    "server.shutdown.drain_wait": {
      "value": "0s",
      "type": "d",
      "description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
    },
    "server.shutdown.query_wait": {
      "value": "10s",
      "type": "d",
      "description": "the server will wait for at least this amount of time for active queries to finish"
    },
    "server.time_until_store_dead": {
      "value": "5m0s",
      "type": "d",
      "description": "the time after which if there is no new gossiped information about a store, it is considered dead"
    },
    "server.web_session_timeout": {
      "value": "168h0m0s",
      "type": "d",
      "description": "the duration that a newly created web session will be valid"
    },
    "sql.defaults.default_int_size": {
      "value": "8",
      "type": "i",
      "description": "the size, in bytes, of an INT type"
    },
    "sql.defaults.distsql": {
      "value": "auto",
      "type": "e",
      "description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
    },
    "sql.defaults.experimental_optimizer_foreign_keys.enabled": {
      "value": "false",
      "type": "b",
      "description": "enables optimizer-driven foreign key checks by default"
    },
    "sql.defaults.reorder_joins_limit": {
      "value": "4",
      "type": "i",
      "description": "default number of joins to reorder"
    },
    "sql.defaults.results_buffer.size": {
      "value": "16 KiB",
      "type": "z",
      "description": "default size of the buffer that accumulates results for a statement or a batch of statements before they are sent to the client. This can be overridden on an individual connection with the 'results_buffer_size' parameter. Note that auto-retries generally only happen while no results have been delivered to the client, so reducing this size can increase the number of retriable errors a client receives. On the other hand, increasing the buffer size can increase the delay until the client receives the first result row. Updating the setting only affects new connections. Setting to 0 disables any buffering."
    },
    "sql.defaults.serial_normalization": {
      "value": "rowid",
      "type": "e",
      "description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
    },
    "sql.defaults.vectorize": {
      "value": "auto",
      "type": "e",
      "description": "default vectorize mode [off = 0, auto = 1, experimental_on = 2]"
    },
    "sql.defaults.vectorize_row_count_threshold": {
      "value": "1000",
      "type": "i",
      "description": "default vectorize row count threshold"
    },
    "sql.distsql.distribute_index_joins": {
      "value": "true",
      "type": "b",
      "description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
    },
    "sql.distsql.flow_stream_timeout": {
      "value": "10s",
      "type": "d",
      "description": "amount of time incoming streams wait for a flow to be set up before erroring out"
    },
    "sql.distsql.interleaved_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set we plan interleaved table joins instead of merge joins when possible"
    },
    "sql.distsql.max_running_flows": {
      "value": "500",
      "type": "i",
      "description": "maximum number of concurrent flows that can be run on a node"
    },
    "sql.distsql.merge_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, we plan merge joins when possible"
    },
    "sql.distsql.temp_storage.joins": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql joins"
    },
    "sql.distsql.temp_storage.sorts": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql sorts"
    },
    "sql.distsql.temp_storage.workmem": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
    },
    "sql.metrics.statement_details.dump_to_logs": {
      "value": "false",
      "type": "b",
      "description": "dump collected statement statistics to node logs when periodically cleared"
    },
    "sql.metrics.statement_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-statement query statistics"
    },
    "sql.metrics.statement_details.plan_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "periodically save a logical plan for each fingerprint"
    },
    "sql.metrics.statement_details.plan_collection.period": {
      "value": "5m0s",
      "type": "d",
      "description": "the time until a new logical plan is collected"
    },
    "sql.metrics.statement_details.threshold": {
      "value": "0s",
      "type": "d",
      "description": "minimum execution time to cause statistics to be collected"
    },
    "sql.metrics.transaction_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-application transaction statistics"
    },
    "sql.parallel_scans.enabled": {
      "value": "true",
      "type": "b",
      "description": "parallelizes scanning different ranges when the maximum result size can be deduced"
    },
    "sql.query_cache.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable the query cache"
    },
    "sql.stats.automatic_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "automatic statistics collection mode"
    },
    "sql.stats.automatic_collection.fraction_stale_rows": {
      "value": "0.2",
      "type": "f",
      "description": "target fraction of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.automatic_collection.max_fraction_idle": {
      "value": "0.9",
      "type": "f",
      "description": "maximum fraction of time that automatic statistics sampler processors are idle"
    },
    "sql.stats.automatic_collection.min_stale_rows": {
      "value": "500",
      "type": "i",
      "description": "target minimum number of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.histogram_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "histogram collection mode"
    },
    "sql.stats.max_timestamp_age": {
      "value": "5m0s",
      "type": "d",
      "description": "maximum age of timestamp during table statistics collection"
    },
    "sql.stats.post_events.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, an event is shown for every CREATE STATISTICS job"
    },
    "sql.tablecache.lease.refresh_limit": {
      "value": "50",
      "type": "i",
      "description": "maximum number of tables to periodically refresh leases for"
    },
    "sql.trace.log_statement_execute": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable logging of executed statements"
    },
    "sql.trace.session_eventlog.enabled": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable session tracing"
    },
    "sql.trace.txn.enable_threshold": {
      "value": "0s",
      "type": "d",
      "description": "duration beyond which all transactions are traced (set to 0 to disable)"
    },
    "timeseries.storage.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
    },
    "timeseries.storage.resolution_10s.ttl": {
      "value": "240h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
    },
    "timeseries.storage.resolution_30m.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
    },
    "trace.debug.enable": {
      "value": "false",
      "type": "b",
      "description": "if set, traces for recent requests can be seen in the /debug page"
    },
    "trace.lightstep.token": {
      "type": "s",
      "description": "if set, traces go to Lightstep using this token"
    },
    "trace.zipkin.collector": {
      "type": "s",
      "description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set"
    },
    "version": {
      "value": "19.2",
      "type": "m",
      "description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'"
    }
  }
}
`
const DefaultClusterSettings19_2_0 string = `
{
  "key_values": {
    "changefeed.experimental_poll_interval": {
      "value": "1s",
      "type": "d",
      "description": "polling interval for the prototype changefeed implementation (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "cloudstorage.gs.default.key": {
      "type": "s",
      "description": "if set, JSON key to use during Google Cloud Storage operations"
    },
    "cloudstorage.http.custom_ca": {
      "type": "s",
      "description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
    },
    "cloudstorage.timeout": {
      "value": "10m0s",
      "type": "d",
      "description": "the timeout for import/export storage operations"
    },
    "cluster.organization": {
      "type": "s",
      "description": "organization name"
    },
    "cluster.preserve_downgrade_option": {
      "type": "s",
      "description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
    },
    "compactor.enabled": {
      "value": "true",
      "type": "b",
      "description": "when false, the system will reclaim space occupied by deleted data less aggressively"
    },
    "compactor.max_record_age": {
      "value": "24h0m0s",
      "type": "d",
      "description": "discard suggestions not processed within this duration (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.min_interval": {
      "value": "15s",
      "type": "d",
      "description": "minimum time interval to wait before compacting (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_available_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the available logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_bytes": {
      "value": "256 MiB",
      "type": "z",
      "description": "minimum expected logical space reclamation required before considering an aggregated suggestion (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_used_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the used logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "debug.panic_on_failed_assertions": {
      "value": "false",
      "type": "b",
      "description": "panic when an assertion fails rather than reporting"
    },
    "diagnostics.forced_stat_reset.interval": {
      "value": "2h0m0s",
      "type": "d",
      "description": "interval after which pending diagnostics statistics should be discarded even if not reported"
    },
    "diagnostics.reporting.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable reporting diagnostic metrics to cockroach labs"
    },
    "diagnostics.reporting.interval": {
      "value": "1h0m0s",
      "type": "d",
      "description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
    },
    "diagnostics.reporting.send_crash_reports": {
      "value": "true",
      "type": "b",
      "description": "send crash and panic reports"
    },
    "external.graphite.endpoint": {
      "type": "s",
      "description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
    },
    "external.graphite.interval": {
      "value": "10s",
      "type": "d",
      "description": "the interval at which metrics are pushed to Graphite (if enabled)"
    },
    "jobs.registry.leniency": {
      "value": "1m0s",
      "type": "d",
      "description": "the amount of time to defer any attempts to reschedule a job"
    },
    "jobs.retention_time": {
      "value": "336h0m0s",
      "type": "d",
      "description": "the amount of time to retain records for completed jobs before"
    },
    "kv.allocator.lease_rebalancing_aggressiveness": {
      "value": "1",
      "type": "f",
      "description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
    },
    "kv.allocator.load_based_lease_rebalancing.enabled": {
      "value": "true",
      "type": "b",
      "description": "set to enable rebalancing of range leases based on load and latency"
    },
    "kv.allocator.load_based_rebalancing": {
      "value": "leases and replicas",
      "type": "e",
      "description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
    },
    "kv.allocator.qps_rebalance_threshold": {
      "value": "0.25",
      "type": "f",
      "description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
    },
    "kv.allocator.range_rebalance_threshold": {
      "value": "0.05",
      "type": "f",
      "description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
    },
    "kv.atomic_replication_changes.enabled": {
      "value": "true",
      "type": "b",
      "description": "use atomic replication changes"
    },
    "kv.bulk_ingest.batch_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "the maximum size of the payload in an AddSSTable request"
    },
    "kv.bulk_ingest.buffer_increment": {
      "value": "32 MiB",
      "type": "z",
      "description": "the size by which the BulkAdder attempts to grow its buffer before flushing"
    },
    "kv.bulk_ingest.index_buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling secondary index imports"
    },
    "kv.bulk_ingest.max_index_buffer_size": {
      "value": "512 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling secondary index imports"
    },
    "kv.bulk_ingest.max_pk_buffer_size": {
      "value": "128 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling primary index imports"
    },
    "kv.bulk_ingest.pk_buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling primary index imports"
    },
    "kv.bulk_io_write.addsstable_max_rate": {
      "value": "1.7976931348623157E+308",
      "type": "f",
      "description": "maximum number of AddSSTable requests per second for a single store"
    },
    "kv.bulk_io_write.concurrent_addsstable_requests": {
      "value": "1",
      "type": "i",
      "description": "number of AddSSTable requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_export_requests": {
      "value": "3",
      "type": "i",
      "description": "number of export requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_import_requests": {
      "value": "1",
      "type": "i",
      "description": "number of import requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.experimental_incremental_export_enabled": {
      "value": "false",
      "type": "b",
      "description": "use experimental time-bound file filter when exporting in BACKUP"
    },
    "kv.bulk_io_write.max_rate": {
      "value": "1.0 TiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
    },
    "kv.bulk_io_write.small_write_size": {
      "value": "400 KiB",
      "type": "z",
      "description": "size below which a 'bulk' write will be performed as a normal write instead"
    },
    "kv.bulk_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
    },
    "kv.closed_timestamp.close_fraction": {
      "value": "0.2",
      "type": "f",
      "description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
    },
    "kv.closed_timestamp.follower_reads_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
    },
    "kv.closed_timestamp.target_duration": {
      "value": "30s",
      "type": "d",
      "description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
    },
    "kv.follower_read.target_multiple": {
      "value": "3",
      "type": "f",
      "description": "if above 1, encourages the distsender to perform a read against the closest replica if a request is older than kv.closed_timestamp.target_duration * (1 + kv.closed_timestamp.close_fraction * this) less a clock uncertainty interval. This value also is used to create follower_timestamp(). (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.raft.command.max_size": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum size of a raft command"
    },
    "kv.raft_log.disable_synchronization_unsafe": {
      "value": "false",
      "type": "b",
      "description": "set to true to disable synchronization on Raft log writes to persistent storage. Setting to true risks data loss or data corruption on server crashes. The setting is meant for internal testing only and SHOULD NOT be used in production."
    },
    "kv.range.backpressure_range_size_multiplier": {
      "value": "2",
      "type": "f",
      "description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
    },
    "kv.range_descriptor_cache.size": {
      "value": "1000000",
      "type": "i",
      "description": "maximum number of entries in the range descriptor and leaseholder caches"
    },
    "kv.range_merge.queue_enabled": {
      "value": "true",
      "type": "b",
      "description": "whether the automatic merge queue is enabled"
    },
    "kv.range_merge.queue_interval": {
      "value": "1s",
      "type": "d",
      "description": "how long the merge queue waits between processing replicas (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.range_split.by_load_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow automatic splits of ranges based on where load is concentrated"
    },
    "kv.range_split.load_qps_threshold": {
      "value": "2500",
      "type": "i",
      "description": "the QPS over which, the range becomes a candidate for load based splitting"
    },
    "kv.rangefeed.concurrent_catchup_iterators": {
      "value": "64",
      "type": "i",
      "description": "number of rangefeeds catchup iterators a store will allow concurrently before queueing"
    },
    "kv.rangefeed.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, rangefeed registration is enabled"
    },
    "kv.replication_reports.interval": {
      "value": "1m0s",
      "type": "d",
      "description": "the frequency for generating the replication_constraint_stats, replication_stats_report and replication_critical_localities reports (set to 0 to disable)"
    },
    "kv.snapshot_rebalance.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for rebalance and upreplication snapshots"
    },
    "kv.snapshot_recovery.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for recovery snapshots"
    },
    "kv.snapshot_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which snapshot SST writes must fsync"
    },
    "kv.transaction.max_intents_bytes": {
      "value": "262144",
      "type": "i",
      "description": "maximum number of bytes used to track write intents in transactions"
    },
    "kv.transaction.max_refresh_spans_bytes": {
      "value": "256000",
      "type": "i",
      "description": "maximum number of bytes used to track refresh spans in serializable transactions"
    },
    "kv.transaction.parallel_commits_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional commits will be parallelized with transactional writes"
    },
    "kv.transaction.write_pipelining_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional writes are pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_batch_size": {
      "value": "128",
      "type": "i",
      "description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_outstanding_size": {
      "value": "256 KiB",
      "type": "z",
      "description": "maximum number of bytes used to track in-flight pipelined writes before disabling pipelining"
    },
    "rocksdb.ingest_backpressure.l0_file_count_threshold": {
      "value": "20",
      "type": "i",
      "description": "number of L0 files after which to backpressure SST ingestions"
    },
    "rocksdb.ingest_backpressure.max_delay": {
      "value": "5s",
      "type": "d",
      "description": "maximum amount of time to backpressure a single SST ingestion"
    },
    "rocksdb.ingest_backpressure.pending_compaction_threshold": {
      "value": "2.0 GiB",
      "type": "z",
      "description": "pending compaction estimate above which to backpressure SST ingestions"
    },
    "rocksdb.min_wal_sync_interval": {
      "value": "0s",
      "type": "d",
      "description": "minimum duration between syncs of the RocksDB WAL"
    },
    "schemachanger.backfiller.buffer_increment": {
      "value": "32 MiB",
      "type": "z",
      "description": "the size by which the BulkAdder attempts to grow its buffer before flushing"
    },
    "schemachanger.backfiller.buffer_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the initial size of the BulkAdder buffer handling index backfills"
    },
    "schemachanger.backfiller.max_buffer_size": {
      "value": "512 MiB",
      "type": "z",
      "description": "the maximum size of the BulkAdder buffer handling index backfills"
    },
    "schemachanger.backfiller.max_sst_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "target size for ingested files during backfills"
    },
    "schemachanger.bulk_index_backfill.batch_size": {
      "value": "50000",
      "type": "i",
      "description": "number of rows to process at a time during bulk index backfill"
    },
    "schemachanger.lease.duration": {
      "value": "5m0s",
      "type": "d",
      "description": "the duration of a schema change lease"
    },
    "schemachanger.lease.renew_fraction": {
      "value": "0.5",
      "type": "f",
      "description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
    },
    "server.clock.forward_jump_check_enabled": {
      "value": "false",
      "type": "b",
      "description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic"
    },
    "server.clock.persist_upper_bound_interval": {
      "value": "0s",
      "type": "d",
      "description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
    },
    "server.consistency_check.interval": {
      "value": "24h0m0s",
      "type": "d",
      "description": "the time between range consistency checks; set to 0 to disable consistency checking"
    },
    "server.declined_reservation_timeout": {
      "value": "1s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
    },
    "server.eventlog.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "if nonzero, event log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.failed_reservation_timeout": {
      "value": "5s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
    },
    "server.goroutine_dump.num_goroutines_threshold": {
      "value": "1000",
      "type": "i",
      "description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
    },
    "server.goroutine_dump.total_dump_size_limit": {
      "value": "500 MiB",
      "type": "z",
      "description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
    },
    "server.heap_profile.max_profiles": {
      "value": "5",
      "type": "i",
      "description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept."
    },
    "server.host_based_authentication.configuration": {
      "type": "s",
      "description": "host-based authentication configuration to use during connection authentication"
    },
    "server.rangelog.ttl": {
      "value": "720h0m0s",
      "type": "d",
      "description": "if nonzero, range log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.remote_debugging.mode": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
    },
    "server.shutdown.drain_wait": {
      "value": "0s",
      "type": "d",
      "description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
    },
    "server.shutdown.query_wait": {
      "value": "10s",
      "type": "d",
      "description": "the server will wait for at least this amount of time for active queries to finish"
    },
    "server.time_until_store_dead": {
      "value": "5m0s",
      "type": "d",
      "description": "the time after which if there is no new gossiped information about a store, it is considered dead"
    },
    "server.web_session_timeout": {
      "value": "168h0m0s",
      "type": "d",
      "description": "the duration that a newly created web session will be valid"
    },
    "sql.defaults.default_int_size": {
      "value": "8",
      "type": "i",
      "description": "the size, in bytes, of an INT type"
    },
    "sql.defaults.distsql": {
      "value": "auto",
      "type": "e",
      "description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
    },
    "sql.defaults.experimental_optimizer_foreign_keys.enabled": {
      "value": "false",
      "type": "b",
      "description": "enables optimizer-driven foreign key checks by default"
    },
    "sql.defaults.reorder_joins_limit": {
      "value": "4",
      "type": "i",
      "description": "default number of joins to reorder"
    },
    "sql.defaults.results_buffer.size": {
      "value": "16 KiB",
      "type": "z",
      "description": "default size of the buffer that accumulates results for a statement or a batch of statements before they are sent to the client. This can be overridden on an individual connection with the 'results_buffer_size' parameter. Note that auto-retries generally only happen while no results have been delivered to the client, so reducing this size can increase the number of retriable errors a client receives. On the other hand, increasing the buffer size can increase the delay until the client receives the first result row. Updating the setting only affects new connections. Setting to 0 disables any buffering."
    },
    "sql.defaults.serial_normalization": {
      "value": "rowid",
      "type": "e",
      "description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
    },
    "sql.defaults.vectorize": {
      "value": "auto",
      "type": "e",
      "description": "default vectorize mode [off = 0, auto = 1, experimental_on = 2]"
    },
    "sql.defaults.vectorize_row_count_threshold": {
      "value": "1000",
      "type": "i",
      "description": "default vectorize row count threshold"
    },
    "sql.distsql.distribute_index_joins": {
      "value": "true",
      "type": "b",
      "description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
    },
    "sql.distsql.flow_stream_timeout": {
      "value": "10s",
      "type": "d",
      "description": "amount of time incoming streams wait for a flow to be set up before erroring out"
    },
    "sql.distsql.interleaved_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set we plan interleaved table joins instead of merge joins when possible"
    },
    "sql.distsql.max_running_flows": {
      "value": "500",
      "type": "i",
      "description": "maximum number of concurrent flows that can be run on a node"
    },
    "sql.distsql.merge_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, we plan merge joins when possible"
    },
    "sql.distsql.temp_storage.joins": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql joins"
    },
    "sql.distsql.temp_storage.sorts": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql sorts"
    },
    "sql.distsql.temp_storage.workmem": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
    },
    "sql.metrics.statement_details.dump_to_logs": {
      "value": "false",
      "type": "b",
      "description": "dump collected statement statistics to node logs when periodically cleared"
    },
    "sql.metrics.statement_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-statement query statistics"
    },
    "sql.metrics.statement_details.plan_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "periodically save a logical plan for each fingerprint"
    },
    "sql.metrics.statement_details.plan_collection.period": {
      "value": "5m0s",
      "type": "d",
      "description": "the time until a new logical plan is collected"
    },
    "sql.metrics.statement_details.threshold": {
      "value": "0s",
      "type": "d",
      "description": "minimum execution time to cause statistics to be collected"
    },
    "sql.metrics.transaction_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-application transaction statistics"
    },
    "sql.parallel_scans.enabled": {
      "value": "true",
      "type": "b",
      "description": "parallelizes scanning different ranges when the maximum result size can be deduced"
    },
    "sql.query_cache.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable the query cache"
    },
    "sql.stats.automatic_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "automatic statistics collection mode"
    },
    "sql.stats.automatic_collection.fraction_stale_rows": {
      "value": "0.2",
      "type": "f",
      "description": "target fraction of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.automatic_collection.max_fraction_idle": {
      "value": "0.9",
      "type": "f",
      "description": "maximum fraction of time that automatic statistics sampler processors are idle"
    },
    "sql.stats.automatic_collection.min_stale_rows": {
      "value": "500",
      "type": "i",
      "description": "target minimum number of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.histogram_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "histogram collection mode"
    },
    "sql.stats.max_timestamp_age": {
      "value": "5m0s",
      "type": "d",
      "description": "maximum age of timestamp during table statistics collection"
    },
    "sql.stats.post_events.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, an event is shown for every CREATE STATISTICS job"
    },
    "sql.tablecache.lease.refresh_limit": {
      "value": "50",
      "type": "i",
      "description": "maximum number of tables to periodically refresh leases for"
    },
    "sql.trace.log_statement_execute": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable logging of executed statements"
    },
    "sql.trace.session_eventlog.enabled": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable session tracing"
    },
    "sql.trace.txn.enable_threshold": {
      "value": "0s",
      "type": "d",
      "description": "duration beyond which all transactions are traced (set to 0 to disable)"
    },
    "timeseries.storage.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
    },
    "timeseries.storage.resolution_10s.ttl": {
      "value": "240h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
    },
    "timeseries.storage.resolution_30m.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
    },
    "trace.debug.enable": {
      "value": "false",
      "type": "b",
      "description": "if set, traces for recent requests can be seen in the /debug page"
    },
    "trace.lightstep.token": {
      "type": "s",
      "description": "if set, traces go to Lightstep using this token"
    },
    "trace.zipkin.collector": {
      "type": "s",
      "description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set"
    },
    "version": {
      "value": "19.2",
      "type": "m",
      "description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'"
    }
  }
}
`
const DefaultClusterSettings19_1_7 string = `
{
  "key_values": {
    "changefeed.experimental_poll_interval": {
      "value": "1s",
      "type": "d",
      "description": "polling interval for the prototype changefeed implementation (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "changefeed.push.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, changed are pushed instead of pulled. This requires the kv.rangefeed.enabled setting. See https://www.cockroachlabs.com/docs/v19.1/change-data-capture.html#enable-rangefeeds-to-reduce-latency"
    },
    "cloudstorage.gs.default.key": {
      "type": "s",
      "description": "if set, JSON key to use during Google Cloud Storage operations"
    },
    "cloudstorage.http.custom_ca": {
      "type": "s",
      "description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
    },
    "cloudstorage.timeout": {
      "value": "10m0s",
      "type": "d",
      "description": "the timeout for import/export storage operations"
    },
    "cluster.organization": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "organization name"
    },
    "cluster.preserve_downgrade_option": {
      "type": "s",
      "description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
    },
    "compactor.enabled": {
      "value": "true",
      "type": "b",
      "description": "when false, the system will reclaim space occupied by deleted data less aggressively"
    },
    "compactor.max_record_age": {
      "value": "24h0m0s",
      "type": "d",
      "description": "discard suggestions not processed within this duration (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.min_interval": {
      "value": "15s",
      "type": "d",
      "description": "minimum time interval to wait before compacting (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_available_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the available logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_bytes": {
      "value": "256 MiB",
      "type": "z",
      "description": "minimum expected logical space reclamation required before considering an aggregated suggestion (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_used_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the used logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "debug.panic_on_failed_assertions": {
      "value": "false",
      "type": "b",
      "description": "panic when an assertion fails rather than reporting"
    },
    "diagnostics.forced_stat_reset.interval": {
      "value": "2h0m0s",
      "type": "d",
      "description": "interval after which pending diagnostics statistics should be discarded even if not reported"
    },
    "diagnostics.reporting.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable reporting diagnostic metrics to cockroach labs"
    },
    "diagnostics.reporting.interval": {
      "value": "1h0m0s",
      "type": "d",
      "description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
    },
    "diagnostics.reporting.send_crash_reports": {
      "value": "true",
      "type": "b",
      "description": "send crash and panic reports"
    },
    "external.graphite.endpoint": {
      "type": "s",
      "description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
    },
    "external.graphite.interval": {
      "value": "10s",
      "type": "d",
      "description": "the interval at which metrics are pushed to Graphite (if enabled)"
    },
    "jobs.registry.leniency": {
      "value": "1m0s",
      "type": "d",
      "description": "the amount of time to defer any attempts to reschedule a job"
    },
    "jobs.retention_time": {
      "value": "336h0m0s",
      "type": "d",
      "description": "the amount of time to retain records for completed jobs before"
    },
    "kv.allocator.lease_rebalancing_aggressiveness": {
      "value": "1",
      "type": "f",
      "description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
    },
    "kv.allocator.load_based_lease_rebalancing.enabled": {
      "value": "true",
      "type": "b",
      "description": "set to enable rebalancing of range leases based on load and latency"
    },
    "kv.allocator.load_based_rebalancing": {
      "value": "2",
      "type": "e",
      "description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
    },
    "kv.allocator.qps_rebalance_threshold": {
      "value": "0.25",
      "type": "f",
      "description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
    },
    "kv.allocator.range_rebalance_threshold": {
      "value": "0.05",
      "type": "f",
      "description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
    },
    "kv.bulk_io_write.addsstable_max_rate": {
      "value": "1.7976931348623157E+308",
      "type": "f",
      "description": "maximum number of AddSSTable requests per second for a single store"
    },
    "kv.bulk_io_write.concurrent_addsstable_requests": {
      "value": "1",
      "type": "i",
      "description": "number of AddSSTable requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_export_requests": {
      "value": "3",
      "type": "i",
      "description": "number of export requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_import_requests": {
      "value": "1",
      "type": "i",
      "description": "number of import requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.max_rate": {
      "value": "1.0 TiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
    },
    "kv.bulk_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
    },
    "kv.closed_timestamp.close_fraction": {
      "value": "0.2",
      "type": "f",
      "description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
    },
    "kv.closed_timestamp.follower_reads_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
    },
    "kv.closed_timestamp.target_duration": {
      "value": "30s",
      "type": "d",
      "description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
    },
    "kv.follower_read.target_multiple": {
      "value": "3",
      "type": "f",
      "description": "if above 1, encourages the distsender to perform a read against the closest replica if a request is older than kv.closed_timestamp.target_duration * (1 + kv.closed_timestamp.close_fraction * this) less a clock uncertainty interval. This value also is used to create follower_timestamp(). (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.import.batch_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the maximum size of the payload in an AddSSTable request (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.raft.command.max_size": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum size of a raft command"
    },
    "kv.raft_log.disable_synchronization_unsafe": {
      "value": "false",
      "type": "b",
      "description": "set to true to disable synchronization on Raft log writes to persistent storage. Setting to true risks data loss or data corruption on server crashes. The setting is meant for internal testing only and SHOULD NOT be used in production."
    },
    "kv.range.backpressure_range_size_multiplier": {
      "value": "2",
      "type": "f",
      "description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
    },
    "kv.range_descriptor_cache.size": {
      "value": "1000000",
      "type": "i",
      "description": "maximum number of entries in the range descriptor and leaseholder caches"
    },
    "kv.range_merge.queue_enabled": {
      "value": "true",
      "type": "b",
      "description": "whether the automatic merge queue is enabled"
    },
    "kv.range_merge.queue_interval": {
      "value": "1s",
      "type": "d",
      "description": "how long the merge queue waits between processing replicas (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.range_split.by_load_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow automatic splits of ranges based on where load is concentrated"
    },
    "kv.range_split.load_qps_threshold": {
      "value": "250",
      "type": "i",
      "description": "the QPS over which, the range becomes a candidate for load based splitting"
    },
    "kv.rangefeed.concurrent_catchup_iterators": {
      "value": "64",
      "type": "i",
      "description": "number of rangefeeds catchup iterators a store will allow concurrently before queueing"
    },
    "kv.rangefeed.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, rangefeed registration is enabled"
    },
    "kv.snapshot_rebalance.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for rebalance and upreplication snapshots"
    },
    "kv.snapshot_recovery.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for recovery snapshots"
    },
    "kv.transaction.max_intents_bytes": {
      "value": "262144",
      "type": "i",
      "description": "maximum number of bytes used to track write intents in transactions"
    },
    "kv.transaction.max_refresh_spans_bytes": {
      "value": "256000",
      "type": "i",
      "description": "maximum number of bytes used to track refresh spans in serializable transactions"
    },
    "kv.transaction.write_pipelining_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional writes are pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_batch_size": {
      "value": "128",
      "type": "i",
      "description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_outstanding_size": {
      "value": "256 KiB",
      "type": "z",
      "description": "maximum number of bytes used to track in-flight pipelined writes before disabling pipelining"
    },
    "rocksdb.ingest_backpressure.delay_l0_file": {
      "value": "200ms",
      "type": "d",
      "description": "delay to add to SST ingestions per file in L0 over the configured limit"
    },
    "rocksdb.ingest_backpressure.l0_file_count_threshold": {
      "value": "20",
      "type": "i",
      "description": "number of L0 files after which to backpressure SST ingestions"
    },
    "rocksdb.ingest_backpressure.max_delay": {
      "value": "5s",
      "type": "d",
      "description": "maximum amount of time to backpressure a single SST ingestion"
    },
    "rocksdb.ingest_backpressure.pending_compaction_threshold": {
      "value": "64 GiB",
      "type": "z",
      "description": "pending compaction estimate above which to backpressure SST ingestions"
    },
    "rocksdb.min_wal_sync_interval": {
      "value": "0s",
      "type": "d",
      "description": "minimum duration between syncs of the RocksDB WAL"
    },
    "schemachanger.backfiller.buffer_size": {
      "value": "196 MiB",
      "type": "z",
      "description": "amount to buffer in memory during backfills"
    },
    "schemachanger.backfiller.max_sst_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "target size for ingested files during backfills"
    },
    "schemachanger.bulk_index_backfill.batch_size": {
      "value": "5000",
      "type": "i",
      "description": "number of rows to process at a time during bulk index backfill"
    },
    "schemachanger.bulk_index_backfill.enabled": {
      "value": "true",
      "type": "b",
      "description": "backfill indexes in bulk via addsstable"
    },
    "schemachanger.lease.duration": {
      "value": "5m0s",
      "type": "d",
      "description": "the duration of a schema change lease"
    },
    "schemachanger.lease.renew_fraction": {
      "value": "0.5",
      "type": "f",
      "description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
    },
    "server.clock.forward_jump_check_enabled": {
      "value": "false",
      "type": "b",
      "description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic"
    },
    "server.clock.persist_upper_bound_interval": {
      "value": "0s",
      "type": "d",
      "description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
    },
    "server.consistency_check.interval": {
      "value": "24h0m0s",
      "type": "d",
      "description": "the time between range consistency checks; set to 0 to disable consistency checking"
    },
    "server.declined_reservation_timeout": {
      "value": "1s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
    },
    "server.eventlog.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "if nonzero, event log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.failed_reservation_timeout": {
      "value": "5s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
    },
    "server.goroutine_dump.num_goroutines_threshold": {
      "value": "1000",
      "type": "i",
      "description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
    },
    "server.goroutine_dump.total_dump_size_limit": {
      "value": "500 MiB",
      "type": "z",
      "description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
    },
    "server.heap_profile.max_profiles": {
      "value": "5",
      "type": "i",
      "description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept."
    },
    "server.heap_profile.system_memory_threshold_fraction": {
      "value": "0.85",
      "type": "f",
      "description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
    },
    "server.host_based_authentication.configuration": {
      "type": "s",
      "description": "host-based authentication configuration to use during connection authentication"
    },
    "server.rangelog.ttl": {
      "value": "720h0m0s",
      "type": "d",
      "description": "if nonzero, range log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.remote_debugging.mode": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
    },
    "server.shutdown.drain_wait": {
      "value": "0s",
      "type": "d",
      "description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
    },
    "server.shutdown.query_wait": {
      "value": "10s",
      "type": "d",
      "description": "the server will wait for at least this amount of time for active queries to finish"
    },
    "server.time_until_store_dead": {
      "value": "5m0s",
      "type": "d",
      "description": "the time after which if there is no new gossiped information about a store, it is considered dead"
    },
    "server.web_session_timeout": {
      "value": "168h0m0s",
      "type": "d",
      "description": "the duration that a newly created web session will be valid"
    },
    "sql.defaults.default_int_size": {
      "value": "8",
      "type": "i",
      "description": "the size, in bytes, of an INT type"
    },
    "sql.defaults.distsql": {
      "value": "1",
      "type": "e",
      "description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
    },
    "sql.defaults.experimental_vectorize": {
      "value": "0",
      "type": "e",
      "description": "default experimental_vectorize mode [off = 0, on = 1, always = 2]"
    },
    "sql.defaults.optimizer": {
      "value": "1",
      "type": "e",
      "description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
    },
    "sql.defaults.reorder_joins_limit": {
      "value": "4",
      "type": "i",
      "description": "default number of joins to reorder"
    },
    "sql.defaults.results_buffer.size": {
      "value": "16 KiB",
      "type": "z",
      "description": "default size of the buffer that accumulates results for a statement or a batch of statements before they are sent to the client. This can be overridden on an individual connection with the 'results_buffer_size' parameter. Note that auto-retries generally only happen while no results have been delivered to the client, so reducing this size can increase the number of retriable errors a client receives. On the other hand, increasing the buffer size can increase the delay until the client receives the first result row. Updating the setting only affects new connections. Setting to 0 disables any buffering."
    },
    "sql.defaults.serial_normalization": {
      "value": "0",
      "type": "e",
      "description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
    },
    "sql.distsql.distribute_index_joins": {
      "value": "true",
      "type": "b",
      "description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
    },
    "sql.distsql.flow_stream_timeout": {
      "value": "10s",
      "type": "d",
      "description": "amount of time incoming streams wait for a flow to be set up before erroring out"
    },
    "sql.distsql.interleaved_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set we plan interleaved table joins instead of merge joins when possible"
    },
    "sql.distsql.max_running_flows": {
      "value": "500",
      "type": "i",
      "description": "maximum number of concurrent flows that can be run on a node"
    },
    "sql.distsql.merge_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, we plan merge joins when possible"
    },
    "sql.distsql.temp_storage.joins": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql joins"
    },
    "sql.distsql.temp_storage.sorts": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql sorts"
    },
    "sql.distsql.temp_storage.workmem": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
    },
    "sql.metrics.statement_details.dump_to_logs": {
      "value": "false",
      "type": "b",
      "description": "dump collected statement statistics to node logs when periodically cleared"
    },
    "sql.metrics.statement_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-statement query statistics"
    },
    "sql.metrics.statement_details.plan_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "periodically save a logical plan for each fingerprint"
    },
    "sql.metrics.statement_details.plan_collection.period": {
      "value": "5m0s",
      "type": "d",
      "description": "the time until a new logical plan is collected"
    },
    "sql.metrics.statement_details.threshold": {
      "value": "0s",
      "type": "d",
      "description": "minimum execution time to cause statistics to be collected"
    },
    "sql.parallel_scans.enabled": {
      "value": "true",
      "type": "b",
      "description": "parallelizes scanning different ranges when the maximum result size can be deduced"
    },
    "sql.query_cache.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable the query cache"
    },
    "sql.stats.automatic_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "automatic statistics collection mode"
    },
    "sql.stats.automatic_collection.fraction_stale_rows": {
      "value": "0.2",
      "type": "f",
      "description": "target fraction of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.automatic_collection.max_fraction_idle": {
      "value": "0.9",
      "type": "f",
      "description": "maximum fraction of time that automatic statistics sampler processors are idle"
    },
    "sql.stats.automatic_collection.min_stale_rows": {
      "value": "500",
      "type": "i",
      "description": "target minimum number of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.max_timestamp_age": {
      "value": "5m0s",
      "type": "d",
      "description": "maximum age of timestamp during table statistics collection"
    },
    "sql.stats.post_events.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, an event is shown for every CREATE STATISTICS job"
    },
    "sql.tablecache.lease.refresh_limit": {
      "value": "50",
      "type": "i",
      "description": "maximum number of tables to periodically refresh leases for"
    },
    "sql.trace.log_statement_execute": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable logging of executed statements"
    },
    "sql.trace.session_eventlog.enabled": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable session tracing"
    },
    "sql.trace.txn.enable_threshold": {
      "value": "0s",
      "type": "d",
      "description": "duration beyond which all transactions are traced (set to 0 to disable)"
    },
    "timeseries.storage.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
    },
    "timeseries.storage.resolution_10s.ttl": {
      "value": "240h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
    },
    "timeseries.storage.resolution_30m.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
    },
    "trace.debug.enable": {
      "value": "false",
      "type": "b",
      "description": "if set, traces for recent requests can be seen in the /debug page"
    },
    "trace.lightstep.token": {
      "type": "s",
      "description": "if set, traces go to Lightstep using this token"
    },
    "trace.zipkin.collector": {
      "type": "s",
      "description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set"
    },
    "version": {
      "value": "19.1",
      "type": "m",
      "description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'"
    }
  }
}
`
const DefaultClusterSettings19_1_6 string = `
{
  "key_values": {
    "changefeed.experimental_poll_interval": {
      "value": "1s",
      "type": "d",
      "description": "polling interval for the prototype changefeed implementation (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "changefeed.push.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, changed are pushed instead of pulled. This requires the kv.rangefeed.enabled setting. See https://www.cockroachlabs.com/docs/v19.1/change-data-capture.html#enable-rangefeeds-to-reduce-latency"
    },
    "cloudstorage.gs.default.key": {
      "type": "s",
      "description": "if set, JSON key to use during Google Cloud Storage operations"
    },
    "cloudstorage.http.custom_ca": {
      "type": "s",
      "description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
    },
    "cloudstorage.timeout": {
      "value": "10m0s",
      "type": "d",
      "description": "the timeout for import/export storage operations"
    },
    "cluster.organization": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "organization name"
    },
    "cluster.preserve_downgrade_option": {
      "type": "s",
      "description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
    },
    "compactor.enabled": {
      "value": "true",
      "type": "b",
      "description": "when false, the system will reclaim space occupied by deleted data less aggressively"
    },
    "compactor.max_record_age": {
      "value": "24h0m0s",
      "type": "d",
      "description": "discard suggestions not processed within this duration (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.min_interval": {
      "value": "15s",
      "type": "d",
      "description": "minimum time interval to wait before compacting (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_available_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the available logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_bytes": {
      "value": "256 MiB",
      "type": "z",
      "description": "minimum expected logical space reclamation required before considering an aggregated suggestion (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_used_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the used logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "debug.panic_on_failed_assertions": {
      "value": "false",
      "type": "b",
      "description": "panic when an assertion fails rather than reporting"
    },
    "diagnostics.forced_stat_reset.interval": {
      "value": "2h0m0s",
      "type": "d",
      "description": "interval after which pending diagnostics statistics should be discarded even if not reported"
    },
    "diagnostics.reporting.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable reporting diagnostic metrics to cockroach labs"
    },
    "diagnostics.reporting.interval": {
      "value": "1h0m0s",
      "type": "d",
      "description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
    },
    "diagnostics.reporting.send_crash_reports": {
      "value": "true",
      "type": "b",
      "description": "send crash and panic reports"
    },
    "external.graphite.endpoint": {
      "type": "s",
      "description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
    },
    "external.graphite.interval": {
      "value": "10s",
      "type": "d",
      "description": "the interval at which metrics are pushed to Graphite (if enabled)"
    },
    "jobs.registry.leniency": {
      "value": "1m0s",
      "type": "d",
      "description": "the amount of time to defer any attempts to reschedule a job"
    },
    "jobs.retention_time": {
      "value": "336h0m0s",
      "type": "d",
      "description": "the amount of time to retain records for completed jobs before"
    },
    "kv.allocator.lease_rebalancing_aggressiveness": {
      "value": "1",
      "type": "f",
      "description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
    },
    "kv.allocator.load_based_lease_rebalancing.enabled": {
      "value": "true",
      "type": "b",
      "description": "set to enable rebalancing of range leases based on load and latency"
    },
    "kv.allocator.load_based_rebalancing": {
      "value": "2",
      "type": "e",
      "description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
    },
    "kv.allocator.qps_rebalance_threshold": {
      "value": "0.25",
      "type": "f",
      "description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
    },
    "kv.allocator.range_rebalance_threshold": {
      "value": "0.05",
      "type": "f",
      "description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
    },
    "kv.bulk_io_write.addsstable_max_rate": {
      "value": "1.7976931348623157E+308",
      "type": "f",
      "description": "maximum number of AddSSTable requests per second for a single store"
    },
    "kv.bulk_io_write.concurrent_addsstable_requests": {
      "value": "1",
      "type": "i",
      "description": "number of AddSSTable requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_export_requests": {
      "value": "3",
      "type": "i",
      "description": "number of export requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_import_requests": {
      "value": "1",
      "type": "i",
      "description": "number of import requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.max_rate": {
      "value": "1.0 TiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
    },
    "kv.bulk_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
    },
    "kv.closed_timestamp.close_fraction": {
      "value": "0.2",
      "type": "f",
      "description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
    },
    "kv.closed_timestamp.follower_reads_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
    },
    "kv.closed_timestamp.target_duration": {
      "value": "30s",
      "type": "d",
      "description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
    },
    "kv.follower_read.target_multiple": {
      "value": "3",
      "type": "f",
      "description": "if above 1, encourages the distsender to perform a read against the closest replica if a request is older than kv.closed_timestamp.target_duration * (1 + kv.closed_timestamp.close_fraction * this) less a clock uncertainty interval. This value also is used to create follower_timestamp(). (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.import.batch_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the maximum size of the payload in an AddSSTable request (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.raft.command.max_size": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum size of a raft command"
    },
    "kv.raft_log.disable_synchronization_unsafe": {
      "value": "false",
      "type": "b",
      "description": "set to true to disable synchronization on Raft log writes to persistent storage. Setting to true risks data loss or data corruption on server crashes. The setting is meant for internal testing only and SHOULD NOT be used in production."
    },
    "kv.range.backpressure_range_size_multiplier": {
      "value": "2",
      "type": "f",
      "description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
    },
    "kv.range_descriptor_cache.size": {
      "value": "1000000",
      "type": "i",
      "description": "maximum number of entries in the range descriptor and leaseholder caches"
    },
    "kv.range_merge.queue_enabled": {
      "value": "true",
      "type": "b",
      "description": "whether the automatic merge queue is enabled"
    },
    "kv.range_merge.queue_interval": {
      "value": "1s",
      "type": "d",
      "description": "how long the merge queue waits between processing replicas (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.range_split.by_load_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow automatic splits of ranges based on where load is concentrated"
    },
    "kv.range_split.load_qps_threshold": {
      "value": "250",
      "type": "i",
      "description": "the QPS over which, the range becomes a candidate for load based splitting"
    },
    "kv.rangefeed.concurrent_catchup_iterators": {
      "value": "64",
      "type": "i",
      "description": "number of rangefeeds catchup iterators a store will allow concurrently before queueing"
    },
    "kv.rangefeed.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, rangefeed registration is enabled"
    },
    "kv.snapshot_rebalance.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for rebalance and upreplication snapshots"
    },
    "kv.snapshot_recovery.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for recovery snapshots"
    },
    "kv.transaction.max_intents_bytes": {
      "value": "262144",
      "type": "i",
      "description": "maximum number of bytes used to track write intents in transactions"
    },
    "kv.transaction.max_refresh_spans_bytes": {
      "value": "256000",
      "type": "i",
      "description": "maximum number of bytes used to track refresh spans in serializable transactions"
    },
    "kv.transaction.write_pipelining_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional writes are pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_batch_size": {
      "value": "128",
      "type": "i",
      "description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_outstanding_size": {
      "value": "256 KiB",
      "type": "z",
      "description": "maximum number of bytes used to track in-flight pipelined writes before disabling pipelining"
    },
    "rocksdb.ingest_backpressure.delay_l0_file": {
      "value": "200ms",
      "type": "d",
      "description": "delay to add to SST ingestions per file in L0 over the configured limit"
    },
    "rocksdb.ingest_backpressure.l0_file_count_threshold": {
      "value": "20",
      "type": "i",
      "description": "number of L0 files after which to backpressure SST ingestions"
    },
    "rocksdb.ingest_backpressure.max_delay": {
      "value": "5s",
      "type": "d",
      "description": "maximum amount of time to backpressure a single SST ingestion"
    },
    "rocksdb.ingest_backpressure.pending_compaction_threshold": {
      "value": "64 GiB",
      "type": "z",
      "description": "pending compaction estimate above which to backpressure SST ingestions"
    },
    "rocksdb.min_wal_sync_interval": {
      "value": "0s",
      "type": "d",
      "description": "minimum duration between syncs of the RocksDB WAL"
    },
    "schemachanger.backfiller.buffer_size": {
      "value": "196 MiB",
      "type": "z",
      "description": "amount to buffer in memory during backfills"
    },
    "schemachanger.backfiller.max_sst_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "target size for ingested files during backfills"
    },
    "schemachanger.bulk_index_backfill.batch_size": {
      "value": "5000",
      "type": "i",
      "description": "number of rows to process at a time during bulk index backfill"
    },
    "schemachanger.bulk_index_backfill.enabled": {
      "value": "true",
      "type": "b",
      "description": "backfill indexes in bulk via addsstable"
    },
    "schemachanger.lease.duration": {
      "value": "5m0s",
      "type": "d",
      "description": "the duration of a schema change lease"
    },
    "schemachanger.lease.renew_fraction": {
      "value": "0.5",
      "type": "f",
      "description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
    },
    "server.clock.forward_jump_check_enabled": {
      "value": "false",
      "type": "b",
      "description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic"
    },
    "server.clock.persist_upper_bound_interval": {
      "value": "0s",
      "type": "d",
      "description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
    },
    "server.consistency_check.interval": {
      "value": "24h0m0s",
      "type": "d",
      "description": "the time between range consistency checks; set to 0 to disable consistency checking"
    },
    "server.declined_reservation_timeout": {
      "value": "1s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
    },
    "server.eventlog.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "if nonzero, event log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.failed_reservation_timeout": {
      "value": "5s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
    },
    "server.goroutine_dump.num_goroutines_threshold": {
      "value": "1000",
      "type": "i",
      "description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
    },
    "server.goroutine_dump.total_dump_size_limit": {
      "value": "500 MiB",
      "type": "z",
      "description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
    },
    "server.heap_profile.max_profiles": {
      "value": "5",
      "type": "i",
      "description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept."
    },
    "server.heap_profile.system_memory_threshold_fraction": {
      "value": "0.85",
      "type": "f",
      "description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
    },
    "server.host_based_authentication.configuration": {
      "type": "s",
      "description": "host-based authentication configuration to use during connection authentication"
    },
    "server.rangelog.ttl": {
      "value": "720h0m0s",
      "type": "d",
      "description": "if nonzero, range log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.remote_debugging.mode": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
    },
    "server.shutdown.drain_wait": {
      "value": "0s",
      "type": "d",
      "description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
    },
    "server.shutdown.query_wait": {
      "value": "10s",
      "type": "d",
      "description": "the server will wait for at least this amount of time for active queries to finish"
    },
    "server.time_until_store_dead": {
      "value": "5m0s",
      "type": "d",
      "description": "the time after which if there is no new gossiped information about a store, it is considered dead"
    },
    "server.web_session_timeout": {
      "value": "168h0m0s",
      "type": "d",
      "description": "the duration that a newly created web session will be valid"
    },
    "sql.defaults.default_int_size": {
      "value": "8",
      "type": "i",
      "description": "the size, in bytes, of an INT type"
    },
    "sql.defaults.distsql": {
      "value": "1",
      "type": "e",
      "description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
    },
    "sql.defaults.experimental_vectorize": {
      "value": "0",
      "type": "e",
      "description": "default experimental_vectorize mode [off = 0, on = 1, always = 2]"
    },
    "sql.defaults.optimizer": {
      "value": "1",
      "type": "e",
      "description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
    },
    "sql.defaults.reorder_joins_limit": {
      "value": "4",
      "type": "i",
      "description": "default number of joins to reorder"
    },
    "sql.defaults.results_buffer.size": {
      "value": "16 KiB",
      "type": "z",
      "description": "default size of the buffer that accumulates results for a statement or a batch of statements before they are sent to the client. This can be overridden on an individual connection with the 'results_buffer_size' parameter. Note that auto-retries generally only happen while no results have been delivered to the client, so reducing this size can increase the number of retriable errors a client receives. On the other hand, increasing the buffer size can increase the delay until the client receives the first result row. Updating the setting only affects new connections. Setting to 0 disables any buffering."
    },
    "sql.defaults.serial_normalization": {
      "value": "0",
      "type": "e",
      "description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
    },
    "sql.distsql.distribute_index_joins": {
      "value": "true",
      "type": "b",
      "description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
    },
    "sql.distsql.flow_stream_timeout": {
      "value": "10s",
      "type": "d",
      "description": "amount of time incoming streams wait for a flow to be set up before erroring out"
    },
    "sql.distsql.interleaved_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set we plan interleaved table joins instead of merge joins when possible"
    },
    "sql.distsql.max_running_flows": {
      "value": "500",
      "type": "i",
      "description": "maximum number of concurrent flows that can be run on a node"
    },
    "sql.distsql.merge_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, we plan merge joins when possible"
    },
    "sql.distsql.temp_storage.joins": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql joins"
    },
    "sql.distsql.temp_storage.sorts": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql sorts"
    },
    "sql.distsql.temp_storage.workmem": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
    },
    "sql.metrics.statement_details.dump_to_logs": {
      "value": "false",
      "type": "b",
      "description": "dump collected statement statistics to node logs when periodically cleared"
    },
    "sql.metrics.statement_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-statement query statistics"
    },
    "sql.metrics.statement_details.plan_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "periodically save a logical plan for each fingerprint"
    },
    "sql.metrics.statement_details.plan_collection.period": {
      "value": "5m0s",
      "type": "d",
      "description": "the time until a new logical plan is collected"
    },
    "sql.metrics.statement_details.threshold": {
      "value": "0s",
      "type": "d",
      "description": "minimum execution time to cause statistics to be collected"
    },
    "sql.parallel_scans.enabled": {
      "value": "true",
      "type": "b",
      "description": "parallelizes scanning different ranges when the maximum result size can be deduced"
    },
    "sql.query_cache.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable the query cache"
    },
    "sql.stats.automatic_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "automatic statistics collection mode"
    },
    "sql.stats.automatic_collection.fraction_stale_rows": {
      "value": "0.2",
      "type": "f",
      "description": "target fraction of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.automatic_collection.max_fraction_idle": {
      "value": "0.9",
      "type": "f",
      "description": "maximum fraction of time that automatic statistics sampler processors are idle"
    },
    "sql.stats.automatic_collection.min_stale_rows": {
      "value": "500",
      "type": "i",
      "description": "target minimum number of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.max_timestamp_age": {
      "value": "5m0s",
      "type": "d",
      "description": "maximum age of timestamp during table statistics collection"
    },
    "sql.stats.post_events.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, an event is shown for every CREATE STATISTICS job"
    },
    "sql.tablecache.lease.refresh_limit": {
      "value": "50",
      "type": "i",
      "description": "maximum number of tables to periodically refresh leases for"
    },
    "sql.trace.log_statement_execute": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable logging of executed statements"
    },
    "sql.trace.session_eventlog.enabled": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable session tracing"
    },
    "sql.trace.txn.enable_threshold": {
      "value": "0s",
      "type": "d",
      "description": "duration beyond which all transactions are traced (set to 0 to disable)"
    },
    "timeseries.storage.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
    },
    "timeseries.storage.resolution_10s.ttl": {
      "value": "240h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
    },
    "timeseries.storage.resolution_30m.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
    },
    "trace.debug.enable": {
      "value": "false",
      "type": "b",
      "description": "if set, traces for recent requests can be seen in the /debug page"
    },
    "trace.lightstep.token": {
      "type": "s",
      "description": "if set, traces go to Lightstep using this token"
    },
    "trace.zipkin.collector": {
      "type": "s",
      "description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set"
    },
    "version": {
      "value": "19.1",
      "type": "m",
      "description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'"
    }
  }
}
`
const DefaultClusterSettings19_1_5 string = `
{
  "key_values": {
    "changefeed.experimental_poll_interval": {
      "value": "1s",
      "type": "d",
      "description": "polling interval for the prototype changefeed implementation (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "changefeed.push.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, changed are pushed instead of pulled. This requires the kv.rangefeed.enabled setting. See https://www.cockroachlabs.com/docs/v19.1/change-data-capture.html#enable-rangefeeds-to-reduce-latency"
    },
    "cloudstorage.gs.default.key": {
      "type": "s",
      "description": "if set, JSON key to use during Google Cloud Storage operations"
    },
    "cloudstorage.http.custom_ca": {
      "type": "s",
      "description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
    },
    "cloudstorage.timeout": {
      "value": "10m0s",
      "type": "d",
      "description": "the timeout for import/export storage operations"
    },
    "cluster.organization": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "organization name"
    },
    "cluster.preserve_downgrade_option": {
      "type": "s",
      "description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
    },
    "compactor.enabled": {
      "value": "true",
      "type": "b",
      "description": "when false, the system will reclaim space occupied by deleted data less aggressively"
    },
    "compactor.max_record_age": {
      "value": "24h0m0s",
      "type": "d",
      "description": "discard suggestions not processed within this duration (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.min_interval": {
      "value": "15s",
      "type": "d",
      "description": "minimum time interval to wait before compacting (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_available_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the available logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_bytes": {
      "value": "256 MiB",
      "type": "z",
      "description": "minimum expected logical space reclamation required before considering an aggregated suggestion (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_used_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the used logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "debug.panic_on_failed_assertions": {
      "value": "false",
      "type": "b",
      "description": "panic when an assertion fails rather than reporting"
    },
    "diagnostics.forced_stat_reset.interval": {
      "value": "2h0m0s",
      "type": "d",
      "description": "interval after which pending diagnostics statistics should be discarded even if not reported"
    },
    "diagnostics.reporting.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable reporting diagnostic metrics to cockroach labs"
    },
    "diagnostics.reporting.interval": {
      "value": "1h0m0s",
      "type": "d",
      "description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
    },
    "diagnostics.reporting.send_crash_reports": {
      "value": "true",
      "type": "b",
      "description": "send crash and panic reports"
    },
    "external.graphite.endpoint": {
      "type": "s",
      "description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
    },
    "external.graphite.interval": {
      "value": "10s",
      "type": "d",
      "description": "the interval at which metrics are pushed to Graphite (if enabled)"
    },
    "jobs.registry.leniency": {
      "value": "1m0s",
      "type": "d",
      "description": "the amount of time to defer any attempts to reschedule a job"
    },
    "jobs.retention_time": {
      "value": "336h0m0s",
      "type": "d",
      "description": "the amount of time to retain records for completed jobs before"
    },
    "kv.allocator.lease_rebalancing_aggressiveness": {
      "value": "1",
      "type": "f",
      "description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
    },
    "kv.allocator.load_based_lease_rebalancing.enabled": {
      "value": "true",
      "type": "b",
      "description": "set to enable rebalancing of range leases based on load and latency"
    },
    "kv.allocator.load_based_rebalancing": {
      "value": "2",
      "type": "e",
      "description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
    },
    "kv.allocator.qps_rebalance_threshold": {
      "value": "0.25",
      "type": "f",
      "description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
    },
    "kv.allocator.range_rebalance_threshold": {
      "value": "0.05",
      "type": "f",
      "description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
    },
    "kv.bulk_io_write.addsstable_max_rate": {
      "value": "1.7976931348623157E+308",
      "type": "f",
      "description": "maximum number of AddSSTable requests per second for a single store"
    },
    "kv.bulk_io_write.concurrent_addsstable_requests": {
      "value": "1",
      "type": "i",
      "description": "number of AddSSTable requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_export_requests": {
      "value": "3",
      "type": "i",
      "description": "number of export requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_import_requests": {
      "value": "1",
      "type": "i",
      "description": "number of import requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.max_rate": {
      "value": "1.0 TiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
    },
    "kv.bulk_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
    },
    "kv.closed_timestamp.close_fraction": {
      "value": "0.2",
      "type": "f",
      "description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
    },
    "kv.closed_timestamp.follower_reads_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
    },
    "kv.closed_timestamp.target_duration": {
      "value": "30s",
      "type": "d",
      "description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
    },
    "kv.follower_read.target_multiple": {
      "value": "3",
      "type": "f",
      "description": "if above 1, encourages the distsender to perform a read against the closest replica if a request is older than kv.closed_timestamp.target_duration * (1 + kv.closed_timestamp.close_fraction * this) less a clock uncertainty interval. This value also is used to create follower_timestamp(). (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.import.batch_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the maximum size of the payload in an AddSSTable request (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.raft.command.max_size": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum size of a raft command"
    },
    "kv.raft_log.disable_synchronization_unsafe": {
      "value": "false",
      "type": "b",
      "description": "set to true to disable synchronization on Raft log writes to persistent storage. Setting to true risks data loss or data corruption on server crashes. The setting is meant for internal testing only and SHOULD NOT be used in production."
    },
    "kv.range.backpressure_range_size_multiplier": {
      "value": "2",
      "type": "f",
      "description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
    },
    "kv.range_descriptor_cache.size": {
      "value": "1000000",
      "type": "i",
      "description": "maximum number of entries in the range descriptor and leaseholder caches"
    },
    "kv.range_merge.queue_enabled": {
      "value": "true",
      "type": "b",
      "description": "whether the automatic merge queue is enabled"
    },
    "kv.range_merge.queue_interval": {
      "value": "1s",
      "type": "d",
      "description": "how long the merge queue waits between processing replicas (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.range_split.by_load_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow automatic splits of ranges based on where load is concentrated"
    },
    "kv.range_split.load_qps_threshold": {
      "value": "250",
      "type": "i",
      "description": "the QPS over which, the range becomes a candidate for load based splitting"
    },
    "kv.rangefeed.concurrent_catchup_iterators": {
      "value": "64",
      "type": "i",
      "description": "number of rangefeeds catchup iterators a store will allow concurrently before queueing"
    },
    "kv.rangefeed.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, rangefeed registration is enabled"
    },
    "kv.snapshot_rebalance.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for rebalance and upreplication snapshots"
    },
    "kv.snapshot_recovery.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for recovery snapshots"
    },
    "kv.transaction.max_intents_bytes": {
      "value": "262144",
      "type": "i",
      "description": "maximum number of bytes used to track write intents in transactions"
    },
    "kv.transaction.max_refresh_spans_bytes": {
      "value": "256000",
      "type": "i",
      "description": "maximum number of bytes used to track refresh spans in serializable transactions"
    },
    "kv.transaction.write_pipelining_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional writes are pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_batch_size": {
      "value": "128",
      "type": "i",
      "description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_outstanding_size": {
      "value": "256 KiB",
      "type": "z",
      "description": "maximum number of bytes used to track in-flight pipelined writes before disabling pipelining"
    },
    "rocksdb.ingest_backpressure.delay_l0_file": {
      "value": "200ms",
      "type": "d",
      "description": "delay to add to SST ingestions per file in L0 over the configured limit"
    },
    "rocksdb.ingest_backpressure.l0_file_count_threshold": {
      "value": "20",
      "type": "i",
      "description": "number of L0 files after which to backpressure SST ingestions"
    },
    "rocksdb.ingest_backpressure.max_delay": {
      "value": "5s",
      "type": "d",
      "description": "maximum amount of time to backpressure a single SST ingestion"
    },
    "rocksdb.ingest_backpressure.pending_compaction_threshold": {
      "value": "64 GiB",
      "type": "z",
      "description": "pending compaction estimate above which to backpressure SST ingestions"
    },
    "rocksdb.min_wal_sync_interval": {
      "value": "0s",
      "type": "d",
      "description": "minimum duration between syncs of the RocksDB WAL"
    },
    "schemachanger.backfiller.buffer_size": {
      "value": "196 MiB",
      "type": "z",
      "description": "amount to buffer in memory during backfills"
    },
    "schemachanger.backfiller.max_sst_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "target size for ingested files during backfills"
    },
    "schemachanger.bulk_index_backfill.batch_size": {
      "value": "5000",
      "type": "i",
      "description": "number of rows to process at a time during bulk index backfill"
    },
    "schemachanger.bulk_index_backfill.enabled": {
      "value": "true",
      "type": "b",
      "description": "backfill indexes in bulk via addsstable"
    },
    "schemachanger.lease.duration": {
      "value": "5m0s",
      "type": "d",
      "description": "the duration of a schema change lease"
    },
    "schemachanger.lease.renew_fraction": {
      "value": "0.5",
      "type": "f",
      "description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
    },
    "server.clock.forward_jump_check_enabled": {
      "value": "false",
      "type": "b",
      "description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic"
    },
    "server.clock.persist_upper_bound_interval": {
      "value": "0s",
      "type": "d",
      "description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
    },
    "server.consistency_check.interval": {
      "value": "24h0m0s",
      "type": "d",
      "description": "the time between range consistency checks; set to 0 to disable consistency checking"
    },
    "server.declined_reservation_timeout": {
      "value": "1s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
    },
    "server.eventlog.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "if nonzero, event log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.failed_reservation_timeout": {
      "value": "5s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
    },
    "server.goroutine_dump.num_goroutines_threshold": {
      "value": "1000",
      "type": "i",
      "description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
    },
    "server.goroutine_dump.total_dump_size_limit": {
      "value": "500 MiB",
      "type": "z",
      "description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
    },
    "server.heap_profile.max_profiles": {
      "value": "5",
      "type": "i",
      "description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept."
    },
    "server.heap_profile.system_memory_threshold_fraction": {
      "value": "0.85",
      "type": "f",
      "description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
    },
    "server.host_based_authentication.configuration": {
      "type": "s",
      "description": "host-based authentication configuration to use during connection authentication"
    },
    "server.rangelog.ttl": {
      "value": "720h0m0s",
      "type": "d",
      "description": "if nonzero, range log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.remote_debugging.mode": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
    },
    "server.shutdown.drain_wait": {
      "value": "0s",
      "type": "d",
      "description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
    },
    "server.shutdown.query_wait": {
      "value": "10s",
      "type": "d",
      "description": "the server will wait for at least this amount of time for active queries to finish"
    },
    "server.time_until_store_dead": {
      "value": "5m0s",
      "type": "d",
      "description": "the time after which if there is no new gossiped information about a store, it is considered dead"
    },
    "server.web_session_timeout": {
      "value": "168h0m0s",
      "type": "d",
      "description": "the duration that a newly created web session will be valid"
    },
    "sql.defaults.default_int_size": {
      "value": "8",
      "type": "i",
      "description": "the size, in bytes, of an INT type"
    },
    "sql.defaults.distsql": {
      "value": "1",
      "type": "e",
      "description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
    },
    "sql.defaults.experimental_vectorize": {
      "value": "0",
      "type": "e",
      "description": "default experimental_vectorize mode [off = 0, on = 1, always = 2]"
    },
    "sql.defaults.optimizer": {
      "value": "1",
      "type": "e",
      "description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
    },
    "sql.defaults.reorder_joins_limit": {
      "value": "4",
      "type": "i",
      "description": "default number of joins to reorder"
    },
    "sql.defaults.results_buffer.size": {
      "value": "16 KiB",
      "type": "z",
      "description": "default size of the buffer that accumulates results for a statement or a batch of statements before they are sent to the client. This can be overridden on an individual connection with the 'results_buffer_size' parameter. Note that auto-retries generally only happen while no results have been delivered to the client, so reducing this size can increase the number of retriable errors a client receives. On the other hand, increasing the buffer size can increase the delay until the client receives the first result row. Updating the setting only affects new connections. Setting to 0 disables any buffering."
    },
    "sql.defaults.serial_normalization": {
      "value": "0",
      "type": "e",
      "description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
    },
    "sql.distsql.distribute_index_joins": {
      "value": "true",
      "type": "b",
      "description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
    },
    "sql.distsql.flow_stream_timeout": {
      "value": "10s",
      "type": "d",
      "description": "amount of time incoming streams wait for a flow to be set up before erroring out"
    },
    "sql.distsql.interleaved_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set we plan interleaved table joins instead of merge joins when possible"
    },
    "sql.distsql.max_running_flows": {
      "value": "500",
      "type": "i",
      "description": "maximum number of concurrent flows that can be run on a node"
    },
    "sql.distsql.merge_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, we plan merge joins when possible"
    },
    "sql.distsql.temp_storage.joins": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql joins"
    },
    "sql.distsql.temp_storage.sorts": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql sorts"
    },
    "sql.distsql.temp_storage.workmem": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
    },
    "sql.metrics.statement_details.dump_to_logs": {
      "value": "false",
      "type": "b",
      "description": "dump collected statement statistics to node logs when periodically cleared"
    },
    "sql.metrics.statement_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-statement query statistics"
    },
    "sql.metrics.statement_details.plan_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "periodically save a logical plan for each fingerprint"
    },
    "sql.metrics.statement_details.plan_collection.period": {
      "value": "5m0s",
      "type": "d",
      "description": "the time until a new logical plan is collected"
    },
    "sql.metrics.statement_details.threshold": {
      "value": "0s",
      "type": "d",
      "description": "minimum execution time to cause statistics to be collected"
    },
    "sql.parallel_scans.enabled": {
      "value": "true",
      "type": "b",
      "description": "parallelizes scanning different ranges when the maximum result size can be deduced"
    },
    "sql.query_cache.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable the query cache"
    },
    "sql.stats.automatic_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "automatic statistics collection mode"
    },
    "sql.stats.automatic_collection.fraction_stale_rows": {
      "value": "0.2",
      "type": "f",
      "description": "target fraction of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.automatic_collection.max_fraction_idle": {
      "value": "0.9",
      "type": "f",
      "description": "maximum fraction of time that automatic statistics sampler processors are idle"
    },
    "sql.stats.automatic_collection.min_stale_rows": {
      "value": "500",
      "type": "i",
      "description": "target minimum number of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.max_timestamp_age": {
      "value": "5m0s",
      "type": "d",
      "description": "maximum age of timestamp during table statistics collection"
    },
    "sql.stats.post_events.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, an event is shown for every CREATE STATISTICS job"
    },
    "sql.tablecache.lease.refresh_limit": {
      "value": "50",
      "type": "i",
      "description": "maximum number of tables to periodically refresh leases for"
    },
    "sql.trace.log_statement_execute": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable logging of executed statements"
    },
    "sql.trace.session_eventlog.enabled": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable session tracing"
    },
    "sql.trace.txn.enable_threshold": {
      "value": "0s",
      "type": "d",
      "description": "duration beyond which all transactions are traced (set to 0 to disable)"
    },
    "timeseries.storage.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
    },
    "timeseries.storage.resolution_10s.ttl": {
      "value": "240h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
    },
    "timeseries.storage.resolution_30m.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
    },
    "trace.debug.enable": {
      "value": "false",
      "type": "b",
      "description": "if set, traces for recent requests can be seen in the /debug page"
    },
    "trace.lightstep.token": {
      "type": "s",
      "description": "if set, traces go to Lightstep using this token"
    },
    "trace.zipkin.collector": {
      "type": "s",
      "description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set"
    },
    "version": {
      "value": "19.1",
      "type": "m",
      "description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'"
    }
  }
}
`
const DefaultClusterSettings19_1_4 string = `
{
  "key_values": {
    "changefeed.experimental_poll_interval": {
      "value": "1s",
      "type": "d",
      "description": "polling interval for the prototype changefeed implementation (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "changefeed.push.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, changed are pushed instead of pulled. This requires the kv.rangefeed.enabled setting. See https://www.cockroachlabs.com/docs/v19.1/change-data-capture.html#enable-rangefeeds-to-reduce-latency"
    },
    "cloudstorage.gs.default.key": {
      "type": "s",
      "description": "if set, JSON key to use during Google Cloud Storage operations"
    },
    "cloudstorage.http.custom_ca": {
      "type": "s",
      "description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
    },
    "cloudstorage.timeout": {
      "value": "10m0s",
      "type": "d",
      "description": "the timeout for import/export storage operations"
    },
    "cluster.organization": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "organization name"
    },
    "cluster.preserve_downgrade_option": {
      "type": "s",
      "description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
    },
    "compactor.enabled": {
      "value": "true",
      "type": "b",
      "description": "when false, the system will reclaim space occupied by deleted data less aggressively"
    },
    "compactor.max_record_age": {
      "value": "24h0m0s",
      "type": "d",
      "description": "discard suggestions not processed within this duration (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.min_interval": {
      "value": "15s",
      "type": "d",
      "description": "minimum time interval to wait before compacting (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_available_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the available logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_bytes": {
      "value": "256 MiB",
      "type": "z",
      "description": "minimum expected logical space reclamation required before considering an aggregated suggestion (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "compactor.threshold_used_fraction": {
      "value": "0.1",
      "type": "f",
      "description": "consider suggestions for at least the given percentage of the used logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "debug.panic_on_failed_assertions": {
      "value": "false",
      "type": "b",
      "description": "panic when an assertion fails rather than reporting"
    },
    "diagnostics.forced_stat_reset.interval": {
      "value": "2h0m0s",
      "type": "d",
      "description": "interval after which pending diagnostics statistics should be discarded even if not reported"
    },
    "diagnostics.reporting.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable reporting diagnostic metrics to cockroach labs"
    },
    "diagnostics.reporting.interval": {
      "value": "1h0m0s",
      "type": "d",
      "description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
    },
    "diagnostics.reporting.send_crash_reports": {
      "value": "true",
      "type": "b",
      "description": "send crash and panic reports"
    },
    "external.graphite.endpoint": {
      "type": "s",
      "description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
    },
    "external.graphite.interval": {
      "value": "10s",
      "type": "d",
      "description": "the interval at which metrics are pushed to Graphite (if enabled)"
    },
    "jobs.registry.leniency": {
      "value": "1m0s",
      "type": "d",
      "description": "the amount of time to defer any attempts to reschedule a job"
    },
    "jobs.retention_time": {
      "value": "336h0m0s",
      "type": "d",
      "description": "the amount of time to retain records for completed jobs before"
    },
    "kv.allocator.lease_rebalancing_aggressiveness": {
      "value": "1",
      "type": "f",
      "description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
    },
    "kv.allocator.load_based_lease_rebalancing.enabled": {
      "value": "true",
      "type": "b",
      "description": "set to enable rebalancing of range leases based on load and latency"
    },
    "kv.allocator.load_based_rebalancing": {
      "value": "2",
      "type": "e",
      "description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
    },
    "kv.allocator.qps_rebalance_threshold": {
      "value": "0.25",
      "type": "f",
      "description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
    },
    "kv.allocator.range_rebalance_threshold": {
      "value": "0.05",
      "type": "f",
      "description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
    },
    "kv.bulk_io_write.addsstable_max_rate": {
      "value": "1.7976931348623157E+308",
      "type": "f",
      "description": "maximum number of AddSSTable requests per second for a single store"
    },
    "kv.bulk_io_write.concurrent_addsstable_requests": {
      "value": "1",
      "type": "i",
      "description": "number of AddSSTable requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_export_requests": {
      "value": "3",
      "type": "i",
      "description": "number of export requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.concurrent_import_requests": {
      "value": "1",
      "type": "i",
      "description": "number of import requests a store will handle concurrently before queuing"
    },
    "kv.bulk_io_write.max_rate": {
      "value": "1.0 TiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
    },
    "kv.bulk_sst.sync_size": {
      "value": "2.0 MiB",
      "type": "z",
      "description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
    },
    "kv.closed_timestamp.close_fraction": {
      "value": "0.2",
      "type": "f",
      "description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
    },
    "kv.closed_timestamp.follower_reads_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
    },
    "kv.closed_timestamp.target_duration": {
      "value": "30s",
      "type": "d",
      "description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
    },
    "kv.follower_read.target_multiple": {
      "value": "3",
      "type": "f",
      "description": "if above 1, encourages the distsender to perform a read against the closest replica if a request is older than kv.closed_timestamp.target_duration * (1 + kv.closed_timestamp.close_fraction * this) less a clock uncertainty interval. This value also is used to create follower_timestamp(). (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.import.batch_size": {
      "value": "32 MiB",
      "type": "z",
      "description": "the maximum size of the payload in an AddSSTable request (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.raft.command.max_size": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum size of a raft command"
    },
    "kv.raft_log.disable_synchronization_unsafe": {
      "value": "false",
      "type": "b",
      "description": "set to true to disable synchronization on Raft log writes to persistent storage. Setting to true risks data loss or data corruption on server crashes. The setting is meant for internal testing only and SHOULD NOT be used in production."
    },
    "kv.range.backpressure_range_size_multiplier": {
      "value": "2",
      "type": "f",
      "description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
    },
    "kv.range_descriptor_cache.size": {
      "value": "1000000",
      "type": "i",
      "description": "maximum number of entries in the range descriptor and leaseholder caches"
    },
    "kv.range_merge.queue_enabled": {
      "value": "true",
      "type": "b",
      "description": "whether the automatic merge queue is enabled"
    },
    "kv.range_merge.queue_interval": {
      "value": "1s",
      "type": "d",
      "description": "how long the merge queue waits between processing replicas (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
    },
    "kv.range_split.by_load_enabled": {
      "value": "true",
      "type": "b",
      "description": "allow automatic splits of ranges based on where load is concentrated"
    },
    "kv.range_split.load_qps_threshold": {
      "value": "250",
      "type": "i",
      "description": "the QPS over which, the range becomes a candidate for load based splitting"
    },
    "kv.rangefeed.concurrent_catchup_iterators": {
      "value": "64",
      "type": "i",
      "description": "number of rangefeeds catchup iterators a store will allow concurrently before queueing"
    },
    "kv.rangefeed.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, rangefeed registration is enabled"
    },
    "kv.snapshot_rebalance.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for rebalance and upreplication snapshots"
    },
    "kv.snapshot_recovery.max_rate": {
      "value": "8.0 MiB",
      "type": "z",
      "description": "the rate limit (bytes/sec) to use for recovery snapshots"
    },
    "kv.transaction.max_intents_bytes": {
      "value": "262144",
      "type": "i",
      "description": "maximum number of bytes used to track write intents in transactions"
    },
    "kv.transaction.max_refresh_spans_bytes": {
      "value": "256000",
      "type": "i",
      "description": "maximum number of bytes used to track refresh spans in serializable transactions"
    },
    "kv.transaction.write_pipelining_enabled": {
      "value": "true",
      "type": "b",
      "description": "if enabled, transactional writes are pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_batch_size": {
      "value": "128",
      "type": "i",
      "description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
    },
    "kv.transaction.write_pipelining_max_outstanding_size": {
      "value": "256 KiB",
      "type": "z",
      "description": "maximum number of bytes used to track in-flight pipelined writes before disabling pipelining"
    },
    "rocksdb.ingest_backpressure.delay_l0_file": {
      "value": "200ms",
      "type": "d",
      "description": "delay to add to SST ingestions per file in L0 over the configured limit"
    },
    "rocksdb.ingest_backpressure.l0_file_count_threshold": {
      "value": "20",
      "type": "i",
      "description": "number of L0 files after which to backpressure SST ingestions"
    },
    "rocksdb.ingest_backpressure.max_delay": {
      "value": "5s",
      "type": "d",
      "description": "maximum amount of time to backpressure a single SST ingestion"
    },
    "rocksdb.ingest_backpressure.pending_compaction_threshold": {
      "value": "64 GiB",
      "type": "z",
      "description": "pending compaction estimate above which to backpressure SST ingestions"
    },
    "rocksdb.min_wal_sync_interval": {
      "value": "0s",
      "type": "d",
      "description": "minimum duration between syncs of the RocksDB WAL"
    },
    "schemachanger.backfiller.buffer_size": {
      "value": "196 MiB",
      "type": "z",
      "description": "amount to buffer in memory during backfills"
    },
    "schemachanger.backfiller.max_sst_size": {
      "value": "16 MiB",
      "type": "z",
      "description": "target size for ingested files during backfills"
    },
    "schemachanger.bulk_index_backfill.batch_size": {
      "value": "5000",
      "type": "i",
      "description": "number of rows to process at a time during bulk index backfill"
    },
    "schemachanger.bulk_index_backfill.enabled": {
      "value": "true",
      "type": "b",
      "description": "backfill indexes in bulk via addsstable"
    },
    "schemachanger.lease.duration": {
      "value": "5m0s",
      "type": "d",
      "description": "the duration of a schema change lease"
    },
    "schemachanger.lease.renew_fraction": {
      "value": "0.5",
      "type": "f",
      "description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
    },
    "server.clock.forward_jump_check_enabled": {
      "value": "false",
      "type": "b",
      "description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic"
    },
    "server.clock.persist_upper_bound_interval": {
      "value": "0s",
      "type": "d",
      "description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
    },
    "server.consistency_check.interval": {
      "value": "24h0m0s",
      "type": "d",
      "description": "the time between range consistency checks; set to 0 to disable consistency checking"
    },
    "server.declined_reservation_timeout": {
      "value": "1s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
    },
    "server.eventlog.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "if nonzero, event log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.failed_reservation_timeout": {
      "value": "5s",
      "type": "d",
      "description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
    },
    "server.goroutine_dump.num_goroutines_threshold": {
      "value": "1000",
      "type": "i",
      "description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
    },
    "server.goroutine_dump.total_dump_size_limit": {
      "value": "500 MiB",
      "type": "z",
      "description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
    },
    "server.heap_profile.max_profiles": {
      "value": "5",
      "type": "i",
      "description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept."
    },
    "server.heap_profile.system_memory_threshold_fraction": {
      "value": "0.85",
      "type": "f",
      "description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
    },
    "server.host_based_authentication.configuration": {
      "type": "s",
      "description": "host-based authentication configuration to use during connection authentication"
    },
    "server.rangelog.ttl": {
      "value": "720h0m0s",
      "type": "d",
      "description": "if nonzero, range log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
    },
    "server.remote_debugging.mode": {
      "value": "\u003credacted\u003e",
      "type": "s",
      "description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
    },
    "server.shutdown.drain_wait": {
      "value": "0s",
      "type": "d",
      "description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
    },
    "server.shutdown.query_wait": {
      "value": "10s",
      "type": "d",
      "description": "the server will wait for at least this amount of time for active queries to finish"
    },
    "server.time_until_store_dead": {
      "value": "5m0s",
      "type": "d",
      "description": "the time after which if there is no new gossiped information about a store, it is considered dead"
    },
    "server.web_session_timeout": {
      "value": "168h0m0s",
      "type": "d",
      "description": "the duration that a newly created web session will be valid"
    },
    "sql.defaults.default_int_size": {
      "value": "8",
      "type": "i",
      "description": "the size, in bytes, of an INT type"
    },
    "sql.defaults.distsql": {
      "value": "1",
      "type": "e",
      "description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
    },
    "sql.defaults.experimental_vectorize": {
      "value": "0",
      "type": "e",
      "description": "default experimental_vectorize mode [off = 0, on = 1, always = 2]"
    },
    "sql.defaults.optimizer": {
      "value": "1",
      "type": "e",
      "description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
    },
    "sql.defaults.reorder_joins_limit": {
      "value": "4",
      "type": "i",
      "description": "default number of joins to reorder"
    },
    "sql.defaults.results_buffer.size": {
      "value": "16 KiB",
      "type": "z",
      "description": "default size of the buffer that accumulates results for a statement or a batch of statements before they are sent to the client. This can be overridden on an individual connection with the 'results_buffer_size' parameter. Note that auto-retries generally only happen while no results have been delivered to the client, so reducing this size can increase the number of retriable errors a client receives. On the other hand, increasing the buffer size can increase the delay until the client receives the first result row. Updating the setting only affects new connections. Setting to 0 disables any buffering."
    },
    "sql.defaults.serial_normalization": {
      "value": "0",
      "type": "e",
      "description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
    },
    "sql.distsql.distribute_index_joins": {
      "value": "true",
      "type": "b",
      "description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
    },
    "sql.distsql.flow_stream_timeout": {
      "value": "10s",
      "type": "d",
      "description": "amount of time incoming streams wait for a flow to be set up before erroring out"
    },
    "sql.distsql.interleaved_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set we plan interleaved table joins instead of merge joins when possible"
    },
    "sql.distsql.max_running_flows": {
      "value": "500",
      "type": "i",
      "description": "maximum number of concurrent flows that can be run on a node"
    },
    "sql.distsql.merge_joins.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, we plan merge joins when possible"
    },
    "sql.distsql.temp_storage.joins": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql joins"
    },
    "sql.distsql.temp_storage.sorts": {
      "value": "true",
      "type": "b",
      "description": "set to true to enable use of disk for distributed sql sorts"
    },
    "sql.distsql.temp_storage.workmem": {
      "value": "64 MiB",
      "type": "z",
      "description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
    },
    "sql.metrics.statement_details.dump_to_logs": {
      "value": "false",
      "type": "b",
      "description": "dump collected statement statistics to node logs when periodically cleared"
    },
    "sql.metrics.statement_details.enabled": {
      "value": "true",
      "type": "b",
      "description": "collect per-statement query statistics"
    },
    "sql.metrics.statement_details.plan_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "periodically save a logical plan for each fingerprint"
    },
    "sql.metrics.statement_details.plan_collection.period": {
      "value": "5m0s",
      "type": "d",
      "description": "the time until a new logical plan is collected"
    },
    "sql.metrics.statement_details.threshold": {
      "value": "0s",
      "type": "d",
      "description": "minimum execution time to cause statistics to be collected"
    },
    "sql.parallel_scans.enabled": {
      "value": "true",
      "type": "b",
      "description": "parallelizes scanning different ranges when the maximum result size can be deduced"
    },
    "sql.query_cache.enabled": {
      "value": "true",
      "type": "b",
      "description": "enable the query cache"
    },
    "sql.stats.automatic_collection.enabled": {
      "value": "true",
      "type": "b",
      "description": "automatic statistics collection mode"
    },
    "sql.stats.automatic_collection.fraction_stale_rows": {
      "value": "0.2",
      "type": "f",
      "description": "target fraction of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.automatic_collection.max_fraction_idle": {
      "value": "0.9",
      "type": "f",
      "description": "maximum fraction of time that automatic statistics sampler processors are idle"
    },
    "sql.stats.automatic_collection.min_stale_rows": {
      "value": "500",
      "type": "i",
      "description": "target minimum number of stale rows per table that will trigger a statistics refresh"
    },
    "sql.stats.max_timestamp_age": {
      "value": "5m0s",
      "type": "d",
      "description": "maximum age of timestamp during table statistics collection"
    },
    "sql.stats.post_events.enabled": {
      "value": "false",
      "type": "b",
      "description": "if set, an event is shown for every CREATE STATISTICS job"
    },
    "sql.tablecache.lease.refresh_limit": {
      "value": "50",
      "type": "i",
      "description": "maximum number of tables to periodically refresh leases for"
    },
    "sql.trace.log_statement_execute": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable logging of executed statements"
    },
    "sql.trace.session_eventlog.enabled": {
      "value": "false",
      "type": "b",
      "description": "set to true to enable session tracing"
    },
    "sql.trace.txn.enable_threshold": {
      "value": "0s",
      "type": "d",
      "description": "duration beyond which all transactions are traced (set to 0 to disable)"
    },
    "timeseries.storage.enabled": {
      "value": "true",
      "type": "b",
      "description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
    },
    "timeseries.storage.resolution_10s.ttl": {
      "value": "240h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
    },
    "timeseries.storage.resolution_30m.ttl": {
      "value": "2160h0m0s",
      "type": "d",
      "description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
    },
    "trace.debug.enable": {
      "value": "false",
      "type": "b",
      "description": "if set, traces for recent requests can be seen in the /debug page"
    },
    "trace.lightstep.token": {
      "type": "s",
      "description": "if set, traces go to Lightstep using this token"
    },
    "trace.zipkin.collector": {
      "type": "s",
      "description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set"
    },
    "version": {
      "value": "19.1",
      "type": "m",
      "description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'"
    }
  }
}
`
const DefaultClusterSettings19_1_3 string = `
{
"key_values": {
"changefeed.experimental_poll_interval": {
"value": "1s",
"type": "d",
"description": "polling interval for the prototype changefeed implementation (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
},
"changefeed.push.enabled": {
"value": "true",
"type": "b",
"description": "if set, changed are pushed instead of pulled. This requires the kv.rangefeed.enabled setting. See https://www.cockroachlabs.com/docs/v19.1/change-data-capture.html#enable-rangefeeds-to-reduce-latency"
},
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"compactor.max_record_age": {
"value": "24h0m0s",
"type": "d",
"description": "discard suggestions not processed within this duration (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
},
"compactor.min_interval": {
"value": "15s",
"type": "d",
"description": "minimum time interval to wait before compacting (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
},
"compactor.threshold_available_fraction": {
"value": "0.1",
"type": "f",
"description": "consider suggestions for at least the given percentage of the available logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
},
"compactor.threshold_bytes": {
"value": "256 MiB",
"type": "z",
"description": "minimum expected logical space reclamation required before considering an aggregated suggestion (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
},
"compactor.threshold_used_fraction": {
"value": "0.1",
"type": "f",
"description": "consider suggestions for at least the given percentage of the used logical space (zero to disable) (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.forced_stat_reset.interval": {
"value": "2h0m0s",
"type": "d",
"description": "interval after which pending diagnostics statistics should be discarded even if not reported"
},
"diagnostics.reporting.enabled": {
"value": "true",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "10s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"jobs.retention_time": {
"value": "336h0m0s",
"type": "d",
"description": "the amount of time to retain records for completed jobs before"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.load_based_rebalancing": {
"value": "2",
"type": "e",
"description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
},
"kv.allocator.qps_rebalance_threshold": {
"value": "0.25",
"type": "f",
"description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
},
"kv.allocator.range_rebalance_threshold": {
"value": "0.05",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.addsstable_max_rate": {
"value": "1.7976931348623157E+308",
"type": "f",
"description": "maximum number of AddSSTable requests per second for a single store"
},
"kv.bulk_io_write.concurrent_addsstable_requests": {
"value": "1",
"type": "i",
"description": "number of AddSSTable requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.concurrent_export_requests": {
"value": "3",
"type": "i",
"description": "number of export requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.concurrent_import_requests": {
"value": "1",
"type": "i",
"description": "number of import requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.max_rate": {
"value": "1.0 TiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.closed_timestamp.close_fraction": {
"value": "0.2",
"type": "f",
"description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
},
"kv.closed_timestamp.follower_reads_enabled": {
"value": "true",
"type": "b",
"description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
},
"kv.closed_timestamp.target_duration": {
"value": "30s",
"type": "d",
"description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
},
"kv.follower_read.target_multiple": {
"value": "3",
"type": "f",
"description": "if above 1, encourages the distsender to perform a read against the closest replica if a request is older than kv.closed_timestamp.target_duration * (1 + kv.closed_timestamp.close_fraction * this) less a clock uncertainty interval. This value also is used to create follower_timestamp(). (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
},
"kv.import.batch_size": {
"value": "32 MiB",
"type": "z",
"description": "the maximum size of the payload in an AddSSTable request (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.disable_synchronization_unsafe": {
"value": "true",
"type": "b",
"description": "set to true to disable synchronization on Raft log writes to persistent storage. Setting to true risks data loss or data corruption on server crashes. The setting is meant for internal testing only and SHOULD NOT be used in production."
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.range_merge.queue_enabled": {
"value": "true",
"type": "b",
"description": "whether the automatic merge queue is enabled"
},
"kv.range_merge.queue_interval": {
"value": "1s",
"type": "d",
"description": "how long the merge queue waits between processing replicas (WARNING: may compromise cluster stability or correctness; do not edit without supervision)"
},
"kv.range_split.by_load_enabled": {
"value": "true",
"type": "b",
"description": "allow automatic splits of ranges based on where load is concentrated"
},
"kv.range_split.load_qps_threshold": {
"value": "250",
"type": "i",
"description": "the QPS over which, the range becomes a candidate for load based splitting"
},
"kv.rangefeed.concurrent_catchup_iterators": {
"value": "64",
"type": "i",
"description": "number of rangefeeds catchup iterators a store will allow concurrently before queueing"
},
"kv.rangefeed.enabled": {
"value": "false",
"type": "b",
"description": "if set, rangefeed registration is enabled"
},
"kv.snapshot_rebalance.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance and upreplication snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "262144",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"kv.transaction.write_pipelining_enabled": {
"value": "true",
"type": "b",
"description": "if enabled, transactional writes are pipelined through Raft consensus"
},
"kv.transaction.write_pipelining_max_batch_size": {
"value": "128",
"type": "i",
"description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
},
"kv.transaction.write_pipelining_max_outstanding_size": {
"value": "256 KiB",
"type": "z",
"description": "maximum number of bytes used to track in-flight pipelined writes before disabling pipelining"
},
"rocksdb.ingest_backpressure.delay_l0_file": {
"value": "200ms",
"type": "d",
"description": "delay to add to SST ingestions per file in L0 over the configured limit"
},
"rocksdb.ingest_backpressure.l0_file_count_threshold": {
"value": "20",
"type": "i",
"description": "number of L0 files after which to backpressure SST ingestions"
},
"rocksdb.ingest_backpressure.max_delay": {
"value": "5s",
"type": "d",
"description": "maximum amount of time to backpressure a single SST ingestion"
},
"rocksdb.ingest_backpressure.pending_compaction_threshold": {
"value": "64 GiB",
"type": "z",
"description": "pending compaction estimate above which to backpressure SST ingestions"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"schemachanger.backfiller.buffer_size": {
"value": "196 MiB",
"type": "z",
"description": "amount to buffer in memory during backfills"
},
"schemachanger.backfiller.max_sst_size": {
"value": "16 MiB",
"type": "z",
"description": "target size for ingested files during backfills"
},
"schemachanger.bulk_index_backfill.batch_size": {
"value": "5000",
"type": "i",
"description": "number of rows to process at a time during bulk index backfill"
},
"schemachanger.bulk_index_backfill.enabled": {
"value": "true",
"type": "b",
"description": "backfill indexes in bulk via addsstable"
},
"schemachanger.lease.duration": {
"value": "5m0s",
"type": "d",
"description": "the duration of a schema change lease"
},
"schemachanger.lease.renew_fraction": {
"value": "0.5",
"type": "f",
"description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic"
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.eventlog.ttl": {
"value": "2160h0m0s",
"type": "d",
"description": "if nonzero, event log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.goroutine_dump.num_goroutines_threshold": {
"value": "1000",
"type": "i",
"description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
},
"server.goroutine_dump.total_dump_size_limit": {
"value": "500 MiB",
"type": "z",
"description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept."
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "0.85",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.host_based_authentication.configuration": {
"type": "s",
"description": "host-based authentication configuration to use during connection authentication"
},
"server.rangelog.ttl": {
"value": "720h0m0s",
"type": "d",
"description": "if nonzero, range log entries older than this duration are deleted every 10m0s. Should not be lowered below 24 hours."
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.default_int_size": {
"value": "8",
"type": "i",
"description": "the size, in bytes, of an INT type"
},
"sql.defaults.distsql": {
"value": "1",
"type": "e",
"description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
},
"sql.defaults.experimental_vectorize": {
"value": "0",
"type": "e",
"description": "default experimental_vectorize mode [off = 0, on = 1, always = 2]"
},
"sql.defaults.optimizer": {
"value": "1",
"type": "e",
"description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
},
"sql.defaults.reorder_joins_limit": {
"value": "4",
"type": "i",
"description": "default number of joins to reorder"
},
"sql.defaults.results_buffer.size": {
"value": "16 KiB",
"type": "z",
"description": "default size of the buffer that accumulates results for a statement or a batch of statements before they are sent to the client. This can be overridden on an individual connection with the 'results_buffer_size' parameter. Note that auto-retries generally only happen while no results have been delivered to the client, so reducing this size can increase the number of retriable errors a client receives. On the other hand, increasing the buffer size can increase the delay until the client receives the first result row. Updating the setting only affects new connections. Setting to 0 disables any buffering."
},
"sql.defaults.serial_normalization": {
"value": "0",
"type": "e",
"description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.flow_stream_timeout": {
"value": "10s",
"type": "d",
"description": "amount of time incoming streams wait for a flow to be set up before erroring out"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.max_running_flows": {
"value": "500",
"type": "i",
"description": "maximum number of concurrent flows that can be run on a node"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.plan_collection.enabled": {
"value": "true",
"type": "b",
"description": "periodically save a logical plan for each fingerprint"
},
"sql.metrics.statement_details.plan_collection.period": {
"value": "5m0s",
"type": "d",
"description": "the time until a new logical plan is collected"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.parallel_scans.enabled": {
"value": "true",
"type": "b",
"description": "parallelizes scanning different ranges when the maximum result size can be deduced"
},
"sql.query_cache.enabled": {
"value": "true",
"type": "b",
"description": "enable the query cache"
},
"sql.stats.automatic_collection.enabled": {
"value": "true",
"type": "b",
"description": "automatic statistics collection mode"
},
"sql.stats.automatic_collection.fraction_stale_rows": {
"value": "0.2",
"type": "f",
"description": "target fraction of stale rows per table that will trigger a statistics refresh"
},
"sql.stats.automatic_collection.max_fraction_idle": {
"value": "0.9",
"type": "f",
"description": "maximum fraction of time that automatic statistics sampler processors are idle"
},
"sql.stats.automatic_collection.min_stale_rows": {
"value": "500",
"type": "i",
"description": "target minimum number of stale rows per table that will trigger a statistics refresh"
},
"sql.stats.max_timestamp_age": {
"value": "5m0s",
"type": "d",
"description": "maximum age of timestamp during table statistics collection"
},
"sql.stats.post_events.enabled": {
"value": "false",
"type": "b",
"description": "if set, an event is shown for every CREATE STATISTICS job"
},
"sql.tablecache.lease.refresh_limit": {
"value": "50",
"type": "i",
"description": "maximum number of tables to periodically refresh leases for"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"timeseries.storage.resolution_10s.ttl": {
"value": "240h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
},
"timeseries.storage.resolution_30m.ttl": {
"value": "2160h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set"
},
"version": {
"value": "19.1",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'"
}
}
}
`
const DefaultClusterSettings2_1_11 string = `
{
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.forced_stat_reset.interval": {
"value": "2h0m0s",
"type": "d",
"description": "interval after which pending diagnostics statistics should be discarded even if not reported"
},
"diagnostics.reporting.enabled": {
"value": "true",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "10s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.load_based_rebalancing": {
"value": "2",
"type": "e",
"description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
},
"kv.allocator.qps_rebalance_threshold": {
"value": "0.25",
"type": "f",
"description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
},
"kv.allocator.range_rebalance_threshold": {
"value": "0.05",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.concurrent_export_requests": {
"value": "5",
"type": "i",
"description": "number of export requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.concurrent_import_requests": {
"value": "1",
"type": "i",
"description": "number of import requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.closed_timestamp.close_fraction": {
"value": "0.2",
"type": "f",
"description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
},
"kv.closed_timestamp.follower_reads_enabled": {
"value": "false",
"type": "b",
"description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
},
"kv.closed_timestamp.target_duration": {
"value": "30s",
"type": "d",
"description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage ('false' risks data loss)"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.rangefeed.enabled": {
"value": "false",
"type": "b",
"description": "if set, rangefeed registration is enabled"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"kv.transaction.write_pipelining_enabled": {
"value": "true",
"type": "b",
"description": "if enabled, transactional writes are pipelined through Raft consensus"
},
"kv.transaction.write_pipelining_max_batch_size": {
"value": "128",
"type": "i",
"description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"schemachanger.lease.duration": {
"value": "5m0s",
"type": "d",
"description": "the duration of a schema change lease"
},
"schemachanger.lease.renew_fraction": {
"value": "0.4",
"type": "f",
"description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "0.85",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "1",
"type": "e",
"description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2, 2.0-off = 3, 2.0-auto = 4]"
},
"sql.defaults.optimizer": {
"value": "1",
"type": "e",
"description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
},
"sql.defaults.serial_normalization": {
"value": "0",
"type": "e",
"description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.flow_stream_timeout": {
"value": "10s",
"type": "d",
"description": "amount of time incoming streams wait for a flow to be set up before erroring out"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.max_running_flows": {
"value": "500",
"type": "i",
"description": "maximum number of concurrent flows that can be run on a node"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.tablecache.lease.refresh_limit": {
"value": "50",
"type": "i",
"description": "maximum number of tables to periodically refresh leases for"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "720h0m0s",
"type": "d",
"description": "deprecated setting: the amount of time to store timeseries data. Replaced by timeseries.storage.10s_resolution_ttl."
},
"timeseries.storage.10s_resolution_ttl": {
"value": "240h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
},
"timeseries.storage.30m_resolution_ttl": {
"value": "2160h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.1",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}

v2_1_4_cluster_settings = {
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.client_cert": {
"type": "s",
"description": "custom client certificate for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.client_key": {
"type": "s",
"description": "custom client key (paired with client cert) for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.http.server_name": {
"type": "s",
"description": "custom server name for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.reporting.enabled": {
"value": "false",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "1m0s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"external.graphite.whitelistpath": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, enable whitelist for which metrics to push based on contents at this path. File contains a regex on each line. If any line matches metric name, it is white-listed."
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1E+00",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.range_rebalance_threshold": {
"value": "5E-02",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.allocator.stat_based_rebalancing.enabled": {
"value": "false",
"type": "b",
"description": "set to enable rebalancing of range replicas based on write load and disk usage"
},
"kv.allocator.stat_rebalance_threshold": {
"value": "2E-01",
"type": "f",
"description": "minimum fraction away from the mean a store's stats (like disk usage or writes per second) can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2E+00",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "If enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.goroutine_dump.num_goroutines_threshold": {
"value": "1000",
"type": "i",
"description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
},
"server.goroutine_dump.total_dump_size_limit": {
"value": "500 MiB",
"type": "z",
"description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "1E-01",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "0",
"type": "e",
"description": "Default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "168h0m0s",
"type": "d",
"description": "the amount of time to store timeseries data"
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.0",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}
`
const DefaultClusterSettings2_1_10 string = `
{
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.forced_stat_reset.interval": {
"value": "2h0m0s",
"type": "d",
"description": "interval after which pending diagnostics statistics should be discarded even if not reported"
},
"diagnostics.reporting.enabled": {
"value": "true",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "10s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.load_based_rebalancing": {
"value": "2",
"type": "e",
"description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
},
"kv.allocator.qps_rebalance_threshold": {
"value": "0.25",
"type": "f",
"description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
},
"kv.allocator.range_rebalance_threshold": {
"value": "0.05",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.concurrent_export_requests": {
"value": "5",
"type": "i",
"description": "number of export requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.concurrent_import_requests": {
"value": "1",
"type": "i",
"description": "number of import requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.closed_timestamp.close_fraction": {
"value": "0.2",
"type": "f",
"description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
},
"kv.closed_timestamp.follower_reads_enabled": {
"value": "false",
"type": "b",
"description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
},
"kv.closed_timestamp.target_duration": {
"value": "30s",
"type": "d",
"description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage ('false' risks data loss)"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.rangefeed.enabled": {
"value": "false",
"type": "b",
"description": "if set, rangefeed registration is enabled"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"kv.transaction.write_pipelining_enabled": {
"value": "true",
"type": "b",
"description": "if enabled, transactional writes are pipelined through Raft consensus"
},
"kv.transaction.write_pipelining_max_batch_size": {
"value": "128",
"type": "i",
"description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"schemachanger.lease.duration": {
"value": "5m0s",
"type": "d",
"description": "the duration of a schema change lease"
},
"schemachanger.lease.renew_fraction": {
"value": "0.4",
"type": "f",
"description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "0.85",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "1",
"type": "e",
"description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2, 2.0-off = 3, 2.0-auto = 4]"
},
"sql.defaults.optimizer": {
"value": "1",
"type": "e",
"description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
},
"sql.defaults.serial_normalization": {
"value": "0",
"type": "e",
"description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.flow_stream_timeout": {
"value": "10s",
"type": "d",
"description": "amount of time incoming streams wait for a flow to be set up before erroring out"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.max_running_flows": {
"value": "500",
"type": "i",
"description": "maximum number of concurrent flows that can be run on a node"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.tablecache.lease.refresh_limit": {
"value": "50",
"type": "i",
"description": "maximum number of tables to periodically refresh leases for"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "720h0m0s",
"type": "d",
"description": "deprecated setting: the amount of time to store timeseries data. Replaced by timeseries.storage.10s_resolution_ttl."
},
"timeseries.storage.10s_resolution_ttl": {
"value": "240h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
},
"timeseries.storage.30m_resolution_ttl": {
"value": "2160h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.1",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}
`
const DefaultClusterSettings2_1_9 string = `
{
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.forced_stat_reset.interval": {
"value": "2h0m0s",
"type": "d",
"description": "interval after which pending diagnostics statistics should be discarded even if not reported"
},
"diagnostics.reporting.enabled": {
"value": "true",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "10s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.load_based_rebalancing": {
"value": "2",
"type": "e",
"description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
},
"kv.allocator.qps_rebalance_threshold": {
"value": "0.25",
"type": "f",
"description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
},
"kv.allocator.range_rebalance_threshold": {
"value": "0.05",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.concurrent_export_requests": {
"value": "5",
"type": "i",
"description": "number of export requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.concurrent_import_requests": {
"value": "1",
"type": "i",
"description": "number of import requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.closed_timestamp.close_fraction": {
"value": "0.2",
"type": "f",
"description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
},
"kv.closed_timestamp.follower_reads_enabled": {
"value": "false",
"type": "b",
"description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
},
"kv.closed_timestamp.target_duration": {
"value": "30s",
"type": "d",
"description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage ('false' risks data loss)"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.rangefeed.enabled": {
"value": "false",
"type": "b",
"description": "if set, rangefeed registration is enabled"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"kv.transaction.write_pipelining_enabled": {
"value": "true",
"type": "b",
"description": "if enabled, transactional writes are pipelined through Raft consensus"
},
"kv.transaction.write_pipelining_max_batch_size": {
"value": "128",
"type": "i",
"description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"schemachanger.lease.duration": {
"value": "5m0s",
"type": "d",
"description": "the duration of a schema change lease"
},
"schemachanger.lease.renew_fraction": {
"value": "0.4",
"type": "f",
"description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "0.85",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "1",
"type": "e",
"description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2, 2.0-off = 3, 2.0-auto = 4]"
},
"sql.defaults.optimizer": {
"value": "1",
"type": "e",
"description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
},
"sql.defaults.serial_normalization": {
"value": "0",
"type": "e",
"description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.flow_stream_timeout": {
"value": "10s",
"type": "d",
"description": "amount of time incoming streams wait for a flow to be set up before erroring out"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.max_running_flows": {
"value": "500",
"type": "i",
"description": "maximum number of concurrent flows that can be run on a node"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.tablecache.lease.refresh_limit": {
"value": "50",
"type": "i",
"description": "maximum number of tables to periodically refresh leases for"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "720h0m0s",
"type": "d",
"description": "deprecated setting: the amount of time to store timeseries data. Replaced by timeseries.storage.10s_resolution_ttl."
},
"timeseries.storage.10s_resolution_ttl": {
"value": "240h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
},
"timeseries.storage.30m_resolution_ttl": {
"value": "2160h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.1",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}

v2_1_4_cluster_settings = {
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.client_cert": {
"type": "s",
"description": "custom client certificate for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.client_key": {
"type": "s",
"description": "custom client key (paired with client cert) for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.http.server_name": {
"type": "s",
"description": "custom server name for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.reporting.enabled": {
"value": "false",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "1m0s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"external.graphite.whitelistpath": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, enable whitelist for which metrics to push based on contents at this path. File contains a regex on each line. If any line matches metric name, it is white-listed."
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1E+00",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.range_rebalance_threshold": {
"value": "5E-02",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.allocator.stat_based_rebalancing.enabled": {
"value": "false",
"type": "b",
"description": "set to enable rebalancing of range replicas based on write load and disk usage"
},
"kv.allocator.stat_rebalance_threshold": {
"value": "2E-01",
"type": "f",
"description": "minimum fraction away from the mean a store's stats (like disk usage or writes per second) can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2E+00",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "If enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.goroutine_dump.num_goroutines_threshold": {
"value": "1000",
"type": "i",
"description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
},
"server.goroutine_dump.total_dump_size_limit": {
"value": "500 MiB",
"type": "z",
"description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "1E-01",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "0",
"type": "e",
"description": "Default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "168h0m0s",
"type": "d",
"description": "the amount of time to store timeseries data"
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.0",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}
`
const DefaultClusterSettings2_1_8 string = `
{
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.forced_stat_reset.interval": {
"value": "2h0m0s",
"type": "d",
"description": "interval after which pending diagnostics statistics should be discarded even if not reported"
},
"diagnostics.reporting.enabled": {
"value": "true",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "10s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.load_based_rebalancing": {
"value": "2",
"type": "e",
"description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
},
"kv.allocator.qps_rebalance_threshold": {
"value": "0.25",
"type": "f",
"description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
},
"kv.allocator.range_rebalance_threshold": {
"value": "0.05",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.concurrent_export_requests": {
"value": "5",
"type": "i",
"description": "number of export requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.concurrent_import_requests": {
"value": "1",
"type": "i",
"description": "number of import requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.closed_timestamp.close_fraction": {
"value": "0.2",
"type": "f",
"description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
},
"kv.closed_timestamp.follower_reads_enabled": {
"value": "false",
"type": "b",
"description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
},
"kv.closed_timestamp.target_duration": {
"value": "30s",
"type": "d",
"description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage ('false' risks data loss)"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.rangefeed.enabled": {
"value": "false",
"type": "b",
"description": "if set, rangefeed registration is enabled"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"kv.transaction.write_pipelining_enabled": {
"value": "true",
"type": "b",
"description": "if enabled, transactional writes are pipelined through Raft consensus"
},
"kv.transaction.write_pipelining_max_batch_size": {
"value": "128",
"type": "i",
"description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"schemachanger.lease.duration": {
"value": "5m0s",
"type": "d",
"description": "the duration of a schema change lease"
},
"schemachanger.lease.renew_fraction": {
"value": "0.4",
"type": "f",
"description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "0.85",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "1",
"type": "e",
"description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2, 2.0-off = 3, 2.0-auto = 4]"
},
"sql.defaults.optimizer": {
"value": "1",
"type": "e",
"description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
},
"sql.defaults.serial_normalization": {
"value": "0",
"type": "e",
"description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.flow_stream_timeout": {
"value": "10s",
"type": "d",
"description": "amount of time incoming streams wait for a flow to be set up before erroring out"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.max_running_flows": {
"value": "500",
"type": "i",
"description": "maximum number of concurrent flows that can be run on a node"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.tablecache.lease.refresh_limit": {
"value": "50",
"type": "i",
"description": "maximum number of tables to periodically refresh leases for"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "720h0m0s",
"type": "d",
"description": "deprecated setting: the amount of time to store timeseries data. Replaced by timeseries.storage.10s_resolution_ttl."
},
"timeseries.storage.10s_resolution_ttl": {
"value": "240h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
},
"timeseries.storage.30m_resolution_ttl": {
"value": "2160h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.1",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}

v2_1_4_cluster_settings = {
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.client_cert": {
"type": "s",
"description": "custom client certificate for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.client_key": {
"type": "s",
"description": "custom client key (paired with client cert) for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.http.server_name": {
"type": "s",
"description": "custom server name for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.reporting.enabled": {
"value": "false",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "1m0s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"external.graphite.whitelistpath": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, enable whitelist for which metrics to push based on contents at this path. File contains a regex on each line. If any line matches metric name, it is white-listed."
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1E+00",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.range_rebalance_threshold": {
"value": "5E-02",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.allocator.stat_based_rebalancing.enabled": {
"value": "false",
"type": "b",
"description": "set to enable rebalancing of range replicas based on write load and disk usage"
},
"kv.allocator.stat_rebalance_threshold": {
"value": "2E-01",
"type": "f",
"description": "minimum fraction away from the mean a store's stats (like disk usage or writes per second) can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2E+00",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "If enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.goroutine_dump.num_goroutines_threshold": {
"value": "1000",
"type": "i",
"description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
},
"server.goroutine_dump.total_dump_size_limit": {
"value": "500 MiB",
"type": "z",
"description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "1E-01",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "0",
"type": "e",
"description": "Default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "168h0m0s",
"type": "d",
"description": "the amount of time to store timeseries data"
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.0",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}
`
const DefaultClusterSettings2_1_7 string = `
{
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.forced_stat_reset.interval": {
"value": "2h0m0s",
"type": "d",
"description": "interval after which pending diagnostics statistics should be discarded even if not reported"
},
"diagnostics.reporting.enabled": {
"value": "true",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "10s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.load_based_rebalancing": {
"value": "2",
"type": "e",
"description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
},
"kv.allocator.qps_rebalance_threshold": {
"value": "0.25",
"type": "f",
"description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
},
"kv.allocator.range_rebalance_threshold": {
"value": "0.05",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.concurrent_export_requests": {
"value": "5",
"type": "i",
"description": "number of export requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.concurrent_import_requests": {
"value": "1",
"type": "i",
"description": "number of import requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.closed_timestamp.close_fraction": {
"value": "0.2",
"type": "f",
"description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
},
"kv.closed_timestamp.follower_reads_enabled": {
"value": "false",
"type": "b",
"description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
},
"kv.closed_timestamp.target_duration": {
"value": "30s",
"type": "d",
"description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage ('false' risks data loss)"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.rangefeed.enabled": {
"value": "false",
"type": "b",
"description": "if set, rangefeed registration is enabled"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"kv.transaction.write_pipelining_enabled": {
"value": "true",
"type": "b",
"description": "if enabled, transactional writes are pipelined through Raft consensus"
},
"kv.transaction.write_pipelining_max_batch_size": {
"value": "128",
"type": "i",
"description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"schemachanger.lease.duration": {
"value": "5m0s",
"type": "d",
"description": "the duration of a schema change lease"
},
"schemachanger.lease.renew_fraction": {
"value": "0.4",
"type": "f",
"description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "0.85",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "1",
"type": "e",
"description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2, 2.0-off = 3, 2.0-auto = 4]"
},
"sql.defaults.optimizer": {
"value": "1",
"type": "e",
"description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
},
"sql.defaults.serial_normalization": {
"value": "0",
"type": "e",
"description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.flow_stream_timeout": {
"value": "10s",
"type": "d",
"description": "amount of time incoming streams wait for a flow to be set up before erroring out"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.max_running_flows": {
"value": "500",
"type": "i",
"description": "maximum number of concurrent flows that can be run on a node"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.tablecache.lease.refresh_limit": {
"value": "50",
"type": "i",
"description": "maximum number of tables to periodically refresh leases for"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "720h0m0s",
"type": "d",
"description": "deprecated setting: the amount of time to store timeseries data. Replaced by timeseries.storage.10s_resolution_ttl."
},
"timeseries.storage.10s_resolution_ttl": {
"value": "240h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
},
"timeseries.storage.30m_resolution_ttl": {
"value": "2160h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.1",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}

v2_1_4_cluster_settings = {
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.client_cert": {
"type": "s",
"description": "custom client certificate for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.client_key": {
"type": "s",
"description": "custom client key (paired with client cert) for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.http.server_name": {
"type": "s",
"description": "custom server name for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.reporting.enabled": {
"value": "false",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "1m0s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"external.graphite.whitelistpath": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, enable whitelist for which metrics to push based on contents at this path. File contains a regex on each line. If any line matches metric name, it is white-listed."
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1E+00",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.range_rebalance_threshold": {
"value": "5E-02",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.allocator.stat_based_rebalancing.enabled": {
"value": "false",
"type": "b",
"description": "set to enable rebalancing of range replicas based on write load and disk usage"
},
"kv.allocator.stat_rebalance_threshold": {
"value": "2E-01",
"type": "f",
"description": "minimum fraction away from the mean a store's stats (like disk usage or writes per second) can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2E+00",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "If enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.goroutine_dump.num_goroutines_threshold": {
"value": "1000",
"type": "i",
"description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
},
"server.goroutine_dump.total_dump_size_limit": {
"value": "500 MiB",
"type": "z",
"description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "1E-01",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "0",
"type": "e",
"description": "Default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "168h0m0s",
"type": "d",
"description": "the amount of time to store timeseries data"
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.0",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}
`
const DefaultClusterSettings2_1_6 string = `
{
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.forced_stat_reset.interval": {
"value": "2h0m0s",
"type": "d",
"description": "interval after which pending diagnostics statistics should be discarded even if not reported"
},
"diagnostics.reporting.enabled": {
"value": "true",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "10s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.load_based_rebalancing": {
"value": "2",
"type": "e",
"description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
},
"kv.allocator.qps_rebalance_threshold": {
"value": "0.25",
"type": "f",
"description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
},
"kv.allocator.range_rebalance_threshold": {
"value": "0.05",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.concurrent_export_requests": {
"value": "5",
"type": "i",
"description": "number of export requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.concurrent_import_requests": {
"value": "1",
"type": "i",
"description": "number of import requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.closed_timestamp.close_fraction": {
"value": "0.2",
"type": "f",
"description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
},
"kv.closed_timestamp.follower_reads_enabled": {
"value": "false",
"type": "b",
"description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
},
"kv.closed_timestamp.target_duration": {
"value": "30s",
"type": "d",
"description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage ('false' risks data loss)"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.rangefeed.enabled": {
"value": "false",
"type": "b",
"description": "if set, rangefeed registration is enabled"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"kv.transaction.write_pipelining_enabled": {
"value": "true",
"type": "b",
"description": "if enabled, transactional writes are pipelined through Raft consensus"
},
"kv.transaction.write_pipelining_max_batch_size": {
"value": "128",
"type": "i",
"description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"schemachanger.lease.duration": {
"value": "5m0s",
"type": "d",
"description": "the duration of a schema change lease"
},
"schemachanger.lease.renew_fraction": {
"value": "0.4",
"type": "f",
"description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "0.85",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "1",
"type": "e",
"description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2, 2.0-off = 3, 2.0-auto = 4]"
},
"sql.defaults.optimizer": {
"value": "1",
"type": "e",
"description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
},
"sql.defaults.serial_normalization": {
"value": "0",
"type": "e",
"description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.flow_stream_timeout": {
"value": "10s",
"type": "d",
"description": "amount of time incoming streams wait for a flow to be set up before erroring out"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.max_running_flows": {
"value": "500",
"type": "i",
"description": "maximum number of concurrent flows that can be run on a node"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.tablecache.lease.refresh_limit": {
"value": "50",
"type": "i",
"description": "maximum number of tables to periodically refresh leases for"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "720h0m0s",
"type": "d",
"description": "deprecated setting: the amount of time to store timeseries data. Replaced by timeseries.storage.10s_resolution_ttl."
},
"timeseries.storage.10s_resolution_ttl": {
"value": "240h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
},
"timeseries.storage.30m_resolution_ttl": {
"value": "2160h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.1",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}

v2_1_4_cluster_settings = {
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.client_cert": {
"type": "s",
"description": "custom client certificate for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.client_key": {
"type": "s",
"description": "custom client key (paired with client cert) for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.http.server_name": {
"type": "s",
"description": "custom server name for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.reporting.enabled": {
"value": "false",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "1m0s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"external.graphite.whitelistpath": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, enable whitelist for which metrics to push based on contents at this path. File contains a regex on each line. If any line matches metric name, it is white-listed."
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1E+00",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.range_rebalance_threshold": {
"value": "5E-02",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.allocator.stat_based_rebalancing.enabled": {
"value": "false",
"type": "b",
"description": "set to enable rebalancing of range replicas based on write load and disk usage"
},
"kv.allocator.stat_rebalance_threshold": {
"value": "2E-01",
"type": "f",
"description": "minimum fraction away from the mean a store's stats (like disk usage or writes per second) can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2E+00",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "If enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.goroutine_dump.num_goroutines_threshold": {
"value": "1000",
"type": "i",
"description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
},
"server.goroutine_dump.total_dump_size_limit": {
"value": "500 MiB",
"type": "z",
"description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "1E-01",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "0",
"type": "e",
"description": "Default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "168h0m0s",
"type": "d",
"description": "the amount of time to store timeseries data"
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.0",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}
`
const DefaultClusterSettings2_1_5 string = `
{
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.forced_stat_reset.interval": {
"value": "2h0m0s",
"type": "d",
"description": "interval after which pending diagnostics statistics should be discarded even if not reported"
},
"diagnostics.reporting.enabled": {
"value": "true",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported (should be shorter than diagnostics.forced_stat_reset.interval)"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "10s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.load_based_rebalancing": {
"value": "2",
"type": "e",
"description": "whether to rebalance based on the distribution of QPS across stores [off = 0, leases = 1, leases and replicas = 2]"
},
"kv.allocator.qps_rebalance_threshold": {
"value": "0.25",
"type": "f",
"description": "minimum fraction away from the mean a store's QPS (such as queries per second) can be before it is considered overfull or underfull"
},
"kv.allocator.range_rebalance_threshold": {
"value": "0.05",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.concurrent_export_requests": {
"value": "5",
"type": "i",
"description": "number of export requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.concurrent_import_requests": {
"value": "1",
"type": "i",
"description": "number of import requests a store will handle concurrently before queuing"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.closed_timestamp.close_fraction": {
"value": "0.2",
"type": "f",
"description": "fraction of closed timestamp target duration specifying how frequently the closed timestamp is advanced"
},
"kv.closed_timestamp.follower_reads_enabled": {
"value": "false",
"type": "b",
"description": "allow (all) replicas to serve consistent historical reads based on closed timestamp information"
},
"kv.closed_timestamp.target_duration": {
"value": "30s",
"type": "d",
"description": "if nonzero, attempt to provide closed timestamp notifications for timestamps trailing cluster time by approximately this duration"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage ('false' risks data loss)"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.rangefeed.enabled": {
"value": "false",
"type": "b",
"description": "if set, rangefeed registration is enabled"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"kv.transaction.write_pipelining_enabled": {
"value": "true",
"type": "b",
"description": "if enabled, transactional writes are pipelined through Raft consensus"
},
"kv.transaction.write_pipelining_max_batch_size": {
"value": "128",
"type": "i",
"description": "if non-zero, defines that maximum size batch that will be pipelined through Raft consensus"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"schemachanger.lease.duration": {
"value": "5m0s",
"type": "d",
"description": "the duration of a schema change lease"
},
"schemachanger.lease.renew_fraction": {
"value": "0.4",
"type": "f",
"description": "the fraction of schemachanger.lease_duration remaining to trigger a renew of the lease"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "if enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "0.85",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "1",
"type": "e",
"description": "default distributed SQL execution mode [off = 0, auto = 1, on = 2, 2.0-off = 3, 2.0-auto = 4]"
},
"sql.defaults.optimizer": {
"value": "1",
"type": "e",
"description": "default cost-based optimizer mode [off = 0, on = 1, local = 2]"
},
"sql.defaults.serial_normalization": {
"value": "0",
"type": "e",
"description": "default handling of SERIAL in table definitions [rowid = 0, virtual_sequence = 1, sql_sequence = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.flow_stream_timeout": {
"value": "10s",
"type": "d",
"description": "amount of time incoming streams wait for a flow to be set up before erroring out"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.max_running_flows": {
"value": "500",
"type": "i",
"description": "maximum number of concurrent flows that can be run on a node"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.tablecache.lease.refresh_limit": {
"value": "50",
"type": "i",
"description": "maximum number of tables to periodically refresh leases for"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "720h0m0s",
"type": "d",
"description": "deprecated setting: the amount of time to store timeseries data. Replaced by timeseries.storage.10s_resolution_ttl."
},
"timeseries.storage.10s_resolution_ttl": {
"value": "240h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 10 second resolution. Data older than this is subject to rollup and deletion."
},
"timeseries.storage.30m_resolution_ttl": {
"value": "2160h0m0s",
"type": "d",
"description": "the maximum age of time series data stored at the 30 minute resolution. Data older than this is subject to deletion."
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.1",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}

v2_1_4_cluster_settings = {
"key_values": {
"cloudstorage.gs.default.key": {
"type": "s",
"description": "if set, JSON key to use during Google Cloud Storage operations"
},
"cloudstorage.http.client_cert": {
"type": "s",
"description": "custom client certificate for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.client_key": {
"type": "s",
"description": "custom client key (paired with client cert) for server-side verification when interacting with HTTPS storage"
},
"cloudstorage.http.custom_ca": {
"type": "s",
"description": "custom root CA (appended to system's default CAs) for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.http.server_name": {
"type": "s",
"description": "custom server name for verifying certificates when interacting with HTTPS storage"
},
"cloudstorage.timeout": {
"value": "10m0s",
"type": "d",
"description": "the timeout for import/export storage operations"
},
"cluster.organization": {
"type": "s",
"description": "organization name"
},
"cluster.preserve_downgrade_option": {
"type": "s",
"description": "disable (automatic or manual) cluster version upgrade from the specified version until reset"
},
"compactor.enabled": {
"value": "true",
"type": "b",
"description": "when false, the system will reclaim space occupied by deleted data less aggressively"
},
"debug.panic_on_failed_assertions": {
"value": "false",
"type": "b",
"description": "panic when an assertion fails rather than reporting"
},
"diagnostics.reporting.enabled": {
"value": "false",
"type": "b",
"description": "enable reporting diagnostic metrics to cockroach labs"
},
"diagnostics.reporting.interval": {
"value": "1h0m0s",
"type": "d",
"description": "interval at which diagnostics data should be reported"
},
"diagnostics.reporting.send_crash_reports": {
"value": "true",
"type": "b",
"description": "send crash and panic reports"
},
"external.graphite.endpoint": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, push server metrics to the Graphite or Carbon server at the specified host:port"
},
"external.graphite.interval": {
"value": "1m0s",
"type": "d",
"description": "the interval at which metrics are pushed to Graphite (if enabled)"
},
"external.graphite.whitelistpath": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "if nonempty, enable whitelist for which metrics to push based on contents at this path. File contains a regex on each line. If any line matches metric name, it is white-listed."
},
"jobs.registry.leniency": {
"value": "1m0s",
"type": "d",
"description": "the amount of time to defer any attempts to reschedule a job"
},
"kv.allocator.lease_rebalancing_aggressiveness": {
"value": "1E+00",
"type": "f",
"description": "set greater than 1.0 to rebalance leases toward load more aggressively, or between 0 and 1.0 to be more conservative about rebalancing leases"
},
"kv.allocator.load_based_lease_rebalancing.enabled": {
"value": "true",
"type": "b",
"description": "set to enable rebalancing of range leases based on load and latency"
},
"kv.allocator.range_rebalance_threshold": {
"value": "5E-02",
"type": "f",
"description": "minimum fraction away from the mean a store's range count can be before it is considered overfull or underfull"
},
"kv.allocator.stat_based_rebalancing.enabled": {
"value": "false",
"type": "b",
"description": "set to enable rebalancing of range replicas based on write load and disk usage"
},
"kv.allocator.stat_rebalance_threshold": {
"value": "2E-01",
"type": "f",
"description": "minimum fraction away from the mean a store's stats (like disk usage or writes per second) can be before it is considered overfull or underfull"
},
"kv.bulk_io_write.max_rate": {
"value": "8.0 EiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for writes to disk on behalf of bulk io ops"
},
"kv.bulk_sst.sync_size": {
"value": "2.0 MiB",
"type": "z",
"description": "threshold after which non-Rocks SST writes must fsync (0 disables)"
},
"kv.raft.command.max_size": {
"value": "64 MiB",
"type": "z",
"description": "maximum size of a raft command"
},
"kv.raft_log.synchronize": {
"value": "true",
"type": "b",
"description": "set to true to synchronize on Raft log writes to persistent storage"
},
"kv.range.backpressure_range_size_multiplier": {
"value": "2E+00",
"type": "f",
"description": "multiple of range_max_bytes that a range is allowed to grow to without splitting before writes to that range are blocked, or 0 to disable"
},
"kv.range_descriptor_cache.size": {
"value": "1000000",
"type": "i",
"description": "maximum number of entries in the range descriptor and leaseholder caches"
},
"kv.snapshot_rebalance.max_rate": {
"value": "2.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for rebalance snapshots"
},
"kv.snapshot_recovery.max_rate": {
"value": "8.0 MiB",
"type": "z",
"description": "the rate limit (bytes/sec) to use for recovery snapshots"
},
"kv.transaction.max_intents_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track write intents in transactions"
},
"kv.transaction.max_refresh_spans_bytes": {
"value": "256000",
"type": "i",
"description": "maximum number of bytes used to track refresh spans in serializable transactions"
},
"rocksdb.min_wal_sync_interval": {
"value": "0s",
"type": "d",
"description": "minimum duration between syncs of the RocksDB WAL"
},
"server.clock.forward_jump_check_enabled": {
"value": "false",
"type": "b",
"description": "If enabled, forward clock jumps \u003e max_offset/2 will cause a panic."
},
"server.clock.persist_upper_bound_interval": {
"value": "0s",
"type": "d",
"description": "the interval between persisting the wall time upper bound of the clock. The clock does not generate a wall time greater than the persisted timestamp and will panic if it sees a wall time greater than this value. When cockroach starts, it waits for the wall time to catch-up till this persisted timestamp. This guarantees monotonic wall time across server restarts. Not setting this or setting a value of 0 disables this feature."
},
"server.consistency_check.interval": {
"value": "24h0m0s",
"type": "d",
"description": "the time between range consistency checks; set to 0 to disable consistency checking"
},
"server.declined_reservation_timeout": {
"value": "1s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a reservation was declined"
},
"server.failed_reservation_timeout": {
"value": "5s",
"type": "d",
"description": "the amount of time to consider the store throttled for up-replication after a failed reservation call"
},
"server.goroutine_dump.num_goroutines_threshold": {
"value": "1000",
"type": "i",
"description": "a threshold beyond which if number of goroutines increases, then goroutine dump can be triggered"
},
"server.goroutine_dump.total_dump_size_limit": {
"value": "500 MiB",
"type": "z",
"description": "total size of goroutine dumps to be kept. Dumps are GC'ed in the order of creation time. The latest dump is always kept even if its size exceeds the limit."
},
"server.heap_profile.max_profiles": {
"value": "5",
"type": "i",
"description": "maximum number of profiles to be kept. Profiles with lower score are GC'ed, but latest profile is always kept"
},
"server.heap_profile.system_memory_threshold_fraction": {
"value": "1E-01",
"type": "f",
"description": "fraction of system memory beyond which if Rss increases, then heap profile is triggered"
},
"server.remote_debugging.mode": {
"value": "\u003credacted\u003e",
"type": "s",
"description": "set to enable remote debugging, localhost-only or disable (any, local, off)"
},
"server.shutdown.drain_wait": {
"value": "0s",
"type": "d",
"description": "the amount of time a server waits in an unready state before proceeding with the rest of the shutdown process"
},
"server.shutdown.query_wait": {
"value": "10s",
"type": "d",
"description": "the server will wait for at least this amount of time for active queries to finish"
},
"server.time_until_store_dead": {
"value": "5m0s",
"type": "d",
"description": "the time after which if there is no new gossiped information about a store, it is considered dead"
},
"server.web_session_timeout": {
"value": "168h0m0s",
"type": "d",
"description": "the duration that a newly created web session will be valid"
},
"sql.defaults.distsql": {
"value": "0",
"type": "e",
"description": "Default distributed SQL execution mode [off = 0, auto = 1, on = 2]"
},
"sql.distsql.distribute_index_joins": {
"value": "true",
"type": "b",
"description": "if set, for index joins we instantiate a join reader on every node that has a stream; if not set, we use a single join reader"
},
"sql.distsql.interleaved_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set we plan interleaved table joins instead of merge joins when possible"
},
"sql.distsql.merge_joins.enabled": {
"value": "true",
"type": "b",
"description": "if set, we plan merge joins when possible"
},
"sql.distsql.temp_storage.joins": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql joins"
},
"sql.distsql.temp_storage.sorts": {
"value": "true",
"type": "b",
"description": "set to true to enable use of disk for distributed sql sorts"
},
"sql.distsql.temp_storage.workmem": {
"value": "64 MiB",
"type": "z",
"description": "maximum amount of memory in bytes a processor can use before falling back to temp storage"
},
"sql.metrics.statement_details.dump_to_logs": {
"value": "false",
"type": "b",
"description": "dump collected statement statistics to node logs when periodically cleared"
},
"sql.metrics.statement_details.enabled": {
"value": "true",
"type": "b",
"description": "collect per-statement query statistics"
},
"sql.metrics.statement_details.threshold": {
"value": "0s",
"type": "d",
"description": "minimum execution time to cause statistics to be collected"
},
"sql.trace.log_statement_execute": {
"value": "false",
"type": "b",
"description": "set to true to enable logging of executed statements"
},
"sql.trace.session_eventlog.enabled": {
"value": "false",
"type": "b",
"description": "set to true to enable session tracing"
},
"sql.trace.txn.enable_threshold": {
"value": "0s",
"type": "d",
"description": "duration beyond which all transactions are traced (set to 0 to disable)"
},
"timeseries.resolution_10s.storage_duration": {
"value": "168h0m0s",
"type": "d",
"description": "the amount of time to store timeseries data"
},
"timeseries.storage.enabled": {
"value": "true",
"type": "b",
"description": "if set, periodic timeseries data is stored within the cluster; disabling is not recommended unless you are storing the data elsewhere"
},
"trace.debug.enable": {
"value": "false",
"type": "b",
"description": "if set, traces for recent requests can be seen in the /debug page"
},
"trace.lightstep.token": {
"type": "s",
"description": "if set, traces go to Lightstep using this token"
},
"trace.zipkin.collector": {
"type": "s",
"description": "if set, traces go to the given Zipkin instance (example: '127.0.0.1:9411'); ignored if trace.lightstep.token is set."
},
"version": {
"value": "2.0",
"type": "m",
"description": "set the active cluster version in the format '\u003cmajor\u003e.\u003cminor\u003e'."
}
}
}
`





