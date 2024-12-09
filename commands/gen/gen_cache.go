package gen

import (
	"strconv"
	"strings"

	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf-cli/library/utils"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
)

func doGenCacheServer() {
	parser, err := gcmd.Parse(g.MapStrBool{
		"t,table":       true,
		"pk,pk":         true, //主键信息
		"prefix,prefix": true, //缓存前缀
		"ttl,ttl":       true, //缓存时长
	})
	if err != nil {
		mlog.Fatal(err)
	}
	newTableName := getOptionOrConfigForDao(-1, parser, "table")
	pkName := getOptionOrConfigForDao(-1, parser, "pk")
	prefixName := getOptionOrConfigForDao(-1, parser, "prefix")
	ttlName := getOptionOrConfigForDao(-1, parser, "ttl")
	newTableName = strings.TrimSpace(newTableName)
	pkName = strings.TrimSpace(pkName)
	prefixName = strings.TrimSpace(prefixName)
	ttlName = strings.TrimSpace(ttlName)
	if len(newTableName) == 0 || len(pkName) == 0 || len(prefixName) == 0 || len(ttlName) == 0 {
		mlog.Fatalf("table or pk or prefix name is empty")
		return
	}
	ttlInt, err := strconv.Atoi(ttlName)
	if err != nil {
		panic(err)
	}
	tableNameCaseCamel := gstr.CaseCamel(newTableName)

	filepath := "rpc/server/internal/cache/gen_" + newTableName + ".go"
	entityContent := gstr.ReplaceByMap(templateCacheServerContent, g.MapStrStr{
		"{TplName}":     tableNameCaseCamel,
		"{TplTable}":    newTableName,
		"{TplPk}":       pkName,
		"{TplStructPk}": gstr.CaseCamel(pkName),
		"{TplPrefix}":   prefixName,
		"{TplTtl}":      strconv.Itoa(ttlInt),
	})
	if err := gfile.PutContents(filepath, strings.TrimSpace(entityContent)); err != nil {
		mlog.Fatalf("writing content to %s failed: %v", filepath, err)
	} else {
		utils.GoFmt(filepath)
		mlog.Print("generated:", filepath)
	}

	filepath = "app/service/codec/gen_codec_" + newTableName + ".go"
	entityContent = gstr.ReplaceByMap(templateCacheCmdContent, g.MapStrStr{
		"{TplName}":     tableNameCaseCamel,
		"{TplTable}":    newTableName,
		"{TplPk}":       pkName,
		"{TplStructPk}": gstr.CaseCamel(pkName),
		"{TplPrefix}":   prefixName,
		"{TplTtl}":      strconv.Itoa(ttlInt),
	})
	if err = gfile.PutContents(filepath, strings.TrimSpace(entityContent)); err != nil {
		mlog.Fatalf("writing content to %s failed: %v", filepath, err)
	} else {
		utils.GoFmt(filepath)
		mlog.Print("generated:", filepath)
	}

	filepath = "cmd/internal/cache_process_" + newTableName + ".go"
	if !gfile.Exists(filepath) {
		entityContent = gstr.ReplaceByMap(templateCachePluginContent, g.MapStrStr{
			"{TplName}":     tableNameCaseCamel,
			"{TplTable}":    newTableName,
			"{TplPk}":       pkName,
			"{TplStructPk}": gstr.CaseCamel(pkName),
			"{TplPrefix}":   prefixName,
			"{TplTtl}":      strconv.Itoa(ttlInt),
		})
		if err = gfile.PutContents(filepath, strings.TrimSpace(entityContent)); err != nil {
			mlog.Fatalf("writing content to %s failed: %v", filepath, err)
		} else {
			utils.GoFmt(filepath)
			mlog.Print("generated:", filepath)
		}
	}

	filepath = "rpc/client/gen_base_cache_" + newTableName + ".go"
	entityContent = gstr.ReplaceByMap(templateCacheClientContent, g.MapStrStr{
		"{TplName}":     tableNameCaseCamel,
		"{TplTable}":    newTableName,
		"{TplPk}":       pkName,
		"{TplStructPk}": gstr.CaseCamel(pkName),
		"{TplPrefix}":   prefixName,
		"{TplTtl}":      strconv.Itoa(ttlInt),
	})
	if err = gfile.PutContents(filepath, strings.TrimSpace(entityContent)); err != nil {
		mlog.Fatalf("writing content to %s failed: %v", filepath, err)
	} else {
		utils.GoFmt(filepath)
		mlog.Print("generated:", filepath)
	}

	filepath = "proto/gen_base_cache_rep_" + newTableName + ".proto"
	entityContent = gstr.ReplaceByMap(templateCacheProtoContent, g.MapStrStr{
		"{TplName}":     tableNameCaseCamel,
		"{TplTable}":    newTableName,
		"{TplPk}":       pkName,
		"{TplStructPk}": gstr.CaseCamel(pkName),
		"{TplPrefix}":   prefixName,
		"{TplTtl}":      strconv.Itoa(ttlInt),
	})
	if err = gfile.PutContents(filepath, strings.TrimSpace(entityContent)); err != nil {
		mlog.Fatalf("writing content to %s failed: %v", filepath, err)
	} else {
		utils.GoFmt(filepath)
		mlog.Print("generated:", filepath)
	}
}
