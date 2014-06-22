package domino

func Head(args ...interface{}) *DomNode {
	return MakeDomNode("head", args...)
}

func (n *DomNode) Head(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("head", args...))
	return n
}

func Title(args ...interface{}) *DomNode {
	return MakeDomNode("title", args...)
}

func (n *DomNode) Title(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("title", args...))
	return n
}

func Base(args ...interface{}) *DomNode {
	return MakeDomNode("base", args...)
}

func (n *DomNode) Base(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("base", args...))
	return n
}

func Link(args ...interface{}) *DomNode {
	return MakeDomNode("link", args...)
}

func (n *DomNode) Link(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("link", args...))
	return n
}

func Meta(args ...interface{}) *DomNode {
	return MakeDomNode("meta", args...)
}

func (n *DomNode) Meta(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("meta", args...))
	return n
}

func Style(args ...interface{}) *DomNode {
	return MakeDomNode("style", args...)
}

func (n *DomNode) Style(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("style", args...))
	return n
}

func Script(args ...interface{}) *DomNode {
	return MakeDomNode("script", args...)
}

func (n *DomNode) Script(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("script", args...))
	return n
}

func Noscript(args ...interface{}) *DomNode {
	return MakeDomNode("noscript", args...)
}

func (n *DomNode) Noscript(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("noscript", args...))
	return n
}

func Body(args ...interface{}) *DomNode {
	return MakeDomNode("body", args...)
}

func (n *DomNode) Body(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("body", args...))
	return n
}

func Section(args ...interface{}) *DomNode {
	return MakeDomNode("section", args...)
}

func (n *DomNode) Section(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("section", args...))
	return n
}

func Nav(args ...interface{}) *DomNode {
	return MakeDomNode("nav", args...)
}

func (n *DomNode) Nav(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("nav", args...))
	return n
}

func Article(args ...interface{}) *DomNode {
	return MakeDomNode("article", args...)
}

func (n *DomNode) Article(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("article", args...))
	return n
}

func Aside(args ...interface{}) *DomNode {
	return MakeDomNode("aside", args...)
}

func (n *DomNode) Aside(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("aside", args...))
	return n
}

func H1(args ...interface{}) *DomNode {
	return MakeDomNode("h1", args...)
}

func (n *DomNode) H1(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("h1", args...))
	return n
}

func H2(args ...interface{}) *DomNode {
	return MakeDomNode("h2", args...)
}

func (n *DomNode) H2(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("h2", args...))
	return n
}

func H3(args ...interface{}) *DomNode {
	return MakeDomNode("h3", args...)
}

func (n *DomNode) H3(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("h3", args...))
	return n
}

func H4(args ...interface{}) *DomNode {
	return MakeDomNode("h4", args...)
}

func (n *DomNode) H4(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("h4", args...))
	return n
}

func H5(args ...interface{}) *DomNode {
	return MakeDomNode("h5", args...)
}

func (n *DomNode) H5(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("h5", args...))
	return n
}

func H6(args ...interface{}) *DomNode {
	return MakeDomNode("h6", args...)
}

func (n *DomNode) H6(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("h6", args...))
	return n
}

func Hgroup(args ...interface{}) *DomNode {
	return MakeDomNode("hgroup", args...)
}

func (n *DomNode) Hgroup(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("hgroup", args...))
	return n
}

func Header(args ...interface{}) *DomNode {
	return MakeDomNode("header", args...)
}

func (n *DomNode) Header(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("header", args...))
	return n
}

func Footer(args ...interface{}) *DomNode {
	return MakeDomNode("footer", args...)
}

func (n *DomNode) Footer(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("footer", args...))
	return n
}

func Address(args ...interface{}) *DomNode {
	return MakeDomNode("address", args...)
}

func (n *DomNode) Address(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("address", args...))
	return n
}

func P(args ...interface{}) *DomNode {
	return MakeDomNode("p", args...)
}

func (n *DomNode) P(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("p", args...))
	return n
}

func Hr(args ...interface{}) *DomNode {
	return MakeDomNode("hr", args...)
}

func (n *DomNode) Hr(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("hr", args...))
	return n
}

func Pre(args ...interface{}) *DomNode {
	return MakeDomNode("pre", args...)
}

func (n *DomNode) Pre(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("pre", args...))
	return n
}

func Blockquote(args ...interface{}) *DomNode {
	return MakeDomNode("blockquote", args...)
}

func (n *DomNode) Blockquote(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("blockquote", args...))
	return n
}

func Ol(args ...interface{}) *DomNode {
	return MakeDomNode("ol", args...)
}

func (n *DomNode) Ol(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("ol", args...))
	return n
}

func Ul(args ...interface{}) *DomNode {
	return MakeDomNode("ul", args...)
}

func (n *DomNode) Ul(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("ul", args...))
	return n
}

func Li(args ...interface{}) *DomNode {
	return MakeDomNode("li", args...)
}

func (n *DomNode) Li(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("li", args...))
	return n
}

func Dl(args ...interface{}) *DomNode {
	return MakeDomNode("dl", args...)
}

func (n *DomNode) Dl(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("dl", args...))
	return n
}

func Dt(args ...interface{}) *DomNode {
	return MakeDomNode("dt", args...)
}

func (n *DomNode) Dt(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("dt", args...))
	return n
}

func Dd(args ...interface{}) *DomNode {
	return MakeDomNode("dd", args...)
}

func (n *DomNode) Dd(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("dd", args...))
	return n
}

func Figure(args ...interface{}) *DomNode {
	return MakeDomNode("figure", args...)
}

func (n *DomNode) Figure(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("figure", args...))
	return n
}

func Figcaption(args ...interface{}) *DomNode {
	return MakeDomNode("figcaption", args...)
}

func (n *DomNode) Figcaption(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("figcaption", args...))
	return n
}

func Div(args ...interface{}) *DomNode {
	return MakeDomNode("div", args...)
}

func (n *DomNode) Div(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("div", args...))
	return n
}

func A(args ...interface{}) *DomNode {
	return MakeDomNode("a", args...)
}

func (n *DomNode) A(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("a", args...))
	return n
}

func Em(args ...interface{}) *DomNode {
	return MakeDomNode("em", args...)
}

func (n *DomNode) Em(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("em", args...))
	return n
}

func Strong(args ...interface{}) *DomNode {
	return MakeDomNode("strong", args...)
}

func (n *DomNode) Strong(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("strong", args...))
	return n
}

func Small(args ...interface{}) *DomNode {
	return MakeDomNode("small", args...)
}

func (n *DomNode) Small(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("small", args...))
	return n
}

func S(args ...interface{}) *DomNode {
	return MakeDomNode("s", args...)
}

func (n *DomNode) S(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("s", args...))
	return n
}

func Cite(args ...interface{}) *DomNode {
	return MakeDomNode("cite", args...)
}

func (n *DomNode) Cite(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("cite", args...))
	return n
}

func Q(args ...interface{}) *DomNode {
	return MakeDomNode("q", args...)
}

func (n *DomNode) Q(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("q", args...))
	return n
}

func Dfn(args ...interface{}) *DomNode {
	return MakeDomNode("dfn", args...)
}

func (n *DomNode) Dfn(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("dfn", args...))
	return n
}

func Abbr(args ...interface{}) *DomNode {
	return MakeDomNode("abbr", args...)
}

func (n *DomNode) Abbr(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("abbr", args...))
	return n
}

func Time(args ...interface{}) *DomNode {
	return MakeDomNode("time", args...)
}

func (n *DomNode) Time(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("time", args...))
	return n
}

func Code(args ...interface{}) *DomNode {
	return MakeDomNode("code", args...)
}

func (n *DomNode) Code(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("code", args...))
	return n
}

func Var(args ...interface{}) *DomNode {
	return MakeDomNode("var", args...)
}

func (n *DomNode) Var(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("var", args...))
	return n
}

func Samp(args ...interface{}) *DomNode {
	return MakeDomNode("samp", args...)
}

func (n *DomNode) Samp(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("samp", args...))
	return n
}

func Kbd(args ...interface{}) *DomNode {
	return MakeDomNode("kbd", args...)
}

func (n *DomNode) Kbd(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("kbd", args...))
	return n
}

func Sub(args ...interface{}) *DomNode {
	return MakeDomNode("sub", args...)
}

func (n *DomNode) Sub(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("sub", args...))
	return n
}

func Sup(args ...interface{}) *DomNode {
	return MakeDomNode("sup", args...)
}

func (n *DomNode) Sup(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("sup", args...))
	return n
}

func I(args ...interface{}) *DomNode {
	return MakeDomNode("i", args...)
}

func (n *DomNode) I(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("i", args...))
	return n
}

func B(args ...interface{}) *DomNode {
	return MakeDomNode("b", args...)
}

func (n *DomNode) B(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("b", args...))
	return n
}

func U(args ...interface{}) *DomNode {
	return MakeDomNode("u", args...)
}

func (n *DomNode) U(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("u", args...))
	return n
}

func Mark(args ...interface{}) *DomNode {
	return MakeDomNode("mark", args...)
}

func (n *DomNode) Mark(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("mark", args...))
	return n
}

func Ruby(args ...interface{}) *DomNode {
	return MakeDomNode("ruby", args...)
}

func (n *DomNode) Ruby(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("ruby", args...))
	return n
}

func Rt(args ...interface{}) *DomNode {
	return MakeDomNode("rt", args...)
}

func (n *DomNode) Rt(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("rt", args...))
	return n
}

func Rp(args ...interface{}) *DomNode {
	return MakeDomNode("rp", args...)
}

func (n *DomNode) Rp(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("rp", args...))
	return n
}

func Bdi(args ...interface{}) *DomNode {
	return MakeDomNode("bdi", args...)
}

func (n *DomNode) Bdi(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("bdi", args...))
	return n
}

func Bdo(args ...interface{}) *DomNode {
	return MakeDomNode("bdo", args...)
}

func (n *DomNode) Bdo(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("bdo", args...))
	return n
}

func Span(args ...interface{}) *DomNode {
	return MakeDomNode("span", args...)
}

func (n *DomNode) Span(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("span", args...))
	return n
}

func Br(args ...interface{}) *DomNode {
	return MakeDomNode("br", args...)
}

func (n *DomNode) Br(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("br", args...))
	return n
}

func Wbr(args ...interface{}) *DomNode {
	return MakeDomNode("wbr", args...)
}

func (n *DomNode) Wbr(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("wbr", args...))
	return n
}

func Ins(args ...interface{}) *DomNode {
	return MakeDomNode("ins", args...)
}

func (n *DomNode) Ins(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("ins", args...))
	return n
}

func Del(args ...interface{}) *DomNode {
	return MakeDomNode("del", args...)
}

func (n *DomNode) Del(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("del", args...))
	return n
}

func Img(args ...interface{}) *DomNode {
	return MakeDomNode("img", args...)
}

func (n *DomNode) Img(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("img", args...))
	return n
}

func Iframe(args ...interface{}) *DomNode {
	return MakeDomNode("iframe", args...)
}

func (n *DomNode) Iframe(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("iframe", args...))
	return n
}

func Embed(args ...interface{}) *DomNode {
	return MakeDomNode("embed", args...)
}

func (n *DomNode) Embed(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("embed", args...))
	return n
}

func Object(args ...interface{}) *DomNode {
	return MakeDomNode("object", args...)
}

func (n *DomNode) Object(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("object", args...))
	return n
}

func Param(args ...interface{}) *DomNode {
	return MakeDomNode("param", args...)
}

func (n *DomNode) Param(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("param", args...))
	return n
}

func Video(args ...interface{}) *DomNode {
	return MakeDomNode("video", args...)
}

func (n *DomNode) Video(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("video", args...))
	return n
}

func Audio(args ...interface{}) *DomNode {
	return MakeDomNode("audio", args...)
}

func (n *DomNode) Audio(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("audio", args...))
	return n
}

func Source(args ...interface{}) *DomNode {
	return MakeDomNode("source", args...)
}

func (n *DomNode) Source(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("source", args...))
	return n
}

func Track(args ...interface{}) *DomNode {
	return MakeDomNode("track", args...)
}

func (n *DomNode) Track(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("track", args...))
	return n
}

func Canvas(args ...interface{}) *DomNode {
	return MakeDomNode("canvas", args...)
}

func (n *DomNode) Canvas(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("canvas", args...))
	return n
}

func Map_(args ...interface{}) *DomNode {
	return MakeDomNode("map_", args...)
}

func (n *DomNode) Map_(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("map_", args...))
	return n
}

func Area(args ...interface{}) *DomNode {
	return MakeDomNode("area", args...)
}

func (n *DomNode) Area(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("area", args...))
	return n
}

func Table(args ...interface{}) *DomNode {
	return MakeDomNode("table", args...)
}

func (n *DomNode) Table(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("table", args...))
	return n
}

func Caption(args ...interface{}) *DomNode {
	return MakeDomNode("caption", args...)
}

func (n *DomNode) Caption(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("caption", args...))
	return n
}

func Colgroup(args ...interface{}) *DomNode {
	return MakeDomNode("colgroup", args...)
}

func (n *DomNode) Colgroup(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("colgroup", args...))
	return n
}

func Col(args ...interface{}) *DomNode {
	return MakeDomNode("col", args...)
}

func (n *DomNode) Col(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("col", args...))
	return n
}

func Tbody(args ...interface{}) *DomNode {
	return MakeDomNode("tbody", args...)
}

func (n *DomNode) Tbody(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("tbody", args...))
	return n
}

func Thead(args ...interface{}) *DomNode {
	return MakeDomNode("thead", args...)
}

func (n *DomNode) Thead(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("thead", args...))
	return n
}

func Tfoot(args ...interface{}) *DomNode {
	return MakeDomNode("tfoot", args...)
}

func (n *DomNode) Tfoot(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("tfoot", args...))
	return n
}

func Tr(args ...interface{}) *DomNode {
	return MakeDomNode("tr", args...)
}

func (n *DomNode) Tr(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("tr", args...))
	return n
}

func Td(args ...interface{}) *DomNode {
	return MakeDomNode("td", args...)
}

func (n *DomNode) Td(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("td", args...))
	return n
}

func Th(args ...interface{}) *DomNode {
	return MakeDomNode("th", args...)
}

func (n *DomNode) Th(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("th", args...))
	return n
}

func Form(args ...interface{}) *DomNode {
	return MakeDomNode("form", args...)
}

func (n *DomNode) Form(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("form", args...))
	return n
}

func Fieldset(args ...interface{}) *DomNode {
	return MakeDomNode("fieldset", args...)
}

func (n *DomNode) Fieldset(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("fieldset", args...))
	return n
}

func Legend(args ...interface{}) *DomNode {
	return MakeDomNode("legend", args...)
}

func (n *DomNode) Legend(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("legend", args...))
	return n
}

func Label(args ...interface{}) *DomNode {
	return MakeDomNode("label", args...)
}

func (n *DomNode) Label(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("label", args...))
	return n
}

func Input_(args ...interface{}) *DomNode {
	return MakeDomNode("input_", args...)
}

func (n *DomNode) Input_(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("input_", args...))
	return n
}

func Button(args ...interface{}) *DomNode {
	return MakeDomNode("button", args...)
}

func (n *DomNode) Button(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("button", args...))
	return n
}

func Select(args ...interface{}) *DomNode {
	return MakeDomNode("select", args...)
}

func (n *DomNode) Select(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("select", args...))
	return n
}

func Datalist(args ...interface{}) *DomNode {
	return MakeDomNode("datalist", args...)
}

func (n *DomNode) Datalist(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("datalist", args...))
	return n
}

func Optgroup(args ...interface{}) *DomNode {
	return MakeDomNode("optgroup", args...)
}

func (n *DomNode) Optgroup(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("optgroup", args...))
	return n
}

func Option(args ...interface{}) *DomNode {
	return MakeDomNode("option", args...)
}

func (n *DomNode) Option(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("option", args...))
	return n
}

func Textarea(args ...interface{}) *DomNode {
	return MakeDomNode("textarea", args...)
}

func (n *DomNode) Textarea(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("textarea", args...))
	return n
}

func Keygen(args ...interface{}) *DomNode {
	return MakeDomNode("keygen", args...)
}

func (n *DomNode) Keygen(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("keygen", args...))
	return n
}

func Output(args ...interface{}) *DomNode {
	return MakeDomNode("output", args...)
}

func (n *DomNode) Output(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("output", args...))
	return n
}

func Progress(args ...interface{}) *DomNode {
	return MakeDomNode("progress", args...)
}

func (n *DomNode) Progress(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("progress", args...))
	return n
}

func Meter(args ...interface{}) *DomNode {
	return MakeDomNode("meter", args...)
}

func (n *DomNode) Meter(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("meter", args...))
	return n
}

func Details(args ...interface{}) *DomNode {
	return MakeDomNode("details", args...)
}

func (n *DomNode) Details(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("details", args...))
	return n
}

func Summary(args ...interface{}) *DomNode {
	return MakeDomNode("summary", args...)
}

func (n *DomNode) Summary(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("summary", args...))
	return n
}

func Command(args ...interface{}) *DomNode {
	return MakeDomNode("command", args...)
}

func (n *DomNode) Command(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("command", args...))
	return n
}

func Menu(args ...interface{}) *DomNode {
	return MakeDomNode("menu", args...)
}

func (n *DomNode) Menu(args ...interface{}) *DomNode {
	n.appendChild(MakeDomNode("menu", args...))
	return n
}
