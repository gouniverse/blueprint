package webtheme

import "strings"

func bootstrapIconList() []struct {
	Name string
	Icon string
} {
	icons := `
0-circle : 0 circle
0-circle-fill : 0 circle fill
0-square : 0 square
0-square-fill : 0 square fill
1-circle : 1 circle
1-circle-fill : 1 circle fill
1-square : 1 square
1-square-fill : 1 square fill
123 : 123
2-circle : 2 circle
2-circle-fill : 2 circle fill
2-square-fill : 2 square fill
3-circle : 3 circle
3-circle-fill : 3 circle fill
3-square : 3 square
3-square-fill : 3 square fill
4-circle : 4 circle
4-circle-fill : 4 circle fill
4-square : 4 square
4-square-fill : 4 square fill
5-circle : 5 circle
5-circle-fill : 5 circle fill
5-square : 5 square
5-square-fill : 5 square fill
6-circle : 6 circle
6-circle-fill : 6 circle fill
6-square : 6 square
6-square-fill : 6 square fill
7-circle : 7 circle
7-circle-fill : 7 circle fill
7-square : 7 square
7-square-fill : 7 square fill
8-circle : 8 circle
8-circle-fill : 8 circle fill
8-square : 8 square
8-square-fill : 8 square fill
9-circle : 9 circle
9-circle-fill : 9 circle fill
9-square : 9 square
9-square-fill : 9 square fill
activity : Activity
airplane : Airplane
airplane-engines : Airplane engines
airplane-engines-fill : Airplane engines fill
airplane-fill : Airplane fill
alarm : Alarm
alarm-fill : Alarm fill
alexa : Alexa
align-bottom : Align bottom
align-center : Align center
align-end : Align end
align-middle : Align middle
align-start : Align start
align-top : Align top
alipay : Alipay
alphabet : Alphabet
alphabet-uppercase : Alphabet uppercase
alt : Alt
amazon : Amazon
amd : Amd
android : Android
android2 : Android2
app : App
app-indicator : App indicator
apple : Apple
archive : Archive
archive-fill : Archive fill
arrow-90deg-down : Arrow 90deg down
arrow-90deg-left : Arrow 90deg left
arrow-90deg-right : Arrow 90deg right
arrow-90deg-up : Arrow 90deg up
arrow-bar-down : Arrow bar down
arrow-bar-left : Arrow bar left
arrow-bar-right : Arrow bar right
arrow-bar-up : Arrow bar up
arrow-clockwise : Arrow clockwise
arrow-counterclockwise : Arrow counterclockwise
arrow-down : Arrow down
arrow-down-circle : Arrow down circle
arrow-down-circle-fill : Arrow down circle fill
arrow-down-left-circle : Arrow down left circle
arrow-down-left-circle-fill : Arrow down left circle fill
arrow-down-left-square : Arrow down left square
arrow-down-left-square-fill : Arrow down left square fill
arrow-down-right-circle : Arrow down right circle
arrow-down-right-circle-fill : Arrow down right circle fill
arrow-down-right-square : Arrow down right square
arrow-down-right-square-fill : Arrow down right square fill
arrow-down-square : Arrow down square
arrow-down-square-fill : Arrow down square fill
arrow-down-left : Arrow down left
arrow-down-right : Arrow down right
arrow-down-short : Arrow down short
arrow-down-up : Arrow down up
arrow-left : Arrow left
arrow-left-circle : Arrow left circle
arrow-left-circle-fill : Arrow left circle fill
arrow-left-square : Arrow left square
arrow-left-square-fill : Arrow left square fill
arrow-left-right : Arrow left right
arrow-left-short : Arrow left short
arrow-repeat : Arrow repeat
arrow-return-left : Arrow return left
arrow-return-right : Arrow return right
arrow-right : Arrow right
arrow-right-circle : Arrow right circle
arrow-right-circle-fill : Arrow right circle fill
arrow-right-square : Arrow right square
arrow-right-square-fill : Arrow right square fill
arrow-right-short : Arrow right short
arrow-through-heart : Arrow through heart
arrow-through-heart-fill : Arrow through heart fill
arrow-up : Arrow up
arrow-up-circle : Arrow up circle
arrow-up-circle-fill : Arrow up circle fill
arrow-up-left-circle : Arrow up left circle
arrow-up-left-circle-fill : Arrow up left circle fill
arrow-up-left-square : Arrow up left square
arrow-up-left-square-fill : Arrow up left square fill
arrow-up-right-circle : Arrow up right circle
arrow-up-right-circle-fill : Arrow up right circle fill
arrow-up-right-square : Arrow up right square
arrow-up-right-square-fill : Arrow up right square fill
arrow-up-square : Arrow up square
arrow-up-square-fill : Arrow up square fill
arrow-up-left : Arrow up left
arrow-up-right : Arrow up right
arrow-up-short : Arrow up short
arrows : Arrows
arrows-angle-contract : Arrows angle contract
arrows-angle-expand : Arrows angle expand
arrows-collapse : Arrows collapse
arrows-collapse-vertical : Arrows collapse vertical
arrows-expand : Arrows expand
arrows-expand-vertical : Arrows expand vertical
arrows-fullscreen : Arrows fullscreen
arrows-move : Arrows move
arrows-vertical : Arrows vertical
aspect-ratio : Aspect ratio
aspect-ratio-fill : Aspect ratio fill
asterisk : Asterisk
at : At
award : Award
award-fill : Award fill
back : Back
backpack : Backpack
backpack-fill : Backpack fill
backpack2 : Backpack2
backpack2-fill : Backpack2 fill
backpack3 : Backpack3
backpack3-fill : Backpack3 fill
backpack4 : Backpack4
backpack4-fill : Backpack4 fill
backspace : Backspace
backspace-fill : Backspace fill
backspace-reverse : Backspace reverse
backspace-reverse-fill : Backspace reverse fill
badge-3d : Badge 3d
badge-3d-fill : Badge 3d fill
badge-4k : Badge 4k
badge-4k-fill : Badge 4k fill
badge-8k : Badge 8k
badge-8k-fill : Badge 8k fill
badge-ad : Badge ad
badge-ad-fill : Badge ad fill
badge-ar : Badge ar
badge-ar-fill : Badge ar fill
badge-cc : Badge cc
badge-cc-fill : Badge cc fill
badge-hd : Badge hd
badge-hd-fill : Badge hd fill
badge-sd : Badge sd
badge-sd-fill : Badge sd fill
badge-tm : Badge tm
badge-tm-fill : Badge tm fill
badge-vo : Badge vo
badge-vo-fill : Badge vo fill
badge-vr : Badge vr
badge-vr-fill : Badge vr fill
badge-wc : Badge wc
badge-wc-fill : Badge wc fill
bag : Bag
bag-check : Bag check
bag-check-fill : Bag check fill
bag-dash : Bag dash
bag-dash-fill : Bag dash fill
bag-fill : Bag fill
bag-heart : Bag heart
bag-heart-fill : Bag heart fill
bag-plus : Bag plus
bag-plus-fill : Bag plus fill
bag-x : Bag x
bag-x-fill : Bag x fill
balloon : Balloon
balloon-fill : Balloon fill
balloon-heart : Balloon heart
balloon-heart-fill : Balloon heart fill
ban : Ban
ban-fill : Ban fill
bandaid : Bandaid
bandaid-fill : Bandaid fill
bank : Bank
bank2 : Bank2
bar-chart : Bar chart
bar-chart-fill : Bar chart fill
bar-chart-line : Bar chart line
bar-chart-line-fill : Bar chart line fill
bar-chart-steps : Bar chart steps
basket : Basket
basket-fill : Basket fill
basket2 : Basket2
basket2-fill : Basket2 fill
basket3 : Basket3
basket3-fill : Basket3 fill
battery : Battery
battery-charging : Battery charging
battery-full : Battery full
battery-half : Battery half
behance : Behance
bell : Bell
bell-fill : Bell fill
bell-slash : Bell slash
bell-slash-fill : Bell slash fill
bezier : Bezier
bezier2 : Bezier2
bicycle : Bicycle
bing : Bing
binoculars : Binoculars
binoculars-fill : Binoculars fill
blockquote-left : Blockquote left
blockquote-right : Blockquote right
bluetooth : Bluetooth
body-text : Body text
book : Book
book-fill : Book fill
book-half : Book half
bookmark : Bookmark
bookmark-check : Bookmark check
bookmark-check-fill : Bookmark check fill
bookmark-dash : Bookmark dash
bookmark-dash-fill : Bookmark dash fill
bookmark-fill : Bookmark fill
bookmark-heart : Bookmark heart
bookmark-heart-fill : Bookmark heart fill
bookmark-plus : Bookmark plus
bookmark-plus-fill : Bookmark plus fill
bookmark-star : Bookmark star
bookmark-star-fill : Bookmark star fill
bookmark-x : Bookmark x
bookmark-x-fill : Bookmark x fill
bookmarks : Bookmarks
bookmarks-fill : Bookmarks fill
bookshelf : Bookshelf
boombox : Boombox
boombox-fill : Boombox fill
bootstrap : Bootstrap
bootstrap-fill : Bootstrap fill
bootstrap-reboot : Bootstrap reboot
border : Border
border-all : Border all
border-bottom : Border bottom
border-center : Border center
border-inner : Border inner
border-left : Border left
border-middle : Border middle
border-outer : Border outer
border-right : Border right
border-style : Border style
border-top : Border top
border-width : Border width
bounding-box : Bounding box
bounding-box-circles : Bounding box circles
box : Box
box-arrow-down-left : Box arrow down left
box-arrow-down-right : Box arrow down right
box-arrow-down : Box arrow down
box-arrow-in-down : Box arrow in down
box-arrow-in-down-left : Box arrow in down left
box-arrow-in-down-right : Box arrow in down right
box-arrow-in-left : Box arrow in left
box-arrow-in-right : Box arrow in right
box-arrow-in-up : Box arrow in up
box-arrow-in-up-left : Box arrow in up left
box-arrow-in-up-right : Box arrow in up right
box-arrow-left : Box arrow left
box-arrow-right : Box arrow right
box-arrow-up : Box arrow up
box-arrow-up-left : Box arrow up left
box-arrow-up-right : Box arrow up right
box-fill : Box fill
box-seam : Box seam
box-seam-fill : Box seam fill
box2 : Box2
box2-fill : Box2 fill
box2-heart : Box2 heart
box2-heart-fill : Box2 heart fill
boxes : Boxes
braces : Braces
braces-asterisk : Braces asterisk
bricks : Bricks
briefcase : Briefcase
briefcase-fill : Briefcase fill
brightness-alt-high : Brightness alt high
brightness-alt-high-fill : Brightness alt high fill
brightness-alt-low : Brightness alt low
brightness-alt-low-fill : Brightness alt low fill
brightness-high : Brightness high
brightness-high-fill : Brightness high fill
brightness-low : Brightness low
brightness-low-fill : Brightness low fill
brilliance : Brilliance
broadcast : Broadcast
broadcast-pin : Broadcast pin
browser-chrome : Browser chrome
browser-edge : Browser edge
browser-firefox : Browser firefox
browser-safari : Browser safari
brush : Brush
brush-fill : Brush fill
bucket : Bucket
bucket-fill : Bucket fill
bug : Bug
bug-fill : Bug fill
building : Building
building-add : Building add
building-check : Building check
building-dash : Building dash
building-down : Building down
building-exclamation : Building exclamation
building-fill : Building fill
building-fill-add : Building fill add
building-fill-check : Building fill check
building-fill-dash : Building fill dash
building-fill-down : Building fill down
building-fill-exclamation : Building fill exclamation
building-fill-gear : Building fill gear
building-fill-lock : Building fill lock
building-fill-slash : Building fill slash
building-fill-up : Building fill up
building-fill-x : Building fill x
building-gear : Building gear
building-lock : Building lock
building-slash : Building slash
building-up : Building up
building-x : Building x
buildings : Buildings
buildings-fill : Buildings fill
bullseye : Bullseye
bus-front : Bus front
bus-front-fill : Bus front fill
c-circle : C circle
c-circle-fill : C circle fill
c-square : C square
c-square-fill : C square fill
cake : Cake
cake-fill : Cake fill
cake2 : Cake2
cake2-fill : Cake2 fill
calculator : Calculator
calculator-fill : Calculator fill
calendar : Calendar
calendar-check : Calendar check
calendar-check-fill : Calendar check fill
calendar-date : Calendar date
calendar-date-fill : Calendar date fill
calendar-day : Calendar day
calendar-day-fill : Calendar day fill
calendar-event : Calendar event
calendar-event-fill : Calendar event fill
calendar-fill : Calendar fill
calendar-heart : Calendar heart
calendar-heart-fill : Calendar heart fill
calendar-minus : Calendar minus
calendar-minus-fill : Calendar minus fill
calendar-month : Calendar month
calendar-month-fill : Calendar month fill
calendar-plus : Calendar plus
calendar-plus-fill : Calendar plus fill
calendar-range : Calendar range
calendar-range-fill : Calendar range fill
calendar-week : Calendar week
calendar-week-fill : Calendar week fill
calendar-x : Calendar x
calendar-x-fill : Calendar x fill
calendar2 : Calendar2
calendar2-check : Calendar2 check
calendar2-check-fill : Calendar2 check fill
calendar2-date : Calendar2 date
calendar2-date-fill : Calendar2 date fill
calendar2-day : Calendar2 day
calendar2-day-fill : Calendar2 day fill
calendar2-event : Calendar2 event
calendar2-event-fill : Calendar2 event fill
calendar2-fill : Calendar2 fill
calendar2-heart : Calendar2 heart
calendar2-heart-fill : Calendar2 heart fill
calendar2-minus : Calendar2 minus
calendar2-minus-fill : Calendar2 minus fill
calendar2-month : Calendar2 month
calendar2-month-fill : Calendar2 month fill
calendar2-plus : Calendar2 plus
calendar2-plus-fill : Calendar2 plus fill
calendar2-range : Calendar2 range
calendar2-range-fill : Calendar2 range fill
calendar2-week : Calendar2 week
calendar2-week-fill : Calendar2 week fill
calendar2-x : Calendar2 x
calendar2-x-fill : Calendar2 x fill
calendar3 : Calendar3
calendar3-event : Calendar3 event
calendar3-event-fill : Calendar3 event fill
calendar3-fill : Calendar3 fill
calendar3-range : Calendar3 range
calendar3-range-fill : Calendar3 range fill
calendar3-week : Calendar3 week
calendar3-week-fill : Calendar3 week fill
calendar4 : Calendar4
calendar4-event : Calendar4 event
calendar4-range : Calendar4 range
calendar4-week : Calendar4 week
camera : Camera
camera2 : Camera2
camera-fill : Camera fill
camera-reels : Camera reels
camera-reels-fill : Camera reels fill
camera-video : Camera video
camera-video-fill : Camera video fill
camera-video-off : Camera video off
camera-video-off-fill : Camera video off fill
capslock : Capslock
capslock-fill : Capslock fill
capsule : Capsule
capsule-pill : Capsule pill
car-front : Car front
car-front-fill : Car front fill
card-checklist : Card checklist
card-heading : Card heading
card-image : Card image
card-list : Card list
card-text : Card text
caret-down : Caret down
caret-down-fill : Caret down fill
caret-down-square : Caret down square
caret-down-square-fill : Caret down square fill
caret-left : Caret left
caret-left-fill : Caret left fill
caret-left-square : Caret left square
caret-left-square-fill : Caret left square fill
caret-right : Caret right
caret-right-fill : Caret right fill
caret-right-square : Caret right square
caret-right-square-fill : Caret right square fill
caret-up : Caret up
caret-up-fill : Caret up fill
caret-up-square : Caret up square
caret-up-square-fill : Caret up square fill
cart : Cart
cart-check : Cart check
cart-check-fill : Cart check fill
cart-dash : Cart dash
cart-dash-fill : Cart dash fill
cart-fill : Cart fill
cart-plus : Cart plus
cart-plus-fill : Cart plus fill
cart-x : Cart x
cart-x-fill : Cart x fill
cart2 : Cart2
cart3 : Cart3
cart4 : Cart4
cash : Cash
cash-coin : Cash coin
cash-stack : Cash stack
cassette : Cassette
cassette-fill : Cassette fill
cast : Cast
cc-circle : Cc circle
cc-circle-fill : Cc circle fill
cc-square : Cc square
cc-square-fill : Cc square fill
chat : Chat
chat-dots : Chat dots
chat-dots-fill : Chat dots fill
chat-fill : Chat fill
chat-heart : Chat heart
chat-heart-fill : Chat heart fill
chat-left : Chat left
chat-left-dots : Chat left dots
chat-left-dots-fill : Chat left dots fill
chat-left-fill : Chat left fill
chat-left-heart : Chat left heart
chat-left-heart-fill : Chat left heart fill
chat-left-quote : Chat left quote
chat-left-quote-fill : Chat left quote fill
chat-left-text : Chat left text
chat-left-text-fill : Chat left text fill
chat-quote : Chat quote
chat-quote-fill : Chat quote fill
chat-right : Chat right
chat-right-dots : Chat right dots
chat-right-dots-fill : Chat right dots fill
chat-right-fill : Chat right fill
chat-right-heart : Chat right heart
chat-right-heart-fill : Chat right heart fill
chat-right-quote : Chat right quote
chat-right-quote-fill : Chat right quote fill
chat-right-text : Chat right text
chat-right-text-fill : Chat right text fill
chat-square : Chat square
chat-square-dots : Chat square dots
chat-square-dots-fill : Chat square dots fill
chat-square-fill : Chat square fill
chat-square-heart : Chat square heart
chat-square-heart-fill : Chat square heart fill
chat-square-quote : Chat square quote
chat-square-quote-fill : Chat square quote fill
chat-square-text : Chat square text
chat-square-text-fill : Chat square text fill
chat-text : Chat text
chat-text-fill : Chat text fill
check : Check
check-all : Check all
check-circle : Check circle
check-circle-fill : Check circle fill
check-lg : Check lg
check-square : Check square
check-square-fill : Check square fill
check2 : Check2
check2-all : Check2 all
check2-circle : Check2 circle
check2-square : Check2 square
chevron-bar-contract : Chevron bar contract
chevron-bar-down : Chevron bar down
chevron-bar-expand : Chevron bar expand
chevron-bar-left : Chevron bar left
chevron-bar-right : Chevron bar right
chevron-bar-up : Chevron bar up
chevron-compact-down : Chevron compact down
chevron-compact-left : Chevron compact left
chevron-compact-right : Chevron compact right
chevron-compact-up : Chevron compact up
chevron-contract : Chevron contract
chevron-double-down : Chevron double down
chevron-double-left : Chevron double left
chevron-double-right : Chevron double right
chevron-double-up : Chevron double up
chevron-down : Chevron down
chevron-expand : Chevron expand
chevron-left : Chevron left
chevron-right : Chevron right
chevron-up : Chevron up
circle : Circle
circle-fill : Circle fill
circle-half : Circle half
slash-circle : Slash circle
circle-square : Circle square
clipboard : Clipboard
clipboard-check : Clipboard check
clipboard-check-fill : Clipboard check fill
clipboard-data : Clipboard data
clipboard-data-fill : Clipboard data fill
clipboard-fill : Clipboard fill
clipboard-heart : Clipboard heart
clipboard-heart-fill : Clipboard heart fill
clipboard-minus : Clipboard minus
clipboard-minus-fill : Clipboard minus fill
clipboard-plus : Clipboard plus
clipboard-plus-fill : Clipboard plus fill
clipboard-pulse : Clipboard pulse
clipboard-x : Clipboard x
clipboard-x-fill : Clipboard x fill
clipboard2 : Clipboard2
clipboard2-check : Clipboard2 check
clipboard2-check-fill : Clipboard2 check fill
clipboard2-data : Clipboard2 data
clipboard2-data-fill : Clipboard2 data fill
clipboard2-fill : Clipboard2 fill
clipboard2-heart : Clipboard2 heart
clipboard2-heart-fill : Clipboard2 heart fill
clipboard2-minus : Clipboard2 minus
clipboard2-minus-fill : Clipboard2 minus fill
clipboard2-plus : Clipboard2 plus
clipboard2-plus-fill : Clipboard2 plus fill
clipboard2-pulse : Clipboard2 pulse
clipboard2-pulse-fill : Clipboard2 pulse fill
clipboard2-x : Clipboard2 x
clipboard2-x-fill : Clipboard2 x fill
clock : Clock
clock-fill : Clock fill
clock-history : Clock history
cloud : Cloud
cloud-arrow-down : Cloud arrow down
cloud-arrow-down-fill : Cloud arrow down fill
cloud-arrow-up : Cloud arrow up
cloud-arrow-up-fill : Cloud arrow up fill
cloud-check : Cloud check
cloud-check-fill : Cloud check fill
cloud-download : Cloud download
cloud-download-fill : Cloud download fill
cloud-drizzle : Cloud drizzle
cloud-drizzle-fill : Cloud drizzle fill
cloud-fill : Cloud fill
cloud-fog : Cloud fog
cloud-fog-fill : Cloud fog fill
cloud-fog2 : Cloud fog2
cloud-fog2-fill : Cloud fog2 fill
cloud-hail : Cloud hail
cloud-hail-fill : Cloud hail fill
cloud-haze : Cloud haze
cloud-haze-fill : Cloud haze fill
cloud-haze2 : Cloud haze2
cloud-haze2-fill : Cloud haze2 fill
cloud-lightning : Cloud lightning
cloud-lightning-fill : Cloud lightning fill
cloud-lightning-rain : Cloud lightning rain
cloud-lightning-rain-fill : Cloud lightning rain fill
cloud-minus : Cloud minus
cloud-minus-fill : Cloud minus fill
cloud-moon : Cloud moon
cloud-moon-fill : Cloud moon fill
cloud-plus : Cloud plus
cloud-plus-fill : Cloud plus fill
cloud-rain : Cloud rain
cloud-rain-fill : Cloud rain fill
cloud-rain-heavy : Cloud rain heavy
cloud-rain-heavy-fill : Cloud rain heavy fill
cloud-slash : Cloud slash
cloud-slash-fill : Cloud slash fill
cloud-sleet : Cloud sleet
cloud-sleet-fill : Cloud sleet fill
cloud-snow : Cloud snow
cloud-snow-fill : Cloud snow fill
cloud-sun : Cloud sun
cloud-sun-fill : Cloud sun fill
cloud-upload : Cloud upload
cloud-upload-fill : Cloud upload fill
clouds : Clouds
clouds-fill : Clouds fill
cloudy : Cloudy
cloudy-fill : Cloudy fill
code : Code
code-slash : Code slash
code-square : Code square
coin : Coin
collection : Collection
collection-fill : Collection fill
collection-play : Collection play
collection-play-fill : Collection play fill
columns : Columns
columns-gap : Columns gap
command : Command
compass : Compass
compass-fill : Compass fill
cone : Cone
cone-striped : Cone striped
controller : Controller
cookie : Cookie
copy : Copy
cpu : Cpu
cpu-fill : Cpu fill
credit-card : Credit card
credit-card-2-back : Credit card 2 back
credit-card-2-back-fill : Credit card 2 back fill
credit-card-2-front : Credit card 2 front
credit-card-2-front-fill : Credit card 2 front fill
credit-card-fill : Credit card fill
crop : Crop
crosshair : Crosshair
crosshair2 : Crosshair2
cup : Cup
cup-fill : Cup fill
cup-hot : Cup hot
cup-hot-fill : Cup hot fill
cup-straw : Cup straw
currency-bitcoin : Currency bitcoin
currency-dollar : Currency dollar
currency-euro : Currency euro
currency-exchange : Currency exchange
currency-pound : Currency pound
currency-rupee : Currency rupee
currency-yen : Currency yen
cursor : Cursor
cursor-fill : Cursor fill
cursor-text : Cursor text
dash : Dash
dash-circle : Dash circle
dash-circle-dotted : Dash circle dotted
dash-circle-fill : Dash circle fill
dash-lg : Dash lg
dash-square : Dash square
dash-square-dotted : Dash square dotted
dash-square-fill : Dash square fill
database : Database
database-add : Database add
database-check : Database check
database-dash : Database dash
database-down : Database down
database-exclamation : Database exclamation
database-fill : Database fill
database-fill-add : Database fill add
database-fill-check : Database fill check
database-fill-dash : Database fill dash
database-fill-down : Database fill down
database-fill-exclamation : Database fill exclamation
database-fill-gear : Database fill gear
database-fill-lock : Database fill lock
database-fill-slash : Database fill slash
database-fill-up : Database fill up
database-fill-x : Database fill x
database-gear : Database gear
database-lock : Database lock
database-slash : Database slash
database-up : Database up
database-x : Database x
device-hdd : Device hdd
device-hdd-fill : Device hdd fill
device-ssd : Device ssd
device-ssd-fill : Device ssd fill
diagram-2 : Diagram 2
diagram-2-fill : Diagram 2 fill
diagram-3 : Diagram 3
diagram-3-fill : Diagram 3 fill
diamond : Diamond
diamond-fill : Diamond fill
diamond-half : Diamond half
dice-1 : Dice 1
dice-1-fill : Dice 1 fill
dice-2 : Dice 2
dice-2-fill : Dice 2 fill
dice-3 : Dice 3
dice-3-fill : Dice 3 fill
dice-4 : Dice 4
dice-4-fill : Dice 4 fill
dice-5 : Dice 5
dice-5-fill : Dice 5 fill
dice-6 : Dice 6
dice-6-fill : Dice 6 fill
disc : Disc
disc-fill : Disc fill
discord : Discord
display : Display
display-fill : Display fill
displayport : Displayport
displayport-fill : Displayport fill
distribute-horizontal : Distribute horizontal
distribute-vertical : Distribute vertical
door-closed : Door closed
door-closed-fill : Door closed fill
door-open : Door open
door-open-fill : Door open fill
dot : Dot
download : Download
dpad : Dpad
dpad-fill : Dpad fill
dribbble : Dribbble
dropbox : Dropbox
droplet : Droplet
droplet-fill : Droplet fill
droplet-half : Droplet half
duffle : Duffle
duffle-fill : Duffle fill
ear : Ear
ear-fill : Ear fill
earbuds : Earbuds
easel : Easel
easel-fill : Easel fill
easel2 : Easel2
easel2-fill : Easel2 fill
easel3 : Easel3
easel3-fill : Easel3 fill
egg : Egg
egg-fill : Egg fill
egg-fried : Egg fried
eject : Eject
eject-fill : Eject fill
emoji-angry : Emoji angry
emoji-angry-fill : Emoji angry fill
emoji-astonished : Emoji astonished
emoji-astonished-fill : Emoji astonished fill
emoji-dizzy : Emoji dizzy
emoji-dizzy-fill : Emoji dizzy fill
emoji-expressionless : Emoji expressionless
emoji-expressionless-fill : Emoji expressionless fill
emoji-frown : Emoji frown
emoji-frown-fill : Emoji frown fill
emoji-grimace : Emoji grimace
emoji-grimace-fill : Emoji grimace fill
emoji-grin : Emoji grin
emoji-grin-fill : Emoji grin fill
emoji-heart-eyes : Emoji heart eyes
emoji-heart-eyes-fill : Emoji heart eyes fill
emoji-kiss : Emoji kiss
emoji-kiss-fill : Emoji kiss fill
emoji-laughing : Emoji laughing
emoji-laughing-fill : Emoji laughing fill
emoji-neutral : Emoji neutral
emoji-neutral-fill : Emoji neutral fill
emoji-smile : Emoji smile
emoji-smile-fill : Emoji smile fill
emoji-smile-upside-down : Emoji smile upside down
emoji-smile-upside-down-fill : Emoji smile upside down fill
emoji-sunglasses : Emoji sunglasses
emoji-sunglasses-fill : Emoji sunglasses fill
emoji-surprise : Emoji surprise
emoji-surprise-fill : Emoji surprise fill
emoji-tear : Emoji tear
emoji-tear-fill : Emoji tear fill
emoji-wink : Emoji wink
emoji-wink-fill : Emoji wink fill
envelope : Envelope
envelope-arrow-down : Envelope arrow down
envelope-arrow-down-fill : Envelope arrow down fill
envelope-arrow-up : Envelope arrow up
envelope-arrow-up-fill : Envelope arrow up fill
envelope-at : Envelope at
envelope-at-fill : Envelope at fill
envelope-check : Envelope check
envelope-check-fill : Envelope check fill
envelope-dash : Envelope dash
envelope-dash-fill : Envelope dash fill
envelope-exclamation : Envelope exclamation
envelope-exclamation-fill : Envelope exclamation fill
envelope-fill : Envelope fill
envelope-heart : Envelope heart
envelope-heart-fill : Envelope heart fill
envelope-open : Envelope open
envelope-open-fill : Envelope open fill
envelope-open-heart : Envelope open heart
envelope-open-heart-fill : Envelope open heart fill
envelope-paper : Envelope paper
envelope-paper-fill : Envelope paper fill
envelope-paper-heart : Envelope paper heart
envelope-paper-heart-fill : Envelope paper heart fill
envelope-plus : Envelope plus
envelope-plus-fill : Envelope plus fill
envelope-slash : Envelope slash
envelope-slash-fill : Envelope slash fill
envelope-x : Envelope x
envelope-x-fill : Envelope x fill
eraser : Eraser
eraser-fill : Eraser fill
escape : Escape
ethernet : Ethernet
ev-front : Ev front
ev-front-fill : Ev front fill
ev-station : Ev station
ev-station-fill : Ev station fill
exclamation : Exclamation
exclamation-circle : Exclamation circle
exclamation-circle-fill : Exclamation circle fill
exclamation-diamond : Exclamation diamond
exclamation-diamond-fill : Exclamation diamond fill
exclamation-lg : Exclamation lg
exclamation-octagon : Exclamation octagon
exclamation-octagon-fill : Exclamation octagon fill
exclamation-square : Exclamation square
exclamation-square-fill : Exclamation square fill
exclamation-triangle : Exclamation triangle
exclamation-triangle-fill : Exclamation triangle fill
exclude : Exclude
explicit : Explicit
explicit-fill : Explicit fill
exposure : Exposure
eye : Eye
eye-fill : Eye fill
eye-slash : Eye slash
eye-slash-fill : Eye slash fill
eyedropper : Eyedropper
eyeglasses : Eyeglasses
facebook : Facebook
fan : Fan
fast-forward : Fast forward
fast-forward-btn : Fast forward btn
fast-forward-btn-fill : Fast forward btn fill
fast-forward-circle : Fast forward circle
fast-forward-circle-fill : Fast forward circle fill
fast-forward-fill : Fast forward fill
feather : Feather
feather2 : Feather2
file : File
file-arrow-down : File arrow down
file-arrow-down-fill : File arrow down fill
file-arrow-up : File arrow up
file-arrow-up-fill : File arrow up fill
file-bar-graph : File bar graph
file-bar-graph-fill : File bar graph fill
file-binary : File binary
file-binary-fill : File binary fill
file-break : File break
file-break-fill : File break fill
file-check : File check
file-check-fill : File check fill
file-code : File code
file-code-fill : File code fill
file-diff : File diff
file-diff-fill : File diff fill
file-earmark : File earmark
file-earmark-arrow-down : File earmark arrow down
file-earmark-arrow-down-fill : File earmark arrow down fill
file-earmark-arrow-up : File earmark arrow up
file-earmark-arrow-up-fill : File earmark arrow up fill
file-earmark-bar-graph : File earmark bar graph
file-earmark-bar-graph-fill : File earmark bar graph fill
file-earmark-binary : File earmark binary
file-earmark-binary-fill : File earmark binary fill
file-earmark-break : File earmark break
file-earmark-break-fill : File earmark break fill
file-earmark-check : File earmark check
file-earmark-check-fill : File earmark check fill
file-earmark-code : File earmark code
file-earmark-code-fill : File earmark code fill
file-earmark-diff : File earmark diff
file-earmark-diff-fill : File earmark diff fill
file-earmark-easel : File earmark easel
file-earmark-easel-fill : File earmark easel fill
file-earmark-excel : File earmark excel
file-earmark-excel-fill : File earmark excel fill
file-earmark-fill : File earmark fill
file-earmark-font : File earmark font
file-earmark-font-fill : File earmark font fill
file-earmark-image : File earmark image
file-earmark-image-fill : File earmark image fill
file-earmark-lock : File earmark lock
file-earmark-lock-fill : File earmark lock fill
file-earmark-lock2 : File earmark lock2
file-earmark-lock2-fill : File earmark lock2 fill
file-earmark-medical : File earmark medical
file-earmark-medical-fill : File earmark medical fill
file-earmark-minus : File earmark minus
file-earmark-minus-fill : File earmark minus fill
file-earmark-music : File earmark music
file-earmark-music-fill : File earmark music fill
file-earmark-pdf : File earmark pdf
file-earmark-pdf-fill : File earmark pdf fill
file-earmark-person : File earmark person
file-earmark-person-fill : File earmark person fill
file-earmark-play : File earmark play
file-earmark-play-fill : File earmark play fill
file-earmark-plus : File earmark plus
file-earmark-plus-fill : File earmark plus fill
file-earmark-post : File earmark post
file-earmark-post-fill : File earmark post fill
file-earmark-ppt : File earmark ppt
file-earmark-ppt-fill : File earmark ppt fill
file-earmark-richtext : File earmark richtext
file-earmark-richtext-fill : File earmark richtext fill
file-earmark-ruled : File earmark ruled
file-earmark-ruled-fill : File earmark ruled fill
file-earmark-slides : File earmark slides
file-earmark-slides-fill : File earmark slides fill
file-earmark-spreadsheet : File earmark spreadsheet
file-earmark-spreadsheet-fill : File earmark spreadsheet fill
file-earmark-text : File earmark text
file-earmark-text-fill : File earmark text fill
file-earmark-word : File earmark word
file-earmark-word-fill : File earmark word fill
file-earmark-x : File earmark x
file-earmark-x-fill : File earmark x fill
file-earmark-zip : File earmark zip
file-earmark-zip-fill : File earmark zip fill
file-easel : File easel
file-easel-fill : File easel fill
file-excel : File excel
file-excel-fill : File excel fill
file-fill : File fill
file-font : File font
file-font-fill : File font fill
file-image : File image
file-image-fill : File image fill
file-lock : File lock
file-lock-fill : File lock fill
file-lock2 : File lock2
file-lock2-fill : File lock2 fill
file-medical : File medical
file-medical-fill : File medical fill
file-minus : File minus
file-minus-fill : File minus fill
file-music : File music
file-music-fill : File music fill
file-pdf : File pdf
file-pdf-fill : File pdf fill
file-person : File person
file-person-fill : File person fill
file-play : File play
file-play-fill : File play fill
file-plus : File plus
file-plus-fill : File plus fill
file-post : File post
file-post-fill : File post fill
file-ppt : File ppt
file-ppt-fill : File ppt fill
file-richtext : File richtext
file-richtext-fill : File richtext fill
file-ruled : File ruled
file-ruled-fill : File ruled fill
file-slides : File slides
file-slides-fill : File slides fill
file-spreadsheet : File spreadsheet
file-spreadsheet-fill : File spreadsheet fill
file-text : File text
file-text-fill : File text fill
file-word : File word
file-word-fill : File word fill
file-x : File x
file-x-fill : File x fill
file-zip : File zip
file-zip-fill : File zip fill
files : Files
files-alt : Files alt
filetype-aac : Filetype aac
filetype-ai : Filetype ai
filetype-bmp : Filetype bmp
filetype-cs : Filetype cs
filetype-css : Filetype css
filetype-csv : Filetype csv
filetype-doc : Filetype doc
filetype-docx : Filetype docx
filetype-exe : Filetype exe
filetype-gif : Filetype gif
filetype-heic : Filetype heic
filetype-html : Filetype html
filetype-java : Filetype java
filetype-jpg : Filetype jpg
filetype-js : Filetype js
filetype-json : Filetype json
filetype-jsx : Filetype jsx
filetype-key : Filetype key
filetype-m4p : Filetype m4p
filetype-md : Filetype md
filetype-mdx : Filetype mdx
filetype-mov : Filetype mov
filetype-mp3 : Filetype mp3
filetype-mp4 : Filetype mp4
filetype-otf : Filetype otf
filetype-pdf : Filetype pdf
filetype-php : Filetype php
filetype-png : Filetype png
filetype-ppt : Filetype ppt
filetype-pptx : Filetype pptx
filetype-psd : Filetype psd
filetype-py : Filetype py
filetype-raw : Filetype raw
filetype-rb : Filetype rb
filetype-sass : Filetype sass
filetype-scss : Filetype scss
filetype-sh : Filetype sh
filetype-sql : Filetype sql
filetype-svg : Filetype svg
filetype-tiff : Filetype tiff
filetype-tsx : Filetype tsx
filetype-ttf : Filetype ttf
filetype-txt : Filetype txt
filetype-wav : Filetype wav
filetype-woff : Filetype woff
filetype-xls : Filetype xls
filetype-xlsx : Filetype xlsx
filetype-xml : Filetype xml
filetype-yml : Filetype yml
film : Film
filter : Filter
filter-circle : Filter circle
filter-circle-fill : Filter circle fill
filter-left : Filter left
filter-right : Filter right
filter-square : Filter square
filter-square-fill : Filter square fill
fingerprint : Fingerprint
fire : Fire
flag : Flag
flag-fill : Flag fill
floppy : Floppy
floppy-fill : Floppy fill
floppy2 : Floppy2
floppy2-fill : Floppy2 fill
flower1 : Flower1
flower2 : Flower2
flower3 : Flower3
folder : Folder
folder-check : Folder check
folder-fill : Folder fill
folder-minus : Folder minus
folder-plus : Folder plus
folder-symlink : Folder symlink
folder-symlink-fill : Folder symlink fill
folder-x : Folder x
folder2 : Folder2
folder2-open : Folder2 open
fonts : Fonts
forward : Forward
forward-fill : Forward fill
front : Front
fuel-pump : Fuel pump
fuel-pump-diesel : Fuel pump diesel
fuel-pump-diesel-fill : Fuel pump diesel fill
fuel-pump-fill : Fuel pump fill
fullscreen : Fullscreen
fullscreen-exit : Fullscreen exit
funnel : Funnel
funnel-fill : Funnel fill
gear : Gear
gear-fill : Gear fill
gear-wide : Gear wide
gear-wide-connected : Gear wide connected
gem : Gem
gender-ambiguous : Gender ambiguous
gender-female : Gender female
gender-male : Gender male
gender-neuter : Gender neuter
gender-trans : Gender trans
geo : Geo
geo-alt : Geo alt
geo-alt-fill : Geo alt fill
geo-fill : Geo fill
gift : Gift
gift-fill : Gift fill
git : Git
github : Github
gitlab : Gitlab
globe : Globe
globe-americas : Globe americas
globe-asia-australia : Globe asia australia
globe-central-south-asia : Globe central south asia
globe-europe-africa : Globe europe africa
globe2 : Globe2
google : Google
google-play : Google play
gpu-card : Gpu card
graph-down : Graph down
graph-down-arrow : Graph down arrow
graph-up : Graph up
graph-up-arrow : Graph up arrow
grid : Grid
grid-1x2 : Grid 1x2
grid-1x2-fill : Grid 1x2 fill
grid-3x2 : Grid 3x2
grid-3x2-gap : Grid 3x2 gap
grid-3x2-gap-fill : Grid 3x2 gap fill
grid-3x3 : Grid 3x3
grid-3x3-gap : Grid 3x3 gap
grid-3x3-gap-fill : Grid 3x3 gap fill
grid-fill : Grid fill
grip-horizontal : Grip horizontal
grip-vertical : Grip vertical
h-circle : H circle
h-circle-fill : H circle fill
h-square : H square
h-square-fill : H square fill
hammer : Hammer
hand-index : Hand index
hand-index-fill : Hand index fill
hand-index-thumb : Hand index thumb
hand-index-thumb-fill : Hand index thumb fill
hand-thumbs-down : Hand thumbs down
hand-thumbs-down-fill : Hand thumbs down fill
hand-thumbs-up : Hand thumbs up
hand-thumbs-up-fill : Hand thumbs up fill
handbag : Handbag
handbag-fill : Handbag fill
hash : Hash
hdd : Hdd
hdd-fill : Hdd fill
hdd-network : Hdd network
hdd-network-fill : Hdd network fill
hdd-rack : Hdd rack
hdd-rack-fill : Hdd rack fill
hdd-stack : Hdd stack
hdd-stack-fill : Hdd stack fill
hdmi : Hdmi
hdmi-fill : Hdmi fill
headphones : Headphones
headset : Headset
headset-vr : Headset vr
heart : Heart
heart-arrow : Heart arrow
heart-fill : Heart fill
heart-half : Heart half
heart-pulse : Heart pulse
heart-pulse-fill : Heart pulse fill
heartbreak : Heartbreak
heartbreak-fill : Heartbreak fill
hearts : Hearts
heptagon : Heptagon
heptagon-fill : Heptagon fill
heptagon-half : Heptagon half
hexagon : Hexagon
hexagon-fill : Hexagon fill
hexagon-half : Hexagon half
highlighter : Highlighter
highlights : Highlights
hospital : Hospital
hospital-fill : Hospital fill
hourglass : Hourglass
hourglass-bottom : Hourglass bottom
hourglass-split : Hourglass split
hourglass-top : Hourglass top
house : House
house-add : House add
house-add-fill : House add fill
house-check : House check
house-check-fill : House check fill
house-dash : House dash
house-dash-fill : House dash fill
house-door : House door
house-door-fill : House door fill
house-down : House down
house-down-fill : House down fill
house-exclamation : House exclamation
house-exclamation-fill : House exclamation fill
house-fill : House fill
house-gear : House gear
house-gear-fill : House gear fill
house-heart : House heart
house-heart-fill : House heart fill
house-lock : House lock
house-lock-fill : House lock fill
house-slash : House slash
house-slash-fill : House slash fill
house-up : House up
house-up-fill : House up fill
house-x : House x
house-x-fill : House x fill
houses : Houses
houses-fill : Houses fill
hr : Hr
hurricane : Hurricane
hypnotize : Hypnotize
image : Image
image-alt : Image alt
image-fill : Image fill
images : Images
inbox : Inbox
inbox-fill : Inbox fill
inboxes-fill : Inboxes fill
inboxes : Inboxes
incognito : Incognito
indent : Indent
infinity : Infinity
info : Info
info-circle : Info circle
info-circle-fill : Info circle fill
info-lg : Info lg
info-square : Info square
info-square-fill : Info square fill
input-cursor : Input cursor
input-cursor-text : Input cursor text
instagram : Instagram
intersect : Intersect
journal : Journal
journal-album : Journal album
journal-arrow-down : Journal arrow down
journal-arrow-up : Journal arrow up
journal-bookmark : Journal bookmark
journal-bookmark-fill : Journal bookmark fill
journal-check : Journal check
journal-code : Journal code
journal-medical : Journal medical
journal-minus : Journal minus
journal-plus : Journal plus
journal-richtext : Journal richtext
journal-text : Journal text
journal-x : Journal x
journals : Journals
joystick : Joystick
justify : Justify
justify-left : Justify left
justify-right : Justify right
kanban : Kanban
kanban-fill : Kanban fill
key : Key
key-fill : Key fill
keyboard : Keyboard
keyboard-fill : Keyboard fill
ladder : Ladder
lamp : Lamp
lamp-fill : Lamp fill
laptop : Laptop
laptop-fill : Laptop fill
layer-backward : Layer backward
layer-forward : Layer forward
layers : Layers
layers-fill : Layers fill
layers-half : Layers half
layout-sidebar : Layout sidebar
layout-sidebar-inset-reverse : Layout sidebar inset reverse
layout-sidebar-inset : Layout sidebar inset
layout-sidebar-reverse : Layout sidebar reverse
layout-split : Layout split
layout-text-sidebar : Layout text sidebar
layout-text-sidebar-reverse : Layout text sidebar reverse
layout-text-window : Layout text window
layout-text-window-reverse : Layout text window reverse
layout-three-columns : Layout three columns
layout-wtf : Layout wtf
life-preserver : Life preserver
lightbulb : Lightbulb
lightbulb-fill : Lightbulb fill
lightbulb-off : Lightbulb off
lightbulb-off-fill : Lightbulb off fill
lightning : Lightning
lightning-charge : Lightning charge
lightning-charge-fill : Lightning charge fill
lightning-fill : Lightning fill
line : Line
link : Link
link-45deg : Link 45deg
linkedin : Linkedin
list : List
list-check : List check
list-columns : List columns
list-columns-reverse : List columns reverse
list-nested : List nested
list-ol : List ol
list-stars : List stars
list-task : List task
list-ul : List ul
lock : Lock
lock-fill : Lock fill
luggage : Luggage
luggage-fill : Luggage fill
lungs : Lungs
lungs-fill : Lungs fill
magic : Magic
magnet : Magnet
magnet-fill : Magnet fill
mailbox : Mailbox
mailbox-flag : Mailbox flag
mailbox2 : Mailbox2
mailbox2-flag : Mailbox2 flag
map : Map
map-fill : Map fill
markdown : Markdown
markdown-fill : Markdown fill
marker-tip : Marker tip
mask : Mask
mastodon : Mastodon
medium : Medium
megaphone : Megaphone
megaphone-fill : Megaphone fill
memory : Memory
menu-app : Menu app
menu-app-fill : Menu app fill
menu-button : Menu button
menu-button-fill : Menu button fill
menu-button-wide : Menu button wide
menu-button-wide-fill : Menu button wide fill
menu-down : Menu down
menu-up : Menu up
messenger : Messenger
meta : Meta
mic : Mic
mic-fill : Mic fill
mic-mute : Mic mute
mic-mute-fill : Mic mute fill
microsoft : Microsoft
microsoft-teams : Microsoft teams
minecart : Minecart
minecart-loaded : Minecart loaded
modem : Modem
modem-fill : Modem fill
moisture : Moisture
moon : Moon
moon-fill : Moon fill
moon-stars : Moon stars
moon-stars-fill : Moon stars fill
mortarboard : Mortarboard
mortarboard-fill : Mortarboard fill
motherboard : Motherboard
motherboard-fill : Motherboard fill
mouse : Mouse
mouse-fill : Mouse fill
mouse2 : Mouse2
mouse2-fill : Mouse2 fill
mouse3 : Mouse3
mouse3-fill : Mouse3 fill
music-note : Music note
music-note-beamed : Music note beamed
music-note-list : Music note list
music-player : Music player
music-player-fill : Music player fill
newspaper : Newspaper
nintendo-switch : Nintendo switch
node-minus : Node minus
node-minus-fill : Node minus fill
node-plus : Node plus
node-plus-fill : Node plus fill
noise-reduction : Noise reduction
nut : Nut
nut-fill : Nut fill
nvidia : Nvidia
nvme : Nvme
nvme-fill : Nvme fill
octagon : Octagon
octagon-fill : Octagon fill
octagon-half : Octagon half
opencollective : Opencollective
optical-audio : Optical audio
optical-audio-fill : Optical audio fill
option : Option
outlet : Outlet
p-circle : P circle
p-circle-fill : P circle fill
p-square : P square
p-square-fill : P square fill
paint-bucket : Paint bucket
palette : Palette
palette-fill : Palette fill
palette2 : Palette2
paperclip : Paperclip
paragraph : Paragraph
pass : Pass
pass-fill : Pass fill
passport : Passport
passport-fill : Passport fill
patch-check : Patch check
patch-check-fill : Patch check fill
patch-exclamation : Patch exclamation
patch-exclamation-fill : Patch exclamation fill
patch-minus : Patch minus
patch-minus-fill : Patch minus fill
patch-plus : Patch plus
patch-plus-fill : Patch plus fill
patch-question : Patch question
patch-question-fill : Patch question fill
pause : Pause
pause-btn : Pause btn
pause-btn-fill : Pause btn fill
pause-circle : Pause circle
pause-circle-fill : Pause circle fill
pause-fill : Pause fill
paypal : Paypal
pc : Pc
pc-display : Pc display
pc-display-horizontal : Pc display horizontal
pc-horizontal : Pc horizontal
pci-card : Pci card
pci-card-network : Pci card network
pci-card-sound : Pci card sound
peace : Peace
peace-fill : Peace fill
pen : Pen
pen-fill : Pen fill
pencil : Pencil
pencil-fill : Pencil fill
pencil-square : Pencil square
pentagon : Pentagon
pentagon-fill : Pentagon fill
pentagon-half : Pentagon half
people : People
person-circle : Person circle
people-fill : People fill
percent : Percent
person : Person
person-add : Person add
person-arms-up : Person arms up
person-badge : Person badge
person-badge-fill : Person badge fill
person-bounding-box : Person bounding box
person-check : Person check
person-check-fill : Person check fill
person-dash : Person dash
person-dash-fill : Person dash fill
person-down : Person down
person-exclamation : Person exclamation
person-fill : Person fill
person-fill-add : Person fill add
person-fill-check : Person fill check
person-fill-dash : Person fill dash
person-fill-down : Person fill down
person-fill-exclamation : Person fill exclamation
person-fill-gear : Person fill gear
person-fill-lock : Person fill lock
person-fill-slash : Person fill slash
person-fill-up : Person fill up
person-fill-x : Person fill x
person-gear : Person gear
person-heart : Person heart
person-hearts : Person hearts
person-lines-fill : Person lines fill
person-lock : Person lock
person-plus : Person plus
person-plus-fill : Person plus fill
person-raised-hand : Person raised hand
person-rolodex : Person rolodex
person-slash : Person slash
person-square : Person square
person-standing : Person standing
person-standing-dress : Person standing dress
person-up : Person up
person-vcard : Person vcard
person-vcard-fill : Person vcard fill
person-video : Person video
person-video2 : Person video2
person-video3 : Person video3
person-walking : Person walking
person-wheelchair : Person wheelchair
person-workspace : Person workspace
person-x : Person x
person-x-fill : Person x fill
phone : Phone
phone-fill : Phone fill
phone-flip : Phone flip
phone-landscape : Phone landscape
phone-landscape-fill : Phone landscape fill
phone-vibrate : Phone vibrate
phone-vibrate-fill : Phone vibrate fill
pie-chart : Pie chart
pie-chart-fill : Pie chart fill
piggy-bank : Piggy bank
piggy-bank-fill : Piggy bank fill
pin : Pin
pin-angle : Pin angle
pin-angle-fill : Pin angle fill
pin-fill : Pin fill
pin-map : Pin map
pin-map-fill : Pin map fill
pinterest : Pinterest
pip : Pip
pip-fill : Pip fill
play : Play
play-btn : Play btn
play-btn-fill : Play btn fill
play-circle : Play circle
play-circle-fill : Play circle fill
play-fill : Play fill
playstation : Playstation
plug : Plug
plug-fill : Plug fill
plugin : Plugin
plus : Plus
plus-circle : Plus circle
plus-circle-dotted : Plus circle dotted
plus-circle-fill : Plus circle fill
plus-lg : Plus lg
plus-slash-minus : Plus slash minus
plus-square : Plus square
plus-square-dotted : Plus square dotted
plus-square-fill : Plus square fill
postage : Postage
postage-fill : Postage fill
postage-heart : Postage heart
postage-heart-fill : Postage heart fill
postcard : Postcard
postcard-fill : Postcard fill
postcard-heart : Postcard heart
postcard-heart-fill : Postcard heart fill
power : Power
prescription : Prescription
prescription2 : Prescription2
printer : Printer
printer-fill : Printer fill
projector : Projector
projector-fill : Projector fill
puzzle : Puzzle
puzzle-fill : Puzzle fill
qr-code : Qr code
qr-code-scan : Qr code scan
question : Question
question-circle : Question circle
question-diamond : Question diamond
question-diamond-fill : Question diamond fill
question-circle-fill : Question circle fill
question-lg : Question lg
question-octagon : Question octagon
question-octagon-fill : Question octagon fill
question-square : Question square
question-square-fill : Question square fill
quora : Quora
quote : Quote
r-circle : R circle
r-circle-fill : R circle fill
r-square : R square
r-square-fill : R square fill
radar : Radar
radioactive : Radioactive
rainbow : Rainbow
receipt : Receipt
receipt-cutoff : Receipt cutoff
reception-0 : Reception 0
reception-1 : Reception 1
reception-2 : Reception 2
reception-3 : Reception 3
reception-4 : Reception 4
record : Record
record-btn : Record btn
record-btn-fill : Record btn fill
record-circle : Record circle
record-circle-fill : Record circle fill
record-fill : Record fill
record2 : Record2
record2-fill : Record2 fill
recycle : Recycle
reddit : Reddit
regex : Regex
repeat : Repeat
repeat-1 : Repeat 1
reply : Reply
reply-all : Reply all
reply-all-fill : Reply all fill
reply-fill : Reply fill
rewind : Rewind
rewind-btn : Rewind btn
rewind-btn-fill : Rewind btn fill
rewind-circle : Rewind circle
rewind-circle-fill : Rewind circle fill
rewind-fill : Rewind fill
robot : Robot
rocket : Rocket
rocket-fill : Rocket fill
rocket-takeoff : Rocket takeoff
rocket-takeoff-fill : Rocket takeoff fill
router : Router
router-fill : Router fill
rss : Rss
rss-fill : Rss fill
rulers : Rulers
safe : Safe
safe-fill : Safe fill
safe2 : Safe2
safe2-fill : Safe2 fill
save : Save
save-fill : Save fill
save2 : Save2
save2-fill : Save2 fill
scissors : Scissors
scooter : Scooter
screwdriver : Screwdriver
sd-card : Sd card
sd-card-fill : Sd card fill
search : Search
search-heart : Search heart
search-heart-fill : Search heart fill
segmented-nav : Segmented nav
send : Send
send-arrow-down : Send arrow down
send-arrow-down-fill : Send arrow down fill
send-arrow-up : Send arrow up
send-arrow-up-fill : Send arrow up fill
send-check : Send check
send-check-fill : Send check fill
send-dash : Send dash
send-dash-fill : Send dash fill
send-exclamation : Send exclamation
send-exclamation-fill : Send exclamation fill
send-fill : Send fill
send-plus : Send plus
send-plus-fill : Send plus fill
send-slash : Send slash
send-slash-fill : Send slash fill
send-x : Send x
send-x-fill : Send x fill
server : Server
shadows : Shadows
share : Share
share-fill : Share fill
shield : Shield
shield-check : Shield check
shield-exclamation : Shield exclamation
shield-fill : Shield fill
shield-fill-check : Shield fill check
shield-fill-exclamation : Shield fill exclamation
shield-fill-minus : Shield fill minus
shield-fill-plus : Shield fill plus
shield-fill-x : Shield fill x
shield-lock : Shield lock
shield-lock-fill : Shield lock fill
shield-minus : Shield minus
shield-plus : Shield plus
shield-shaded : Shield shaded
shield-slash : Shield slash
shield-slash-fill : Shield slash fill
shield-x : Shield x
shift : Shift
shift-fill : Shift fill
shop : Shop
shop-window : Shop window
shuffle : Shuffle
sign-dead-end : Sign dead end
sign-dead-end-fill : Sign dead end fill
sign-do-not-enter : Sign do not enter
sign-do-not-enter-fill : Sign do not enter fill
sign-intersection : Sign intersection
sign-intersection-fill : Sign intersection fill
sign-intersection-side : Sign intersection side
sign-intersection-side-fill : Sign intersection side fill
sign-intersection-t : Sign intersection t
sign-intersection-t-fill : Sign intersection t fill
sign-intersection-y : Sign intersection y
sign-intersection-y-fill : Sign intersection y fill
sign-merge-left : Sign merge left
sign-merge-left-fill : Sign merge left fill
sign-merge-right : Sign merge right
sign-merge-right-fill : Sign merge right fill
sign-no-left-turn : Sign no left turn
sign-no-left-turn-fill : Sign no left turn fill
sign-no-parking : Sign no parking
sign-no-parking-fill : Sign no parking fill
sign-no-right-turn : Sign no right turn
sign-no-right-turn-fill : Sign no right turn fill
sign-railroad : Sign railroad
sign-railroad-fill : Sign railroad fill
sign-stop : Sign stop
sign-stop-fill : Sign stop fill
sign-stop-lights : Sign stop lights
sign-stop-lights-fill : Sign stop lights fill
sign-turn-left : Sign turn left
sign-turn-left-fill : Sign turn left fill
sign-turn-right : Sign turn right
sign-turn-right-fill : Sign turn right fill
sign-turn-slight-left : Sign turn slight left
sign-turn-slight-left-fill : Sign turn slight left fill
sign-turn-slight-right : Sign turn slight right
sign-turn-slight-right-fill : Sign turn slight right fill
sign-yield : Sign yield
sign-yield-fill : Sign yield fill
signal : Signal
signpost : Signpost
signpost-2 : Signpost 2
signpost-2-fill : Signpost 2 fill
signpost-fill : Signpost fill
signpost-split : Signpost split
signpost-split-fill : Signpost split fill
sim : Sim
sim-fill : Sim fill
sim-slash : Sim slash
sim-slash-fill : Sim slash fill
sina-weibo : Sina weibo
skip-backward : Skip backward
skip-backward-btn : Skip backward btn
skip-backward-btn-fill : Skip backward btn fill
skip-backward-circle : Skip backward circle
skip-backward-circle-fill : Skip backward circle fill
skip-backward-fill : Skip backward fill
skip-end : Skip end
skip-end-btn : Skip end btn
skip-end-btn-fill : Skip end btn fill
skip-end-circle : Skip end circle
skip-end-circle-fill : Skip end circle fill
skip-end-fill : Skip end fill
skip-forward : Skip forward
skip-forward-btn : Skip forward btn
skip-forward-btn-fill : Skip forward btn fill
skip-forward-circle : Skip forward circle
skip-forward-circle-fill : Skip forward circle fill
skip-forward-fill : Skip forward fill
skip-start : Skip start
skip-start-btn : Skip start btn
skip-start-btn-fill : Skip start btn fill
skip-start-circle : Skip start circle
skip-start-circle-fill : Skip start circle fill
skip-start-fill : Skip start fill
skype : Skype
slack : Slack
slash : Slash
slash-circle-fill : Slash circle fill
slash-lg : Slash lg
slash-square : Slash square
slash-square-fill : Slash square fill
sliders : Sliders
sliders2 : Sliders2
sliders2-vertical : Sliders2 vertical
smartwatch : Smartwatch
snapchat : Snapchat
snow : Snow
snow2 : Snow2
snow3 : Snow3
sort-alpha-down : Sort alpha down
sort-alpha-down-alt : Sort alpha down alt
sort-alpha-up : Sort alpha up
sort-alpha-up-alt : Sort alpha up alt
sort-down : Sort down
sort-down-alt : Sort down alt
sort-numeric-down : Sort numeric down
sort-numeric-down-alt : Sort numeric down alt
sort-numeric-up : Sort numeric up
sort-numeric-up-alt : Sort numeric up alt
sort-up : Sort up
sort-up-alt : Sort up alt
soundwave : Soundwave
sourceforge : Sourceforge
speaker : Speaker
speaker-fill : Speaker fill
speedometer : Speedometer
speedometer2 : Speedometer2
spellcheck : Spellcheck
spotify : Spotify
square : Square
square-fill : Square fill
square-half : Square half
stack : Stack
stack-overflow : Stack overflow
star : Star
star-fill : Star fill
star-half : Star half
stars : Stars
steam : Steam
stickies : Stickies
stickies-fill : Stickies fill
sticky : Sticky
sticky-fill : Sticky fill
stop : Stop
stop-btn : Stop btn
stop-btn-fill : Stop btn fill
stop-circle : Stop circle
stop-circle-fill : Stop circle fill
stop-fill : Stop fill
stoplights : Stoplights
stoplights-fill : Stoplights fill
stopwatch : Stopwatch
stopwatch-fill : Stopwatch fill
strava : Strava
stripe : Stripe
subscript : Subscript
substack : Substack
subtract : Subtract
suit-club : Suit club
suit-club-fill : Suit club fill
suit-diamond : Suit diamond
suit-diamond-fill : Suit diamond fill
suit-heart : Suit heart
suit-heart-fill : Suit heart fill
suit-spade : Suit spade
suit-spade-fill : Suit spade fill
suitcase : Suitcase
suitcase-fill : Suitcase fill
suitcase-lg : Suitcase lg
suitcase-lg-fill : Suitcase lg fill
suitcase2 : Suitcase2
suitcase2-fill : Suitcase2 fill
sun : Sun
sun-fill : Sun fill
sunglasses : Sunglasses
sunrise : Sunrise
sunrise-fill : Sunrise fill
sunset : Sunset
sunset-fill : Sunset fill
superscript : Superscript
symmetry-horizontal : Symmetry horizontal
symmetry-vertical : Symmetry vertical
table : Table
tablet : Tablet
tablet-fill : Tablet fill
tablet-landscape : Tablet landscape
tablet-landscape-fill : Tablet landscape fill
tag : Tag
tag-fill : Tag fill
tags : Tags
tags-fill : Tags fill
taxi-front : Taxi front
taxi-front-fill : Taxi front fill
telegram : Telegram
telephone : Telephone
telephone-fill : Telephone fill
telephone-forward : Telephone forward
telephone-forward-fill : Telephone forward fill
telephone-inbound : Telephone inbound
telephone-inbound-fill : Telephone inbound fill
telephone-minus : Telephone minus
telephone-minus-fill : Telephone minus fill
telephone-outbound : Telephone outbound
telephone-outbound-fill : Telephone outbound fill
telephone-plus : Telephone plus
telephone-plus-fill : Telephone plus fill
telephone-x : Telephone x
telephone-x-fill : Telephone x fill
tencent-qq : Tencent qq
terminal : Terminal
terminal-dash : Terminal dash
terminal-fill : Terminal fill
terminal-plus : Terminal plus
terminal-split : Terminal split
terminal-x : Terminal x
text-center : Text center
text-indent-left : Text indent left
text-indent-right : Text indent right
text-left : Text left
text-paragraph : Text paragraph
text-right : Text right
text-wrap : Text wrap
textarea : Textarea
textarea-resize : Textarea resize
textarea-t : Textarea t
thermometer : Thermometer
thermometer-half : Thermometer half
thermometer-high : Thermometer high
thermometer-low : Thermometer low
thermometer-snow : Thermometer snow
thermometer-sun : Thermometer sun
threads : Threads
threads-fill : Threads fill
three-dots : Three dots
three-dots-vertical : Three dots vertical
thunderbolt : Thunderbolt
thunderbolt-fill : Thunderbolt fill
ticket : Ticket
ticket-detailed : Ticket detailed
ticket-detailed-fill : Ticket detailed fill
ticket-fill : Ticket fill
ticket-perforated : Ticket perforated
ticket-perforated-fill : Ticket perforated fill
tiktok : Tiktok
toggle-off : Toggle off
toggle-on : Toggle on
toggle2-off : Toggle2 off
toggle2-on : Toggle2 on
toggles : Toggles
toggles2 : Toggles2
tools : Tools
tornado : Tornado
train-freight-front : Train freight front
train-freight-front-fill : Train freight front fill
train-front : Train front
train-front-fill : Train front fill
train-lightrail-front : Train lightrail front
train-lightrail-front-fill : Train lightrail front fill
translate : Translate
transparency : Transparency
trash : Trash
trash-fill : Trash fill
trash2 : Trash2
trash2-fill : Trash2 fill
trash3 : Trash3
trash3-fill : Trash3 fill
tree : Tree
tree-fill : Tree fill
trello : Trello
triangle : Triangle
triangle-fill : Triangle fill
triangle-half : Triangle half
trophy : Trophy
trophy-fill : Trophy fill
tropical-storm : Tropical storm
truck : Truck
truck-flatbed : Truck flatbed
truck-front : Truck front
truck-front-fill : Truck front fill
tsunami : Tsunami
tv : Tv
tv-fill : Tv fill
twitch : Twitch
twitter : Twitter
twitter-x : Twitter x
type : Type
type-bold : Type bold
type-h1 : Type h1
type-h2 : Type h2
type-h3 : Type h3
type-h4 : Type h4
type-h5 : Type h5
type-h6 : Type h6
type-italic : Type italic
type-strikethrough : Type strikethrough
type-underline : Type underline
ubuntu : Ubuntu
ui-checks : Ui checks
ui-checks-grid : Ui checks grid
ui-radios : Ui radios
ui-radios-grid : Ui radios grid
umbrella : Umbrella
umbrella-fill : Umbrella fill
unindent : Unindent
union : Union
unity : Unity
universal-access : Universal access
universal-access-circle : Universal access circle
unlock : Unlock
unlock-fill : Unlock fill
upc : Upc
upc-scan : Upc scan
upload : Upload
usb : Usb
usb-c : Usb c
usb-c-fill : Usb c fill
usb-drive : Usb drive
usb-drive-fill : Usb drive fill
usb-fill : Usb fill
usb-micro : Usb micro
usb-micro-fill : Usb micro fill
usb-mini : Usb mini
usb-mini-fill : Usb mini fill
usb-plug : Usb plug
usb-plug-fill : Usb plug fill
usb-symbol : Usb symbol
valentine : Valentine
valentine2 : Valentine2
vector-pen : Vector pen
view-list : View list
view-stacked : View stacked
vignette : Vignette
vimeo : Vimeo
vinyl : Vinyl
vinyl-fill : Vinyl fill
virus : Virus
virus2 : Virus2
voicemail : Voicemail
volume-down : Volume down
volume-down-fill : Volume down fill
volume-mute : Volume mute
volume-mute-fill : Volume mute fill
volume-off : Volume off
volume-off-fill : Volume off fill
volume-up : Volume up
volume-up-fill : Volume up fill
vr : Vr
wallet : Wallet
wallet-fill : Wallet fill
wallet2 : Wallet2
watch : Watch
water : Water
webcam : Webcam
webcam-fill : Webcam fill
wechat : Wechat
whatsapp : Whatsapp
wifi : Wifi
wifi-1 : Wifi 1
wifi-2 : Wifi 2
wifi-off : Wifi off
wikipedia : Wikipedia
wind : Wind
window : Window
window-dash : Window dash
window-desktop : Window desktop
window-dock : Window dock
window-fullscreen : Window fullscreen
window-plus : Window plus
window-sidebar : Window sidebar
window-split : Window split
window-stack : Window stack
window-x : Window x
windows : Windows
wordpress : Wordpress
wrench : Wrench
wrench-adjustable : Wrench adjustable
wrench-adjustable-circle : Wrench adjustable circle
wrench-adjustable-circle-fill : Wrench adjustable circle fill
x : X
x-circle : X circle
x-circle-fill : X circle fill
x-diamond : X diamond
x-diamond-fill : X diamond fill
x-lg : X lg
x-octagon : X octagon
x-octagon-fill : X octagon fill
x-square : X square
x-square-fill : X square fill
xbox : Xbox
yelp : Yelp
yin-yang : Yin yang
youtube : Youtube
zoom-in : Zoom in
zoom-out : Zoom out
	`

	rows := strings.Split(strings.TrimSpace(icons), "\n")
	list := []struct {
		Name string
		Icon string
	}{}

	for _, row := range rows {
		parts := strings.Split(row, " : ")
		list = append(list, struct {
			Name string
			Icon string
		}{
			Name: parts[1],
			Icon: `bi bi-` + parts[0],
		})
	}

	return list
}
