// Code generated by qtc from "util.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line util.html:1
package templates

//line util.html:1
import "strconv"

//line util.html:2
import "fmt"

//line util.html:3
import "github.com/bakape/meguca/common"

//line util.html:4
import "github.com/bakape/meguca/assets"

//line util.html:5
import "github.com/bakape/meguca/lang"

// Renders the tab selection butts in tabbed windows

//line util.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line util.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line util.html:8
func streamtabButts(qw422016 *qt422016.Writer, names []string) {
//line util.html:8
	qw422016.N().S(`<div class="tab-butts">`)
//line util.html:10
	for i, n := range names {
//line util.html:10
		qw422016.N().S(`<a class="tab-link`)
//line util.html:11
		if i == 0 {
//line util.html:11
			qw422016.N().S(` `)
//line util.html:11
			qw422016.N().S(`tab-sel`)
//line util.html:11
		}
//line util.html:11
		qw422016.N().S(`" data-id="`)
//line util.html:11
		qw422016.N().D(i)
//line util.html:11
		qw422016.N().S(`">`)
//line util.html:12
		qw422016.N().S(n)
//line util.html:12
		qw422016.N().S(`</a>`)
//line util.html:14
	}
//line util.html:14
	qw422016.N().S(`</div><hr>`)
//line util.html:17
}

//line util.html:17
func writetabButts(qq422016 qtio422016.Writer, names []string) {
//line util.html:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:17
	streamtabButts(qw422016, names)
//line util.html:17
	qt422016.ReleaseWriter(qw422016)
//line util.html:17
}

//line util.html:17
func tabButts(names []string) string {
//line util.html:17
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:17
	writetabButts(qb422016, names)
//line util.html:17
	qs422016 := string(qb422016.B)
//line util.html:17
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:17
	return qs422016
//line util.html:17
}

// Render a link to another post. Can optionally be cross-thread.

//line util.html:20
func streampostLink(qw422016 *qt422016.Writer, link common.Link, cross, boardPage bool) {
//line util.html:21
	idBuf := strconv.AppendUint(make([]byte, 0, 16), link.ID, 10)

//line util.html:22
	url := make([]byte, 0, 64)

//line util.html:23
	if cross {
//line util.html:24
		url = append(url, '/')

//line util.html:25
		url = append(url, link.Board...)

//line util.html:26
		url = append(url, '/')

//line util.html:27
		url = strconv.AppendUint(url, link.OP, 10)

//line util.html:28
	}
//line util.html:29
	url = append(url, "#p"...)

//line util.html:30
	url = append(url, idBuf...)

//line util.html:30
	qw422016.N().S(`<a class="post-link" data-id="`)
//line util.html:31
	qw422016.N().Z(idBuf)
//line util.html:31
	qw422016.N().S(`" href="`)
//line util.html:31
	qw422016.N().Z(url)
//line util.html:31
	qw422016.N().S(`">>>`)
//line util.html:33
	qw422016.N().Z(idBuf)
//line util.html:34
	if cross && !boardPage {
//line util.html:35
		qw422016.N().S(` `)
//line util.html:35
		qw422016.N().S(`➡`)
//line util.html:36
	}
//line util.html:36
	qw422016.N().S(`</a><a class="hash-link" href="`)
//line util.html:38
	qw422016.N().Z(url)
//line util.html:38
	qw422016.N().S(`"> #</a>`)
//line util.html:39
}

//line util.html:39
func writepostLink(qq422016 qtio422016.Writer, link common.Link, cross, boardPage bool) {
//line util.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:39
	streampostLink(qw422016, link, cross, boardPage)
//line util.html:39
	qt422016.ReleaseWriter(qw422016)
//line util.html:39
}

//line util.html:39
func postLink(link common.Link, cross, boardPage bool) string {
//line util.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:39
	writepostLink(qb422016, link, cross, boardPage)
//line util.html:39
	qs422016 := string(qb422016.B)
//line util.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:39
	return qs422016
//line util.html:39
}

//line util.html:41
func streamexpandLink(qw422016 *qt422016.Writer, board, id string) {
//line util.html:41
	qw422016.N().S(`<span class="act"><a href="/`)
//line util.html:43
	qw422016.N().S(board)
//line util.html:43
	qw422016.N().S(`/`)
//line util.html:43
	qw422016.N().S(id)
//line util.html:43
	qw422016.N().S(`">`)
//line util.html:44
	qw422016.N().S(lang.Get().Common.Posts["expand"])
//line util.html:44
	qw422016.N().S(`</a></span>`)
//line util.html:47
}

//line util.html:47
func writeexpandLink(qq422016 qtio422016.Writer, board, id string) {
//line util.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:47
	streamexpandLink(qw422016, board, id)
//line util.html:47
	qt422016.ReleaseWriter(qw422016)
//line util.html:47
}

//line util.html:47
func expandLink(board, id string) string {
//line util.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:47
	writeexpandLink(qb422016, board, id)
//line util.html:47
	qs422016 := string(qb422016.B)
//line util.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:47
	return qs422016
//line util.html:47
}

//line util.html:49
func streamlast100Link(qw422016 *qt422016.Writer, board, id string) {
//line util.html:49
	qw422016.N().S(`<span class="act"><a href="/`)
//line util.html:51
	qw422016.N().S(board)
//line util.html:51
	qw422016.N().S(`/`)
//line util.html:51
	qw422016.N().S(id)
//line util.html:51
	qw422016.N().S(`?last=100#bottom">`)
//line util.html:52
	qw422016.N().S(lang.Get().Common.UI["last"])
//line util.html:52
	qw422016.N().S(` `)
//line util.html:52
	qw422016.N().S(`100</a></span>`)
//line util.html:55
}

//line util.html:55
func writelast100Link(qq422016 qtio422016.Writer, board, id string) {
//line util.html:55
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:55
	streamlast100Link(qw422016, board, id)
//line util.html:55
	qt422016.ReleaseWriter(qw422016)
//line util.html:55
}

//line util.html:55
func last100Link(board, id string) string {
//line util.html:55
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:55
	writelast100Link(qb422016, board, id)
//line util.html:55
	qs422016 := string(qb422016.B)
//line util.html:55
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:55
	return qs422016
//line util.html:55
}

// Render the class attribute of a post

//line util.html:58
func streampostClass(qw422016 *qt422016.Writer, p common.Post, op uint64) {
//line util.html:58
	qw422016.N().S(`class="glass`)
//line util.html:60
	if p.Editing {
//line util.html:61
		qw422016.N().S(` `)
//line util.html:61
		qw422016.N().S(`editing`)
//line util.html:62
	}
//line util.html:63
	if p.IsDeleted() {
//line util.html:64
		qw422016.N().S(` `)
//line util.html:64
		qw422016.N().S(`deleted hidden`)
//line util.html:65
	}
//line util.html:66
	if p.Image != nil {
//line util.html:67
		qw422016.N().S(` `)
//line util.html:67
		qw422016.N().S(`media`)
//line util.html:68
	}
//line util.html:69
	if p.ID == op {
//line util.html:70
		qw422016.N().S(` `)
//line util.html:70
		qw422016.N().S(`op`)
//line util.html:71
	}
//line util.html:71
	qw422016.N().S(`"`)
//line util.html:73
}

//line util.html:73
func writepostClass(qq422016 qtio422016.Writer, p common.Post, op uint64) {
//line util.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:73
	streampostClass(qw422016, p, op)
//line util.html:73
	qt422016.ReleaseWriter(qw422016)
//line util.html:73
}

//line util.html:73
func postClass(p common.Post, op uint64) string {
//line util.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:73
	writepostClass(qb422016, p, op)
//line util.html:73
	qs422016 := string(qb422016.B)
//line util.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:73
	return qs422016
//line util.html:73
}

// Renders a stylized deleted post display toggle

//line util.html:76
func streamdeletedToggle(qw422016 *qt422016.Writer) {
//line util.html:76
	qw422016.N().S(`<input type="checkbox" class="deleted-toggle">`)
//line util.html:78
}

//line util.html:78
func writedeletedToggle(qq422016 qtio422016.Writer) {
//line util.html:78
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:78
	streamdeletedToggle(qw422016)
//line util.html:78
	qt422016.ReleaseWriter(qw422016)
//line util.html:78
}

//line util.html:78
func deletedToggle() string {
//line util.html:78
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:78
	writedeletedToggle(qb422016)
//line util.html:78
	qs422016 := string(qb422016.B)
//line util.html:78
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:78
	return qs422016
//line util.html:78
}

// Notice widget, that reveals text on hover

//line util.html:82
func streamhoverReveal(qw422016 *qt422016.Writer, tag, text, label string) {
//line util.html:83
	if text == "" {
//line util.html:84
		return
//line util.html:85
	}
//line util.html:85
	qw422016.N().S(`<`)
//line util.html:86
	qw422016.N().S(tag)
//line util.html:86
	qw422016.N().S(` `)
//line util.html:86
	qw422016.N().S(`class="hover-reveal`)
//line util.html:86
	if tag == "aside" {
//line util.html:86
		qw422016.N().S(` `)
//line util.html:86
		qw422016.N().S(`glass`)
//line util.html:86
	}
//line util.html:86
	qw422016.N().S(`"><span class="act">`)
//line util.html:88
	qw422016.N().S(label)
//line util.html:88
	qw422016.N().S(`</span><span class="popup-menu glass">`)
//line util.html:91
	qw422016.E().S(text)
//line util.html:91
	qw422016.N().S(`</span></`)
//line util.html:93
	qw422016.N().S(tag)
//line util.html:93
	qw422016.N().S(`>`)
//line util.html:94
}

//line util.html:94
func writehoverReveal(qq422016 qtio422016.Writer, tag, text, label string) {
//line util.html:94
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:94
	streamhoverReveal(qw422016, tag, text, label)
//line util.html:94
	qt422016.ReleaseWriter(qw422016)
//line util.html:94
}

//line util.html:94
func hoverReveal(tag, text, label string) string {
//line util.html:94
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:94
	writehoverReveal(qb422016, tag, text, label)
//line util.html:94
	qs422016 := string(qb422016.B)
//line util.html:94
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:94
	return qs422016
//line util.html:94
}

// Render pin signifying a thread is sticky

//line util.html:97
func streamrenderSticky(qw422016 *qt422016.Writer, sticky bool) {
//line util.html:98
	if !sticky {
//line util.html:99
		return
//line util.html:100
	}
//line util.html:100
	qw422016.N().S(`<svg class="sticky" xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M1.34 0a.5.5 0 0 0 .16 1h.5v2h-1c-.55 0-1 .45-1 1h3v3l.44 1 .56-1v-3h3c0-.55-.45-1-1-1h-1v-2h.5a.5.5 0 1 0 0-1h-4a.5.5 0 0 0-.09 0 .5.5 0 0 0-.06 0z" /></svg>`)
//line util.html:104
}

//line util.html:104
func writerenderSticky(qq422016 qtio422016.Writer, sticky bool) {
//line util.html:104
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:104
	streamrenderSticky(qw422016, sticky)
//line util.html:104
	qt422016.ReleaseWriter(qw422016)
//line util.html:104
}

//line util.html:104
func renderSticky(sticky bool) string {
//line util.html:104
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:104
	writerenderSticky(qb422016, sticky)
//line util.html:104
	qs422016 := string(qb422016.B)
//line util.html:104
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:104
	return qs422016
//line util.html:104
}

// Render lock signifying a thread has posting disabled

//line util.html:107
func streamrenderLocked(qw422016 *qt422016.Writer, locked bool) {
//line util.html:108
	if !locked {
//line util.html:109
		return
//line util.html:110
	}
//line util.html:110
	qw422016.N().S(`<svg class="locked" xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M3 0c-1.1 0-2 .9-2 2v1h-1v4h6v-4h-1v-1c0-1.1-.9-2-2-2zm0 1c.56 0 1 .44 1 1v1h-2v-1c0-.56.44-1 1-1z" transform="translate(1)" /></svg>`)
//line util.html:114
}

//line util.html:114
func writerenderLocked(qq422016 qtio422016.Writer, locked bool) {
//line util.html:114
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:114
	streamrenderLocked(qw422016, locked)
//line util.html:114
	qt422016.ReleaseWriter(qw422016)
//line util.html:114
}

//line util.html:114
func renderLocked(locked bool) string {
//line util.html:114
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:114
	writerenderLocked(qb422016, locked)
//line util.html:114
	qs422016 := string(qb422016.B)
//line util.html:114
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:114
	return qs422016
//line util.html:114
}

// Render an image or video asset

//line util.html:117
func streamasset(qw422016 *qt422016.Writer, url, mime string) {
//line util.html:118
	if mime == "video/webm" {
//line util.html:118
		qw422016.N().S(`<video src="`)
//line util.html:119
		qw422016.N().S(url)
//line util.html:119
		qw422016.N().S(`" autoplay loop>`)
//line util.html:120
	} else {
//line util.html:120
		qw422016.N().S(`<img src="`)
//line util.html:121
		qw422016.N().S(url)
//line util.html:121
		qw422016.N().S(`">`)
//line util.html:122
	}
//line util.html:123
}

//line util.html:123
func writeasset(qq422016 qtio422016.Writer, url, mime string) {
//line util.html:123
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:123
	streamasset(qw422016, url, mime)
//line util.html:123
	qt422016.ReleaseWriter(qw422016)
//line util.html:123
}

//line util.html:123
func asset(url, mime string) string {
//line util.html:123
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:123
	writeasset(qb422016, url, mime)
//line util.html:123
	qs422016 := string(qb422016.B)
//line util.html:123
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:123
	return qs422016
//line util.html:123
}

// Render Banners NFT Banner

//line util.html:126
func streambanners_nft(qw422016 *qt422016.Writer, id string) {
//line util.html:126
	qw422016.N().S(`<a href="https://blur.io/eth/asset/0x1352149cd78d686043b504e7e7d96c5946b0c39c/`)
//line util.html:127
	qw422016.N().S(id)
//line util.html:127
	qw422016.N().S(`"><img src="https://miladymaker.net/banners/nft/`)
//line util.html:128
	qw422016.N().S(id)
//line util.html:128
	qw422016.N().S(`.png"></a>`)
//line util.html:130
}

//line util.html:130
func writebanners_nft(qq422016 qtio422016.Writer, id string) {
//line util.html:130
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:130
	streambanners_nft(qw422016, id)
//line util.html:130
	qt422016.ReleaseWriter(qw422016)
//line util.html:130
}

//line util.html:130
func banners_nft(id string) string {
//line util.html:130
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:130
	writebanners_nft(qb422016, id)
//line util.html:130
	qs422016 := string(qb422016.B)
//line util.html:130
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:130
	return qs422016
//line util.html:130
}

//line util.html:132
func streamloadingImage(qw422016 *qt422016.Writer, board string) {
//line util.html:132
	qw422016.N().S(`<div id="loading-image" class="noscript-hide">`)
//line util.html:134
	streamasset(qw422016, fmt.Sprintf("/assets/loading/%s", board), assets.Loading.Get(board).Mime)
//line util.html:134
	qw422016.N().S(`</div>`)
//line util.html:136
}

//line util.html:136
func writeloadingImage(qq422016 qtio422016.Writer, board string) {
//line util.html:136
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:136
	streamloadingImage(qw422016, board)
//line util.html:136
	qt422016.ReleaseWriter(qw422016)
//line util.html:136
}

//line util.html:136
func loadingImage(board string) string {
//line util.html:136
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:136
	writeloadingImage(qb422016, board)
//line util.html:136
	qs422016 := string(qb422016.B)
//line util.html:136
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:136
	return qs422016
//line util.html:136
}

// Render localized table headers by UI translation ID

//line util.html:139
func streamtableHeaders(qw422016 *qt422016.Writer, ids ...string) {
//line util.html:140
	ln := lang.Get()

//line util.html:140
	qw422016.N().S(`<tr>`)
//line util.html:142
	for _, id := range ids {
//line util.html:143
		label := ln.UI[id]

//line util.html:144
		if label == "" {
//line util.html:145
			label = ln.Common.UI[id]

//line util.html:146
		}
//line util.html:146
		qw422016.N().S(`<th>`)
//line util.html:147
		qw422016.N().S(label)
//line util.html:147
		qw422016.N().S(`</th>`)
//line util.html:148
	}
//line util.html:148
	qw422016.N().S(`</tr>`)
//line util.html:150
}

//line util.html:150
func writetableHeaders(qq422016 qtio422016.Writer, ids ...string) {
//line util.html:150
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:150
	streamtableHeaders(qw422016, ids...)
//line util.html:150
	qt422016.ReleaseWriter(qw422016)
//line util.html:150
}

//line util.html:150
func tableHeaders(ids ...string) string {
//line util.html:150
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:150
	writetableHeaders(qb422016, ids...)
//line util.html:150
	qs422016 := string(qb422016.B)
//line util.html:150
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:150
	return qs422016
//line util.html:150
}

//line util.html:152
func streamthreadWatcherToggle(qw422016 *qt422016.Writer, id uint64) {
//line util.html:152
	qw422016.N().S(`<a class="watcher-toggle svg-link noscript-hide" title="`)
//line util.html:153
	qw422016.N().S(lang.Get().Common.UI["watchThread"])
//line util.html:153
	qw422016.N().S(`" data-id="`)
//line util.html:153
	qw422016.N().S(strconv.FormatUint(id, 10))
//line util.html:153
	qw422016.N().S(`"><svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M4.03 0c-2.53 0-4.03 3-4.03 3s1.5 3 4.03 3c2.47 0 3.97-3 3.97-3s-1.5-3-3.97-3zm-.03 1c1.11 0 2 .9 2 2 0 1.11-.89 2-2 2-1.1 0-2-.89-2-2 0-1.1.9-2 2-2zm0 1c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1c0-.1-.04-.19-.06-.28-.08.16-.24.28-.44.28-.28 0-.5-.22-.5-.5 0-.2.12-.36.28-.44-.09-.03-.18-.06-.28-.06z" transform="translate(0 1)" /></svg></a>`)
//line util.html:158
}

//line util.html:158
func writethreadWatcherToggle(qq422016 qtio422016.Writer, id uint64) {
//line util.html:158
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:158
	streamthreadWatcherToggle(qw422016, id)
//line util.html:158
	qt422016.ReleaseWriter(qw422016)
//line util.html:158
}

//line util.html:158
func threadWatcherToggle(id uint64) string {
//line util.html:158
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:158
	writethreadWatcherToggle(qb422016, id)
//line util.html:158
	qs422016 := string(qb422016.B)
//line util.html:158
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:158
	return qs422016
//line util.html:158
}

//line util.html:160
func streamcontrolLink(qw422016 *qt422016.Writer) {
//line util.html:160
	qw422016.N().S(`<a class="control svg-link noscript-hide"><svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M1.5 0l-1.5 1.5 4 4 4-4-1.5-1.5-2.5 2.5-2.5-2.5z" transform="translate(0 1)" /></svg></a>`)
//line util.html:166
}

//line util.html:166
func writecontrolLink(qq422016 qtio422016.Writer) {
//line util.html:166
	qw422016 := qt422016.AcquireWriter(qq422016)
//line util.html:166
	streamcontrolLink(qw422016)
//line util.html:166
	qt422016.ReleaseWriter(qw422016)
//line util.html:166
}

//line util.html:166
func controlLink() string {
//line util.html:166
	qb422016 := qt422016.AcquireByteBuffer()
//line util.html:166
	writecontrolLink(qb422016)
//line util.html:166
	qs422016 := string(qb422016.B)
//line util.html:166
	qt422016.ReleaseByteBuffer(qb422016)
//line util.html:166
	return qs422016
//line util.html:166
}
