package reply

import (
	"context"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestReplyInsert(t *testing.T) {
	convey.Convey("Insert", t, func(ctx convey.C) {
		var (
			c        = context.Background()
			adminid  = int64(0)
			oid      = int64(0)
			rpID     = int64(0)
			tp       = int8(0)
			result   = ""
			remark   = ""
			isnew    = int8(0)
			isreport = int8(0)
			state    = int8(0)
			now      = time.Now()
		)
		ctx.Convey("When everything gose positive", func(ctx convey.C) {
			id, err := d.Admin.Insert(c, adminid, oid, rpID, tp, result, remark, isnew, isreport, state, now)
			ctx.Convey("Then err should be nil.id should not be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldBeNil)
				ctx.So(id, convey.ShouldNotBeNil)
			})
		})
	})
}

func TestReplyUpIsNotNew(t *testing.T) {
	convey.Convey("UpIsNotNew", t, func(ctx convey.C) {
		var (
			c    = context.Background()
			rpID = int64(0)
			now  = time.Now()
		)
		ctx.Convey("When everything gose positive", func(ctx convey.C) {
			rows, err := d.Admin.UpIsNotNew(c, rpID, now)
			ctx.Convey("Then err should be nil.rows should not be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldBeNil)
				ctx.So(rows, convey.ShouldNotBeNil)
			})
		})
	})
}
