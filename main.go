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
    index := strings.IndexAny(str, "=")
    rangeType := string([]byte(str)[:index])
    
    if index == -1 {
        return []Range{Range{Type: rangeType, Start: 0, End: -2}}
    }

    // split the range string
    arr := strings.Split(strings.Trim(str, string([]byte(str)[:(index + 1)])), ",")
    var ranges []Range

    // parse all ranges
    for i := 0; i < len(arr); i++ {
        rng := strings.Split(arr[i], "-")
        
        start, err := strconv.ParseFloat(rng[0], 64)
        if err != nil {
            start = math.NaN()
        }

        end, err := strconv.ParseFloat(rng[1], 64)
        if err != nil {
            end = math.NaN()
        }

        // -nnn
        if math.IsNaN(start) == true {
            start = float64(size) - end
            end = float64(size) - 1
        // nnn-
        } else if math.IsNaN(end) == true {
            end = float64(size) - 1
        }

        // limit last-byte-pos to current length
        if end > float64(size) - 1 {
            end = float64(size) - 1
        }

        // invalid or unsatisifiable
        if math.IsNaN(start) || math.IsNaN(end) || start > end || start < 0 {
            continue
        }

        // add range
        ranges = append(ranges, Range{ Type: rangeType, Start: int64(start), End: int64(end) })
    }

    if len(ranges) > 0 {
        return ranges
    }

    return []Range{Range{Type: rangeType, Start: 0, End: -1}}
}
