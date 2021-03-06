TAGS = '''
html
head
title
base
link
meta
style
script
noscript
body
section
nav
article
aside
h1
h2
h3
h4
h5
h6
hgroup
header
footer
address
p
hr
pre
blockquote
ol
ul
li
dl
dt
dd
figure
figcaption
div
a
em
strong
small
s
cite
q
dfn
abbr
time
code
var
samp
kbd
sub
sup
i
b
u
mark
ruby
rt
rp
bdi
bdo
span
br
wbr
ins
del
img
iframe
embed
object
param
video
audio
source
track
canvas
map_
area
table
caption
colgroup
col
tbody
thead
tfoot
tr
td
th
form
fieldset
legend
label
input_
button
select
datalist
optgroup
option
textarea
keygen
output
progress
meter
details
summary
command
menu
'''


if __name__ == '__main__':

  print('package domino\n\n')

  for tag in filter(None, TAGS.split()):
    fname = tag.title()

    print(
      '''func %s(args ...interface{}) *DomNode {
  return NewDomNode("%s", args...)
}\n\n''' % (fname, tag))

    print(
      '''func (n *DomNode) %s(args ...interface{}) *DomNode {
   n.Add(NewDomNode("%s", args...))
   return n
}
\n\n''' % (fname, tag))








