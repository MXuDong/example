package pkg

import (
	"github.com/sirupsen/logrus"
)

// logs provide some method to export info in run time.

func init() {
	logrus.Trace("Something very low level.")
}
