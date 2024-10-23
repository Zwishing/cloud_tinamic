package pkg

// Minio中存储数据的存储桶的名称
const (
	OriginalSourceBucketName       = "original-source"
	CloudOptimizedSourceBucketName = "cloud-optimized-source"
	VectorTileBucketName           = "vector-tile"
)

const (
	SourceSchema              = "source"
	SourceInfoTable           = "source.info"
	SourceOriginalTable       = "source.original"
	SourceCloudOptimizedTable = "source.cloud_optimized"
)

const (
	VectorSchema           = "vector"
	ServiceSchema          = "service"
	ServiceCollectionTable = "service.collection"
	ServiceInfoTable       = "service.info"
	ServiceVectorTable     = "service.vector"
)

const (
	GeometryFieldName = "geom"
)

const (
	VectorProcessingTaskQueue = "vector-processing-pipeline"
)

const (
	VectorThumbnailWidth  = 300
	VectorThumbnailHeight = 300
)
