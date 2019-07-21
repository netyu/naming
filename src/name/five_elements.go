/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2019 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file five_elements.go
 * @package name
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 06/06/2019
 */

package name

import (
	"texts"
	"utils"
)

// GanzhiSound : Get sound five-elements by given gan-zhi
func GanzhiSound(ganzhi utils.GanzhiPair) int {
	return ganzhi.Value()
}

// GanzhiSoundAlias : Get sound five-elements id / name / description by given gan-zhi
func GanzhiSoundAlias(ganzhi utils.GanzhiPair, language int) (int, string, string) {
	id := ganzhi.Value()
	if id >= 0 && id < 60 {
		return id,
			texts.GetAlias(texts.AliasSoundFiveElement, id/2, language),
			texts.GetMessage(texts.MessageSoundFiveElementDescription, id, language)
	}

	return 0, "", ""
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
