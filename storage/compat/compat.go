package compat

import (
	"github.com/influxdata/platform/storage"
	"github.com/influxdata/platform/toml"
	"github.com/influxdata/platform/tsdb/tsm1"
)

// Old Config structure
type Config struct {
	Dir string `toml:"dir"` // This doesn't exist in the new config

	WALDir        string        `toml:"wal-dir"`
	WALFsyncDelay toml.Duration `toml:"wal-fsync-delay"`
	ValidateKeys  bool          `toml:"validate-keys"`

	CacheMaxMemorySize             toml.Size     `toml:"cache-max-memory-size"`
	CacheSnapshotMemorySize        toml.Size     `toml:"cache-snapshot-memory-size"`
	CacheSnapshotWriteColdDuration toml.Duration `toml:"cache-snapshot-write-cold-duration"`
	CompactFullWriteColdDuration   toml.Duration `toml:"compact-full-write-cold-duration"`
	CompactThroughput              toml.Size     `toml:"compact-throughput"`
	CompactThroughputBurst         toml.Size     `toml:"compact-throughput-burst"`

	MaxConcurrentCompactions int  `toml:"max-concurrent-compactions"`
	TraceLoggingEnabled      bool `toml:"trace-logging-enabled"`
	TSMWillNeed              bool `toml:"tsm-use-madv-willneed"`
}

func Create() Config {
	return Config{
		WALDir:                         tsm1.DefaultWALPath,
		WALFsyncDelay:                  toml.Duration(tsm1.DefaultWALFsyncDelay),
		ValidateKeys:                   storage.DefaultValidateKeys,
		CacheMaxMemorySize:             toml.Size(tsm1.DefaultCacheMaxMemorySize),
		CacheSnapshotMemorySize:        toml.Size(tsm1.DefaultCacheSnapshotMemorySize),
		CacheSnapshotWriteColdDuration: toml.Duration(tsm1.DefaultCacheSnapshotWriteColdDuration),
		CompactFullWriteColdDuration:   toml.Duration(tsm1.DefaultCompactFullWriteColdDuration),
		CompactThroughput:              toml.Size(tsm1.DefaultCompactThroughput),
		CompactThroughputBurst:         toml.Size(tsm1.DefaultCompactThroughputBurst),
		MaxConcurrentCompactions:       tsm1.DefaultCompactMaxConcurrent,
		TraceLoggingEnabled:            tsm1.DefaultTraceLoggingEnabled,
		TSMWillNeed:                    tsm1.DefaultMADVWillNeed,
	}
}

func Convert(oldConfig Config) storage.Config {
	newConfig := storage.NewConfig()
	newConfig.ValidateKeys = oldConfig.ValidateKeys
	newConfig.Engine.MADVWillNeed = oldConfig.TSMWillNeed
	newConfig.Engine.TraceLoggingEnabled = oldConfig.TraceLoggingEnabled
	newConfig.Engine.Cache.MaxMemorySize = oldConfig.CacheMaxMemorySize
	newConfig.Engine.Cache.SnapshotMemorySize = oldConfig.CacheSnapshotMemorySize
	newConfig.Engine.Cache.SnapshotWriteColdDuration = oldConfig.CacheSnapshotWriteColdDuration
	newConfig.Engine.Compaction.FullWriteColdDuration = oldConfig.CompactFullWriteColdDuration
	newConfig.Engine.Compaction.Throughput = oldConfig.CompactThroughput
	newConfig.Engine.Compaction.ThroughputBurst = oldConfig.CompactThroughputBurst
	newConfig.Engine.Compaction.MaxConcurrent = oldConfig.MaxConcurrentCompactions
	newConfig.Engine.WAL.Path = oldConfig.WALDir
	newConfig.Engine.WAL.FsyncDelay = oldConfig.WALFsyncDelay
	return newConfig
}
