package RangeParser

import (
    "strings"
    "strconv"
    "math"
)

type Range struct {
    Type string
    Start int64
    End int64
}

func Parse(size int64, str string) []Range {
    rangeType := string([]byte(str)[:index])

    index := strings.IndexAny(str, "=")
    if index == -1 {
        return []Range{Range{Type: rangeType, Start: 0, End: -2}}
    }

    // split the range string
    arr := strings.Split(strings.Trim(str, string([]byte(str)[:(index + 1)])), ",")
    var ranges []Range

    // parse all ranges
    for i := 0; i < len(arr); i++ {
        rng := strings.Split(arr[i], "-")
        start, _ := strconv.ParseInt(rng[0], 10, 64)
        end, _ := strconv.ParseInt(rng[1], 10, 64)

        // -nnn
        if math.IsNaN(float64(start)) == true {
            start = size - end
            end = size - 1
        // nnn-
        } else if math.IsNaN(float64(end)) == true {
            end = size - 1
        }

        // limit last-byte-pos to current length
        if end > size - 1 {
            end = size - 1
        }

        // invalid or unsatisifiable
        if math.IsNaN(float64(start)) || math.IsNaN(float64(end)) || start > end || start < 0 {
            continue
        }

        // add range
        ranges = append(ranges, Range{Type: rangeType, Start: start, End: end})
    }

    if len(ranges) > 0 {
        return ranges
    }

    return []Range{Range{Type: rangeType, Start: 0, End: -1}}
}