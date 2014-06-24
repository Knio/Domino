package domino

func Html(args ...interface{}) *DomNode {
	return NewDomNode("html", args...)
}

func (n *DomNode) Html(args ...interface{}) *DomNode {
	n.Add(NewDomNode("html", args...))
	return n
}

func Head(args ...interface{}) *DomNode {
	return NewDomNode("head", args...)
}

func (n *DomNode) Head(args ...interface{}) *DomNode {
	n.Add(NewDomNode("head", args...))
	return n
}

func Title(args ...interface{}) *DomNode {
	return NewDomNode("title", args...)
}

func (n *DomNode) Title(args ...interface{}) *DomNode {
	n.Add(NewDomNode("title", args...))
	return n
}

func Base(args ...interface{}) *DomNode {
	return NewDomNode("base", args...)
}

func (n *DomNode) Base(args ...interface{}) *DomNode {
	n.Add(NewDomNode("base", args...))
	return n
}

func Link(args ...interface{}) *DomNode {
	return NewDomNode("link", args...)
}

func (n *DomNode) Link(args ...interface{}) *DomNode {
	n.Add(NewDomNode("link", args...))
	return n
}

func Meta(args ...interface{}) *DomNode {
	return NewDomNode("meta", args...)
}

func (n *DomNode) Meta(args ...interface{}) *DomNode {
	n.Add(NewDomNode("meta", args...))
	return n
}

func Style(args ...interface{}) *DomNode {
	return NewDomNode("style", args...)
}

func (n *DomNode) Style(args ...interface{}) *DomNode {
	n.Add(NewDomNode("style", args...))
	return n
}

func Script(args ...interface{}) *DomNode {
	return NewDomNode("script", args...)
}

func (n *DomNode) Script(args ...interface{}) *DomNode {
	n.Add(NewDomNode("script", args...))
	return n
}

func Noscript(args ...interface{}) *DomNode {
	return NewDomNode("noscript", args...)
}

func (n *DomNode) Noscript(args ...interface{}) *DomNode {
	n.Add(NewDomNode("noscript", args...))
	return n
}

func Body(args ...interface{}) *DomNode {
	return NewDomNode("body", args...)
}

func (n *DomNode) Body(args ...interface{}) *DomNode {
	n.Add(NewDomNode("body", args...))
	return n
}

func Section(args ...interface{}) *DomNode {
	return NewDomNode("section", args...)
}

func (n *DomNode) Section(args ...interface{}) *DomNode {
	n.Add(NewDomNode("section", args...))
	return n
}

func Nav(args ...interface{}) *DomNode {
	return NewDomNode("nav", args...)
}

func (n *DomNode) Nav(args ...interface{}) *DomNode {
	n.Add(NewDomNode("nav", args...))
	return n
}

func Article(args ...interface{}) *DomNode {
	return NewDomNode("article", args...)
}

func (n *DomNode) Article(args ...interface{}) *DomNode {
	n.Add(NewDomNode("article", args...))
	return n
}

func Aside(args ...interface{}) *DomNode {
	return NewDomNode("aside", args...)
}

func (n *DomNode) Aside(args ...interface{}) *DomNode {
	n.Add(NewDomNode("aside", args...))
	return n
}

func H1(args ...interface{}) *DomNode {
	return NewDomNode("h1", args...)
}

func (n *DomNode) H1(args ...interface{}) *DomNode {
	n.Add(NewDomNode("h1", args...))
	return n
}

func H2(args ...interface{}) *DomNode {
	return NewDomNode("h2", args...)
}

func (n *DomNode) H2(args ...interface{}) *DomNode {
	n.Add(NewDomNode("h2", args...))
	return n
}

func H3(args ...interface{}) *DomNode {
	return NewDomNode("h3", args...)
}

func (n *DomNode) H3(args ...interface{}) *DomNode {
	n.Add(NewDomNode("h3", args...))
	return n
}

func H4(args ...interface{}) *DomNode {
	return NewDomNode("h4", args...)
}

func (n *DomNode) H4(args ...interface{}) *DomNode {
	n.Add(NewDomNode("h4", args...))
	return n
}

func H5(args ...interface{}) *DomNode {
	return NewDomNode("h5", args...)
}

func (n *DomNode) H5(args ...interface{}) *DomNode {
	n.Add(NewDomNode("h5", args...))
	return n
}

func H6(args ...interface{}) *DomNode {
	return NewDomNode("h6", args...)
}

func (n *DomNode) H6(args ...interface{}) *DomNode {
	n.Add(NewDomNode("h6", args...))
	return n
}

func Hgroup(args ...interface{}) *DomNode {
	return NewDomNode("hgroup", args...)
}

func (n *DomNode) Hgroup(args ...interface{}) *DomNode {
	n.Add(NewDomNode("hgroup", args...))
	return n
}

func Header(args ...interface{}) *DomNode {
	return NewDomNode("header", args...)
}

func (n *DomNode) Header(args ...interface{}) *DomNode {
	n.Add(NewDomNode("header", args...))
	return n
}

func Footer(args ...interface{}) *DomNode {
	return NewDomNode("footer", args...)
}

func (n *DomNode) Footer(args ...interface{}) *DomNode {
	n.Add(NewDomNode("footer", args...))
	return n
}

func Address(args ...interface{}) *DomNode {
	return NewDomNode("address", args...)
}

func (n *DomNode) Address(args ...interface{}) *DomNode {
	n.Add(NewDomNode("address", args...))
	return n
}

func P(args ...interface{}) *DomNode {
	return NewDomNode("p", args...)
}

func (n *DomNode) P(args ...interface{}) *DomNode {
	n.Add(NewDomNode("p", args...))
	return n
}

func Hr(args ...interface{}) *DomNode {
	return NewDomNode("hr", args...)
}

func (n *DomNode) Hr(args ...interface{}) *DomNode {
	n.Add(NewDomNode("hr", args...))
	return n
}

func Pre(args ...interface{}) *DomNode {
	return NewDomNode("pre", args...)
}

func (n *DomNode) Pre(args ...interface{}) *DomNode {
	n.Add(NewDomNode("pre", args...))
	return n
}

func Blockquote(args ...interface{}) *DomNode {
	return NewDomNode("blockquote", args...)
}

func (n *DomNode) Blockquote(args ...interface{}) *DomNode {
	n.Add(NewDomNode("blockquote", args...))
	return n
}

func Ol(args ...interface{}) *DomNode {
	return NewDomNode("ol", args...)
}

func (n *DomNode) Ol(args ...interface{}) *DomNode {
	n.Add(NewDomNode("ol", args...))
	return n
}

func Ul(args ...interface{}) *DomNode {
	return NewDomNode("ul", args...)
}

func (n *DomNode) Ul(args ...interface{}) *DomNode {
	n.Add(NewDomNode("ul", args...))
	return n
}

func Li(args ...interface{}) *DomNode {
	return NewDomNode("li", args...)
}

func (n *DomNode) Li(args ...interface{}) *DomNode {
	n.Add(NewDomNode("li", args...))
	return n
}

func Dl(args ...interface{}) *DomNode {
	return NewDomNode("dl", args...)
}

func (n *DomNode) Dl(args ...interface{}) *DomNode {
	n.Add(NewDomNode("dl", args...))
	return n
}

func Dt(args ...interface{}) *DomNode {
	return NewDomNode("dt", args...)
}

func (n *DomNode) Dt(args ...interface{}) *DomNode {
	n.Add(NewDomNode("dt", args...))
	return n
}

func Dd(args ...interface{}) *DomNode {
	return NewDomNode("dd", args...)
}

func (n *DomNode) Dd(args ...interface{}) *DomNode {
	n.Add(NewDomNode("dd", args...))
	return n
}

func Figure(args ...interface{}) *DomNode {
	return NewDomNode("figure", args...)
}

func (n *DomNode) Figure(args ...interface{}) *DomNode {
	n.Add(NewDomNode("figure", args...))
	return n
}

func Figcaption(args ...interface{}) *DomNode {
	return NewDomNode("figcaption", args...)
}

func (n *DomNode) Figcaption(args ...interface{}) *DomNode {
	n.Add(NewDomNode("figcaption", args...))
	return n
}

func Div(args ...interface{}) *DomNode {
	return NewDomNode("div", args...)
}

func (n *DomNode) Div(args ...interface{}) *DomNode {
	n.Add(NewDomNode("div", args...))
	return n
}

func A(args ...interface{}) *DomNode {
	return NewDomNode("a", args...)
}

func (n *DomNode) A(args ...interface{}) *DomNode {
	n.Add(NewDomNode("a", args...))
	return n
}

func Em(args ...interface{}) *DomNode {
	return NewDomNode("em", args...)
}

func (n *DomNode) Em(args ...interface{}) *DomNode {
	n.Add(NewDomNode("em", args...))
	return n
}

func Strong(args ...interface{}) *DomNode {
	return NewDomNode("strong", args...)
}

func (n *DomNode) Strong(args ...interface{}) *DomNode {
	n.Add(NewDomNode("strong", args...))
	return n
}

func Small(args ...interface{}) *DomNode {
	return NewDomNode("small", args...)
}

func (n *DomNode) Small(args ...interface{}) *DomNode {
	n.Add(NewDomNode("small", args...))
	return n
}

func S(args ...interface{}) *DomNode {
	return NewDomNode("s", args...)
}

func (n *DomNode) S(args ...interface{}) *DomNode {
	n.Add(NewDomNode("s", args...))
	return n
}

func Cite(args ...interface{}) *DomNode {
	return NewDomNode("cite", args...)
}

func (n *DomNode) Cite(args ...interface{}) *DomNode {
	n.Add(NewDomNode("cite", args...))
	return n
}

func Q(args ...interface{}) *DomNode {
	return NewDomNode("q", args...)
}

func (n *DomNode) Q(args ...interface{}) *DomNode {
	n.Add(NewDomNode("q", args...))
	return n
}

func Dfn(args ...interface{}) *DomNode {
	return NewDomNode("dfn", args...)
}

func (n *DomNode) Dfn(args ...interface{}) *DomNode {
	n.Add(NewDomNode("dfn", args...))
	return n
}

func Abbr(args ...interface{}) *DomNode {
	return NewDomNode("abbr", args...)
}

func (n *DomNode) Abbr(args ...interface{}) *DomNode {
	n.Add(NewDomNode("abbr", args...))
	return n
}

func Time(args ...interface{}) *DomNode {
	return NewDomNode("time", args...)
}

func (n *DomNode) Time(args ...interface{}) *DomNode {
	n.Add(NewDomNode("time", args...))
	return n
}

func Code(args ...interface{}) *DomNode {
	return NewDomNode("code", args...)
}

func (n *DomNode) Code(args ...interface{}) *DomNode {
	n.Add(NewDomNode("code", args...))
	return n
}

func Var(args ...interface{}) *DomNode {
	return NewDomNode("var", args...)
}

func (n *DomNode) Var(args ...interface{}) *DomNode {
	n.Add(NewDomNode("var", args...))
	return n
}

func Samp(args ...interface{}) *DomNode {
	return NewDomNode("samp", args...)
}

func (n *DomNode) Samp(args ...interface{}) *DomNode {
	n.Add(NewDomNode("samp", args...))
	return n
}

func Kbd(args ...interface{}) *DomNode {
	return NewDomNode("kbd", args...)
}

func (n *DomNode) Kbd(args ...interface{}) *DomNode {
	n.Add(NewDomNode("kbd", args...))
	return n
}

func Sub(args ...interface{}) *DomNode {
	return NewDomNode("sub", args...)
}

func (n *DomNode) Sub(args ...interface{}) *DomNode {
	n.Add(NewDomNode("sub", args...))
	return n
}

func Sup(args ...interface{}) *DomNode {
	return NewDomNode("sup", args...)
}

func (n *DomNode) Sup(args ...interface{}) *DomNode {
	n.Add(NewDomNode("sup", args...))
	return n
}

func I(args ...interface{}) *DomNode {
	return NewDomNode("i", args...)
}

func (n *DomNode) I(args ...interface{}) *DomNode {
	n.Add(NewDomNode("i", args...))
	return n
}

func B(args ...interface{}) *DomNode {
	return NewDomNode("b", args...)
}

func (n *DomNode) B(args ...interface{}) *DomNode {
	n.Add(NewDomNode("b", args...))
	return n
}

func U(args ...interface{}) *DomNode {
	return NewDomNode("u", args...)
}

func (n *DomNode) U(args ...interface{}) *DomNode {
	n.Add(NewDomNode("u", args...))
	return n
}

func Mark(args ...interface{}) *DomNode {
	return NewDomNode("mark", args...)
}

func (n *DomNode) Mark(args ...interface{}) *DomNode {
	n.Add(NewDomNode("mark", args...))
	return n
}

func Ruby(args ...interface{}) *DomNode {
	return NewDomNode("ruby", args...)
}

func (n *DomNode) Ruby(args ...interface{}) *DomNode {
	n.Add(NewDomNode("ruby", args...))
	return n
}

func Rt(args ...interface{}) *DomNode {
	return NewDomNode("rt", args...)
}

func (n *DomNode) Rt(args ...interface{}) *DomNode {
	n.Add(NewDomNode("rt", args...))
	return n
}

func Rp(args ...interface{}) *DomNode {
	return NewDomNode("rp", args...)
}

func (n *DomNode) Rp(args ...interface{}) *DomNode {
	n.Add(NewDomNode("rp", args...))
	return n
}

func Bdi(args ...interface{}) *DomNode {
	return NewDomNode("bdi", args...)
}

func (n *DomNode) Bdi(args ...interface{}) *DomNode {
	n.Add(NewDomNode("bdi", args...))
	return n
}

func Bdo(args ...interface{}) *DomNode {
	return NewDomNode("bdo", args...)
}

func (n *DomNode) Bdo(args ...interface{}) *DomNode {
	n.Add(NewDomNode("bdo", args...))
	return n
}

func Span(args ...interface{}) *DomNode {
	return NewDomNode("span", args...)
}

func (n *DomNode) Span(args ...interface{}) *DomNode {
	n.Add(NewDomNode("span", args...))
	return n
}

func Br(args ...interface{}) *DomNode {
	return NewDomNode("br", args...)
}

func (n *DomNode) Br(args ...interface{}) *DomNode {
	n.Add(NewDomNode("br", args...))
	return n
}

func Wbr(args ...interface{}) *DomNode {
	return NewDomNode("wbr", args...)
}

func (n *DomNode) Wbr(args ...interface{}) *DomNode {
	n.Add(NewDomNode("wbr", args...))
	return n
}

func Ins(args ...interface{}) *DomNode {
	return NewDomNode("ins", args...)
}

func (n *DomNode) Ins(args ...interface{}) *DomNode {
	n.Add(NewDomNode("ins", args...))
	return n
}

func Del(args ...interface{}) *DomNode {
	return NewDomNode("del", args...)
}

func (n *DomNode) Del(args ...interface{}) *DomNode {
	n.Add(NewDomNode("del", args...))
	return n
}

func Img(args ...interface{}) *DomNode {
	return NewDomNode("img", args...)
}

func (n *DomNode) Img(args ...interface{}) *DomNode {
	n.Add(NewDomNode("img", args...))
	return n
}

func Iframe(args ...interface{}) *DomNode {
	return NewDomNode("iframe", args...)
}

func (n *DomNode) Iframe(args ...interface{}) *DomNode {
	n.Add(NewDomNode("iframe", args...))
	return n
}

func Embed(args ...interface{}) *DomNode {
	return NewDomNode("embed", args...)
}

func (n *DomNode) Embed(args ...interface{}) *DomNode {
	n.Add(NewDomNode("embed", args...))
	return n
}

func Object(args ...interface{}) *DomNode {
	return NewDomNode("object", args...)
}

func (n *DomNode) Object(args ...interface{}) *DomNode {
	n.Add(NewDomNode("object", args...))
	return n
}

func Param(args ...interface{}) *DomNode {
	return NewDomNode("param", args...)
}

func (n *DomNode) Param(args ...interface{}) *DomNode {
	n.Add(NewDomNode("param", args...))
	return n
}

func Video(args ...interface{}) *DomNode {
	return NewDomNode("video", args...)
}

func (n *DomNode) Video(args ...interface{}) *DomNode {
	n.Add(NewDomNode("video", args...))
	return n
}

func Audio(args ...interface{}) *DomNode {
	return NewDomNode("audio", args...)
}

func (n *DomNode) Audio(args ...interface{}) *DomNode {
	n.Add(NewDomNode("audio", args...))
	return n
}

func Source(args ...interface{}) *DomNode {
	return NewDomNode("source", args...)
}

func (n *DomNode) Source(args ...interface{}) *DomNode {
	n.Add(NewDomNode("source", args...))
	return n
}

func Track(args ...interface{}) *DomNode {
	return NewDomNode("track", args...)
}

func (n *DomNode) Track(args ...interface{}) *DomNode {
	n.Add(NewDomNode("track", args...))
	return n
}

func Canvas(args ...interface{}) *DomNode {
	return NewDomNode("canvas", args...)
}

func (n *DomNode) Canvas(args ...interface{}) *DomNode {
	n.Add(NewDomNode("canvas", args...))
	return n
}

func Map_(args ...interface{}) *DomNode {
	return NewDomNode("map_", args...)
}

func (n *DomNode) Map_(args ...interface{}) *DomNode {
	n.Add(NewDomNode("map_", args...))
	return n
}

func Area(args ...interface{}) *DomNode {
	return NewDomNode("area", args...)
}

func (n *DomNode) Area(args ...interface{}) *DomNode {
	n.Add(NewDomNode("area", args...))
	return n
}

func Table(args ...interface{}) *DomNode {
	return NewDomNode("table", args...)
}

func (n *DomNode) Table(args ...interface{}) *DomNode {
	n.Add(NewDomNode("table", args...))
	return n
}

func Caption(args ...interface{}) *DomNode {
	return NewDomNode("caption", args...)
}

func (n *DomNode) Caption(args ...interface{}) *DomNode {
	n.Add(NewDomNode("caption", args...))
	return n
}

func Colgroup(args ...interface{}) *DomNode {
	return NewDomNode("colgroup", args...)
}

func (n *DomNode) Colgroup(args ...interface{}) *DomNode {
	n.Add(NewDomNode("colgroup", args...))
	return n
}

func Col(args ...interface{}) *DomNode {
	return NewDomNode("col", args...)
}

func (n *DomNode) Col(args ...interface{}) *DomNode {
	n.Add(NewDomNode("col", args...))
	return n
}

func Tbody(args ...interface{}) *DomNode {
	return NewDomNode("tbody", args...)
}

func (n *DomNode) Tbody(args ...interface{}) *DomNode {
	n.Add(NewDomNode("tbody", args...))
	return n
}

func Thead(args ...interface{}) *DomNode {
	return NewDomNode("thead", args...)
}

func (n *DomNode) Thead(args ...interface{}) *DomNode {
	n.Add(NewDomNode("thead", args...))
	return n
}

func Tfoot(args ...interface{}) *DomNode {
	return NewDomNode("tfoot", args...)
}

func (n *DomNode) Tfoot(args ...interface{}) *DomNode {
	n.Add(NewDomNode("tfoot", args...))
	return n
}

func Tr(args ...interface{}) *DomNode {
	return NewDomNode("tr", args...)
}

func (n *DomNode) Tr(args ...interface{}) *DomNode {
	n.Add(NewDomNode("tr", args...))
	return n
}

func Td(args ...interface{}) *DomNode {
	return NewDomNode("td", args...)
}

func (n *DomNode) Td(args ...interface{}) *DomNode {
	n.Add(NewDomNode("td", args...))
	return n
}

func Th(args ...interface{}) *DomNode {
	return NewDomNode("th", args...)
}

func (n *DomNode) Th(args ...interface{}) *DomNode {
	n.Add(NewDomNode("th", args...))
	return n
}

func Form(args ...interface{}) *DomNode {
	return NewDomNode("form", args...)
}

func (n *DomNode) Form(args ...interface{}) *DomNode {
	n.Add(NewDomNode("form", args...))
	return n
}

func Fieldset(args ...interface{}) *DomNode {
	return NewDomNode("fieldset", args...)
}

func (n *DomNode) Fieldset(args ...interface{}) *DomNode {
	n.Add(NewDomNode("fieldset", args...))
	return n
}

func Legend(args ...interface{}) *DomNode {
	return NewDomNode("legend", args...)
}

func (n *DomNode) Legend(args ...interface{}) *DomNode {
	n.Add(NewDomNode("legend", args...))
	return n
}

func Label(args ...interface{}) *DomNode {
	return NewDomNode("label", args...)
}

func (n *DomNode) Label(args ...interface{}) *DomNode {
	n.Add(NewDomNode("label", args...))
	return n
}

func Input_(args ...interface{}) *DomNode {
	return NewDomNode("input_", args...)
}

func (n *DomNode) Input_(args ...interface{}) *DomNode {
	n.Add(NewDomNode("input_", args...))
	return n
}

func Button(args ...interface{}) *DomNode {
	return NewDomNode("button", args...)
}

func (n *DomNode) Button(args ...interface{}) *DomNode {
	n.Add(NewDomNode("button", args...))
	return n
}

func Select(args ...interface{}) *DomNode {
	return NewDomNode("select", args...)
}

func (n *DomNode) Select(args ...interface{}) *DomNode {
	n.Add(NewDomNode("select", args...))
	return n
}

func Datalist(args ...interface{}) *DomNode {
	return NewDomNode("datalist", args...)
}

func (n *DomNode) Datalist(args ...interface{}) *DomNode {
	n.Add(NewDomNode("datalist", args...))
	return n
}

func Optgroup(args ...interface{}) *DomNode {
	return NewDomNode("optgroup", args...)
}

func (n *DomNode) Optgroup(args ...interface{}) *DomNode {
	n.Add(NewDomNode("optgroup", args...))
	return n
}

func Option(args ...interface{}) *DomNode {
	return NewDomNode("option", args...)
}

func (n *DomNode) Option(args ...interface{}) *DomNode {
	n.Add(NewDomNode("option", args...))
	return n
}

func Textarea(args ...interface{}) *DomNode {
	return NewDomNode("textarea", args...)
}

func (n *DomNode) Textarea(args ...interface{}) *DomNode {
	n.Add(NewDomNode("textarea", args...))
	return n
}

func Keygen(args ...interface{}) *DomNode {
	return NewDomNode("keygen", args...)
}

func (n *DomNode) Keygen(args ...interface{}) *DomNode {
	n.Add(NewDomNode("keygen", args...))
	return n
}

func Output(args ...interface{}) *DomNode {
	return NewDomNode("output", args...)
}

func (n *DomNode) Output(args ...interface{}) *DomNode {
	n.Add(NewDomNode("output", args...))
	return n
}

func Progress(args ...interface{}) *DomNode {
	return NewDomNode("progress", args...)
}

func (n *DomNode) Progress(args ...interface{}) *DomNode {
	n.Add(NewDomNode("progress", args...))
	return n
}

func Meter(args ...interface{}) *DomNode {
	return NewDomNode("meter", args...)
}

func (n *DomNode) Meter(args ...interface{}) *DomNode {
	n.Add(NewDomNode("meter", args...))
	return n
}

func Details(args ...interface{}) *DomNode {
	return NewDomNode("details", args...)
}

func (n *DomNode) Details(args ...interface{}) *DomNode {
	n.Add(NewDomNode("details", args...))
	return n
}

func Summary(args ...interface{}) *DomNode {
	return NewDomNode("summary", args...)
}

func (n *DomNode) Summary(args ...interface{}) *DomNode {
	n.Add(NewDomNode("summary", args...))
	return n
}

func Command(args ...interface{}) *DomNode {
	return NewDomNode("command", args...)
}

func (n *DomNode) Command(args ...interface{}) *DomNode {
	n.Add(NewDomNode("command", args...))
	return n
}

func Menu(args ...interface{}) *DomNode {
	return NewDomNode("menu", args...)
}

func (n *DomNode) Menu(args ...interface{}) *DomNode {
	n.Add(NewDomNode("menu", args...))
	return n
}
