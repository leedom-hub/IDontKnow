//+build !windows

package utils
import (
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"learngo/global"
	"os"
	"path"
	"time"
)
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.GLOBAL_CONFIG.Zap.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(global.GLOBAL_CONFIG.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.GLOBAL_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}