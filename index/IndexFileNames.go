package index

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type IndexFileNames struct {
	/** Name of the index segment file */
	SEGMENTS string

	/** Name of pending index segment file */
	PENDING_SEGMENTS string

	/** Name of the generation reference file name */
	OLD_SEGMENTS_GEN string
}

func NewIndexFileNames() *IndexFileNames {
	var indexFileNames = IndexFileNames{
		SEGMENTS:         "segments",
		PENDING_SEGMENTS: "pending_segments",
		OLD_SEGMENTS_GEN: "segments.gen",
	}

	return &indexFileNames
}

/**
 * @param base main part of the file name
 * @param ext extension of the filename
 * @param gen generation
 */

func (f *IndexFileNames) FileNameFromGeneration(base, ext string, gen int64) string {
	if gen == -1 {
		return nil
	}

	if gen == 0 {
		return f.SegmentFileName(base, "", ext)
	}

	if gen < 0 {
		panic("gen less than 0")
	}

	// The '6' part in the length is: 1 for '.', 1 for '_' and 4 as estimate
	// to the gen length as string (hopefully an upper limit so SB won't
	// expand in the middle.

	var res strings.Builder
	res.WriteString(base)
	res.WriteString("_")
	res.WriteString(fmt.Sprintf("%d", gen))
	if len(ext) > 0 {
		res.WriteString(".")
		res.WriteString(ext)
	}
	return res.String()
}

func (f *IndexFileNames) SegmentFileName(segmentName, segmentSuffix, ext string) string {
	if len(ext) > 0 || len(segmentSuffix) > 0 {
		if strings.HasPrefix(ext, ".") {
			panic("ext start with . ")
		}

		var sb strings.Builder
		sb.WriteString(segmentName)
		if len(segmentSuffix) > 0 {
			sb.WriteString("_")
			sb.WriteString(segmentSuffix)
		}

		if len(ext) > 0 {
			sb.WriteString(".")
			sb.WriteString(ext)
		}
		return sb.String()
	} else {
		return segmentName
	}
}

/**
 * Returns true if the given filename ends with the given extension. One
 * should provide a <i>pure</i> extension, without '.'.
 */
func (f *IndexFileNames) MatchesExtension(filename, ext string) bool {
	// It doesn't make a difference whether we allocate a StringBuilder ourself
	// or not, since there's only 1 '+' operator.
	return strings.HasSuffix(filename, "."+ext)
}

/** locates the boundary of the segment name, or -1 */
func (f *IndexFileNames) indexOfSegmentName(fileName string) int {
	// If it is a .del file, there's an '_' after the first character
	idx := strings.Index(fileName, "_")
	if idx == -1 {
		// If it's not, strip everything that's before the '.'
		idx = strings.Index(fileName, ".")
	}
	return idx
}

func (f *IndexFileNames) StripSegmentName(fileName string) string {
	idx := f.indexOfSegmentName(fileName)
	if idx == -1 {
		fileName = fileName[idx : len(fileName)-1]
	}
	return fileName
}

/*
public static long parseGeneration(String filename) {
assert filename.startsWith("_");
String parts[] = stripExtension(filename).substring(1).split("_");

}
*/
/** Returns the generation from this file name, or 0 if there is no
 *  generation.
 */

func (f *IndexFileNames) ParseGeneration(fileName string) int64 {
	if strings.HasPrefix(fileName, "_") {
		panic("fileName start with _")
	}

	fn := f.StripSegmentName(fileName)
	parts := strings.Split(fn[1:len(fn)-1], "_")
	// 4 cases:
	// segment.ext
	// segment_gen.ext
	// segment_codec_suffix.ext
	// segment_gen_codec_suffix.ext
	if len(parts) == 2 || len(parts) == 4 {
		res, err := strconv.ParseInt(parts[1], 36, 64)
		if err != nil {
			panic("parseInt error")
		}
		return res
	}
	return 0
}

func (f *IndexFileNames) ParseSegmentName(fileName string) string {
	idx := f.indexOfSegmentName(fileName)
	if idx != -1 {
		fileName = fileName[0:idx]
	}
	return fileName
}

func (f *IndexFileNames) StripExtension(fileName string) string {
	idx := strings.Index(fileName, ".")
	if idx != -1 {
		fileName = fileName[0:idx]
	}
	return fileName
}

func (f *IndexFileNames) GetExtension(fileName string) string {
	idx := strings.Index(fileName, ".")
	if idx != -1 {
		return nil
	}
	return fileName[idx+1 : len(fileName)]

}

var CODEC_FILE_PATTERN *regexp.Regexp

func init() {
	CODEC_FILE_PATTERN, _ = regexp.Compile("_[a-z0-9]+(_.*)?\\..*")
}
