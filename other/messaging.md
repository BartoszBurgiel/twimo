# Messaging specs 

here you can find the formats and types of all messages 
that will be send and recieved via the websocket on all (affected) routes  

# /listscreen 

* json with the locations (see on /backend/server/assets/login.html)
* `{"Name":"Betties Place","Coords":{"X":0,"Y":0},"Desc":"Betties Places description...","Webpage":"","Comments":null,"Users":0,"Rating":2.5590277777777777,"Price":4,"Features":{"Coffee":false,"Wifi":false,"Power":false,"Music":false,"Food":false},"ID":"6dc2eb46-5fb4-44a6-8957-5a432885860d"}`

# /login 

## front end sends

* **user name** and **password** as singe messages **IN THIS VERY ORDER** 

## backend responds 

* single message confirming the login -> **bool** (0/1)
* (if login successful) json sending the **user's username** and the **hash key** 
    * these must be stored in phone's local memory 
    * user's username to search in the database and hash hey to verify
    * `{"Username":"Louis Beul","Key":".[GxU3fb.r55LbsF*%mmbgDbBubFz\u003c_;{c84,o2kmbVkb7XcnbWjbw0g2GTn1?nbokbo0GTl04ZrjbY1^SS3T4p1m3CTi1Wibt0oZoYf4c4i1i4l2xWTlbf4u1uWuXckb\u003cmbinbdlb)Sjjbi3CXo0(SZZpZCXe4DTWjb2YjkbXib\u0026SvW84LjbzlbTjb/1F2F2U0H0Anb/W-.IGhU[f\u003c.b5;LMsL*.md*Ur^b^H=bizlbfaH4}pgb%oGbf0=ribanLbi2#4rV/3\u003esRb(nlbY2yVu212xmpb04mV46v7;3255VW3BmIbs2Y1S0D6a6T316\u00264nY9pUb46P3EY2Z,nHbNrHb;rebVo8btV{mFb356Zu2EV4211CZO6IVynEbz1Rn?b^m7bDVMYx6,nFb0pzb8nYbY4n4Y5C3A3*subUZ\u003cSZD{tR^EXVTrDx\u003eP:EUimYZK{P?E1+[H\u0026z+/,9YnOoGrtOkmA*OsGL75Y.sfGULG_.-yByHvvmFfHtrrHrzoaipbob0rbnb24V3sbnb2V22mb4V6735V3mb21066364Ypb63YZnbrbrbobVmb5Z2V21Z6Vnb1nbmbVY6nbpbnb44533sbZ"}`


