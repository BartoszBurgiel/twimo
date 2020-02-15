# TWIMO

...**T**he **W**orld **I**s **M**y **O**ffice

## The challenge

_"Remote Work"_ is one of the biggest buzzwords in the 21st century and that's mainly because it's f\*cking awesome. However, remote work brings up some caveats, some requirements that one must fulfill in order to enjoy the ability to work from whatever place he likes. These requirements are...

- power supply
- decent wifi speed
- comfortable location

While it's relatively easy to achieve this at home, a lot of people like to keep their private life separated from work - which brings up the demand for an external location that provides the three points mentioned above.

If you live in a digital hub like Berlin or London, you could simply sign up for a co-working space that is specialized in serving remote workers as a daily business. _But what about the people who live in more rural, less digitalized areas?_ Well, they rely heavily on locations such as coffeeshops or bars that also fulfill the three requirements mentioned earlier.

Coffeeshops and bars, however, do not always provide fast & free wifi or power supplies for their customers. So one needs an easy and fast solution to evaluate different locations. **And this is what _twimo_ does for them.**

## The solution

With some help of public APIs like _Google Maps_ or _tripadvisor_, **twimo** is able to retrieve information about potential locations that then can be evaluated by the users.

Through that principle, **TWIMO** gets better and better with every new user.

## The implementation

#### FrontEnd

- **React Native** for cross-plattform compatibility
  - React Navigate for navigation from screen to screen
  - ExpoIcons for... well, icons.
- **Expo** for testing

#### BackEnd

- **Go** as server-side language to connect DB & Frontend
- **REST API** build with go

#### DataStore

- **mySQL** to store user, location and rating data
