package objectstore

import (
	"context"

	"gocloud.dev/blob/s3blob"
	//   "gocloud.dev/blob/s3blob"
)

/*
provides a mapping between the filesystem object, the inode, and the
dir object with the blob?
*/

func NewBucket() {
	ctx = context.Background()
	bucket, err := s3blob.OpenBucket(ctx, sess, "mybucket")

	if err != nil {
		return bucket
	}
}
